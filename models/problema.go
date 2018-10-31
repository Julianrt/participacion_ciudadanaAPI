package models

//Problema struct
type Problema struct {
	ID            int64  `json:"id"`
	Foto          string `json:"foto"`
	IDUsuario     int64  `json:"id_usuario"`
	Descripcion   string `json:"descripcion"`
	FechaRegistro string `json:"fecha"`
	IDDependencia int64  `json:"id_dependencia"`
	IDEstado      int64  `json:"id_estado"`
	Direccion     string `json:"direccion"`
}

const problemaScheme string = `CREATE TABLE PRO_PROBLEMAS(
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    foto BLOB,
    id_usuario INT NOT NULL,
    descripcion TEXT,
    fecha_registro TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    id_dependencia INT NOT NULL,
    id_estado INT NOT NULL,
    direccion TEXT);`

//Problemas slice from problema
type Problemas []Problema

//NewProblema functino
func NewProblema() *Problema {
	problema := &Problema{
		Foto:          "",
		IDUsuario:     0,
		Descripcion:   "",
		FechaRegistro: "",
		IDDependencia: 0,
		IDEstado:      0,
		Direccion:     "",
	}
	return problema
}
