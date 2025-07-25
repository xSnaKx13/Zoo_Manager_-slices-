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
	var allZooAnimals []string
	var lion string
	
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
func printAllAnimal(){

}
