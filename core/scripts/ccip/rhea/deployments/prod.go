package deployments

import (
	gethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/core/scripts/ccip/rhea"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/arm_contract"
	"github.com/smartcontractkit/chainlink/v2/core/store/models"
)

var ProdChains = map[rhea.Chain]rhea.EvmDeploymentConfig{
	rhea.AvaxFuji:       {ChainConfig: Prod_AvaxFuji},
	rhea.OptimismGoerli: {ChainConfig: Prod_OptimismGoerli},
	rhea.Sepolia:        {ChainConfig: Prod_Sepolia},
	rhea.ArbitrumGoerli: {ChainConfig: Prod_ArbitrumGoerli},
	rhea.PolygonMumbai:  {ChainConfig: Prod_PolygonMumbai},
}

var ProdChainMapping = map[rhea.Chain]map[rhea.Chain]rhea.EvmDeploymentConfig{
	rhea.Sepolia: {
		rhea.AvaxFuji:       Prod_SepoliaToAvaxFuji,
		rhea.OptimismGoerli: Prod_SepoliaToOptimismGoerli,
		rhea.ArbitrumGoerli: Prod_SepoliaToArbitrumGoerli,
		rhea.PolygonMumbai:  Prod_SepoliaToPolygonMumbai,
	},
	rhea.AvaxFuji: {
		rhea.Sepolia:        Prod_AvaxFujiToSepolia,
		rhea.OptimismGoerli: Prod_AvaxFujiToOptimismGoerli,
		rhea.PolygonMumbai:  Prod_AvaxFujiToPolygonMumbai,
	},
	rhea.OptimismGoerli: {
		rhea.Sepolia:        Prod_OptimismGoerliToSepolia,
		rhea.AvaxFuji:       Prod_OptimismGoerliToAvaxFuji,
		rhea.ArbitrumGoerli: Prod_OptimismGoerliToArbitrumGoerli,
	},
	rhea.ArbitrumGoerli: {
		rhea.Sepolia:        Prod_ArbitrumGoerliToSepolia,
		rhea.OptimismGoerli: Prod_ArbitrumGoerliToOptimismGoerli,
	},
	rhea.PolygonMumbai: {
		rhea.Sepolia:  Prod_PolygonMumbaiToSepolia,
		rhea.AvaxFuji: Prod_PolygonMumbaiToAvaxFuji,
	},
}

