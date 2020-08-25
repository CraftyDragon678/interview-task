package cmd

import "xsi/framework"

// HiCommand say hi
func HiCommand(ctx *framework.Context) {
	ctx.Reply("안녕!!")
}
