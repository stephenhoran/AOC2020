package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

const (
	offsetX = 31
)

func checkCoordinates(x int, y int, board [][]string) bool {
	return board[y][x%offsetX] == "#"
}

func part1(input [][]string, moveX int, moveY int) (int, time.Duration) {
	t := time.Now()
	var x, y int
	var win bool
	var count int

	//offsetX = len(input[0])
	winCondition := len(input)

	for !win {
		x += moveX
		y += moveY

		if y > winCondition-1 {
			win = true
			continue
		}

		if checkCoordinates(x, y, input) {
			count++
		}
	}

	return count, time.Since(t)
}

func part2(input [][]string) (int, time.Duration) {
	t := time.Now()

	var solution int = 1
	var wg sync.WaitGroup

	c := make(chan int)
	done := make(chan bool)

	paths := []struct {
		x, y int
	}{
		{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2},
	}

	go func(solution *int) {
		for v := range c {
			*solution *= v
		}
		done <- true
	}(&solution)

	for _, coordinates := range paths {
		wg.Add(1)
		go func(c chan int, coordinates struct{ x, y int }) {
			trees, _ := part1(input, coordinates.x, coordinates.y)
			c <- trees

			wg.Done()
		}(c, coordinates)
	}

	wg.Wait()
	close(c)
	<-done

	return solution, time.Since(t)
}

func input(file string) [][]string {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal("unable to open input file")
	}
	defer f.Close()

	response := make([][]string, 0)

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		s := make([]string, 0)
		for _, i := range strings.Split(scanner.Text(), "") {
			s = append(s, i)
		}

		response = append(response, s)
	}

	return response
}

func main() {
	f := flag.String("input", "input.txt", "input file")
	rounds := flag.Int("rounds", 1, "number of rounds")

	flag.Parse()

	input := input(*f)

	var totalTime1, totalTime2 time.Duration
	var solution1, solution2 int

	for i := 0; i < *rounds; i++ {
		s, t := part1(input, 3, 1)
		totalTime1 += t
		solution1 = s
	}

	for i := 0; i < *rounds; i++ {
		s, t := part2(input)
		totalTime2 += t
		solution2 = s
	}

	fmt.Printf("solution 1: \t%v\n", solution1)
	fmt.Printf("elapsed: \t%v\n\n", totalTime1/time.Duration(*rounds))
	fmt.Printf("solution 2: \t%v\n", solution2)
	fmt.Printf("elapsed: \t%v\n", totalTime2/time.Duration(*rounds))

}
