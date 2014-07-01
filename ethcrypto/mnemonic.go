package ethcrypto

import (
	"fmt"
	"io/ioutil"
	"path"
	"runtime"
	"strconv"
	"strings"
)

func InitWords() []string {
	_, thisfile, _, _ := runtime.Caller(1)
	filename := path.Join(path.Dir(thisfile), "mnemonic.words.lst")
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(fmt.Errorf("reading mnemonic word list file 'mnemonic.words.lst' failed: ", err))
	}
	return strings.Split(string(content), "\n")
}

var words = InitWords()

// TODO: See if we can refactor this into a shared util lib if we need it multiple times
func IndexOf(slice []string, value string) int64 {
	for p, v := range slice {
		if v == value {
			return int64(p)
		}
	}
	return -1
}

func MnemonicEncode(message string) []string {
	var out []string
	n := int64(len(words))

	for i := 0; i < len(message); i += (len(message) / 8) {
		x := message[i : i+8]
		bit, _ := strconv.ParseInt(x, 16, 64)
		w1 := (bit % n)
		w2 := ((bit / n) + w1) % n
		w3 := ((bit / n / n) + w2) % n
		out = append(out, words[w1], words[w2], words[w3])
	}
	return out
}

func MnemonicDecode(wordsar []string) string {
	var out string
	n := int64(len(words))

	for i := 0; i < len(wordsar); i += 3 {
		word1 := wordsar[i]
		word2 := wordsar[i+1]
		word3 := wordsar[i+2]
		w1 := IndexOf(words, word1)
		w2 := IndexOf(words, word2)
		w3 := IndexOf(words, word3)

		y := (w2 - w1) % n
		z := (w3 - w2) % n

		// Golang handles modulo with negative numbers different then most languages
		// The modulo can be negative, we don't want that.
		if z < 0 {
			z += n
		}
		if y < 0 {
			y += n
		}
		x := w1 + n*(y) + n*n*(z)
		out += fmt.Sprintf("%08x", x)
	}
	return out
}