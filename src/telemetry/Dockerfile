FROM golang:alpine as build

ENV GOOS=linux \
    GOARCH=amd64

# The default location is /go/src
WORKDIR /go/src/spacef

# We simply copy everything to /go/src/spacef    
COPY . .

# Add the required build libraries
RUN apk update && apk add gcc librdkafka-dev zstd-libs libsasl lz4-dev libc-dev musl-dev 

# Run the build
RUN go build -o main -tags musl

# Run test
RUN go test -v ./... -tags musl

FROM alpine

# We use dockerize to make sure the kafka sever is up and running before the command starts.
ENV DOCKERIZE_VERSION v0.6.1
ENV KAFKA kafka

# Add dockerize
RUN apk --no-cache upgrade && apk --no-cache --virtual .get add curl \
    && curl -L -O https://github.com/jwilder/dockerize/releases/download/${DOCKERIZE_VERSION}/dockerize-linux-amd64-${DOCKERIZE_VERSION}.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && apk del .get \
    # Add the runtime dependency.
    && apk add --no-cache librdkafka

# Fetch the binary 
COPY --from=build /go/src/spacef/main /

# Wait for kafka to come up, only then start /main
ENTRYPOINT ["sh","-c","/usr/local/bin/dockerize -wait tcp://${KAFKA}:9092 /main kafka test"]