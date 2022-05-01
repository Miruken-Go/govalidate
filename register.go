package govalidate

import "github.com/miruken-go/miruken"

// GoValidationInstaller enables validation support
// for https://github.com/asaskevich/govalidator
type GoValidationInstaller struct{}

func (v *GoValidationInstaller) Install(registration *miruken.Registration) {
	if registration.CanInstall(&_registrationTag) {
		registration.Install(miruken.WithValidation())
		registration.AddHandlerTypes(miruken.TypeOf[*validator]())
	}
}

func WithGoValidation(
	config ...func(installer *GoValidationInstaller),
) miruken.Installer {
	installer := &GoValidationInstaller{}
	for _, configure := range config {
		if configure != nil {
			configure(installer)
		}
	}
	return installer
}

var _registrationTag byte
