package main

import "math/rand"

var bigBazoka = []string{
	"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n",
	"p", "q", "r", "s", "t", "v", "w", "x", "y", "z",
}

var smallBazoka = []string{
	"a", "e", "i", "o", "u",
}

func randProject() string {
	bigSize := 2 + rand.Intn(2)   // 2 or 3
	smallSize := 2 + rand.Intn(2) // 2 or 3

	project := ""
	for i := 0; i < bigSize; i += 1 {
		bigIndex := rand.Intn(len(bigBazoka))
		project += bigBazoka[bigIndex]

		for j := 0; j < smallSize; j += 1 {
			smallIndex := rand.Intn(len(smallBazoka))
			project += smallBazoka[smallIndex]
		}
	}

	return project
}
