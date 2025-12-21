package employees

import "fmt"

type FullTime struct{}

func (f *FullTime) GetDetails() {
	fmt.Println("Full time employee")
}
