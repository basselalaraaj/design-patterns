package hospital

type patient struct {
	name     string
	age      int
	symptoms []string
}

func NewPatient(name string, age int, symptoms []string) *patient {
	return &patient{
		name:     name,
		age:      age,
		symptoms: symptoms,
	}
}
