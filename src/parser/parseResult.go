package parser

type ParseResult struct {
	InitialFacts []Set
	Query        []*string
}

// implementation parse result

func NewParseResult() *ParseResult {
	return &ParseResult{
		[]Set{},
		[]*string{},
	}
}

func (p *ParseResult) addSet(newSet Set) {
	p.InitialFacts = append(p.InitialFacts, newSet)
}

func (p *ParseResult) addRequest(req *string) {
	p.Query = append(p.Query, req)
}
