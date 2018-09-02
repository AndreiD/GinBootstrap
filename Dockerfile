FROM golang:1.11 AS builder

# Download and install the latest release of dep
#ADD https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64 /usr/bin/dep
#RUN chmod +x /usr/bin/dep

# Copy the code from the host and compile it
WORKDIR $GOPATH/src/github.com/AndreiD/GinBootstrap
#COPY Gopkg.toml Gopkg.lock ./
#RUN dep ensure --vendor-only

COPY . ./
RUN go get -d ./...
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w' -o /go/bin/backend

# In scratch, there is nothing, except the project binary
FROM scratch
LABEL maintainer="Andy <andy@motionwerk.com>"
COPY --from=builder /go/bin/backend /bin/backend
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
EXPOSE 8080
ENTRYPOINT [ "/bin/backend" ]
CMD [ "--produce" ]
