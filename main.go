// Description: lolcat but using Golang
// References:
//	https://flaviocopes.com/go-tutorial-lolcat/
//	https://github.com/dmgk/faker
//		run 'go get -u syreclabs.com/go/faker' to install
//		this just generates random text for us to use

package main

import (
	"fmt"
	"strings"

	"syreclabs.com/go/faker"
)

func main() {
	var phrases []string

	// this generates all of the text
	for i := 1; i < 3; i++ {
		phrases = append(phrases, faker.Hacker().Phrases()...)
	}

	// we print the text joined
	fmt.Println(strings.Join(phrases[:], "; "))
}
