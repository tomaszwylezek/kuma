package clusters

import (
	envoy_api_v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	envoy_core_v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	envoy_cluster "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	envoy_core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
)

func EdsCluster(name string) ClusterBuilderOpt {
	return ClusterBuilderOptFunc(func(config *ClusterBuilderConfig) {
		config.Add(&edsClusterConfigurer{
			name: name,
		})
		config.Add(&altStatNameConfigurer{})
		config.Add(&timeoutConfigurer{})
	})
}

type edsClusterConfigurer struct {
	name string
}

func (e *edsClusterConfigurer) ConfigureV2(c *envoy_api_v2.Cluster) error {
	c.Name = e.name
	c.ClusterDiscoveryType = &envoy_api_v2.Cluster_Type{Type: envoy_api_v2.Cluster_EDS}
	c.EdsClusterConfig = &envoy_api_v2.Cluster_EdsClusterConfig{
		EdsConfig: &envoy_core_v2.ConfigSource{
			ConfigSourceSpecifier: &envoy_core_v2.ConfigSource_Ads{
				Ads: &envoy_core_v2.AggregatedConfigSource{},
			},
		},
	}
	return nil
}

func (e *edsClusterConfigurer) ConfigureV3(c *envoy_cluster.Cluster) error {
	c.Name = e.name
	c.ClusterDiscoveryType = &envoy_cluster.Cluster_Type{Type: envoy_cluster.Cluster_EDS}
	c.EdsClusterConfig = &envoy_cluster.Cluster_EdsClusterConfig{
		EdsConfig: &envoy_core.ConfigSource{
			ConfigSourceSpecifier: &envoy_core.ConfigSource_Ads{
				Ads: &envoy_core.AggregatedConfigSource{},
			},
		},
	}
	return nil
}
