package parser

import (
	"fmt"
	"log"
	"regexp"
)

// type Parser struct

func (p *Parser) SaveState() int {
	return p.Index
}

func (p *Parser) RestoreState(state int, b bool) {
	if !b {
		p.Index = state
	}
}

func (p *Parser) ReachedEnd() bool {
	if p.Index == p.Tokens.Len() {
		return false
	}
	return true
}

func (p *Parser) Optional(_ interface{}) bool {
	return true
}

func (p *Parser) TokenIsType(regexp *regexp.Regexp) (b bool) {
	if b = regexp.MatchString(*(*p.Tokens)[p.Index].Content); b {
		p.Index++
	}
	return
}

func (p *Parser) RuleEmptyLine() (b bool) {
	state := p.SaveState()
	b = p.Optional(p.TokenIsType((*p.TokenInfos)[Comment].Regex)) &&
		p.TokenIsType((*p.TokenInfos)[EndLine].Regex)
	p.RestoreState(state, b)
	return
}

func (p *Parser) RuleAxiom() (b bool) {
	state := p.SaveState()
	if b = p.TokenIsType((*p.TokenInfos)[Axiom].Regex); b {
		p.Stack.PushRight((*p.Tokens)[p.Index-1].Content)
	}
	p.RestoreState(state, b)
	return
}

func (p *Parser) RuleNot() (b bool) {
	p.SaveState()
	b = p.TokenIsType((*p.TokenInfos)[Not].Regex) && p.RuleAxiom()
	return
}

func (p *Parser) RuleValue() (b bool) {
	p.SaveState()
	b = p.RuleNot()
	return
}

func (p *Parser) RuleExpression() (b bool) {
	p.SaveState()
	b = p.RuleValue()
	return
}

func (p *Parser) RuleInstruction() (b bool) {
	state := p.SaveState()
	if b = p.RuleExpression() &&
		p.TokenIsType((*p.TokenInfos)[Implies].Regex) &&
		p.RuleExpression() &&
		p.Optional(p.TokenIsType((*p.TokenInfos)[Comment].Regex)) &&
		p.TokenIsType((*p.TokenInfos)[EndLine].Regex); b {
		if err := p.AddInstructions(p.PopStackTwo()); err != nil {
			log.Fatal(err)
		}
		p.Stack.Reset()
	}
	p.RestoreState(state, b)
	return
}

func (p *Parser) RuleInitialFacts() (b bool) {
	state := p.SaveState()
	b = p.TokenIsType((*p.TokenInfos)[InitialFacts].Regex)
	newSet := NewSet()
	token := (*p.Tokens)[p.Index-1].Type
	for token == Axiom {
		stringToken := *(*p.Tokens)[p.Index-1].Content
		if _, ok := (*newSet)[stringToken]; ok {
			(*newSet)[stringToken] = true
		} else {
			log.Fatal("Issue Rule Initial Facts")
		}
		p.Index ++
		token = (*p.Tokens)[p.Index-1].Type
	}
	p.TokenIsType((*p.TokenInfos)[Comment].Regex)
	if b = b && p.TokenIsType((*p.TokenInfos)[EndLine].Regex); b {
		p.InitialFacts = append(p.InitialFacts, *newSet)
	}
	p.RestoreState(state, b)
	return
}

func (p *Parser) Process() {
	carryOn := true
	for carryOn && !p.ReachedEnd() {
		carryOn = p.RuleEmptyLine() ||
			p.RuleInstruction() ||
			p.RuleInitialFacts()
	}
	fmt.Println(p.ParseResult)
}

func (p *Parser) PopStackTwo() (rg, lg interface{}) {
	if p.Stack.Size() <= 1 {
		log.Fatal("Error parsing")
	}
	return p.Stack.PopLeft(), p.Stack.PopLeft()
}

func (p *Parser) AddInstructions(rg, lg interface{}) error {
	p.Instructions = append(p.Instructions, map[string]interface{}{
		"rg": rg,
		"lg": lg,
	})
	return nil
}
