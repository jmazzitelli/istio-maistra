// Code generated by xns-informer-gen. DO NOT EDIT.

package v1alpha1

import (
	xnsinformers "github.com/maistra/xns-informer/pkg/informers"
	v1alpha1 "k8s.io/api/rbac/v1alpha1"
	informers "k8s.io/client-go/informers/rbac/v1alpha1"
	listers "k8s.io/client-go/listers/rbac/v1alpha1"
	"k8s.io/client-go/tools/cache"
)

type roleBindingInformer struct {
	informer cache.SharedIndexInformer
}

var _ informers.RoleBindingInformer = &roleBindingInformer{}

func NewRoleBindingInformer(f xnsinformers.SharedInformerFactory) informers.RoleBindingInformer {
	resource := v1alpha1.SchemeGroupVersion.WithResource("rolebindings")
	converter := xnsinformers.NewListWatchConverter(
		f.GetScheme(),
		&v1alpha1.RoleBinding{},
		&v1alpha1.RoleBindingList{},
	)

	informer := f.ForResource(resource, xnsinformers.ResourceOptions{
		ClusterScoped:      false,
		ListWatchConverter: converter,
	})

	return &roleBindingInformer{informer: informer.Informer()}
}

func (i *roleBindingInformer) Informer() cache.SharedIndexInformer {
	return i.informer
}

func (i *roleBindingInformer) Lister() listers.RoleBindingLister {
	return listers.NewRoleBindingLister(i.informer.GetIndexer())
}