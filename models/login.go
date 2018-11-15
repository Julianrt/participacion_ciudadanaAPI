package models

import "log"

//Login struct
type Login struct {
	ID            int64  `json:"id"`
	IDUsuario     int64  `json:"id_usuario"`
	Token         string `json:"token"`
	FechaRegistro string `json:"fecha_registro"`
	SesionActiva  bool   `json:"sesion_activa"`
}

//LoginResponse struct to response
type LoginResponse struct {
	Token     string `json:"token"`
	IDUsuario int64  `json:"id_usuario"`
	Username  string `json:"username"`
}

const loginScheme string = `CREATE TABLE USR_LOGINS(
	id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	id_usuario INT NOT NULL UNIQUE,
	token VARCHAR(50) NOT NULL UNIQUE,
	fecha_registro TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	sesion_activa BIT DEFAULT 1);`

//Logins slice of Login
type Logins []Login

//NewLogin function Constructor
func NewLogin(idUsuario int64, token string) *Login {
	login := &Login{IDUsuario: idUsuario, Token: token}
	return login
}

//CreateLogin function
func CreateLogin(idUsuario int64) (*Login, error) {
	token, _ := RandomHex(20)
	login := NewLogin(idUsuario, token)
	err := login.Save()
	return login, err
}

//CreateToken function
func CreateToken(idUsuario int64) string {
	token, _ := RandomHex(20)
	sql := "INSERT USR_LOGINS SET id_usuario=?, token=?"
	if _, err := Exec(sql, idUsuario, token); err != nil {
		log.Println(err)
	}
	return token
}

//GetToken function
func GetToken(idUsuario int64) string {
	var token string
	sql := "SELECT token FROM USR_LOGINS WHERE id_usuario=?"
	if rows, err := Query(sql, idUsuario); err != nil {
		return token
	} else {
		for rows.Next() {
			rows.Scan(&token)
		}
		return token
	}
}

//GetTokenWithUser function
func GetTokenWithUser(token string) (*LoginResponse, error) {
	tokenlogin := &LoginResponse{"", 0, ""}
	sql := "SELECT l.token, u.id, u.username FROM USR_LOGINS l, USR_USUARIOS u WHERE u.id = l.id_usuario AND l.token = ?"
	rows, err := Query(sql, token)
	if err != nil {
		return tokenlogin, err
	}
	for rows.Next() {
		rows.Scan(&tokenlogin.Token, &tokenlogin.IDUsuario, &tokenlogin.Username)
	}
	return tokenlogin, err
}

//Save function
func (l *Login) Save() error {
	if l.ID == 0 {
		return l.insert()
	}
	return l.update()
}

func (l *Login) insert() error {
	sql := "INSERT USR_LOGINS SET id_usuario=?, token=?"
	loginID, err := InsertData(sql, l.IDUsuario, l.Token)
	l.ID = loginID
	return err
}

func (l *Login) update() error {
	sql := "UPDATE USR_LOGINS SET id_usuario=?, token=? WHERE id=?"
	_, err := Exec(sql, l.IDUsuario, l.Token, l.ID)
	return err
}

//Delete method
func (l *Login) Delete() error {
	sql := "DELETE FROM USR_LOGINS WHERE id=?"
	_, err := Exec(sql, l.ID)
	return err
}
