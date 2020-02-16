/*
Copyright The Stash Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	"time"

	v1beta1 "stash.appscode.dev/apimachinery/apis/stash/v1beta1"
	scheme "stash.appscode.dev/apimachinery/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// RestoreSessionsGetter has a method to return a RestoreSessionInterface.
// A group's client should implement this interface.
type RestoreSessionsGetter interface {
	RestoreSessions(namespace string) RestoreSessionInterface
}

// RestoreSessionInterface has methods to work with RestoreSession resources.
type RestoreSessionInterface interface {
	Create(*v1beta1.RestoreSession) (*v1beta1.RestoreSession, error)
	Update(*v1beta1.RestoreSession) (*v1beta1.RestoreSession, error)
	UpdateStatus(*v1beta1.RestoreSession) (*v1beta1.RestoreSession, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.RestoreSession, error)
	List(opts v1.ListOptions) (*v1beta1.RestoreSessionList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.RestoreSession, err error)
	RestoreSessionExpansion
}

// restoreSessions implements RestoreSessionInterface
type restoreSessions struct {
	client rest.Interface
	ns     string
}

// newRestoreSessions returns a RestoreSessions
func newRestoreSessions(c *StashV1beta1Client, namespace string) *restoreSessions {
	return &restoreSessions{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the restoreSession, and returns the corresponding restoreSession object, and an error if there is any.
func (c *restoreSessions) Get(name string, options v1.GetOptions) (result *v1beta1.RestoreSession, err error) {
	result = &v1beta1.RestoreSession{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("restoresessions").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of RestoreSessions that match those selectors.
func (c *restoreSessions) List(opts v1.ListOptions) (result *v1beta1.RestoreSessionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.RestoreSessionList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("restoresessions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested restoreSessions.
func (c *restoreSessions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("restoresessions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a restoreSession and creates it.  Returns the server's representation of the restoreSession, and an error, if there is any.
func (c *restoreSessions) Create(restoreSession *v1beta1.RestoreSession) (result *v1beta1.RestoreSession, err error) {
	result = &v1beta1.RestoreSession{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("restoresessions").
		Body(restoreSession).
		Do().
		Into(result)
	return
}

// Update takes the representation of a restoreSession and updates it. Returns the server's representation of the restoreSession, and an error, if there is any.
func (c *restoreSessions) Update(restoreSession *v1beta1.RestoreSession) (result *v1beta1.RestoreSession, err error) {
	result = &v1beta1.RestoreSession{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("restoresessions").
		Name(restoreSession.Name).
		Body(restoreSession).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *restoreSessions) UpdateStatus(restoreSession *v1beta1.RestoreSession) (result *v1beta1.RestoreSession, err error) {
	result = &v1beta1.RestoreSession{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("restoresessions").
		Name(restoreSession.Name).
		SubResource("status").
		Body(restoreSession).
		Do().
		Into(result)
	return
}

// Delete takes name of the restoreSession and deletes it. Returns an error if one occurs.
func (c *restoreSessions) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("restoresessions").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *restoreSessions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("restoresessions").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched restoreSession.
func (c *restoreSessions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.RestoreSession, err error) {
	result = &v1beta1.RestoreSession{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("restoresessions").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
