package main
import "fmt"
const nombre string = "aArman"
const newNombre string = ""

func stringExamples(){
	// array := [3]int {1,2,3}


	// for i:= range array{
	// 	var resultado int = i * array[i]
	// 	fmt.Println(" %i por %i y el resultado %i",i,array[i],resultado)
	// }

	for i ,c  := range nombre{
		fmt.Printf("range = %i y %i",i,c)
	}

	
}