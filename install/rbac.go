// SPDX-License-Identifier: Apache-2.0
// Copyright 2021 Authors of Cilium

package install

import (
	"github.com/cilium/cilium/pkg/versioncheck"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	"sigs.k8s.io/yaml"

	"github.com/cilium/cilium-cli/defaults"
)

func (k *K8sInstaller) NewServiceAccount(name string) *corev1.ServiceAccount {
	var (
		saFileName string
	)

	ciliumVer := k.getCiliumVersion()
	switch {
	case versioncheck.MustCompile(">=1.9.0")(ciliumVer):
		switch name {
		case defaults.AgentServiceAccountName:
			saFileName = "templates/cilium-agent/serviceaccount.yaml"
		case defaults.OperatorServiceAccountName:
			saFileName = "templates/cilium-operator/serviceaccount.yaml"
		}
	default:
		panic("unsupported version")
	}

	saFile := k.manifests[saFileName]

	var sa corev1.ServiceAccount
	err := yaml.Unmarshal([]byte(saFile), &sa)
	if err != nil {
		// Developer mistake, this shouldn't happen
		panic(err)
	}
	return &sa
}

func (k *K8sInstaller) NewClusterRole(name string) *rbacv1.ClusterRole {
	var (
		crFileName string
	)

	ciliumVer := k.getCiliumVersion()
	switch {
	case versioncheck.MustCompile(">=1.9.0")(ciliumVer):
		switch name {
		case defaults.AgentServiceAccountName:
			crFileName = "templates/cilium-agent/clusterrole.yaml"
		case defaults.OperatorServiceAccountName:
			crFileName = "templates/cilium-operator/clusterrole.yaml"
		}
	default:
		panic("unsupported version")
	}

	crFile := k.manifests[crFileName]

	var cr rbacv1.ClusterRole
	err := yaml.Unmarshal([]byte(crFile), &cr)
	if err != nil {
		// Developer mistake, this shouldn't happen
		panic(err)
	}
	return &cr
}

func (k *K8sInstaller) NewClusterRoleBinding(crbName string) *rbacv1.ClusterRoleBinding {
	var (
		crbFileName string
	)

	ciliumVer := k.getCiliumVersion()
	switch {
	case versioncheck.MustCompile(">=1.9.0")(ciliumVer):
		switch crbName {
		case defaults.AgentClusterRoleName:
			crbFileName = "templates/cilium-agent/clusterrolebinding.yaml"
		case defaults.OperatorClusterRoleName:
			crbFileName = "templates/cilium-operator/clusterrolebinding.yaml"
		}
	default:
		panic("unsupported version")
	}

	crbFile := k.manifests[crbFileName]

	var crb rbacv1.ClusterRoleBinding
	err := yaml.Unmarshal([]byte(crbFile), &crb)
	if err != nil {
		// Developer mistake, this shouldn't happen
		panic(err)
	}
	return &crb
}
