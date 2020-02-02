FROM golang:1.13.6-buster AS builder

ARG PROJECT_NAME
ARG PROJECT_NAMESPACE
ARG TAG
ARG COMMIT

ADD . $GOPATH/src/github.com/${PROJECT_NAMESPACE}/${PROJECT_NAME}
WORKDIR $GOPATH/src/github.com/${PROJECT_NAMESPACE}/${PROJECT_NAME}
RUN go get -v -t -d ./...
RUN env GOOS=linux GOARCH=arm go build -o /go/bin/${PROJECT_NAME} .

FROM alpine
COPY --from=builder /go/bin/${PROJECT_NAME} /app/${PROJECT_NAME}
# COPY ./${PROJECT_NAME} /app/${PROJECT_NAME}
ADD videos.list info.list config.yml /app/
WORKDIR /app
RUN chmod +x ./${PROJECT_NAME}

# ENTRYPOINT ./${PROJECT_NAME}
