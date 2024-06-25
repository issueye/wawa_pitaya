package main

import (
	"pitaya/internal/initialize"
	"pitaya/pages/home"

	// _ "github.com/ying32/govcl/pkgs/winappres"
	"github.com/ying32/govcl/vcl"
)

func main() {
	initialize.Initialize()

	vcl.RunApp(&home.Frm_pitaya)
}
