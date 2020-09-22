/*
Copyright 2020 The Kubernetes Authors.

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

package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "sigs.k8s.io/secrets-store-csi-driver/apis/v1alpha1"
)

// FakeSecretProviderClasses implements SecretProviderClassInterface
type FakeSecretProviderClasses struct {
	Fake *FakeSecretsstoreV1alpha1
	ns   string
}

var secretproviderclassesResource = schema.GroupVersionResource{Group: "secrets-store.csi.x-k8s.io", Version: "v1alpha1", Resource: "secretproviderclasses"}

var secretproviderclassesKind = schema.GroupVersionKind{Group: "secrets-store.csi.x-k8s.io", Version: "v1alpha1", Kind: "SecretProviderClass"}

// Get takes name of the secretProviderClass, and returns the corresponding secretProviderClass object, and an error if there is any.
func (c *FakeSecretProviderClasses) Get(name string, options v1.GetOptions) (result *v1alpha1.SecretProviderClass, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(secretproviderclassesResource, c.ns, name), &v1alpha1.SecretProviderClass{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SecretProviderClass), err
}

// List takes label and field selectors, and returns the list of SecretProviderClasses that match those selectors.
func (c *FakeSecretProviderClasses) List(opts v1.ListOptions) (result *v1alpha1.SecretProviderClassList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(secretproviderclassesResource, secretproviderclassesKind, c.ns, opts), &v1alpha1.SecretProviderClassList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SecretProviderClassList{ListMeta: obj.(*v1alpha1.SecretProviderClassList).ListMeta}
	for _, item := range obj.(*v1alpha1.SecretProviderClassList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested secretProviderClasses.
func (c *FakeSecretProviderClasses) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(secretproviderclassesResource, c.ns, opts))

}

// Create takes the representation of a secretProviderClass and creates it.  Returns the server's representation of the secretProviderClass, and an error, if there is any.
func (c *FakeSecretProviderClasses) Create(secretProviderClass *v1alpha1.SecretProviderClass) (result *v1alpha1.SecretProviderClass, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(secretproviderclassesResource, c.ns, secretProviderClass), &v1alpha1.SecretProviderClass{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SecretProviderClass), err
}

// Update takes the representation of a secretProviderClass and updates it. Returns the server's representation of the secretProviderClass, and an error, if there is any.
func (c *FakeSecretProviderClasses) Update(secretProviderClass *v1alpha1.SecretProviderClass) (result *v1alpha1.SecretProviderClass, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(secretproviderclassesResource, c.ns, secretProviderClass), &v1alpha1.SecretProviderClass{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SecretProviderClass), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSecretProviderClasses) UpdateStatus(secretProviderClass *v1alpha1.SecretProviderClass) (*v1alpha1.SecretProviderClass, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(secretproviderclassesResource, "status", c.ns, secretProviderClass), &v1alpha1.SecretProviderClass{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SecretProviderClass), err
}

// Delete takes name of the secretProviderClass and deletes it. Returns an error if one occurs.
func (c *FakeSecretProviderClasses) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(secretproviderclassesResource, c.ns, name), &v1alpha1.SecretProviderClass{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSecretProviderClasses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(secretproviderclassesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.SecretProviderClassList{})
	return err
}

// Patch applies the patch and returns the patched secretProviderClass.
func (c *FakeSecretProviderClasses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SecretProviderClass, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(secretproviderclassesResource, c.ns, name, pt, data, subresources...), &v1alpha1.SecretProviderClass{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SecretProviderClass), err
}
