FROM golang:1.22.4-alpine AS backend

COPY . /app
WORKDIR /app

RUN apk --no-cache add curl

RUN go mod download && go mod verify
RUN go build -v -o /app/outclimb-event-manager cmd/main.go

CMD ["/app/outclimb-event-manager"]
