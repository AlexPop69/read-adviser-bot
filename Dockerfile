FROM golang:1.22

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o read-adviser-bot

CMD ./read-adviser-bot -tg-bot-token $TG_BOT_TOKEN