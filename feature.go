package govalidate

import "github.com/miruken-go/miruken"

// GoValidationInstaller integrates validation
// support for the go validator.
// for https://github.com/asaskevich/govalidator
type GoValidationInstaller struct{}

func (v *GoValidationInstaller) DependsOn() []miruken.Feature {
	return []miruken.Feature{miruken.WithValidation()}
}

func (v *GoValidationInstaller) Install(setup *miruken.SetupBuilder) error {
	if setup.CanInstall(&_featureTag) {
		setup.RegisterHandlers(&validator{})
	}
	return nil
}

func WithGoValidation(
	config ...func(installer *GoValidationInstaller),
) miruken.Feature {
	installer := &GoValidationInstaller{}
	for _, configure := range config {
		if configure != nil {
			configure(installer)
		}
	}
	return installer
}

var _featureTag byte
