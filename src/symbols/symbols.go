package symbols

type Set struct {
	A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W, X, Y, Z bool
}

// Symbols regroups all symbols
type Symbols struct {
	OpenBrace  string
	CloseBrace string
	Not        string
	And        string
	Or         string
	Xor        string
	Implies    string
	IfOnlyIf   string
	Comment    string
	Initial    string
	Query      string
}

// Appendix struct for Symbols
type Appendix struct {
	Symbols
	Set
}

// Tokens var
var Tokens = Appendix{
	Symbols: Symbols{
		"(",
		")",
		"!",
		"+",
		"|",
		"^",
		"=>",
		"<=>",
		"#",
		"=",
		"?",
	},
	Set: Set{
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
	},
}
