FROM golang:1.22-alpine as build

WORKDIR /app
COPY . .

RUN go install github.com/a-h/templ/cmd/templ@latest &&\
    apk add make npm nodejs &&\
    make

FROM alpine:latest as run

WORKDIR /app
COPY --from=build /app/spendings ./run
COPY --from=build /app/db.json ./db.json

EXPOSE 8080

CMD ["./run"]
