package main

import "fmt"

// Dependency Inversion Principle
// HLM should not depend on LLM
// Both should depend on abstractions

type Relationship int

const (
	Parent Relationship = iota
	Child
	Sibling
)

type Person struct {
	name string
}

type Info struct {
	from         *Person
	relationship Relationship
	to           *Person
}

// low-level module

type RelationshipBrowser interface {
	FindAllChildrenOf(name string) []*Person
}

type Relationships struct {
	relations []Info
}

func (r *Relationships) FindAllChildrenOf(name string) []*Person {
	result := make([]*Person, 0, len(r.relations))
	for _, relation := range r.relations {
		if relation.from.name == name &&
			relation.relationship == Parent {
			result = append(result, relation.to)
		}
	}
	return result
}

func (r *Relationships) AddParentAndChild(parent *Person, child *Person) {
	r.relations = append(r.relations, Info{parent, Parent, child})
	r.relations = append(r.relations, Info{child, Child, parent})
}

// high-level module

type Research struct {
	// break DIP
	// relationships
	browser RelationshipBrowser
}

func (r *Research) Investigate() {
	for _, person := range r.browser.FindAllChildrenOf("John") {
		fmt.Println(person.name)
	}
}

func main() {

}
