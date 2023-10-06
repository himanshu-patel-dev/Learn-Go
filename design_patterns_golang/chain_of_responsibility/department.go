package main

import "fmt"

// Patient interface on which execution is done
type Patient struct {
	name              string
	registrationDone  bool
	doctorCheckUpDone bool
	medicineDone      bool
	paymentDone       bool
}

// Department interface, which operates on patient
type Department interface {
	execute(*Patient)
	setNext(Department)
}

// ---------- concrete handler ----------

// Reception first stage of patient
type Reception struct {
	next Department
}

func (r *Reception) execute(p *Patient) {
	if p.registrationDone {
		fmt.Println("Registration already done")
	} else {
		fmt.Println("Done Registration")
		p.registrationDone = true
	}
	r.next.execute(p)
}

func (r *Reception) setNext(d Department) {
	r.next = d
}

// Doctor second stage of patient
type Doctor struct {
	next Department
}

func (d *Doctor) execute(p *Patient) {
	if p.doctorCheckUpDone {
		fmt.Println("Doctor checkup done already")
	} else {
		fmt.Println("Doctor checkup done")
		p.doctorCheckUpDone = true
	}
	d.next.execute(p)
}

func (d *Doctor) setNext(dpt Department) {
	d.next = dpt
}

// Medical third stage of patient
type Medical struct {
	next Department
}

func (m *Medical) execute(p *Patient) {
	if p.medicineDone {
		fmt.Println("Medicine already done")
	} else {
		fmt.Println("Medicine Done")
		p.medicineDone = true
	}
	m.next.execute(p)
}

func (m *Medical) setNext(d Department) {
	m.next = d
}

// Cashier fourth stage of patient
type Cashier struct {
	next Department
}

func (c *Cashier) execute(p *Patient) {
	if p.paymentDone {
		fmt.Println("Payment already done")
	} else {
		fmt.Println("Payment Done")
		p.paymentDone = true
	}
}

func (c *Cashier) setNext(d Department) {
	c.next = d
}
