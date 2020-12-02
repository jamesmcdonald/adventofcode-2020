package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type passent struct {
	minimum  int
	maximum  int
	match    rune
	password string
}

func readinput(filename string) []passent {
	result := []passent{}
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var p passent
		fmt.Sscanf(scanner.Text(), "%d-%d %c: %s", &p.minimum, &p.maximum, &p.match, &p.password)
		result = append(result, p)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}

func main() {
	values := readinput("input")
	valid := 0
	newvalid := 0

	for _, pass := range values {
		// Valid if there are at least minimum and at most maximum matches
		matchcount := 0
		for _, c := range pass.password {
			if c == pass.match {
				matchcount++
			}
		}
		if matchcount <= pass.maximum && matchcount >= pass.minimum {
			valid++
		}

		// New valid if exactly one of minimum, maximum is match
		var m byte = byte(pass.match)
		if (pass.password[pass.minimum-1] == m && pass.password[pass.maximum-1] != m) ||
			(pass.password[pass.maximum-1] == m && pass.password[pass.minimum-1] != m) {
			newvalid++
		}

	}

	fmt.Println(valid)
	fmt.Println(newvalid)
}
