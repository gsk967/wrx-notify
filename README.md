# WRX NOTIFY 
> it will notify the currency value 



## Exchanges Json File Format
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
 go build -o ./build/wrx-notify cmd/main.go 
```

```shell
./build/wrx-notify --exchanges-list ./exchanges.json --watch 10
```




