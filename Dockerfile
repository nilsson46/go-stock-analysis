FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/mypackage/myapp/
COPY . .

RUN go get -d -v
RUN go build -o /app/cmd/stock-analysis

FROM alpine:latest
COPY --from=builder /app/cmd/stock-analysis /stock-analysis
COPY wait-for-it.sh /wait-for-it.sh
COPY *.yml ./

RUN chmod +x /wait-for-it.sh
RUN ls -l /wait-for-it.sh  # Lägg till denna rad för att verifiera att filen kopieras korrekt

EXPOSE 8085
ENTRYPOINT ["/wait-for-it.sh", "db", "5432", "--", "/stock-analysis"]