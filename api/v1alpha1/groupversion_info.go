// Package v1alpha1 contains the API types for the
// platform.example.com group at version v1alpha1.
//
// +kubebuilder:object:generate=true
// +groupName=platform.example.com
package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

// GroupVersion for platform.example.com API group
var GroupVersion = schema.GroupVersion{Group: "platform.example.com", Version: "v1alpha1"}

var (
	// SchemeBuilder registers our API types
	SchemeBuilder = &scheme.Builder{GroupVersion: GroupVersion}
	// AddToScheme adds the types to the scheme
	AddToScheme = SchemeBuilder.AddToScheme
)

func init() { // Register API types
	SchemeBuilder.Register(&Tenant{}, &TenantList{})
}
