# Building the binary of the App
FROM golang:1.19 AS build

# `boilerplate` should be replaced with your project name
WORKDIR /go/src/pmgo

# Copy all the Code and stuff to compile everything
COPY . .

# Downloads all the dependencies in advance (could be left out, but it's more clear this way)
RUN go mod download

# Builds the application as a staticly linked one, to allow it to run on alpine
RUN cd cmd/gapi && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o gapi && mv gapi ../../gapi && cd ../..


# Moving the binary to the 'final Image' to make it smaller
FROM alpine:latest as final

WORKDIR /app

# `boilerplate` should be replaced here as well
COPY --from=build /go/src/pmgo/gapi .

# Add packages
RUN apk -U upgrade \
    && apk add --no-cache dumb-init ca-certificates \
    && chmod +x /app/gapi

ADD 	 https://releases.hashicorp.com/consul-template/${CONSUL_TEMPLATE_VERSION}/consul-template_${CONSUL_TEMPLATE_VERSION}_linux_amd64.zip /usr/bin/
RUN 	 unzip /usr/bin/consul-template_${CONSUL_TEMPLATE_VERSION}_linux_amd64.zip && \
    	 mv consul-template /usr/local/bin/consul-template && \
    	 rm -rf /usr/bin/consul-template_${CONSUL_TEMPLATE_VERSION}_linux_amd64.zip

# Exposes port 3000 because our program listens on that port
EXPOSE 9900

ENTRYPOINT ["/usr/bin/dumb-init", "--"]
CMD ["./gapi"]
