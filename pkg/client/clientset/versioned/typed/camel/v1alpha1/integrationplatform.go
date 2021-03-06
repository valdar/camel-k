/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"time"

	v1alpha1 "github.com/apache/camel-k/pkg/apis/camel/v1alpha1"
	scheme "github.com/apache/camel-k/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// IntegrationPlatformsGetter has a method to return a IntegrationPlatformInterface.
// A group's client should implement this interface.
type IntegrationPlatformsGetter interface {
	IntegrationPlatforms(namespace string) IntegrationPlatformInterface
}

// IntegrationPlatformInterface has methods to work with IntegrationPlatform resources.
type IntegrationPlatformInterface interface {
	Create(*v1alpha1.IntegrationPlatform) (*v1alpha1.IntegrationPlatform, error)
	Update(*v1alpha1.IntegrationPlatform) (*v1alpha1.IntegrationPlatform, error)
	UpdateStatus(*v1alpha1.IntegrationPlatform) (*v1alpha1.IntegrationPlatform, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.IntegrationPlatform, error)
	List(opts v1.ListOptions) (*v1alpha1.IntegrationPlatformList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.IntegrationPlatform, err error)
	IntegrationPlatformExpansion
}

// integrationPlatforms implements IntegrationPlatformInterface
type integrationPlatforms struct {
	client rest.Interface
	ns     string
}

// newIntegrationPlatforms returns a IntegrationPlatforms
func newIntegrationPlatforms(c *CamelV1alpha1Client, namespace string) *integrationPlatforms {
	return &integrationPlatforms{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the integrationPlatform, and returns the corresponding integrationPlatform object, and an error if there is any.
func (c *integrationPlatforms) Get(name string, options v1.GetOptions) (result *v1alpha1.IntegrationPlatform, err error) {
	result = &v1alpha1.IntegrationPlatform{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("integrationplatforms").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of IntegrationPlatforms that match those selectors.
func (c *integrationPlatforms) List(opts v1.ListOptions) (result *v1alpha1.IntegrationPlatformList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.IntegrationPlatformList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("integrationplatforms").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested integrationPlatforms.
func (c *integrationPlatforms) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("integrationplatforms").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a integrationPlatform and creates it.  Returns the server's representation of the integrationPlatform, and an error, if there is any.
func (c *integrationPlatforms) Create(integrationPlatform *v1alpha1.IntegrationPlatform) (result *v1alpha1.IntegrationPlatform, err error) {
	result = &v1alpha1.IntegrationPlatform{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("integrationplatforms").
		Body(integrationPlatform).
		Do().
		Into(result)
	return
}

// Update takes the representation of a integrationPlatform and updates it. Returns the server's representation of the integrationPlatform, and an error, if there is any.
func (c *integrationPlatforms) Update(integrationPlatform *v1alpha1.IntegrationPlatform) (result *v1alpha1.IntegrationPlatform, err error) {
	result = &v1alpha1.IntegrationPlatform{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("integrationplatforms").
		Name(integrationPlatform.Name).
		Body(integrationPlatform).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *integrationPlatforms) UpdateStatus(integrationPlatform *v1alpha1.IntegrationPlatform) (result *v1alpha1.IntegrationPlatform, err error) {
	result = &v1alpha1.IntegrationPlatform{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("integrationplatforms").
		Name(integrationPlatform.Name).
		SubResource("status").
		Body(integrationPlatform).
		Do().
		Into(result)
	return
}

// Delete takes name of the integrationPlatform and deletes it. Returns an error if one occurs.
func (c *integrationPlatforms) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("integrationplatforms").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *integrationPlatforms) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("integrationplatforms").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched integrationPlatform.
func (c *integrationPlatforms) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.IntegrationPlatform, err error) {
	result = &v1alpha1.IntegrationPlatform{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("integrationplatforms").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
