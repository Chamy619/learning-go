package main

import "fmt"

type Employee struct {
	Name string
	ID   string
}

func (e Employee) Description() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

type Manager struct {
	Employee
	Reports []Employee
}

func (m Manager) FindNewEmployees() []Employee {
	return m.Reports
}

type Inner struct {
	X int
}

type Outer struct {
	Inner
	X int
}

func main() {
	m := Manager{
		Employee: Employee{
			Name: "Bob bobson",
			ID:   "12345",
		},
		Reports: []Employee{},
	}
	fmt.Println(m.ID)
	fmt.Println(m.Description())

	o := Outer{
		Inner: Inner{
			X: 10,
		},
		X: 20,
	}

	fmt.Println(o.X)
	fmt.Println(o.Inner.X)
}
