package main

// Interface Separation Principle

type Document struct {
}

type Machine interface {
	Print(d Document)
	Fax(d Document)
	Scan(d Document)
}

type MultiFunctionPrinter struct{}

func (m *MultiFunctionPrinter) Print(d Document) {
}

func (m *MultiFunctionPrinter) Fax(d Document) {
}

func (m *MultiFunctionPrinter) Scan(d Document) {
}

type OldFunctionPrinter struct{}

func (o *OldFunctionPrinter) Print(d Document) {
}

func (o *OldFunctionPrinter) Fax(d Document) {
}

// Deprecated Method

func (o *OldFunctionPrinter) Scan(d Document) {
}

// ISP

type Printer interface {
	Print(d Document)
}

type Scanner interface {
	Scan(d Document)
}

type MyPrinter struct{}

func (m MyPrinter) Print(d Document) {
}

type Photocopier struct{}

func (p Photocopier) Print() {
}

func (p Photocopier) Scan() {
}

type MultiFunctionDevice interface {
	Printer
	Scanner
	// FAX
}

// decorator

type MultiFunctionMachine struct {
	printer Printer
	scanner Scanner
}

func (m MultiFunctionMachine) Print(d Document) {
	m.printer.Print(d)
}

func (m MultiFunctionMachine) Scan(d Document) {
	m.scanner.Scan(d)
}

func main() {

}
