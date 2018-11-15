package models

//Usuario struct
type Usuario struct {
	ID            int64  `json:"id"`
	Nombre        string `json:"nombre"`
	Username      string `json:"username"`
	Correo        string `json:"correo"`
	Pass          string `json:"pass"`
	Telefono      string `json:"telefono"`
	IDColonia     int64  `json:"id_colonia"`
	FechaRegistro string `json:"fecha_registro"`
}

const usuarioScheme string = `CREATE TABLE USR_USUARIOS(
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    nombre VARCHAR(50) NOT NULL,
    username VARCHAR(50) NOT NULL UNIQUE,
    correo VARCHAR(100) NOT NULL,
    pass VARCHAR(50) NOT NULL,
    telefono VARCHAR(20),
    id_colonia INT NOT NULL,
    fecha_registro TIMESTAMP DEFAULT CURRENT_TIMESTAMP);`

//Usuarios struct
type Usuarios []Usuario

//NewUsuario function
func NewUsuario(nombre, username, correo, pass, telefono string, idColonia int64) *Usuario {
	usuario := &Usuario{
		Nombre:    nombre,
		Username:  username,
		Correo:    correo,
		Pass:      pass,
		Telefono:  telefono,
		IDColonia: idColonia,
	}
	return usuario
}

//CreateUsuario function
func CreateUsuario(nombre, username, correo, pass, telefono string, idColonia int64) (*Usuario, error) {
	usuario := NewUsuario(nombre, username, correo, pass, telefono, idColonia)
	if err := usuario.Save(); err != nil {
		return usuario, err
	}
	return usuario, nil
}

//GetUsuario function
func GetUsuario(id int) (*Usuario, error) {
	usuario := NewUsuario("", "", "", "", "", 0)
	sql := "SELECT * FROM USR_USUARIOS WHERE id=?"
	rows, err := Query(sql, id)
	if err != nil {
		return usuario, err
	}
	for rows.Next() {
		rows.Scan(&usuario.ID, &usuario.Nombre, &usuario.Username, &usuario.Correo, &usuario.Pass, &usuario.Telefono, &usuario.IDColonia, &usuario.FechaRegistro)
	}
	return usuario, err
}

func GetUsuarioByUsername(username, pass string) (*Usuario, error) {
	usuario := NewUsuario("", "", "", "", "", 0)
	sql := "SELECT * FROM USR_USUARIOS WHERE username=? AND pass=?"
	if rows, err := Query(sql, username, pass); err != nil {
		return usuario, err
	} else {
		for rows.Next() {
			rows.Scan(&usuario.ID, &usuario.Nombre, &usuario.Username, &usuario.Correo, &usuario.Pass, &usuario.Telefono, &usuario.IDColonia, &usuario.FechaRegistro)
		}
		return usuario, err
	}
}

//GetUsuarios function
func GetUsuarios() Usuarios {
	var usuarios Usuarios
	sql := "SELECT * FROM USR_USUARIOS"
	rows, _ := Query(sql)
	for rows.Next() {
		var usuario Usuario
		rows.Scan(&usuario.ID, &usuario.Nombre, &usuario.Username, &usuario.Correo, &usuario.Pass, &usuario.Telefono, &usuario.IDColonia, &usuario.FechaRegistro)
		usuarios = append(usuarios, usuario)
	}
	return usuarios
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
func (u *Usuario) Delete() {
	sql := "DELETE FROM USR_USUARIOS WHERE id=?"
	Exec(sql, u.ID)
}
