FROM golang:latest as builder

WORKDIR /src

COPY . .
RUN CGO_ENABLED=0 go build

FROM scratch
COPY --from=builder /src/goweb /goweb

ENTRYPOINT [ "/goweb" ]
