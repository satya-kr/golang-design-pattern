package main

import "fmt"

// Dependency Inversion Principle
// HLM(High-Level Modules) should not depend on LLM(Low-Level Modules)
// Both should depend on abstraction

type Relationship int

const (
	Parent Relationship = iota
	Child
	Sibling
)

type Person struct {
	name string
	// age  int
}

type Info struct {
	from         *Person      //here is a example kiran is
	relationship Relationship // parent of
	to           *Person      // of satya
}

// low-level module - this is a storage object or a database
type RelationshipBrowser interface {
	FindAllChildrenOf(name string) []*Person
}
type Relationships struct {
	relations []Info
}

func (r *Relationships) FindAllChildrenOf(name string) []*Person {
	result := make([]*Person, 0)
	for i, v := range r.relations {
		if v.relationship == Parent && v.from.name == name {
			result = append(result, r.relations[i].to)
		}
	}

	return result
}

func (r *Relationships) AddParentAndChild(parent, child *Person) {
	r.relations = append(r.relations, Info{
		parent, Parent, child,
	})

	r.relations = append(r.relations, Info{
		child, Child, parent,
	})
}

// high-level module - and here this is operate data an make some kind of research
type Research struct {
	/*
		Here we break DIP
		because Research(HLM) and we make this depend on
		Relationships(LLM)
	*/
	// relationships Relationships

	browser RelationshipBrowser
}

func (r *Research) Investigate(name string) {
	// relations := r.relationships.relations
	// for _, rel := range relations {
	// 	if rel.from.name == "John" && rel.relationship == Parent {
	// 		fmt.Println("John has a child called ", rel.to.name)
	// 	}
	// }

	for _, p := range r.browser.FindAllChildrenOf(name) {
		fmt.Println("John has a child called ", p.name)
	}
}

func main() {

	parent := Person{"John"}
	child1 := Person{"Chris"}
	child2 := Person{"Mathew"}

	relationships := Relationships{}
	relationships.AddParentAndChild(&parent, &child1)
	relationships.AddParentAndChild(&parent, &child2)

	r := Research{&relationships}
	r.Investigate("John")
	fmt.Println("Parent: ", parent)

	for _, v := range relationships.relations {
		fmt.Printf("\n============================\n")
		fmt.Println("Relationships From: ", v.from)
		fmt.Println("Relationships: ", v.relationship)
		fmt.Println("Relationships To: ", v.to)
		fmt.Printf("============================\n")
	}

}
