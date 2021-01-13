package main

import (
    _ "fmt"
    "os"
    "log"
    "strconv"
    "bufio"
    "strings"

    telebot "gopkg.in/tucnak/telebot.v2"
)

func main() {

    botToken := os.Getenv("TGBOTTOKEN"); if len(botToken) == 0 {
        log.Fatal("missing TGBOTTOKEN environment variable")
    }

    chatId, err := strconv.Atoi(os.Getenv("TGCHATID")); if err != nil {
        log.Fatal("missing or invalid CHATID environment variable")
    }

    recipient := telebot.User{ID: chatId}

    bot, err := telebot.NewBot(telebot.Settings{
        Token:  botToken,
        ParseMode: "Markdown",
    }); if err != nil {
        log.Fatal(err)
    }

    scanner := bufio.NewScanner(os.Stdin)
    var sb strings.Builder
    for scanner.Scan() {
        sb.WriteString(scanner.Text())
        sb.WriteString("\n")
    }

    _, err = bot.Send(&recipient, sb.String()); if err != nil {
        log.Fatal(err)
    }

}
