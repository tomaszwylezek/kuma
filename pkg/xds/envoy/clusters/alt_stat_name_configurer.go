package clusters

import (
	envoy_api_v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	envoy_cluster "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"

	util_xds "github.com/kumahq/kuma/pkg/util/xds"
)

func AltStatName() ClusterBuilderOpt {
	return ClusterBuilderOptFunc(func(config *ClusterBuilderConfig) {
		config.Add(&altStatNameConfigurer{})
	})
}

type altStatNameConfigurer struct {
}

func (e *altStatNameConfigurer) ConfigureV2(cluster *envoy_api_v2.Cluster) error {
	sanitizedName := util_xds.SanitizeMetric(cluster.Name)
	if sanitizedName != cluster.Name {
		cluster.AltStatName = sanitizedName
	}
	return nil
}

func (e *altStatNameConfigurer) ConfigureV3(cluster *envoy_cluster.Cluster) error {
	sanitizedName := util_xds.SanitizeMetric(cluster.Name)
	if sanitizedName != cluster.Name {
		cluster.AltStatName = sanitizedName
	}
	return nil
}
