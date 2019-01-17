package parser

type Set map[string]bool

type Result struct {
	InitialFacts map[int]Set
	Instructions map[int]string
	Queries      map[string]bool
}
