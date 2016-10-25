// Package userip provides functions for extracting a user IP address from
// a request and associating it with a Context
package userip

import (
	"context"
	"fmt"
	"net"
	"net/http"
)

// FromRequest extracts the use IP address from req, if present
func FromRequest(req *http.Request) (net.IP, error) {
	ip, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		return nil, fmt.Errorf("userip: %q is not IP:port", req.RemoteAddr)
	}

	userIP := net.ParseIP(ip)
	if err != nil {
		return nil, fmt.Errorf("userip: %q is not IP:port", req.RemoteAddr)
	}

	return userIP, nil
}

// The key type is unexported to prevent collision with context keys
// defined in other package
type key int

// userIPKey is the context key for the user IP address. Its value of zero is
// arbitrary.
const userIPKey key = 0

// NewContext extracts the user IP address from ctx, if present
func NewContext(ctx context.Context, userIP net.IP) context.Context {
	return context.WithValue(ctx, userIPKey, userIP)
}

// FromContext extract the user IP address from ctx, if present
func FromContext(ctx context.Context) (net.IP, bool) {
	// ctx.Value returns nil if ctx has no value for the key;
	// the net.IP type assertion returns of=false for nil
	userIP, ok := ctx.Value(userIPKey).(net.IP)
	return userIP, ok
}
