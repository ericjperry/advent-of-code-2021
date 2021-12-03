package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

func sum(window []int) int {
	total := 0
	for i := range window {
		total += window[i]
	}
	return total
}

func calculate_depth_increases(window_size int) {
	last_measurement := make([]int, window_size)
	current_measurement := make([]int, window_size)
	for i := range last_measurement {
		last_measurement[i] = 0
	}
	for i := range current_measurement {
		current_measurement[i] = 0
	}
	num_of_increases := 0
	file, err := os.Open("aoc1.txt")

	if err != nil {
		log.Fatalf("Could not open the data file.")
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		for i := range last_measurement {
			last_measurement[i] = current_measurement[i]
		}
		for i := 1; i < len(current_measurement); i++ {
			current_measurement[i-1] = current_measurement[i]
		}
		has_zero := false
		for i := range last_measurement {
			if last_measurement[i] == 0 {
				has_zero = true
				break
			}
		}
		current_measurement[len(current_measurement)-1], _ = strconv.Atoi(scanner.Text())
		if !has_zero && sum(current_measurement) > sum(last_measurement) {
			num_of_increases++
		}
	}

	// The method os.File.Close() is called
	// on the os.File object to close the file
	file.Close()

	result := fmt.Sprintf("There were %d increases in measurement.", num_of_increases)
	fmt.Println(result)
}

func main() {
	window_size_ptr := flag.Int("window", 1, "The window size to calculate depth increases.")
	flag.Parse()
	calculate_depth_increases(*window_size_ptr)
}
