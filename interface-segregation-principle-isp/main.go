package main

import "log"

type Document struct {
}

type Machine interface {
	Print(d Document)
	Scan(d Document)
	Fax(d Document)
}

type MultiFunctionPrinter struct {
}

func (m MultiFunctionPrinter) Print(d Document) {}
func (m MultiFunctionPrinter) Scan(d Document)  {}
func (m MultiFunctionPrinter) Fax(d Document)   {}

type OldFashionPrinter struct {
}

func (o OldFashionPrinter) Print(d Document) {

}

// Deprecated: ...
func (o OldFashionPrinter) Scan(d Document) {
	log.Fatal("Operation not supported.")
}

// Deprecated: ...
func (o OldFashionPrinter) Fax(d Document) {
	log.Fatal("Operation not supported.")
}

// here is the ISP implementation

type Printer interface {
	Print(d Document)
}

type Scanner interface {
	Scan(d Document)
}

type MyPrinter struct {
}

func (m MyPrinter) Print(d Document) {

}

type Photocopier struct {
}

func (p Photocopier) Scan(d Document) {

}

type MultiFunctionDevice interface {
	Printer
	Scanner
	// Fax
}

// decorator
type MultiFunctionMachine struct {
	printer Printer
	scanner Scanner
}

func (m MultiFunctionMachine) Print(d Document) {
	m.printer.Print(d)
}

func (p Photocopier) Print(d Document) {

}

func main() {

}
