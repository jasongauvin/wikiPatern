package export

import "fmt"

type Csv struct {
}

func (c *Csv) export() {
	fmt.Println("Export in csv")
}