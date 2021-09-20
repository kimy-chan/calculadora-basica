package main

import "fmt"

func suma(a int, b int) int {
	return a + b

}
func resta(a int, b int) int {
	return a - b

}
func multi(a int, b int) int {
	return a * b
}
func divi(a int, b int) int {
	return a / b
}
func menu() {
	var menu = "\n 1 suma \n 2 resta \n 3 multiplicacion \n 4 divicion \n 5 salir"

	fmt.Printf(menu)

}
func main() {
	var (
		datos    int
		a        int
		b        int
		contador bool = true
	)

	for contador {
		menu()
		fmt.Print("\n ingrese su eleccion: ")
		fmt.Scanln(&datos)
		switch datos {
		case 1:

			fmt.Print("ingrese el primer valor: ")
			fmt.Scanln(&a)
			fmt.Print("ingrese su segundo valor : ")
			fmt.Scanln(&b)
			resultado := suma(a, b)
			fmt.Print("el resultado es: ", resultado)
			break

		case 2:
			fmt.Print("ingrese el primer valor: ")
			fmt.Scanln(&a)
			fmt.Print("ingrese su segundo valor : ")
			fmt.Scanln(&b)
			resultado := resta(a, b)
			fmt.Print("el resultado es: ", resultado)
			break
		case 3:
			fmt.Print("ingrese el primer valor: ")
			fmt.Scanln(&a)
			fmt.Print("ingrese su segundo valor : ")
			fmt.Scanln(&b)
			resultado := multi(a, b)
			fmt.Print("el resultado es: ", resultado)
			break
		case 4:
			fmt.Print("ingrese el primer valor: ")
			fmt.Scanln(&a)
			fmt.Print("ingrese su segundo valor : ")
			fmt.Scanln(&b)
			resultado := divi(a, b)
			fmt.Print("el resultado es: ", resultado)
			break
		case 5:
			contador = false
			break

		default:
			fmt.Println("opcion invalida")

		}
	}

}
