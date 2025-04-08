package namespace_api

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type NamespaceAPI struct {
	client kubernetes.Interface
}

func NewNamespaceAPI(client kubernetes.Interface) *NamespaceAPI {
	return &NamespaceAPI{
		client: client,
	}
}

func (n *NamespaceAPI) GetNamespaceByName(ctx context.Context, name string) (*corev1.Namespace, error) {
	ns, err := n.client.CoreV1().Namespaces().Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get namespace %q: %w", name, err)
	}
	return ns, nil
}

func (n *NamespaceAPI) ListNamespacesByLabel(ctx context.Context, labelSelector string) ([]corev1.Namespace, error) {
	opts := metav1.ListOptions{
		LabelSelector: labelSelector,
	}

	list, err := n.client.CoreV1().Namespaces().List(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to list namespaces by label %q: %w", labelSelector, err)
	}

	return list.Items, nil
}

func (n *NamespaceAPI) ListNamespacesByField(ctx context.Context, fieldSelector string) ([]corev1.Namespace, error) {
	opts := metav1.ListOptions{
		FieldSelector: fieldSelector,
	}

	list, err := n.client.CoreV1().Namespaces().List(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to list namespaces by field %q: %w", fieldSelector, err)
	}

	return list.Items, nil
}
