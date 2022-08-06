
FROM golang:latest as builder


WORKDIR /app

COPY . .

RUN go build -o main -a -ldflags "-linkmode external -extldflags '-static' -s -w" .


FROM alpine:3.15

WORKDIR /app

COPY --from=builder /app/main .
RUN ls
EXPOSE 4000
# executable
# arguments that can be overridden
CMD [ "./main" ]

