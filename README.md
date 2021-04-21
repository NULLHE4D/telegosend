# telegosend

A command-line tool written in Golang that simply sends a message from your [Telegram Bot](https://core.telegram.org/bots) to a Telegram chat using the [Telegram Bot API](https://core.telegram.org/bots/api). Yup..that's it.

## Table of contents

   * [Usage](#usage)
      * [Command line options](#command-line-options)
      * [Examples](#examples)
   * [Installation](#installation)
      * [Dependencies](#dependencies)
      * [Instructions](#instructions)
   * [Gotchas](#gotchas)
   * [Links](#links)
   * [License](#license)

Created by [gh-md-toc](https://github.com/ekalinin/github-markdown-toc)

## Usage

:warning: **IMPORTANT: telegosend will read your [Telegram Bot Token](https://core.telegram.org/bots#3-how-do-i-create-a-bot) and [Chat ID](https://telegram.me/userinfobot) from the environment variables `TG_BOT_TOKEN` and `TG_CHAT_ID` respectively.**

### Command line options

```console
foo@bar:~$ telegosend -h
Usage of telegosend:
  -f string
    	file containing the message to send
  -m string
    	the message to send
```

### Examples

In Linux, you can get notified when a certain process exits by running the following (maybe in a tmux session):
```
tail --pid 1337 -f /dev/null; telegosend -f done-msg.txt
```

Or maybe you want to get notified if a certain command exited with a non-zero exit code:
```
./my-awesome-script.sh || telegosend -m "not so awesome after all"
```

You can also specify the message to send via stdin:
```
echo "Hello World!" | telegosend
```

## Installation

### Dependencies

- Go v1.10 or later

### Instructions

Assuming you have [Go](https://golang.org/doc/install) installed and configured, run the following command in a shell to download and install telegosend:
```
go get -v github.com/NULLHE4D/telegosend
```

## Gotchas

* In case Telegram decides to change the behaviour or structure (unlikely) of their API, this tool may no longer work until it is updated accordingly.
* Messages sent using telegosend are parsed as [MarkdownV2](https://core.telegram.org/bots/api#markdownv2-style).

## Links

* [Telegram Bot API](https://core.telegram.org/bots/api)
* [How do I create a bot](https://core.telegram.org/bots/#3-how-do-i-create-a-bot)
* [userinfobot](https://telegram.me/userinfobot)

## License

This project is unworthy of a license. Do anything you like with it. :wink:
