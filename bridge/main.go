package main

func main() {
	mac := &mac{}

	hpPrinter := &hp{}

	mac.setPrinter(hpPrinter)
	mac.print()

	mac.setPrinter(&epson{})
	mac.print()
}
