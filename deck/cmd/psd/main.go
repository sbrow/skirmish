package main

import (
	"fmt"
	"github.com/solovev/gopsd"
)

func main() {
	d, err := gopsd.ParseFromPath("./TurnSheetGuide.psd")
	if err != nil {
		panic(err)
	}
	fmt.Println(d.ToJSON())
	// fmt.Println(d.GetLayersByName("name")[0].Name)
}
