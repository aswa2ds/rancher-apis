/*
Copyright 2022 Rancher Labs, Inc.

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

// Code generated by main. DO NOT EDIT.

// +k8s:deepcopy-gen=package
// +groupName=project.cattle.io
package v3

import (
	project "github.com/rancher/rancher/pkg/apis/project.cattle.io"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	AppResourceName                           = "apps"
	AppRevisionResourceName                   = "apprevisions"
	BasicAuthResourceName                     = "basicauths"
	CertificateResourceName                   = "certificates"
	DockerCredentialResourceName              = "dockercredentials"
	NamespacedBasicAuthResourceName           = "namespacedbasicauths"
	NamespacedCertificateResourceName         = "namespacedcertificates"
	NamespacedDockerCredentialResourceName    = "namespaceddockercredentials"
	NamespacedSSHAuthResourceName             = "namespacedsshauths"
	NamespacedServiceAccountTokenResourceName = "namespacedserviceaccounttokens"
	PipelineResourceName                      = "pipelines"
	PipelineExecutionResourceName             = "pipelineexecutions"
	PipelineSettingResourceName               = "pipelinesettings"
	SSHAuthResourceName                       = "sshauths"
	ServiceAccountTokenResourceName           = "serviceaccounttokens"
	SourceCodeCredentialResourceName          = "sourcecodecredentials"
	SourceCodeProviderResourceName            = "sourcecodeproviders"
	SourceCodeProviderConfigResourceName      = "sourcecodeproviderconfigs"
	SourceCodeRepositoryResourceName          = "sourcecoderepositories"
	WorkloadResourceName                      = "workloads"
)

// SchemeGroupVersion is group version used to register these objects
var SchemeGroupVersion = schema.GroupVersion{Group: project.GroupName, Version: "v3"}

// Kind takes an unqualified kind and returns back a Group qualified GroupKind
func Kind(kind string) schema.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

// Resource takes an unqualified resource and returns a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

var (
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	AddToScheme   = SchemeBuilder.AddToScheme
)

// Adds the list of known types to Scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&App{},
		&AppList{},
		&AppRevision{},
		&AppRevisionList{},
		&BasicAuth{},
		&BasicAuthList{},
		&Certificate{},
		&CertificateList{},
		&DockerCredential{},
		&DockerCredentialList{},
		&NamespacedBasicAuth{},
		&NamespacedBasicAuthList{},
		&NamespacedCertificate{},
		&NamespacedCertificateList{},
		&NamespacedDockerCredential{},
		&NamespacedDockerCredentialList{},
		&NamespacedSSHAuth{},
		&NamespacedSSHAuthList{},
		&NamespacedServiceAccountToken{},
		&NamespacedServiceAccountTokenList{},
		&Pipeline{},
		&PipelineList{},
		&PipelineExecution{},
		&PipelineExecutionList{},
		&PipelineSetting{},
		&PipelineSettingList{},
		&SSHAuth{},
		&SSHAuthList{},
		&ServiceAccountToken{},
		&ServiceAccountTokenList{},
		&SourceCodeCredential{},
		&SourceCodeCredentialList{},
		&SourceCodeProvider{},
		&SourceCodeProviderList{},
		&SourceCodeProviderConfig{},
		&SourceCodeProviderConfigList{},
		&SourceCodeRepository{},
		&SourceCodeRepositoryList{},
		&Workload{},
		&WorkloadList{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
