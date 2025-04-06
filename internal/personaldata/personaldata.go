package personaldata

import "fmt"

type Personal struct {
	Name   string
	Weight float64
	Height float64
}

func (personal Personal) Print() {
	fmt.Printf("Имя: %s\nВес: %.2f кг.\nРост: %.2f м.\n\n", personal.Name, personal.Weight, personal.Height)
}
