package models

//Dependencia struct
type Dependencia struct {
	ID          int64  `json:"id"`
	Nombre      string `json:"nombre"`
	NombreCorto string `json:"nombre_corto"`
	Correo      string `json:"correo"`
	Direccion   string `json:"direccion"`
	Telefono    string `json:"telefono"`
}

const dependenciaScheme string = `CREATE TABLE CTL_DEPENDENCIAS(
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    nombre VARCHAR(50) NOT NULL,
    nombre_corto VARCHAR(20),
    correo VARCHAR(100) NOT NULL,
    direccion TEXT,
	telefono VARCHAR(20));`

//Dependencias struct
type Dependencias []Dependencia

//NewDependencia method
func NewDependencia(nombre, nombreCorto, correo, direccion, telefono string) *Dependencia {
	dependencia := &Dependencia{
		Nombre:      nombre,
		NombreCorto: nombreCorto,
		Correo:      correo,
		Direccion:   direccion,
		Telefono:    telefono,
	}
	return dependencia
}

//CreateDependencia method
func CreateDependencia(nombre, nombreCorto, correo, direccion, telefono string) (*Dependencia, error) {
	dependencia := NewDependencia(nombre, nombreCorto, correo, direccion, telefono)
	err := dependencia.Save()
	return dependencia, err
}

//GetDependencia method
func GetDependencia(id int) (*Dependencia, error) {
	dependencia := NewDependencia("", "", "", "", "")
	sql := "SELECT * FROM CTL_DEPENDENCIAS WHERE id=?"
	rows, err := Query(sql, id)
	if err != nil {
		return dependencia, err
	}
	for rows.Next() {
		rows.Scan(&dependencia.ID, &dependencia.Nombre, &dependencia.NombreCorto, &dependencia.Correo, &dependencia.Direccion, &dependencia.Telefono)
	}
	return dependencia, err
}

//GetDependencias method
func GetDependencias() Dependencias {
	var dependencias Dependencias
	sql := "SELECT * FROM CTL_DEPENDENCIAS"
	rows, _ := Query(sql)
	for rows.Next() {
		var dependencia Dependencia
		rows.Scan(&dependencia.ID, &dependencia.Nombre, &dependencia.NombreCorto, &dependencia.Correo, &dependencia.Direccion, &dependencia.Telefono)
		dependencias = append(dependencias, dependencia)
	}
	return dependencias
}

//Save method
func (d *Dependencia) Save() error {
	if d.ID == 0 {
		return d.insert()
	}
	return d.update()
}

func (d *Dependencia) insert() error {
	sql := "INSERT CTL_DEPENDENCIAS SET nombre=?, nombre_corto=?, correo=?, direccion=?, telefono=?"
	id, err := InsertData(sql, d.Nombre, d.NombreCorto, d.Correo, d.Direccion, d.Telefono)
	d.ID = id
	return err
}

func (d *Dependencia) update() error {
	sql := "UPDATE CTL_DEPENDENCIAS SET nombre=?, nombre_corto=?, correo=?, direccion=?, telefono=? WHERE id=?"
	_, err := Exec(sql, d.Nombre, d.NombreCorto, d.Correo, d.Direccion, d.Telefono, d.ID)
	return err
}

//Delete method
func (d *Dependencia) Delete() {
	sql := "DELETE FROM CTL_DEPENDENCIAS WHERE id=?"
	Exec(sql, d.ID)
}
