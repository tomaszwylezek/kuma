package clusters

import (
	envoy_api "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	envoy_api_v2_core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	envoy_cluster "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
)

func Http2() ClusterBuilderOpt {
	return ClusterBuilderOptFunc(func(config *ClusterBuilderConfig) {
		config.Add(&http2Configurer{})
	})
}

type http2Configurer struct {
}

func (p *http2Configurer) ConfigureV2(c *envoy_api.Cluster) error {
	c.Http2ProtocolOptions = &envoy_api_v2_core.Http2ProtocolOptions{}
	return nil
}

func (p *http2Configurer) ConfigureV3(c *envoy_cluster.Cluster) error {
	//c.Http2ProtocolOptions = &envoy_api_v2_core.Http2ProtocolOptions{} fixme what is equivalent
	return nil
}
