package main

import (
	"log"
	"strings"
	"time"

	tele "gopkg.in/telebot.v4"
)

func main() {
	pref := tele.Settings{
		Token:  "TG_TOKEN",
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	btn := tele.Btn{
		Text: "Проверить подписку",
	}

	menu := &tele.ReplyMarkup{}
	menu.Reply(
		menu.Row(btn),
	)

	b.Handle(&btn, func(c tele.Context) error {
		channelUsername := "@CHANNEL"
		channelID := int64(-100) // first was -100 and contiue your group ID

		channel := &tele.Chat{
			ID:       channelID,
			Type:     "channel",
			Username: strings.TrimPrefix(channelUsername, "@"),
		}

		chatMember, err := b.ChatMemberOf(channel, c.Sender())
		if err != nil {
			log.Printf("Ошибка при проверке подписки: %v", err)
			return c.Send("❌ Вы не подписаны на канал. Пожалуйста, подпишитесь: " + channelUsername)
		}

		log.Printf("Статус пользователя: %v", chatMember.Role)

		switch chatMember.Role {
		case tele.Member, tele.Administrator, tele.Creator:
			return c.Send("✅ Спасибо за подписку!")
		default:
			return c.Send("❌ Чтобы увидеть название, подпишитесь на канал " + channelUsername)
		}
	})

	b.Handle("/start", func(c tele.Context) error {
		return c.Send("Привет! Нажмите кнопку, чтобы проверить подписку на канал:", menu)
	})

	log.Println("Бот запущен...")
	b.Start()
}
