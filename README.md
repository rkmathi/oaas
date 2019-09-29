# oaas

Omikuji As A Service


# Usage

```bash
$ cd /path/to/oaas

$ protoc -I proto proto/omikuji.proto --go_out=plugins=grpc:proto

$ export GO111MODULE=ON

$ go run cmd/omikujiservice/main.go
#=> START omikuji service

$ go run cmd/bankservice/main.go
#=> START bank service

$ go run cmd/gateway/main.go
#=> START gateway
```

```bash
# UpdateBalance
$ curl -s -XPUT -d '{"delta": 300}' http://127.0.0.1:8000/balance
{"status":"OK","amount":300}

# GetBalance
$ curl -s -XGET http://127.0.0.1:8000/balance
{"amount":300}

# DrawOmikuji
$ curl -s -XPOST http://127.0.0.1:8000/omikuji
{"status":"OK","value":"chu-kichi"}
```
