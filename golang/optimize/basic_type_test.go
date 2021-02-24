package optimize

import (
	"fmt"
	"testing"
)

type Employee struct {
	name string
	city string
}

func (e *Employee) String() string {
	return fmt.Sprintf("Employee[address: %p, name: %s, city: %s]", e, e.name, e.city)
}

func add(employee *Employee) {
	g_employees = append(g_employees, employee)
}

func debug() {
	count := len(g_employees)
	for i := 0; i < count; i++ {
		fmt.Printf("%d: %s\n", i, g_employees[i])
	}
	fmt.Println()
}

var g_employees = []*Employee{}

func TestAppend(t *testing.T) {
	one := Employee{name: "name1", city: "city1"}
	two := Employee{name: "name2", city: "city2"}
	three := Employee{name: "name3", city: "city3"}
	four := Employee{name: "name4", city: "city4"}

	fmt.Printf("add:%v len: %v cap:%v\n\n", g_employees, len(g_employees), cap(g_employees))

	add(&one)
	debug()

	var p *Employee
	p = g_employees[0]
	fmt.Printf("p: %s\n\n", p)
	fmt.Printf("add:%v len: %v cap:%v\n\n", g_employees, len(g_employees), cap(g_employees))

	add(&two)
	debug()
	fmt.Printf("p: %s\n\n", p)
	fmt.Printf("add:%v len: %v cap:%v\n\n", g_employees, len(g_employees), cap(g_employees))

	add(&three)
	debug()
	fmt.Printf("p: %s\n\n", p)
	fmt.Printf("add:%v len: %v cap:%v\n\n", g_employees, len(g_employees), cap(g_employees))

	add(&four)
	debug()
	fmt.Printf("p: %s\n\n", p)
	fmt.Printf("add:%p len: %v cap:%v\n\n", g_employees, len(g_employees), cap(g_employees))

}
