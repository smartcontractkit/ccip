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
	rhea.Quorum:         {ChainConfig: Prod_Quorum},
}

var ProdChainMapping = map[rhea.Chain]map[rhea.Chain]rhea.EvmDeploymentConfig{
	rhea.Sepolia: {
		rhea.AvaxFuji:       Prod_SepoliaToAvaxFuji,
		rhea.OptimismGoerli: Prod_SepoliaToOptimismGoerli,
		rhea.ArbitrumGoerli: Prod_SepoliaToArbitrumGoerli,
		rhea.PolygonMumbai:  Prod_SepoliaToPolygonMumbai,
		rhea.Quorum:         Prod_SepoliaToQuorum,
	},
	rhea.AvaxFuji: {
		rhea.Sepolia:        Prod_AvaxFujiToSepolia,
		rhea.OptimismGoerli: Prod_AvaxFujiToOptimismGoerli,
	},
	rhea.OptimismGoerli: {
		rhea.Sepolia:  Prod_OptimismGoerliToSepolia,
		rhea.AvaxFuji: Prod_OptimismGoerliToAvaxFuji,
	},
	rhea.ArbitrumGoerli: {
		rhea.Sepolia: Prod_ArbitrumGoerliToSepolia,
	},
	rhea.PolygonMumbai: {
		rhea.Sepolia: Prod_PolygonMumbaiToSepolia,
	},
	rhea.Quorum: {
		rhea.Sepolia: Prod_QuorumToSepolia,
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
		gethcommon.HexToAddress("0x4d2bd64a51c84fef0bac9090473fb9b7fe652a66"), // SepoliaToAvaxFuji.PingPongDapp,
		gethcommon.HexToAddress("0x63142ee8aa67fcfe478c8da84a51380a5510f01b"), // SepoliaToOptimismGoerli.PingPongDapp,
		gethcommon.HexToAddress("0x6c1b166f191bb923865647320f560cf329fe4839"), // SepoliaToArbitrum.PingPongDapp,
		gethcommon.HexToAddress("0xc82f067bb4d66a5a878e88ca96b3bc6a40068dc4"), // SepoliaToQuorum.PingPongDapp,
		// Personal
		gethcommon.HexToAddress("0xEa94AA1318796b5C01a9A37faCBc65423fb2c520"), // Anindita Ghosh
		gethcommon.HexToAddress("0x25D7214ae75F169263921a1cAaf7E6F033210E24"), // Chris Cushman
		gethcommon.HexToAddress("0x498533848239DDc6Bb5Cf7aEF63c97f3f5513ed2"), // Pramod - DApp Sepolia->Fuji

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
		gethcommon.HexToAddress("0x87F45de79da4c3356591d74619693E372D525F1b"), // BankToken (BANK) - internal testing contract for SWIFT POC
		gethcommon.HexToAddress("0x844257f146a586dAdD5A3a46cd8c069Bb565a80e"), // BetaUser - SWIFT (BondToken)
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
			TokenPoolType: rhea.BurnMint,
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
			TokenPoolType: rhea.BurnMint,
		},
	},
	FeeTokens:     []rhea.Token{rhea.LINK, rhea.WETH},
	WrappedNative: rhea.WETH,
	Router:        gethcommon.HexToAddress("0xa5bd184d05c7535c8a022905558974752e646a88"),
	Afn:           gethcommon.HexToAddress("0x21fcd6874124c893d001949e182203b24c58c8fa"),
	PriceRegistry: gethcommon.HexToAddress("0x0466a27cfd908f179e4dccdc3e3ee13102650e2d"),
	TunableChainValues: rhea.TunableChainValues{
		BlockConfirmations:       getBlockConfirmations(rhea.Sepolia),
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
		gethcommon.HexToAddress("0x9ec0b177259c98df498096096dbbbef1696fa58e"), // OptimismGoerliToAvaxFuji.PingPongDapp,
		gethcommon.HexToAddress("0x63a646443180ba96888e2e3f1a047658612c9d67"), // OptimismGoerliToSepolia.PingPongDapp,
		// Personal
		gethcommon.HexToAddress("0xEa94AA1318796b5C01a9A37faCBc65423fb2c520"), // Anindita Ghosh

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
			TokenPoolType: rhea.BurnMint,
		},
		rhea.ZUSD: {
			Token:         gethcommon.HexToAddress("0x740ba2E7f25c036ED0b19b83c9Da2cB8D756f9D5"),
			Pool:          gethcommon.HexToAddress("0x77b5794c89f8161b958838f56702f9c9923967da"),
			Price:         rhea.ZUSD.Price(),
			TokenPoolType: rhea.BurnMint,
		},
		rhea.STEADY: {
			Token:         gethcommon.HexToAddress("0x615c83D5FEdafAEa641f1cC1a91ea09111EF0158"),
			Pool:          gethcommon.HexToAddress("0xaa9c6a9dd369a4c2ff83661acfb23c3be993ae3e"),
			Price:         rhea.STEADY.Price(),
			TokenPoolType: rhea.BurnMint,
		},
	},
	FeeTokens:     []rhea.Token{rhea.LINK, rhea.WETH},
	WrappedNative: rhea.WETH,
	Router:        gethcommon.HexToAddress("0x6a9ccb433615caaf0ef20a9f7f04e339dca8f219"),
	Afn:           gethcommon.HexToAddress("0x2b09ca26549d6afc37c1fbfd29c94161e157fa84"),
	PriceRegistry: gethcommon.HexToAddress("0x84dd647ce149c1ada3afb5ceb441789ab80c5bad"),
	TunableChainValues: rhea.TunableChainValues{
		BlockConfirmations:       getBlockConfirmations(rhea.OptimismGoerli),
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
		gethcommon.HexToAddress("0x01d45e872d24f3ffa4693c0a4aaeb3008ed972ae"), // AvaxFujiToSepolia.PingPongDapp,
		gethcommon.HexToAddress("0x8ec5ad6cb5496dae52d2e80cd59fa78d066626d6"), // AvaxFujiToOptimismGoerli.PingPongDapp,
		// Personal
		gethcommon.HexToAddress("0xEa94AA1318796b5C01a9A37faCBc65423fb2c520"), // Anindita Ghosh
		gethcommon.HexToAddress("0x594D8E57D8801069C77AAB90222a9162E908AA63"), // Pramod - Dapp Fuji->OptimismGoerli
		gethcommon.HexToAddress("0xFE5394A63433A3975b1936dEc92DAa161FEE7463"), // Pramod - DApp Fuji->Sepolia

		// ==============  EXTERNAL ==============
		gethcommon.HexToAddress("0x1b5D803Be089e43110Faf54c6b4eC40409Cc7450"), // BetaUser - Multimedia
		gethcommon.HexToAddress("0xE8Cc2Bd6082387a7AC749176b1Fe19377f420740"), // BetaUser - Multimedia (AA wallet)
		gethcommon.HexToAddress("0xa78ceF54da82D6279b20457F4D46294AfF59C871"), // BetaUser - Flash Liquidity
		gethcommon.HexToAddress("0x6613fd61bbfEF3291f2D7C7203Ceab212e880DbB"), // BetaUser - Flash Liquidity
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
		gethcommon.HexToAddress("0x9344AeA9b3270d51c9603d3054E421386dFaacB8"), // BetaUser - Fidelity Digital Assets
		gethcommon.HexToAddress("0x89Eccc61B2d35eACCe08284CF22c2D6487B80A3A"), // BetaUser - Robolabs
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
			TokenPoolType: rhea.BurnMint,
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
			TokenPoolType: rhea.BurnMint,
		},
	},
	FeeTokens:     []rhea.Token{rhea.LINK, rhea.WAVAX},
	WrappedNative: rhea.WAVAX,
	Router:        gethcommon.HexToAddress("0x9b45eda197971e5fc1eba5b51e6c8b3b9f2578cc"),
	Afn:           gethcommon.HexToAddress("0x49f70545287b480ac0b3c798f3dabb5b40cb7561"),
	PriceRegistry: gethcommon.HexToAddress("0x02b4cf36b51ac3e0abed26c1e6a5aee11c624117"),
	TunableChainValues: rhea.TunableChainValues{
		BlockConfirmations:       getBlockConfirmations(rhea.AvaxFuji),
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
		gethcommon.HexToAddress("0x13c2d66a8023e2feb9cbe28e26f32b8d2dae3bd0"), // ArbitrumGoerliToSepolia.PingPongDapp,

		// ==============  EXTERNAL ==============
		gethcommon.HexToAddress("0xF5022eDd1B827E6EA4bBdb961212ECD7F315ed88"), // BetaUser - RiseWorks
		gethcommon.HexToAddress("0x0D7a3a17E2E160287D3e7e74c4A1B22422156642"), // BetaUser - RiseWorks
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
	Afn:           gethcommon.HexToAddress("0xa237b21fb55eecbc6edb76017e3dab5c9587173e"),
	PriceRegistry: gethcommon.HexToAddress("0xbacf5cb76b2abc6b754bcffae8209c76bae731aa"),
	TunableChainValues: rhea.TunableChainValues{
		BlockConfirmations:       getBlockConfirmations(rhea.ArbitrumGoerli),
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
	AllowList: []gethcommon.Address{},
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
	},
	FeeTokens:     []rhea.Token{rhea.LINK, rhea.WMATIC},
	WrappedNative: rhea.WMATIC,
	Router:        gethcommon.HexToAddress("0x8a710bbd77661d168d5a6725bd2e514ba1bff59d"),
	Afn:           gethcommon.HexToAddress("0xdca67fd8324990792c0bfae95903b8a64097754f"),
	PriceRegistry: gethcommon.HexToAddress("0x9527e2d01a3064ef6b50c1da1c0cc523803bcff2"),
	TunableChainValues: rhea.TunableChainValues{
		BlockConfirmations:       getBlockConfirmations(rhea.PolygonMumbai),
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

var Prod_Quorum = rhea.EVMChainConfig{
	EvmChainId: 1337,
	GasSettings: rhea.EVMGasSettings{
		EIP1559:   false,
		GasTipCap: rhea.DefaultGasTipFee,
	},
	AllowList: []gethcommon.Address{},
	SupportedTokens: map[rhea.Token]rhea.EVMBridgedToken{
		rhea.LINK: {
			Token:         gethcommon.HexToAddress("0x8a710bBd77661D168D5A6725bD2E514ba1bFf59d"),
			Pool:          gethcommon.HexToAddress("0x3df2913ec50702957ce34f1d22ffe98edd9efefc"),
			Price:         rhea.LINK.Price(),
			TokenPoolType: rhea.LockRelease,
		},
		rhea.WETH: {
			Token:         gethcommon.HexToAddress("0xe5BD6BCb6fa8e5236984D5ea127dE5047f93B5fF"),
			Price:         rhea.WETH.Price(),
			TokenPoolType: rhea.FeeTokenOnly,
		},
	},
	FeeTokens:     []rhea.Token{rhea.LINK, rhea.WETH},
	WrappedNative: rhea.WETH,
	Router:        gethcommon.HexToAddress("0xd74033b0466e26ac984c6165dd0dc017ca05ff91"),
	Afn:           gethcommon.HexToAddress("0x10ea990fc39f4a140f107fc87bf7d4da3fd37095"),
	PriceRegistry: gethcommon.HexToAddress("0x632ad295b9cb955d70feb33e6a3ee2bddbf0582d"),
	TunableChainValues: rhea.TunableChainValues{
		BlockConfirmations:       getBlockConfirmations(rhea.Quorum),
		BatchGasLimit:            BATCH_GAS_LIMIT,
		RelativeBoostPerWaitHour: RELATIVE_BOOST_PER_WAIT_HOUR,
		FeeUpdateHeartBeat:       models.MustMakeDuration(FEE_UPDATE_HEARTBEAT),
		FeeUpdateDeviationPPB:    FEE_UPDATE_DEVIATION_PPB_FAST_CHAIN,
		MaxGasPrice:              getMaxGasPrice(rhea.Quorum),
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

var Prod_SepoliaToAvaxFuji = rhea.EvmDeploymentConfig{
	ChainConfig: Prod_Sepolia,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress("0x095c837a00eb3323b57849ac1950352172f81f9f"),
		OffRamp:      gethcommon.HexToAddress("0xb0e7f0fccd3c961c473e7c44d939c1cdb4cec1cb"),
		CommitStore:  gethcommon.HexToAddress("0x4b56d8d53f1a6e0117b09700067de99581aa5542"),
		PingPongDapp: gethcommon.HexToAddress("0x4d2bd64a51c84fef0bac9090473fb9b7fe652a66"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    3491247,
		},
	},
}

var Prod_AvaxFujiToSepolia = rhea.EvmDeploymentConfig{
	ChainConfig: Prod_AvaxFuji,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress("0xd58fedcb99e2d1274203489801695f9392713877"),
		OffRamp:      gethcommon.HexToAddress("0x8973c9c10ca2fcb7b3bde3253f5ac7b290425d0a"),
		CommitStore:  gethcommon.HexToAddress("0x5eadd4ed3b0e80f95343609226b7b7dd13197224"),
		PingPongDapp: gethcommon.HexToAddress("0x01d45e872d24f3ffa4693c0a4aaeb3008ed972ae"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    21936491,
		},
	},
}

var Prod_SepoliaToOptimismGoerli = rhea.EvmDeploymentConfig{
	ChainConfig: Prod_Sepolia,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress("0x5fcc6941d1685c7115e257cdcbda258af85f0c83"),
		OffRamp:      gethcommon.HexToAddress("0x8815e7090090c90f56e2eb5e79c3ff8fadf4815e"),
		CommitStore:  gethcommon.HexToAddress("0x0e1136cc3a2147ca178d265ae336602217988f48"),
		PingPongDapp: gethcommon.HexToAddress("0x63142ee8aa67fcfe478c8da84a51380a5510f01b"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    3491307,
		},
	},
}

var Prod_OptimismGoerliToSepolia = rhea.EvmDeploymentConfig{
	ChainConfig: Prod_OptimismGoerli,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress("0xaa0e745da1711a7a0699f65988b9bb4b9539c3d6"),
		OffRamp:      gethcommon.HexToAddress("0x68ff78c02af20a71054cd6cdfa685f62f9bbd375"),
		CommitStore:  gethcommon.HexToAddress("0xbf38262aab8bfe63fe14fcbe67573fb2c270e143"),
		PingPongDapp: gethcommon.HexToAddress("0x63a646443180ba96888e2e3f1a047658612c9d67"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    9366229,
		},
	},
}

