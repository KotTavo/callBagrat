package boot
import (
	viper "github.com/spf13/viper"
	tele "gopkg.in/telebot.v3"
	"time"
	"log"
	"GoProject/bot_on_telebot/handlers"

)
func Start(){

	pref := tele.Settings{
		Token:  loadCfg(),
		Poller: &tele.LongPoller{Timeout: 2 * time.Second},
	}

	bot, err := tele.NewBot(pref)
	if err != nil{
		log.Fatal(err)
	}

	handlers.NewHandler(bot)

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