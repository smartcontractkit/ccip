package deployments

import (
	gethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/rhea"
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
		rhea.Sepolia:  Prod_OptimismGoerliToSepolia,
		rhea.AvaxFuji: Prod_OptimismGoerliToAvaxFuji,
	},
	rhea.ArbitrumGoerli: {
		rhea.Sepolia: Prod_ArbitrumGoerliToSepolia,
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
		gethcommon.HexToAddress("0x5fc725ce3857f46dd53615f852333866a1faacdb"), // SepoliaToAvaxFuji.PingPongDapp,
		gethcommon.HexToAddress("0x1b74b67a4a5adf52ee41d33eb5af9005deee595a"), // SepoliaToOptimismGoerli.PingPongDapp,
		gethcommon.HexToAddress("0xf300a0bcb47c7b169f9309615caee69889b05f3f"), // SepoliaToArbitrum.PingPongDapp,
		gethcommon.HexToAddress("0xa546a40b5f6468901bb321b4836a813af42d50ee"), // SepoliaToPolygonMumbai.PingPongDapp,
		// Personal
		gethcommon.HexToAddress("0xEa94AA1318796b5C01a9A37faCBc65423fb2c520"), // Anindita Ghosh
		gethcommon.HexToAddress("0x25D7214ae75F169263921a1cAaf7E6F033210E24"), // Chris Cushman
		gethcommon.HexToAddress("0x498533848239DDc6Bb5Cf7aEF63c97f3f5513ed2"), // Pramod - DApp Sepolia->Fuji
		gethcommon.HexToAddress("0x8e5267453b0aa137Be1Fc976755E6A9bD2a2E029"), // Amine (DevRel) 1
		gethcommon.HexToAddress("0x9d087fC03ae39b088326b67fA3C788236645b717"), // Amine (DevRel) 2
		gethcommon.HexToAddress("0x8fDEA7A82D7861144D027e4eb2acCCf4eB37bb05"), // Andrej Rakic
		gethcommon.HexToAddress("0x208AA722Aca42399eaC5192EE778e4D42f4E5De3"), // Zubin Pratap
		gethcommon.HexToAddress("0x52eE5a881287486573cF5CB5e7E7D92F30b03014"), // Zubin Pratap

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
	},
	SupportedTokens: map[rhea.Token]rhea.EVMBridgedToken{
		rhea.LINK: {
			Token:         gethcommon.HexToAddress("0x779877A7B0D9E8603169DdbD7836e478b4624789"),
			Pool:          gethcommon.HexToAddress("0xc1c76a8c5bfde1be034bbcd930c668726e7c1987"),
			Price:         rhea.LINK.Price(),
			TokenPoolType: rhea.LockRelease,
		},
		rhea.WETH: {
			Token:         gethcommon.HexToAddress("0x097D90c9d3E0B50Ca60e1ae45F6A81010f9FB534"),
			Pool:          gethcommon.HexToAddress("0xc049f9902e580df50438ebeeae87d76c8a0c91f9"),
			Price:         rhea.WETH.Price(),
			TokenPoolType: rhea.FeeTokenOnly,
		},
		rhea.CACHEGOLD: {
			Token:         gethcommon.HexToAddress("0x997BCCAE553112CD023592691d41687a3f1EfA7C"),
			Pool:          gethcommon.HexToAddress("0x85d2616e94ff408967959a3abe957d78d566b234"),
			Price:         rhea.CACHEGOLD.Price(),
			TokenPoolType: rhea.LockRelease,
		},
		rhea.ANZ: {
			Token:         gethcommon.HexToAddress("0x92eA346B7a2AaB84e6AaB03b80E2421eeFB04685"),
			Pool:          gethcommon.HexToAddress("0x3054a06e89d83317e9b15f943da87bfa67979935"),
			Price:         rhea.ANZ.Price(),
			TokenPoolType: rhea.LockRelease,
		},
		rhea.InsurAce: {
			Token:         gethcommon.HexToAddress("0xb7c8bCA891143221a34DB60A26639785C4839040"),
			Pool:          gethcommon.HexToAddress("0x9fd866891732eebd989f52b75b0a21a11c271dc8"),
			Price:         rhea.InsurAce.Price(),
			TokenPoolType: rhea.LockRelease,
		},
		rhea.ZUSD: {
			Token:         gethcommon.HexToAddress("0x09ae935D80E190403C61Cc5d854Fbf6a7b4a559a"),
			Pool:          gethcommon.HexToAddress("0x674fcad1a94f611a6a15b995e05707011c29606b"),
			Price:         rhea.ZUSD.Price(),
			TokenPoolType: rhea.Legacy,
		},
		rhea.STEADY: {
			Token:         gethcommon.HexToAddress("0x82abB1864326A8A7e1A357FFA2270D09CCb867B9"),
			Pool:          gethcommon.HexToAddress("0x546212f0bfa34cbc5ef3c20a7d483fa2d6d7aadc"),
			Price:         rhea.STEADY.Price(),
			TokenPoolType: rhea.LockRelease,
		},
		rhea.SUPER: {
			Token:         gethcommon.HexToAddress("0xCb4B3f72B5b6D0b7072aFDDf18FE61A0d569EC39"),
			Pool:          gethcommon.HexToAddress("0x790967db00f34ca3c6711acf0488f0f321aea6cf"),
			Price:         rhea.SUPER.Price(),
			TokenPoolType: rhea.Legacy,
		},
		rhea.BankToken: {
			Token:         gethcommon.HexToAddress("0x784c400D6fF625051d2f587dC0276E3A1ffD9cda"),
			Pool:          gethcommon.HexToAddress("0x314ba6767d54e4eab62733583deeaeaa9c5f9f24"),
			Price:         rhea.BankToken.Price(),
			TokenPoolType: rhea.LockRelease,
		},
		rhea.BondToken: {
			Token:         gethcommon.HexToAddress("0xF92E4b278380f39fADc24483C7baC61b73EE93F2"),
			Pool:          gethcommon.HexToAddress("0x3173def68e8445e6f0c83d0a014ea49cc091527f"),
			Price:         rhea.BondToken.Price(),
			TokenPoolType: rhea.LockRelease,
		},
		rhea.SNXUSD: {
			Token:         gethcommon.HexToAddress("0x5B33A61Fe23260b55f1Fa9c586001a630C048BF4"),
			Pool:          gethcommon.HexToAddress("0xCCd34b2A4496eB6b764134bD509d04E8261C6242"),
			Price:         rhea.SNXUSD.Price(),
			TokenPoolType: rhea.BurnMint,
			PoolAllowList: []gethcommon.Address{
				gethcommon.HexToAddress("0xda9e8e71bb750a996af33ebb8abb18cd9eb9dc75"),
				gethcommon.HexToAddress("0x2A45BaE1E58AaD3261af187b7dAde90889c039Dc"),
			},
		},
		rhea.FUGAZIUSDC: {
			Token:         gethcommon.HexToAddress("0x832bA6abcAdC68812be372F4ef20aAC268bA20B7"),
			Pool:          gethcommon.HexToAddress("0x72d33705441b03c85a715e2c9a452a1d18ec6b25"),
			Price:         rhea.FUGAZIUSDC.Price(),
			TokenPoolType: rhea.Legacy,
		},
	},
	FeeTokens:     []rhea.Token{rhea.LINK, rhea.WETH},
	WrappedNative: rhea.WETH,
	Router:        gethcommon.HexToAddress("0xa5bd184d05c7535c8a022905558974752e646a88"),
	Afn:           gethcommon.HexToAddress("0xc61c21bc5d89ddc53a222b7f2cb76f2975b73a7f"),
	PriceRegistry: gethcommon.HexToAddress("0xbe70fa4caff9b7a9d93f409d78b82d3c8257fb2b"),
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
	DeploySettings: rhea.ChainDeploySettings{
		DeployAFN:           false,
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
		gethcommon.HexToAddress("0xaf9b895f87043c037551968128a779301fa5669e"), // OptimismGoerliToAvaxFuji.PingPongDapp,
		gethcommon.HexToAddress("0x008e39086d09594e3d8de90101e1a7a18aed1a0c"), // OptimismGoerliToSepolia.PingPongDapp,
		// Personal
		gethcommon.HexToAddress("0xEa94AA1318796b5C01a9A37faCBc65423fb2c520"), // Anindita Ghosh
		gethcommon.HexToAddress("0x8fDEA7A82D7861144D027e4eb2acCCf4eB37bb05"), // Andrej Rakic
		gethcommon.HexToAddress("0x208AA722Aca42399eaC5192EE778e4D42f4E5De3"), // Zubin Pratap
		gethcommon.HexToAddress("0x52eE5a881287486573cF5CB5e7E7D92F30b03014"), // Zubin Pratap

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
	},
	SupportedTokens: map[rhea.Token]rhea.EVMBridgedToken{
		rhea.LINK: {
			Token:         gethcommon.HexToAddress("0xdc2CC710e42857672E7907CF474a69B63B93089f"),
			Pool:          gethcommon.HexToAddress("0xaaf300b21536ce583851442574e08a871ffbe874"),
			Price:         rhea.LINK.Price(),
			TokenPoolType: rhea.LockRelease,
		},
		rhea.WETH: {
			Token:         gethcommon.HexToAddress("0x4200000000000000000000000000000000000006"),
			Pool:          gethcommon.HexToAddress("0xb4de051bfa993dac12a5c276b816133aac7d7951"),
			Price:         rhea.WETH.Price(),
			TokenPoolType: rhea.FeeTokenOnly,
		},
		rhea.CACHEGOLD: {
			Token:         gethcommon.HexToAddress("0xa6446C6f492f31A33bC68249ae59F8871123a777"),
			Pool:          gethcommon.HexToAddress("0x7b3e2db33d32286a16d96f635e1513619aa63a64"),
			Price:         rhea.CACHEGOLD.Price(),
			TokenPoolType: rhea.Legacy,
		},
		rhea.ZUSD: {
			Token:         gethcommon.HexToAddress("0x740ba2E7f25c036ED0b19b83c9Da2cB8D756f9D5"),
			Pool:          gethcommon.HexToAddress("0x77b5794c89f8161b958838f56702f9c9923967da"),
			Price:         rhea.ZUSD.Price(),
			TokenPoolType: rhea.Legacy,
		},
		rhea.STEADY: {
			Token:         gethcommon.HexToAddress("0x615c83D5FEdafAEa641f1cC1a91ea09111EF0158"),
			Pool:          gethcommon.HexToAddress("0xaa9c6a9dd369a4c2ff83661acfb23c3be993ae3e"),
			Price:         rhea.STEADY.Price(),
			TokenPoolType: rhea.Legacy,
		},
	},
	FeeTokens:     []rhea.Token{rhea.LINK, rhea.WETH},
	WrappedNative: rhea.WETH,
	Router:        gethcommon.HexToAddress("0x6a9ccb433615caaf0ef20a9f7f04e339dca8f219"),
	Afn:           gethcommon.HexToAddress("0xbe38f896eab25c19bd5a2b4a5b03982043d6f1ae"),
	PriceRegistry: gethcommon.HexToAddress("0xd1f0bf6f4b448a97785d998e6eca7c2ae45a0cc8"),
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
	DeploySettings: rhea.ChainDeploySettings{
		DeployAFN:           false,
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
		gethcommon.HexToAddress("0x76c1ed1199e3ff88e78c032842b9d758b39d2d19"), // AvaxFujiToSepolia.PingPongDapp,
		gethcommon.HexToAddress("0x0ded743d54f462f0cd4f18fbd416631d97efd4b3"), // AvaxFujiToOptimismGoerli.PingPongDapp,
		gethcommon.HexToAddress("0x6114310b5730ece2ed7558c7a13d1de2eec728ab"), // AvaxFujiToPolygonMumbai.PingPongDapp,
		// Personal
		gethcommon.HexToAddress("0xEa94AA1318796b5C01a9A37faCBc65423fb2c520"), // Anindita Ghosh
		gethcommon.HexToAddress("0x594D8E57D8801069C77AAB90222a9162E908AA63"), // Pramod - Dapp Fuji->OptimismGoerli
		gethcommon.HexToAddress("0xFE5394A63433A3975b1936dEc92DAa161FEE7463"), // Pramod - DApp Fuji->Sepolia
		gethcommon.HexToAddress("0x912519a7E5e2e2309b1e60F540683c6661757A0C"), // Amine (DevRel) 1
		gethcommon.HexToAddress("0x9d087fC03ae39b088326b67fA3C788236645b717"), // Amine (DevRel) 2
		gethcommon.HexToAddress("0x8fDEA7A82D7861144D027e4eb2acCCf4eB37bb05"), // Andrej Rakic
		gethcommon.HexToAddress("0x208AA722Aca42399eaC5192EE778e4D42f4E5De3"), // Zubin Pratap
		gethcommon.HexToAddress("0x52eE5a881287486573cF5CB5e7E7D92F30b03014"), // Zubin Pratap

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
	},
	SupportedTokens: map[rhea.Token]rhea.EVMBridgedToken{
		rhea.LINK: {
			Token:         gethcommon.HexToAddress("0x0b9d5D9136855f6FEc3c0993feE6E9CE8a297846"),
			Pool:          gethcommon.HexToAddress("0xbcb9674b30041a30cc206faa10ffd1d256f0522a"),
			Price:         rhea.LINK.Price(),
			TokenPoolType: rhea.LockRelease,
		},
		rhea.WAVAX: {
			Token:         gethcommon.HexToAddress("0xd00ae08403B9bbb9124bB305C09058E32C39A48c"),
			Pool:          gethcommon.HexToAddress("0xf2b9147c77e67fd0122fc50ac89565a74085638b"),
			Price:         rhea.WAVAX.Price(),
			TokenPoolType: rhea.FeeTokenOnly,
		},
		rhea.CACHEGOLD: {
			Token:         gethcommon.HexToAddress("0xD16eD805F3eCe986d9541afaD3E59De2F3732517"),
			Pool:          gethcommon.HexToAddress("0x00dacc32abcfafa1128076213c32ea9859075f03"),
			Price:         rhea.CACHEGOLD.Price(),
			TokenPoolType: rhea.Legacy,
		},
		rhea.ANZ: {
			Token:         gethcommon.HexToAddress("0xe3d06cb8eac016749281f45e779ac2976baa02ed"),
			Pool:          gethcommon.HexToAddress("0xe3d06cb8eac016749281f45e779ac2976baa02ed"),
			Price:         rhea.ANZ.Price(),
			TokenPoolType: rhea.Wrapped,
		},
		rhea.InsurAce: {
			Token:         gethcommon.HexToAddress("0xda305ab72858939758d5a711494cd447d2d8842e"),
			Pool:          gethcommon.HexToAddress("0xda305ab72858939758d5a711494cd447d2d8842e"),
			Price:         rhea.InsurAce.Price(),
			TokenPoolType: rhea.Wrapped,
		},
		rhea.SUPER: {
			Token:         gethcommon.HexToAddress("0xCb4B3f72B5b6D0b7072aFDDf18FE61A0d569EC39"),
			Pool:          gethcommon.HexToAddress("0xa546a40b5f6468901bb321b4836a813af42d50ee"),
			Price:         rhea.SUPER.Price(),
			TokenPoolType: rhea.Legacy,
		},
		rhea.BankToken: {
			Token:         gethcommon.HexToAddress("0x7130aac4827a8b085ffe701a7d4749e2b452a837"),
			Pool:          gethcommon.HexToAddress("0x7130aac4827a8b085ffe701a7d4749e2b452a837"),
			Price:         rhea.BankToken.Price(),
			TokenPoolType: rhea.Wrapped,
		},
		rhea.BondToken: {
			Token:         gethcommon.HexToAddress("0x56e01ecb119c45ff14248f6ebc27c05d4a72d4f9"),
			Pool:          gethcommon.HexToAddress("0x56e01ecb119c45ff14248f6ebc27c05d4a72d4f9"),
			Price:         rhea.BondToken.Price(),
			TokenPoolType: rhea.Wrapped,
		},
		rhea.FUGAZIUSDC: {
			Token:         gethcommon.HexToAddress("0x150a0ee7393294442EE4d4F5C7d637af01dF93ee"),
			Pool:          gethcommon.HexToAddress("0xbf9073bc263c880b15ecda52a4f79f3c7c4d6ac0"),
			Price:         rhea.FUGAZIUSDC.Price(),
			TokenPoolType: rhea.Legacy,
		},
	},
	FeeTokens:     []rhea.Token{rhea.LINK, rhea.WAVAX},
	WrappedNative: rhea.WAVAX,
	Router:        gethcommon.HexToAddress("0x9b45eda197971e5fc1eba5b51e6c8b3b9f2578cc"),
	Afn:           gethcommon.HexToAddress("0x943e33dc8a45d666fe4cce76a8355574b282d1c5"),
	PriceRegistry: gethcommon.HexToAddress("0x3b7a30028bf7ce52ad75b0afb142beef02deeecd"),
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
	DeploySettings: rhea.ChainDeploySettings{
		DeployAFN:           false,
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
		gethcommon.HexToAddress("0x7854e73c73e7f9bb5b0d5b4861e997f4c6e8dcc6"), // ArbitrumGoerliToSepolia.PingPongDapp,
		// Personal
		gethcommon.HexToAddress("0x8fDEA7A82D7861144D027e4eb2acCCf4eB37bb05"), // Andrej Rakic
		gethcommon.HexToAddress("0x208AA722Aca42399eaC5192EE778e4D42f4E5De3"), // Zubin Pratap
		gethcommon.HexToAddress("0x52eE5a881287486573cF5CB5e7E7D92F30b03014"), // Zubin Pratap
		// ==============  EXTERNAL ==============
		gethcommon.HexToAddress("0xF5022eDd1B827E6EA4bBdb961212ECD7F315ed88"), // BetaUser - RiseWorks
		gethcommon.HexToAddress("0x0D7a3a17E2E160287D3e7e74c4A1B22422156642"), // BetaUser - RiseWorks
		gethcommon.HexToAddress("0x63e430dBd88C1bBFBc97336b4357Aa5Aea83367e"), // BetaUser - RiseWorks
	},
	SupportedTokens: map[rhea.Token]rhea.EVMBridgedToken{
		rhea.LINK: {
			Token:         gethcommon.HexToAddress("0xd14838A68E8AFBAdE5efb411d5871ea0011AFd28"),
			Pool:          gethcommon.HexToAddress("0x2f81c1003366249f1fd94127f5d9527c9da30dfd"),
			Price:         rhea.LINK.Price(),
			TokenPoolType: rhea.LockRelease,
		},
		rhea.WETH: {
			Token:         gethcommon.HexToAddress("0x32d5D5978905d9c6c2D4C417F0E06Fe768a4FB5a"),
			Pool:          gethcommon.HexToAddress("0xb339c3ee63dfedf0eab481700417db6848a2e66a"),
			Price:         rhea.WETH.Price(),
			TokenPoolType: rhea.FeeTokenOnly,
		},
	},
	FeeTokens:     []rhea.Token{rhea.LINK, rhea.WETH},
	WrappedNative: rhea.WETH,
	Router:        gethcommon.HexToAddress("0xf9b7595d64a380ffa605a1d11bff5cd629fb7189"),
	Afn:           gethcommon.HexToAddress("0x0d7ef99c39f5d6669cc19e8443b22e92ec5225c8"),
	PriceRegistry: gethcommon.HexToAddress("0x6fe3e48057a39311725964840af164c890752a47"),
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
	DeploySettings: rhea.ChainDeploySettings{
		DeployAFN:           false,
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
		gethcommon.HexToAddress("0xdaad0c8b4b3030b9b866fa4528abd9e1eec9082b"), // PolygonMumbaiToSepolia.PingPongDapp,
		gethcommon.HexToAddress("0xa237b21fb55eecbc6edb76017e3dab5c9587173e"), // PolygonMumbaiToAvax.PingPongDapp,
		// Personal
		gethcommon.HexToAddress("0xEa94AA1318796b5C01a9A37faCBc65423fb2c520"), // Anindita Ghosh
		gethcommon.HexToAddress("0x8fDEA7A82D7861144D027e4eb2acCCf4eB37bb05"), // Andrej Rakic
		gethcommon.HexToAddress("0x208AA722Aca42399eaC5192EE778e4D42f4E5De3"), // Zubin Pratap
		gethcommon.HexToAddress("0x52eE5a881287486573cF5CB5e7E7D92F30b03014"), // Zubin Pratap
		// ==============  EXTERNAL ==============
		gethcommon.HexToAddress("0xe764C455e3Bd05Eb7Cf53Ec8491dca0e91486D24"), // BetaUser - Synthetix v3 core
		gethcommon.HexToAddress("0x8e52262f91ef7049adfD8d1E608172fAC57995c3"), // BetaUser - Synthetix v3 core
		gethcommon.HexToAddress("0x6De1e981d2137f7839840e2140dBB3A05F05B770"), // BetaUser - Flash Liquidity
		gethcommon.HexToAddress("0x2A45BaE1E58AaD3261af187b7dAde90889c039Dc"), // BetaUser - Synthetix v3 core
	},
	SupportedTokens: map[rhea.Token]rhea.EVMBridgedToken{
		rhea.LINK: {
			Token:         gethcommon.HexToAddress("0x326C977E6efc84E512bB9C30f76E30c160eD06FB"),
			Pool:          gethcommon.HexToAddress("0x4c10d67e4b8e18a67a7606defdce42ccc281d39b"),
			Price:         rhea.LINK.Price(),
			TokenPoolType: rhea.LockRelease,
		},
		rhea.WMATIC: {
			Token:         gethcommon.HexToAddress("0x9c3C9283D3e44854697Cd22D3Faa240Cfb032889"),
			Price:         rhea.WMATIC.Price(),
			TokenPoolType: rhea.FeeTokenOnly,
		},
		rhea.SNXUSD: {
			Token:         gethcommon.HexToAddress("0x5B33A61Fe23260b55f1Fa9c586001a630C048BF4"),
			Pool:          gethcommon.HexToAddress("0x4D5F14561e949127Bf435d2Ac884f512F4C058FC"),
			Price:         rhea.SNXUSD.Price(),
			TokenPoolType: rhea.BurnMint,
			PoolAllowList: []gethcommon.Address{
				gethcommon.HexToAddress("0xda9e8e71bb750a996af33ebb8abb18cd9eb9dc75"),
				gethcommon.HexToAddress("0x2A45BaE1E58AaD3261af187b7dAde90889c039Dc"),
			},
		},
	},
	FeeTokens:     []rhea.Token{rhea.LINK, rhea.WMATIC},
	WrappedNative: rhea.WMATIC,
	Router:        gethcommon.HexToAddress("0x8a710bbd77661d168d5a6725bd2e514ba1bff59d"),
	Afn:           gethcommon.HexToAddress("0xd6b8378092f590a39c360e8196101290551a66ea"),
	PriceRegistry: gethcommon.HexToAddress("0x1e590a4066e6747ae099078264d9449fc10de85e"),
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
	DeploySettings: rhea.ChainDeploySettings{
		DeployAFN:           false,
		DeployTokenPools:    false,
		DeployRouter:        false,
		DeployPriceRegistry: false,
	},
}

