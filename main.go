// Description: lolcat but using Golang
// References:
//	https://flaviocopes.com/go-tutorial-lolcat/
//	https://github.com/dmgk/faker
//		run 'go get -u syreclabs.com/go/faker' to install
//		this just generates random text for us to use
// 	https://en.wikipedia.org/wiki/ANSI_escape_code
//		to understand how we color the text

package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
)

// this function generates the rgb rainbow
func rgb(i int) (int, int, int) {
	var f = 0.1
	ifloat := float64(i)
	return int(math.Sin(f*ifloat+0)*127 + 128),
		int(math.Sin(f*ifloat+2*math.Pi/3)*127 + 128),
		int(math.Sin(f*ifloat+4*math.Pi/3)*127 + 128)
}

func main() {
	// get stdin info
	info, err := os.Stdin.Stat()
	if err != nil {
		fmt.Println("Error!")
		return
	}
	// rune is an alias for type int32 and represents a code point
	var output []rune

	// checks that we piped something
	if info.Mode()&os.ModeCharDevice != 0 {
		fmt.Println("The command is intended to work with pipes.")
		fmt.Println("Usage: fortune | lolgopher")
	}

	reader := bufio.NewReader(os.Stdin)
	j := 0
	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		output = append(output, input)
		// generate color
		r, g, b := rgb(j)

		// we print each character in rgb rainbow
		// \033 is the ASCII escape character
		// begin ANSI Escape sequence
		// [ - control sequence introducer
		// 38 - set foreground color
		// 		is followed by 2;r;g;b
		// m - sets the appearance of the following characters
		// 0 - reset / set all attributes back to normal
		fmt.Printf("\033[38;2;%d;%d;%dm%c\033[0m", r, g, b, output[j])
		j++
	}
}
