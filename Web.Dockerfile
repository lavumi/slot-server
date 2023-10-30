# Stage 1: 빌드 스테이지
FROM golang:1.19 AS builder

WORKDIR /
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o web-server cmd/web.go

FROM alpine:latest

# Creates an app directory to hold your app’s source code
WORKDIR /

ADD ./web ./web
COPY --from=builder ./web-server ./

# Tells Docker which network port your container listens on
EXPOSE 8081

RUN apk add --no-cache bash
# Specifies the executable command that runs when the container starts
CMD [ "/web-server" ]