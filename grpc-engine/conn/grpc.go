package conn

import (
	"bytes"
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Connect(target string) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(target,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		// connect timeout
		grpc.WithTimeout(60*time.Second),
	)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func Close(c *grpc.ClientConn) {
	// ignore close err
	c.Close()
}

func Call(c *grpc.ClientConn, method string, req []byte) ([]byte, error) {
	out := &bytes.Buffer{}
	err := c.Invoke(context.Background(), method, req, out)
	if err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}
