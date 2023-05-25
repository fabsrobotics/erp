FROM golang:1.20-alpine

WORKDIR /usr/src/app

COPY ./server .
RUN go mod download && go mod verify

RUN go install -mod=mod github.com/githubnemo/CompileDaemon

ENTRYPOINT [ "sh", "entrypoint.sh"]