package models

//Colonia struct
type Colonia struct {
	ID           int64  `json:"id"`
	Nombre       string `json:"nombre"`
	CodigoPostal string `json:"codigo_postal"`
	IDCiudad     int64  `json:"id_ciudad"`
}

const coloniaScheme string = `CREATE TABLE CTL_COLONIAS(
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    nombre VARCHAR(50) NOT NULL,
    codigo_postal VARCHAR(10),
	id_ciudad INT NOT NULL);`

//Colonias struct
type Colonias []Colonia

//NewColonia method
func NewColonia(nombre, codigoPostal string, idCiudad int64) *Colonia {
	colonia := &Colonia{Nombre: nombre, CodigoPostal: codigoPostal, IDCiudad: idCiudad}
	return colonia
}

//CreateColonia method
func CreateColonia(nombre, codigoPostal string, idCiudad int64) (*Colonia, error) {
	colonia := NewColonia(nombre, codigoPostal, idCiudad)
	err := colonia.Save()
	return colonia, err
}

//GetColonia method
func GetColonia(id int) (*Colonia, error) {
	colonia := NewColonia("", "", 0)
	sql := "SELECT * FROM CTL_COLONIAS WHERE id=?"
	rows, err := Query(sql, id)
	if err != nil {
		return colonia, err
	}
	for rows.Next() {
		rows.Scan(&colonia.ID, &colonia.Nombre, &colonia.CodigoPostal, &colonia.IDCiudad)
	}
	return colonia, nil
}

//GetColonias method
func GetColonias() Colonias {
	var colonias Colonias
	sql := "SELECT * FROM CTL_COLONIAS"
	rows, _ := Query(sql)
	for rows.Next() {
		var colonia Colonia
		rows.Scan(&colonia.ID, &colonia.Nombre, &colonia.CodigoPostal, &colonia.IDCiudad)
		colonias = append(colonias, colonia)
	}
	return colonias
}

//Save method
func (c *Colonia) Save() error {
	if c.ID == 0 {
		return c.insert()
	}
	return c.update()
}

func (c *Colonia) insert() error {
	sql := "INSERT CTL_COLONIAS SET nombre=?, codigo_postal=?, id_ciudad=?"
	id, err := InsertData(sql, c.Nombre, c.CodigoPostal, c.IDCiudad)
	c.ID = id
	return err
}

func (c *Colonia) update() error {
	sql := "UPDATE CTL_COLONIAS SET nombre=?, codigo_postal=?, id_ciudad=? WHERE id=?"
	_, err := Exec(sql, c.Nombre, c.CodigoPostal, c.IDCiudad, c.ID)
	return err
}

//Delete method
func (c *Colonia) Delete() {
	sql := "DELETE FROM CTL_COLONIAS WHERE id=?"
	Exec(sql, c.ID)
}
