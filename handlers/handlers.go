package handlers

import (
	"GoProject/bot_on_telebot/keyboard"

	tele "gopkg.in/telebot.v3"
)

type DetectBagrat struct{
	Button tele.MenuButton
	Message string
}

func NewHandler(bot *tele.Bot) *DetectBagrat{
	keyboard.NewKeyboard().Menu1.Reply(
		keyboard.NewKeyboard().Menu1.Row(keyboard.NewButton().Button1, keyboard.NewButton().Button2 ),
	)

	keyboard.NewKeyboard().Menu2.Reply(
		keyboard.NewKeyboard().Menu2.Row(keyboard.NewButton().Button3, keyboard.NewButton().Button4 ),
		keyboard.NewKeyboard().Menu2.Row(keyboard.NewButton().Button5, keyboard.NewButton().Button6 ),
		keyboard.NewKeyboard().Menu2.Row(keyboard.NewButton().Button0),
	)

	keyboard.NewKeyboard().Menu3.Reply(
		keyboard.NewKeyboard().Menu3.Row(keyboard.NewButton().Button7, keyboard.NewButton().Button8 ),
		keyboard.NewKeyboard().Menu3.Row(keyboard.NewButton().Button9, keyboard.NewButton().Button10 ),
		keyboard.NewKeyboard().Menu2.Row(keyboard.NewButton().Button0),
	)



	bot.Handle("/start", func(ctx tele.Context) error {
		return ctx.Send("Начнем...", keyboard.NewKeyboard().Menu1)
	})



	bot.Handle(keyboard.NewButton().Button1, func(ctx tele.Context) error {
		return ctx.Send("Хорошо...", keyboard.NewKeyboard().Menu2)
	})

	bot.Handle(keyboard.NewButton().Button2, func(ctx tele.Context) error {
		return ctx.Send("Хорошо...", keyboard.NewKeyboard().Menu3)
	})

	bot.Handle(keyboard.NewButton().Button3, func(ctx tele.Context) error {
		return ctx.Send("Баграт уёбище")
	})

	bot.Handle(keyboard.NewButton().Button4, func(ctx tele.Context) error {
		return ctx.Send("Баграт уёбище")
	})

	bot.Handle(keyboard.NewButton().Button5, func(ctx tele.Context) error {
		return ctx.Send("Баграт уёбище")
	})

	bot.Handle(keyboard.NewButton().Button6, func(ctx tele.Context) error {
		return ctx.Send("Баграт уёбище")
	})

	bot.Handle(keyboard.NewButton().Button7, func(ctx tele.Context) error {
		return ctx.Send("Я уёбище")
	})

	bot.Handle(keyboard.NewButton().Button8, func(ctx tele.Context) error {
		return ctx.Send("Я уёбище")
	})

	bot.Handle(keyboard.NewButton().Button9, func(ctx tele.Context) error {
		return ctx.Send("Я уёбище")
	})

	bot.Handle(keyboard.NewButton().Button10, func(ctx tele.Context) error {
		return ctx.Send("Я уёбище")
	})

	bot.Handle(keyboard.NewButton().Button10, func(ctx tele.Context) error {
		return ctx.Send("Хорошо...", keyboard.NewKeyboard().Menu1)
	})


	
	return &DetectBagrat{
		
	}
}