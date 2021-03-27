#builder
FROM golang:alpine as builder
WORKDIR /home
COPY . .
RUN go build -o ecommerce-app app/main.go

#final image
FROM alpine
RUN apk add tzdata
COPY --from=builder /home/ecommerce-app .
EXPOSE 5050
CMD ./ecommerce-app