package main

import (
	"cmp"
	// passwordvalidator "github.com/wagslane/go-password-validator"
	zxcvbn "github.com/wneessen/zxcvbn-go"
	scoring "github.com/wneessen/zxcvbn-go/scoring"
)

type passRated struct {
	password string
	// entropy   float64
	zxcvbnRes scoring.MinEntropyMatch
}

func ratePass(password string) passRated {
	// entropy := passwordvalidator.GetEntropy(password)
	zxcvbnRes := zxcvbn.PasswordStrength(password, nil)
	// return passRated{password, entropy, zxcvbnRes}
	return passRated{password, zxcvbnRes}
}

func fpassSort(a, b passRated) int {
	zScoreCmp := cmp.Compare(a.zxcvbnRes.Score, b.zxcvbnRes.Score)
	if zScoreCmp != 0 {
		return zScoreCmp
	}
	zEntropyCmp := cmp.Compare(a.zxcvbnRes.Entropy, b.zxcvbnRes.Entropy)
	if zEntropyCmp != 0 {
		return zEntropyCmp
	}
	/*
		entropyCmp := cmp.Compare(a.entropy, b.entropy)
		if entropyCmp != 0 {
			return entropyCmp
		}
	*/
	zTimeCmp := cmp.Compare(a.zxcvbnRes.CrackTime, b.zxcvbnRes.CrackTime)
	if zTimeCmp != 0 {
		return zTimeCmp
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
