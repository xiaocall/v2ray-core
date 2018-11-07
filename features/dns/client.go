package dns

import (
	"v2ray.com/core/common/net"
	"v2ray.com/core/features"
)

// Client is a V2Ray feature for querying DNS information.
type Client interface {
	features.Feature
	LookupIP(host string) ([]net.IP, error)
}

// ClientType returns the type of Client interface. Can be used for implementing common.HasType.
func ClientType() interface{} {
	return (*Client)(nil)
}

// LocalClient is an implementation of Client, which queries localhost for DNS.
type LocalClient struct{}

// Type implements common.HasType.
func (LocalClient) Type() interface{} {
	return ClientType()
}

// Start implements common.Runnable.
func (LocalClient) Start() error { return nil }

// Close implements common.Closable.
func (LocalClient) Close() error { return nil }

// LookupIP implements Client.
func (LocalClient) LookupIP(host string) ([]net.IP, error) {
	return net.LookupIP(host)
}
