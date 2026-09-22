package main

import "fmt"

type Person struct {
	Name   string
	Age    int
	Job    string
	Salary float64
}

func (p *Person) ReadData() {
	fmt.Print("Enter Name: ")
	fmt.Scanln(&p.Name)

	fmt.Print("Enter Age: ")
	fmt.Scanln(&p.Age)

	fmt.Print("Enter Job: ")
	fmt.Scanln(&p.Job)

	fmt.Print("Enter Salary: ")
	fmt.Scanln(&p.Salary)
}

func (p Person) PrintData() {
	fmt.Println("\n--- Person Details ---")
	fmt.Printf("Name   : %s\n", p.Name)
	fmt.Printf("Age    : %d\n", p.Age)
	fmt.Printf("Job    : %s\n", p.Job)
	fmt.Printf("Salary : %.2f\n", p.Salary)
}

func main() {
	var person1 Person
	var person2 Person

	fmt.Println("Enter details for Person 1:")
	person1.ReadData()
	person1.PrintData()

	fmt.Println("\nEnter details for Person 2:")
	person2.ReadData()
	person2.PrintData()
}