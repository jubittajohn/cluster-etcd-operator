// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1 "github.com/openshift/api/operator/v1"
	operatorv1 "github.com/openshift/client-go/operator/applyconfigurations/operator/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeConsoles implements ConsoleInterface
type FakeConsoles struct {
	Fake *FakeOperatorV1
}

var consolesResource = v1.SchemeGroupVersion.WithResource("consoles")

var consolesKind = v1.SchemeGroupVersion.WithKind("Console")

// Get takes name of the console, and returns the corresponding console object, and an error if there is any.
func (c *FakeConsoles) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Console, err error) {
	emptyResult := &v1.Console{}
	obj, err := c.Fake.
		Invokes(testing.NewRootGetActionWithOptions(consolesResource, name, options), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Console), err
}

// List takes label and field selectors, and returns the list of Consoles that match those selectors.
func (c *FakeConsoles) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ConsoleList, err error) {
	emptyResult := &v1.ConsoleList{}
	obj, err := c.Fake.
		Invokes(testing.NewRootListActionWithOptions(consolesResource, consolesKind, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.ConsoleList{ListMeta: obj.(*v1.ConsoleList).ListMeta}
	for _, item := range obj.(*v1.ConsoleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested consoles.
func (c *FakeConsoles) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchActionWithOptions(consolesResource, opts))
}

// Create takes the representation of a console and creates it.  Returns the server's representation of the console, and an error, if there is any.
func (c *FakeConsoles) Create(ctx context.Context, console *v1.Console, opts metav1.CreateOptions) (result *v1.Console, err error) {
	emptyResult := &v1.Console{}
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateActionWithOptions(consolesResource, console, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Console), err
}

// Update takes the representation of a console and updates it. Returns the server's representation of the console, and an error, if there is any.
func (c *FakeConsoles) Update(ctx context.Context, console *v1.Console, opts metav1.UpdateOptions) (result *v1.Console, err error) {
	emptyResult := &v1.Console{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateActionWithOptions(consolesResource, console, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Console), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeConsoles) UpdateStatus(ctx context.Context, console *v1.Console, opts metav1.UpdateOptions) (result *v1.Console, err error) {
	emptyResult := &v1.Console{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceActionWithOptions(consolesResource, "status", console, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Console), err
}

// Delete takes name of the console and deletes it. Returns an error if one occurs.
func (c *FakeConsoles) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(consolesResource, name, opts), &v1.Console{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeConsoles) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewRootDeleteCollectionActionWithOptions(consolesResource, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1.ConsoleList{})
	return err
}

// Patch applies the patch and returns the patched console.
func (c *FakeConsoles) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Console, err error) {
	emptyResult := &v1.Console{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(consolesResource, name, pt, data, opts, subresources...), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Console), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied console.
func (c *FakeConsoles) Apply(ctx context.Context, console *operatorv1.ConsoleApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Console, err error) {
	if console == nil {
		return nil, fmt.Errorf("console provided to Apply must not be nil")
	}
	data, err := json.Marshal(console)
	if err != nil {
		return nil, err
	}
	name := console.Name
	if name == nil {
		return nil, fmt.Errorf("console.Name must be provided to Apply")
	}
	emptyResult := &v1.Console{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(consolesResource, *name, types.ApplyPatchType, data, opts.ToPatchOptions()), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Console), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeConsoles) ApplyStatus(ctx context.Context, console *operatorv1.ConsoleApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Console, err error) {
	if console == nil {
		return nil, fmt.Errorf("console provided to Apply must not be nil")
	}
	data, err := json.Marshal(console)
	if err != nil {
		return nil, err
	}
	name := console.Name
	if name == nil {
		return nil, fmt.Errorf("console.Name must be provided to Apply")
	}
	emptyResult := &v1.Console{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(consolesResource, *name, types.ApplyPatchType, data, opts.ToPatchOptions(), "status"), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Console), err
}
