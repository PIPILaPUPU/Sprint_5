package personaldata

import "fmt"

type Personal struct {
	// TODO: done
	Name   string
	Weight float64
	Height float64
}

func (p Personal) Print() {
	// TODO: done
	fmt.Printf("Имя: %s\nВес: %.2f кг.\nРост: %.2f м.\n", p.Name, p.Weight, p.Height)
}
