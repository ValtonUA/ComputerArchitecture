package main

import (
	"fmt"
	"lab1/morze"
	"time"
)

func main() {
	task1()
	task2()
}

func task1() {
	var text = "Hello, I`m 20 years old."
	var sequence string

	fmt.Println("Initial string:\n", text)

	sequence = morze.Encode(text)

	fmt.Println("Convert to Morze alphabet:\n", sequence)

	text = morze.Decode(sequence)

	fmt.Println("Back to English:\n", text)
}

func task2() {
	var text, sequence string

	fmt.Println("Generating MorzeSequence...")
	var startTime = time.Now()
	var startPartTime = startTime

	sequence = morze.GetRandomMorzeSequence(10000)

	var engGenTime = time.Since(startPartTime)

	fmt.Println("Generated sequence:\n", sequence)

	fmt.Println("Decoding MorzeSequence to plain text...")

	startPartTime = time.Now()

	text = morze.Decode(sequence)

	fmt.Println("Decoded plain text:\n", text)

	var endTime = time.Since(startTime)
	var endDecodeTime = time.Since(startPartTime)

	fmt.Printf("Generating time: %v\n", engGenTime)
	fmt.Printf("Decode time: %v\n", endDecodeTime)
	fmt.Printf("Whole process time: %v\n", endTime)
}
