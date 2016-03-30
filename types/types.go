package types

type LispExpression interface {
	isAtom() bool
}

type LispInt int
type LispString string
type LispList []LispExpression

func (expr LispInt) isAtom() bool {
	return true
}

func (expr LispString) isAtom() bool {
	return true
}

func (expr LispList) isAtom() bool {
	return false
}