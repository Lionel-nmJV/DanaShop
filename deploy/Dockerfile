FROM golang:alpine

WORKDIR /apps

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o starfish ./main.go

CMD [ "./starfish" ]