//var Prod_Quorum = rhea.EVMChainConfig{
//	EvmChainId: 1337,
//	GasSettings: rhea.EVMGasSettings{
//		EIP1559:   false,
//		GasTipCap: rhea.DefaultGasTipFee,
//	},
//	AllowList: []gethcommon.Address{},
//	SupportedTokens: map[rhea.Token]rhea.EVMBridgedToken{
//		rhea.LINK: {
//			Token:         gethcommon.HexToAddress("0x8a710bBd77661D168D5A6725bD2E514ba1bFf59d"),
//			Pool:          gethcommon.HexToAddress("0x3df2913ec50702957ce34f1d22ffe98edd9efefc"),
//			Price:         rhea.LINK.Price(),
//			TokenPoolType: rhea.LockRelease,
//		},
//		rhea.WETH: {
//			Token:         gethcommon.HexToAddress("0xe5BD6BCb6fa8e5236984D5ea127dE5047f93B5fF"),
//			Price:         rhea.WETH.Price(),
//			TokenPoolType: rhea.FeeTokenOnly,
//		},
//		rhea.BankToken: {
//			Token:         gethcommon.HexToAddress("0x1041d5b06b9f1e73a4d6ae4f4d2a3d5c9e0aa88e"),
//			Pool:          gethcommon.HexToAddress("0x1041d5b06b9f1e73a4d6ae4f4d2a3d5c9e0aa88e"),
//			Price:         rhea.BankToken.Price(),
//			TokenPoolType: rhea.Wrapped,
//		},
//		rhea.BondToken: {
//			Token:         gethcommon.HexToAddress("0x98c73259170aa9bd680ff897eaf93955b2902955"),
//			Pool:          gethcommon.HexToAddress("0x98c73259170aa9bd680ff897eaf93955b2902955"),
//			Price:         rhea.BondToken.Price(),
//			TokenPoolType: rhea.Wrapped,
//		},
//	},
//	FeeTokens:     []rhea.Token{rhea.LINK, rhea.WETH},
//	WrappedNative: rhea.WETH,
//	Router:        gethcommon.HexToAddress("0xd74033b0466e26ac984c6165dd0dc017ca05ff91"),
//	Afn:           gethcommon.HexToAddress("0x10ea990fc39f4a140f107fc87bf7d4da3fd37095"),
//	PriceRegistry: gethcommon.HexToAddress("0x632ad295b9cb955d70feb33e6a3ee2bddbf0582d"),
//	TunableChainValues: rhea.TunableChainValues{
//		FinalityDepth:            getFinalityDepth(rhea.Quorum),
//		OptimisticConfirmations:  getOptimisticConfirmations(rhea.Quorum),
//		BatchGasLimit:            BATCH_GAS_LIMIT,
//		RelativeBoostPerWaitHour: RELATIVE_BOOST_PER_WAIT_HOUR,
//		FeeUpdateHeartBeat:       models.MustMakeDuration(FEE_UPDATE_HEARTBEAT),
//		FeeUpdateDeviationPPB:    FEE_UPDATE_DEVIATION_PPB_FAST_CHAIN,
//		MaxGasPrice:              getMaxGasPrice(rhea.Quorum),
//		InflightCacheExpiry:      models.MustMakeDuration(INFLIGHT_CACHE_EXPIRY),
//		RootSnoozeTime:           models.MustMakeDuration(ROOT_SNOOZE_TIME),
//	},
//	DeploySettings: rhea.ChainDeploySettings{
//		DeployAFN:           false,
//		DeployTokenPools:    false,
//		DeployRouter:        false,
//		DeployPriceRegistry: false,
//	},
//}

