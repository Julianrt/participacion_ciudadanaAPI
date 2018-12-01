package models

import "errors"

type ValidationError error

var (
	ErrorUsuarioNoEncontrado = ValidationError(errors.New("El usuario no existe en la base de datos"))

	errorUsernameVacio = ValidationError(errors.New("El username no debe de estar vacio"))
	errorUsernameCorto = ValidationError(errors.New("El username es muy corto"))
	errorUsernameLargo = ValidationError(errors.New("El username es muy largo"))
	errorUsernameExistente = ValidationError(errors.New("El username ya existe"))

	errorFormatoCorreo = ValidationError(errors.New("Formato de correo invalido"))
	errorCorreoExistente = ValidationError(errors.New("Ese correo ya se encuentra registrado en nuestra base de datos"))

	errorEncriptacionPassword = ValidationError(errors.New("No es posible cifrar el password"))

	errorLogin = ValidationError(errors.New("Usuario o password incorrectos"))
)
