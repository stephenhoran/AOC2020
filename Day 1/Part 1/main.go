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
	input := input()

	t := time.Now()

	for _, num := range input {
		for _, i := range input {
			if num+i == 2020 {
				fmt.Println(num * i)
				fmt.Println(time.Since(t))
				return
			}
		}
	}

}
