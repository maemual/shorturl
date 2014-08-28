package main

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"net/http"
	"runtime"
	"strconv"
	"time"
)

var chars = [...]string{
	"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k",
	"l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v",
	"w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G",
	"H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R",
	"S", "T", "U", "V", "W", "X", "Y", "Z"}

func Shorter(data string) string {
	md5_hex := md5.Sum([]byte(data))
	md5_hash := fmt.Sprintf("%x", md5_hex)
	subLen := len(md5_hash) / 8
	var shortStr [4]string
	for i := 0; i < subLen; i++ {
		tmpStr := ""
		subStr := md5_hash[i*8 : (i+1)*8]
		x, _ := strconv.ParseInt(subStr, 16, 0)
		x = x & 0x3fffffff
		for k := 0; k < 6; k++ {
			index := 0x0000003d & x
			tmpStr += chars[index]
			x = x >> 5
		}
		shortStr[i] = tmpStr
	}
	rand.Seed(time.Now().Unix())
	return shortStr[rand.Intn(4)]
}

func EncodeHandler(w http.ResponseWriter, r *http.Request) {
	url := r.PostFormValue("url")
	if url != "" {
		w.Write([]byte(Shorter(url)))
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	http.HandleFunc("/", EncodeHandler)
	http.ListenAndServe(":8001", nil)
}
