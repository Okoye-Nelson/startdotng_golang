package main
import (
	"fmt"
	"strconv"
)

func main() {
	var dec_value int //variable declaration of dec_value as an integer
	fmt.Println("What is the decimal value: ")
	fmt.Scanln(&dec_value) //the value input by the user is stored in dec_value
	result := strconv.FormatInt(dec_value, 2) //the stored input is converted to binary digits an stored in result
	fmt.Println("The binary digit of "+dec_value" is: "+result)

}
