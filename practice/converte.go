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
var booleano bool = false
var nombre string = "german"


func converte(){

	fmt.Println("Conversor de unidades")
	fmt.Println("1. Millas a Kilometros")
	fmt.Println("2. Kilometros a Millas")
	fmt.Println("3. Farenheit a Celsius")
	fmt.Println("4. Celsius a Farenheit")
	fmt.Println("5. Libras a Kilogramos")
	fmt.Println("6. Kilogramos a Libras")
	fmt.Scanf("%d", &option)
	fmt.Println("Ingrese el valor a convertir")
	fmt.Scanf("%f",&value)
	if(option == 1){
		value = value * mileToKm
		fmt.Println(value,"Km")
	}
	if(option ==2){
		value = value * kmToMile
		fmt.Println(value,"ml")
	}	

}