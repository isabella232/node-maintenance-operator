// +build !ignore_autogenerated

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1beta1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"kubevirt.io/node-maintenance-operator/pkg/apis/kubevirt/v1beta1.NodeMaintenance": schema_pkg_apis_kubevirt_v1beta1_NodeMaintenance(ref),
	}
}

func schema_pkg_apis_kubevirt_v1beta1_NodeMaintenance(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NodeMaintenance is the Schema for the nodemaintenances API kubebuilder:subresource:status",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("kubevirt.io/node-maintenance-operator/pkg/apis/kubevirt/v1beta1.NodeMaintenanceSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("kubevirt.io/node-maintenance-operator/pkg/apis/kubevirt/v1beta1.NodeMaintenanceStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta", "kubevirt.io/node-maintenance-operator/pkg/apis/kubevirt/v1beta1.NodeMaintenanceSpec", "kubevirt.io/node-maintenance-operator/pkg/apis/kubevirt/v1beta1.NodeMaintenanceStatus"},
	}
}
