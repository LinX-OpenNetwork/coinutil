package wallet

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/LinX-OpenNetwork/coinutil/bip32"
	"github.com/LinX-OpenNetwork/coinutil/bip44"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// EthereumWallet is a wallet with ethereum private key
type EthereumWallet struct {
	privateKey *ecdsa.PrivateKey
}

func NewEthereumWalletFromHDWallet(hdWallet *HDWallet) (*EthereumWallet, error) {
	key, err := hdWallet.KeyForDerivePath(bip44.Ethereum)
	if err != nil {
		return nil, err
	}
	return NewEthereumWalletFromKey(key)
}

func NewEthereumWalletFromKey(key *bip32.Key) (*EthereumWallet, error) {
	privateKey, err := crypto.ToECDSA(key.Key)
	if err != nil {
		return nil, err
	}
	return &EthereumWallet{privateKey: privateKey}, nil
}

// PrivateKey the private key of this ethereum wallet
func (w *EthereumWallet) PrivateKey() *ecdsa.PrivateKey {
	return w.privateKey
}

// PublicKey the public key of this ethereum wallet
func (w *EthereumWallet) PublicKey() ecdsa.PublicKey {
	return w.privateKey.PublicKey
}

// Address the address of this ethereum wallet
func (w *EthereumWallet) Address() common.Address {
	return crypto.PubkeyToAddress(w.privateKey.PublicKey)
}

// Sign a transaction data
func (w *EthereumWallet) Sign(data []byte) ([]byte, error) {
	return crypto.Sign(data, w.privateKey)
}

// PersonalSign use personal sign to sign any data
func (w *EthereumWallet) PersonalSign(data []byte) ([]byte, error) {
	prefix := []byte(fmt.Sprintf("\x19Ethereum Signed Message:\n%d", len(data)))
	data = append(prefix, data...)

	return w.Sign(data)
}

// PersonalSignString use personal sign to sign a custom string message
func (w *EthereumWallet) PersonalSignString(s string) ([]byte, error) {
	return w.PersonalSign([]byte(s))
}
