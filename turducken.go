package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/gobuffalo/packr"
)

//This packs all the template(s) into the binary
var templateBox  packr.Box


func main(){
	templateBox = packr.NewBox( "./templates")
	println("I'm a turtle, hi.")
	boxStr := spew.Sdump(templateBox)
	println(boxStr)
	//render boring template

	//take input

	//rerender after network call

}
