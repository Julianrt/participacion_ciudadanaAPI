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

//NewProblema function
func NewProblema(foto string, idUsuario int64, descripcion string, idDependencia, idEstado int64, direccion string) *Problema {
	problema := &Problema{
		Foto:          foto,
		IDUsuario:     idUsuario,
		Descripcion:   descripcion,
		IDDependencia: idDependencia,
		IDEstado:      idEstado,
		Direccion:     direccion,
	}
	return problema
}

//CreateProblema function
func CreateProblema(foto string, idUsuario int64, descripcion string, idDependencia, idEstado int64, direccion string) (*Problema, error) {
	problema := NewProblema(foto, idUsuario, descripcion, idDependencia, idEstado, direccion)
	err := problema.Save()
	return problema, err
}

//GetProblema function
func GetProblema(id int) (*Problema, error) {
	problema := NewProblema("", 0, "", 0, 0, "")
	sql := "SELECT * FROM PRO_PROBLEMAS WHERE id=?"
	rows, err := Query(sql, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		rows.Scan(&problema.ID, &problema.Foto, &problema.IDUsuario, &problema.Descripcion, &problema.FechaRegistro, &problema.IDDependencia, &problema.IDEstado, &problema.Direccion)
	}
	return problema, err
}

//GetProblemas function
func GetProblemas() Problemas {
	var problemas Problemas
	sql := "SELECT * FROM PRO_PROBLEMAS"
	rows, _ := Query(sql)
	for rows.Next() {
		var problema Problema
		rows.Scan(&problema.ID, &problema.Foto, &problema.IDUsuario, &problema.Descripcion, &problema.FechaRegistro, &problema.IDDependencia, &problema.IDEstado, &problema.Direccion)
		problemas = append(problemas, problema)
	}
	return problemas
}

//Save method
func (p *Problema) Save() error {
	if p.ID == 0 {
		return p.insert()
	}
	return p.update()
}

func (p *Problema) insert() error {
	sql := "INSERT PRO_PROBLEMAS SET foto=?, id_usuario=?, descripcion=?, id_dependencia=?, id_estado=?, direccion=?"
	id, err := InsertData(sql, p.Foto, p.IDUsuario, p.Descripcion, p.IDDependencia, p.IDEstado, p.Direccion)
	p.ID = id
	return err
}

func (p *Problema) update() error {
	sql := "UPDATE PRO_PROBLEMAS SET foto=?, id_usuario=?, descripcion=?, id_dependencia=?, id_estado=?, direccion=? WHERE id=?"
	_, err := Exec(sql, p.Foto, p.IDUsuario, p.Descripcion, p.IDDependencia, p.IDEstado, p.Direccion, p.ID)
	return err
}

//Delete method
func (p *Problema) Delete() {
	sql := "DELETE FROM PRO_PROBLEMAS WHERE id=?"
	Exec(sql, p.ID)
}
