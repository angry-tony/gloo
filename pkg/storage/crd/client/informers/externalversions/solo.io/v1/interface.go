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

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	internalinterfaces "github.com/solo-io/gloo/pkg/storage/crd/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Attributes returns a AttributeInformer.
	Attributes() AttributeInformer
	// Roles returns a RoleInformer.
	Roles() RoleInformer
	// Upstreams returns a UpstreamInformer.
	Upstreams() UpstreamInformer
	// VirtualServices returns a VirtualServiceInformer.
	VirtualServices() VirtualServiceInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// Attributes returns a AttributeInformer.
func (v *version) Attributes() AttributeInformer {
	return &attributeInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Roles returns a RoleInformer.
func (v *version) Roles() RoleInformer {
	return &roleInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Upstreams returns a UpstreamInformer.
func (v *version) Upstreams() UpstreamInformer {
	return &upstreamInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// VirtualServices returns a VirtualServiceInformer.
func (v *version) VirtualServices() VirtualServiceInformer {
	return &virtualServiceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
