package precedence

import "github.com/cyanial/go-lox/token"

const (
	None       Precedence = iota
	Assignment            // =
	Or                    // or
	And                   // and
	Equality              // == !=
	Comparison            // < > <= >=
	Term                  // + -
	Factor                // * /
	Unary                 // ! -
	Call                  // . ()
	Primary
)

type Precedence uint8

type ParseFn func()

type ParserRule struct {
	Prefix     ParseFn
	Infix      ParseFn
	Precedence Precedence
}

type Rules struct {
	r map[token.Type]*ParserRule
}

func NewRules(unary, binary, grouping, number, literal ParseFn) *Rules {
	return &Rules{
		r: map[token.Type]*ParserRule{
			token.LeftParen:    {grouping, nil, None},
			token.RightParen:   {nil, nil, None},
			token.LeftBrace:    {nil, nil, None},
			token.RightBrace:   {nil, nil, None},
			token.Comma:        {nil, nil, None},
			token.Dot:          {nil, nil, None},
			token.Minus:        {unary, binary, Term},
			token.Plus:         {nil, binary, Term},
			token.Semicolon:    {nil, nil, None},
			token.Slash:        {nil, binary, Factor},
			token.Star:         {nil, binary, Factor},
			token.Bang:         {unary, nil, None},
			token.BangEqual:    {nil, nil, None},
			token.Equal:        {nil, nil, None},
			token.EqualEqual:   {nil, nil, None},
			token.Greater:      {nil, nil, None},
			token.GreaterEqual: {nil, nil, None},
			token.Less:         {nil, nil, None},
			token.LessEqual:    {nil, nil, None},
			token.Identify:     {nil, nil, None},
			token.String:       {nil, nil, None},
			token.Number:       {number, nil, None},
			token.And:          {nil, nil, None},
			token.Class:        {nil, nil, None},
			token.Else:         {nil, nil, None},
			token.False:        {literal, nil, None},
			token.For:          {nil, nil, None},
			token.Func:         {nil, nil, None},
			token.If:           {nil, nil, None},
			token.Nil:          {literal, nil, None},
			token.Or:           {nil, nil, None},
			token.Print:        {nil, nil, None},
			token.Return:       {nil, nil, None},
			token.Super:        {nil, nil, None},
			token.This:         {nil, nil, None},
			token.True:         {literal, nil, None},
			token.Var:          {nil, nil, None},
			token.While:        {nil, nil, None},
			token.Error:        {nil, nil, None},
			token.EOF:          {nil, nil, None},
		},
	}
}

func (r *Rules) Get(typ token.Type) *ParserRule {
	return r.r[typ]
}
