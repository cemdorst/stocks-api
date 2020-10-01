FROM golang:1.14 as build

WORKDIR /go/src/app
COPY . .

RUN go get -d -v github.com/gorilla/mux github.com/cemdorst/stocks-api github.com/montanaflynn/stats github.com/gorilla/handlers

RUN CGO_ENABLED=0 go build -o /bin/stocks-api

FROM scratch
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /go/src/app/config.toml /config.toml
COPY --from=build /bin/stocks-api /bin/stocks-api
ENTRYPOINT ["/bin/stocks-api"]
