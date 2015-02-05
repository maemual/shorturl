/*
	A very esay short url package.
*/
package shorturl

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
)

var chars = [...]string{
	"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k",
	"l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v",
	"w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G",
	"H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R",
	"S", "T", "U", "V", "W", "X", "Y", "Z"}

// Normally, one url can have 4 short urls.
// So, idx meanings which one you need(0-3).
func New(url string, idx int) (shortStr string) {
	md5_hex := md5.Sum([]byte(url))
	md5_hash := hex.EncodeToString(md5_hex[:])
	shortStr = ""
	subStr := md5_hash[idx*8 : (idx+1)*8]
	x, _ := strconv.ParseInt(subStr, 16, 0)
	x = x & 0x3fffffff
	for k := 0; k < 6; k++ {
		index := 0x0000003d & x
		shortStr += chars[index]
		x = x >> 5
	}
	return
}
