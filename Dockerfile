FROM golang:1.21 AS base
RUN apt-get update && \
    apt-get install -y zip

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

FROM base AS build_api

WORKDIR /build

RUN go build -o /tmp/service ./cmd/api/...


FROM golang:1.21 AS app

COPY ["scripts/entrypoint.sh", "/entrypoint.sh"]
COPY ["migration", "/app/migration"]

COPY --from=build_api /tmp/service /app/service

CMD ["/entrypoint.sh", "/app/service"]
