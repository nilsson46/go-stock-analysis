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
COPY init.sql /docker-entrypoint-initdb.d/

RUN chmod +x /wait-for-it.sh

EXPOSE 8085
ENTRYPOINT ["/wait-for-it.sh", "db", "5432", "--", "sh", "-c", "psql -U user -d database -f /docker-entrypoint-initdb.d/init.sql && /stock-analysis"]