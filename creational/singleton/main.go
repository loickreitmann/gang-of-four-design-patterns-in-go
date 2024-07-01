/*
Singleton Example:
Ensures a class has only one instance and provides
a global point of access to it.
*/
package main

import (
	"fmt"
	"sync"
)

type Singleton struct {
	Name string
}

var instance *Singleton
var once sync.Once

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{Name: "Singleton Instance"}
	})
	return instance
}

func main() {
	s1 := GetInstance()
	s2 := GetInstance()

	fmt.Println("Instance 1:", s1.Name)
	fmt.Println("Instance 2:", s2.Name)

	s2.Name = "Updated Singleton Instance"
	fmt.Println("Instance 1:", s1.Name)
	fmt.Println("Instance 2:", s2.Name)
}
