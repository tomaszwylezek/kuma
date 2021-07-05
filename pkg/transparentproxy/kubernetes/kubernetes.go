/*
Copyright 2021 Kuma authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package kubernetes

import (
	"fmt"
	"strconv"

	kube_core "k8s.io/api/core/v1"

	"github.com/kumahq/kuma/pkg/plugins/runtime/k8s/metadata"
	"github.com/kumahq/kuma/pkg/transparentproxy/config"
)

type PodRedirect struct {
	BuiltinDNSEnabled      bool
	BuiltinDNSPort         uint32
	ExcludeOutboundPorts   string
	RedirectPortOutbound   uint32
	RedirectInbound        bool
	ExcludeInboundPorts    string
	RedirectPortInbound    uint32
	RedirectPortInboundV6  uint32
	UID                    string
	RedirectAllDNSTraffic  bool
	DNSUpstreamTargetChain string
}

func NewPodRedirectForPod(pod *kube_core.Pod) (*PodRedirect, error) {
	var err error
	podRedirect := &PodRedirect{}

	podRedirect.BuiltinDNSEnabled, _, err = metadata.Annotations(pod.Annotations).GetEnabled(metadata.KumaBuiltinDNS)
	if err != nil {
		return nil, err
	}

	if podRedirect.BuiltinDNSEnabled {
		// default to redirect all DNS traffic, as this is what we need in Kubernetes.
		// if this code is used in Universal, override this in the caller.
		podRedirect.RedirectAllDNSTraffic = true
	}

	podRedirect.BuiltinDNSPort, _, err = metadata.Annotations(pod.Annotations).GetUint32(metadata.KumaBuiltinDNSPort)
	if err != nil {
		return nil, err
	}

	podRedirect.ExcludeOutboundPorts, _ = metadata.Annotations(pod.Annotations).GetString(metadata.KumaTrafficExcludeOutboundPorts)

	podRedirect.RedirectPortOutbound, _, err = metadata.Annotations(pod.Annotations).GetUint32(metadata.KumaTransparentProxyingOutboundPortAnnotation)
	if err != nil {
		return nil, err
	}

	podRedirect.RedirectInbound = true
	enabled, exist, err := metadata.Annotations(pod.Annotations).GetEnabled(metadata.KumaGatewayAnnotation)
	if err != nil {
		return nil, err
	}
	if exist && enabled {
		podRedirect.RedirectInbound = false
	}

	podRedirect.ExcludeInboundPorts, _ = metadata.Annotations(pod.Annotations).GetString(metadata.KumaTrafficExcludeInboundPorts)

	podRedirect.RedirectPortInbound, _, err = metadata.Annotations(pod.Annotations).GetUint32(metadata.KumaTransparentProxyingInboundPortAnnotation)
	if err != nil {
		return nil, err
	}

	podRedirect.RedirectPortInboundV6, _, err = metadata.Annotations(pod.Annotations).GetUint32(metadata.KumaTransparentProxyingInboundPortAnnotationV6)
	if err != nil {
		return nil, err
	}

	podRedirect.UID, _ = metadata.Annotations(pod.Annotations).GetString(metadata.KumaSidecarUID)

	return podRedirect, nil
}

func (pr *PodRedirect) AsTransparentProxyConfig() *config.TransparentProxyConfig {
	dnsUpstreamTargetChain := pr.DNSUpstreamTargetChain
	if dnsUpstreamTargetChain == "" {
		dnsUpstreamTargetChain = "RETURN"
	}

	return &config.TransparentProxyConfig{
		DryRun:                 false,
		Verbose:                false,
		RedirectPortOutBound:   fmt.Sprintf("%d", pr.RedirectPortOutbound),
		RedirectInBound:        pr.RedirectInbound,
		RedirectPortInBound:    fmt.Sprintf("%d", pr.RedirectPortInbound),
		RedirectPortInBoundV6:  fmt.Sprintf("%d", pr.RedirectPortInboundV6),
		ExcludeInboundPorts:    pr.ExcludeInboundPorts,
		ExcludeOutboundPorts:   pr.ExcludeOutboundPorts,
		UID:                    pr.UID,
		GID:                    pr.UID, // TODO: shall we have a separate annotation here?
		RedirectDNS:            pr.BuiltinDNSEnabled,
		RedirectAllDNSTraffic:  pr.RedirectAllDNSTraffic,
		AgentDNSListenerPort:   strconv.FormatInt(int64(pr.BuiltinDNSPort), 10),
		DNSUpstreamTargetChain: dnsUpstreamTargetChain,
	}
}

func (pr *PodRedirect) AsKumactlCommandLine() []string {
	result := []string{
		"--redirect-outbound-port",
		fmt.Sprintf("%d", pr.RedirectPortOutbound),
		"--redirect-inbound=" + fmt.Sprintf("%t", pr.RedirectInbound),
		"--redirect-inbound-port",
		fmt.Sprintf("%d", pr.RedirectPortInbound),
		"--redirect-inbound-port-v6",
		fmt.Sprintf("%d", pr.RedirectPortInboundV6),
		"--kuma-dp-uid",
		pr.UID,
		"--exclude-inbound-ports",
		pr.ExcludeInboundPorts,
		"--exclude-outbound-ports",
		pr.ExcludeOutboundPorts,
		"--verbose",
		"--skip-resolv-conf",
	}

	if pr.BuiltinDNSEnabled {
		if pr.RedirectAllDNSTraffic {
			result = append(result,
				"--redirect-all-dns-traffic",
			)
		}
		result = append(result,
			"--redirect-dns-port", strconv.FormatInt(int64(pr.BuiltinDNSPort), 10),
		)
	}

	if pr.DNSUpstreamTargetChain != "" {
		result = append(result,
			"--redirect-dns-upstream-target-chain", pr.DNSUpstreamTargetChain,
		)
	}

	return result
}
