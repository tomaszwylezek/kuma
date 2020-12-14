package xds

import (
	"context"

	"github.com/golang/protobuf/proto"
)

type DiscoveryRequest interface {
	NodeId() string
	GetTypeUrl() string
	GetResponseNonce() string
	HasErrors() bool
	Proto() proto.Message
}

type DiscoveryResponse interface {

}

type Callbacks interface {
	// OnStreamOpen is called once an xDS stream is open with a stream ID and the type URL (or "" for ADS).
	// Returning an error will end processing and close the stream. OnStreamClosed will still be called.
	OnStreamOpen(context.Context, int64, string) error
	// OnStreamClosed is called immediately prior to closing an xDS stream with a stream ID.
	OnStreamClosed(int64)
	// OnStreamRequest is called once a request is received on a stream.
	// Returning an error will end processing and close the stream. OnStreamClosed will still be called.
	OnStreamRequest(int64, DiscoveryRequest) error
	// OnStreamResponse is called immediately prior to sending a response on a stream.
	OnStreamResponse(int64, DiscoveryRequest, DiscoveryResponse)
}
