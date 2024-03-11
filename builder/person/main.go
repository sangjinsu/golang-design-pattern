package main

import "fmt"

type Person struct {
	StreetAddress, Postcode, City string
	CompanyName, Position         string
	AnnualIncome                  int
}

type PersonBuilder struct {
	person *Person
}

func (b *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{*b}
}

type PersonJobBuilder struct {
	PersonBuilder
}

func (b *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{*b}
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{&Person{}}
}

type PersonAddressBuilder struct {
	PersonBuilder
}

func (it *PersonAddressBuilder) At(streetAddress string) *PersonAddressBuilder {
	it.person.StreetAddress = streetAddress
	return it
}

func (it *PersonAddressBuilder) In(city string) *PersonAddressBuilder {
	it.person.City = city
	return it
}
func (it *PersonAddressBuilder) WithPostcode(postcode string) *PersonAddressBuilder {
	it.person.Postcode = postcode
	return it
}

func (b *PersonBuilder) Build() *Person {
	return b.person
}

func main() {
	person := NewPersonBuilder().Lives().At("").Build()
	fmt.Println(person)
}
