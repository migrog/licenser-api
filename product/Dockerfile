FROM golang:1.16-alpine3.13 AS build
WORKDIR /app

COPY . .
COPY go.mod .
COPY go.sum .

RUN go clean --modcache
RUN GOOS=linux go build -o product product/product.go

FROM alpine:3.13
WORKDIR /app

RUN apk add --no-cache tzdata

ENV TZ=America/Lima

COPY --from=build /app/product .

COPY ["product/setting.development.json", "."]
COPY ["product/setting.production.json", "."]
COPY ["product/setting.staging.json", "."]

EXPOSE 3001

CMD [ "/app/product" ]