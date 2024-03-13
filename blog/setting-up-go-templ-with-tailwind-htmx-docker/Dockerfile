FROM golang:1.21-alpine as build

WORKDIR /app
COPY . .

RUN apk add make npm nodejs &&\
    make

FROM alpine:latest as run

WORKDIR /app
COPY --from=build /app/spendings ./run
COPY --from=build /app/static ./static

EXPOSE 8080

CMD ["./run"]
