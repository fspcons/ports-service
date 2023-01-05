FROM golang:1.19-alpine AS build
WORKDIR /var/app

COPY . /var/app
RUN env GOOS=linux go build -ldflags='-s -w' -o /var/app/bin/app /var/app/src/main.go


## Runtime image
FROM alpine:latest AS runtime

COPY --from=build /var/app/bin/app /usr/local/bin/app
COPY --from=build /var/app/data/ports.json /usr/local/bin/ports.json

ENV PORTS_FILE_PATH=/usr/local/bin/ports.json

EXPOSE 8080

ENTRYPOINT [ "/usr/local/bin/app" ]