FROM golang:1.19-alpine as build
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
#ENV GOARCH=amd64
ARG PATH_PROJECT

RUN apk add --no-cache git

WORKDIR /services/app
# Copy dependencies
COPY pkg /pkg

COPY $PATH_PROJECT/go.mod $PATH_PROJECT/go.sum ./
RUN go mod download
COPY $PATH_PROJECT .

#RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/app
RUN go build -o /go/bin/app


FROM alpine:3.7

RUN adduser -D -u 1000 app
RUN apk --no-cache add tzdata
RUN apk add --no-cache ca-certificates

COPY --from=build /go/bin/app /go/bin/app

USER app
EXPOSE 8081
ENTRYPOINT ["./go/bin/app"]
