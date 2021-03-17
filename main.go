package main

import (
    "fmt"
    "os"
    "log"
    "strconv"
    "io/ioutil"
    "bufio"
    "strings"
    "time"
    "net/http"
    "net/url"
)

func checkErr(err error) {
    if err != nil {
        urlErr := err.(*url.Error)
        if urlErr.Timeout() {
            log.Fatal("request timed out")
        } else {
            log.Fatal("something went wrong: ", err)
        }
    }
}

func checkOk(data []byte) {
    if dataStr := string(data); strings.Contains(dataStr, "\"ok\":false") {
        log.Fatal("something went wrong: ", dataStr)
    }
}

func main() {

    botToken := os.Getenv("TGBOTTOKEN"); if len(botToken) == 0 {
        log.Fatal("missing TGBOTTOKEN environment variable")
    }

    chatId, err := strconv.Atoi(os.Getenv("TGCHATID")); if err != nil {
        log.Fatal("missing or invalid TGCHATID environment variable")
    }

    client := &http.Client {
        Timeout: 10 * time.Second,
    }

    scanner := bufio.NewScanner(os.Stdin)
    var sb strings.Builder
    scanner.Scan()
    sb.WriteString(scanner.Text())
    for scanner.Scan() {
        sb.WriteString("\n")
        sb.WriteString(scanner.Text())
    }

    baseUrl := fmt.Sprintf("https://api.telegram.org/bot%s/", botToken)

    getMeUrl := baseUrl + "getMe"
    res, err := client.Get(getMeUrl)
    checkErr(err)
    data, _ := ioutil.ReadAll(res.Body)
    res.Body.Close()
    checkOk(data)

    sendMessageUrl := baseUrl + "sendMessage"
    reqBody := fmt.Sprintf("chat_id=%d&text=%s&parse_mode=MarkdownV2", chatId, sb.String())
    res, err = client.Post(sendMessageUrl, "application/x-www-form-urlencoded", strings.NewReader(reqBody))
    checkErr(err)
    data, _ = ioutil.ReadAll(res.Body)
    res.Body.Close()
    checkOk(data)

}
