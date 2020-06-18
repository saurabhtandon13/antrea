// Copyright 2020 Antrea Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"time"

	v1alpha1 "github.com/vmware-tanzu/antrea/pkg/apis/ops/v1alpha1"
	scheme "github.com/vmware-tanzu/antrea/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// TraceflowsGetter has a method to return a TraceflowInterface.
// A group's client should implement this interface.
type TraceflowsGetter interface {
	Traceflows() TraceflowInterface
}

// TraceflowInterface has methods to work with Traceflow resources.
type TraceflowInterface interface {
	Create(*v1alpha1.Traceflow) (*v1alpha1.Traceflow, error)
	Update(*v1alpha1.Traceflow) (*v1alpha1.Traceflow, error)
	UpdateStatus(*v1alpha1.Traceflow) (*v1alpha1.Traceflow, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Traceflow, error)
	List(opts v1.ListOptions) (*v1alpha1.TraceflowList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Traceflow, err error)
	TraceflowExpansion
}

// traceflows implements TraceflowInterface
type traceflows struct {
	client rest.Interface
}

// newTraceflows returns a Traceflows
func newTraceflows(c *OpsV1alpha1Client) *traceflows {
	return &traceflows{
		client: c.RESTClient(),
	}
}

// Get takes name of the traceflow, and returns the corresponding traceflow object, and an error if there is any.
func (c *traceflows) Get(name string, options v1.GetOptions) (result *v1alpha1.Traceflow, err error) {
	result = &v1alpha1.Traceflow{}
	err = c.client.Get().
		Resource("traceflows").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Traceflows that match those selectors.
func (c *traceflows) List(opts v1.ListOptions) (result *v1alpha1.TraceflowList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.TraceflowList{}
	err = c.client.Get().
		Resource("traceflows").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested traceflows.
func (c *traceflows) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("traceflows").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a traceflow and creates it.  Returns the server's representation of the traceflow, and an error, if there is any.
func (c *traceflows) Create(traceflow *v1alpha1.Traceflow) (result *v1alpha1.Traceflow, err error) {
	result = &v1alpha1.Traceflow{}
	err = c.client.Post().
		Resource("traceflows").
		Body(traceflow).
		Do().
		Into(result)
	return
}

// Update takes the representation of a traceflow and updates it. Returns the server's representation of the traceflow, and an error, if there is any.
func (c *traceflows) Update(traceflow *v1alpha1.Traceflow) (result *v1alpha1.Traceflow, err error) {
	result = &v1alpha1.Traceflow{}
	err = c.client.Put().
		Resource("traceflows").
		Name(traceflow.Name).
		Body(traceflow).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *traceflows) UpdateStatus(traceflow *v1alpha1.Traceflow) (result *v1alpha1.Traceflow, err error) {
	result = &v1alpha1.Traceflow{}
	err = c.client.Put().
		Resource("traceflows").
		Name(traceflow.Name).
		SubResource("status").
		Body(traceflow).
		Do().
		Into(result)
	return
}

// Delete takes name of the traceflow and deletes it. Returns an error if one occurs.
func (c *traceflows) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("traceflows").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *traceflows) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("traceflows").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched traceflow.
func (c *traceflows) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Traceflow, err error) {
	result = &v1alpha1.Traceflow{}
	err = c.client.Patch(pt).
		Resource("traceflows").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
