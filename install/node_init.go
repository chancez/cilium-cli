// SPDX-License-Identifier: Apache-2.0
// Copyright 2021 Authors of Cilium

package install

import (
	"github.com/cilium/cilium/pkg/versioncheck"
	appsv1 "k8s.io/api/apps/v1"
	"sigs.k8s.io/yaml"

	"github.com/cilium/cilium-cli/internal/k8s"
)

var (
	nodeInitScript = map[k8s.Kind]string{
		k8s.KindEKS: "",
		k8s.KindGKE: "",
	}
)

func (k *K8sInstaller) generateNodeInitDaemonSet(_ k8s.Kind) *appsv1.DaemonSet {
	var (
		dsFileName string
	)

	ciliumVer := k.getCiliumVersion()
	switch {
	case versioncheck.MustCompile(">=1.9.0")(ciliumVer):
		dsFileName = "templates/cilium-nodeinit/daemonset.yaml"
	default:
		panic("unsupported version")
	}

	dsFile := k.manifests[dsFileName]

	var ds appsv1.DaemonSet
	err := yaml.Unmarshal([]byte(dsFile), &ds)
	if err != nil {
		// Developer mistake, this shouldn't happen
		panic(err)
	}
	return &ds
}
