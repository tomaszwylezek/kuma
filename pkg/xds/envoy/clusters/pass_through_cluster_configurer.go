package clusters

import (
	envoy_api_v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	envoy_cluster "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
)

func PassThroughCluster(name string) ClusterBuilderOpt {
	return ClusterBuilderOptFunc(func(config *ClusterBuilderConfig) {
		config.Add(&passThroughClusterConfigurer{
			name: name,
		})
		config.Add(&altStatNameConfigurer{})
		config.Add(&timeoutConfigurer{})
	})
}

type passThroughClusterConfigurer struct {
	name string
}

func (p *passThroughClusterConfigurer) ConfigureV2(c *envoy_api_v2.Cluster) error {
	c.Name = p.name
	c.ClusterDiscoveryType = &envoy_api_v2.Cluster_Type{Type: envoy_api_v2.Cluster_ORIGINAL_DST}
	c.LbPolicy = envoy_api_v2.Cluster_CLUSTER_PROVIDED
	return nil
}

func (p *passThroughClusterConfigurer) ConfigureV3(c *envoy_cluster.Cluster) error {
	c.Name = p.name
	c.ClusterDiscoveryType = &envoy_cluster.Cluster_Type{Type: envoy_cluster.Cluster_ORIGINAL_DST}
	c.LbPolicy = envoy_cluster.Cluster_CLUSTER_PROVIDED
	return nil
}
