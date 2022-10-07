package urlGenerator

import (
	"math/big"
	"strconv"
	"crypto/rand"
)

var characters = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
	"0123456789"

func GenerateShortUrl() string {
	var res string
	for i := 0; i < 8; i++ {
		res += getRandomCharacter()
	}
	return res
}

func getRandomCharacter() string {
	t := big.NewInt(int64(len(characters)))
	RandomCrypto, _ := rand.Int(rand.Reader, t)
	randomized, _ := strconv.Atoi(RandomCrypto.String())
	return characters[randomized:randomized + 1]
}
