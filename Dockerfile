FROM golang:1.21.3-alpine3.18 AS build

RUN apk --no-cache add gcc g++ make git

WORKDIR /go/src/app

COPY . .

RUN GOOS=linux go build -ldflags="-s -w" -o ./bin/menu ./main.go

# ==============

FROM alpine:3.18

RUN apk --no-cache add ca-certificates

WORKDIR /usr/bin

COPY --from=build /go/src/app/bin /go/bin

EXPOSE 8080

ENTRYPOINT /go/bin/menu --port 8080
