package parser

type Set map[string]bool

type Result struct {
	InitialFacts map[int]Set
	Instructions map[int]string
	Queries      map[string]bool
}

func NewSet() *Set {
	return &Set{
		"A": false,
		"B": false,
		"C": false,
		"D": false,
		"E": false,
		"F": false,
		"G": false,
		"H": false,
		"I": false,
		"J": false,
		"K": false,
		"L": false,
		"M": false,
		"N": false,
		"O": false,
		"P": false,
		"Q": false,
		"R": false,
		"S": false,
		"T": false,
		"U": false,
		"V": false,
		"W": false,
		"X": false,
		"Y": false,
		"Z": false,
	}
}