var Prod_Sepolia = rhea.EVMChainConfig{
	EvmChainId: 11155111,
	GasSettings: rhea.EVMGasSettings{
		EIP1559: false,
	},
	AllowList: []gethcommon.Address{
		// ==============  INTERNAL ==============
		gethcommon.HexToAddress("0xda9e8e71bb750a996af33ebb8abb18cd9eb9dc75"), // deployer key
		gethcommon.HexToAddress("0x9BE566ad50021129F00Ee7219FcEE28490a85656"), // batch testing key
		gethcommon.HexToAddress("0xd54ba5d998479352f375940E5A2A18272714d434"), // batch testing key
		gethcommon.HexToAddress("0x28C70D03e471a2f1D1cad1DC35e7D90AAd2Ac512"), // batch testing key
		gethcommon.HexToAddress("0x5d39fF1Ae4Ab23E3640aa87a5C050483b53b9030"), // batch testing key
		gethcommon.HexToAddress("0x50C38847c059a7c829F7AEee969C652922bd139B"), // batch testing key
		gethcommon.HexToAddress("0x63fc8eE3Dc2326BC17A5E618872C1a4342Bcca09"), // batch testing key
		gethcommon.HexToAddress("0x68f740b79B9abe81628a654f8f733dd4ccE44DFB"), // batch testing key
		gethcommon.HexToAddress("0x0c55B0d8f41E6094a3d0F737c73E892ED0A52D8f"), // batch testing key
		gethcommon.HexToAddress("0x37ffDEe6Dc234E0D1d66571E2c2405aEfd661A6f"), // batch testing key
		gethcommon.HexToAddress("0x450F58153db2289B422e7629Eb4a70cFF77aA72f"), // batch testing key
		// Ping pong
		gethcommon.HexToAddress("0x7a7783e6073175f58db4d5f8bb40ea44065246db"), // SepoliaToAvaxFuji.PingPongDapp,
		gethcommon.HexToAddress("0x37b27863a14781acf41b787cf9ec3fb65d1c5885"), // SepoliaToOptimismGoerli.PingPongDapp,
		gethcommon.HexToAddress("0x65b51ba5c9233465f118285e5fb2110c52ad6b27"), // SepoliaToArbitrum.PingPongDapp,
		gethcommon.HexToAddress("0xf66fcb898e838a997547ae58fd6882b9bbfdc399"), // SepoliaToPolygonMumbai.PingPongDapp,
		// Personal
		gethcommon.HexToAddress("0xEa94AA1318796b5C01a9A37faCBc65423fb2c520"), // Anindita Ghosh
		gethcommon.HexToAddress("0x25D7214ae75F169263921a1cAaf7E6F033210E24"), // Chris Cushman
		gethcommon.HexToAddress("0x498533848239DDc6Bb5Cf7aEF63c97f3f5513ed2"), // Pramod - DApp Sepolia->Fuji
		gethcommon.HexToAddress("0x8e5267453b0aa137Be1Fc976755E6A9bD2a2E029"), // Amine (DevRel) 1
		gethcommon.HexToAddress("0x9d087fC03ae39b088326b67fA3C788236645b717"), // Amine (DevRel) 2
		gethcommon.HexToAddress("0x8fDEA7A82D7861144D027e4eb2acCCf4eB37bb05"), // Andrej Rakic
		gethcommon.HexToAddress("0x208AA722Aca42399eaC5192EE778e4D42f4E5De3"), // Zubin Pratap
		gethcommon.HexToAddress("0x52eE5a881287486573cF5CB5e7E7D92F30b03014"), // Zubin Pratap
		gethcommon.HexToAddress("0x44794725885F23cf36deE43554Ad204fb634A057"), // Frank Kong
		gethcommon.HexToAddress("0x0F1aF0A5d727b53dA44bBDE16843B3BA7F98Af68"), // Frank Kong
		gethcommon.HexToAddress("0x5803a251D118899dF7B403769e72532dBE854712"), // Amine El Manaa
		gethcommon.HexToAddress("0xcD936a39336a2E2c5a011137E46c8120dcaE0d65"), // Internal devrel proxy
		gethcommon.HexToAddress("0x03f6D53c4B337EE4D121db358baf33dF8c71108C"), // Zubin Pratap
		gethcommon.HexToAddress("0xeA3977eCC954a820cc709f6Edf8b467FDbe085A9"), // Richard Gottleber
		gethcommon.HexToAddress("0xF7b4ef69E7Cf13C205566345CcFAd1aB5fdCc49F"), // Harry Papacharissiou
		gethcommon.HexToAddress("0x9680201d9c93d65a3603d2088d125e955c73BD65"), // Patrick Collins

		// ==============  EXTERNAL ==============
		gethcommon.HexToAddress("0xd65113b9B1EeD81113EaF41DC0D2d34fCa31522C"), // BetaUser - Multimedia
		gethcommon.HexToAddress("0x217F4Eb693C54cA36Cfd80DA4DAAE6f7A5535e9C"), // BetaUser - Cozy Labs
		gethcommon.HexToAddress("0xB22107572f5A5352dDC1B4fc9630083FBfAE2022"), // BetaUser - Cozy Labs
		gethcommon.HexToAddress("0xB0AC8F6AF9712CF369934A811A79550DA046Fc51"), // BetaUser - InsurAce
		gethcommon.HexToAddress("0x244d07fe4DFa30b4EE376751FDC793aE844c5dE6"), // BetaUser - CACHE.gold
		gethcommon.HexToAddress("0x8264AcEE321ac02549aff7fA05A4Ae7a2e92A6f1"), // BetaUser - CACHE.gold
		gethcommon.HexToAddress("0x012a3fda37649945Cc72D725168FcB57A469bA6A"), // BetaUser - CACHE.gold
		gethcommon.HexToAddress("0x552acA1343A6383aF32ce1B7c7B1b47959F7ad90"), // BetaUser - Sommelier Finance
		gethcommon.HexToAddress("0x8e0866aacCF880E45249e932a094c821Ef4dE5f7"), // BetaUser - OpenZeppelin
		gethcommon.HexToAddress("0x9bf889acd6dd651bd897b6ff7a6ecde84a4b29aa"), // BetaUser - ANZ
		gethcommon.HexToAddress("0x9E945BB44B7E264c579e7f0c1FC28FBb39a32386"), // BetaUser - ANZ
		gethcommon.HexToAddress("0x309bdb4F7608584653D1bE804E8420fA0302911b"), // BetaUser - ANZ
		gethcommon.HexToAddress("0x066AFe67f2762C4009637c5ac10C789738cc7488"), // BetaUser - Tristero
		gethcommon.HexToAddress("0x6d818effaE3B40a89AEEb0e0FbA1827EFf77e0E1"), // BetaUser - Tristero
		gethcommon.HexToAddress("0x1C4310602DEFc04117980080b1807eac15687649"), // BetaUser - Zaros (ZD Labs)
		gethcommon.HexToAddress("0x4d2F1C99BCE324B9Ba486d704A0235A754D188a2"), // BetaUser - Aave (BGD Labs)
		gethcommon.HexToAddress("0x289F4D1e83BE7bb8A493D55622cE09D72D2A16e6"), // BetaUser - Steadefi
		gethcommon.HexToAddress("0x651c84ACc85D7a4506FD5dd6EB94d050c7ED2fe7"), // BetaUser - Lendvest
		gethcommon.HexToAddress("0xf62FD6119EBAFEdAAa7a75C1713Bca98729f163D"), // BetaUser - Fidelity Digital Assets
		gethcommon.HexToAddress("0x0D7a3a17E2E160287D3e7e74c4A1B22422156642"), // BetaUser - RiseWorks
		gethcommon.HexToAddress("0xc5f502Ae5972c938940b33308f8845cbe80211B5"), // BetaUser - Robolabs
		gethcommon.HexToAddress("0x87F45de79da4c3356591d74619693E372D525F1b"), // BankToken 1 (BANK) - internal testing contract for SWIFT POC
		gethcommon.HexToAddress("0x784c400D6fF625051d2f587dC0276E3A1ffD9cda"), // BankToken 2 (BANK) - internal testing contract for SWIFT POC
		gethcommon.HexToAddress("0xF92E4b278380f39fADc24483C7baC61b73EE93F2"), // BetaUser - SWIFT (BondToken)
		gethcommon.HexToAddress("0xAa6f663a14b8dA1EB9CF021379f4Ba6BF536268A"), // BetaUser - Fidelity Digital Assets
		gethcommon.HexToAddress("0xB781A9EFC6bd4Cf0dbE547D20151A405673F4CDe"), // BetaUser - RiseWorks
		gethcommon.HexToAddress("0xe764C455e3Bd05Eb7Cf53Ec8491dca0e91486D24"), // BetaUser - Synthetix v3 core
		gethcommon.HexToAddress("0x8e52262f91ef7049adfD8d1E608172fAC57995c3"), // BetaUser - Synthetix v3 core
		gethcommon.HexToAddress("0x2A45BaE1E58AaD3261af187b7dAde90889c039Dc"), // BetaUser - Synthetix v3 core
		gethcommon.HexToAddress("0xB3c3977B0aC329A9035889929482a4c635B50573"), // BetaUser - Alongside
		gethcommon.HexToAddress("0x0819BBae96c2C0F15477D212e063303221Cf24b9"), // BetaUser - Oddz
		gethcommon.HexToAddress("0x38104E1bB27A06306B72162047F585B3e6D27484"), // BetaUser - Oddz
		gethcommon.HexToAddress("0x789d7f3e2eaA6de41133A7fB11d7390603645F31"), // BetaUser - Galaxis
		gethcommon.HexToAddress("0xB1d5b1A03b2A76b80990F58551e95Fe7b29255BB"), // BetaUser - OpenXYZ
		gethcommon.HexToAddress("0x9B4a955444776f24FD976B9044Faf2FF902cD035"), // BetaUser - Yieldification
	},
	SupportedTokens: map[rhea.Token]rhea.EVMBridgedToken{
		rhea.LINK: {
			Token:         gethcommon.HexToAddress("0x779877A7B0D9E8603169DdbD7836e478b4624789"),
			Pool:          gethcommon.HexToAddress("0x5344b4bf5ae39038a591866d2853b2b1db622911"),
			Price:         rhea.LINK.Price(),
			Decimals:      rhea.LINK.Decimals(),
			TokenPoolType: rhea.LockRelease,
		},
		rhea.ANZ: {
			Token:         gethcommon.HexToAddress("0x92eA346B7a2AaB84e6AaB03b80E2421eeFB04685"),
			Pool:          gethcommon.HexToAddress("0x6c4cf212a5d074bb4d9055279b1fb2f1f37265db"),
			Price:         rhea.ANZ.Price(),
			Decimals:      rhea.ANZ.Decimals(),
			TokenPoolType: rhea.LockRelease,
		},
		rhea.Alongside: {
			Token:         gethcommon.HexToAddress("0xB3c3977B0aC329A9035889929482a4c635B50573"),
			Pool:          gethcommon.HexToAddress("0xac8cfc3762a979628334a0e4c1026244498e821b"),
			Price:         rhea.Alongside.Price(),
			Decimals:      rhea.Alongside.Decimals(),
			TokenPoolType: rhea.LockRelease,
		},
		rhea.BankToken: {
			Token:         gethcommon.HexToAddress("0x784c400D6fF625051d2f587dC0276E3A1ffD9cda"),
			Pool:          gethcommon.HexToAddress("0x5f217ce93e206d6f13b342aeef53a084fa957745"),
			Price:         rhea.BankToken.Price(),
			Decimals:      rhea.BankToken.Decimals(),
			TokenPoolType: rhea.LockRelease,
		},
		rhea.BondToken: {
			Token:         gethcommon.HexToAddress("0xF92E4b278380f39fADc24483C7baC61b73EE93F2"),
			Pool:          gethcommon.HexToAddress("0x919b1d308e4477c88350c336537ec5ac9ee76d9a"),
			Price:         rhea.BondToken.Price(),
			Decimals:      rhea.BondToken.Decimals(),
			TokenPoolType: rhea.LockRelease,
		},
		rhea.CCIP_BnM: {
			// NOTE this should be the custom burn_mint_erc677_helper contract, not the default burn_mint_erc677
			Token:    gethcommon.HexToAddress("0xFd57b4ddBf88a4e07fF4e34C487b99af2Fe82a05"),
			Pool:     gethcommon.HexToAddress("0x38d1ef9619cd40cf5482c045660ae7c82ada062c"),
			Price:    rhea.CCIP_BnM.Price(),
			Decimals: rhea.CCIP_BnM.Decimals(),
			// Wrapped is used to ensure new pool deployments will automatically grant burn/mint permissions
			TokenPoolType: rhea.Wrapped,
		},
		rhea.FUGAZIUSDC: {
			Token:         gethcommon.HexToAddress("0x832bA6abcAdC68812be372F4ef20aAC268bA20B7"),
			Pool:          gethcommon.HexToAddress("0x0ea0d7b2b78dd3a926fc76d6875a287f0aeb158f"),
			Price:         rhea.FUGAZIUSDC.Price(),
			Decimals:      rhea.FUGAZIUSDC.Decimals(),
			TokenPoolType: rhea.BurnMint,
		},
		rhea.InsurAce: {
			Token:         gethcommon.HexToAddress("0xb7c8bCA891143221a34DB60A26639785C4839040"),
			Pool:          gethcommon.HexToAddress("0xa04c2cbbfa7bf7adcbde911216a0ba7e3f1e36b3"),
			Price:         rhea.InsurAce.Price(),
			Decimals:      rhea.InsurAce.Decimals(),
			TokenPoolType: rhea.LockRelease,
		},
		rhea.STEADY: {
			Token:         gethcommon.HexToAddress("0x82abB1864326A8A7e1A357FFA2270D09CCb867B9"),
			Pool:          gethcommon.HexToAddress("0x5c0b55dbd1335a7c96653788cf545a8c08148496"),
			Price:         rhea.STEADY.Price(),
			Decimals:      rhea.STEADY.Decimals(),
			TokenPoolType: rhea.LockRelease,
		},
		rhea.WETH: {
			Token:         gethcommon.HexToAddress("0x097D90c9d3E0B50Ca60e1ae45F6A81010f9FB534"),
			Price:         rhea.WETH.Price(),
			Decimals:      rhea.WETH.Decimals(),
			TokenPoolType: rhea.FeeTokenOnly,
		},
		rhea.CCIP_LnM: {
			// NOTE this should be the custom burn_mint_erc677_helper contract, not the default burn_mint_erc677
			Token:    gethcommon.HexToAddress("0x466D489b6d36E7E3b824ef491C225F5830E81cC1"),
			Pool:     gethcommon.HexToAddress("0x3637220fccd067927766a40475f2e8fade33f590"),
			Price:    rhea.CCIP_LnM.Price(),
			Decimals: rhea.CCIP_LnM.Decimals(),
			// Wrapped is used to ensure new pool deployments will automatically grant burn/mint permissions
			TokenPoolType: rhea.Wrapped,
		},
		rhea.SNXUSD: {
			Token:         gethcommon.HexToAddress("0x585d8E269A250aCBf7D4884A1a31D3b596B46D8B"),
			Pool:          gethcommon.HexToAddress("0x2291909e328dd00e6b3284730d478f39dbde4c88"),
			Price:         rhea.SNXUSD.Price(),
			Decimals:      rhea.SNXUSD.Decimals(),
			TokenPoolType: rhea.BurnMint,
			PoolAllowList: []gethcommon.Address{
				gethcommon.HexToAddress("0xda9e8e71bb750a996af33ebb8abb18cd9eb9dc75"),
				gethcommon.HexToAddress("0x2A45BaE1E58AaD3261af187b7dAde90889c039Dc"),
			},
		},
		//rhea.CACHEGOLD: {
		//	Token:         gethcommon.HexToAddress("0x997BCCAE553112CD023592691d41687a3f1EfA7C"),
		//	Pool:          gethcommon.HexToAddress("0x23183132966fd8c91dbf6c13830b357ffeba98a2"),
		//	Price:         rhea.CACHEGOLD.Price(),
		//	Decimals:      rhea.CACHEGOLD.Decimals(),
		//	TokenPoolType: rhea.LockRelease,
		//},
		//rhea.ZUSD: {
		//	Token:         gethcommon.HexToAddress("0x09ae935D80E190403C61Cc5d854Fbf6a7b4a559a"),
		//	Pool:          gethcommon.HexToAddress(""),
		//	Price:         rhea.ZUSD.Price(),
		//	Decimals:      rhea.ZUSD.Decimals(),
		//	TokenPoolType: rhea.Legacy,
		//},
		//rhea.SUPER: {
		//	Token:         gethcommon.HexToAddress("0xCb4B3f72B5b6D0b7072aFDDf18FE61A0d569EC39"),
		//	Pool:          gethcommon.HexToAddress(""),
		//	Price:         rhea.SUPER.Price(),
		//	Decimals:      rhea.SUPER.Decimals(),
		//	TokenPoolType: rhea.Legacy,
		//},
	},
	FeeTokens:     []rhea.Token{rhea.LINK, rhea.WETH},
	WrappedNative: rhea.WETH,
	Router:        gethcommon.HexToAddress("0xd0daae2231e9cb96b94c8512223533293c3693bf"),
	ARM:           gethcommon.HexToAddress("0xb4d360459f32dd641ef5a6985ffbac5c4e5521aa"),
	ARMProxy:      gethcommon.HexToAddress("0xba3f6251de62ded61ff98590cb2fdf6871fbb991"),
	PriceRegistry: gethcommon.HexToAddress("0x8737a1c3d55779d03b7a08188e97af87b4110946"),
	TunableChainValues: rhea.TunableChainValues{
		FinalityDepth:            getFinalityDepth(rhea.Sepolia),
		OptimisticConfirmations:  getOptimisticConfirmations(rhea.Sepolia),
		BatchGasLimit:            BATCH_GAS_LIMIT,
		RelativeBoostPerWaitHour: RELATIVE_BOOST_PER_WAIT_HOUR,
		FeeUpdateHeartBeat:       models.MustMakeDuration(FEE_UPDATE_HEARTBEAT),
		FeeUpdateDeviationPPB:    FEE_UPDATE_DEVIATION_PPB,
		MaxGasPrice:              getMaxGasPrice(rhea.Sepolia),
		InflightCacheExpiry:      models.MustMakeDuration(INFLIGHT_CACHE_EXPIRY),
		RootSnoozeTime:           models.MustMakeDuration(ROOT_SNOOZE_TIME),
	},
	ARMConfig: &arm_contract.ARMConfig{
		Voters: []arm_contract.ARMVoter{
			// Infra-testnet-1
			{
				BlessVoteAddr:   gethcommon.HexToAddress("0x8e327be2a8bb2e95b7e281ec8fbdb327ea7cbbb1"),
				CurseVoteAddr:   gethcommon.HexToAddress("0x1936092090584fdf4542df8cc9b3ba695ef2cf88"),
				CurseUnvoteAddr: gethcommon.HexToAddress("0x0000000000000000000000000000000000000001"),
				BlessWeight:     1,
				CurseWeight:     1,
			},
			// Infra-testnet-2
			{
				BlessVoteAddr:   gethcommon.HexToAddress("0x671dea470d173e3a3fab9a463a9d85c9032da5e5"),
				CurseVoteAddr:   gethcommon.HexToAddress("0xca1dc0e1ef2a413c4672f3dfa28922a097ba32a2"),
				CurseUnvoteAddr: gethcommon.HexToAddress("0x0000000000000000000000000000000000000002"),
				BlessWeight:     1,
				CurseWeight:     1,
			},
			// Kostis-0
			{
				BlessVoteAddr:   gethcommon.HexToAddress("0x4679297e452b4b09ff2e351ddac3eff9c7999a17"),
				CurseVoteAddr:   gethcommon.HexToAddress("0xaa4d66d0db8ac802ba5ecfa4291c2b7aabe23a3f"),
				CurseUnvoteAddr: gethcommon.HexToAddress("0x0000000000000000000000000000000000000003"),
				BlessWeight:     1,
				CurseWeight:     1,
			},
			// Xueyuan-0
			{
				BlessVoteAddr:   gethcommon.HexToAddress("0x254090c0355c60aa7409c362196f65925393760e"),
				CurseVoteAddr:   gethcommon.HexToAddress("0x916bfcdb65d7216e869fea39eb6bbc5b15e61768"),
				CurseUnvoteAddr: gethcommon.HexToAddress("0x0000000000000000000000000000000000000004"),
				BlessWeight:     1,
				CurseWeight:     1,
			},
		},
		BlessWeightThreshold: 2,
		CurseWeightThreshold: 2,
	},
	DeploySettings: rhea.ChainDeploySettings{
		DeployARM:           false,
		DeployTokenPools:    false,
		DeployRouter:        false,
		DeployPriceRegistry: false,
	},
}