var Prod_OptimismGoerliToAvaxFuji = rhea.EvmDeploymentConfig{
	ChainConfig: Prod_OptimismGoerli,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress("0x37fb5ae9f8e95879a1d67d5bd451b5d1358aadc8"),
		OffRamp:      gethcommon.HexToAddress("0xd306412fd23a797b3ff73c3bf846c8fd70bb58a4"),
		CommitStore:  gethcommon.HexToAddress("0xce16b4e4acdae2d96c5f25baae1ff3e17c244fe1"),
		PingPongDapp: gethcommon.HexToAddress("0x9ec0b177259c98df498096096dbbbef1696fa58e"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    9366543,
		},
	},
}

var Prod_AvaxFujiToOptimismGoerli = rhea.EvmDeploymentConfig{
	ChainConfig: Prod_AvaxFuji,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress("0x14a4665dc1b0b5e6b00b1c7dd6c83eafa35b01d5"),
		OffRamp:      gethcommon.HexToAddress("0xf300a0bcb47c7b169f9309615caee69889b05f3f"),
		CommitStore:  gethcommon.HexToAddress("0x48f449bf38bfcdd7236bb28f260124fb222d22ac"),
		PingPongDapp: gethcommon.HexToAddress("0x8ec5ad6cb5496dae52d2e80cd59fa78d066626d6"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    21937135,
		},
	},
}

