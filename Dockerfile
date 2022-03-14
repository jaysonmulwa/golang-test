# FROM golang:1.16 as builder

# RUN mkdir /app
# ADD . /app
# WORKDIR /app

# RUN CGO_ENABLED=1 GOOS=linux go build -o app cmd/main.go

# FROM alpine:latest AS production
# COPY --from=builder /app .
# CMD ["./app"]


FROM golang:1.16 as builder
RUN mkdir /build
ADD . /build/
WORKDIR /build

RUN go build -o app cmd/main.go

FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build /app/

WORKDIR /app
CMD ["./app"]