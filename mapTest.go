package main

import (
	"crypto/sha1"
	"fmt"
	"hash"
	"io"
	"math/rand"
	"time"
)

type myStructT struct {
	name string
	pets []string
}

func main() {
	var myMap map[string]myStructT
	myMap = make(map[string]myStructT)

	rand.Seed(time.Now().UTC().UnixNano())
	mySlice := initSlice()
	myMap = initMap(mySlice)
	printMap(myMap)
	myMap = editMap(myMap, "Dewey")
	fmt.Println("----- After edit -----")
	printMap(myMap)
}

func editMap(in map[string]myStructT, index string) map[string]myStructT {
	var h hash.Hash
	var hash string
	var old myStructT

	h = sha1.New()
	io.WriteString(h, index)
	hash = fmt.Sprintf("%x", h.Sum(nil))

	old = in[hash]
	delete(in, hash)
	old.pets = append(old.pets, "Alligator")
	in[hash] = old
	return in
}

func printMap(in map[string]myStructT) {
	for i := range in {
		fmt.Printf("Name: %s\n", in[i].name)
		if len(in[i].pets) > 0 {
			fmt.Printf("Pets: %d\n", len(in[i].pets))
			for j := range in[i].pets {
				fmt.Printf("- %s\n", in[i].pets[j])
			}
		} else {
			fmt.Println("No Pets :(")
		}
		fmt.Println("----------")
	}
}

func initSlice() []myStructT {
	var s myStructT
	var out []myStructT

	name := [3]string{"Huey", "Luie", "Dewey"}
	pets := [6]string{"Cat", "Dog", "Ferret", "Bunny", "Hamster", "Parrot"}

	for i := range name {
		s.name = name[i]
		s.pets = nil
		x := randInt(0, 3)
		for j := 0; j < x; j++ {
			y := randInt(0, 6)
			s.pets = append(s.pets, pets[y])
		}
		out = append(out, s)
	}
	return out
}

func initMap(in []myStructT) map[string]myStructT {
	var h hash.Hash
	var hash string
	var s myStructT

	out := make(map[string]myStructT)

	for i := range in {
		h = sha1.New()
		io.WriteString(h, in[i].name)
		hash = fmt.Sprintf("%x", h.Sum(nil))
		s = in[i]
		out[hash] = s
	}
	return out
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
