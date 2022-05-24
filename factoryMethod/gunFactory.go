package main

import "fmt"

func getGun(name string) (iGun, error) {
	switch name {
	case "ak47":
		return newAk47(), nil
	case "m16":
		return newM16(), nil

	default:
		return nil, fmt.Errorf("wrong gun type")
	}
}
