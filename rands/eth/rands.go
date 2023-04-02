package rands

import (
	"context"
	"math/big"
	"math/rand"

	ec "github.com/ethereum/go-ethereum/ethclient"
)

func Rands(n []int, num int) ([]int, error) {
	client, err := ec.Dial("https://mainnet.infura.io")
	if err != nil {
		return nil, err
	}
	defer client.Close()
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	seed := big.NewInt(0).SetBytes(header.Hash().Bytes())
	rand.Seed(seed.Int64())
	idxs := rand.Perm(len(n))[:num]
	result := make([]int, 0, num)
	for _, idx := range idxs {
		result = append(result, n[idx])
	}
	return result, nil
}
