package wallet

import (
	"github.com/LinX-OpenNetwork/coinutil/bip32"
	"github.com/LinX-OpenNetwork/coinutil/bip39"
	"github.com/LinX-OpenNetwork/coinutil/bip44"
)

// HDWallet represents a Hierarchical deterministic wallet
type HDWallet struct {
	seed []byte
	key  *bip32.Key
}

// NewHDWalletPassphrase creates a HDWallet with specified mnemonic and a passphrase
func NewHDWalletPassphrase(mnemonic *bip39.Mnemonic, passphrase string) (*HDWallet, error) {
	seed := mnemonic.GenerateSeed(passphrase)
	masterKey, err := bip32.NewMasterKey(seed.Bytes)
	if err != nil {
		return nil, err
	}

	return &HDWallet{seed: seed.Bytes, key: masterKey}, nil
}

// NewHDWallet creates a HDWallet with specified mnemonic and a passphrase
func NewHDWallet(mnemonic *bip39.Mnemonic) (*HDWallet, error) {
	return NewHDWalletPassphrase(mnemonic, "")
}

// KeyForDerivePath derive a private key from specified DerivePath.
func (m *HDWallet) KeyForDerivePath(path bip44.DerivePath) (*bip32.Key, error) {
	params, err := path.ToParams()
	if err != nil {
		return nil, err
	}
	return m.KeyForDeriveParams(params)
}

// KeyForDeriveParams derive a private key from specified DerivePathParams.
func (m *HDWallet) KeyForDeriveParams(params *bip44.DerivePathParams) (key *bip32.Key, err error) {
	indexes := params.Indexes()
	key = m.key
	for _, idx := range indexes {
		key, err = key.NewChildKey(idx)
		if err != nil {
			return nil, err
		}
	}
	return
}

// KeyForCoin derive a private key from coin, account, change and address index.
func (m *HDWallet) KeyForCoin(coin, account, change, address uint32) (key *bip32.Key, err error) {
	return m.KeyForDeriveParams(&bip44.DerivePathParams{
		Purpose: 44, CoinType: coin, Account: account, Change: change, AddressIndex: address,
	})
}
