package boot

import (
	//"GoProject/bot_on_telebot/keyboard"
	"GoProject/bot_on_telebot/handlers"
	"GoProject/bot_on_telebot/keyboard"

	"log"
	"time"

	viper "github.com/spf13/viper"
	tele "gopkg.in/telebot.v3"
)

func Start(){
	pref := tele.Settings{
		Token:  loadCfg(),
		Poller: &tele.LongPoller{Timeout: 2 * time.Second},
	}

	bot, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	k := keyboard.NewKeyboard()

	k.LoadCfgKeyboard()

	handler := handlers.NewHandler(bot, k)

	bot.Handle("/start", handler.Start)
	bot.Handle(&keyboard.NewButton().Button0, handler.Cancel)
	bot.Handle(&keyboard.NewButton().Button1, handler.IsBagrat)
	bot.Handle(&keyboard.NewButton().Button2, handler.NotBagrat)
	bot.Handle(&keyboard.NewButton().Button3, handler.BagratUeba)
	bot.Handle(&keyboard.NewButton().Button4, handler.BagratUeba)
	bot.Handle(&keyboard.NewButton().Button5, handler.BagratUeba)
	bot.Handle(&keyboard.NewButton().Button6, handler.BagratUeba)
	bot.Handle(&keyboard.NewButton().Button7, handler.YaUeba)
	bot.Handle(&keyboard.NewButton().Button8, handler.YaUeba)
	bot.Handle(&keyboard.NewButton().Button9, handler.YaUeba)
	bot.Handle(&keyboard.NewButton().Button10, handler.YaUeba)

	

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