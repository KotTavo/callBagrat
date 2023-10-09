package handlers

import (
	"GoProject/bot_on_telebot/keyboard"


	tele "gopkg.in/telebot.v3"
)

type Handler struct {
	Bot      *tele.Bot
	Keyboard *keyboard.Keyboard
}

func NewHandler(bot *tele.Bot, keyboard *keyboard.Keyboard) *Handler {
	return &Handler{
		Bot:      bot,
		Keyboard: keyboard,
	}
}

func (h *Handler) Start(ctx tele.Context) error {
	return ctx.Send("Начнем...", h.Keyboard.Menu1)
	
}

func (h *Handler) Cancel(ctx tele.Context) error {
	return ctx.Send("Возвращаемся назад...", h.Keyboard.Menu1)
}

func (h *Handler) NotBagrat(ctx tele.Context) error{
	return ctx.Send("Хорошо...", h.Keyboard.Menu2)
}

func (h *Handler) IsBagrat(ctx tele.Context) error{
	return ctx.Send("Хорошо...", h.Keyboard.Menu3)
}
 
func (h *Handler) BagratUeba(ctx tele.Context) error{
	return ctx.Send("Баграт уёбище")
}

func (h *Handler)  YaUeba(ctx tele.Context) error{
	return ctx.Send("Я уёбище")
}