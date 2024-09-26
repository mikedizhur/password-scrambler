package main

import (
	"bufio"
	"flag"
	"fmt"
	"slices"
	"os"
)

func main() {

	var number int
	var lines int
	var recursive bool
	var strip bool
	var private bool
	var inPassword string


	flag.IntVar(&number, "n", 1024, "number of variations")
	flag.IntVar(&lines, "l", 10, "number of lines to print")
	flag.BoolVar(&recursive, "r", false, "scramble the same password or recursively")
	flag.BoolVar(&strip, "s", false, "only print the password without entropy")
	flag.BoolVar(&private, "p", false, "private input method")

	flag.Parse()

	for _, arg := range flag.Args() {
		inPassword += arg
	}

	if private {
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		inPassword = line[:len(line)-1]
	}

	if lines > number || lines <= 0 {
		lines = number
	}
	if number <= 0 {
		panic("Won't generate 0 passwords")
	}
	if len(inPassword) <= 0 {
		panic("No password to scramble")
	}

	translations, _ := formatTranslations(defaultTranslations)
	aliases := genTransMaps(translations)
	passwords, _ := genPasswords(inPassword, number, aliases, recursive)
	slices.SortFunc(passwords, fpassSort)

	for _, password := range passwords[len(passwords) - lines:] {
		fmt.Print(password.password)
		if !strip {
			fmt.Print(" ", password.entropy)
		}
		fmt.Println()
	}

}