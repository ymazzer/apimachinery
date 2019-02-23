/*
Copyright 2019 The Stash Authors.

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
	v1beta1 "github.com/appscode/stash/apis/stash/v1beta1"
	scheme "github.com/appscode/stash/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// BackupConfigurationsGetter has a method to return a BackupConfigurationInterface.
// A group's client should implement this interface.
type BackupConfigurationsGetter interface {
	BackupConfigurations(namespace string) BackupConfigurationInterface
}

// BackupConfigurationInterface has methods to work with BackupConfiguration resources.
type BackupConfigurationInterface interface {
	Create(*v1beta1.BackupConfiguration) (*v1beta1.BackupConfiguration, error)
	Update(*v1beta1.BackupConfiguration) (*v1beta1.BackupConfiguration, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.BackupConfiguration, error)
	List(opts v1.ListOptions) (*v1beta1.BackupConfigurationList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.BackupConfiguration, err error)
	BackupConfigurationExpansion
}

// backupConfigurations implements BackupConfigurationInterface
type backupConfigurations struct {
	client rest.Interface
	ns     string
}

// newBackupConfigurations returns a BackupConfigurations
func newBackupConfigurations(c *StashV1beta1Client, namespace string) *backupConfigurations {
	return &backupConfigurations{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the backupConfiguration, and returns the corresponding backupConfiguration object, and an error if there is any.
func (c *backupConfigurations) Get(name string, options v1.GetOptions) (result *v1beta1.BackupConfiguration, err error) {
	result = &v1beta1.BackupConfiguration{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("backupconfigurations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of BackupConfigurations that match those selectors.
func (c *backupConfigurations) List(opts v1.ListOptions) (result *v1beta1.BackupConfigurationList, err error) {
	result = &v1beta1.BackupConfigurationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("backupconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested backupConfigurations.
func (c *backupConfigurations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("backupconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a backupConfiguration and creates it.  Returns the server's representation of the backupConfiguration, and an error, if there is any.
func (c *backupConfigurations) Create(backupConfiguration *v1beta1.BackupConfiguration) (result *v1beta1.BackupConfiguration, err error) {
	result = &v1beta1.BackupConfiguration{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("backupconfigurations").
		Body(backupConfiguration).
		Do().
		Into(result)
	return
}

// Update takes the representation of a backupConfiguration and updates it. Returns the server's representation of the backupConfiguration, and an error, if there is any.
func (c *backupConfigurations) Update(backupConfiguration *v1beta1.BackupConfiguration) (result *v1beta1.BackupConfiguration, err error) {
	result = &v1beta1.BackupConfiguration{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("backupconfigurations").
		Name(backupConfiguration.Name).
		Body(backupConfiguration).
		Do().
		Into(result)
	return
}

// Delete takes name of the backupConfiguration and deletes it. Returns an error if one occurs.
func (c *backupConfigurations) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("backupconfigurations").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *backupConfigurations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("backupconfigurations").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched backupConfiguration.
func (c *backupConfigurations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.BackupConfiguration, err error) {
	result = &v1beta1.BackupConfiguration{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("backupconfigurations").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
