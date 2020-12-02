package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func input() []int {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("unable to open input file")
	}
	defer f.Close()

	response := make([]int, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		response = append(response, i)
	}

	return response
}

func main() {
	f, err := os.Create("memprofile")
	if err != nil {
		log.Fatalln("Cannot open file for memory profiling")
	}
	defer f.Close()

	input := input()

	t := time.Now()

	for _, num := range input {
		for _, i := range input {
			for _, n := range input {
				if num+i+n == 2020 {
					fmt.Println(time.Since(t))
					fmt.Println(num * i * n)

					return
				}
			}
		}
	}
}
