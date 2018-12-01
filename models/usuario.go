package models

import (
	"time"
	"regexp"
	"golang.org/x/crypto/bcrypt"
)

//Usuario struct
type Usuario struct {
	ID            int64  `json:"id"`
	Nombre        string `json:"nombre"`
	Username      string `json:"username"`
	Correo        string `json:"correo"`
	Pass          string `json:"pass"`
	Telefono      string `json:"telefono"`
	IDColonia     int64  `json:"id_colonia"`
	fechaRegistro time.Time
}

var emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

const usuarioScheme string = `CREATE TABLE USR_USUARIOS(
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    nombre VARCHAR(50) NOT NULL,
    username VARCHAR(50) NOT NULL UNIQUE,
    correo VARCHAR(100) NOT NULL UNIQUE,
    pass VARCHAR(60) NOT NULL,
    telefono VARCHAR(20),
    id_colonia INT NOT NULL,
    fecha_registro TIMESTAMP DEFAULT CURRENT_TIMESTAMP);`

//Usuarios struct
type Usuarios []Usuario

//NewUsuario function
func NewUsuario(nombre, username, correo, pass, telefono string, idColonia int64) (*Usuario, error) {
	usuario := &Usuario{
		Nombre:    nombre,
		Username:  username,
		Correo:    correo,
		Pass:      pass,
		Telefono:  telefono,
		IDColonia: idColonia,
	}
	if err := usuario.Valid(); err != nil {
		return &Usuario{}, err
	}
	err := usuario.SetPassword(pass)
	return usuario, err
}

//CreateUsuario function
func CreateUsuario(nombre, username, correo, pass, telefono string, idColonia int64) (*Usuario, error) {
	usuario, err := NewUsuario(nombre, username, correo, pass, telefono, idColonia)
	if err != nil {
		return &Usuario{}, err
	}
	err = usuario.Save()
	if err != nil && err.Error() == "Error 1062: Duplicate entry '"+usuario.Username+"' for key 'username'" {
		err = errorUsernameExistente
	}
	return usuario, err
}

//GetUsuario function
func getUsuario(sql string, condicion interface{}) (*Usuario, error) {
	usuario := &Usuario{}
	rows, err := Query(sql, condicion)
	for rows.Next() {
		rows.Scan(&usuario.ID, &usuario.Nombre, &usuario.Username, &usuario.Correo, &usuario.Pass, &usuario.Telefono, &usuario.IDColonia)
	}
	return usuario, err
}

func GetUsuarioByID(id int) (*Usuario, error) {
	sql := "SELECT id, nombre, username, correo, pass, telefono, id_colonia FROM USR_USUARIOS WHERE id=?"
	return getUsuario(sql, id)
}

func GetUsuarioByUsername(username string) (*Usuario, error) {
	sql := "SELECT id, nombre, username, correo, pass, telefono, id_colonia FROM USR_USUARIOS WHERE username=?"
	return getUsuario(sql, username)
}

//GetUsuarios function
func GetUsuarios() Usuarios {
	var usuarios Usuarios
	sql := "SELECT id, nombre, username, correo, pass, telefono, id_colonia FROM USR_USUARIOS"
	rows, _ := Query(sql)
	for rows.Next() {
		var usuario Usuario
		rows.Scan(&usuario.ID, &usuario.Nombre, &usuario.Username, &usuario.Correo, &usuario.Pass, &usuario.Telefono, &usuario.IDColonia)
		usuarios = append(usuarios, usuario)
	}
	return usuarios
}

func LoginUsuario(username, password string) (*Usuario, error) {
	usuario,_ := GetUsuarioByUsername(username)
	err := bcrypt.CompareHashAndPassword([]byte(usuario.Pass),[]byte(password))
	if err != nil {
		return &Usuario{}, errorLogin
	}
	return usuario, nil
}

func ValidEmail(email string) error {
	if !emailRegexp.MatchString(email){
		return errorFormatoCorreo
	}
	return nil
}

func ValidUsername(username string) error {
	if username == "" {
		return errorUsernameVacio
	}
	if len(username) < 5 {
		return errorUsernameCorto
	}
	if len(username) > 30 {
		return errorUsernameLargo
	}
	return nil
}

func (u *Usuario) Valid() error {
	if err := ValidEmail(u.Correo); err != nil {
		return err
	}
	if err := ValidUsername(u.Username); err != nil {
		return err
	}
	return nil
}

//Save method
func (u *Usuario) Save() error {
	if u.ID == 0 {
		return u.insert()
	}
	return u.update()
}

func (u *Usuario) insert() error {
	sql := "INSERT USR_USUARIOS SET nombre=?, username=?, correo=?, pass=?, telefono=?, id_colonia=?"
	usuarioID, err := InsertData(sql, u.Nombre, u.Username, u.Correo, u.Pass, u.Telefono, u.IDColonia)
	u.ID = usuarioID
	return err
}

func (u *Usuario) update() error {
	sql := "UPDATE USR_USUARIOS SET nombre=?, username=?, correo=?, pass=?, telefono=?, id_colonia=? WHERE id=?"
	_, err := Exec(sql, u.Nombre, u.Username, u.Correo, u.Pass, u.Telefono, u.IDColonia, u.ID)
	return err
}

//Delete method
func (u *Usuario) Delete() error {
	sql := "DELETE FROM USR_USUARIOS WHERE id=?"
	_, err := Exec(sql, u.ID)
	return err
}

func (u *Usuario) SetPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errorEncriptacionPassword
	}
	u.Pass = string(hash)
	return nil
}

func (u *Usuario) GetFechaRegistro() time.Time {
	return u.fechaRegistro
}