var Prod_OptimismGoerli = rhea.EVMChainConfig{
	EvmChainId: 420,
	GasSettings: rhea.EVMGasSettings{
		EIP1559: true,
	},
	AllowList: []gethcommon.Address{
		// ==============  INTERNAL ==============
		gethcommon.HexToAddress("0xda9e8e71bb750a996af33ebb8abb18cd9eb9dc75"), // deployer key
		// Ping pong
		gethcommon.HexToAddress("0x227c3699f9f0d6d55c38551a7d7feaea82efdd66"), // OptimismGoerliToAvaxFuji.PingPongDapp,
		gethcommon.HexToAddress("0x3b7a30028bf7ce52ad75b0afb142beef02deeecd"), // OptimismGoerliToSepolia.PingPongDapp,
		gethcommon.HexToAddress("0x2af63f50fa3f97f4aa94d28327a759ca86b33bf8"), // OptimismGoerliToArbitrumGoerli.PingPongDapp,
		// Personal
		gethcommon.HexToAddress("0xEa94AA1318796b5C01a9A37faCBc65423fb2c520"), // Anindita Ghosh
		gethcommon.HexToAddress("0x8fDEA7A82D7861144D027e4eb2acCCf4eB37bb05"), // Andrej Rakic
		gethcommon.HexToAddress("0x208AA722Aca42399eaC5192EE778e4D42f4E5De3"), // Zubin Pratap
		gethcommon.HexToAddress("0x52eE5a881287486573cF5CB5e7E7D92F30b03014"), // Zubin Pratap
		gethcommon.HexToAddress("0xabFD23063251A6481D65e8244237996d3D4d7b59"), // Internal devrel proxy
		gethcommon.HexToAddress("0x03f6D53c4B337EE4D121db358baf33dF8c71108C"), // Zubin Pratap
		gethcommon.HexToAddress("0x9680201d9c93d65a3603d2088d125e955c73BD65"), // Patrick Collins

		// ==============  EXTERNAL ==============
		gethcommon.HexToAddress("0x3FcFF7d9f88C64905e2cD9960c7452b5E6690E13"), // BetaUser - AAVE
		gethcommon.HexToAddress("0x1b5D803Be089e43110Faf54c6b4eC40409Cc7450"), // BetaUser - Multimedia
		gethcommon.HexToAddress("0xE8Cc2Bd6082387a7AC749176b1Fe19377f420740"), // BetaUser - Multimedia (AA wallet)
		gethcommon.HexToAddress("0x244d07fe4DFa30b4EE376751FDC793aE844c5dE6"), // BetaUser - CACHE.gold
		gethcommon.HexToAddress("0x8264AcEE321ac02549aff7fA05A4Ae7a2e92A6f1"), // BetaUser - CACHE.gold
		gethcommon.HexToAddress("0x012a3fda37649945Cc72D725168FcB57A469bA6A"), // BetaUser - CACHE.gold
		gethcommon.HexToAddress("0xF7726C9F7D2a9433CF8E46640821bebAbbE020b3"), // BetaUser - Zaros (ZD Labs)
		gethcommon.HexToAddress("0xF640cEA278E94708c358D79e5872AFda56010117"), // BetaUser - Aave (BGD Labs)
		gethcommon.HexToAddress("0x69D235A7E01aBdf463D7d886492229b75A4F1BC6"), // BetaUser - Steadefi
		gethcommon.HexToAddress("0xDdcE30979147091F26513C495EEE1bfa6C0a6730"), // BetaUser - RiseWorks
		gethcommon.HexToAddress("0xB3c3977B0aC329A9035889929482a4c635B50573"), // BetaUser - Alongside
		gethcommon.HexToAddress("0x2A45BaE1E58AaD3261af187b7dAde90889c039Dc"), // BetaUser - Synthetix v3 core
	},
	SupportedTokens: map[rhea.Token]rhea.EVMBridgedToken{
		rhea.Alongside: {
			Token:         gethcommon.HexToAddress("0xB3c3977B0aC329A9035889929482a4c635B50573"),
			Pool:          gethcommon.HexToAddress("0x8bcd622ac003160ea239c82e1b0e09364d77b1ac"),
			Price:         rhea.Alongside.Price(),
			Decimals:      rhea.Alongside.Decimals(),
			TokenPoolType: rhea.BurnMint,
		},
		rhea.CCIP_BnM: {
			// NOTE this should be the custom burn_mint_erc677_helper contract, not the default burn_mint_erc677
			Token:    gethcommon.HexToAddress("0xaBfE9D11A2f1D61990D1d253EC98B5Da00304F16"),
			Pool:     gethcommon.HexToAddress("0x8668ab4eb1dffe11db7491ebce633b050bb29cda"),
			Price:    rhea.CCIP_BnM.Price(),
			Decimals: rhea.CCIP_BnM.Decimals(),
			// Wrapped is used to ensure new pool deployments will automatically grant burn/mint permissions
			TokenPoolType: rhea.Wrapped,
		},
		rhea.LINK: {
			Token:         gethcommon.HexToAddress("0xdc2CC710e42857672E7907CF474a69B63B93089f"),
			Pool:          gethcommon.HexToAddress("0xdecfaf632175915bdf38c00d9d9746e8a90a56c4"),
			Price:         rhea.LINK.Price(),
			Decimals:      rhea.LINK.Decimals(),
			TokenPoolType: rhea.LockRelease,
		},
		rhea.WETH: {
			Token:         gethcommon.HexToAddress("0x4200000000000000000000000000000000000006"),
			Price:         rhea.WETH.Price(),
			Decimals:      rhea.WETH.Decimals(),
			TokenPoolType: rhea.FeeTokenOnly,
		},
		rhea.CCIP_LnM: {
			Token:         gethcommon.HexToAddress("0x835833d556299cdec623e7980e7369145b037591"),
			Pool:          gethcommon.HexToAddress("0xf66d20ac7b981e249fce8fb8ddae3974f5559735"),
			Price:         rhea.CCIP_LnM.Price(),
			Decimals:      rhea.CCIP_LnM.Decimals(),
			TokenPoolType: rhea.Wrapped,
		},
		rhea.SNXUSD: {
			Token:         gethcommon.HexToAddress("0x585d8E269A250aCBf7D4884A1a31D3b596B46D8B"),
			Pool:          gethcommon.HexToAddress("0x7a81e36b40b0a778afb2ae3c8384b6195233d093"),
			Price:         rhea.SNXUSD.Price(),
			Decimals:      rhea.SNXUSD.Decimals(),
			TokenPoolType: rhea.BurnMint,
			PoolAllowList: []gethcommon.Address{
				gethcommon.HexToAddress("0xda9e8e71bb750a996af33ebb8abb18cd9eb9dc75"),
				gethcommon.HexToAddress("0x2A45BaE1E58AaD3261af187b7dAde90889c039Dc"),
			},
		},
		//rhea.CACHEGOLD: {
		//	Token:         gethcommon.HexToAddress("0xa6446C6f492f31A33bC68249ae59F8871123a777"),
		//	Pool:          gethcommon.HexToAddress(""),
		//	Price:         rhea.CACHEGOLD.Price(),
		//	Decimals:      rhea.CACHEGOLD.Decimals(),
		//	TokenPoolType: rhea.Legacy,
		//},
		//rhea.ZUSD: {
		//	Token:         gethcommon.HexToAddress("0x740ba2E7f25c036ED0b19b83c9Da2cB8D756f9D5"),
		//	Pool:          gethcommon.HexToAddress(""),
		//	Price:         rhea.ZUSD.Price(),
		//	Decimals:      rhea.ZUSD.Decimals(),
		//	TokenPoolType: rhea.Legacy,
		//},
		//rhea.STEADY: {
		//	Token:         gethcommon.HexToAddress("0x615c83D5FEdafAEa641f1cC1a91ea09111EF0158"),
		//	Pool:          gethcommon.HexToAddress(""),
		//	Price:         rhea.STEADY.Price(),
		//	Decimals:      rhea.STEADY.Decimals(),
		//	TokenPoolType: rhea.Legacy,
		//},
	},
	FeeTokens:     []rhea.Token{rhea.LINK, rhea.WETH},
	WrappedNative: rhea.WETH,
	Router:        gethcommon.HexToAddress("0xeb52e9ae4a9fb37172978642d4c141ef53876f26"),
	ARM:           gethcommon.HexToAddress("0xeaf6968fab9c54ac31c3679f120705b5019d3546"),
	ARMProxy:      gethcommon.HexToAddress("0x4eb4dbdb3c3b56e5e209abf9c424a3834f2087d0"),
	PriceRegistry: gethcommon.HexToAddress("0x490f3b46fba6af0d7499867a73469a077251c2bb"),
	TunableChainValues: rhea.TunableChainValues{
		FinalityDepth:            getFinalityDepth(rhea.OptimismGoerli),
		OptimisticConfirmations:  getOptimisticConfirmations(rhea.OptimismGoerli),
		BatchGasLimit:            BATCH_GAS_LIMIT,
		RelativeBoostPerWaitHour: RELATIVE_BOOST_PER_WAIT_HOUR,
		FeeUpdateHeartBeat:       models.MustMakeDuration(FEE_UPDATE_HEARTBEAT),
		FeeUpdateDeviationPPB:    FEE_UPDATE_DEVIATION_PPB_FAST_CHAIN,
		MaxGasPrice:              getMaxGasPrice(rhea.OptimismGoerli),
		InflightCacheExpiry:      models.MustMakeDuration(INFLIGHT_CACHE_EXPIRY),
		RootSnoozeTime:           models.MustMakeDuration(ROOT_SNOOZE_TIME),
	},
	CustomerSettings: rhea.CustomerSettings{
		CacheGoldFeeAddress:  gethcommon.HexToAddress("0x8264AcEE321ac02549aff7fA05A4Ae7a2e92A6f1"),
		CacheGoldFeeEnforcer: gethcommon.HexToAddress("0x194E7a932663f11AC0790bfC44dBdd8339f0ED65"),
	},
	ARMConfig: &arm_contract.ARMConfig{
		Voters: []arm_contract.ARMVoter{
			// Infra-testnet-1
			{
				BlessVoteAddr:   gethcommon.HexToAddress("0x4d9987720ec678aa1271621ffe617771288e436f"),
				CurseVoteAddr:   gethcommon.HexToAddress("0x0e3594b19fb2b7ceb4e6872a9393b407579702b8"),
				CurseUnvoteAddr: gethcommon.HexToAddress("0x0000000000000000000000000000000000000001"),
				BlessWeight:     1,
				CurseWeight:     1,
			},
			// Infra-testnet-2
			{
				BlessVoteAddr:   gethcommon.HexToAddress("0xebe7b72da8ade2e1ed2077d51a933767029bf513"),
				CurseVoteAddr:   gethcommon.HexToAddress("0xb83efe598c14c80004dc75d2728ecc52c0113315"),
				CurseUnvoteAddr: gethcommon.HexToAddress("0x0000000000000000000000000000000000000002"),
				BlessWeight:     1,
				CurseWeight:     1,
			},
			// Kostis-0
			{
				BlessVoteAddr:   gethcommon.HexToAddress("0x7cf0592c4eda6b839b635ec0269df4d2e51ba1f4"),
				CurseVoteAddr:   gethcommon.HexToAddress("0xf5c77f0fbf0be8559b6ddf752fca4342eb8e254b"),
				CurseUnvoteAddr: gethcommon.HexToAddress("0x0000000000000000000000000000000000000003"),
				BlessWeight:     1,
				CurseWeight:     1,
			},
			// Xueyuan-0
			{
				BlessVoteAddr:   gethcommon.HexToAddress("0xc4e34bcc4b46b8e7fe02c8ac6fa8129f5027f3ef"),
				CurseVoteAddr:   gethcommon.HexToAddress("0xb94388d24dde6e7155cecec5d6474a1ce14f0127"),
				CurseUnvoteAddr: gethcommon.HexToAddress("0x0000000000000000000000000000000000000004"),
				BlessWeight:     1,
				CurseWeight:     1,
			},
		},
		BlessWeightThreshold: 2,
		CurseWeightThreshold: 2,
	},
	DeploySettings: rhea.ChainDeploySettings{
		DeployARM:           false,
		DeployTokenPools:    false,
		DeployRouter:        false,
		DeployPriceRegistry: false,
	},
}

