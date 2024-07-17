FROM golang:1.22.0-alpine3.19 AS builder
LABEL stage-category=builder

RUN apk update && apk add --no-cache git

WORKDIR /go/src/CategoryService

COPY . .

RUN go build -mod=readonly -o main
RUN rm -rf vendor

FROM alpine:latest

RUN apk add --no-cache tzdata
RUN apk add busybox-extras

RUN mkdir /app
WORKDIR /app

COPY --from=builder /go/src/CategoryService/main /app
COPY --from=builder /go/src/CategoryService/resources /app/resources

RUN mkdir logs

CMD /app/main
