package linters

type Linter interface {
	Apply(input string) (string, error)
	SetNext(next Linter)
}

type BaseLinter struct {
	next Linter
}

func (l *BaseLinter) SetNext(next Linter) {
	l.next = next
}

func (l *BaseLinter) Apply(input string) (string, error) {
	// Base implementation (if needed)
	if l.next != nil {
		return l.next.Apply(input)
	}
	return input, nil
}
