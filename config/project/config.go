/*
Copyright 2021 Upbound Inc.
*/

package project

import (
	ujconfig "github.com/upbound/upjet/pkg/config"
)

// Configure configures the null group
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("azuredevops_project", func(r *ujconfig.Resource) {
		r.Kind = "Project"
		r.ShortGroup = ""
		// And other overrides.
	})
}
