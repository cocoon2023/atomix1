// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by client-gen. DO NOT EDIT.

package v1beta2

import (
	"context"
	"time"

	v1beta2 "github.com/atomix/atomix/stores/raft/pkg/apis/raft/v1beta2"
	scheme "github.com/atomix/atomix/stores/raft/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// RaftStoresGetter has a method to return a RaftStoreInterface.
// A group's client should implement this interface.
type RaftStoresGetter interface {
	RaftStores(namespace string) RaftStoreInterface
}

// RaftStoreInterface has methods to work with RaftStore resources.
type RaftStoreInterface interface {
	Create(ctx context.Context, raftStore *v1beta2.RaftStore, opts v1.CreateOptions) (*v1beta2.RaftStore, error)
	Update(ctx context.Context, raftStore *v1beta2.RaftStore, opts v1.UpdateOptions) (*v1beta2.RaftStore, error)
	UpdateStatus(ctx context.Context, raftStore *v1beta2.RaftStore, opts v1.UpdateOptions) (*v1beta2.RaftStore, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta2.RaftStore, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta2.RaftStoreList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta2.RaftStore, err error)
	RaftStoreExpansion
}

// raftStores implements RaftStoreInterface
type raftStores struct {
	client rest.Interface
	ns     string
}

// newRaftStores returns a RaftStores
func newRaftStores(c *RaftV1beta2Client, namespace string) *raftStores {
	return &raftStores{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the raftStore, and returns the corresponding raftStore object, and an error if there is any.
func (c *raftStores) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta2.RaftStore, err error) {
	result = &v1beta2.RaftStore{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("raftstores").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of RaftStores that match those selectors.
func (c *raftStores) List(ctx context.Context, opts v1.ListOptions) (result *v1beta2.RaftStoreList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta2.RaftStoreList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("raftstores").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested raftStores.
func (c *raftStores) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("raftstores").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a raftStore and creates it.  Returns the server's representation of the raftStore, and an error, if there is any.
func (c *raftStores) Create(ctx context.Context, raftStore *v1beta2.RaftStore, opts v1.CreateOptions) (result *v1beta2.RaftStore, err error) {
	result = &v1beta2.RaftStore{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("raftstores").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(raftStore).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a raftStore and updates it. Returns the server's representation of the raftStore, and an error, if there is any.
func (c *raftStores) Update(ctx context.Context, raftStore *v1beta2.RaftStore, opts v1.UpdateOptions) (result *v1beta2.RaftStore, err error) {
	result = &v1beta2.RaftStore{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("raftstores").
		Name(raftStore.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(raftStore).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *raftStores) UpdateStatus(ctx context.Context, raftStore *v1beta2.RaftStore, opts v1.UpdateOptions) (result *v1beta2.RaftStore, err error) {
	result = &v1beta2.RaftStore{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("raftstores").
		Name(raftStore.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(raftStore).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the raftStore and deletes it. Returns an error if one occurs.
func (c *raftStores) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("raftstores").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *raftStores) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("raftstores").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched raftStore.
func (c *raftStores) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta2.RaftStore, err error) {
	result = &v1beta2.RaftStore{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("raftstores").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
