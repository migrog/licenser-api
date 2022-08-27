FROM golang:1.16-alpine3.13 AS build
WORKDIR /app

COPY . .
COPY go.mod .
COPY go.sum .

RUN go clean --modcache
RUN GOOS=linux go build -o security security/security.go

FROM alpine:3.13
WORKDIR /app/

RUN apk add --no-cache tzdata

ENV TZ=America/Lima

COPY --from=build /app/security .

COPY ["security/setting.development.json", "."]
COPY ["security/setting.production.json", "."]
COPY ["security/setting.staging.json", "."]

EXPOSE 3000

CMD [ "/app/security" ]