package shortener

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"os"

	"github.com/itchyny/base58-go"
)

// GenerateShortLink 生成短链接
func GenerateShortLink(initialLink string, userId string) string {
	// 将初始链接和用户ID连接起来，然后计算SHA256哈希值
	urlHashBytes := sha256Of(initialLink + userId)
	// 将哈希值转换为大整数，然后将其转换为base58编码的字符串
	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
	finalString := base58Encoded([]byte(fmt.Sprintf("%d", generatedNumber)))
	// 返回前8个字符
	return finalString[:8]
}

// sha256Of 将输入字符串转换为sha256哈希值
func sha256Of(input string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(input))
	return algorithm.Sum(nil)
}

// base58Encoded 将字节数组转换为base58编码的字符串
func base58Encoded(bytes []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(encoded)
}
