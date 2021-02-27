FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /build/zero

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
COPY jobs/etc /app/etc
RUN go build -ldflags="-s -w" -o /app/jobs jobs/jobs.go


FROM alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/jobs /app/jobs
COPY --from=builder /app/etc /app/etc

EXPOSE 8889

CMD ["./jobs", "-f", "etc/jobs.yaml"]
