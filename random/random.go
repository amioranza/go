/*
Purpose:
Number guessing by random generators using simple and advanced methods

Author:
amioranza@mdcnet.ninja

https://github.com/amioranza/random

Usage:

	random <MAX>

	<MAX> is the highest number to guess

	random <MIN> <MAX>

	<MIN> is the lowest number to guess
	<MAX> is the highest number to guess

*/
package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// genRandom generatas the random number based on the nanosec of time
func genRandom(max uint) uint {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(int(max) + 1)
	if r == 0 {
		log.Println("Rand generated number zero, changing to number one")
		r++
	} else {
		log.Printf("Rand generated number %v is ok\n", r)
	}
	return uint(r)
}

// genRandom2 fills a slice with the range of the numbers specified as parameters and calls genRandom
// to guess the index of the slice to get the number guessed
func genRandom2(beg, end int) int {
	var numRange []int

	if beg > end {
		log.Fatal("ERROR: Begin greater than end.")
	}

	size := (end - beg) + 1
	log.Printf("Random number generator: Begin=%v and End=%v, Slice size=%v\n", beg, end, size)

	for i := 1; i <= size; i++ {
		num := (beg + i) - 1
		numRange = append(numRange, num)
		//fmt.Printf("idx: %v, num: %v\n", i, num)
	}

	log.Println("Slice: ", numRange)

	idx := genRandom(uint(size))

	return numRange[idx-1]
}

func parseArg(param string) int {
	arg, err := strconv.Atoi(param)
	if err != nil {
		log.Fatal(err)
	}
	return arg
}

func main() {
	switch {
	case len(os.Args) == 2:
		max := uint(parseArg(os.Args[1]))
		random := genRandom(uint(max))
		fmt.Printf("The guessed number is >>>>>>>>>> %v <<<<<<<<<<<<\n", random)
	case len(os.Args) == 3:
		min := parseArg(os.Args[1])
		max := parseArg(os.Args[2])
		random := genRandom2(int(min), int(max))
		fmt.Println("")
		fmt.Printf("The guessed number is >>>>>>>>>> %v <<<<<<<<<<<<\n", random)
	default:
		fmt.Println("Usage:")
		fmt.Println("\trandom <MAX> or random <MIN> <MAX>")
		fmt.Println("\tPlease inform at least one value to guess the number.")
	}
}
