package govalidate

import "github.com/miruken-go/miruken"

// GoValidationInstaller enables validation support
// for https://github.com/asaskevich/govalidator
type GoValidationInstaller struct{}

func (v *GoValidationInstaller) Install(setup *miruken.SetupBuilder) error {
	if setup.CanInstall(&_featureTag) {
		setup.Install(miruken.WithValidation())
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
