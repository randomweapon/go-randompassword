package main

import (
	"math/rand"
)

func randomByWeight(keys []string, weights []int) string {

	length := 11
	keyBytes := keys
	keyWeights := weights
	random := rand.Intn(length)
	cursor := 0

	for i := range keyWeights {
		cursor += keyWeights[i]
		if cursor > random {
			return keyBytes[i]
		}
	}

	return ""
}
