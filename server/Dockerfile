# syntax=docker/dockerfile:1

FROM golang:1.21.1-bullseye AS build

WORKDIR /app

# copy Go modules and dependencies to image
COPY go.mod go.sum ./

# download Go modules and dependencies
RUN go mod download

# copy directory files i.e all files ending with .go
COPY . ./

# compile application
RUN CGO_ENABLED=0 GOOS=linux go build -o /server ./cmd/main.go 

FROM scratch

WORKDIR /

COPY --from=build /server /server

EXPOSE 50051

CMD [ "/server" ]
