FROM golang:alpine AS builder

RUN mkdir /app 
ADD . /app/
WORKDIR /app 
RUN go build ./cmd/*.go


FROM alpine
RUN mkdir /app
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /app/main /app/
EXPOSE 8080
WORKDIR /app
CMD ["./main"]
