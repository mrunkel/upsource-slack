FROM golang:1.8-alpine
ADD .  $GOPATH/src/upsource-slack
WORKDIR $GOPATH/src/upsource-slack
RUN  apk add --no-cache bash git openssh && \
     go get -u github.com/kardianos/govendor && \
     rm -rf vendor/**/* && \
     govendor init && \
     govendor fetch -v +outside && \
     govendor add +external && \
     go build -o $GOPATH/src/upsource-slack/bin/upsource-slack

EXPOSE 8989

ENTRYPOINT ["bin/upsource-slack"]