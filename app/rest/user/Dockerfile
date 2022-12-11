FROM gghtrt520/golang:alpine AS builder

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
COPY ./app/rest/user/etc /app/etc
RUN go build -ldflags="-s -w" -o /app/rest/user/user.go .


FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/user /app/user
COPY --from=builder /app/etc /app/etc

CMD ["./user", "-f", "etc/user.yaml"]