var Prod_SepoliaToAvaxFuji = rhea.EvmDeploymentConfig{
	ChainConfig: Prod_Sepolia,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress("0x667035db697ef3d64c6cc1aecb93b4ad98aed507"),
		OffRamp:      gethcommon.HexToAddress("0x532030a8b3a966c5ddc9c1c0afafef7a3fd96e06"),
		CommitStore:  gethcommon.HexToAddress("0xb7aff55a785ad727a6f77850ef84786381a63617"),
		PingPongDapp: gethcommon.HexToAddress("0x5fc725ce3857f46dd53615f852333866a1faacdb"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    3655211,
		},
	},
}

var Prod_AvaxFujiToSepolia = rhea.EvmDeploymentConfig{
	ChainConfig: Prod_AvaxFuji,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress("0x39c73ea5356633750c46ae92d009368d31bcb12a"),
		OffRamp:      gethcommon.HexToAddress("0x9f8d452b2b9a0a93193ce84c352b09ea7b5e3048"),
		CommitStore:  gethcommon.HexToAddress("0x9228ba7390b47ed490676e9dd77865d5d6c7356c"),
		PingPongDapp: gethcommon.HexToAddress("0x76c1ed1199e3ff88e78c032842b9d758b39d2d19"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    22876561,
		},
	},
}

