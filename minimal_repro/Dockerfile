
FROM golang:1.20.5-buster
WORKDIR /usr/src/app

COPY go.mod go.* ./
RUN go mod download && go mod verify
RUN go install github.com/go-delve/delve/cmd/dlv@latest

COPY . .
RUN go build -gcflags="all=-N -l" -v -o /usr/local/bin/app ./...

EXPOSE 2345
EXPOSE 5001

CMD ["dlv","debug","--listen=:2345","--headless=true","--api-version=2","--accept-multiclient","--log","--log-output=rpc,dap"]
# CMD ["dlv","debug","--api-version=2"]
# ENTRYPOINT ["bash"]
