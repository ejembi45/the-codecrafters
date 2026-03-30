package main
import "fmt"

func main() {
	for {
	var sum1 int
	fmt.Println("enter sum1:")
	fmt.Scanln(&sum1)
	var sum2 int
	fmt.Println("enter sum2: ")
	fmt.Scanln(&sum2)
	fmt.Println("1.adition")
    fmt.Println("2.subtraction")
    fmt.Println("3.multiplication")
    fmt.Println("4.division")


var operation int
fmt.Println("choose your input")
fmt.Scan(&operation)
if operation != 1 && operation != 2 && operation != 3 && operation !=4 {
	fmt.Print("invalid input, try again")
	continue
}
switch operation {
case 1:
	fmt.Println(sum1 + sum2)

case 2: 
fmt.Println(sum1 - sum2)

case 3:
	fmt.Println(sum1 * sum2)

case 4:
	fmt.Println(sum1 / sum2)
}
break
}
}









	
