package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readinput(filename string) []int {
	result := []int{}
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, val)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}

func main() {
	values := readinput("input")

	// Super lazy nested loops win!
	outerdone := false
	innerdone := false
	for i := range values {
		for j := range values {
			if i == j {
				continue
			}
			if !outerdone && values[i]+values[j] == 2020 {
				fmt.Println(values[i] * values[j])
				outerdone = true
				break
			}
			if innerdone {
				continue
			}
			for k := range values {
				if k == j || k == i {
					continue
				}
				if values[i]+values[j]+values[k] == 2020 {
					fmt.Println(values[i] * values[j] * values[k])
					innerdone = true
					break
				}
			}
		}
		if innerdone && outerdone {
			break
		}
	}
}
