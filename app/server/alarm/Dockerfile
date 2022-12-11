FROM gghtrt520/golang:alpine AS builder


WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
COPY ./app/server/alarm/etc /app/etc
RUN cd ./app/server/alarm
RUN go build -ldflags="-s -w" -o /app/alarm .


FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/alarm /app/alarm
COPY --from=builder /app/etc /app/etc

CMD ["./alarm", "-f", "etc/alarm.yaml"]
