package api

import (
	"context"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

type DeploymentAPI interface {
	GetDeploymentByName(ctx context.Context, namespace, name string) (*appsv1.Deployment, error)
	ListDeploymentsByLabel(ctx context.Context, namespace string, labelSelector string) ([]appsv1.Deployment, error)
	ListDeploymentsByField(ctx context.Context, namespace string, fieldSelector string) ([]appsv1.Deployment, error)
}

type NamespaceAPI interface {
	GetNamespaceByName(ctx context.Context, name string) (*corev1.Namespace, error)
	ListNamespacesByLabel(ctx context.Context, labelSelector string) ([]corev1.Namespace, error)
	ListNamespacesByField(ctx context.Context, fieldSelector string) ([]corev1.Namespace, error)
}

type ServiceAPI interface {
	GetServiceByName(ctx context.Context, namespace, name string) (*corev1.Service, error)
	ListServicesByLabel(ctx context.Context, namespace string, labelSelector string) ([]corev1.Service, error)
	ListServicesByField(ctx context.Context, namespace string, fieldSelector string) ([]corev1.Service, error)
}

type PodAPI interface {
	GetPodByName(ctx context.Context, namespace, name string) (*corev1.Pod, error)
	ListPodsByLabel(ctx context.Context, namespace string, labelSelector string) ([]corev1.Pod, error)
	ListPodsByField(ctx context.Context, namespace string, fieldSelector string) ([]corev1.Pod, error)
}