var Prod_SepoliaToOptimismGoerli = rhea.EvmDeploymentConfig{
	ChainConfig: Prod_Sepolia,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress("0x9297aab356e7d387080b856b20231ccea5d33bfe"),
		OffRamp:      gethcommon.HexToAddress("0x5604335ae01d8318100b5e08c85c6b8df7e13b1b"),
		CommitStore:  gethcommon.HexToAddress("0x04953b444da6416c664f80b3c441ecc3e37169f6"),
		PingPongDapp: gethcommon.HexToAddress("0x1b74b67a4a5adf52ee41d33eb5af9005deee595a"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    3655106,
		},
	},
}

var Prod_OptimismGoerliToSepolia = rhea.EvmDeploymentConfig{
	ChainConfig: Prod_OptimismGoerli,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress("0xc015d280fbd26211bac450e95e56ffbba1fd49eb"),
		OffRamp:      gethcommon.HexToAddress("0xda143f0bace5aa90604b235e5eea44f67cb3d4b9"),
		CommitStore:  gethcommon.HexToAddress("0x133dd8a18088688d79be5a47619714b6cd4df297"),
		PingPongDapp: gethcommon.HexToAddress("0x008e39086d09594e3d8de90101e1a7a18aed1a0c"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    10439785,
		},
	},
}

