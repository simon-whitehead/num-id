package numid

import (
	"bytes"
	"math"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Encode(n int64) string {
	result := bytes.NewBuffer([]byte{})
	f := float64(n)
	al := float64(len(alphabet))
	for i := math.Floor(math.Log(f) / math.Log(al)); i >= 0; i-- {
		idx := int(math.Mod(math.Floor(f/bpow(al, i)), al))
		result.WriteString(alphabet[idx : idx+1])
	}

	return reverseString(result.String())
}

func Decode(id string) int64 {
	str := reverseString(id)
	result := int64(0)
	end := len(str) - 1
	al := float64(len(alphabet))
	for i := 0; i <= end; i++ {
		result += int64(float64(strings.Index(alphabet, str[i:i+1])) * bpow(al, float64(end-i)))
	}

	return result
}

func reverseString(s string) string {
	n := len(s)
	runes := make([]rune, n)
	for _, rune := range s {
		n--
		runes[n] = rune
	}
	return string(runes[n:])
}

func bpow(a float64, b float64) float64 {
	return math.Floor(math.Pow(a, b))
}
