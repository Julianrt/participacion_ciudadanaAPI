package models

//Voto struct
type Voto struct {
	ID            int64  `json:"id"`
	IDProblema    int64  `json:"id_problema"`
	IDUsuario     int64  `json:"id_usuario"`
	Calificacion  int    `json:"calificacion"`
	FechaRegistro string `json:"fecha_registro"`
}

const votoScheme string = `CREATE TABLE PRO_VOTOS(
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    id_problema INT NOT NULL,
    id_usuario INT NOT NULL,
    calificacion INT NOT NULL,
    fecha_registro TIMESTAMP DEFAULT CURRENT_TIMESTAMP);`

//Votos a slice from voto
type Votos []Voto

//NewVoto function, return a voto struct
func NewVoto(idProblema, idUsuario int64, calificacion int) *Voto {
	voto := &Voto{
		IDProblema:   idProblema,
		IDUsuario:    idUsuario,
		Calificacion: calificacion,
	}
	return voto
}

//CreateVoto function, return a store in db a voto struct
func CreateVoto(idProblema, idUsuario int64, calificacion int) (*Voto, error) {
	voto := NewVoto(idProblema, idUsuario, calificacion)
	err := voto.Save()
	return voto, err
}

//GetVoto function
func GetVoto(id int) (*Voto, error) {
	voto := NewVoto(0, 0, 0)
	sql := "SELECT * FROM PRO_VOTOS WHERE id=?"
	rows, err := Query(sql, id)
	if err != nil {
		return voto, err
	}
	for rows.Next() {
		rows.Scan(&voto.ID, &voto.IDProblema, &voto.IDUsuario, &voto.Calificacion, &voto.FechaRegistro)
	}
	return voto, err
}

//GetVotos function
func GetVotos() Votos {
	var votos Votos
	sql := "SELECT * FROM PRO_VOTOS"
	rows, _ := Query(sql)
	for rows.Next() {
		var voto Voto
		rows.Scan(&voto.ID, &voto.IDProblema, &voto.IDUsuario, &voto.Calificacion, &voto.FechaRegistro)
		votos = append(votos, voto)
	}
	return votos
}

//Save method, insert or update a voto struct in db and return error
func (v *Voto) Save() error {
	if v.ID == 0 {
		return v.insert()
	}
	return v.update()
}

func (v *Voto) insert() error {
	sql := "INSERT PRO_VOTOS SET id_problema=?, id_usuario=?, calificacion=?"
	id, err := InsertData(sql, v.IDProblema, v.IDUsuario, v.Calificacion)
	v.ID = id
	return err
}

func (v *Voto) update() error {
	sql := "UPDATE PRO_VOTOS SET id_problema=?, id_usuario=?, calificacion=? WHERE id=?"
	_, err := Exec(sql, v.IDProblema, v.IDUsuario, v.Calificacion, v.ID)
	return err
}

//Delete method, delete a voto from db
func (v *Voto) Delete() {
	sql := "DELETE FROM PRO_VOTOS WHERE id=?"
	Exec(sql, v.ID)
}
