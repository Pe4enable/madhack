FROM golang:1.10-alpine as build
ARG github_auth_token
RUN apk add --no-cache git wget gcc musl-dev curl
WORKDIR /go/src/github.com/BankEx/RatesReader/
COPY . .
RUN git config --global url."https://${github_auth_token}:x-oauth-basic@github.com/".insteadOf 'https://github.com/'
RUN go get -v
#if you want include default config
#RUN mv config.y*ml /go/bin
RUN go build -o /go/bin/RatesReader .

FROM alpine:3.8
RUN mkdir /app && chmod 755 /app
WORKDIR /app/
RUN apk add --no-cache wget ca-certificates
COPY --from=build /go/bin/RatesReader  /app/rates-reader
RUN chmod +x /app/rates-reader
RUN mkdir /etc/bankex && chmod 777 /etc/bankex
VOLUME /etc/bankex

#if you want include default config
COPY ./config ./config
COPY ./docs ./docs

EXPOSE 8002
CMD ["/app/rates-reader"]

