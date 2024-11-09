FROM alpine:3.20.0 AS certs

RUN apk add ca-certificates

FROM ubuntu

WORKDIR /app

COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY ./snippy /usr/bin/snippy

ENV DATABASE_CONNECTION="file:/app/snippy.db"
ENV APPLICATION_COMMAND="serve"
ENV PORT="8080"
ENV HOST="0.0.0.0"

EXPOSE 8080

ENTRYPOINT ["/usr/bin/snippy"]
