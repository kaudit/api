# K8s API Library

A Go library that provides high-level abstractions for interacting with Kubernetes resources using clean, domain-specific interfaces.

## Overview

This library offers a simplified and consistent way to interact with Kubernetes resources through a facade pattern. It provides strongly-typed interfaces for working with:

- Pods
- Services
- Deployments
- Namespaces

## Key Features

- **Type-Safe Interfaces**: All operations are defined through clear interface contracts.
- **Validation Built-in**: Input validation is integrated into all methods.
- **Error Handling**: Detailed error messages with proper context wrapping.
- **Thread-Safe**: All API implementations are stateless and safe for concurrent use.
- **Simplified API Surface**: Focused on common operations with consistent patterns.

## Installation

Make sure you have Go 1.23 or later installed.

To add this library to your Go project, use the `go get` command:

```bash
# Install dependencies
go get github.com/kaudit/auth
go get github.com/kaudit/api
go get github.com/kaudit/val
go get k8s.io/client-go
go get k8s.io/api
go get k8s.io/apimachinery

# Install the main library
go get github.com/kaudit/k8s_api
```

## Architecture

The library follows a layered architecture:

- **K8sApi**: A facade that provides centralized access to all resource types
- **Resource APIs**: Individual typed interfaces for each Kubernetes resource
- **Authentication**: Pluggable authentication mechanism through interfaces

## Usage

### Initialization

```go
import (
    "github.com/kaudit/k8s_api"
    "github.com/kaudit/auth"
)

// Create an authenticator (implementation depends on your environment)
authenticator := myauth.NewKubeConfigAuth("/path/to/kubeconfig")

// Initialize the K8s API with your authenticator
k8sAPI, err := k8s_api.NewK8sApi(authenticator)
if err != nil {
    // handle error
}
```

### Working with Pods

```go
// Get the Pod API
podAPI := k8sAPI.GetPodAPI()

// Retrieve a specific pod
pod, err := podAPI.GetPodByName(ctx, "default", "my-pod-name")
if err != nil {
    // handle error
}

// List pods by label
pods, err := podAPI.ListPodsByLabel(ctx, "default", "app=myapp")
if err != nil {
    // handle error
}
```

### Working with Deployments

```go
// Get the Deployment API
deploymentAPI := k8sAPI.GetDeploymentAPI()

// Get a specific deployment
deployment, err := deploymentAPI.GetDeploymentByName(ctx, "default", "my-deployment")
if err != nil {
    // handle error
}
```

### Working with Services

```go
// Get the Service API
serviceAPI := k8sAPI.GetServiceAPI()

// List services by field selector
services, err := serviceAPI.ListServicesByField(ctx, "default", "metadata.name=frontend")
if err != nil {
    // handle error
}
```

### Working with Namespaces

```go
// Get the Namespace API
namespaceAPI := k8sAPI.GetNamespaceAPI()

// Get a specific namespace
namespace, err := namespaceAPI.GetNamespaceByName(ctx, "default")
if err != nil {
    // handle error
}

// List namespaces by label
namespaces, err := namespaceAPI.ListNamespacesByLabel(ctx, "environment=production")
if err != nil {
    // handle error
}
```

## API Documentation

### K8sApi

#### `NewK8sApi(auth auth.Authenticator) (*K8sApi, error)`
Initializes a K8sApi facade by constructing all typed clients behind interface boundaries.
- Takes an `auth.Authenticator` to establish the Kubernetes client connection
- Returns a fully wired K8sApi instance or an error if initialization fails

#### `GetPodAPI() api.PodAPI`
Exposes the PodAPI interface, allowing access to pod-specific operations.

#### `GetServiceAPI() api.ServiceAPI`
Exposes the ServiceAPI interface for service-level operations.

#### `GetDeploymentAPI() api.DeploymentAPI`
Exposes the DeploymentAPI interface for managing deployments.

#### `GetNamespaceAPI() api.NamespaceAPI`
Exposes the NamespaceAPI interface for managing namespaces.

### PodAPI

#### `GetPodByName(ctx context.Context, namespace, name string) (*corev1.Pod, error)`
Retrieves a specific Pod by namespace and name.
- `ctx`: Context for cancellation
- `namespace`: Namespace of the pod (must be non-empty)
- `name`: Name of the pod (must be non-empty)
- Returns the matched *corev1.Pod or an error if not found or invalid

#### `ListPodsByLabel(ctx context.Context, namespace string, labelSelector string) ([]corev1.Pod, error)`
Lists pods by namespace and label selector.
- `ctx`: Context for cancellation
- `namespace`: Namespace scope
- `labelSelector`: Kubernetes label selector syntax
- Returns all matching pods or an error

