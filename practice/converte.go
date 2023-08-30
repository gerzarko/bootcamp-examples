package main
import "fmt"

const (
	mileToKm = 1.60934
	kmToMile = 0.621371
	farenheitToCelsius = 1.8
	celsiusToFarenheit = 0.5555555555555556
	poundsToKg = 0.453592
	kgToPounds = 2.20462
)

var option int = 0
var value float32 = 0 
var qsy int64 = 10
var booleano bool = true



func main(){
	
	fmt.Println("<!DOCTYPE html>")
	fmt.Println("<head></head>")
	fmt.Println("<body>")
	
	fmt.Println("Conversor de unidades")
	fmt.Println("1. Millas a Kilometros")
	
	fmt.Scanf("%d", &option)
	fmt.Println("Ingrese el valor a convertir")
	fmt.Scanf("%f",&value)
	
	if(option == 1){
		var resultado float32 =  value * mileToKm
		fmt.Printf("+-------------------------+ \n| Miles: %0.2f \n|+-------------------------+|\n Kilometers: %0.2f       \n|+-------------------------+",resultado,value)
	}

	fmt.Println("</body>")
}	