var Prod_ArbitrumGoerliToSepolia = rhea.EvmDeploymentConfig{
	ChainConfig: Prod_ArbitrumGoerli,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress("0x17e0950f3358a3d6735baea5c9d9c9e4c377a111"),
		OffRamp:      gethcommon.HexToAddress("0xbfa2acd33ed6eec0ed3cc06bf1ac38d22b36b9e9"),
		CommitStore:  gethcommon.HexToAddress("0x86000bff3465c579dba5703b2dba6117ce022576"),
		PingPongDapp: gethcommon.HexToAddress("0x13c2d66a8023e2feb9cbe28e26f32b8d2dae3bd0"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    19905564,
		},
	},
}

var Prod_SepoliaToArbitrumGoerli = rhea.EvmDeploymentConfig{
	ChainConfig: Prod_Sepolia,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress("0xc277d7d76861a51c0782ac5b6b7fd61cfe30132b"),
		OffRamp:      gethcommon.HexToAddress("0xa9de3f7a617d67bc50c56baacb9e0373c15ebfc6"),
		CommitStore:  gethcommon.HexToAddress("0x652285058b413aa3abf1e8c50a0e074b3ddf9de4"),
		PingPongDapp: gethcommon.HexToAddress("0x6c1b166f191bb923865647320f560cf329fe4839"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    3491412,
		},
	},
}

