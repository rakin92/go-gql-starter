FROM golang:1.12 as builder

WORKDIR /usr/app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-gql-starter .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o migrate ./migrations/

# Multi stage

FROM alpine:latest

RUN apk --no-cache add ca-certificates

ENV PATH /root

WORKDIR /root/

COPY --from=builder /usr/app/go-gql-starter .

COPY --from=builder /usr/app/migrate .

EXPOSE 8080

CMD ["./go-gql-starter"] 
