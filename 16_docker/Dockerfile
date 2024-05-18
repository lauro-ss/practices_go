## stage 1 - golang build
FROM golang:1.22-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

# copy the application source code
COPY . .

# build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -o myapp cmd/main.go


## stage 2 - runtime
FROM alpine

# copy all the content from /app/myapp into root of alpine
COPY --from=build /app/myapp /

# just to let know
EXPOSE 8080

ENTRYPOINT [ "/myapp" ]

