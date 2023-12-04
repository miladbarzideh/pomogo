# Build the application from source
FROM golang:1.21-alpine as build-stage

LABEL maintainer="Milad Barzideh <barzidehmilad@gmail.com>"

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o pomogo ./cmd/pomogo

# Deploy the application binary into a lean image
FROM gcr.io/distroless/static-debian12 as build-release-stage

WORKDIR /app

COPY --from=build-stage /app/pomogo .
COPY --from=build-stage /app/configs ./configs

ENTRYPOINT ["./pomogo"]