package main

import (
	"cmp"
	passwordvalidator "github.com/wagslane/go-password-validator"
)

type passRated struct {
	password	string
	entropy		float64
}

func ratePass(password string) passRated {
	entropy := passwordvalidator.GetEntropy(password)
	return passRated{password, entropy}
}

func fpassSort(a, b passRated) int {
	entropyCmp := cmp.Compare(a.entropy, b.entropy)
	if entropyCmp != 0 {
		return entropyCmp
	}
	return cmp.Compare(a.password, b.password)
}

/*
func ratePasswords(passwords []string) []passRated {
	var ratedPasswords []passRated
	for _, pass := range passwords {
		ratedPasswords = append(ratedPasswords, ratePass(pass))
	}
	return ratedPasswords
}*/