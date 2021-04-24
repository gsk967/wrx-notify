# WRX NOTIFY 
> It will notify the cryptocurrency price value in wazirx exchange , it can send telegram notifications also



## Exchanges Json File Format
1. Get the info about  wazirx apis in github [Click link for wazirx api ](https://github.com/WazirX/wazirx-api)
2. Wazirx ticker url [Click link for get the crypto currency exchange info](https://api.wazirx.com/api/v2/tickers)

```json
{
  "exchanges": [
    {
      "coin": "WAZIRX/WRX",
      "exchange": "wrxinr",
      "high": 140,
      "low": 131
    }
  ]
}
```

### Build and Run 
```shell
# It will build the binary to copy that one into go bin folder 
make install 
```

```shell
./build/wrx-notify --exchanges-list ./exchanges.json --watch 10

or 

wrx-notify --exchanges-list ./exchanges.json --watch 10
```


### With Telegram Notification 
```shell
wrx-notify --exchanges-list ./exchanges.json --watch 10 --tg-notify true  --tg-token XXXXX --tg-chat-id XXX
```

#### Create telegram bot 
1. [Follow this link for creating the telegram bot ](https://core.telegram.org/bots#6-botfather)
2. [For Chat id Follow this link](https://sean-bradley.medium.com/get-telegram-chat-id-80b575520659) 



