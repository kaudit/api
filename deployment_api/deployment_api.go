package deployment_api

import (
	"context"

	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/client-go/kubernetes"
)

type DeploymentAPI struct {
	client kubernetes.Interface
}

func NewDeploymentAPI(client kubernetes.Interface) *DeploymentAPI {
	return &DeploymentAPI{
		client: client,
	}
}

func (d DeploymentAPI) GetDeploymentByName(ctx context.Context, namespace, name string) (*appsv1.Deployment, error) {
	// TODO implement me
	panic("implement me")
}

func (d DeploymentAPI) ListDeploymentsByLabel(ctx context.Context, namespace string, labelSelector string) ([]appsv1.Deployment, error) {
	// TODO implement me
	panic("implement me")
}

func (d DeploymentAPI) ListDeploymentsByField(ctx context.Context, namespace string, fieldSelector string) ([]appsv1.Deployment, error) {
	// TODO implement me
	panic("implement me")
}
