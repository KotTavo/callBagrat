package main

import (
	"log"
	"time"

	tele "gopkg.in/telebot.v3"

	viper "github.com/spf13/viper"
)

func main() {
	pref := tele.Settings{
		Token:  loadCfg(),
		Poller: &tele.LongPoller{Timeout: 2 * time.Second},
	}

	bot, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	var (
		menu  = &tele.ReplyMarkup{ResizeKeyboard: true}
		menu2 = &tele.ReplyMarkup{ResizeKeyboard: true}
		menu3 = &tele.ReplyMarkup{ResizeKeyboard: true}
		btn1  = menu2.Text("Назвать Баграта уёбищем")
		btn2  = menu2.Text(("Назвать Баграта уёбищем"))
		btn3  = menu2.Text(("Назвать Баграта уёбищем"))
		btn4  = menu2.Text(("Назвать Баграта уёбищем"))
		btn5  = menu.Text("Я Баграт")
		btn6  = menu.Text(("Я не Баграт"))
		btn7  = menu.Text(("Назвать себя уёбищем"))
		btn8  = menu.Text(("Назвать себя уёбищем"))
		btn9  = menu.Text(("Назвать себя уёбищем"))
		btn10 = menu.Text(("Назвать себя уёбищем"))
		btn11 = menu.Text("Вернуться в начало")
	)
	menu3.Reply(
		menu3.Row(btn7, btn8),
		menu3.Row(btn9, btn10),
		menu3.Row(btn11),
	)

	menu.Reply(
		menu.Row(btn5, btn6),
	)

	menu2.Reply(
		menu2.Row(btn1, btn2),
		menu2.Row(btn3, btn4),
		menu3.Row(btn11),
	)

	bot.Handle("/start", func(ctx tele.Context) error {
		return ctx.Send("Начнем", menu)
	})

	bot.Handle(&btn1, func(ctx tele.Context) error {
		return ctx.Send("Баграт уёбище")
	})

	bot.Handle(&btn2, func(ctx tele.Context) error {
		return ctx.Send("Баграт уёбище")
	})

	bot.Handle(&btn3, func(ctx tele.Context) error {
		return ctx.Send("Баграт уёбище")
	})

	bot.Handle(&btn4, func(ctx tele.Context) error {
		return ctx.Send("Баграт уёбище")
	})
	bot.Handle(&btn5, func(ctx tele.Context) error {
		return ctx.Send("Ладно..", menu3)
	})
	bot.Handle(&btn6, func(ctx tele.Context) error {
		return ctx.Send("Ладно..", menu2)
	})
	bot.Handle(&btn7, func(ctx tele.Context) error {
		return ctx.Send("Я уёбище")
	})
	bot.Handle(&btn8, func(ctx tele.Context) error {
		return ctx.Send("Я уёбище")
	})
	bot.Handle(&btn9, func(ctx tele.Context) error {
		return ctx.Send("Я уёбище")
	})
	bot.Handle(&btn10, func(ctx tele.Context) error {
		return ctx.Send("Я уёбище")
	})
	bot.Handle(&btn11, func(ctx tele.Context) error {
		return ctx.Send("Вернемся в начало...", menu)
	})

	bot.Start()
}
func loadCfg() string {
	config_path := "C:/GoProject/bot_on_telebot/config.json"
	viper.SetConfigType("json")
	viper.SetConfigFile(config_path)
	viper.ReadInConfig()
	token := viper.GetString("Token")
	return token
}
