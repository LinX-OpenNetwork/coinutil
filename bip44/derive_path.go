package bip44

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/LinX-OpenNetwork/coinutil/bip32"
)

// DerivePath is the key path of Hierarchical Deterministic Wallets
type DerivePath string

// DerivePathParams is the BIP44 params of a derived path
type DerivePathParams struct {
	// Purpose is the purpose field of BIP44. Must be 44'
	Purpose uint32
	// CoinType is the coin type for BIP44.
	// 0' for Bitcoin, 60' for Ethereum, 501' for Solana
	// See https://github.com/satoshilabs/slips/blob/master/slip-0044.md for detail.
	CoinType                      uint32
	Account, Change, AddressIndex uint32
	// Extra for any possible path
	Extra []uint32
	// Depth for the actual depth of the path
	Depth uint
}

func parseIndex(s string) (uint32, error) {
	index := uint32(0)
	if strings.HasSuffix(s, "'") {
		index |= bip32.FirstHardenedChild
		s = s[:len(s)-1]
	}
	idx, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0, err
	}
	return index | uint32(idx), nil

}

// ToParams converts a string derive path to DerivePathParams
func (p DerivePath) ToParams() (*DerivePathParams, error) {
	components := strings.Split(string(p), "/")
	n := len(components)
	if n < 2 || components[0] != "m" {
		return nil, fmt.Errorf("invalid derive path: %s. must have at least 2 components and starts with 'm/'", p)
	}
	params := DerivePathParams{
		Depth: uint(n - 1),
	}
	var err error
	params.Purpose, err = parseIndex(components[1])
	if err != nil {
		return nil, err
	}
	if n >= 3 {
		params.CoinType, err = parseIndex(components[2])
		if err != nil {
			return nil, err
		}
	}
	if n >= 4 {
		params.Account, err = parseIndex(components[3])
		if err != nil {
			return nil, err
		}
	}
	if n >= 5 {
		params.Change, err = parseIndex(components[4])
		if err != nil {
			return nil, err
		}
	}
	if n >= 6 {
		params.AddressIndex, err = parseIndex(components[5])
		if err != nil {
			return nil, err
		}
	}
	if n > 6 {
		extras := components[6:]
		for _, extra := range extras {
			idx, err := parseIndex(extra)
			if err != nil {
				return nil, err
			}
			params.Extra = append(params.Extra, idx)
		}
	}
	return &params, nil
}

// Indexes returns child indexes of this derived path.
func (p *DerivePathParams) Indexes() []uint32 {
	indexes := make([]uint32, 0)
	indexes = append(indexes, p.Purpose)
	if p.Depth >= 2 {
		indexes = append(indexes, p.CoinType)
	}
	if p.Depth >= 3 {
		indexes = append(indexes, p.Account)
	}
	if p.Depth >= 4 {
		indexes = append(indexes, p.Change)
	}
	if p.Depth >= 5 {
		indexes = append(indexes, p.AddressIndex)
	}
	if p.Depth > 5 {
		indexes = append(indexes, p.Extra...)
	}
	return indexes
}
