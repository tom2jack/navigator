package util

import (
	"fmt"

	apiv1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"
	appslisters "k8s.io/client-go/listers/apps/v1beta1"

	v1alpha1 "github.com/jetstack-experimental/navigator/pkg/apis/navigator/v1alpha1"
)

const (
	NodePoolNameLabelKey      = "navigator.jetstack.io/elasticsearch-node-pool-name"
	NodePoolHashAnnotationKey = "navigator.jetstack.io/elasticsearch-node-pool-hash"
)

func ClusterLabels(c *v1alpha1.ElasticsearchCluster) map[string]string {
	return map[string]string{
		"app":               "elasticsearch",
		ClusterNameLabelKey: c.Name,
	}
}

func NodePoolLabels(c *v1alpha1.ElasticsearchCluster, poolName string, roles ...string) map[string]string {
	labels := ClusterLabels(c)
	if poolName != "" {
		labels[NodePoolNameLabelKey] = poolName
	}
	for _, role := range roles {
		labels[role] = "true"
	}
	return labels
}

func NodePoolResourceName(c *v1alpha1.ElasticsearchCluster, np *v1alpha1.ElasticsearchClusterNodePool) string {
	return fmt.Sprintf("%s-%s", ResourceBaseName(c), np.Name)
}

func SelectorForNodePool(c *v1alpha1.ElasticsearchCluster, np *v1alpha1.ElasticsearchClusterNodePool) (labels.Selector, error) {
	nodePoolNameReq, err := labels.NewRequirement(NodePoolNameLabelKey, selection.Equals, []string{np.Name})
	if err != nil {
		return nil, err
	}
	clusterSelector, err := SelectorForCluster(c)
	if err != nil {
		return nil, err
	}
	return clusterSelector.Add(*nodePoolNameReq), nil
}

func PodControlledByCluster(c *v1alpha1.ElasticsearchCluster, pod *apiv1.Pod, ssLister appslisters.StatefulSetLister) (bool, error) {
	ownerRef := metav1.GetControllerOf(pod)
	if ownerRef == nil || ownerRef.Kind != "StatefulSet" {
		return false, nil
	}
	ss, err := ssLister.StatefulSets(pod.Namespace).Get(ownerRef.Name)
	if apierrors.IsNotFound(err) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return metav1.IsControlledBy(ss, c), nil
}
