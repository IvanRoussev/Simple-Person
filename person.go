package main

import "fmt"

type Person struct {
	fname string
	lname string
	age   int
}

func CreatePerson(fname string, lname string, age int) *Person {
	return &Person{
		fname: fname,
		lname: lname,
		age:   age,
	}
}

func (p *Person) FullName() string {
	fn := p.fname + " " + p.lname
	fmt.Println(fn)
	return fn
}




func main() {

}