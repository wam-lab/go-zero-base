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
COPY jobs/email/etc /app/etc
RUN go build -ldflags="-s -w" -o /app/emailJob jobs/email/emailJob.go


FROM alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/emailJob /app/emailJob
COPY --from=builder /app/etc /app/etc

EXPOSE 8889

CMD ["./emailJob", "-f", "etc/emailJob.yaml"]
