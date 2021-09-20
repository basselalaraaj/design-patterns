package employee

type position int

const (
	Developer position = iota
	Manager
	Boss
)

func NewEmployee(position position) *employee {
	switch position {
	case Developer:
		return &employee{position: "developer", annualIncome: 60000}
	case Manager:
		return &employee{position: "manager", annualIncome: 80000}
	case Boss:
		return &employee{position: "boss", annualIncome: 100000}
	}
	return nil
}

type employee struct {
	name, position string
	annualIncome   int
}

func (e *employee) SetName(name string) {
	e.name = name
}

func (e *employee) Name() string {
	return e.name
}

func (e *employee) Position() string {
	return e.position
}
