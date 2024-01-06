package compiler

//
//const (
//	PrecedenceNone       Precedence = iota
//	PrecedenceAssignment            // =
//	PrecedenceOr                    // or
//	PrecedenceAnd                   // and
//	PrecedenceEquality              // == !=
//	PrecedenceComparison            // < > <= >=
//	PrecedenceTerm                  // + -
//	PrecedenceFactor                // * /
//	PrecedenceUnary                 // ! -
//	PrecedenceCall                  // . ()
//	PrecedencePrimary
//)
//
//type Precedence uint8
//
//type Parser struct {
//	Current   *Token
//	Previous  *Token
//	HadError  bool
//	PanicMode bool
//}
//
//type ParseFn func()
//
//type ParserRule struct {
//	Prefix     ParseFn
//	infix      ParseFn
//	Precedence Precedence
//}
