package color

import "github.com/fatih/color"

var (
	Notice  = color.New(color.FgCyan, color.Bold).PrintlnFunc()
	Success = color.New(color.FgMagenta, color.Bold).PrintlnFunc()
	Warn    = color.New(color.FgBlue, color.Bold).PrintlnFunc()
	Debug   = color.New(color.FgBlack, color.Bold).PrintlnFunc()
	Info    = color.New(color.FgGreen, color.Bold).PrintlnFunc()
)
