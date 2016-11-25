// This program makes use of randmap package to get random keys/values from map

// Dhriti Shikhar

package main

import (
	"fmt"

	"github.com/lukechampine/randmap"
)

func main() {
	m := map[string]string{
		"a": "mango",
		"b": "apple",
	}

	// The non-randomness of range
	fmt.Println("Use of random:")
	for i := range m {
		fmt.Println(i, "\n")
		break
	}

	// The use of randmap
	fmt.Println("Use of randmap:")
	fmt.Println("random key--->", randmap.Key(m).(string), "\n")   // Select a random key
	fmt.Println("random value--->", randmap.Val(m).(string), "\n") // Select a random value

	// Pseudorandom is something which appears to be random but is not
	// https://boallen.com/random-numbers.html
	fmt.Println("random value--->", randmap.FastKey(m).(string), "\n") // Select a pseudorandom key
	fmt.Println("random value--->", randmap.FastVal(m).(string), "\n") // Select a pseudorandom value
}
