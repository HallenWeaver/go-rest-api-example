FROM    golang:latest as builder
RUN     apk --no-cache add ca-certificates git
WORKDIR /build/api
COPY    go.mod ./
RUN     go mod download
COPY    . ./
RUN     CGO_ENABLED=0 go build -o api

# Post-build stage
FROM    alpine
WORKDIR /root
COPY    --from=builder /build/api/api .
EXPOSE  8080
CMD     ["./api"]