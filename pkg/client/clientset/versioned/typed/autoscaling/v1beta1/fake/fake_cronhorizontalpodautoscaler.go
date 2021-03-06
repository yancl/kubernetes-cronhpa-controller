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
	v1beta1 "github.com/AliyunContainerService/kubernetes-cronhpa-controller/pkg/apis/autoscaling/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCronHorizontalPodAutoscalers implements CronHorizontalPodAutoscalerInterface
type FakeCronHorizontalPodAutoscalers struct {
	Fake *FakeAutoscalingV1beta1
	ns   string
}

var cronhorizontalpodautoscalersResource = schema.GroupVersionResource{Group: "autoscaling.alibabacloud.com", Version: "v1beta1", Resource: "cronhorizontalpodautoscalers"}

var cronhorizontalpodautoscalersKind = schema.GroupVersionKind{Group: "autoscaling.alibabacloud.com", Version: "v1beta1", Kind: "CronHorizontalPodAutoscaler"}

// Get takes name of the cronHorizontalPodAutoscaler, and returns the corresponding cronHorizontalPodAutoscaler object, and an error if there is any.
func (c *FakeCronHorizontalPodAutoscalers) Get(name string, options v1.GetOptions) (result *v1beta1.CronHorizontalPodAutoscaler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(cronhorizontalpodautoscalersResource, c.ns, name), &v1beta1.CronHorizontalPodAutoscaler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CronHorizontalPodAutoscaler), err
}

// List takes label and field selectors, and returns the list of CronHorizontalPodAutoscalers that match those selectors.
func (c *FakeCronHorizontalPodAutoscalers) List(opts v1.ListOptions) (result *v1beta1.CronHorizontalPodAutoscalerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(cronhorizontalpodautoscalersResource, cronhorizontalpodautoscalersKind, c.ns, opts), &v1beta1.CronHorizontalPodAutoscalerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.CronHorizontalPodAutoscalerList{ListMeta: obj.(*v1beta1.CronHorizontalPodAutoscalerList).ListMeta}
	for _, item := range obj.(*v1beta1.CronHorizontalPodAutoscalerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cronHorizontalPodAutoscalers.
func (c *FakeCronHorizontalPodAutoscalers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(cronhorizontalpodautoscalersResource, c.ns, opts))

}

// Create takes the representation of a cronHorizontalPodAutoscaler and creates it.  Returns the server's representation of the cronHorizontalPodAutoscaler, and an error, if there is any.
func (c *FakeCronHorizontalPodAutoscalers) Create(cronHorizontalPodAutoscaler *v1beta1.CronHorizontalPodAutoscaler) (result *v1beta1.CronHorizontalPodAutoscaler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(cronhorizontalpodautoscalersResource, c.ns, cronHorizontalPodAutoscaler), &v1beta1.CronHorizontalPodAutoscaler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CronHorizontalPodAutoscaler), err
}

// Update takes the representation of a cronHorizontalPodAutoscaler and updates it. Returns the server's representation of the cronHorizontalPodAutoscaler, and an error, if there is any.
func (c *FakeCronHorizontalPodAutoscalers) Update(cronHorizontalPodAutoscaler *v1beta1.CronHorizontalPodAutoscaler) (result *v1beta1.CronHorizontalPodAutoscaler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(cronhorizontalpodautoscalersResource, c.ns, cronHorizontalPodAutoscaler), &v1beta1.CronHorizontalPodAutoscaler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CronHorizontalPodAutoscaler), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCronHorizontalPodAutoscalers) UpdateStatus(cronHorizontalPodAutoscaler *v1beta1.CronHorizontalPodAutoscaler) (*v1beta1.CronHorizontalPodAutoscaler, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(cronhorizontalpodautoscalersResource, "status", c.ns, cronHorizontalPodAutoscaler), &v1beta1.CronHorizontalPodAutoscaler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CronHorizontalPodAutoscaler), err
}

// Delete takes name of the cronHorizontalPodAutoscaler and deletes it. Returns an error if one occurs.
func (c *FakeCronHorizontalPodAutoscalers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(cronhorizontalpodautoscalersResource, c.ns, name), &v1beta1.CronHorizontalPodAutoscaler{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCronHorizontalPodAutoscalers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(cronhorizontalpodautoscalersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1beta1.CronHorizontalPodAutoscalerList{})
	return err
}

// Patch applies the patch and returns the patched cronHorizontalPodAutoscaler.
func (c *FakeCronHorizontalPodAutoscalers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.CronHorizontalPodAutoscaler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(cronhorizontalpodautoscalersResource, c.ns, name, pt, data, subresources...), &v1beta1.CronHorizontalPodAutoscaler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CronHorizontalPodAutoscaler), err
}
