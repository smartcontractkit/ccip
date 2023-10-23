package testconfig

import (
	"bytes"
	_ "embed"
	"encoding/base64"
	"fmt"

	"github.com/pelletier/go-toml/v2"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/smartcontractkit/chainlink-testing-framework/blockchain"
	"github.com/smartcontractkit/chainlink-testing-framework/networks"
	"github.com/smartcontractkit/chainlink-testing-framework/utils"

	"github.com/smartcontractkit/chainlink/v2/core/store/models"
	"github.com/smartcontractkit/chainlink/v2/core/utils/config"
)

const (
	ErrReadConfig      = "failed to read TOML config"
	ErrUnmarshalConfig = "failed to unmarshal TOML config"

	Load  string = "Load"
	Chaos string = "Chaos"
	Smoke string = "Smoke"

	CCIP = "CCIP"
)

var (
	//go:embed tomls/default.toml
	DefaultConfig    []byte
	GlobalTestConfig *Config
)

func initv1() {
	var err error
	GlobalTestConfig, err = NewConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load config")
	}
}

// GenericConfig is an interface for all product based config types to implement
type GenericConfig interface {
	ReadSecrets() error
	Validate() error
}

type Group interface {
	ApplyOverrides(from Group) error
	GenericConfig
}

// Config is the top level config struct. It contains config for all product based tests.
type Config struct {
	Products map[string]ProductTest `toml:",omitempty"`
}

// CCIP is the config for CCIP tests for a particular test group
func (c Config) CCIP(group string) (ProductTest, CCIPTestConfig, error) {
	ccip, ok := c.Products[CCIP]
	if !ok {
		return ProductTest{}, CCIPTestConfig{}, errors.New("no CCIP config found")
	}
	testCfg, exists := ccip.Groups[group]
	if !exists {
		return ccip, CCIPTestConfig{}, errors.Errorf("no CCIP test config found for test type %s", group)
	}
	testGroupCfg, ok := testCfg.(CCIPTestConfig)
	if !ok {
		return ccip, CCIPTestConfig{}, errors.Errorf("invalid CCIP test config type %T", testCfg)
	}
	return ccip, testGroupCfg, nil
}

func NewConfig() (*Config, error) {
	var cfg *Config
	var override *Config
	// load config from default file
	err := config.DecodeTOML(bytes.NewReader(DefaultConfig), cfg)
	if err != nil {
		return nil, errors.Wrap(err, ErrReadConfig)
	}

	// load config from env var if specified
	rawConfig, _ := utils.GetEnv("TEST_CONFIG_OVERRIDE")
	if rawConfig != "" {
		d, err := base64.StdEncoding.DecodeString(rawConfig)
		err = toml.Unmarshal(d, &override)
		if err != nil {
			return nil, errors.Wrap(err, ErrUnmarshalConfig)
		}
	}
	if override != nil {
		// apply overrides for all products
		for name, product := range override.Products {
			if existing, ok := cfg.Products[name]; ok {
				err := existing.ApplyOverrides(product)
				if err != nil {
					return nil, err
				}
			} else {
				cfg.Products[name] = product
			}
		}
	}
	// read secrets for all products
	for _, product := range cfg.Products {
		err := product.ReadSecrets()
		if err != nil {
			return nil, err
		}
	}

	// validate all products
	for _, product := range cfg.Products {
		err := product.Validate()
		if err != nil {
			return nil, err
		}
	}
	log.Debug().Interface("Config", cfg).Msg("Parsed config")
	return cfg, nil
}

// ProductTest is the generic config struct which can be used with product specific configs.
// It contains generic DON and networks config which can be applied to all product based tests.
type ProductTest struct {
	EnvUser   string           `toml:",omitempty"`
	TTL       *models.Duration `toml:",omitempty"`
	Chainlink *Chainlink       `toml:",omitempty"`
	Networks  []string         `toml:",omitempty"`
	Groups    map[string]Group `toml:",omitempty"`
}

func (p ProductTest) ReadSecrets() error {
	// read secrets for all products and test types
	// TODO: as of now we read network secrets through networks.SetNetworks, change this to generic secret reading mechanism
	err := p.Chainlink.ReadSecrets()
	if err != nil {
		return err
	}
	for _, group := range p.Groups {
		err := group.ReadSecrets()
		if err != nil {
			return err
		}
	}
	return nil
}

