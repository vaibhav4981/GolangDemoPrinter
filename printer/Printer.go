package vb_print

import (
	pterm "github.com/pterm/pterm"
)

func PrintUsingStyle1(text string) {
	pterm.Success.Printfln(text)
}
