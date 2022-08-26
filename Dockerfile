FROM golang:1.18.4-alpine
WORKDIR /
ADD . / 
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM scratch
COPY --from=0 main /
COPY --from=0 .env /
CMD ["/main"]
