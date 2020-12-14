package clusters

import (
	envoy_api_v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	"github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	envoy_api "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"

	envoy_endpoints "github.com/kumahq/kuma/pkg/xds/envoy/endpoints"
)

func DNSCluster(name string, address string, port uint32) ClusterBuilderOpt {
	return ClusterBuilderOptFunc(func(config *ClusterBuilderConfig) {
		config.Add(&dnsClusterConfigurer{
			name:    name,
			address: address,
			port:    port,
		})
		config.Add(&altStatNameConfigurer{})
		config.Add(&timeoutConfigurer{})
	})
}

type dnsClusterConfigurer struct {
	name    string
	address string
	port    uint32
}

func (e *dnsClusterConfigurer) ConfigureV2(c *envoy_api_v2.Cluster) error {
	c.Name = e.name
	c.ClusterDiscoveryType = &envoy_api_v2.Cluster_Type{Type: envoy_api_v2.Cluster_STRICT_DNS}
	c.LbPolicy = envoy_api_v2.Cluster_ROUND_ROBIN
	c.LoadAssignment = envoy_endpoints.CreateStaticEndpoint(e.name, e.address, e.port)
	return nil
}

func (e *dnsClusterConfigurer) ConfigureV3(c *envoy_config_cluster_v3.Cluster) error {
	c.Name = e.name
	c.ClusterDiscoveryType = &envoy_api.Cluster_Type{Type: envoy_api.Cluster_STRICT_DNS}
	c.LbPolicy = envoy_api.Cluster_ROUND_ROBIN
	//c.LoadAssignment = envoy_endpoints.CreateStaticEndpoint(e.name, e.address, e.port) // fixme
	return nil
}
