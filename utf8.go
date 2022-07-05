package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func encode(utf32 []rune) []byte {
	var (
		result []byte
		i      int
		pl     int
	)
	result = make([]byte, len(utf32)*4)
	i = 0
	pl = 128
	for x := 0; x < len(utf32); x++ {
		if utf32[x] <= rune(math.Pow(2, 7)-1) {
			result[i] = byte(utf32[x])
			i++
		} else if utf32[x] <= rune(math.Pow(2, 10)-1) {
			result[i] = byte(utf32[x]>>6) + 192
			i++
			result[i] = byte((uint64(utf32[x])<<26)>>26) + byte(pl)
			i++
		} else if utf32[x] <= rune(math.Pow(2, 16)-1) {
			result[i] = byte(utf32[x]>>12) + 224
			i++
			result[i] = byte((uint64(utf32[x])<<20)>>26) + byte(pl)
			i++
			result[i] = byte((uint64(utf32[x])<<26)>>26) + byte(pl)
			i++
		} else if utf32[x] <= rune(math.Pow(2, 21)-1) {
			result[i] = byte(utf32[x]>>18) + 240
			i++
			result[i] = byte((uint64(utf32[x])<<14)>>26) + byte(pl)
			i++
			result[i] = byte((uint64(utf32[x])<<20)>>26) + byte(pl)
			i++
			result[i] = byte((uint64(utf32[x])<<26)>>26) + byte(pl)
			i++
		}
	}

	return result[0:i]
}

func decode(utf8 []byte) []rune {
	var (
		result []rune
		i      int
	)
	result = make([]rune, len(utf8))
	i = 0
	for x := 0; x < len(utf8); i++ {
		if utf8[x]>>7 == 0 {
			result[i] = rune(utf8[x])
			x += 1
		} else if utf8[x]>>5 == 6 && utf8[x+1]>>6 == 2 {
			result[i] = (rune((utf8[x]<<3)>>3) << 6) + (rune((utf8[x+1] << 2) >> 2))
			x += 2
		} else if utf8[x]>>4 == 14 && utf8[x+1]>>6 == 2 && utf8[x+2]>>6 == 2 {
			result[i] = (rune((utf8[x]<<4)>>4) << 12) + (rune((utf8[x+1]<<2)>>2) << 6) + (rune((utf8[x+2] << 2) >> 2))
			x += 3
		} else if utf8[x]>>4 == 15 && utf8[x+1]>>6 == 2 && utf8[x+2]>>6 == 2 && utf8[x+3]>>6 == 2 {
			result[i] = (rune((utf8[x]<<5)>>5) << 18) + (rune((utf8[x+1]<<2)>>2) << 12) + (rune((utf8[x+2]<<2)>>2) << 6) + (rune((utf8[x+3] << 2) >> 2))
			x += 4
		}
	}

	return result[0:i]
}

func main() {
	var F string
	myscanner := bufio.NewScanner(os.Stdin)
	myscanner.Scan()
	F = myscanner.Text()

	encodedText := encode([]rune(F))
	for i := 0; i < len(encodedText); i++ {
		fmt.Print(encodedText[i], " ")
	}

	fmt.Println()
	decodedText := decode(encodedText)
	for i := 0; i < len(decodedText); i++ {
		fmt.Print(string(decodedText[i]))
	}
}
