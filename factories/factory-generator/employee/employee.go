package employee

type employeeFac func(name string) *employee

func NewEmployee(position string, annualIncome int) employeeFac {
	return func(name string) *employee {
		return &employee{
			position:     position,
			annualIncome: annualIncome,
			name:         name,
		}
	}
}

type employee struct {
	position, name string
	annualIncome   int
}

func (e *employee) Name() string {
	return e.name
}

func (e *employee) Position() string {
	return e.position
}
