FROM golang:1.12-alpine

ENV APP_DIR /slackconcierge

RUN apk update \
    && apk add curl git alpine-sdk \
    && curl https://glide.sh/get | sh

RUN go get -v -u github.com/DiSiqueira/CompileDaemon

COPY . ${APP_DIR}
WORKDIR ${APP_DIR}

CMD CompileDaemon -build="go build ." -command="slackconcierge" -exclude-dir=".git" -exclude-dir=".idea" -exclude-dir="vendor" -verbose=true
