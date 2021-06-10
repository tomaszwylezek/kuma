package v3

import (
	envoy_listener "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	envoy_local_ratelimit "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/local_ratelimit/v3"
	envoy_type_v3 "github.com/envoyproxy/go-control-plane/envoy/type/v3"
	"github.com/golang/protobuf/ptypes/duration"
	"github.com/golang/protobuf/ptypes/wrappers"

	"github.com/kumahq/kuma/pkg/util/proto"
)

type LocalRateLimitConfigurer struct {
	StatsName string
}

func (c *LocalRateLimitConfigurer) Configure(filterChain *envoy_listener.FilterChain) error {
	tcpProxy := c.rateLimit()

	pbst, err := proto.MarshalAnyDeterministic(tcpProxy)
	if err != nil {
		return err
	}

	filterChain.Filters = append([]*envoy_listener.Filter{
		{
			Name: "envoy.filters.network.local_ratelimit",
			ConfigType: &envoy_listener.Filter_TypedConfig{
				TypedConfig: pbst,
			},
		},
	}, filterChain.Filters...)

	//
	return nil
}

func (c *LocalRateLimitConfigurer) rateLimit() *envoy_local_ratelimit.LocalRateLimit {
	return &envoy_local_ratelimit.LocalRateLimit{
		StatPrefix: "rate_limit_" + c.StatsName,
		TokenBucket: &envoy_type_v3.TokenBucket{
			MaxTokens: 10,
			TokensPerFill: &wrappers.UInt32Value{
				Value: 10,
			},
			FillInterval: &duration.Duration{
				Seconds: 1,
			},
		},
	}
}
