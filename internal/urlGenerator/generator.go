package urlGenerator

import (
	"math/big"
	"strconv"
	"crypto/rand"
)

var (
	lowers = "abcdefghijklmnopqrstuvwxyz"
	uppers = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits = "0123456789"
	)

func GenerateShortUrl() string {
	var res string
	for i := 0; i < 3; i++ {
		res += getRandomCharacter(lowers)
		res += getRandomCharacter(digits)
		if i > 0 {
			res += getRandomCharacter(uppers)
		}
	}
	return res
}

func getRandomCharacter(str string) string {
	t := big.NewInt(int64(len(str)))
	RandomCrypto, _ := rand.Int(rand.Reader, t)
	randomized, _ := strconv.Atoi(RandomCrypto.String())
	return str[randomized:randomized + 1]
}
