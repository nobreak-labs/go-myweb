FROM golang:alpine AS gobuilder

# Set ARG for architecture
ARG TARGETARCH

# Set Golang environmet variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=$TARGETARCH

# Move to working directory /source
WORKDIR /source

# Copy the code into the container
COPY . .

# Get Go Module
RUN go get -u -d -v .

# Test the application
RUN go test -v

# Build the application
RUN go build -ldflags '-s -w' -o myweb ./myweb.go

# Move to /release directory for builded binary
WORKDIR /release

# Copy binary from source
RUN cp /source/myweb .

EXPOSE 8080

# Build a scratch image
FROM scratch AS release
COPY --from=gobuilder /release/myweb /
ENTRYPOINT ["/myweb"]
