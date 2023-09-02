package ethereum

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ChainSafe/chainbridge-utils/crypto/secp256k1"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

const (
	decimalBase = 10
	hexBase     = 16
)

// valueToBig converts a string value to a *big.Int in the provided base
func valueToBig(value string, base int) (*big.Int, error) {
	val, ok := new(big.Int).SetString(value, base)
	if !ok {
		return nil, fmt.Errorf("unable to parse value")
	}

	return val, nil
}

// ParseUint256OrHex parses a string value as either a base 10 number
// or as a hex value
func ParseUint256OrHex(value *string) (*big.Int, error) {
	// Check if the value is valid
	if value == nil {
		return nil, fmt.Errorf("invalid value")
	}

	// Check if the value is hex
	if strings.HasPrefix(*value, "0x") {
		// Hex value, remove the prefix and parse it
		clipped := (*value)[2:]
		return valueToBig(clipped, hexBase)
	}

	// Decimal number, parse it
	return valueToBig(*value, decimalBase)
}

func padWithZeros(key []byte, targetLength int) []byte {
	res := make([]byte, targetLength-len(key))
	return append(res, key...)
}

func makeEthRing() map[string]*secp256k1.Keypair {
	ring := map[string]*secp256k1.Keypair{}
	for _, key := range Keys {
		bz := padWithZeros([]byte(key), secp256k1.PrivateKeyLength)
		kp, err := secp256k1.NewKeypairFromPrivateKey(bz)
		if err != nil {
			panic(err)
		}
		ring[key] = kp
	}

	return ring
}

func makeKey(privateKey string) (*secp256k1.Keypair, error) {
	bz, err := hexutil.Decode(privateKey)
	kp, err := secp256k1.NewKeypairFromPrivateKey(bz)
	if err != nil {
		panic(err)
	}
	return kp, nil
}
