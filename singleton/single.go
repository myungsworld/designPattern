package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type single struct {
}

var singleton *single

func getInstance() {
	if singleton == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleton == nil {
			fmt.Println("first assign singleton variable")
			singleton = &single{}
		} else {
			fmt.Println("it is assigned after first")
		}
	} else {
		fmt.Println("it is assigned after first")
	}
}
