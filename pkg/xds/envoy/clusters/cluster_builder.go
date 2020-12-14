package clusters

import (
	"errors"

	envoy_api_v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	envoy_api "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"

	"github.com/kumahq/kuma/pkg/xds/envoy"
)

// ClusterConfigurer is responsible for configuring a single aspect of the entire Envoy cluster,
// such as filter chain, transparent proxying, etc.
type ClusterConfigurer interface {
	// Configure configures a single aspect on a given Envoy cluster.
	ConfigureV2(cluster *envoy_api_v2.Cluster) error
	ConfigureV3(cluster *envoy_api.Cluster) error
}

// ClusterBuilderOpt is a configuration option for ClusterBuilder.
//
// The goal of ClusterBuilderOpt is to facilitate fluent ClusterBuilder API.
type ClusterBuilderOpt interface {
	// ApplyTo adds ClusterConfigurer(s) to the ClusterBuilder.
	ApplyTo(config *ClusterBuilderConfig)
}

func NewClusterBuilder(apiVersion envoy.APIVersion) *ClusterBuilder {
	return &ClusterBuilder{
		apiVersion: apiVersion,
	}
}

// ClusterBuilder is responsible for generating an Envoy cluster
// by applying a series of ClusterConfigurers.
type ClusterBuilder struct {
	apiVersion envoy.APIVersion
	config     ClusterBuilderConfig
}

// Configure configures ClusterBuilder by adding individual ClusterConfigurers.
func (b *ClusterBuilder) Configure(opts ...ClusterBuilderOpt) *ClusterBuilder {
	for _, opt := range opts {
		opt.ApplyTo(&b.config)
	}
	return b
}

// Build generates an Envoy cluster by applying a series of ClusterConfigurers.
func (b *ClusterBuilder) Build() (envoy.NamedResource, error) {
	switch b.apiVersion {
	case envoy.APIV2:
		cluster := envoy_api_v2.Cluster{}
		for _, configurer := range b.config.Configurers {
			if err := configurer.ConfigureV2(&cluster); err != nil {
				return nil, err
			}
		}
		return &cluster, nil
	case envoy.APIV3:
		cluster := envoy_api.Cluster{}
		for _, configurer := range b.config.Configurers {
			if err := configurer.ConfigureV3(&cluster); err != nil {
				return nil, err
			}
		}
		return &cluster, nil
	default:
		return nil, errors.New("unknown API")
	}
}

// ClusterBuilderConfig holds configuration of a ClusterBuilder.
type ClusterBuilderConfig struct {
	// A series of ClusterConfigurers to apply to Envoy cluster.
	Configurers []ClusterConfigurer
}

// Add appends a given ClusterConfigurer to the end of the chain.
func (c *ClusterBuilderConfig) Add(configurer ClusterConfigurer) {
	c.Configurers = append(c.Configurers, configurer)
}

// ClusterBuilderOptFunc is a convenience type adapter.
type ClusterBuilderOptFunc func(config *ClusterBuilderConfig)

func (f ClusterBuilderOptFunc) ApplyTo(config *ClusterBuilderConfig) {
	f(config)
}
