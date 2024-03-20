package main

import "fmt"

type ModernPrinter interface {
	PrintStored() string
}

type LegacyPrinter interface {
	Print(s string) string
}

type MyLegacyPrinter struct{}

func (p *MyLegacyPrinter) Print(s string) string {
	result := "Legacy Printer: " + s
	fmt.Println(result)
	return result
}

type PrinterAdapter struct {
	OldPrinter LegacyPrinter
	Msg        string
}

func (p *PrinterAdapter) PrintStored() string {
	if p.OldPrinter != nil {
		return p.OldPrinter.Print("Adapter: " + p.Msg)
	} else {
		return p.Msg
	}
}

func main() {
	// Using the new system directly
	modernPrinter := &PrinterAdapter{Msg: "Hello World"}
	fmt.Println(modernPrinter.PrintStored())

	// Adapting the old system
	legacyPrinter := &MyLegacyPrinter{}
	adapter := &PrinterAdapter{
		OldPrinter: legacyPrinter,
		Msg:        "Hello World",
	}
	fmt.Println(adapter.PrintStored())
}
