FROM golang as build

ENV GOPROXY=https://goproxy.io

ADD . /cmall

WORKDIR /cmall

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api_server

FROM alpine:3.7

ENV REDIS_ADDR=""
ENV REDIS_PW=""
ENV REDIS_DB=""
ENV MysqlDSN=""
ENV GIN_MODE="release"
ENV JWT_SECRET = ""
ENV OSS_END_POINT=""
ENV OSS_ACCESS_KEY_ID=""
ENV OSS_ACCESS_KEY_SECRET=""
ENV OSS_BUCKET=""
ENV PORT=3000


RUN echo "http://mirrors.aliyun.com/alpine/v3.7/main/" > /etc/apk/repositories && \
    apk update && \
    apk add ca-certificates && \
    echo "hosts: files dns" > /etc/nsswitch.conf && \
    mkdir -p /www/conf

WORKDIR /www

COPY --from=build /cmall/api_server /usr/bin/api_server
ADD ./conf /www/conf

RUN chmod +x /usr/bin/api_server

ENTRYPOINT ["api_server"]