var Prod_OptimismGoerliToAvaxFuji = rhea.EvmDeploymentConfig{
	ChainConfig: Prod_OptimismGoerli,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress("0x15c6e8b043a530299cb369abcade9c2a9f5df181"),
		OffRamp:      gethcommon.HexToAddress("0xa356123bd67657fe26b176f179d6cf6ea96f6aea"),
		CommitStore:  gethcommon.HexToAddress("0x5484f1d6d6780ec11062f6d347ed155b8719386c"),
		PingPongDapp: gethcommon.HexToAddress("0xaf9b895f87043c037551968128a779301fa5669e"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    10439147,
		},
	},
}

var Prod_AvaxFujiToOptimismGoerli = rhea.EvmDeploymentConfig{
	ChainConfig: Prod_AvaxFuji,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress("0x9684916b779002642ecf8c6bd62ff9afbb2ab9a2"),
		OffRamp:      gethcommon.HexToAddress("0x366e98df2881543185ece04dacffde417c7d2f1d"),
		CommitStore:  gethcommon.HexToAddress("0x9f9f04e6dd1e95e0edd197923afc20a9c0e5a550"),
		PingPongDapp: gethcommon.HexToAddress("0x0ded743d54f462f0cd4f18fbd416631d97efd4b3"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    22875290,
		},
	},
}

