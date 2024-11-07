FROM golang:1.23.1 AS builder

WORKDIR /auth_service

ENV DB_HOST=localhost
ENV DB_PORT=5432
ENV DB_USER=postgres
ENV DB_PASS=postgres
ENV DB_NAME=auth_service
ENV JWT_KEY=1292505d912b1eac350324e78290b150a30ce3d7d67ded121f4ccda8cb1bb67472be425d926ca3a1ee7232fb5e89b763c781ef15acc430cc2aaf4c5cb96cafe3

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go mod tidy
RUN go build -o auth_service cmd/app/main.go

FROM golang:1.23.1

WORKDIR /app

COPY --from=builder /app/auth_service .

EXPOSE 8000

CMD [ "./auth_service" ]
