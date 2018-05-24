package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type myStructT struct {
	name string
	pets []string
}

var myMap map[int]myStructT

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	//myMap := make(map[int]myStructT)
	myMap = initMap()
	outBytes, _ := json.MarshalIndent(myMap, "", "	")
	fmt.Print(string(outBytes[:]))
}

func initMap() map[int]myStructT {
	var s myStructT
	m := make(map[int]myStructT)
	name := [3]string{"Ole", "Dole", "Doffen"}
	pets := [6]string{"Cat", "Dog", "Ferret", "Bunny", "Hamster", "Ferret"}
	for i := range name {
		s.name = name[i]
		x := randInt(0, 3)
		fmt.Printf("Name: %s X: %d\n", s.name, x)
		if x > 0 {
			s.pets = nil
			for j := 0; j < x; j++ {
				y := randInt(0, 5)
				s.pets = append(s.pets, pets[y])
				fmt.Printf("Pet %d : %s\n", j+1, s.pets[j])
			}
		}
		m[i] = s
	}
	return m
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
