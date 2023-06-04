FROM golang:1.19.8-alpine3.16 AS builder

# ENV GOPROXY      https://goproxy.io

RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build -o dingtalk-wenxin .

FROM alpine:3.16

ARG TZ="Asia/Shanghai"

ENV TZ ${TZ}

RUN mkdir /app && apk upgrade \
    && apk add bash tzdata \
    && ln -sf /usr/share/zoneinfo/${TZ} /etc/localtime \
    && echo ${TZ} > /etc/timezone

WORKDIR /app
COPY --from=builder /app/ .
RUN chmod +x dingtalk-wenxin && cp config.example.yml config.yml

CMD ./dingtalk-wenxin