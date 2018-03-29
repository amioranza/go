/*
Purpose:
Number guessing by random generators using simple and advanced methods

Author:
amioranza@mdcnet.ninja

https://github.com/amioranza/random

Usage:

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

func genRandom(max uint) uint {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(int(max) + 1)
	switch {
	case r == 0:
		log.Println("Rand generated number zero, changing to number one")
		r++
	default:
		log.Printf("Rand generated number %v is ok\n", r)
	}
	return uint(r)
}

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

func main() {
	// Validates the number of arguments
	if len(os.Args) <= 1 || len(os.Args) <= 2 || len(os.Args) > 3 {
		fmt.Println("Usage:")
		fmt.Println("\trandom <MIN> <MAX>")
		log.Fatal("Please inform the <MIN> and <MAX> values")
	}
	// Parse arguments as unsigned integers
	min, err1 := strconv.ParseUint(os.Args[1], 0, 64)
	if err1 != nil {
		panic(err1)
	}
	max, err2 := strconv.ParseUint(os.Args[2], 0, 64)
	if err2 != nil {
		panic(err2)
	}

	// Simple random generator, only work with max
	//random := genRandom(uint(max))
	//fmt.Printf("1: The generated number is: %v\n", random)

	// Advanced random generator using a slice and calculating the size of the slice, filling slice
	// and using the simple rand generator to guess the index of the sliec to return the real number guessed
	random2 := genRandom2(int(min), int(max))
	fmt.Println("")
	fmt.Printf("The guessed number is >>>>>>>>>> %v <<<<<<<<<<<<\n", random2)
}
