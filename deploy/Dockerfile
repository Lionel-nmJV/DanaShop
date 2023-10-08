FROM golang:alpine as builder

WORKDIR /apps

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o starfish ./main.go

FROM scratch

WORKDIR /apps

COPY --from=builder /apps/starfish .

CMD [ "./starfish" ]