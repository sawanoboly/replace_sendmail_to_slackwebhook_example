# replace_sendmail_to_slack_example


## build fake sendmail bin

```
GOOS=linux go build -ldflags="-X main.webhookURL=${YOUR_WEBHOOK_URL}" -o override_bin/sendmail main.go
```

## crond demo with docker

```
docker build -f Dockerfile -t fakesendmail:latest . && docker run -it --rm fakesendmail:latest
```
