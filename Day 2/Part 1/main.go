package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func minMax(s string) (int, int) {
	numbers := strings.Split(s, "-")
	min, _ := strconv.Atoi(numbers[0])
	max, _ := strconv.Atoi(numbers[1])

	return min, max
}

func checkPolicy(passwd string) bool {
	policy := strings.Split(passwd, " ")
	min, max := minMax(policy[0])
	letter := policy[1][0]

	//for i := 0; i <= len(password)-1; i++ {
	//	if password[i] == letter {
	//		count++
	//	}
	//}

	count := strings.Count(policy[2], string(letter))

	if count >= min && count <= max {
		return true
	}

	return false
}

func input() []string {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("unable to open input file")
	}
	defer f.Close()

	response := make([]string, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		response = append(response, scanner.Text())
	}

	return response
}

func exec() {
	input := input()

	t := time.Now()

	var count int

	for _, passwd := range input {
		if checkPolicy(passwd) {
			count++
		}
	}

	fmt.Println(time.Since(t))
}

func main() {
	exec()
}
