package clusters

import (
	envoy_api_v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	envoy_cluster "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"

	"github.com/kumahq/kuma/pkg/core/xds"

	envoy_endpoints "github.com/kumahq/kuma/pkg/xds/envoy/endpoints"
)

func StrictDNSCluster(name string, endpoints []xds.Endpoint) ClusterBuilderOpt {
	return ClusterBuilderOptFunc(func(config *ClusterBuilderConfig) {
		config.Add(&StricDNSClusterConfigurer{
			name:      name,
			endpoints: endpoints,
		})
		config.Add(&altStatNameConfigurer{})
		config.Add(&timeoutConfigurer{})
	})
}

type StricDNSClusterConfigurer struct {
	name      string
	endpoints []xds.Endpoint
}

func (e *StricDNSClusterConfigurer) ConfigureV2(c *envoy_api_v2.Cluster) error {
	c.Name = e.name
	c.ClusterDiscoveryType = &envoy_api_v2.Cluster_Type{Type: envoy_api_v2.Cluster_STRICT_DNS}
	c.LbPolicy = envoy_api_v2.Cluster_ROUND_ROBIN
	c.LoadAssignment = envoy_endpoints.CreateClusterLoadAssignment(e.name, e.endpoints)
	return nil
}

func (e *StricDNSClusterConfigurer) ConfigureV3(c *envoy_cluster.Cluster) error {
	c.Name = e.name
	c.ClusterDiscoveryType = &envoy_cluster.Cluster_Type{Type: envoy_cluster.Cluster_STRICT_DNS}
	c.LbPolicy = envoy_cluster.Cluster_ROUND_ROBIN
	//c.LoadAssignment = envoy_endpoints.CreateClusterLoadAssignment(e.name, e.endpoints) fixme
	return nil
}
