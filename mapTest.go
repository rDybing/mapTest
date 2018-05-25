package main

import (
	"crypto/sha1"
	"encoding/json"
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

var myMap map[string]myStructT

func main() {

	rand.Seed(time.Now().UTC().UnixNano())
	mySlice := initSlice()
	myMap = initMap(mySlice)
	outBytes, _ := json.MarshalIndent(myMap, "", "	")
	fmt.Print(string(outBytes[:]))
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
		fmt.Printf("Name  : %s X: %d\n", s.name, x)
		for j := 0; j < x; j++ {
			y := randInt(0, 5)
			s.pets = append(s.pets, pets[y])
			fmt.Printf("Pet %d : %s\n", j+1, s.pets[j])
		}
		fmt.Println("---")
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
		fmt.Println(out[hash])
	}
	return out
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