var Prod_PolygonMumbaiToSepolia = rhea.EvmDeploymentConfig{
	ChainConfig: Prod_PolygonMumbai,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress("0xa8bfa4fc9f97793297793291e7f63da18eccbe05"),
		OffRamp:      gethcommon.HexToAddress("0x5d3ced1297d948658b104e70eeb9d2594a7e9b1a"),
		CommitStore:  gethcommon.HexToAddress("0x001a05467dd427b538126f7221f44183a23a94bc"),
		PingPongDapp: gethcommon.HexToAddress("0x56edc4d8367932f0e36b966cbbd95df48e9db40f"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    35662240,
		},
	},
}

var Prod_SepoliaToPolygonMumbai = rhea.EvmDeploymentConfig{
	ChainConfig: Prod_Sepolia,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress("0x1d4a253fcb48a7fd433d305ed426a94767e464b1"),
		OffRamp:      gethcommon.HexToAddress("0x6659611bae71a6763644f8c74c16426a4737e925"),
		CommitStore:  gethcommon.HexToAddress("0x75366f24921d589a6c47d73fc2bccae1c47af884"),
		PingPongDapp: gethcommon.HexToAddress("0x92edc8d0408d795095015dfdb50aab59f80110ae"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    3497174,
		},
	},
}

