package main

import (
	"errors"
	"fmt"
)

// Pi const
const Pi = 3.1416

func circulo(radio float64) (float64, float64, error) {

	if radio <= 0 {
		return 0, 0, errors.New("El cero no es valido")
	}

	area := Pi * radio * radio
	perimetro := 2 * Pi * radio

	return area, perimetro, nil
}

func main() {
	a, p, err := circulo(0)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Area: %v,\nPerimitro: %v\n", a, p)
}
