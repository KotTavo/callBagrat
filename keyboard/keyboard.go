package keyboard

import (
	tele "gopkg.in/telebot.v3"
)

type Buttons struct {
	Button1  tele.Btn
	Button2  tele.Btn
	Button3  tele.Btn
	Button4  tele.Btn
	Button5  tele.Btn
	Button6  tele.Btn
	Button7  tele.Btn
	Button8  tele.Btn
	Button9  tele.Btn
	Button10 tele.Btn
	Button0 tele.Btn
}

type Keyboard struct {
	Menu1 *tele.ReplyMarkup
	Menu2 *tele.ReplyMarkup
	Menu3 *tele.ReplyMarkup
}

func(k *Keyboard) LoadCfgKeyboard() {
	k.Menu1.Reply(
		NewKeyboard().Menu1.Row(NewButton().Button1, NewButton().Button2),
	)

	k.Menu2.Reply(
		NewKeyboard().Menu2.Row(NewButton().Button3, NewButton().Button4),
		NewKeyboard().Menu2.Row(NewButton().Button5, NewButton().Button6),
		NewKeyboard().Menu2.Row(NewButton().Button0),
	)

	k.Menu3.Reply(
		NewKeyboard().Menu3.Row(NewButton().Button7, NewButton().Button8),
		NewKeyboard().Menu3.Row(NewButton().Button9, NewButton().Button10),
		NewKeyboard().Menu3.Row(NewButton().Button0),
	)
}

func NewKeyboard() *Keyboard {
	return &Keyboard{
		Menu1: &tele.ReplyMarkup{ResizeKeyboard: true},
		Menu2: &tele.ReplyMarkup{ResizeKeyboard: true},
		Menu3: &tele.ReplyMarkup{ResizeKeyboard: true},
	}
}

func NewButton() *Buttons {
	return &Buttons{
		Button1: tele.Btn{
			Text: "Я Баграт",
		},
		Button2: tele.Btn{
			Text: "Я не Баграт",
		},
		Button3: tele.Btn{
			Text: "Назвать Баграта уебищем",
		},
		Button4: tele.Btn{
			Text: "Назвать Баграта уебищем",
		},
		Button5: tele.Btn{
			Text: "Назвать Баграта уебищем",
		},
		Button6: tele.Btn{
			Text: "Назвать Баграта уебищем",
		},
		Button7: tele.Btn{
			Text: "Назвать себя уебищем",
		},
		Button8: tele.Btn{
			Text: "Назвать себя уебищем",
		},
		Button9: tele.Btn{
			Text: "Назвать себя уебищем",
		},
		Button10: tele.Btn{
			Text: "Назвать себя уебищем",
		},
		Button0: tele.Btn{
			Text: "Вернуться в начало",
		},
	}
}
