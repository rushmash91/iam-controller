// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// OpenIDConnectProviderSpec defines the desired state of OpenIDConnectProvider.
type OpenIDConnectProviderSpec struct {

	// Provides a list of client IDs, also known as audiences. When a mobile or
	// web app registers with an OpenID Connect provider, they establish a value
	// that identifies the application. This is the value that's sent as the client_id
	// parameter on OAuth requests.
	//
	// You can register multiple client IDs with the same provider. For example,
	// you might have multiple applications that use the same OIDC provider. You
	// cannot register more than 100 client IDs with a single IAM OIDC provider.
	//
	// There is no defined format for a client ID. The CreateOpenIDConnectProviderRequest
	// operation accepts client IDs up to 255 characters long.
	ClientIDs []*string `json:"clientIDs,omitempty"`
	// A list of tags that you want to attach to the new IAM OpenID Connect (OIDC)
	// provider. Each tag consists of a key name and an associated value. For more
	// information about tagging, see Tagging IAM resources (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html)
	// in the IAM User Guide.
	//
	// If any one of the tags is invalid or if you exceed the allowed maximum number
	// of tags, then the entire request fails and the resource is not created.
	Tags []*Tag `json:"tags,omitempty"`
	// A list of server certificate thumbprints for the OpenID Connect (OIDC) identity
	// provider's server certificates. Typically this list includes only one entry.
	// However, IAM lets you have up to five thumbprints for an OIDC provider. This
	// lets you maintain multiple thumbprints if the identity provider is rotating
	// certificates.
	//
	// This parameter is optional. If it is not included, IAM will retrieve and
	// use the top intermediate certificate authority (CA) thumbprint of the OpenID
	// Connect identity provider server certificate.
	//
	// The server certificate thumbprint is the hex-encoded SHA-1 hash value of
	// the X.509 certificate used by the domain where the OpenID Connect provider
	// makes its keys available. It is always a 40-character string.
	//
	// For example, assume that the OIDC provider is server.example.com and the
	// provider stores its keys at https://keys.server.example.com/openid-connect.
	// In that case, the thumbprint string would be the hex-encoded SHA-1 hash value
	// of the certificate used by https://keys.server.example.com.
	//
	// For more information about obtaining the OIDC provider thumbprint, see Obtaining
	// the thumbprint for an OpenID Connect provider (https://docs.aws.amazon.com/IAM/latest/UserGuide/identity-providers-oidc-obtain-thumbprint.html)
	// in the IAM user Guide.
	Thumbprints []*string `json:"thumbprints,omitempty"`
	// The URL of the identity provider. The URL must begin with https:// and should
	// correspond to the iss claim in the provider's OpenID Connect ID tokens. Per
	// the OIDC standard, path components are allowed but query parameters are not.
	// Typically the URL consists of only a hostname, like https://server.example.org
	// or https://example.com. The URL should not contain a port number.
	//
	// You cannot register the same provider multiple times in a single Amazon Web
	// Services account. If you try to submit a URL that has already been used for
	// an OpenID Connect provider in the Amazon Web Services account, you will get
	// an error.
	// +kubebuilder:validation:Required
	URL *string `json:"url"`
}

// OpenIDConnectProviderStatus defines the observed state of OpenIDConnectProvider
type OpenIDConnectProviderStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	// +kubebuilder:validation:Optional
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRS managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	// +kubebuilder:validation:Optional
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
}

// OpenIDConnectProvider is the Schema for the OpenIDConnectProviders API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type OpenIDConnectProvider struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OpenIDConnectProviderSpec   `json:"spec,omitempty"`
	Status            OpenIDConnectProviderStatus `json:"status,omitempty"`
}

// OpenIDConnectProviderList contains a list of OpenIDConnectProvider
// +kubebuilder:object:root=true
type OpenIDConnectProviderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OpenIDConnectProvider `json:"items"`
}

func init() {
	SchemeBuilder.Register(&OpenIDConnectProvider{}, &OpenIDConnectProviderList{})
}
