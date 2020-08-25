package cmd

import (
	"strconv"
	"xsi/framework"
)

// MoneyCommand check money, give money
func MoneyCommand(ctx *framework.Context) {
	if len(ctx.Args) == 0 {
		ctx.Reply(strconv.Itoa(framework.GetMoney(ctx.User.ID)) + "원")
	} else if len(ctx.Args) >= 1 {
		if ctx.Args[0] == "give" || ctx.Args[0] == "줘" {
			giveMoney(ctx)
		}
	}
}

func giveMoney(ctx *framework.Context) {
	if len(ctx.Args) >= 2 {
		if amount, err := strconv.Atoi(ctx.Args[1]); err == nil {
			framework.GiveMoney(ctx.User.ID, amount)
			return
		}
		ctx.Reply("숫자가 아니예요ㅠㅠ")
		return
	}
	ctx.Reply("얼마를 주실건지,,")
}
