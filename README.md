# tgbot-msg
A simple program that send message with telegram bot API

This project provide two executable and one sample env file

tgbot-msg:
go binary, simply send message to telegram
use 'TGBOT_MSG_TOKEN' and 'TGBOT_MSG_CHATID' env var

process-notify:
bash script that wrap other command, and notify through tgbot-msg after
command finished.

Follow the tutorial to get your token and chat ID
- [From BotFather to 'Hello World'](https://core.telegram.org/bots/tutorial)
- [Introduction to the API · python-telegram-bot/python-telegram-bot Wiki](https://github.com/python-telegram-bot/python-telegram-bot/wiki/Introduction-to-the-API)

## Installation

```
tgbot-msg_Linux_x86_64/
├── bin/
│  ├── process-notify*
│  └── tgbot-msg*
└── tgbot.env
```

- move two binaries into PATH
- move tgbot.env to ~/.tgbot.env

## Usage

put `process-notify` before the command that you want to monitor

```bash
process-notify sleep 3
```

you will receive the message

```
host: machine
PWD: /home/foo/bar
$ sleep 3
ret = 0, take 00:00:03
```

if the command fail,

```bash
process-notify false
```

you will receive the message

```
host: machine
PWD: /home/foo/bar
$ false
ret = 1, take 00:00:00
```

Change `process-notify` to custom the message format you want