var Prod_AvaxFuji = rhea.EVMChainConfig{
	EvmChainId: 43113,
	GasSettings: rhea.EVMGasSettings{
		EIP1559: false,
	},
	AllowList: []gethcommon.Address{
		// ==============  INTERNAL ==============
		gethcommon.HexToAddress("0xda9e8e71bb750a996af33ebb8abb18cd9eb9dc75"), // deployer key
		gethcommon.HexToAddress("0xEa94AA1318796b5C01a9A37faCBc65423fb2c520"), // Test Script 0xEa94AA1318796b5C01a9A37faCBc65423fb2c520
		// Ping pong
		gethcommon.HexToAddress("0xc1f01a6d0e8382f2c5d394923a7e79693354934b"), // AvaxFujiToSepolia.PingPongDapp,
		gethcommon.HexToAddress("0x4ce7b0782966d58ebc5e1804ca6de3244dac9ad9"), // AvaxFujiToOptimismGoerli.PingPongDapp,
		gethcommon.HexToAddress("0xcffb4c676a996daefa2a8a6d404a55f59ecc7ce8"), // AvaxFujiToPolygonMumbai.PingPongDapp,
		// Personal
		gethcommon.HexToAddress("0xEa94AA1318796b5C01a9A37faCBc65423fb2c520"), // Anindita Ghosh
		gethcommon.HexToAddress("0x594D8E57D8801069C77AAB90222a9162E908AA63"), // Pramod - Dapp Fuji->OptimismGoerli
		gethcommon.HexToAddress("0xFE5394A63433A3975b1936dEc92DAa161FEE7463"), // Pramod - DApp Fuji->Sepolia
		gethcommon.HexToAddress("0x912519a7E5e2e2309b1e60F540683c6661757A0C"), // Amine (DevRel) 1
		gethcommon.HexToAddress("0x9d087fC03ae39b088326b67fA3C788236645b717"), // Amine (DevRel) 2
		gethcommon.HexToAddress("0x8fDEA7A82D7861144D027e4eb2acCCf4eB37bb05"), // Andrej Rakic
		gethcommon.HexToAddress("0x208AA722Aca42399eaC5192EE778e4D42f4E5De3"), // Zubin Pratap
		gethcommon.HexToAddress("0x52eE5a881287486573cF5CB5e7E7D92F30b03014"), // Zubin Pratap
		gethcommon.HexToAddress("0x00104e54E037453daE202d6a92A3f75B2fdC2737"), // Amine El Manaa
		gethcommon.HexToAddress("0x447Fd5eC2D383091C22B8549cb231a3bAD6d3fAf"), // Internal devrel proxy
		gethcommon.HexToAddress("0x03f6D53c4B337EE4D121db358baf33dF8c71108C"), // Zubin Pratap
		gethcommon.HexToAddress("0xeA3977eCC954a820cc709f6Edf8b467FDbe085A9"), // Richard Gottleber
		gethcommon.HexToAddress("0xF7b4ef69E7Cf13C205566345CcFAd1aB5fdCc49F"), // Harry Papacharissiou
		gethcommon.HexToAddress("0x9680201d9c93d65a3603d2088d125e955c73BD65"), // Patrick Collins

		// ==============  EXTERNAL ==============
		gethcommon.HexToAddress("0x1b5D803Be089e43110Faf54c6b4eC40409Cc7450"), // BetaUser - Multimedia
		gethcommon.HexToAddress("0xE8Cc2Bd6082387a7AC749176b1Fe19377f420740"), // BetaUser - Multimedia (AA wallet)
		gethcommon.HexToAddress("0xa78ceF54da82D6279b20457F4D46294AfF59C871"), // BetaUser - Flash Liquidity
		gethcommon.HexToAddress("0x6613fd61bbfEF3291f2D7C7203Ceab212e880DbB"), // BetaUser - Flash Liquidity
		gethcommon.HexToAddress("0xa294275E5Bb4A786a3305f4276645290cCC7419B"), // BetaUser - Flash Liquidity
		gethcommon.HexToAddress("0xcA218DCFD26990223a2eDA70f3A568eaae22c051"), // BetaUser - Cozy Labs
		gethcommon.HexToAddress("0xD0fB066847d5DBc760E9575f79d9A044385e4079"), // BetaUser - Cozy Labs
		gethcommon.HexToAddress("0xD93C3Ae0949f905846FdfFc2b5b8A0a047dda59f"), // BetaUser - InsurAce
		gethcommon.HexToAddress("0x244d07fe4DFa30b4EE376751FDC793aE844c5dE6"), // BetaUser - CACHE.gold
		gethcommon.HexToAddress("0x8264AcEE321ac02549aff7fA05A4Ae7a2e92A6f1"), // BetaUser - CACHE.gold
		gethcommon.HexToAddress("0x012a3fda37649945Cc72D725168FcB57A469bA6A"), // BetaUser - CACHE.gold
		gethcommon.HexToAddress("0x1b38148B8DfdeA0B3D80C45F0d8569889504f0B5"), // BetaUser - Sommelier Finance
		gethcommon.HexToAddress("0xe0534662Ff1182a1C32E400d2b64723817344Ab4"), // BetaUser - Sommelier Finance
		gethcommon.HexToAddress("0x4986fD36b6b16f49b43282Ee2e24C5cF90ed166d"), // BetaUser - Sommelier Finance
		gethcommon.HexToAddress("0xc7a5d29248cf53b094106ca1d29634b34ad0fede"), // BetaUser - Tristero
		gethcommon.HexToAddress("0x4A5D71F7027684d473a1110a412B510354aF33e7"), // BetaUser - Aave (BGD Labs)
		gethcommon.HexToAddress("0x44eb6D97e98CE35eEFBD5764aa786f10121bC5e4"), // BetaUser - ANZ
		gethcommon.HexToAddress("0xa707480A11f12569b888306F2F118716d3BC29A1"), // BetaUser - Lendvest
		gethcommon.HexToAddress("0xbcFA8eAB1fCe576F1Ef71772E46519e0ADC06623"), // BetaUser - Lendvest
		gethcommon.HexToAddress("0xd35468ab2547a5ba9c9b809e67a35bcc5b89d2fe"), // BetaUser - Lendvest
		gethcommon.HexToAddress("0x9344AeA9b3270d51c9603d3054E421386dFaacB8"), // BetaUser - Fidelity Digital Assets
		gethcommon.HexToAddress("0x89Eccc61B2d35eACCe08284CF22c2D6487B80A3A"), // BetaUser - Robolabs
		gethcommon.HexToAddress("0xAa6f663a14b8dA1EB9CF021379f4Ba6BF536268A"), // BetaUser - Fidelity Digital Assets
		gethcommon.HexToAddress("0xB3c3977B0aC329A9035889929482a4c635B50573"), // BetaUser - Alongside
		gethcommon.HexToAddress("0xA9cb37191A089C8f8c24fC3e1F2f761De93FA827"), // BetaUser - MintDAO - Sender (EOA)
		gethcommon.HexToAddress("0xaFE336062eD69c108232c303fBa9b2b1c709fd9d"), // BetaUser - MintDAO - Sender (Proxy)
		gethcommon.HexToAddress("0x3e44ba19e932F4985983994DFa1Bd01f2a2f8eE2"), // BetaUser - Folks Finance
	},
	SupportedTokens: map[rhea.Token]rhea.EVMBridgedToken{
		rhea.ANZ: {
			Token:         gethcommon.HexToAddress("0xcdf7d0e9b7216160907722f9e10cccb0362c3b9c"),
			Pool:          gethcommon.HexToAddress("0x232794176678d6b500e263ef65be956dbac32692"),
			Price:         rhea.ANZ.Price(),
			Decimals:      rhea.ANZ.Decimals(),
			TokenPoolType: rhea.Wrapped,
		},
		rhea.Alongside: {
			Token:         gethcommon.HexToAddress("0xB3c3977B0aC329A9035889929482a4c635B50573"),
			Pool:          gethcommon.HexToAddress("0xaef84b05e96e2aafac9a347e60ae7e9a414fb649"),
			Price:         rhea.Alongside.Price(),
			Decimals:      rhea.Alongside.Decimals(),
			TokenPoolType: rhea.BurnMint,
		},
		rhea.BankToken: {
			Token:         gethcommon.HexToAddress("0x0147cba76c478aa46a76b8e2d2fdbd789d63b773"),
			Pool:          gethcommon.HexToAddress("0xd12e98b53446048e5ec614df514bc6838c5a8010"),
			Price:         rhea.BankToken.Price(),
			Decimals:      rhea.BankToken.Decimals(),
			TokenPoolType: rhea.Wrapped,
		},
		rhea.BondToken: {
			Token:         gethcommon.HexToAddress("0x8737a1c3d55779d03b7a08188e97af87b4110946"),
			Pool:          gethcommon.HexToAddress("0xfac166b229ca504c254bf89449da10d08d44cf69"),
			Price:         rhea.BondToken.Price(),
			Decimals:      rhea.BondToken.Decimals(),
			TokenPoolType: rhea.Wrapped,
		},
		rhea.CCIP_BnM: {
			// NOTE this should be the custom burn_mint_erc677_helper contract, not the default burn_mint_erc677
			Token:    gethcommon.HexToAddress("0xD21341536c5cF5EB1bcb58f6723cE26e8D8E90e4"),
			Pool:     gethcommon.HexToAddress("0xec1062cbdf4fbf31b3a6aac62b6f6f123bb70e12"),
			Price:    rhea.CCIP_BnM.Price(),
			Decimals: rhea.CCIP_BnM.Decimals(),
			// Wrapped is used to ensure new pool deployments will automatically grant burn/mint permissions
			TokenPoolType: rhea.Wrapped,
		},
		rhea.FUGAZIUSDC: {
			Token:         gethcommon.HexToAddress("0x150a0ee7393294442EE4d4F5C7d637af01dF93ee"),
			Pool:          gethcommon.HexToAddress("0x0040e5e502fe84de97b1b1cd7d33ca729d8b2a8b"),
			Price:         rhea.FUGAZIUSDC.Price(),
			Decimals:      rhea.FUGAZIUSDC.Decimals(),
			TokenPoolType: rhea.BurnMint,
		},
		rhea.InsurAce: {
			Token:         gethcommon.HexToAddress("0x005d27b4ee87e1f7362916ffa54bc37a30729554"),
			Pool:          gethcommon.HexToAddress("0x96a0c308a3293f0a4425bab68f37bdf6661eedad"),
			Price:         rhea.InsurAce.Price(),
			Decimals:      rhea.InsurAce.Decimals(),
			TokenPoolType: rhea.Wrapped,
		},
		rhea.LINK: {
			Token:         gethcommon.HexToAddress("0x0b9d5D9136855f6FEc3c0993feE6E9CE8a297846"),
			Pool:          gethcommon.HexToAddress("0x658af0d8ecbb13c5fd5b545ac7316e50cc07cf6e"),
			Price:         rhea.LINK.Price(),
			Decimals:      rhea.LINK.Decimals(),
			TokenPoolType: rhea.LockRelease,
		},
		rhea.WAVAX: {
			Token:         gethcommon.HexToAddress("0xd00ae08403B9bbb9124bB305C09058E32C39A48c"),
			Price:         rhea.WAVAX.Price(),
			Decimals:      rhea.WAVAX.Decimals(),
			TokenPoolType: rhea.FeeTokenOnly,
		},
		rhea.CCIP_LnM: {
			Token:         gethcommon.HexToAddress("0x70f5c5c40b873ea597776da2c21929a8282a3b35"),
			Pool:          gethcommon.HexToAddress("0x583dbe5f15dea93f321826d856994e53e01cd498"),
			Price:         rhea.CCIP_LnM.Price(),
			Decimals:      rhea.CCIP_LnM.Decimals(),
			TokenPoolType: rhea.Wrapped,
		},
		//rhea.CACHEGOLD: {
		//	Token:         gethcommon.HexToAddress("0xD16eD805F3eCe986d9541afaD3E59De2F3732517"),
		//	Pool:          gethcommon.HexToAddress(""),
		//	Price:         rhea.CACHEGOLD.Price(),
		//	Decimals:      rhea.CACHEGOLD.Decimals(),
		//	TokenPoolType: rhea.Legacy,
		//},
		//rhea.SUPER: {
		//	Token:         gethcommon.HexToAddress("0xCb4B3f72B5b6D0b7072aFDDf18FE61A0d569EC39"),
		//	Pool:          gethcommon.HexToAddress(""),
		//	Price:         rhea.SUPER.Price(),
		//	Decimals:      rhea.SUPER.Decimals(),
		//	TokenPoolType: rhea.Legacy,
		//},
	},
	FeeTokens:     []rhea.Token{rhea.LINK, rhea.WAVAX},
	WrappedNative: rhea.WAVAX,
	Router:        gethcommon.HexToAddress("0x554472a2720e5e7d5d3c817529aba05eed5f82d8"),
	ARM:           gethcommon.HexToAddress("0x0ea0d7b2b78dd3a926fc76d6875a287f0aeb158f"),
	ARMProxy:      gethcommon.HexToAddress("0xac8cfc3762a979628334a0e4c1026244498e821b"),
	PriceRegistry: gethcommon.HexToAddress("0xe42ecce39ce5bd2bbf2443660ba6979eeafd48df"),
	TunableChainValues: rhea.TunableChainValues{
		FinalityDepth:            getFinalityDepth(rhea.AvaxFuji),
		OptimisticConfirmations:  getOptimisticConfirmations(rhea.AvaxFuji),
		BatchGasLimit:            BATCH_GAS_LIMIT,
		RelativeBoostPerWaitHour: RELATIVE_BOOST_PER_WAIT_HOUR,
		FeeUpdateHeartBeat:       models.MustMakeDuration(FEE_UPDATE_HEARTBEAT),
		FeeUpdateDeviationPPB:    FEE_UPDATE_DEVIATION_PPB_FAST_CHAIN,
		MaxGasPrice:              getMaxGasPrice(rhea.AvaxFuji),
		InflightCacheExpiry:      models.MustMakeDuration(INFLIGHT_CACHE_EXPIRY),
		RootSnoozeTime:           models.MustMakeDuration(ROOT_SNOOZE_TIME),
	},
	CustomerSettings: rhea.CustomerSettings{
		CacheGoldFeeAddress:  gethcommon.HexToAddress("0x8264AcEE321ac02549aff7fA05A4Ae7a2e92A6f1"),
		CacheGoldFeeEnforcer: gethcommon.HexToAddress("0x194E7a932663f11AC0790bfC44dBdd8339f0ED65"),
	},
	ARMConfig: &arm_contract.ARMConfig{
		Voters: []arm_contract.ARMVoter{
			// Infra-testnet-1
			{
				BlessVoteAddr:   gethcommon.HexToAddress("0xac591a80ff5a81c512a5bb52c77e2513eca77245"),
				CurseVoteAddr:   gethcommon.HexToAddress("0xc5202b4be2f03ec895773c7ade8e79e9794ac214"),
				CurseUnvoteAddr: gethcommon.HexToAddress("0x0000000000000000000000000000000000000001"),
				BlessWeight:     1,
				CurseWeight:     1,
			},
			// Infra-testnet-2
			{
				BlessVoteAddr:   gethcommon.HexToAddress("0x722915bc6373d35bd051e1d61ff15edd2f0b0aae"),
				CurseVoteAddr:   gethcommon.HexToAddress("0x829a26e035a0d5e217960db184f9742c076c9ddd"),
				CurseUnvoteAddr: gethcommon.HexToAddress("0x0000000000000000000000000000000000000002"),
				BlessWeight:     1,
				CurseWeight:     1,
			},
			// Kostis-0
			{
				BlessVoteAddr:   gethcommon.HexToAddress("0x1874d82f4a25e2f2633106afd08baecbf3b52468"),
				CurseVoteAddr:   gethcommon.HexToAddress("0xbf131623483b4f0ac00371cce9c4f3b59339390e"),
				CurseUnvoteAddr: gethcommon.HexToAddress("0x0000000000000000000000000000000000000003"),
				BlessWeight:     1,
				CurseWeight:     1,
			},
			// Xueyuan-0
			{
				BlessVoteAddr:   gethcommon.HexToAddress("0x13e103d39e1970317e8d9dc05583bafb7b08f79e"),
				CurseVoteAddr:   gethcommon.HexToAddress("0x112f239515839d8349a62fbfb48bae79ffbbad74"),
				CurseUnvoteAddr: gethcommon.HexToAddress("0x0000000000000000000000000000000000000004"),
				BlessWeight:     1,
				CurseWeight:     1,
			},
		},
		BlessWeightThreshold: 2,
		CurseWeightThreshold: 2,
	},
	DeploySettings: rhea.ChainDeploySettings{
		DeployARM:           false,
		DeployTokenPools:    false,
		DeployRouter:        false,
		DeployPriceRegistry: false,
	},
}

var Prod_ArbitrumGoerli = rhea.EVMChainConfig{
	EvmChainId: 421613,
	GasSettings: rhea.EVMGasSettings{
		EIP1559:   true,
		GasTipCap: rhea.DefaultGasTipFee,
	},
	AllowList: []gethcommon.Address{
		// ==============  INTERNAL ==============
		gethcommon.HexToAddress("0xda9e8e71bb750a996af33ebb8abb18cd9eb9dc75"), // deployer key
		// Ping pong
		gethcommon.HexToAddress("0x9b451300c94c7328bdb56a514f83205ea789136f"), // ArbitrumGoerliToSepolia.PingPongDapp,
		gethcommon.HexToAddress("0x57c0059fc3f98aa0a5ce4fb5d2882d81d839e74f"), // ArbitrumGoerliToOptimismGoerli.PingPongDapp,
		// Personal
		gethcommon.HexToAddress("0x8fDEA7A82D7861144D027e4eb2acCCf4eB37bb05"), // Andrej Rakic
		gethcommon.HexToAddress("0x208AA722Aca42399eaC5192EE778e4D42f4E5De3"), // Zubin Pratap
		gethcommon.HexToAddress("0x52eE5a881287486573cF5CB5e7E7D92F30b03014"), // Zubin Pratap
		gethcommon.HexToAddress("0x82FAB72c5Baf6f15f89540EfBb7A62Cb410c300C"), // Internal devrel proxy
		gethcommon.HexToAddress("0x03f6D53c4B337EE4D121db358baf33dF8c71108C"), // Zubin Pratap
		gethcommon.HexToAddress("0x9680201d9c93d65a3603d2088d125e955c73BD65"), // Patrick Collins
		gethcommon.HexToAddress("0xEa94AA1318796b5C01a9A37faCBc65423fb2c520"), // Anindita Ghosh

		// ==============  EXTERNAL ==============
		gethcommon.HexToAddress("0xF5022eDd1B827E6EA4bBdb961212ECD7F315ed88"), // BetaUser - RiseWorks
		gethcommon.HexToAddress("0x0D7a3a17E2E160287D3e7e74c4A1B22422156642"), // BetaUser - RiseWorks
		gethcommon.HexToAddress("0x63e430dBd88C1bBFBc97336b4357Aa5Aea83367e"), // BetaUser - RiseWorks
		gethcommon.HexToAddress("0x2A45BaE1E58AaD3261af187b7dAde90889c039Dc"), // Synthetix v3 core
		gethcommon.HexToAddress("0x3e44ba19e932F4985983994DFa1Bd01f2a2f8eE2"), // BetaUser - Folks Finance
		gethcommon.HexToAddress("0xB573315f912a7e31A8d37054641c031c3c6A9eb9"), // BetaUser - Whitehole Finance
		gethcommon.HexToAddress("0x3D3F14b31eb5a86a64Eb7951CcAdBE9471eaAd23"), // BetaUser - Yieldification
	},
	SupportedTokens: map[rhea.Token]rhea.EVMBridgedToken{
		rhea.CCIP_BnM: {
			// NOTE this should be the custom burn_mint_erc677_helper contract, not the default burn_mint_erc677
			Token:    gethcommon.HexToAddress("0x0579b4c1C8AcbfF13c6253f1B10d66896Bf399Ef"),
			Pool:     gethcommon.HexToAddress("0xf399f6a4ea83442f97f480118ebd56d1aed767b9"),
			Price:    rhea.CCIP_BnM.Price(),
			Decimals: rhea.CCIP_BnM.Decimals(),
			// Wrapped is used to ensure new pool deployments will automatically grant burn/mint permissions
			TokenPoolType: rhea.Wrapped,
		},
		rhea.LINK: {
			Token:         gethcommon.HexToAddress("0xd14838A68E8AFBAdE5efb411d5871ea0011AFd28"),
			Pool:          gethcommon.HexToAddress("0x044a6b4b561af69d2319a2f4be5ec327a6975d0a"),
			Price:         rhea.LINK.Price(),
			Decimals:      rhea.LINK.Decimals(),
			TokenPoolType: rhea.LockRelease,
		},
		rhea.WETH: {
			Token:         gethcommon.HexToAddress("0x32d5D5978905d9c6c2D4C417F0E06Fe768a4FB5a"),
			Price:         rhea.WETH.Price(),
			Decimals:      rhea.WETH.Decimals(),
			TokenPoolType: rhea.FeeTokenOnly,
		},
		rhea.CCIP_LnM: {
			Token:         gethcommon.HexToAddress("0x0e14dbe2c8e1121902208be173a3fb91bb125cdb"),
			Pool:          gethcommon.HexToAddress("0xa77aefaba6161f907299dc2be79a60c9e80e9b91"),
			Price:         rhea.CCIP_LnM.Price(),
			Decimals:      rhea.CCIP_LnM.Decimals(),
			TokenPoolType: rhea.Wrapped,
		},
		rhea.SNXUSD: {
			Token:         gethcommon.HexToAddress("0x585d8E269A250aCBf7D4884A1a31D3b596B46D8B"),
			Pool:          gethcommon.HexToAddress("0x7990f84bbe3e6638907d5fcc217555146304a8d9"),
			Price:         rhea.SNXUSD.Price(),
			Decimals:      rhea.SNXUSD.Decimals(),
			TokenPoolType: rhea.BurnMint,
			PoolAllowList: []gethcommon.Address{
				gethcommon.HexToAddress("0xda9e8e71bb750a996af33ebb8abb18cd9eb9dc75"),
				gethcommon.HexToAddress("0x2A45BaE1E58AaD3261af187b7dAde90889c039Dc"),
			},
		},
	},
	FeeTokens:     []rhea.Token{rhea.LINK, rhea.WETH},
	WrappedNative: rhea.WETH,
	Router:        gethcommon.HexToAddress("0x88e492127709447a5abefdab8788a15b4567589e"),
	ARM:           gethcommon.HexToAddress("0x8af4204e30565df93352fe8e1de78925f6664da7"),
	ARMProxy:      gethcommon.HexToAddress("0x3cc9364260d80f09ccac1ee6b07366db598900e6"),
	PriceRegistry: gethcommon.HexToAddress("0x114a20a10b43d4115e5aeef7345a1a71d2a60c57"),
	TunableChainValues: rhea.TunableChainValues{
		FinalityDepth:            getFinalityDepth(rhea.ArbitrumGoerli),
		OptimisticConfirmations:  getOptimisticConfirmations(rhea.ArbitrumGoerli),
		BatchGasLimit:            BATCH_GAS_LIMIT,
		RelativeBoostPerWaitHour: RELATIVE_BOOST_PER_WAIT_HOUR,
		FeeUpdateHeartBeat:       models.MustMakeDuration(FEE_UPDATE_HEARTBEAT),
		FeeUpdateDeviationPPB:    FEE_UPDATE_DEVIATION_PPB_FAST_CHAIN,
		MaxGasPrice:              getMaxGasPrice(rhea.ArbitrumGoerli),
		InflightCacheExpiry:      models.MustMakeDuration(INFLIGHT_CACHE_EXPIRY),
		RootSnoozeTime:           models.MustMakeDuration(ROOT_SNOOZE_TIME),
	},
	ARMConfig: &arm_contract.ARMConfig{
		Voters: []arm_contract.ARMVoter{
			// Infra-testnet-1
			{
				BlessVoteAddr:   gethcommon.HexToAddress("0x9d2fefa5791ee383884df019e08d2fc307e776b1"),
				CurseVoteAddr:   gethcommon.HexToAddress("0x7203cf710b0fa2128c13e12672286a287890ec22"),
				CurseUnvoteAddr: gethcommon.HexToAddress("0x0000000000000000000000000000000000000001"),
				BlessWeight:     1,
				CurseWeight:     1,
			},
			// Infra-testnet-2
			{
				BlessVoteAddr:   gethcommon.HexToAddress("0x3e8eee517dba13675fff1b4f2ea4210902cba81b"),
				CurseVoteAddr:   gethcommon.HexToAddress("0x02cd3b5011567b70d41809907e672dcdd05285ee"),
				CurseUnvoteAddr: gethcommon.HexToAddress("0x0000000000000000000000000000000000000002"),
				BlessWeight:     1,
				CurseWeight:     1,
			},
			// Kostis-0
			{
				BlessVoteAddr:   gethcommon.HexToAddress("0x0a6e792d5c6f813e399341740cf7a368b6d66e6f"),
				CurseVoteAddr:   gethcommon.HexToAddress("0x263a09238c91cabe5e19642cd1c81fa567406c36"),
				CurseUnvoteAddr: gethcommon.HexToAddress("0x0000000000000000000000000000000000000003"),
				BlessWeight:     1,
				CurseWeight:     1,
			},
			// Xueyuan-0
			{
				BlessVoteAddr:   gethcommon.HexToAddress("0x1db6a4237b54fdfc7bd3c50c055e468c85832fea"),
				CurseVoteAddr:   gethcommon.HexToAddress("0x153a6331278bda690b4f7365d514a7a27017463a"),
				CurseUnvoteAddr: gethcommon.HexToAddress("0x0000000000000000000000000000000000000004"),
				BlessWeight:     1,
				CurseWeight:     1,
			},
		},
		BlessWeightThreshold: 2,
		CurseWeightThreshold: 2,
	},
	DeploySettings: rhea.ChainDeploySettings{
		DeployARM:           false,
		DeployTokenPools:    false,
		DeployRouter:        false,
		DeployPriceRegistry: false,
	},
}

