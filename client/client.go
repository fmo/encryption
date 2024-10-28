package main

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
)

func getTlsCredentials() (credentials.TransportCredentials, error) {
	clientCert, clientCertErr := tls.LoadX509KeyPair("cert/client.crt", "cert/client.key")
	if clientCertErr != nil {
		return nil, fmt.Errorf("could not load client key pair : %v", clientCertErr)
	}

	certPool := x509.NewCertPool()
	caCert, caCertErr := ioutil.ReadFile("cert/ca.crt")
	if caCertErr != nil {
		return nil, fmt.Errorf("could not read Cert CA : %v", caCertErr)
	}

	if ok := certPool.AppendCertsFromPEM(caCert); !ok {
		return nil, errors.New("failed to append CA cert")
	}

	return credentials.NewTLS(&tls.Config{
		ServerName:   "*.microservices.dev",
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      certPool,
	}), nil
}
