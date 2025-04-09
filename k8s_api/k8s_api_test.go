package k8sapi

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"

	"k8s.io/client-go/kubernetes"

	"github.com/kaudit/api"
	mockauth "github.com/kaudit/api/mocks/Authenticator"
)

// TestNewK8sApi tests the successful creation of K8sAPI
func TestNewK8sApi_Success(t *testing.T) {
	// Setup
	mockAuthenticator := mockauth.NewMockAuthenticator(t)
	fakeClient := kubernetes.Clientset{}

	// Expect the authenticator to be called and return a fake client
	mockAuthenticator.EXPECT().NativeAPI().Return(&fakeClient, nil)

	// Call the function under test
	k8sAPI, err := NewK8sAPI(mockAuthenticator)

	// Assertions
	require.NoError(t, err)
	require.NotNil(t, k8sAPI)

	// Verify the returned object is of the correct type
	assert.IsType(t, &K8sAPI{}, k8sAPI, "k8sApi should be of type *K8sAPI")

	// Test PodAPI
	t.Run("GetPodAPI", func(t *testing.T) {
		podAPI := k8sAPI.GetPodAPI()
		assert.NotNil(t, podAPI)
		assert.Implements(t, (*api.PodAPI)(nil), podAPI)
	})

	// Test ServiceAPI
	t.Run("GetServiceAPI", func(t *testing.T) {
		serviceAPI := k8sAPI.GetServiceAPI()
		assert.NotNil(t, serviceAPI)
		assert.Implements(t, (*api.ServiceAPI)(nil), serviceAPI)
	})

	// Test DeploymentAPI
	t.Run("GetDeploymentAPI", func(t *testing.T) {
		deploymentAPI := k8sAPI.GetDeploymentAPI()
		assert.NotNil(t, deploymentAPI)
		assert.Implements(t, (*api.DeploymentAPI)(nil), deploymentAPI)
	})

	// Test NamespaceAPI
	t.Run("GetNamespaceAPI", func(t *testing.T) {
		namespaceAPI := k8sAPI.GetNamespaceAPI()
		assert.NotNil(t, namespaceAPI)
		assert.Implements(t, (*api.NamespaceAPI)(nil), namespaceAPI)
	})

}

// TestNewK8sApi_AuthFailure tests the case when authentication fails
func TestNewK8sApi_AuthFailure(t *testing.T) {
	// Setup
	mockAuthenticator := mockauth.NewMockAuthenticator(t)
	expectedError := errors.New("auth error")

	// Expect the authenticator to be called and return an error
	mockAuthenticator.EXPECT().NativeAPI().Return(nil, expectedError)

	// Call the function under test
	k8sAPI, err := NewK8sAPI(mockAuthenticator)

	// Assertions
	require.Error(t, err)
	assert.Nil(t, k8sAPI)
	assert.Contains(t, err.Error(), "failed to init k8s client")
}

// TestK8sAPI_GetAPIs tests each of the getter methods
func TestK8sAPI_GetAPIs(t *testing.T) {
	// Setup
	mockAuthenticator := mockauth.NewMockAuthenticator(t)
	fakeClient := kubernetes.Clientset{}

	// Expect the authenticator to be called
	mockAuthenticator.EXPECT().NativeAPI().Return(&fakeClient, nil)

	// Create the K8sAPI instance
	k8sAPI, err := NewK8sAPI(mockAuthenticator)
	require.NoError(t, err)

	// Test each getter method returns the correct interface type
	t.Run("GetPodAPI", func(t *testing.T) {
		podAPI := k8sAPI.GetPodAPI()
		assert.NotNil(t, podAPI)
		assert.Implements(t, (*api.PodAPI)(nil), podAPI)
	})

	t.Run("GetServiceAPI", func(t *testing.T) {
		serviceAPI := k8sAPI.GetServiceAPI()
		assert.NotNil(t, serviceAPI)
		assert.Implements(t, (*api.ServiceAPI)(nil), serviceAPI)
	})

	t.Run("GetDeploymentAPI", func(t *testing.T) {
		deploymentAPI := k8sAPI.GetDeploymentAPI()
		assert.NotNil(t, deploymentAPI)
		assert.Implements(t, (*api.DeploymentAPI)(nil), deploymentAPI)
	})

	t.Run("GetNamespaceAPI", func(t *testing.T) {
		namespaceAPI := k8sAPI.GetNamespaceAPI()
		assert.NotNil(t, namespaceAPI)
		assert.Implements(t, (*api.NamespaceAPI)(nil), namespaceAPI)
	})
}

