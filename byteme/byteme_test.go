package main

import "encoding/hex"

func ExampleDecodeHex() {
	src := []byte("hello")
	input := hex.EncodeToString(src)

	Decode(input)
	// Output: hello
}
