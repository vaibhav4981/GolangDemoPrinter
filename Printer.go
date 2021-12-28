package main

import (
	pterm "github.com/pterm/pterm"
)

func PrintUsingStyle1(text string) {
	pterm.Success.Printfln(text)
}