// Test PodAPI Implementation
func TestPodAPIImpl_GetPodByName(t *testing.T) {
	// Setup
	mockAuthenticator := mockauth.NewMockAuthenticator(t)
	fakeClientset := fake.NewClientset(
		&corev1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-pod",
				Namespace: "default",
			},
			Status: corev1.PodStatus{
				Phase: corev1.PodRunning,
			},
		},
	)

	mockAuthenticator.EXPECT().NativeAPI().Return(fakeClientset, nil)

	k8sAPI, err := NewK8sAPI(mockAuthenticator)
	require.NoError(t, err)

	podAPI := k8sAPI.GetPodAPI()
	require.NotNil(t, podAPI)

	// Test successful pod retrieval
	t.Run("GetPodByName_Success", func(t *testing.T) {
		ctx := context.Background()
		pod, err := podAPI.GetPodByName(ctx, "default", "test-pod")

		require.NoError(t, err)
		assert.NotNil(t, pod)
		assert.Equal(t, "test-pod", pod.Name)
		assert.Equal(t, "default", pod.Namespace)
		assert.Equal(t, corev1.PodRunning, pod.Status.Phase)
	})

	// Test pod not found
	t.Run("GetPodByName_NotFound", func(t *testing.T) {
		ctx := context.Background()
		pod, err := podAPI.GetPodByName(ctx, "default", "nonexistent-pod")

		require.Error(t, err)
		assert.Nil(t, pod)
		assert.Contains(t, err.Error(), "not found")
	})
}

func TestPodAPIImpl_ListPodsByLabel(t *testing.T) {
	// Setup
	mockAuthenticator := mockauth.NewMockAuthenticator(t)
	fakeClientset := fake.NewClientset(
		&corev1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-pod-1",
				Namespace: "default",
				Labels: map[string]string{
					"app": "test-app",
				},
			},
		},
		&corev1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-pod-2",
				Namespace: "default",
				Labels: map[string]string{
					"app": "test-app",
				},
			},
		},
		&corev1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "other-pod",
				Namespace: "default",
				Labels: map[string]string{
					"app": "other-app",
				},
			},
		},
	)

	mockAuthenticator.EXPECT().NativeAPI().Return(fakeClientset, nil)

	k8sAPI, err := NewK8sAPI(mockAuthenticator)
	require.NoError(t, err)

	podAPI := k8sAPI.GetPodAPI()
	require.NotNil(t, podAPI)

	// Test listing pods by label
	t.Run("ListPodsByLabel_Success", func(t *testing.T) {
		ctx := context.Background()
		pods, err := podAPI.ListPodsByLabel(ctx, "default", "app=test-app")

		require.NoError(t, err)
		assert.Len(t, pods, 2)
		assert.Contains(t, []string{pods[0].Name, pods[1].Name}, "test-pod-1")
		assert.Contains(t, []string{pods[0].Name, pods[1].Name}, "test-pod-2")
	})

	// Test listing pods with no results
	t.Run("ListPodsByLabel_NoResults", func(t *testing.T) {
		ctx := context.Background()
		pods, err := podAPI.ListPodsByLabel(ctx, "default", "app=nonexistent")

		require.NoError(t, err)
		assert.Empty(t, pods)
	})
}

func TestPodAPIImpl_ListPodsByField(t *testing.T) {
	// Setup
	mockAuthenticator := mockauth.NewMockAuthenticator(t)
	fakeClientset := fake.NewClientset(
		&corev1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-pod-1",
				Namespace: "default",
			},
			Spec: corev1.PodSpec{
				NodeName: "node-1",
			},
		},
		&corev1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-pod-2",
				Namespace: "default",
			},
			Spec: corev1.PodSpec{
				NodeName: "node-1",
			},
		},
	)

	mockAuthenticator.EXPECT().NativeAPI().Return(fakeClientset, nil)

	k8sAPI, err := NewK8sAPI(mockAuthenticator)
	require.NoError(t, err)

	podAPI := k8sAPI.GetPodAPI()
	require.NotNil(t, podAPI)

	// Test listing pods by field
	t.Run("ListPodsByField_Success", func(t *testing.T) {
		ctx := context.Background()
		pods, err := podAPI.ListPodsByField(ctx, "default", "spec.nodeName=node-1")

		require.NoError(t, err)
		assert.Len(t, pods, 2)
	})

	// Test invalid field selector
	t.Run("ListPodsByField_InvalidSelector", func(t *testing.T) {
		ctx := context.Background()
		pods, err := podAPI.ListPodsByField(ctx, "default", "invalid=selector")

		// The behavior for invalid field selectors may vary based on implementation
		// This assertion is just an example
		require.Error(t, err)
		assert.Empty(t, pods)
	})
}

