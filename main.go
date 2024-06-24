package main

import (
	"pitaya/pages/home"

	// _ "github.com/ying32/govcl/pkgs/winappres"
	"github.com/ying32/govcl/vcl"
)

func main() {
	vcl.RunApp(&home.FrmPitaya)
}
