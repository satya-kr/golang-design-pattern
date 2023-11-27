package main

type Document struct{}

// Printer interface
type Printer interface {
	Print(d Document)
}

// Scanner interface
type Scanner interface {
	Scan(d Document)
}

// Faxer interface (optional, commented out to show how to add it if needed)
// type Faxer interface {
// 	Fax(d Document)
// }

// MultiFunctionDevice interface
type MultiFunctionDevice interface {
	Printer
	Scanner
	// Faxer
}

// MyPrinter implements Printer
type MyPrinter struct{}

func (m MyPrinter) Print(d Document) {
	// Implementation for printing
}

// Photocopier implements Scanner
type Photocopier struct{}

func (p Photocopier) Scan(d Document) {
	// Implementation for scanning
}

// MultiFunctionMachine implements MultiFunctionDevice
type MultiFunctionMachine struct {
	printer Printer
	scanner Scanner
	// faxer   Faxer
}

func (m MultiFunctionMachine) Print(d Document) {
	m.printer.Print(d)
}

func (m MultiFunctionMachine) Scan(d Document) {
	m.scanner.Scan(d)
}

// Uncomment the following method if Fax functionality is added to MultiFunctionMachine
// func (m MultiFunctionMachine) Fax(d Document) {
// 	m.faxer.Fax(d)
// }

func main() {
	// Example usage
	myPrinter := MyPrinter{}
	photocopier := Photocopier{}
	multiFunctionMachine := MultiFunctionMachine{
		printer: myPrinter,
		scanner: photocopier,
		// faxer:   faxMachine,
	}

	document := Document{}

	// Using Printer
	multiFunctionMachine.Print(document)

	// Using Scanner
	multiFunctionMachine.Scan(document)

	// Uncomment the following line if Fax functionality is added to MultiFunctionMachine
	// multiFunctionMachine.Fax(document)
}
