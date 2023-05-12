FROM golang:alpine
MAINTAINER NeterAlex "neteralex@foxmail.com"

ENV GO111MODULE=on \
    CGO_ENABLE=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY="https://goproxy.cn,direct"

WORKDIR /refined_server
#COPY biz/config/config.yml ./app/biz/config/
COPY config/config.yml ./config/
COPY biz/dal/database/refined.db ./biz/dal/database/
COPY out/* ./

EXPOSE 8022
RUN chmod +xwr /refined_server/
ENTRYPOINT ["/refined_server/go_build_linux_amd64_linux"]