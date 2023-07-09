FROM golang:1.20 AS build

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 go build -o server

FROM gcr.io/distroless/base-debian10

WORKDIR /app

COPY --from=build /app/ ./

USER nonroot:nonroot

ENTRYPOINT [ "/app/server" ]