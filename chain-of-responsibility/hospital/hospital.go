package hospital

func VisitDoctor(p *patient) {
	c := &cashier{}

	m := &medical{}
	m.setNext(c)

	d := &doctor{}
	d.setNext(m)

	r := &reception{}
	r.setNext(d)

	r.handle(p)
}
