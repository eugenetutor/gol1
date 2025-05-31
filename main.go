package main

import (
	"fmt"
)

//структури - це визначений користувачем тип, який дозволяє групувати, комбінувати елементи різних типів в один тип. 
//struct is value type. NOT REFERENCE TYPE
type Student struct{
	Name string
	Age int 
	Group string
	Avarage float64
}

type School struct{
	Students []Student
}

func main() {
	student1 := Student{
		Name: "Alex",
		Age: 13,
		Group: "IT-101",
		Avarage: 8.5,
	}
	student2 := Student{
		Name: "Maria",
		Age: 12,
		//group, avg -> null
	}

	var student3 Student
	student3.Name = "Andrew"
	student3.Age = 18
	student3.Group = "IT-102"
	student3.Avarage = 9.2

	fmt.Printf("Student 1: %+v\n", student1)
	fmt.Printf("Student 2: %+v\n", student2)
	fmt.Printf("Student 2: %+v\n", student3)

	fmt.Printf("Student1 name: %s\n", student1.Name)
 
	studentsSlice := []Student{student1, student2, student3}
	school := School{
		Students: studentsSlice,
	}
	fmt.Printf("School: %v\n", school)
}
