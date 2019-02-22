FROM golang:1.11-alpine as build
RUN apk --no-cache add git
WORKDIR /build/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o api-server .

FROM alpine:3.9
COPY --from=build /build/api-server /
CMD ["/api-server"]
