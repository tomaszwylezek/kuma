package clusters

import (
	envoy_api_v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	envoy_cluster "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"

	envoy_endpoints "github.com/kumahq/kuma/pkg/xds/envoy/endpoints"
)

func StaticCluster(name string, address string, port uint32) ClusterBuilderOpt {
	return ClusterBuilderOptFunc(func(config *ClusterBuilderConfig) {
		config.Add(&staticClusterConfigurer{
			name:    name,
			address: address,
			port:    port,
		})
		config.Add(&altStatNameConfigurer{})
		config.Add(&timeoutConfigurer{})
	})
}

type staticClusterConfigurer struct {
	name    string
	address string
	port    uint32
}

func (e *staticClusterConfigurer) ConfigureV2(c *envoy_api_v2.Cluster) error {
	c.Name = e.name
	c.ClusterDiscoveryType = &envoy_api_v2.Cluster_Type{Type: envoy_api_v2.Cluster_STATIC}
	c.LoadAssignment = envoy_endpoints.CreateStaticEndpoint(e.name, e.address, e.port)
	return nil
}

func (e *staticClusterConfigurer) ConfigureV3(c *envoy_cluster.Cluster) error {
	c.Name = e.name
	c.ClusterDiscoveryType = &envoy_cluster.Cluster_Type{Type: envoy_cluster.Cluster_STATIC}
	//c.LoadAssignment = envoy_endpoints.CreateStaticEndpoint(e.name, e.address, e.port) fixme
	return nil
}
