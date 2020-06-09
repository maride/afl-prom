FROM golang:1.14
COPY . .
RUN go get github.com/prometheus/client_golang/prometheus
RUN go build -o afl-prom .
CMD [ "/afl-prom" ]
