/*
Copyright 2021 Upbound Inc.
*/

package serviceendpointkubernetes

import (
	ujconfig "github.com/upbound/upjet/pkg/config"
)

// Configure configures the null group
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("azuredevops_serviceendpoint_kubernetes", func(r *ujconfig.Resource) {
		r.Kind = "ServiceEndpointKubernetes"
		r.ShortGroup = ""
		r.References["project_id"] = ujconfig.Reference{
            Type: "github.com/xtrasolver/provider-azuredevops/apis/azuredevops/v1alpha1.Project",
        }
		r.LateInitializer = ujconfig.LateInitializer{
			IgnoredFields: []string{
				"authorization",
			},
		}		
		// And other overrides.		
	})
}
