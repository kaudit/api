package k8s_api

import (
	"fmt"

	"github.com/kaudit/auth"

	"github.com/kaudit/api"
	"github.com/kaudit/api/deployment_api"
	"github.com/kaudit/api/namespace_api"
	"github.com/kaudit/api/pod_api"
	"github.com/kaudit/api/service_api"
)

// K8sApi provides a centralized access point to high-level Kubernetes API abstractions.
//
// It encapsulates typed interfaces for interacting with Pods, Services, Deployments,
// and Namespaces — each exposed through domain-specific interface contracts.
//
// All API implementations are stateless, thread-safe, and validated via typed input contracts.
type K8sApi struct {
	pods        api.PodAPI
	services    api.ServiceAPI
	deployments api.DeploymentAPI
	namespaces  api.NamespaceAPI
}

// NewK8sApi initializes a K8sApi facade by constructing all typed clients behind interface boundaries.
//
// This function:
//   - Initializes a client using the provided auth.Authenticator (via NativeAPI()).
//   - Injects the client into each module's constructor (e.g., pod_api.NewPodAPI).
//   - Assembles a fully wired K8sApi instance.
func NewK8sApi(auth auth.Authenticator) (*K8sApi, error) {
	client, err := auth.NativeAPI()
	if err != nil {
		return nil, fmt.Errorf("failed to init k8s client: %w", err)
	}

	return &K8sApi{
		pods:        pod_api.NewPodAPI(client),
		services:    service_api.NewServiceAPI(client),
		deployments: deployment_api.NewDeploymentAPI(client),
		namespaces:  namespace_api.NewNamespaceAPI(client),
	}, nil
}

// GetPodAPI exposes the PodAPI interface, allowing access to pod-specific operations.
func (k *K8sApi) GetPodAPI() api.PodAPI {
	return k.pods
}

// GetServiceAPI exposes the ServiceAPI interface for service-level operations.
func (k *K8sApi) GetServiceAPI() api.ServiceAPI {
	return k.services
}

// GetDeploymentAPI exposes the DeploymentAPI interface for managing deployments.
func (k *K8sApi) GetDeploymentAPI() api.DeploymentAPI {
	return k.deployments
}

// GetNamespaceAPI exposes the NamespaceAPI interface for managing namespaces.
func (k *K8sApi) GetNamespaceAPI() api.NamespaceAPI {
	return k.namespaces
}
