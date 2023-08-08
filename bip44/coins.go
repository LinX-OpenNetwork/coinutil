package bip44

const (
	// DeterministicWalletsPurpose Purpose 44'
	DeterministicWalletsPurpose = uint32(0x8000002C)

	// CoinTypeBitcoin Coin type 0'
	CoinTypeBitcoin = uint32(0x80000000)
	// CoinTypeEther Coin type 60'
	CoinTypeEther = uint32(0x8000003c)
	// CoinTypeSolana Coin type 501'
	CoinTypeSolana = uint32(0x800001f5)

	// Bitcoin the derive key path of Bitcoin
	Bitcoin = DerivePath("m/44’/0’/0’/0/0")
	// Ethereum the derive key path of Ethereum
	Ethereum = DerivePath("m/44'/60'/0'/0/0")

	Solana = DerivePath("m/44'/501'/0'/0'/0")
)
