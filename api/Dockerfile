FROM golang:1.20.1-alpine3.17 as builder

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -o /usr/local/bin/app .

FROM alpine
WORKDIR /usr/src/app
COPY --from=builder /usr/local/bin/app ./goal-tracker
COPY --from=builder /usr/src/app/docs ./docs
RUN chmod +x ./goal-tracker

EXPOSE 8000
CMD ["/usr/src/app/goal-tracker"]
