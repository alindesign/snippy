FROM node:lts-alpine AS ui-builder

WORKDIR /app

RUN corepack enable

COPY ./package.json ./pnpm-lock.yaml ./
RUN pnpm install

COPY . .
RUN pnpm css

FROM golang:alpine AS app-builder

RUN apk update && apk upgrade
RUN apk add --no-cache sqlite sqlite-dev gcc musl-dev

WORKDIR /app
COPY . .
COPY --from=ui-builder /app/web/assets/main.css /app/web/assets/main.css
RUN CGO_ENABLED=1 go build -ldflags "-s -w" -v -o /usr/bin/snippy .

FROM alpine:latest

WORKDIR /app

COPY --from=app-builder /usr/bin/snippy /usr/bin/snippy

ENV APPLICATION_COMMAND="serve"
ENV PORT="8080"
ENV HOST="0.0.0.0"

EXPOSE 8080

CMD ["snippy"]
