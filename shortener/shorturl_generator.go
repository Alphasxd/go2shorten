package shortener

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"os"

	"github.com/itchyny/base58-go"
)

// sha256Of 计算字符串的 sha256 值
func sha256Of(input string) []byte {
	algo := sha256.New()
	algo.Write([]byte(input))
	return algo.Sum(nil)
}

// base58Encode 将字节数组转换为 base58 编码的字符串
func base58Encode(bytes []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(encoded)
}

// GenerateShortURL 生成短链接
func GenerateShortURL(initialURL string, userID string) string {
	// 将 initialURL 和 userID 拼接起来，计算 sha256 值
	urlHash := sha256Of(initialURL + userID)
	// 将 sha256 值转换为 uint64 类型的数字
	generatedNum := new(big.Int).SetBytes(urlHash).Uint64()
	// 将数字转换为 base58 编码的字符串
	finalURL := base58Encode([]byte(fmt.Sprintf("%d", generatedNum)))
	// 返回前 8 位
	return finalURL[:8]
}
