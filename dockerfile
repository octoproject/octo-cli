# build stage
FROM golang:alpine AS build-env
RUN mkdir -p src/github.com/octoproject/octo-cli
WORKDIR src/github.com/octoproject/octo-cli
ADD . .
RUN go mod download
RUN go version
WORKDIR $GOPATH/src/github.com/octoproject/octo-cli
RUN  GOOS=linux go build -o octo-cli .  \
    && chmod +x octo-cli \
    && mv octo-cli /

FROM alpine 
COPY --from=build-env octo-cli .
CMD ["./octo-cli"]
