package main

import "fmt"

func main() {

	ak47, _ := getGun("ak47")
	fmt.Println(ak47.getName())
	fmt.Println(ak47.getPower())

	m16, _ := getGun("m16")
	fmt.Println(m16.getName())
	fmt.Println(m16.getPower())
	m16.setName("m116")
	fmt.Println(m16.getName())
}