func (p ProductTest) ApplyOverrides(from ProductTest) error {
	if from.EnvUser != "" {
		p.EnvUser = from.EnvUser
	}
	if from.TTL != nil {
		p.TTL = from.TTL
	}
	if from.Networks != nil {
		p.Networks = from.Networks
	}
	if from.Chainlink != nil {
		if p.Chainlink == nil {
			p.Chainlink = &Chainlink{}
		}
		p.Chainlink.ApplyOverrides(from.Chainlink)
	}
	if from.Groups != nil {
		if p.Groups == nil {
			p.Groups = make(map[string]Group)
		}
		for name, group := range from.Groups {
			if existing, ok := p.Groups[name]; ok {
				err := existing.ApplyOverrides(group)
				if err != nil {
					return err
				}
			} else {
				p.Groups[name] = group
			}
		}
	}
	return nil
}

func (p ProductTest) Validate() error {
	if p.Networks == nil {
		return errors.New("no networks specified")
	}
	err := p.Chainlink.Validate()
	if err != nil {
		return err
	}
	for _, group := range p.Groups {
		err := group.Validate()
		if err != nil {
			return err
		}
	}
	return nil
}

func (p ProductTest) EVMNetworks() []blockchain.EVMNetwork {
	return networks.SetNetworks(p.Networks)
}

type Chainlink struct {
	Common     *Node   `toml:",omitempty"`
	NodeMemory string  `toml:",omitempty"`
	NodeCPU    string  `toml:",omitempty"`
	DBMemory   string  `toml:",omitempty"`
	DBCPU      string  `toml:",omitempty"`
	DBArgs     string  `toml:",omitempty"`
	NoOfNodes  *int    `toml:",omitempty"`
	Nodes      []*Node `toml:",omitempty"` // to be mentioned only if diff nodes follow diff configs; not required if all nodes follow CommonConfig
}

func (c *Chainlink) ApplyOverrides(from *Chainlink) {
	if from.NoOfNodes != nil {
		c.NoOfNodes = from.NoOfNodes
	}
	if from.Common != nil {
		c.Common.ApplyOverrides(from.Common)
	}
	if from.Nodes != nil {
		for i, node := range from.Nodes {
			if len(c.Nodes) > i {
				c.Nodes[i].ApplyOverrides(node)
			} else {
				c.Nodes = append(c.Nodes, node)
			}
		}
	}
	if from.NodeMemory != "" {
		c.NodeMemory = from.NodeMemory
	}
	if from.NodeCPU != "" {
		c.NodeCPU = from.NodeCPU
	}
	if from.DBMemory != "" {
		c.DBMemory = from.DBMemory
	}
	if from.DBCPU != "" {
		c.DBCPU = from.DBCPU
	}
	if from.DBArgs != "" {
		c.DBArgs = from.DBArgs
	}
}

func (c *Chainlink) ReadSecrets() error {
	image, _ := utils.GetEnv("CHAINLINK_IMAGE")
	if image != "" {
		c.Common.Image = image
	}
	tag, _ := utils.GetEnv("CHAINLINK_VERSION")
	if tag != "" {
		c.Common.Tag = tag
	}
	for i, node := range c.Nodes {
		image, _ := utils.GetEnv(fmt.Sprintf("CHAINLINK_IMAGE-%d", i+1))
		if image != "" {
			node.Image = image
		}
		tag, _ := utils.GetEnv(fmt.Sprintf("CHAINLINK_VERSION-%d", i+1))
		if tag != "" {
			node.Tag = tag
		}
	}
	return nil
}

func (c *Chainlink) Validate() error {
	if c.Common == nil && c.Nodes == nil {
		return errors.New("chainlink config is empty, either Common or Nodes should be specified")
	}
	if c.Common != nil && c.Nodes != nil {
		return errors.New("chainlink config is invalid, either Common or Nodes should be specified")
	}
	return nil
}

type Node struct {
	Name       string `toml:",omitempty"`
	Image      string `toml:",omitempty"`
	Tag        string `toml:",omitempty"`
	NodeConfig string `toml:",omitempty"`
	DBImage    string `toml:",omitempty"`
	DBTag      string `toml:",omitempty"`
}

func (n *Node) ApplyOverrides(from *Node) {
	if from.Name != "" {
		n.Name = from.Name
	}
	if from.Image != "" {
		n.Image = from.Image
	}
	if from.Tag != "" {
		n.Tag = from.Tag
	}
	if from.DBImage != "" {
		n.DBImage = from.DBImage
	}
	if from.DBTag != "" {
		n.DBTag = from.DBTag
	}
	if from.NodeConfig != "" {
		n.NodeConfig = from.NodeConfig
	}
}
