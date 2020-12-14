package clusters

import (
	"time"

	envoy_api_v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	envoy_cluster "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	"github.com/golang/protobuf/ptypes"
)

const defaultConnectTimeout = 5 * time.Second

func ConnectTimeout(timeout time.Duration) ClusterBuilderOpt {
	return ClusterBuilderOptFunc(func(config *ClusterBuilderConfig) {
		config.Add(&timeoutConfigurer{
			connectTimeout: timeout,
		})
	})
}

type timeoutConfigurer struct {
	connectTimeout time.Duration
}

func (t *timeoutConfigurer) ConfigureV2(cluster *envoy_api_v2.Cluster) error {
	if t.connectTimeout.Nanoseconds() == 0 {
		t.connectTimeout = defaultConnectTimeout
	}
	cluster.ConnectTimeout = ptypes.DurationProto(t.connectTimeout)
	return nil
}

func (t *timeoutConfigurer) ConfigureV3(cluster *envoy_cluster.Cluster) error {
	if t.connectTimeout.Nanoseconds() == 0 {
		t.connectTimeout = defaultConnectTimeout
	}
	cluster.ConnectTimeout = ptypes.DurationProto(t.connectTimeout)
	return nil
}
