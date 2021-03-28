#builder
FROM golang:alpine as builder
WORKDIR /home
COPY . .
RUN go build -o ecommerce-app app/main.go

#final image
FROM alpine
RUN apk update && \
    apk add --no-cache tzdata && \
    apk add --no-cache curl
ENV TZ=Asia/Jakarta
RUN rm -rf /var/cache/apk/* && date
COPY --from=builder /home/ecommerce-app .
EXPOSE 5050
CMD ./ecommerce-app