package clusters

import (
	envoy_api_v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	envoy_core_v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	envoy_cluster "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	envoy_core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	pstruct "github.com/golang/protobuf/ptypes/struct"

	"github.com/kumahq/kuma/pkg/xds/envoy/tls"

	"github.com/kumahq/kuma/pkg/core/xds"
	"github.com/kumahq/kuma/pkg/util/proto"
	"github.com/kumahq/kuma/pkg/xds/envoy"
)

func ClientSideTLS(endpoints []xds.Endpoint) ClusterBuilderOpt {
	return ClusterBuilderOptFunc(func(config *ClusterBuilderConfig) {
		config.Add(&clientSideTLSConfigurer{
			endpoints: endpoints,
		})
	})
}

type clientSideTLSConfigurer struct {
	endpoints []xds.Endpoint
}

func (c *clientSideTLSConfigurer) ConfigureV2(cluster *envoy_api_v2.Cluster) error {
	for _, ep := range c.endpoints {
		if ep.ExternalService.TLSEnabled {
			tlsContext, err := tls.UpstreamTlsContextOutsideMesh(
				ep.ExternalService.CaCert,
				ep.ExternalService.ClientCert,
				ep.ExternalService.ClientKey,
				ep.Target)
			if err != nil {
				return err
			}

			pbst, err := proto.MarshalAnyDeterministic(tlsContext)
			if err != nil {
				return err
			}

			transportSocket := &envoy_core_v2.TransportSocket{
				Name: "envoy.transport_sockets.tls",
				ConfigType: &envoy_core_v2.TransportSocket_TypedConfig{
					TypedConfig: pbst,
				},
			}

			cluster.TransportSocketMatches = append(cluster.TransportSocketMatches, &envoy_api_v2.Cluster_TransportSocketMatch{
				Name: ep.Target,
				Match: &pstruct.Struct{
					Fields: envoy.MetadataFields(ep.Tags),
				},
				TransportSocket: transportSocket,
			})
		}
	}

	return nil
}

func (c *clientSideTLSConfigurer) ConfigureV3(cluster *envoy_cluster.Cluster) error {
	for _, ep := range c.endpoints {
		if ep.ExternalService.TLSEnabled {
			tlsContext, err := tls.UpstreamTlsContextOutsideMesh(
				ep.ExternalService.CaCert,
				ep.ExternalService.ClientCert,
				ep.ExternalService.ClientKey,
				ep.Target)
			if err != nil {
				return err
			}

			pbst, err := proto.MarshalAnyDeterministic(tlsContext)
			if err != nil {
				return err
			}

			transportSocket := &envoy_core.TransportSocket{
				Name: "envoy.transport_sockets.tls",
				ConfigType: &envoy_core.TransportSocket_TypedConfig{
					TypedConfig: pbst,
				},
			}

			cluster.TransportSocketMatches = append(cluster.TransportSocketMatches, &envoy_cluster.Cluster_TransportSocketMatch{
				Name: ep.Target,
				Match: &pstruct.Struct{
					Fields: envoy.MetadataFields(ep.Tags),
				},
				TransportSocket: transportSocket,
			})
		}
	}

	return nil
}