var Prod_ArbitrumGoerliToSepolia = rhea.EvmDeploymentConfig{
	ChainConfig: Prod_ArbitrumGoerli,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress("0x201d1843707764ca2f236bd69e37ccbeff0827d4"),
		OffRamp:      gethcommon.HexToAddress("0x65120af1c7ecaa90294758aafbb87226d2b3b798"),
		CommitStore:  gethcommon.HexToAddress("0xdaebcfbf1a27bd416ebe20207495a9086a52c22d"),
		PingPongDapp: gethcommon.HexToAddress("0x7854e73c73e7f9bb5b0d5b4861e997f4c6e8dcc6"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    25022239,
		},
	},
}

var Prod_SepoliaToArbitrumGoerli = rhea.EvmDeploymentConfig{
	ChainConfig: Prod_Sepolia,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress("0x038cb2d6d3b57dd43fabc8b2d2257b4755d8b3ec"),
		OffRamp:      gethcommon.HexToAddress("0xa1f9a83eaadd801ea91c7121894e8eb9cb54499d"),
		CommitStore:  gethcommon.HexToAddress("0xea2a136c99b1582ce912a131b249aa5ac99c3095"),
		PingPongDapp: gethcommon.HexToAddress("0xf300a0bcb47c7b169f9309615caee69889b05f3f"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    3654960,
		},
	},
}

