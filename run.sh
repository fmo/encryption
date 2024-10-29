#!/bin/bash

protoc \
  --go_out=. \
  --go_opt=paths=source_relative \
  --go-grpc_out=. \
  --go-grpc_opt=paths=source_relative \
  order.proto

go mod tidy

mkdir -p cert && cd cert
echo "Generating private key and self-signed certificate for CA..."
openssl req -x509 \
  -sha256 \
  -newkey rsa:4096 \
  -days 365 \
  -keyout ca.key \
  -out ca.crt \
  -subj "/C=TR/ST=EURASIA/L=ISTANBUL/O=Software/OU=Microservices/CN=*.microservices.dev/emailAddress=huseyin@microservices.dev" \
  -nodes
