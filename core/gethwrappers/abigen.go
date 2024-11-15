package gethwrappers

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	gethParams "github.com/ethereum/go-ethereum/params"
	"golang.org/x/tools/go/ast/astutil"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils"
)

const headerComment = `// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

`

// AbigenArgs is the arguments to the abigen executable. E.g., Bin is the -bin
// arg.
type AbigenArgs struct {
	Bin, ABI, Out, Type, Pkg, ZkBinPath string
}

var zkDeployCode = (`
type CustomTransaction struct {
	*types.Transaction
	CustomHash common.Hash
}

func (tx *CustomTransaction) Hash() common.Hash {
	return tx.CustomHash
}

func ConvertToTransaction(resp zktypes.TransactionResponse) *CustomTransaction {
	dtx := &types.DynamicFeeTx{
		ChainID:   resp.ChainID.ToInt(),
		Nonce:     uint64(resp.Nonce),
		GasTipCap: resp.MaxPriorityFeePerGas.ToInt(),
		GasFeeCap: resp.MaxFeePerGas.ToInt(),
		To:        &resp.To,
		Value:     resp.Value.ToInt(),
		Data:      resp.Data,
		Gas:       uint64(resp.Gas),
	}

	// Create the transaction
	tx := types.NewTx(dtx)
	customTransaction := CustomTransaction{Transaction: tx, CustomHash: resp.Hash}
	return &customTransaction
}
// this should generated.CustomTransaction
func DeployZkSync%s(auth *bind.TransactOpts, backend bind.ContractBackend, params ...interface{}) (common.Address, *CustomTransaction, *%s, error) {
	client, ok := backend.(*ethclient.Client)
	if !ok {
		return common.Address{}, nil, nil, errors.New("backend is not an ethclient")
	}
	fmt.Println("Deploying zksync contract")
	zksyncClient := zkSyncClient.NewClient(client.Client())
	fmt.Println("getting wallet")
	wallet := auth.Context.Value("wallet").(*zkSyncAccounts.Wallet)
	fmt.Println("got wallet")
	fmt.Println("getting bytes")
	decodedBytes := common.FromHex(%sZkBin)
	fmt.Println("deploying")
	%sAbi, err := %sMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	constructor, err := %sAbi.Pack("", params...)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	hash, err := wallet.DeployWithCreate(nil, zkSyncAccounts.CreateTransaction{
		Bytecode: decodedBytes,
		Calldata: constructor,
	})
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	fmt.Println("hash of tx", hash)
	receipt, err := zksyncClient.WaitMined(context.Background(), hash)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	tx, _, err := zksyncClient.TransactionByHash(context.Background(), hash)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	fmt.Println("tx hash", tx.Hash)
	// this should generated.ConvertToTransaction
	ethTx := ConvertToTransaction(*tx)
	address := receipt.ContractAddress

	parsed, err := %sMetaData.GetAbi()
	contractBind := bind.NewBoundContract(address, *parsed, backend, backend, backend)

	contractReturn := &%s{address: address, abi: *parsed, %sCaller: %sCaller{contract: contractBind}, %sTransactor: %sTransactor{contract: contractBind}, %sFilterer: %sFilterer{contract: contractBind}}

	return address, ethTx, contractReturn, err
}
`)

