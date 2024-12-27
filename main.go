package main

import (
	"fmt"
)

func main() {

	// scans user input of floating point number
	var input float64
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println(err)
		return
	}

	// from floating point to integer
	var truncated = int64(input)

	// prints output
	fmt.Println(truncated)

}
