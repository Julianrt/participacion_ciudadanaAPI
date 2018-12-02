package models

import "log"

type Login struct {
	ID            	int64  `json:"id"`
	IDUsuario     	int64  `json:"id_usuario"`
	Uuid        	string `json:"uuid"`
	FechaRegistro 	string `json:"fecha_registro"`
	SesionActiva  	bool   `json:"sesion_activa"`
}

type TokenResponse struct {
	Uuid     	string `json:"uuid"`
	IDUsuario 	int64  `json:"id_usuario"`
	Username  	string `json:"username"`
}

const loginScheme string = `CREATE TABLE USR_LOGINS(
	id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	id_usuario INT NOT NULL UNIQUE,
	token VARCHAR(50) NOT NULL UNIQUE,
	fecha_registro TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	sesion_activa BIT DEFAULT 1);`

type Logins []Login

func NewLogin(idUsuario int64, uuid string) *Login {
	login := &Login{IDUsuario: idUsuario, Uuid: uuid}
	return login
}

func CreateLogin(idUsuario int64) (*Login, error) {
	token, _ := RandomHex(20)
	login := NewLogin(idUsuario, token)
	err := login.Save()
	return login, err
}

func CreateUuid(idUsuario int64) string {
	uuid, _ := RandomHex(20)
	sql := "INSERT USR_LOGINS SET id_usuario=?, token=?"
	if _, err := Exec(sql, idUsuario, uuid); err != nil {
		log.Println(err)
	}
	return uuid
}

func GetUuidByUserID(idUsuario int64) string {
	var uuid string
	sql := "SELECT token FROM USR_LOGINS WHERE id_usuario=?"
	if rows, err := Query(sql, idUsuario); err != nil {
		return uuid
	} else {
		for rows.Next() {
			rows.Scan(&uuid)
		}
		return uuid
	}
}

func GetTokenWithUuid(uuid string) *TokenResponse {
	tokenResponse := &TokenResponse{"", 0, ""}
	sql := "SELECT l.token, u.id, u.username FROM USR_LOGINS l, USR_USUARIOS u WHERE u.id = l.id_usuario AND l.token = ?"
	rows, _ := Query(sql, uuid)
	for rows.Next() {
		rows.Scan(&tokenResponse.Uuid, &tokenResponse.IDUsuario, &tokenResponse.Username)
	}
	return tokenResponse
}

func (l *Login) Save() error {
	if l.ID == 0 {
		return l.insert()
	}
	return l.update()
}

func (l *Login) insert() error {
	sql := "INSERT USR_LOGINS SET id_usuario=?, token=?"
	loginID, err := InsertData(sql, l.IDUsuario, l.Uuid)
	l.ID = loginID
	return err
}

func (l *Login) update() error {
	sql := "UPDATE USR_LOGINS SET id_usuario=?, token=? WHERE id=?"
	_, err := Exec(sql, l.IDUsuario, l.Uuid, l.ID)
	return err
}

func (l *Login) Delete() error {
	sql := "DELETE FROM USR_LOGINS WHERE id=?"
	_, err := Exec(sql, l.ID)
	return err
}
