package personaldata

import "fmt"

// Personal holds user's personal data.
type Personal struct {
	Name   string
	Weight float64
	Height float64
}

// Print outputs the personal data to stdout.
func (p Personal) Print() {
	fmt.Printf("Имя: %s\nВес: %.2f\nРост: %.2f\n", p.Name, p.Weight, p.Height)
}
