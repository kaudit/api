package pod_api

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
)

type PodAPI struct {
	client kubernetes.Interface
}

func (p PodAPI) GetPodByName(ctx context.Context, namespace, name string) (*corev1.Pod, error) {
	// TODO implement me
	panic("implement me")
}

func (p PodAPI) ListPodsByLabel(ctx context.Context, namespace string, labelSelector string) ([]corev1.Pod, error) {
	// TODO implement me
	panic("implement me")
}

func (p PodAPI) ListPodsByField(ctx context.Context, namespace string, fieldSelector string) ([]corev1.Pod, error) {
	// TODO implement me
	panic("implement me")
}

func NewPodAPI(client kubernetes.Interface) *PodAPI {
	return &PodAPI{
		client: client,
	}
}