var Prod_PolygonMumbai = rhea.EVMChainConfig{
	EvmChainId: 80001,
	GasSettings: rhea.EVMGasSettings{
		EIP1559:   true,
		GasTipCap: rhea.DefaultGasTipFee,
	},
	AllowList: []gethcommon.Address{
		// ==============  INTERNAL ==============
		gethcommon.HexToAddress("0xda9e8e71bb750a996af33ebb8abb18cd9eb9dc75"), // deployer key
		gethcommon.HexToAddress("0xEa94AA1318796b5C01a9A37faCBc65423fb2c520"), // Test Script 0xEa94AA1318796b5C01a9A37faCBc65423fb2c520
		// Ping pong
		gethcommon.HexToAddress("0x044a6b4b561af69d2319a2f4be5ec327a6975d0a"), // PolygonMumbaiToSepolia.PingPongDapp,
		gethcommon.HexToAddress("0x3bd38d308045a39253b502f1789e95c703e27f77"), // PolygonMumbaiToAvax.PingPongDapp,
		// Personal
		gethcommon.HexToAddress("0xEa94AA1318796b5C01a9A37faCBc65423fb2c520"), // Anindita Ghosh
		gethcommon.HexToAddress("0x8fDEA7A82D7861144D027e4eb2acCCf4eB37bb05"), // Andrej Rakic
		gethcommon.HexToAddress("0x208AA722Aca42399eaC5192EE778e4D42f4E5De3"), // Zubin Pratap
		gethcommon.HexToAddress("0x52eE5a881287486573cF5CB5e7E7D92F30b03014"), // Zubin Pratap
		gethcommon.HexToAddress("0x44794725885F23cf36deE43554Ad204fb634A057"), // Frank Kong
		gethcommon.HexToAddress("0x0F1aF0A5d727b53dA44bBDE16843B3BA7F98Af68"), // Frank Kong
		gethcommon.HexToAddress("0xA4285EC042b198aeb0C68679c94a615c4d82DAd0"), // Amine El Manaa
		gethcommon.HexToAddress("0xBdcc3f1D0B4c78F1fe03C91D9498f3DAEeE6948B"), // Internal devrel proxy
		gethcommon.HexToAddress("0x03f6D53c4B337EE4D121db358baf33dF8c71108C"), // Zubin Pratap
		gethcommon.HexToAddress("0xeA3977eCC954a820cc709f6Edf8b467FDbe085A9"), // Richard Gottleber
		gethcommon.HexToAddress("0xF7b4ef69E7Cf13C205566345CcFAd1aB5fdCc49F"), // Harry Papacharissiou
		gethcommon.HexToAddress("0x9680201d9c93d65a3603d2088d125e955c73BD65"), // Patrick Collins

		// ==============  EXTERNAL ==============
		gethcommon.HexToAddress("0xe764C455e3Bd05Eb7Cf53Ec8491dca0e91486D24"), // BetaUser - Synthetix v3 core
		gethcommon.HexToAddress("0x8e52262f91ef7049adfD8d1E608172fAC57995c3"), // BetaUser - Synthetix v3 core
		gethcommon.HexToAddress("0x6De1e981d2137f7839840e2140dBB3A05F05B770"), // BetaUser - Flash Liquidity
		gethcommon.HexToAddress("0x2A45BaE1E58AaD3261af187b7dAde90889c039Dc"), // BetaUser - Synthetix v3 core
		gethcommon.HexToAddress("0x0819BBae96c2C0F15477D212e063303221Cf24b9"), // BetaUser - Oddz
		gethcommon.HexToAddress("0x38104E1bB27A06306B72162047F585B3e6D27484"), // BetaUser - Oddz
		gethcommon.HexToAddress("0xA9cb37191A089C8f8c24fC3e1F2f761De93FA827"), // MintDAO - Sender (EOA)
		gethcommon.HexToAddress("0xaFE336062eD69c108232c303fBa9b2b1c709fd9d"), // MintDAO - Sender (Proxy)
		gethcommon.HexToAddress("0x789d7f3e2eaA6de41133A7fB11d7390603645F31"), // BetaUser - Galaxis
	},
	SupportedTokens: map[rhea.Token]rhea.EVMBridgedToken{
		rhea.CCIP_BnM: {
			// NOTE this should be the custom burn_mint_erc677_helper contract, not the default burn_mint_erc677
			Token:    gethcommon.HexToAddress("0xf1E3A5842EeEF51F2967b3F05D45DD4f4205FF40"),
			Pool:     gethcommon.HexToAddress("0xa6c88f12ae1aa9c333e86ccbdd2957cac2e5f58c"),
			Price:    rhea.CCIP_BnM.Price(),
			Decimals: rhea.CCIP_BnM.Decimals(),
			// Wrapped is used to ensure new pool deployments will automatically grant burn/mint permissions
			TokenPoolType: rhea.Wrapped,
		},
		rhea.LINK: {
			Token:         gethcommon.HexToAddress("0x326C977E6efc84E512bB9C30f76E30c160eD06FB"),
			Pool:          gethcommon.HexToAddress("0x6fce09b2e74f649a4494a1844219cb0d86cfe8b7"),
			Price:         rhea.LINK.Price(),
			Decimals:      rhea.LINK.Decimals(),
			TokenPoolType: rhea.LockRelease,
		},
		rhea.WMATIC: {
			Token:         gethcommon.HexToAddress("0x9c3C9283D3e44854697Cd22D3Faa240Cfb032889"),
			Price:         rhea.WMATIC.Price(),
			Decimals:      rhea.WMATIC.Decimals(),
			TokenPoolType: rhea.FeeTokenOnly,
		},
		rhea.CCIP_LnM: {
			Token:         gethcommon.HexToAddress("0xc1c76a8c5bfde1be034bbcd930c668726e7c1987"),
			Pool:          gethcommon.HexToAddress("0x83369f8586ba000a87db278549b9a2370dc626b6"),
			Price:         rhea.CCIP_LnM.Price(),
			Decimals:      rhea.CCIP_LnM.Decimals(),
			TokenPoolType: rhea.Wrapped,
		},
		rhea.SNXUSD: {
			Token:         gethcommon.HexToAddress("0x585d8E269A250aCBf7D4884A1a31D3b596B46D8B"),
			Pool:          gethcommon.HexToAddress("0x56ffa3ae8bb98b120067b1be50136f681aa3addf"),
			Price:         rhea.SNXUSD.Price(),
			Decimals:      rhea.SNXUSD.Decimals(),
			TokenPoolType: rhea.BurnMint,
			PoolAllowList: []gethcommon.Address{
				gethcommon.HexToAddress("0xda9e8e71bb750a996af33ebb8abb18cd9eb9dc75"),
				gethcommon.HexToAddress("0x2A45BaE1E58AaD3261af187b7dAde90889c039Dc"),
			},
		},
	},
	FeeTokens:     []rhea.Token{rhea.LINK, rhea.WMATIC},
	WrappedNative: rhea.WMATIC,
	Router:        gethcommon.HexToAddress("0x70499c328e1e2a3c41108bd3730f6670a44595d1"),
	ARM:           gethcommon.HexToAddress("0x917a6913f785094f8b06785aa8a884f922a650d8"),
	ARMProxy:      gethcommon.HexToAddress("0x235ce3408845a4767a2eaa2a3d8ef0848d283f1f"),
	PriceRegistry: gethcommon.HexToAddress("0x9bd312170aa145ef98453940dc9ab894235b063e"),
	TunableChainValues: rhea.TunableChainValues{
		FinalityDepth:            getFinalityDepth(rhea.PolygonMumbai),
		OptimisticConfirmations:  getOptimisticConfirmations(rhea.PolygonMumbai),
		BatchGasLimit:            BATCH_GAS_LIMIT,
		RelativeBoostPerWaitHour: RELATIVE_BOOST_PER_WAIT_HOUR,
		FeeUpdateHeartBeat:       models.MustMakeDuration(FEE_UPDATE_HEARTBEAT),
		FeeUpdateDeviationPPB:    FEE_UPDATE_DEVIATION_PPB_FAST_CHAIN,
		MaxGasPrice:              getMaxGasPrice(rhea.PolygonMumbai),
		InflightCacheExpiry:      models.MustMakeDuration(INFLIGHT_CACHE_EXPIRY),
		RootSnoozeTime:           models.MustMakeDuration(ROOT_SNOOZE_TIME),
	},
	ARMConfig: &arm_contract.ARMConfig{
		Voters: []arm_contract.ARMVoter{
			// Infra-testnet-1
			{
				BlessVoteAddr:   gethcommon.HexToAddress("0x6b33111e07a15f51a82bf30708e95dc2169ec100"),
				CurseVoteAddr:   gethcommon.HexToAddress("0x4e6239f0fdda2b81d4bc790c959caffcc47d8436"),
				CurseUnvoteAddr: gethcommon.HexToAddress("0x0000000000000000000000000000000000000001"),
				BlessWeight:     1,
				CurseWeight:     1,
			},
			// Infra-testnet-2
			{
				BlessVoteAddr:   gethcommon.HexToAddress("0x19efeb472fc308b777bef6282ad77688ff954181"),
				CurseVoteAddr:   gethcommon.HexToAddress("0xf69d7ba7a60b148a08881744cd1a415582703ad5"),
				CurseUnvoteAddr: gethcommon.HexToAddress("0x0000000000000000000000000000000000000002"),
				BlessWeight:     1,
				CurseWeight:     1,
			},
			// Kostis-0
			{
				BlessVoteAddr:   gethcommon.HexToAddress("0x8bccd4cd06f0b50e78da446ef0ac61f3b43aefc5"),
				CurseVoteAddr:   gethcommon.HexToAddress("0xe5c553af74b9badb9e7d52ebf5a52a6c556fba10"),
				CurseUnvoteAddr: gethcommon.HexToAddress("0x0000000000000000000000000000000000000003"),
				BlessWeight:     1,
				CurseWeight:     1,
			},
			// Xueyuan-0
			{
				BlessVoteAddr:   gethcommon.HexToAddress("0x76118c18bcaa561cea1bedf558cb9a11bfd7bf2c"),
				CurseVoteAddr:   gethcommon.HexToAddress("0xa19eadc34225c8a521047a7027421e868e73a577"),
				CurseUnvoteAddr: gethcommon.HexToAddress("0x0000000000000000000000000000000000000004"),
				BlessWeight:     1,
				CurseWeight:     1,
			},
		},
		BlessWeightThreshold: 2,
		CurseWeightThreshold: 2,
	},
	DeploySettings: rhea.ChainDeploySettings{
		DeployARM:           false,
		DeployTokenPools:    false,
		DeployRouter:        false,
		DeployPriceRegistry: false,
	},
}

