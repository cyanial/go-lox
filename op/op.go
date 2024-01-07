package op

const (
	Constant Code = iota
	ConstantLong
	Nil
	True
	False
	Add
	Subtract
	Multiply
	Divide
	Negate
	Return
)

type Code byte

func (c *Code) String() string {
	switch *c {
	case Constant:
		return "Constant"
	case ConstantLong:
		return "ConstantLong"
	case Nil:
		return "Nil"
	case True:
		return "True"
	case False:
		return "False"
	case Add:
		return "Add"
	case Subtract:
		return "Subtract"
	case Multiply:
		return "Multiply"
	case Divide:
		return "Divide"
	case Negate:
		return "Negate"
	case Return:
		return "Return"
	}
	return ""
}