#### `ListPodsByField(ctx context.Context, namespace string, fieldSelector string) ([]corev1.Pod, error)`
Lists pods by namespace and field selector.
- `ctx`: Context for cancellation
- `namespace`: Namespace scope
- `fieldSelector`: Kubernetes field selector syntax
- Returns all matching pods or an error

### ServiceAPI

#### `GetServiceByName(ctx context.Context, namespace, name string) (*corev1.Service, error)`
Retrieves a specific Service by namespace and name.
- `ctx`: Context for cancellation
- `namespace`: Namespace of the service (must be non-empty)
- `name`: Name of the service (must be non-empty)
- Returns the matched *corev1.Service or an error if not found or invalid

#### `ListServicesByLabel(ctx context.Context, namespace string, labelSelector string) ([]corev1.Service, error)`
Lists services by namespace and label selector.
- `ctx`: Context for cancellation
- `namespace`: Namespace scope
- `labelSelector`: Kubernetes label selector syntax
- Returns all matching services or an error

#### `ListServicesByField(ctx context.Context, namespace string, fieldSelector string) ([]corev1.Service, error)`
Lists services by namespace and field selector.
- `ctx`: Context for cancellation
- `namespace`: Namespace scope
- `fieldSelector`: Kubernetes field selector syntax
- Returns all matching services or an error

### DeploymentAPI

#### `GetDeploymentByName(ctx context.Context, namespace, name string) (*appsv1.Deployment, error)`
Retrieves a specific Deployment by namespace and name.
- `ctx`: Context for cancellation
- `namespace`: Namespace of the deployment (must be non-empty)
- `name`: Name of the deployment (must be non-empty)
- Returns the matched *appsv1.Deployment or an error if not found or invalid

#### `ListDeploymentsByLabel(ctx context.Context, namespace string, labelSelector string) ([]appsv1.Deployment, error)`
Lists deployments by namespace and label selector.
- `ctx`: Context for cancellation
- `namespace`: Namespace scope
- `labelSelector`: Kubernetes label selector syntax
- Returns all matching deployments or an error

#### `ListDeploymentsByField(ctx context.Context, namespace string, fieldSelector string) ([]appsv1.Deployment, error)`
Lists deployments by namespace and field selector.
- `ctx`: Context for cancellation
- `namespace`: Namespace scope
- `fieldSelector`: Kubernetes field selector syntax
- Returns all matching deployments or an error

### NamespaceAPI

#### `GetNamespaceByName(ctx context.Context, name string) (*corev1.Namespace, error)`
Retrieves a single Namespace object by its name.
- `ctx`: The context to use for cancellation
- `name`: The name of the Kubernetes namespace to retrieve
- Returns a pointer to a corev1.Namespace object or an error if the namespace is not found

#### `ListNamespacesByLabel(ctx context.Context, labelSelector string) ([]corev1.Namespace, error)`
Retrieves a list of Namespace objects filtered by a label selector.
- `ctx`: The context to use for cancellation
- `labelSelector`: The Kubernetes-compliant label selector string
- Returns a slice of corev1.Namespace objects matching the label selector, or an error

#### `ListNamespacesByField(ctx context.Context, fieldSelector string) ([]corev1.Namespace, error)`
Retrieves a list of Namespace objects filtered by a field selector.
- `ctx`: The context to use for cancellation
- `fieldSelector`: The Kubernetes-compliant field selector string
- Returns a slice of corev1.Namespace objects matching the field selector, or an error

## Validation

The library uses the `github.com/kaudit/val` package for input validation with the following custom validation rules:

### `k8s_label_selector`
Ensures that a string represents a valid Kubernetes label selector. This validator rejects empty values and checks the syntax against Kubernetes label-selector parsing rules.

### `k8s_field_selector`
Ensures that a string represents a valid Kubernetes field selector. This validator enforces a whitelist of recognized field keys to prevent invalid fields.

## Design Principles

1. **Interface-driven design**: All components interact through clearly defined interfaces.
2. **Proper validation**: All inputs are validated with descriptive error messages.
3. **Consistent error handling**: Errors are wrapped with context for better debugging.
4. **Clean API boundaries**: Each resource type has its own focused API surface.

## Requirements

- Go 1.23+
- Access to a Kubernetes cluster

## License

This project is licensed under the [MIT License](./LICENSE)

## Thanks

We would like to express our gratitude to the Kubernetes team and contributors for creating and maintaining the excellent `k8s.io` packages that this library builds upon. Their work on `client-go`, `api`, and `apimachinery` provides the robust foundation that makes this abstraction layer possible.