var Prod_QuorumToSepolia = rhea.EvmDeploymentConfig{
	ChainConfig: Prod_Quorum,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress("0x6ea155fc77566d9dce01b8aa5d7968665dc4f0c5"),
		OffRamp:      gethcommon.HexToAddress("0xb602b6e5caf08ac0c920eae585aed100a8cf6f3b"),
		CommitStore:  gethcommon.HexToAddress("0x89d5b13908b9063abcc6791dc724bf7b7c93634c"),
		PingPongDapp: gethcommon.HexToAddress("0x77f10a563eef8996694fb0325ac110a4083b9002"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    77706,
		},
	},
}

var Prod_SepoliaToQuorum = rhea.EvmDeploymentConfig{
	ChainConfig: Prod_Sepolia,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress("0x4f7a6e0045d78630bd49394df292aedbd1631bd6"),
		OffRamp:      gethcommon.HexToAddress("0x43e9295190e3bfe51d6a8568aaa42d4520fc4476"),
		CommitStore:  gethcommon.HexToAddress("0x0a5781e8912c0fa09d2029e221ec7ab392edc7f3"),
		PingPongDapp: gethcommon.HexToAddress("0xc82f067bb4d66a5a878e88ca96b3bc6a40068dc4"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    3515969,
		},
	},
}
