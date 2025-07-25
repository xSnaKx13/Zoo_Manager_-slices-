package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
type animal struct{
	animalName string
	animalSpecies string
}
func main(){
	var zoo []animal
	var lion animal
	var wolf animal
	wolf.animalName = "Волк"
	wolf.animalSpecies = "Собачьи"
	zoo = append(zoo, wolf)

	addNewAnumal(&lion)
	zoo = append(zoo, lion)

	printAllAnimal(zoo)
	
}
func addNewAnumal(animal *animal){
	for{
	    var err error
	    fmt.Print("Имя животного: ")
	    animal.animalName, err = bufio.NewReader(os.Stdin).ReadString('\n')
	    animal.animalName = strings.TrimSpace(animal.animalName)
	    if err != nil{
		    fmt.Println("Введены невалидные данные")
		}
		fmt.Print("Введите вид животного (кошачьи/собачьи/попугай...: )")
		animal.animalSpecies, err = bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil{
			fmt.Println("Введены невалидные данные")
		}
		if err == nil{
			break
		}
	}
}
func printAllAnimal(zoo[]animal){
	for _, animal := range zoo{
		fmt.Println(animal)
	}
}
