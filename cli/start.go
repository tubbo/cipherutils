package cli

import (
	"github.com/yuya-takeyama/argf"
	"io/ioutil"
	"log"
)

// Start a new CLI
func Start(callback func(input string)) {
	reader, err := argf.Argf()

	if err != nil {
		log.Fatal(err)
	}

	buf, err := ioutil.ReadAll(reader)

	if err != nil {
		log.Fatal(err)
	}

	input := string(buf)

	callback(input)
}