// Abigen calls Abigen  with the given arguments
//
// It might seem like a shame, to shell out to another golang program like
// this, but the abigen executable is the stable public interface to the
// geth contract-wrapper machinery.
//
// Check whether native abigen is installed, and has correct version
func Abigen(a AbigenArgs) {
	var versionResponse bytes.Buffer
	abigenExecutablePath := filepath.Join(GetProjectRoot(), "tools/bin/abigen")
	abigenVersionCheck := exec.Command(abigenExecutablePath, "--version")
	abigenVersionCheck.Stdout = &versionResponse
	if err := abigenVersionCheck.Run(); err != nil {
		Exit("no native abigen; you must install it (`make abigen` in the "+
			"chainlink root dir)", err)
	}
	version := string(regexp.MustCompile(`[0-9]+\.[0-9]+\.[0-9]+`).Find(
		versionResponse.Bytes()))
	if version != gethParams.Version {
		Exit(fmt.Sprintf("wrong version (%s) of abigen; install the correct one "+
			"(%s) with `make abigen` in the chainlink root dir", version,
			gethParams.Version),
			nil)
	}
	args := []string{
		"-abi", a.ABI,
		"-out", a.Out,
		"-type", a.Type,
		"-pkg", a.Pkg,
	}
	if a.Bin != "-" {
		args = append(args, "-bin", a.Bin)
	}
	buildCommand := exec.Command(abigenExecutablePath, args...)
	var buildResponse bytes.Buffer
	buildCommand.Stderr = &buildResponse
	if err := buildCommand.Run(); err != nil {
		Exit("failure while building "+a.Pkg+" wrapper, stderr: "+buildResponse.String(), err)
	}

	zkbytes, err := os.ReadFile(a.ZkBinPath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	zkHexString := string(zkbytes)

	ImproveAbigenOutput(a.Out, a.ABI, zkHexString)
}

func ImproveAbigenOutput(path string, abiPath string, zkHexString string) {
	abiBytes, err := os.ReadFile(abiPath)
	if err != nil {
		Exit("Error while improving abigen output", err)
	}
	abi, err := abi.JSON(strings.NewReader(string(abiBytes)))
	if err != nil {
		Exit("Error while improving abigen output", err)
	}

	bs, err := os.ReadFile(path)
	if err != nil {
		Exit("Error while improving abigen output", err)
	}

	fset, fileNode := parseFile(bs)
	logNames := getLogNames(fileNode)
	if len(logNames) > 0 {
		astutil.AddImport(fset, fileNode, "fmt")
		astutil.AddImport(fset, fileNode, "github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated")
	}
	contractName := getContractName(fileNode)
	fileNode = addContractStructFields(contractName, fileNode)

	// zksync
	fileNode = addZKSyncImports(fset, fileNode)
	fileNode = addZKSyncBin(fileNode, contractName, zkHexString)
	fileNode = addZKSyncLogic(contractName, fset, fileNode)

	fileNode = replaceAnonymousStructs(contractName, fileNode)
	bs = generateCode(fset, fileNode)
	bs = writeAdditionalMethods(contractName, logNames, abi, bs)

	// zksync
	zkSyncDeployFunction := strings.ReplaceAll(zkDeployCode, "%s", contractName)
	bs = append(bs, []byte(fmt.Sprintf("%s\n", zkSyncDeployFunction))...)

	err = os.WriteFile(path, bs, 0600)
	if err != nil {
		Exit("Error while writing improved abigen source", err)
	}

	fset, fileNode = parseFile(bs)
	fileNode = writeInterface(contractName, fileNode)
	bs = generateCode(fset, fileNode)
	bs = addHeader(bs)

	err = os.WriteFile(path, bs, 0600)
	if err != nil {
		Exit("Error while writing improved abigen source", err)
	}
}

func parseFile(bs []byte) (*token.FileSet, *ast.File) {
	fset := token.NewFileSet()
	fileNode, err := parser.ParseFile(fset, "", string(bs), parser.AllErrors)
	if err != nil {
		Exit("Error while improving abigen output", err)
	}
	return fset, fileNode
}

func generateCode(fset *token.FileSet, fileNode *ast.File) []byte {
	var buf bytes.Buffer
	err := format.Node(&buf, fset, fileNode)
	if err != nil {
		Exit("Error while writing improved abigen source", err)
	}
	return buf.Bytes()
}

func getContractName(fileNode *ast.File) string {
	// Search for the ABI const e.g. VRFCoordinatorV2ABI = "0x..."
	var contractName string
	astutil.Apply(fileNode, func(cursor *astutil.Cursor) bool {
		x, is := cursor.Node().(*ast.ValueSpec)
		if !is {
			return true
		}
		if len(x.Names) > 0 {
			for _, n := range x.Names {
				if len(n.Name) < 3 {
					return true
				}
				if n.Name[len(n.Name)-3:] != "ABI" {
					return true
				}
				contractName = n.Name[:len(n.Name)-3]
			}
		}
		return false
	}, nil)
	return contractName
}

// Add the `.address` and `.abi` fields to the contract struct.
func addContractStructFields(contractName string, fileNode *ast.File) *ast.File {
	fileNode = addContractStructFieldsToStruct(contractName, fileNode)
	fileNode = addContractStructFieldsToConstructor(contractName, fileNode)
	fileNode = addContractStructFieldsToDeployMethod(contractName, fileNode)
	return fileNode
}

// Add the fields to the contract struct.
func addContractStructFieldsToStruct(contractName string, fileNode *ast.File) *ast.File {
	return astutil.Apply(fileNode, func(cursor *astutil.Cursor) bool {
		x, is := cursor.Node().(*ast.StructType)
		if !is {
			return true
		}
		theType, is := cursor.Parent().(*ast.TypeSpec)
		if !is {
			return false
		} else if theType.Name.Name != contractName {
			return false
		}

		addrField := &ast.Field{
			Names: []*ast.Ident{ast.NewIdent("address")},
			Type: &ast.SelectorExpr{
				X:   ast.NewIdent("common"),
				Sel: ast.NewIdent("Address"),
			},
		}

		abiField := &ast.Field{
			Names: []*ast.Ident{ast.NewIdent("abi")},
			Type: &ast.SelectorExpr{
				X:   ast.NewIdent("abi"),
				Sel: ast.NewIdent("ABI"),
			},
		}
		x.Fields.List = append([]*ast.Field{addrField, abiField}, x.Fields.List...)
		return false
	}, nil).(*ast.File)
}

// Add the fields to the return value of the constructor.
func addContractStructFieldsToConstructor(contractName string, fileNode *ast.File) *ast.File {
	return astutil.Apply(fileNode, func(cursor *astutil.Cursor) bool {
		x, is := cursor.Node().(*ast.FuncDecl)
		if !is {
			return true
		} else if x.Name.Name != "New"+contractName {
			return false
		}

		for _, stmt := range x.Body.List {
			returnStmt, is := stmt.(*ast.ReturnStmt)
			if !is {
				continue
			}
			lit, is := returnStmt.Results[0].(*ast.UnaryExpr).X.(*ast.CompositeLit)
			if !is {
				continue
			}
			addressExpr := &ast.KeyValueExpr{
				Key:   ast.NewIdent("address"),
				Value: ast.NewIdent("address"),
			}
			abiExpr := &ast.KeyValueExpr{
				Key:   ast.NewIdent("abi"),
				Value: ast.NewIdent("abi"),
			}
			lit.Elts = append([]ast.Expr{addressExpr, abiExpr}, lit.Elts...)
		}

		parseABIStmt := &ast.AssignStmt{
			Lhs: []ast.Expr{ast.NewIdent("abi"), ast.NewIdent("err")},
			Tok: token.DEFINE,
			Rhs: []ast.Expr{
				&ast.CallExpr{
					Fun: &ast.SelectorExpr{
						X:   ast.NewIdent("abi"),
						Sel: ast.NewIdent("JSON"),
					},
					Args: []ast.Expr{
						&ast.CallExpr{
							Fun: &ast.SelectorExpr{
								X:   ast.NewIdent("strings"),
								Sel: ast.NewIdent("NewReader"),
							},
							Args: []ast.Expr{ast.NewIdent(contractName + "ABI")},
						},
					},
				},
			},
		}
		checkParseABIErrStmt := &ast.IfStmt{
			Cond: &ast.BinaryExpr{
				X:  ast.NewIdent("err"),
				Op: token.NEQ,
				Y:  ast.NewIdent("nil"),
			},
			Body: &ast.BlockStmt{
				List: []ast.Stmt{
					&ast.ReturnStmt{
						Results: []ast.Expr{ast.NewIdent("nil"), ast.NewIdent("err")},
					},
				},
			},
		}

		x.Body.List = append([]ast.Stmt{parseABIStmt, checkParseABIErrStmt}, x.Body.List...)
		return false
	}, nil).(*ast.File)
}

// Add the fields to the returned struct in the 'Deploy<contractName>' method.
func addContractStructFieldsToDeployMethod(contractName string, fileNode *ast.File) *ast.File {
	return astutil.Apply(fileNode, func(cursor *astutil.Cursor) bool {
		x, is := cursor.Node().(*ast.FuncDecl)
		if !is {
			return true
		} else if x.Name.Name != "Deploy"+contractName {
			return false
		}

		for _, stmt := range x.Body.List {
			returnStmt, is := stmt.(*ast.ReturnStmt)
			if !is {
				continue
			}
			if len(returnStmt.Results) < 3 {
				continue
			}
			rs, is := returnStmt.Results[2].(*ast.UnaryExpr)
			if !is {
				return true
			}
			lit, is := rs.X.(*ast.CompositeLit)
			if !is {
				continue
			}
			addressExpr := &ast.KeyValueExpr{
				Key:   ast.NewIdent("address"),
				Value: ast.NewIdent("address"),
			}
			abiExpr := &ast.KeyValueExpr{
				Key:   ast.NewIdent("abi"),
				Value: ast.NewIdent("*parsed"),
			}
			lit.Elts = append([]ast.Expr{addressExpr, abiExpr}, lit.Elts...)
		}
		return false
	}, nil).(*ast.File)
}

func updateReturnStatements(funcDecl *ast.FuncDecl, params []ast.Expr) {
	// Use ast.Inspect to traverse the entire function body
	ast.Inspect(funcDecl.Body, func(n ast.Node) bool {
		if returnStmt, ok := n.(*ast.ReturnStmt); ok {
			fmt.Println("Found return statement")
			if len(returnStmt.Results) > 0 {
				if callExpr, ok := returnStmt.Results[0].(*ast.CallExpr); ok {
					// Update the arguments of the DeployZkSync function call
					fmt.Println("Updating DeployZkSync function call arguments")
					callExpr.Args = params
				}
			}
		}
		return true
	})
}

// Add the fields to the returned struct in the 'Deploy<contractName>' method.
func addZKSyncLogic(contractName string, fset *token.FileSet, fileNode *ast.File) *ast.File {
	return astutil.Apply(fileNode, func(cursor *astutil.Cursor) bool {
		x, is := cursor.Node().(*ast.FuncDecl)
		if !is {
			return true
		} else if x.Name.Name != "Deploy"+contractName {
			return false
		}

		// Extract the parameters from the existing function x
		var params []ast.Expr
		for _, param := range x.Type.Params.List {
			for _, name := range param.Names {
				params = append(params, ast.NewIdent(name.Name))
			}
		}

		newCode := fmt.Sprintf(`
		package main
		func tempFunc() {
			client, ok := backend.(*ethclient.Client)
			if !ok {
				return common.Address{}, nil, nil, errors.New("backend is not an ethclient")
			}
			chainId, err := client.ChainID(context.Background())
			if err != nil {
				return common.Address{}, nil, nil, err
			}
			switch chainId.Uint64() {
			// this is not sustainable, but it's a quick fix for now
			case 324, 280, 300:
				return DeployZkSync%s(auth, backend)
			}
		}
		`, contractName)
		// Parse the new code snippet as a temporary function to get the statements
		tempNode, err := parser.ParseFile(fset, "", newCode, parser.ParseComments)
		if err != nil {
			panic(err)
		}

		for _, decl := range tempNode.Decls {
			if funcDecl, ok := decl.(*ast.FuncDecl); ok && funcDecl.Name.Name == "tempFunc" {
				fmt.Println("Found function:", funcDecl.Name.Name)
				// Update return statements
				updateReturnStatements(funcDecl, params)
			}
		}

		// Extract the body of the temporary function as statements
		var newStatements []ast.Stmt
		for _, decl := range tempNode.Decls {
			if funcDecl, ok := decl.(*ast.FuncDecl); ok && funcDecl.Name.Name == "tempFunc" {

				newStatements = funcDecl.Body.List
				break
			}
		}
		x.Body.List = append(newStatements, x.Body.List...)

		// zksync
		// x.Type.Results.List[1].Type = &ast.StarExpr{
		// 	X: &ast.SelectorExpr{
		// 		X:   &ast.Ident{Name: "generated"},
		// 		Sel: &ast.Ident{Name: "CustomTransaction"},
		// 	},
		// }
		x.Type.Results.List[1].Type = &ast.StarExpr{
			X: &ast.Ident{Name: "CustomTransaction"},
		}

		for _, stmt := range x.Body.List {
			returnStmt, is := stmt.(*ast.ReturnStmt)
			if !is {
				continue
			}
			if len(returnStmt.Results) < 3 {
				continue
			}

			// zksync
			// convert tx to &CustomTransaction{Transaction: tx, customHash: tx.Hash()}
			txExpr, ok := returnStmt.Results[1].(*ast.Ident)
			if !ok {
				return true
			}
			if txExpr.Name != "tx" {
				return true
			}
			txField := &ast.KeyValueExpr{
				Key:   ast.NewIdent("Transaction"),
				Value: ast.NewIdent("tx"),
			}
			hashField := &ast.KeyValueExpr{
				Key: ast.NewIdent("CustomHash"),
				Value: &ast.CallExpr{
					Fun: &ast.SelectorExpr{
						X:   ast.NewIdent("tx"),
						Sel: ast.NewIdent("Hash"),
					},
				},
			}
			newRet := &ast.CompositeLit{
				Type: &ast.Ident{Name: "CustomTransaction"},
				Elts: []ast.Expr{txField, hashField},
			}
			pointerRet := &ast.UnaryExpr{Op: token.AND, X: newRet}
			returnStmt.Results[1] = pointerRet
		}

		return false
	}, nil).(*ast.File)
}

func addZKSyncImports(fset *token.FileSet, fileNode *ast.File) *ast.File {
	astutil.AddImport(fset, fileNode, "github.com/ethereum/go-ethereum/ethclient")
	astutil.AddImport(fset, fileNode, "context")
	astutil.AddNamedImport(fset, fileNode, "zkSyncClient", "github.com/zksync-sdk/zksync2-go/clients")
	astutil.AddNamedImport(fset, fileNode, "zkSyncAccounts", "github.com/zksync-sdk/zksync2-go/accounts")
	astutil.AddNamedImport(fset, fileNode, "zktypes", "github.com/zksync-sdk/zksync2-go/types")
	return fileNode
}

func addZKSyncBin(fileNode *ast.File, contractName string, zkHexString string) *ast.File {
	// zksync
	newVarSpec := &ast.ValueSpec{
		Names: []*ast.Ident{ast.NewIdent(fmt.Sprintf("%sZkBin", contractName))},
		Type:  ast.NewIdent("string"),
		Values: []ast.Expr{
			&ast.BasicLit{
				Kind:  token.STRING,
				Value: fmt.Sprintf("(\"0x%s\")", zkHexString),
			},
		},
	}
	newVarDecl := &ast.GenDecl{
		Tok:   token.VAR,
		Specs: []ast.Spec{newVarSpec},
	}

	// Insert the new variable declaration at the top of the file (before existing functions)
	fileNode.Decls = append(fileNode.Decls, newVarDecl)
	return fileNode
}

func getLogNames(fileNode *ast.File) []string {
	var logNames []string
	astutil.Apply(fileNode, func(cursor *astutil.Cursor) bool {
		x, is := cursor.Node().(*ast.FuncDecl)
		if !is {
			return true
		} else if !strings.HasPrefix(x.Name.Name, "Parse") {
			return false
		}
		logNames = append(logNames, x.Name.Name[len("Parse"):])
		return false
	}, nil)
	return logNames
}

func replaceAnonymousStructs(contractName string, fileNode *ast.File) *ast.File {
	done := map[string]bool{}
	return astutil.Apply(fileNode, func(cursor *astutil.Cursor) bool {
		// Replace all anonymous structs with named structs
		x, is := cursor.Node().(*ast.FuncDecl)
		if !is {
			return true
		} else if len(x.Type.Results.List) == 0 {
			return false
		}
		theStruct, is := x.Type.Results.List[0].Type.(*ast.StructType)
		if !is {
			return false
		}

		methodName := x.Name.Name
		x.Type.Results.List[0].Type = ast.NewIdent(methodName)

		x.Body = astutil.Apply(x.Body, func(cursor *astutil.Cursor) bool {
			if _, is := cursor.Node().(*ast.StructType); !is {
				return true
			}
			if call, is := cursor.Parent().(*ast.CallExpr); is {
				for i, arg := range call.Args {
					if arg == cursor.Node() {
						call.Args[i] = ast.NewIdent(methodName)
						break
					}
				}
			}
			return true
		}, nil).(*ast.BlockStmt)

		if done[contractName+methodName] {
			return true
		}

		// Add the named structs to the bottom of the file
		fileNode.Decls = append(fileNode.Decls, &ast.GenDecl{
			Tok: token.TYPE,
			Specs: []ast.Spec{
				&ast.TypeSpec{
					Name: ast.NewIdent(methodName),
					Type: theStruct,
				},
			},
		})

		done[contractName+methodName] = true
		return false
	}, nil).(*ast.File)
}

func writeAdditionalMethods(contractName string, logNames []string, abi abi.ABI, bs []byte) []byte {
	// Write the ParseLog method
	if len(logNames) > 0 {
		var logSwitchBody string
		for _, logName := range logNames {
			logSwitchBody += fmt.Sprintf(`case _%v.abi.Events["%v"].ID:
        return _%v.Parse%v(log)
`, contractName, logName, contractName, logName)
		}

		bs = append(bs, []byte(fmt.Sprintf(`
func (_%v *%v) ParseLog(log types.Log) (generated.AbigenLog, error) {
    switch log.Topics[0] {
    %v
    default:
        return nil, fmt.Errorf("abigen wrapper received unknown log topic: %%v", log.Topics[0])
    }
}
`, contractName, contractName, logSwitchBody))...)
	}

	// Write the Topic method
	for _, logName := range logNames {
		bs = append(bs, []byte(fmt.Sprintf(`
func (%v%v) Topic() common.Hash {
    return common.HexToHash("%v")
}
`, contractName, logName, abi.Events[logName].ID.Hex()))...)
	}

	// Write the Address method to the bottom of the file
	bs = append(bs, []byte(fmt.Sprintf(`
func (_%v *%v) Address() common.Address {
    return _%v.address
}
`, contractName, contractName, contractName))...)

	return bs
}

func writeInterface(contractName string, fileNode *ast.File) *ast.File {
	// Generate an interface for the contract
	var methods []*ast.Field
	ast.Inspect(fileNode, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.FuncDecl:
			if x.Recv == nil {
				return true
			}
			star, is := x.Recv.List[0].Type.(*ast.StarExpr)
			if !is {
				return false
			}

			typeName := star.X.(*ast.Ident).String()
			if typeName != contractName && typeName != contractName+"Caller" && typeName != contractName+"Transactor" && typeName != contractName+"Filterer" {
				return true
			}

			methods = append(methods, &ast.Field{
				Names: []*ast.Ident{x.Name},
				Type:  x.Type,
			})
		}
		return true
	})

	fileNode.Decls = append(fileNode.Decls, &ast.GenDecl{
		Tok: token.TYPE,
		Specs: []ast.Spec{
			&ast.TypeSpec{
				Name: ast.NewIdent(contractName + "Interface"),
				Type: &ast.InterfaceType{
					Methods: &ast.FieldList{
						List: methods,
					},
				},
			},
		},
	})

	return fileNode
}

func addHeader(code []byte) []byte {
	return utils.ConcatBytes([]byte(headerComment), code)
}
