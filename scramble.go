package main

import (
	"math/rand/v2"
	"errors"
)

func scramble(str string, aliases aliasMap) string {
	var strout string
	for _, char := range str {
		als, ok := aliases[string(char)]
		if !ok {
			strout += string(char)
			continue
		}
		alias := als[rand.IntN(len(als))]
		strout += alias
	}
	return strout
} // this probably works

func genPasswords(inPassword string, num int, aliases aliasMap, recursive bool) ([]passRated, error) {

	var outPasswords []passRated

	if len(inPassword) == 0 {
		return outPasswords, errors.New("Can't scramble string of length 0")
	} else if num <= 0 {
		return outPasswords, errors.New("Generate less than 0 passwords")
	}

	outPass := ratePass(scramble(inPassword, aliases))
	outPasswords = append(outPasswords, outPass)
	num--

	for range num {
		if recursive {
			outPass = ratePass(scramble(outPass.password, aliases))
		} else {
			outPass = ratePass(scramble(inPassword, aliases))
		}
		outPasswords = append(outPasswords, outPass)
	}

	return outPasswords, nil
}