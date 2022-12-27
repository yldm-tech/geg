FROM golang:1.19-alpine as build

ENV RUN_ENV local

WORKDIR /app

COPY . /app

RUN export GOPRIVATE=github.com/myetc
RUN apk add git openssh
RUN git config --global url."git@github.com:".insteadOf "https://github.com"

RUN go mod download \
    && go build -o server /app/cmd/api/*

FROM alpine

COPY --from=build /app/server /app/server
COPY --from=build /app/envs /app/envs

WORKDIR /app
EXPOSE 11000

CMD ./server $RUN_ENV