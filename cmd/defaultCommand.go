package cmd

import "xsi/framework"

// DefaultCommand help?
func DefaultCommand(ctx *framework.Context) {
	ctx.Reply("왜?")
}
