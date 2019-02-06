package parser

type ParseResult struct {
	InitialFacts []Set
	Query        []*string
	Instructions []map[string]interface{}
}

// implementation parse result
func NewParseResult() *ParseResult {
	return &ParseResult{
		InitialFacts: []Set{},
		Query:        []*string{},
		Instructions: []map[string]interface{}{},
	}
}

func (p *ParseResult) addSet(newSet Set) {
	p.InitialFacts = append(p.InitialFacts, newSet)
}

func (p *ParseResult) addRequest(req *string) {
	p.Query = append(p.Query, req)
}
