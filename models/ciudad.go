package models

//Ciudad struct
type Ciudad struct {
	ID          int64  `json:"id"`
	Nombre      string `json:"nombre"`
	NombreCorto string `json:"nombre_corto"`
}

const ciudadScheme string = `CREATE TABLE CTL_CIUDADES(
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    nombre VARCHAR(50) NOT NULL,
	nombre_corto VARCHAR(20));`

//Ciudades struct
type Ciudades []Ciudad

//NewCiudad method
func NewCiudad(nombre, nombreCorto string) *Ciudad {
	ciudad := &Ciudad{Nombre: nombre, NombreCorto: nombreCorto}
	return ciudad
}

//CreateCiudad method
func CreateCiudad(nombre, nombreCorto string) (*Ciudad, error) {
	ciudad := NewCiudad(nombre, nombreCorto)
	err := ciudad.Save()
	return ciudad, err
}

//GetCiudad method
func GetCiudad(id int) (*Ciudad, error) {
	ciudad := NewCiudad("", "")
	sql := "SELECT id, nombre, nombre_corto FROM CTL_CIUDADES WHERE id=?"
	rows, err := Query(sql, id)
	if err != nil {
		return ciudad, err
	}
	for rows.Next() {
		rows.Scan(&ciudad.ID, &ciudad.Nombre, &ciudad.NombreCorto)
	}
	return ciudad, nil
}

//GetCiudades method
func GetCiudades() Ciudades {
	var ciudades Ciudades
	sql := "SELECT * FROM CTL_CIUDADES"
	rows, _ := Query(sql)
	for rows.Next() {
		var ciudad Ciudad
		rows.Scan(&ciudad.ID, &ciudad.Nombre, &ciudad.NombreCorto)
		ciudades = append(ciudades, ciudad)
	}
	return ciudades
}

//Save method
func (c *Ciudad) Save() error {
	if c.ID == 0 {
		return c.insert()
	}
	return c.update()
}

func (c *Ciudad) insert() error {
	sql := "INSERT CTL_CIUDADES SET nombre=?, nombre_corto=?"
	id, err := InsertData(sql, c.Nombre, c.NombreCorto)
	c.ID = id
	return err
}

func (c *Ciudad) update() error {
	sql := "UPDATE CTL_CIUDADES SET nombre=?, nombre_corto=? WHERE id=?"
	_, err := Exec(sql, c.Nombre, c.NombreCorto, c.ID)
	return err
}

//Delete method
func (c *Ciudad) Delete() {
	sql := "DELETE FROM CTL_CIUDADES WHERE id=?"
	Exec(sql, c.ID)
}
