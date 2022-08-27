FROM golang:1.16-alpine3.13 AS build
WORKDIR /app

COPY . .
COPY go.mod .
COPY go.sum .

RUN go clean --modcache
RUN GOOS=linux go build -o license license/license.go

FROM alpine:3.13
WORKDIR /app

RUN apk add --no-cache tzdata

ENV TZ=America/Lima

COPY --from=build /app/license .

COPY ["license/setting.development.json", "."]
COPY ["license/setting.production.json", "."]
COPY ["license/setting.staging.json", "."]

EXPOSE 3002

CMD [ "/app/license" ]