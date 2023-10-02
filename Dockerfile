# syntax=docker/dockerfile:1
FROM golang:1.21 AS base
WORKDIR /src
# TODO change to folder by folder
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/url-shortener ./cmd/url-shortener/main.go

# TODO change to distroless
FROM scratch
COPY --from=base /bin/url-shortener /bin/url-shortener
EXPOSE 8080
CMD ["/bin/url-shortener"]