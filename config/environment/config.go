/*
Copyright 2021 Upbound Inc.
*/

package environment

import (
	ujconfig "github.com/upbound/upjet/pkg/config"
)

// Configure configures the null group
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("azuredevops_environment", func(r *ujconfig.Resource) {
		r.Kind = "Environment"
		r.ShortGroup = ""
		r.References["project_id"] = ujconfig.Reference{
            Type: "github.com/xtrasolver/provider-azuredevops/apis/azuredevops/v1alpha1.Project",
        }
		// And other overrides.		
	})
}
