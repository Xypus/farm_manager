package main

import "fmt"
import "github.com/Xypus/farm_manager/species"


// farmAnimals := make(map[int]string)

type Animal struct {
	species string
	name string
	age int
}

var animal Animal

func main() {
	fmt.Printf("Enter animal's species:\n Cow :1 \n Pig: 2 \n Horse: 3\n Chicken: 4\n")
    var input int
    fmt.Scanf("%d", &input)
    animal.species = species.AnimalSpecies(input)
	fmt.Println(animal.species)

    // fmt.Printf("Enter animal's age: \n")
    // var input int
    // fmt.Scanf("%d", &input)
    // animal.age = input
   
}











