// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	"net/http"

	v1beta1 "github.com/atomix/atomix/stores/pod-memory/pkg/apis/podmemory/v1beta1"
	"github.com/atomix/atomix/stores/pod-memory/pkg/client/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type PodmemoryV1beta1Interface interface {
	RESTClient() rest.Interface
	PodMemoryStoresGetter
}

// PodmemoryV1beta1Client is used to interact with features provided by the podmemory.atomix.io group.
type PodmemoryV1beta1Client struct {
	restClient rest.Interface
}

func (c *PodmemoryV1beta1Client) PodMemoryStores(namespace string) PodMemoryStoreInterface {
	return newPodMemoryStores(c, namespace)
}

// NewForConfig creates a new PodmemoryV1beta1Client for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*PodmemoryV1beta1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	httpClient, err := rest.HTTPClientFor(&config)
	if err != nil {
		return nil, err
	}
	return NewForConfigAndClient(&config, httpClient)
}

// NewForConfigAndClient creates a new PodmemoryV1beta1Client for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*PodmemoryV1beta1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientForConfigAndClient(&config, h)
	if err != nil {
		return nil, err
	}
	return &PodmemoryV1beta1Client{client}, nil
}

// NewForConfigOrDie creates a new PodmemoryV1beta1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *PodmemoryV1beta1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new PodmemoryV1beta1Client for the given RESTClient.
func New(c rest.Interface) *PodmemoryV1beta1Client {
	return &PodmemoryV1beta1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1beta1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *PodmemoryV1beta1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}