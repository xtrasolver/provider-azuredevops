/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	environment "github.com/xtrasolver/provider-azuredevops/internal/controller/azuredevops/environment"
	project "github.com/xtrasolver/provider-azuredevops/internal/controller/azuredevops/project"
	variablegroup "github.com/xtrasolver/provider-azuredevops/internal/controller/azuredevops/variablegroup"
	providerconfig "github.com/xtrasolver/provider-azuredevops/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		environment.Setup,
		project.Setup,
		variablegroup.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
