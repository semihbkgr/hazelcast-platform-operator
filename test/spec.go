package test

import (
	. "github.com/onsi/gomega"

	hazelcastv1alpha1 "github.com/hazelcast/hazelcast-platform-operator/api/v1alpha1"
)

type HazelcastSpecValues struct {
	ClusterSize int32
	Repository  string
	Version     string
	LicenseKey  string
}

func HazelcastSpec(values *HazelcastSpecValues, ee bool) hazelcastv1alpha1.HazelcastSpec {
	spec := hazelcastv1alpha1.HazelcastSpec{
		ClusterSize: values.ClusterSize,
		Repository:  values.Repository,
		Version:     values.Version,
	}
	if ee {
		spec.LicenseKeySecret = values.LicenseKey
	}
	return spec
}

func CheckHazelcastCR(hazelcast *hazelcastv1alpha1.Hazelcast, expected *HazelcastSpecValues, ee bool) {
	Expect(hazelcast.Spec.ClusterSize).Should(Equal(expected.ClusterSize))
	Expect(hazelcast.Spec.Repository).Should(Equal(expected.Repository))
	Expect(hazelcast.Spec.Version).Should(Equal(expected.Version))
	if ee {
		Expect(hazelcast.Spec.LicenseKeySecret).Should(Equal(expected.LicenseKey))
	}
}

type MCSpecValues struct {
	Repository string
	Version    string
	LicenseKey string
}

func ManagementCenterSpec(values *MCSpecValues, ee bool) hazelcastv1alpha1.ManagementCenterSpec {
	spec := hazelcastv1alpha1.ManagementCenterSpec{
		Repository: values.Repository,
		Version:    values.Version,
		ExternalConnectivity: hazelcastv1alpha1.ExternalConnectivityConfiguration{
			Type: hazelcastv1alpha1.ExternalConnectivityTypeLoadBalancer,
		},
		HazelcastClusters: []hazelcastv1alpha1.HazelcastClusterConfig{},
		Persistence: hazelcastv1alpha1.PersistenceConfiguration{
			StorageClass: nil,
		},
	}
	if ee {
		spec.LicenseKeySecret = values.LicenseKey
	}
	return spec
}

func CheckManagementCenterCR(mc *hazelcastv1alpha1.ManagementCenter, expected *MCSpecValues, ee bool) {
	Expect(mc.Spec.Repository).Should(Equal(expected.Repository))
	Expect(mc.Spec.Version).Should(Equal(expected.Version))
	if ee {
		Expect(mc.Spec.LicenseKeySecret).Should(Equal(expected.LicenseKey))
	}
}