FROM golang:1.20-alpine as golang

WORKDIR /app
COPY . .

RUN go mod download
RUN go mod verify

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /server ./cmd/pickup-locations-server/main.go

FROM gcr.io/distroless/static-debian11

COPY --from=golang /server .

CMD ["/server"]
