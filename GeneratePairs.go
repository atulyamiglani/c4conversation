package main

import "math/rand"

// takes in a array of member strings
// returns in a [][2 | 3] array with pairs
func GeneratePairs(ids []string) ([][]string) {
	rand.Shuffle(len(ids), func(i, j int) { ids[i], ids[j] = ids[j], ids[i] })

	var pairs [][]string 
	if len(ids) == 1 {
		return pairs; 
	}

	var lonelyPerson string
	if (len(ids) % 2 == 1) {
		lonelyPerson = ids[0]
		ids = ids[1:]

	}
	for i := 0; i < len(ids); i+= 2 {
		pair := []string{ids[i], ids[i+1]}
		pairs = append(pairs, pair)
	}

	// add the lonely person to an existing group
	if lonelyPerson != "" {
		pairs[0] = append(pairs[0], lonelyPerson)
	}
	
	return pairs; 
}

