package employees

import "fmt"

type PartTime struct{}

func (p *PartTime) GetDetails() {
	fmt.Println("Part time employee")
}
