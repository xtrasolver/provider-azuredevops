/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	"github.com/pkg/errors"

	"github.com/upbound/upjet/pkg/resource"
	"github.com/upbound/upjet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this Environment
func (mg *Environment) GetTerraformResourceType() string {
	return "azuredevops_environment"
}

// GetConnectionDetailsMapping for this Environment
func (tr *Environment) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this Environment
func (tr *Environment) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this Environment
func (tr *Environment) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this Environment
func (tr *Environment) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this Environment
func (tr *Environment) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this Environment
func (tr *Environment) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this Environment using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *Environment) LateInitialize(attrs []byte) (bool, error) {
	params := &EnvironmentParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *Environment) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this Project
func (mg *Project) GetTerraformResourceType() string {
	return "azuredevops_project"
}

// GetConnectionDetailsMapping for this Project
func (tr *Project) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this Project
func (tr *Project) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this Project
func (tr *Project) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this Project
func (tr *Project) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this Project
func (tr *Project) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this Project
func (tr *Project) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this Project using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *Project) LateInitialize(attrs []byte) (bool, error) {
	params := &ProjectParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *Project) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this ServiceEndpointKubernetes
func (mg *ServiceEndpointKubernetes) GetTerraformResourceType() string {
	return "azuredevops_serviceendpoint_kubernetes"
}

// GetConnectionDetailsMapping for this ServiceEndpointKubernetes
func (tr *ServiceEndpointKubernetes) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"kubeconfig[*].kube_config": "spec.forProvider.kubeconfig[*].kubeConfigSecretRef", "service_account[*].ca_cert": "spec.forProvider.serviceAccount[*].caCertSecretRef", "service_account[*].ca_cert_hash": "status.atProvider.serviceAccount[*].caCertHash", "service_account[*].token": "spec.forProvider.serviceAccount[*].tokenSecretRef", "service_account[*].token_hash": "status.atProvider.serviceAccount[*].tokenHash"}
}

// GetObservation of this ServiceEndpointKubernetes
func (tr *ServiceEndpointKubernetes) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this ServiceEndpointKubernetes
func (tr *ServiceEndpointKubernetes) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this ServiceEndpointKubernetes
func (tr *ServiceEndpointKubernetes) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this ServiceEndpointKubernetes
func (tr *ServiceEndpointKubernetes) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this ServiceEndpointKubernetes
func (tr *ServiceEndpointKubernetes) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this ServiceEndpointKubernetes using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *ServiceEndpointKubernetes) LateInitialize(attrs []byte) (bool, error) {
	params := &ServiceEndpointKubernetesParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}
	opts = append(opts, resource.WithNameFilter("Authorization"))

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *ServiceEndpointKubernetes) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this VariableGroup
func (mg *VariableGroup) GetTerraformResourceType() string {
	return "azuredevops_variable_group"
}

// GetConnectionDetailsMapping for this VariableGroup
func (tr *VariableGroup) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"variable[*].secret_value": "spec.forProvider.variable[*].secretValueSecretRef"}
}

// GetObservation of this VariableGroup
func (tr *VariableGroup) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this VariableGroup
func (tr *VariableGroup) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this VariableGroup
func (tr *VariableGroup) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this VariableGroup
func (tr *VariableGroup) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this VariableGroup
func (tr *VariableGroup) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this VariableGroup using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *VariableGroup) LateInitialize(attrs []byte) (bool, error) {
	params := &VariableGroupParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *VariableGroup) GetTerraformSchemaVersion() int {
	return 0
}
