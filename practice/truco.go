package main
import (
	"fmt"
	"math/rand"
) 




type Cards struct{
	number uint8
	palo string	
}

type Player struct{
	name string
	cartas Cards
}

var mazoDeCartas [40] Cards;


for i:= 0;i <=12; i++{

	mazoDeCartas[i] = 
}



func truco(){

var aleatorio int = rand.Intn(100);
fmt.Println("bienvenido al truco")
fmt.Println(aleatorio)



}



