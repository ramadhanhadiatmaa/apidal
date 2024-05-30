FROM golang:1.22.2 AS build

WORKDIR /app

COPY go.mod go.sum ./
COPY . .

RUN go mod download

RUN CGO_ENABLED=0 go build -o apidal

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/apidal .

EXPOSE 8062

CMD [ "./apidal" ]