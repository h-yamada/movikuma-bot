# About Movikuma
iOS App for sharing movies to your friend.
https://www.movikuma.tv/

# movikuma-bot

## Overview
provide popular movies of Movikuma by Facebook Messenger, written in golang.

Facebook Messenger : [quickstart](https://developers.facebook.com/docs/messenger-platform/quickstart)

## Installation

```
go get github.com/h-yamada/movikuma-bot
```

- Libraries
 - gin
 ``` go get github.com/gin-gonic/gin ```
 - toml
 ``` go get github.com/BurntSushi/toml ```
 - govalidator
 ``` go get github.com/asaskevich/govalidator ```
 - elastic
 ``` go get gopkg.in/olivere/elastic.v2 ```
 - facebook/messanger
 ``` go get github.com/ymd38/facebook/messenger ```

## Usage
```
$movikuma-bot -h
Usage of movikuma-bot:
  -config-path string
    	config-path (default "./movikuma-bot.toml")
```

### Configuration

see [sample-file](https://github.com/h-yamada/movikuma-bot/blob/master/movikuma-bot.toml).
