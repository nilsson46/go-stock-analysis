FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/mypackage/myapp/
COPY . .

RUN go get -d -v

RUN go build -o /app/cmd/stock-analysis
FROM scratch

COPY --from=builder /app/cmd/stock-analysis /stock-analysis
COPY *.yml ./ 


EXPOSE 8080
ENTRYPOINT ["/stock-analysis"]