var Prod_PolygonMumbaiToSepolia = rhea.EvmDeploymentConfig{
	ChainConfig: Prod_PolygonMumbai,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress("0x51298c07ef8849f89552c2b3184741a759d4b37c"),
		OffRamp:      gethcommon.HexToAddress("0x0363be02ebe4e90af03e7c45ff90d4be974032ba"),
		CommitStore:  gethcommon.HexToAddress("0xb5fee1152fe189978dc9914299114ea1c4b1d5f5"),
		PingPongDapp: gethcommon.HexToAddress("0xdaad0c8b4b3030b9b866fa4528abd9e1eec9082b"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    36631143,
		},
	},
}

var Prod_SepoliaToPolygonMumbai = rhea.EvmDeploymentConfig{
	ChainConfig: Prod_Sepolia,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress("0xfe5394a63433a3975b1936dec92daa161fee7463"),
		OffRamp:      gethcommon.HexToAddress("0xd6cdc2fc1e02a757d3216b64ed707ae0285e1654"),
		CommitStore:  gethcommon.HexToAddress("0x0e97fb9a0ef1e39a85577b7fb9da86cdc1bf1b8c"),
		PingPongDapp: gethcommon.HexToAddress("0xa546a40b5f6468901bb321b4836a813af42d50ee"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    3654802,
		},
	},
}

