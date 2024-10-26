## Описание

Этот бот для Telegram проверяет подписку пользователя на указанный канал. При запуске, бот приветствует пользователя и предлагает проверить статус подписки с помощью кнопки.

## Требования

- Go 1.16+
- Библиотека для работы с Telegram API: `gopkg.in/telebot.v4`

## Установка

1. Скачайте зависимости:
   ```bash
   go get gopkg.in/telebot.v4
   ```

2. Укажите `TG_TOKEN` (токен вашего бота) и `@CHANNEL` (имя канала) в коде.

## Использование

1. Запустите бота:
   ```bash
   go run main.go
   ```

2. В Telegram, отправьте `/start` для начала работы с ботом.

## Основные функции

- **Проверка подписки** — при нажатии кнопки бот проверяет, подписан ли пользователь на канал.
- **Ответы** — бот уведомляет пользователя о статусе его подписки.

## Логгирование

Логирует запуск, статус пользователя и ошибки.
