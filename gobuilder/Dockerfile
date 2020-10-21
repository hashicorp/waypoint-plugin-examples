FROM golang:1.15.3-alpine3.12 as build

# Install the Protocol Buffers compiler and Go plugin
RUN apk add protobuf git make zip
RUN go get github.com/golang/protobuf/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc

# Create the source folder
RUN mkdir /go/plugin
WORKDIR /go/plugin

# Copy the source to the build folder
COPY . /go/plugin

# Build the plugin
RUN chmod +x ./print_arch
RUN make all

# Create the zipped binaries
RUN make zip

FROM scratch as export_stage

COPY --from=build /go/plugin/bin/*.zip .
