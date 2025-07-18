package keyoperations

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/binary"
	"math/big"
	"github.com/btcsuite/btcd/btcec/v2"
	"fmt"
)

var curve = btcec.S256()
var curveN = curve.Params().N

func uint32ToBytes(i uint32) []byte {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, i)
	return b
}

func DeriveHardenedChilds(parentPrivKey, parentChainCode []byte, index uint32) ([]byte, []byte, error) {
	if index < 0x80000000 {
		return nil, nil, fmt.Errorf("index must be >= 0x80000000 for hardened derivation")
	}

	data := append([]byte{0x00}, parentPrivKey...)
	data = append(data, uint32ToBytes(index)...)

	mac := hmac.New(sha512.New, parentChainCode)
	mac.Write(data)
	I := mac.Sum(nil)

	IL := I[:32]
	IR := I[32:]

	// Convert bytes to big integers
	ilInt := new(big.Int).SetBytes(IL)
	kParInt := new(big.Int).SetBytes(parentPrivKey)

	childKeyInt := new(big.Int).Add(ilInt, kParInt)
	childKeyInt.Mod(childKeyInt, curveN)

	if childKeyInt.Sign() == 0 {
		return nil, nil, fmt.Errorf("derived key is zero")
	}

	childKey := childKeyInt.Bytes()

	// Ensure the child key is 32 bytes
	if len(childKey) < 32 {
		padded := make([]byte, 32)
		copy(padded[32-len(childKey):], childKey)
		childKey = padded
	}

	return childKey, IR, nil
}
