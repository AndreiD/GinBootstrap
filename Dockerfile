FROM golang:1.11-alpine AS builder

# Download and install the latest release of dep
#ADD https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64 /usr/bin/dep
#RUN chmod +x /usr/bin/dep

# Copy the code from the host and compile it
WORKDIR $GOPATH/src/github.com/AndreiD/GinBootstrap
#COPY Gopkg.toml Gopkg.lock ./
#RUN dep ensure --vendor-only
RUN go get -d -v ./...
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix netgo -o backend .

FROM scratch
COPY --from=builder /backend ./
ENTRYPOINT ["./backend"]
