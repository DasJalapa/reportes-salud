/*Package lib implementa todo el conjunto de errores y expresiones regulares que pueden
ser utilizadas a nivel global de la aplicacion
*/
package lib

import "errors"

var (
	// ErrUnauthenticated error de inicio no correcto
	ErrUnauthenticated = errors.New("Unauthenticated")

	// ErrTokenExpired error de token expirado
	ErrTokenExpired = errors.New("The token was expired")

	// ErrInvalidsignature error de firma inválida
	ErrInvalidsignature = errors.New("The signature is invalid")

	// ErrInvalidToken controlador de cualquier otro error
	ErrInvalidToken = errors.New("Invalid Token")
)

var (
	// ErrUserNotFound error de usuario no encontrado
	ErrUserNotFound = errors.New("User not found")

	//ErrInvalidEmail error de email invalido
	ErrInvalidEmail = errors.New("Invalid email")

	// ErrInvalideUsername error de nombre de usuario invalido
	ErrInvalideUsername = errors.New("Username is invalid")

	//ErrDuplicateUser  error de usuario invalido por que ya existe
	ErrDuplicateUser = errors.New("User already exists")

	//ErrUserOK es una confirmacion de ingresado correctamente
	ErrUserOK = errors.New("User successfully created")

	//ErrNoSeller es error de rolo no vendedor
	ErrNoSeller = errors.New("Request permission to change role")
)

var (
	//ErrFileBig error de maximo peso superado
	ErrFileBig = errors.New("The file exceeds the weight")

	//ErrFileNotSuch error de archivo no encontrado en la peticion
	ErrFileNotSuch = errors.New("File not found in the request")

	//ErrFileNoSoported error de archivo no soportado
	ErrFileNoSoported = errors.New("Invalid file")

	//ErrFileUploadSuccess es resultado satisfactorio de subida
	ErrFileUploadSuccess = errors.New("File successfully uploaded")
)

var (
	// ErrAnyProductFound es el resultado de consulta vacía
	ErrAnyProductFound = errors.New("No product found for this user")
)

var (
	// ErrAnyServiceFound es el resutado de consulta de servicio vacía
	ErrAnyServiceFound = errors.New("No service found for this user")
)

var (
	// ErrNoFoodFound es el resultado de una consulta de food storage vacía
	ErrNoFoodFound = errors.New("No food service found for this user")
)

var (
	// ErrInvalidID error de un id invalido
	ErrInvalidID = errors.New("The ID is invalid")
)

var (
	// ErrNotFound error de ningun registro encontrado
	ErrNotFound = errors.New("No existe ningun registro")
)