// Test ServiceAPI Implementation
func TestServiceAPIImpl_GetServiceByName(t *testing.T) {
	// Setup
	mockAuthenticator := mockauth.NewMockAuthenticator(t)
	fakeClientset := fake.NewClientset(
		&corev1.Service{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-service",
				Namespace: "default",
			},
			Spec: corev1.ServiceSpec{
				Type: corev1.ServiceTypeClusterIP,
				Ports: []corev1.ServicePort{
					{
						Port: 80,
					},
				},
			},
		},
	)

	mockAuthenticator.EXPECT().NativeAPI().Return(fakeClientset, nil)

	k8sAPI, err := NewK8sAPI(mockAuthenticator)
	require.NoError(t, err)

	serviceAPI := k8sAPI.GetServiceAPI()
	require.NotNil(t, serviceAPI)

	// Test successful service retrieval
	t.Run("GetServiceByName_Success", func(t *testing.T) {
		ctx := context.Background()
		service, err := serviceAPI.GetServiceByName(ctx, "default", "test-service")

		require.NoError(t, err)
		assert.NotNil(t, service)
		assert.Equal(t, "test-service", service.Name)
		assert.Equal(t, "default", service.Namespace)
		assert.Equal(t, corev1.ServiceTypeClusterIP, service.Spec.Type)
		assert.Len(t, service.Spec.Ports, 1)
		assert.Equal(t, int32(80), service.Spec.Ports[0].Port)
	})

	// Test service not found
	t.Run("GetServiceByName_NotFound", func(t *testing.T) {
		ctx := context.Background()
		service, err := serviceAPI.GetServiceByName(ctx, "default", "nonexistent-service")

		require.Error(t, err)
		assert.Nil(t, service)
	})
}

// Test DeploymentAPI Implementation
func TestDeploymentAPIImpl_GetDeploymentByName(t *testing.T) {
	// Setup
	mockAuthenticator := mockauth.NewMockAuthenticator(t)
	replicas := int32(3)
	fakeClientset := fake.NewClientset(
		&appsv1.Deployment{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-deployment",
				Namespace: "default",
			},
			Spec: appsv1.DeploymentSpec{
				Replicas: &replicas,
			},
			Status: appsv1.DeploymentStatus{
				ReadyReplicas: 3,
			},
		},
	)

	mockAuthenticator.EXPECT().NativeAPI().Return(fakeClientset, nil)

	k8sAPI, err := NewK8sAPI(mockAuthenticator)
	require.NoError(t, err)

	deploymentAPI := k8sAPI.GetDeploymentAPI()
	require.NotNil(t, deploymentAPI)

	// Test successful deployment retrieval
	t.Run("GetDeploymentByName_Success", func(t *testing.T) {
		ctx := context.Background()
		deployment, err := deploymentAPI.GetDeploymentByName(ctx, "default", "test-deployment")

		require.NoError(t, err)
		assert.NotNil(t, deployment)
		assert.Equal(t, "test-deployment", deployment.Name)
		assert.Equal(t, "default", deployment.Namespace)
		assert.Equal(t, int32(3), *deployment.Spec.Replicas)
		assert.Equal(t, int32(3), deployment.Status.ReadyReplicas)
	})

	// Test deployment not found
	t.Run("GetDeploymentByName_NotFound", func(t *testing.T) {
		ctx := context.Background()
		deployment, err := deploymentAPI.GetDeploymentByName(ctx, "default", "nonexistent-deployment")

		require.Error(t, err)
		assert.Nil(t, deployment)
	})
}

// Test NamespaceAPI Implementation
func TestNamespaceAPIImpl_GetNamespaceByName(t *testing.T) {
	// Setup
	mockAuthenticator := mockauth.NewMockAuthenticator(t)
	fakeClientset := fake.NewClientset(
		&corev1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				Name: "test-namespace",
				Labels: map[string]string{
					"environment": "test",
				},
			},
			Status: corev1.NamespaceStatus{
				Phase: corev1.NamespaceActive,
			},
		},
	)

	mockAuthenticator.EXPECT().NativeAPI().Return(fakeClientset, nil)

	k8sAPI, err := NewK8sAPI(mockAuthenticator)
	require.NoError(t, err)

	namespaceAPI := k8sAPI.GetNamespaceAPI()
	require.NotNil(t, namespaceAPI)

	// Test successful namespace retrieval
	t.Run("GetNamespaceByName_Success", func(t *testing.T) {
		ctx := context.Background()
		namespace, err := namespaceAPI.GetNamespaceByName(ctx, "test-namespace")

		require.NoError(t, err)
		assert.NotNil(t, namespace)
		assert.Equal(t, "test-namespace", namespace.Name)
		assert.Equal(t, "test", namespace.Labels["environment"])
		assert.Equal(t, corev1.NamespaceActive, namespace.Status.Phase)
	})

	// Test namespace not found
	t.Run("GetNamespaceByName_NotFound", func(t *testing.T) {
		ctx := context.Background()
		namespace, err := namespaceAPI.GetNamespaceByName(ctx, "nonexistent-namespace")

		require.Error(t, err)
		assert.Nil(t, namespace)
	})
}