var Prod_SepoliaToAvaxFuji = rhea.EvmDeploymentConfig{
	ChainConfig: Prod_Sepolia,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress("0xe42ecce39ce5bd2bbf2443660ba6979eeafd48df"),
		OffRamp:      gethcommon.HexToAddress("0x1f06781450e994b0005ce2922fca78e2c72d4353"),
		CommitStore:  gethcommon.HexToAddress("0xf3855a07bf75c4e0b4ddbeb7784badc9dd2ca274"),
		PingPongDapp: gethcommon.HexToAddress("0x7a7783e6073175f58db4d5f8bb40ea44065246db"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployLane:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    3816446,
		},
	},
}

var Prod_AvaxFujiToSepolia = rhea.EvmDeploymentConfig{
	ChainConfig: Prod_AvaxFuji,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress("0xa799c1855875e79b2e1752412058b485ee51aec4"),
		OffRamp:      gethcommon.HexToAddress("0x61c67e7b7c90ed1a44dabb26c33900270df7a144"),
		CommitStore:  gethcommon.HexToAddress("0xb4407405465a5dab21fe6e6b748b42a2dccc5e9d"),
		PingPongDapp: gethcommon.HexToAddress("0xc1f01a6d0e8382f2c5d394923a7e79693354934b"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployLane:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    23753755,
		},
	},
}

