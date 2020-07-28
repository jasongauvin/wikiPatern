package export

import "fmt"

type Xlsx struct {
}

func (x *Xlsx) export() {
	fmt.Println("Export in csv")
}