package clusters

import (
	envoy_api_v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	envoy_core_v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"

	envoy_cluster "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	envoy_core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
)

func UpstreamBindConfig(address string, port uint32) ClusterBuilderOpt {
	return ClusterBuilderOptFunc(func(config *ClusterBuilderConfig) {
		config.Add(&upstreamBindConfigConfigurer{
			address: address,
			port:    port,
		})
	})
}

type upstreamBindConfigConfigurer struct {
	address string
	port    uint32
}

func (u *upstreamBindConfigConfigurer) ConfigureV2(c *envoy_api_v2.Cluster) error {
	c.UpstreamBindConfig = &envoy_core_v2.BindConfig{
		SourceAddress: &envoy_core_v2.SocketAddress{
			Address: u.address,
			PortSpecifier: &envoy_core_v2.SocketAddress_PortValue{
				PortValue: u.port,
			},
		},
	}
	return nil
}

func (u *upstreamBindConfigConfigurer) ConfigureV3(c *envoy_cluster.Cluster) error {
	c.UpstreamBindConfig = &envoy_core.BindConfig{
		SourceAddress: &envoy_core.SocketAddress{
			Address: u.address,
			PortSpecifier: &envoy_core.SocketAddress_PortValue{
				PortValue: u.port,
			},
		},
	}
	return nil
}