var Prod_SepoliaToOptimismGoerli = rhea.EvmDeploymentConfig{
	ChainConfig: Prod_Sepolia,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress("0x365408d655a6cdf8bc668fd6cebe1bb16a403ca6"),
		OffRamp:      gethcommon.HexToAddress("0x0d3299ee55d493b8d9aafc834a6fd5dcbc4a409a"),
		CommitStore:  gethcommon.HexToAddress("0x1576d23f986ecb572a4c839ba6758ca05c1eadc2"),
		PingPongDapp: gethcommon.HexToAddress("0x37b27863a14781acf41b787cf9ec3fb65d1c5885"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployLane:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    3816777,
		},
	},
}

var Prod_OptimismGoerliToSepolia = rhea.EvmDeploymentConfig{
	ChainConfig: Prod_OptimismGoerli,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress("0x64877f0b53e801adeb8d65f9706f7b134b82971c"),
		OffRamp:      gethcommon.HexToAddress("0xdc4606e96c37b877f2c9ddda82104c85a198a82d"),
		CommitStore:  gethcommon.HexToAddress("0xb7019c10bd604768c9cf5b3d086a2e661559a189"),
		PingPongDapp: gethcommon.HexToAddress("0x3b7a30028bf7ce52ad75b0afb142beef02deeecd"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployLane:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    11482670,
		},
	},
}

var Prod_OptimismGoerliToAvaxFuji = rhea.EvmDeploymentConfig{
	ChainConfig: Prod_OptimismGoerli,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress("0xf33303166911ff86ad1a5ea94d459b4ba0ba8cc9"),
		OffRamp:      gethcommon.HexToAddress("0xee8ce182ea0c0edecf06c2a032a17b2058fc5a04"),
		CommitStore:  gethcommon.HexToAddress("0x2980de4ce178bc8bb6840abd2ef0e2a7c8e7272f"),
		PingPongDapp: gethcommon.HexToAddress("0x227c3699f9f0d6d55c38551a7d7feaea82efdd66"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployLane:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    11482157,
		},
	},
}

var Prod_AvaxFujiToOptimismGoerli = rhea.EvmDeploymentConfig{
	ChainConfig: Prod_AvaxFuji,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress("0x097076fbd8573418c77d2600606ad063c0e3cc7c"),
		OffRamp:      gethcommon.HexToAddress("0x0f287140d86335b37ae2ad0707992ecd4202d5b7"),
		CommitStore:  gethcommon.HexToAddress("0x5a7fa03e52628a0a6f0ab637f10ba45b68f9ad33"),
		PingPongDapp: gethcommon.HexToAddress("0x4ce7b0782966d58ebc5e1804ca6de3244dac9ad9"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployLane:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    23755228,
		},
	},
}

var Prod_ArbitrumGoerliToSepolia = rhea.EvmDeploymentConfig{
	ChainConfig: Prod_ArbitrumGoerli,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress("0xc8b93b46bf682c39b3f65aa1c135bc8a95a5e43a"),
		OffRamp:      gethcommon.HexToAddress("0x7a0bb92bc8663abe6296d0162a9b41a2cb2e0358"),
		CommitStore:  gethcommon.HexToAddress("0x7eef73aca8657aaefd509a97ee75aa6740046e75"),
		PingPongDapp: gethcommon.HexToAddress("0x9b451300c94c7328bdb56a514f83205ea789136f"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployLane:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    29170089,
		},
	},
}

var Prod_SepoliaToArbitrumGoerli = rhea.EvmDeploymentConfig{
	ChainConfig: Prod_Sepolia,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress("0x2c0b51f491ceaefe8c24c0199fce62d7b040470a"),
		OffRamp:      gethcommon.HexToAddress("0x1d649a11fa14024f9fa2058a6b5b473ea308b688"),
		CommitStore:  gethcommon.HexToAddress("0xc677d898f06cee7b5f6ecbd0f72df5125cebbfc9"),
		PingPongDapp: gethcommon.HexToAddress("0x65b51ba5c9233465f118285e5fb2110c52ad6b27"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployLane:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    3816651,
		},
	},
}

var Prod_PolygonMumbaiToSepolia = rhea.EvmDeploymentConfig{
	ChainConfig: Prod_PolygonMumbai,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress("0xa83f2cecb391779b59022eded6ebba0d7ec01f20"),
		OffRamp:      gethcommon.HexToAddress("0xbe582db704bd387222c70ca2e5a027e5e2c06fb7"),
		CommitStore:  gethcommon.HexToAddress("0xf06ff5d2084295909119ca541e93635e7d582ffc"),
		PingPongDapp: gethcommon.HexToAddress("0x044a6b4b561af69d2319a2f4be5ec327a6975d0a"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployLane:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    37526794,
		},
	},
}

var Prod_SepoliaToPolygonMumbai = rhea.EvmDeploymentConfig{
	ChainConfig: Prod_Sepolia,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress("0x23438be3256369316e53fd2ef1cd2bdfaf22f6ad"),
		OffRamp:      gethcommon.HexToAddress("0x026fb7c16f1d0082809ff2335715f27e1e074ff6"),
		CommitStore:  gethcommon.HexToAddress("0x290789a55e2e26480f9c04c583d1d5c682aba49a"),
		PingPongDapp: gethcommon.HexToAddress("0xf66fcb898e838a997547ae58fd6882b9bbfdc399"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployLane:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    3819110,
		},
	},
}

var Prod_AvaxFujiToPolygonMumbai = rhea.EvmDeploymentConfig{
	ChainConfig: Prod_AvaxFuji,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress("0x762aabc808270fadfdd9e4186739920d68106673"),
		OffRamp:      gethcommon.HexToAddress("0x31cf2040d53f178d168997c658d1a7fc5fa7d215"),
		CommitStore:  gethcommon.HexToAddress("0xa60821b061116054672d102c0b59290910fb51e2"),
		PingPongDapp: gethcommon.HexToAddress("0xcffb4c676a996daefa2a8a6d404a55f59ecc7ce8"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployLane:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    23754063,
		},
	},
}

var Prod_PolygonMumbaiToAvaxFuji = rhea.EvmDeploymentConfig{
	ChainConfig: Prod_PolygonMumbai,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress("0x48601a62aa4cb289d006ff4c14023e9d8a7e5a88"),
		OffRamp:      gethcommon.HexToAddress("0xf11e96f85e1038c429d32a877e2225d37cde10e2"),
		CommitStore:  gethcommon.HexToAddress("0x6b6b328cb1467d906389a1bbe54359c56000422c"),
		PingPongDapp: gethcommon.HexToAddress("0x3bd38d308045a39253b502f1789e95c703e27f77"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployLane:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    37511834,
		},
	},
}

var Prod_OptimismGoerliToArbitrumGoerli = rhea.EvmDeploymentConfig{
	ChainConfig: Prod_OptimismGoerli,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress("0x6bb8d729c35f29df532eb3998ddace336187c84b"),
		OffRamp:      gethcommon.HexToAddress("0xee55842b1d68224d9eef238d4736e851db613630"),
		CommitStore:  gethcommon.HexToAddress("0x4f57b2d4b3b42f09cd7ef48254d2c31b6b525763"),
		PingPongDapp: gethcommon.HexToAddress("0x2af63f50fa3f97f4aa94d28327a759ca86b33bf8"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployLane:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    11481052,
		},
	},
}

var Prod_ArbitrumGoerliToOptimismGoerli = rhea.EvmDeploymentConfig{
	ChainConfig: Prod_ArbitrumGoerli,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress("0x782a7ba95215f2f7c3dd4c153cbb2ae3ec2d3215"),
		OffRamp:      gethcommon.HexToAddress("0xff4b0c64c50d2d7b444cb28699df03ed4bbaf44f"),
		CommitStore:  gethcommon.HexToAddress("0xb69923dfb790e622084b774b99bb45f68904d6a4"),
		PingPongDapp: gethcommon.HexToAddress("0x57c0059fc3f98aa0a5ce4fb5d2882d81d839e74f"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployLane:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    29167620,
		},
	},
}
