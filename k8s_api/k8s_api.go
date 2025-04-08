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

type K8sApi struct {
	pods        api.PodAPI
	services    api.ServiceAPI
	deployments api.DeploymentAPI
	namespaces  api.NamespaceAPI
}

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

func (k *K8sApi) GetPodAPI() api.PodAPI {
	return k.pods
}

func (k *K8sApi) GetServiceAPI() api.ServiceAPI {
	return k.services
}

func (k *K8sApi) GetDeploymentAPI() api.DeploymentAPI {
	return k.deployments
}

func (k *K8sApi) GetNamespaceAPI() api.NamespaceAPI {
	return k.namespaces
}
