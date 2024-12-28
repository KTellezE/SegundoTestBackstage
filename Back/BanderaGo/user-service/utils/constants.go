package utils

type Constants struct {
	MessageErrorID           string
	MessageErrorJson         string
	MessageErrorCreation     string
	MessageErrorGetUsers     string
	MessageErrorUserNotFount string
	MessageErrorUpdateUser   string
	MessageErrorDeleteUser   string
}

var DefaultConstants = Constants{
	MessageErrorID:           "ID de usuario inválido",
	MessageErrorJson:         "Error al decodificar el JSON",
	MessageErrorCreation:     "Error al crear el usuario",
	MessageErrorGetUsers:     "Error al obtener los usuarios",
	MessageErrorUserNotFount: "Usuario no encontrado",
	MessageErrorUpdateUser:   "No fue posible actualizar el usuario",
	MessageErrorDeleteUser:   "No fue posible eliminar el usuario",
}