var Prod_AvaxFujiToPolygonMumbai = rhea.EvmDeploymentConfig{
	ChainConfig: Prod_AvaxFuji,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress("0xa58e254bfb1c0967196ef4b5736a965c887c5950"),
		OffRamp:      gethcommon.HexToAddress("0x108ce393b732fcdc7bbd8055f1b71c00e61ad894"),
		CommitStore:  gethcommon.HexToAddress("0x9d86ab7bf10a22e6c3721e8951f900fed238ab72"),
		PingPongDapp: gethcommon.HexToAddress("0x6114310b5730ece2ed7558c7a13d1de2eec728ab"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    22874574,
		},
	},
}

var Prod_PolygonMumbaiToAvaxFuji = rhea.EvmDeploymentConfig{
	ChainConfig: Prod_PolygonMumbai,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress("0x91002494008dcf53bfb59924cc1eb29f607436e8"),
		OffRamp:      gethcommon.HexToAddress("0xc5ccb84c3d8ead52c081ddb24e7add615c0c9daf"),
		CommitStore:  gethcommon.HexToAddress("0x405770cd319120175ea929bf305795af5ac5bea1"),
		PingPongDapp: gethcommon.HexToAddress("0xa237b21fb55eecbc6edb76017e3dab5c9587173e"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    36631628,
		},
	},
}
