package keyoperations

import (
	"crypto/sha256"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/ripemd160"
)

func PrivateKeyToPublicKey(privateKey []byte) []byte {
	priv, _ := btcec.PrivKeyFromBytes(privateKey)
	pubKey := priv.PubKey()
	return pubKey.SerializeCompressed()
}

func publicKeyHash(pubKey []byte) []byte {
	sha := sha256.Sum256(pubKey)

	ripmd := ripemd160.New()
	ripmd.Write(sha[:])
	return ripmd.Sum(nil)
}

func addVersion(pubKeyHash []byte) []byte {
	return append([]byte{0x00}, pubKeyHash...)
}

func checkSum(data []byte) []byte {
	first := sha256.Sum256(data)
	second := sha256.Sum256(first[:])
	return second[:4]
}

func base58CheckEncode(data []byte) string {
	ck := checkSum(data)
	
	full := append(data, ck...)
	return base58.Encode(full)
}

func GenerateP2PKeyAddress(privateKey []byte) string {
	pubKey := PrivateKeyToPublicKey(privateKey)
	pubKeyHashed := publicKeyHash(pubKey)
	versioned := addVersion(pubKeyHashed)
	return base58CheckEncode(versioned)
}
