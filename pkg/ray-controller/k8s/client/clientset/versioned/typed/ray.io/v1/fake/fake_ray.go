/*
Copyright The Kubernetes Authors.

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
	rayiov1 "github.com/ray-operator/pkg/ray-controller/k8s/apis/ray.io/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRays implements RayInterface
type FakeRays struct {
	Fake *FakeRayV1
	ns   string
}

var raysResource = schema.GroupVersionResource{Group: "ray.io", Version: "v1", Resource: "rays"}

var raysKind = schema.GroupVersionKind{Group: "ray.io", Version: "v1", Kind: "Ray"}

// Get takes name of the ray, and returns the corresponding ray object, and an error if there is any.
func (c *FakeRays) Get(name string, options v1.GetOptions) (result *rayiov1.Ray, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(raysResource, c.ns, name), &rayiov1.Ray{})

	if obj == nil {
		return nil, err
	}
	return obj.(*rayiov1.Ray), err
}

// List takes label and field selectors, and returns the list of Rays that match those selectors.
func (c *FakeRays) List(opts v1.ListOptions) (result *rayiov1.RayList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(raysResource, raysKind, c.ns, opts), &rayiov1.RayList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &rayiov1.RayList{ListMeta: obj.(*rayiov1.RayList).ListMeta}
	for _, item := range obj.(*rayiov1.RayList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested rays.
func (c *FakeRays) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(raysResource, c.ns, opts))

}

// Create takes the representation of a ray and creates it.  Returns the server's representation of the ray, and an error, if there is any.
func (c *FakeRays) Create(ray *rayiov1.Ray) (result *rayiov1.Ray, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(raysResource, c.ns, ray), &rayiov1.Ray{})

	if obj == nil {
		return nil, err
	}
	return obj.(*rayiov1.Ray), err
}

// Update takes the representation of a ray and updates it. Returns the server's representation of the ray, and an error, if there is any.
func (c *FakeRays) Update(ray *rayiov1.Ray) (result *rayiov1.Ray, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(raysResource, c.ns, ray), &rayiov1.Ray{})

	if obj == nil {
		return nil, err
	}
	return obj.(*rayiov1.Ray), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRays) UpdateStatus(ray *rayiov1.Ray) (*rayiov1.Ray, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(raysResource, "status", c.ns, ray), &rayiov1.Ray{})

	if obj == nil {
		return nil, err
	}
	return obj.(*rayiov1.Ray), err
}

// Delete takes name of the ray and deletes it. Returns an error if one occurs.
func (c *FakeRays) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(raysResource, c.ns, name), &rayiov1.Ray{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRays) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(raysResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &rayiov1.RayList{})
	return err
}

// Patch applies the patch and returns the patched ray.
func (c *FakeRays) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *rayiov1.Ray, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(raysResource, c.ns, name, data, subresources...), &rayiov1.Ray{})

	if obj == nil {
		return nil, err
	}
	return obj.(*rayiov1.Ray), err
}