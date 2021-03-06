package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"

	pb "go.protobuf.foo.alis.exchange/foo/br/resources/books/v1"
)

// The booksClient is defined as a global variable. It is declared once on init and used to call the various methods of the BooksService
var (
	booksClient pb.BooksServiceClient
)

func init() {

	// Pre-declare err to avoid shadowing.
	var err error

	// Declare the server host url and port.
	// This follows the format {{neuronID}}-{{majorVersion}}-{{hash}}-{{region}}.a.run.app:{{port}}
	// Typical predefined values are:
	//  - region: "ew"
	//	- port: "443"
	serverHost := "resources-books-v1-z5x5ywf7za-ew.a.run.app:443"

	// Initialise connection to the books service.
	conn, err := NewConn(context.Background(), serverHost, false)
	if err != nil {
		log.Fatal(err)
	}

	// Initialise the booksClient.
	booksClient = pb.NewBooksServiceClient(conn)
}

type grpcTokenSource struct {
	oauth.TokenSource
}

// Code generated by alis.exchange CLI. DO NOT EDIT.
//
// NewConn creates a new gRPC connection.
// host should be of the form domain:port, e.g., example.com:443
func NewConn(ctx context.Context, host string, insecure bool) (*grpc.ClientConn, error) {
	var opts []grpc.DialOption
	if host != "" {
		opts = append(opts, grpc.WithAuthority(host))
	}

	if insecure {
		opts = append(opts, grpc.WithInsecure())
	} else {
		systemRoots, err := x509.SystemCertPool()
		if err != nil {
			return nil, err
		}
		cred := credentials.NewTLS(&tls.Config{
			RootCAs: systemRoots,
		})
		opts = append(opts, grpc.WithTransportCredentials(cred))
	}

	return grpc.Dial(host, opts...)
}
