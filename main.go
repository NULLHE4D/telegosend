package main

import (
    "fmt"
    "flag"
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
        log.Fatal("something went wrong: ", err)
    }
}

func checkReqErr(err error) {
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

    flag.Usage = func() {
        fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
        flag.PrintDefaults()
    }

    var msg string
    flag.StringVar(&msg, "m", "", "the message to send")

    var file string
    flag.StringVar(&file, "f", "", "read message from given file")

    flag.Parse()

    client := &http.Client {
        Timeout: 10 * time.Second,
    }

    // check if command is run from pipe
    fi, _ := os.Stdin.Stat()
    if (fi.Mode() & os.ModeCharDevice) == 0 {

        scanner := bufio.NewScanner(os.Stdin)
        var sb strings.Builder
        scanner.Scan()
        sb.WriteString(scanner.Text())
        for scanner.Scan() {
            sb.WriteString("\n")
            sb.WriteString(scanner.Text())
        }

        msg = sb.String()

    } else {

        if len(msg) == 0 && len(file) == 0 {
            flag.Usage()
            log.Fatal("no message specified")
        } else if len(msg) > 0 && len(file) > 0 {
            flag.Usage()
            log.Fatal("cannot specify both -m and -f flags at the same time")
        } else if len(file) > 0 {
            data, err := ioutil.ReadFile(file)
            checkErr(err)
            msg = string(data)
        }

        // no need to check for -m flag because that's the default

    }


    baseUrl := fmt.Sprintf("https://api.telegram.org/bot%s/", botToken)

    getMeUrl := baseUrl + "getMe"
    res, err := client.Get(getMeUrl)
    checkReqErr(err)
    data, _ := ioutil.ReadAll(res.Body)
    res.Body.Close()
    checkOk(data)

    sendMessageUrl := baseUrl + "sendMessage"
    reqBody := fmt.Sprintf("chat_id=%d&text=%s&parse_mode=MarkdownV2", chatId, msg)
    res, err = client.Post(sendMessageUrl, "application/x-www-form-urlencoded", strings.NewReader(reqBody))
    checkReqErr(err)
    data, _ = ioutil.ReadAll(res.Body)
    res.Body.Close()
    checkOk(data)

}
