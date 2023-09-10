FROM golang:1.21.1-alpine3.18 as builder

WORKDIR /app
ADD . /app
RUN go test main_test.go -c -o challenge
RUN chmod +x challenge
ENV LANGUAGE=go
CMD ["/app/challenge"]

FROM php:8.2-cli-alpine3.17 as php
COPY --from=builder /app/challenge /opt/challenge
COPY --from=builder /app/tests /opt/tests
COPY --from=builder /app/run-php.sh /opt/run-php.sh
VOLUME /opt/php
WORKDIR /opt
ENV LANGUAGE=php
CMD ["/opt/challenge"]

FROM golang:1.21.1-alpine3.18 as go
COPY --from=builder /app/challenge /opt/challenge
COPY --from=builder /app/tests /opt/tests
COPY --from=builder /app/run-go.sh /opt/run-go.sh
VOLUME /opt/go
WORKDIR /opt
ENV LANGUAGE=go
CMD ["/opt/challenge"]

FROM node:18.17.1-alpine3.18 as js
COPY --from=builder /app/challenge /opt/challenge
COPY --from=builder /app/tests /opt/tests
COPY --from=builder /app/run-js.sh /opt/run-js.sh
VOLUME /opt/js
WORKDIR /opt
ENV LANGUAGE=js
CMD ["/opt/challenge"]

