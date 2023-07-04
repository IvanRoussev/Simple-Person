package main

import "testing"

func TestFullName(t *testing.T) {
	type TestCase struct {
		fname    string
		lname    string
		age      int
		fullName string
	}

	PersonTestCases := []TestCase{
		{fname: "Ivan", lname: "Roussev", age: 19, fullName: "Ivan Roussev"},
		{fname: "Joe", lname: "Smith", age: 21, fullName: "Joe Smith"},
		{fname: "Garret", lname: "Jones", age: 29, fullName: "Garret Jones"},
		{fname: "Owen", lname: "Payne", age: 54, fullName: "Owen Payne"},
		{fname: "Jeremy", lname: "dawn", age: 91, fullName: "Jeremy dawn"},
		{fname: "Alex", lname: "Stuart", age: 8, fullName: "Alex Stuart"},
	}

	for _, personCase:= range PersonTestCases {
		t.Run("Full Name Testing", func(t *testing.T){
			person := CreatePerson(personCase.fname, personCase.lname, personCase.age)
			fullName := person.FullName()

			if fullName != personCase.fullName {
				t.Errorf("Expected %v, Got %v", personCase.fullName, fullName)
			}
		})
	}


}


