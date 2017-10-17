package kubernikus

import (
	"fmt"

	"github.com/sapcc/kubernikus/pkg/apis/kubernikus/v1"
)

const (
	DEFAULT_CLUSTER_CIDR       = "198.19.0.0/16"
	DEFAULT_SERVICE_CIDR       = "198.18.128.0/17"
	DEFAULT_ADVERTISE_ADDRESS  = "198.18.128.1"
	DEFAULT_CLUSTER_DNS        = "198.18.254.254"
	DEFAULT_CLUSTER_DNS_DOMAIN = "cluster.local"
)

type KlusterFactory interface {
	KlusterFor(v1.KlusterSpec) (*v1.Kluster, error)
}

type klusterFactory struct {
}

func NewKlusterFactory() KlusterFactory {
	return &klusterFactory{}
}

func (klusterFactory) KlusterFor(spec v1.KlusterSpec) (*v1.Kluster, error) {
	if spec.Name == "" {
		return nil, fmt.Errorf("unabled to create cluster. missing name")
	}

	k := &v1.Kluster{
		Spec: spec,
		Status: v1.KlusterStatus{
			Kluster: v1.KlusterInfo{
				State: v1.KlusterPending,
			},
			NodePools: []v1.NodePoolInfo{},
		},
	}

	if k.Spec.ClusterCIDR == "" {
		k.Spec.ClusterCIDR = DEFAULT_CLUSTER_CIDR
	}

	if k.Spec.ServiceCIDR == "" {
		k.Spec.ServiceCIDR = DEFAULT_SERVICE_CIDR
	}

	if k.Spec.AdvertiseAddress == "" {
		k.Spec.AdvertiseAddress = DEFAULT_ADVERTISE_ADDRESS
	}

	if k.Spec.ClusterDNS == "" {
		k.Spec.ClusterDNS = DEFAULT_CLUSTER_DNS
	}

	if k.Spec.ClusterDNSDomain == "" {
		k.Spec.ClusterDNSDomain = DEFAULT_CLUSTER_DNS_DOMAIN
	}

	if k.ObjectMeta.Name == "" {
		k.ObjectMeta.Name = spec.Name
	}

	for _, nodePool := range k.Spec.NodePools {
		k.Status.NodePools = append(k.Status.NodePools, v1.NodePoolInfo{
			Name:        nodePool.Name,
			Size:        nodePool.Size,
			Running:     0,
			Healthy:     0,
			Schedulable: 0,
		})
	}

	return k, nil
}
