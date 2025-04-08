package service_api

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
)

type ServiceAPI struct {
	client kubernetes.Interface
}

func (s ServiceAPI) GetServiceByName(ctx context.Context, namespace, name string) (*corev1.Service, error) {
	// TODO implement me
	panic("implement me")
}

func (s ServiceAPI) ListServicesByLabel(ctx context.Context, namespace string, labelSelector string) ([]corev1.Service, error) {
	// TODO implement me
	panic("implement me")
}

func (s ServiceAPI) ListServicesByField(ctx context.Context, namespace string, fieldSelector string) ([]corev1.Service, error) {
	// TODO implement me
	panic("implement me")
}

func NewServiceAPI(client kubernetes.Interface) *ServiceAPI {
	return &ServiceAPI{
		client: client,
	}
}
