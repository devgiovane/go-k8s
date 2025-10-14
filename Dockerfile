FROM golang:1.24.2
COPY . .
RUN go build -o server .
CMD [ "./server" ]