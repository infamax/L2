package main

import "fmt"

type Person struct {
	name    string
	surname string
	weight  float64
	height  float64
	job     string
}

type PersonBuilderI interface {
	Name(name string) PersonBuilderI
	Surname(surname string) PersonBuilderI
	Weight(weight float64) PersonBuilderI
	Height(height float64) PersonBuilderI
	Job(job string) PersonBuilderI
	Build() Person
}

type PersonBuilder struct {
	name    string
	surname string
	weight  float64
	height  float64
	job     string
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{}
}

func (pb *PersonBuilder) Name(name string) PersonBuilderI {
	pb.name = name
	return pb
}

func (pb *PersonBuilder) Surname(surname string) PersonBuilderI {
	pb.surname = surname
	return pb
}

func (pb *PersonBuilder) Height(height float64) PersonBuilderI {
	pb.height = height
	return pb
}

func (pb *PersonBuilder) Weight(weight float64) PersonBuilderI {
	pb.weight = weight
	return pb
}

func (pb *PersonBuilder) Job(job string) PersonBuilderI {
	pb.job = job
	return pb
}

func (pb *PersonBuilder) Build() Person {
	return Person{
		name:    pb.name,
		surname: pb.surname,
		weight:  pb.weight,
		height:  pb.height,
		job:     pb.job,
	}
}

func main() {
	personBuilder := NewPersonBuilder()
	person1 := personBuilder.Name("Ivan").Surname("Ivanov").Job("developer").Build()
	fmt.Println(person1)
}
