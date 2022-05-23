package main

import "fmt"

func main() {
	aFactory, _ := getSportsFactory("adidas")
	nFactory, _ := getSportsFactory("nike")

	adidasShoe := aFactory.makeShoe()
	nikeShirt := nFactory.makeShirt()

	printShoeDetail(adidasShoe)
	printShirtDetail(nikeShirt)
}

func printShoeDetail(s iShoe) {
	fmt.Println(s.getLogo())
	fmt.Println(s.getSize())
}

func printShirtDetail(s iShoe) {
	fmt.Println(s.getLogo())
	fmt.Println(s.getSize())
}
