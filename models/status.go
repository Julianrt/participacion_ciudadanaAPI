package models

//Status struct
type Status struct {
	ID     int64  `json:"id"`
	Nombre string `json:"nombre"`
}

const statusScheme string = `CREATE TABLE CTL_STATUS(
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    nombre VARCHAR(50));`

//Statuses a slice from Status
type Statuses []Status

//NewStatus function
func NewStatus(nombre string) *Status {
	status := &Status{Nombre: nombre}
	return status
}

//CreateStatus function
func CreateStatus(nombre string) (*Status, error) {
	status := NewStatus(nombre)
	err := status.Save()
	return status, err
}

//GetStatus function
func GetStatus(id int) (*Status, error) {
	status := NewStatus("")
	sql := "SELECT * FROM CTL_STATUS WHERE id=?"
	rows, err := Query(sql, id)
	if err != nil {
		return status, err
	}
	for rows.Next() {
		rows.Scan(&status.ID, &status.Nombre)
	}
	return status, err
}

//GetAllStatus function
func GetAllStatus() Statuses {
	var statuses Statuses
	sql := "SELECT * FROM CTL_STATUS"
	rows, _ := Query(sql)
	for rows.Next() {
		var status Status
		rows.Scan(&status.ID, &status.Nombre)
		statuses = append(statuses, status)
	}
	return statuses
}

//Save method
func (s *Status) Save() error {
	if s.ID == 0 {
		return s.insert()
	}
	return s.update()
}

func (s *Status) insert() error {
	sql := "INSERT CTL_STATUS SET nombre=?"
	id, err := InsertData(sql, s.Nombre)
	s.ID = id
	return err
}

func (s *Status) update() error {
	sql := "UPDATE CTL_STATUS SET nombre=? WHERE id=?"
	_, err := Exec(sql, s.Nombre, s.ID)
	return err
}

//Delete method
func (s *Status) Delete() {
	sql := "DELETE FROM CTL_STATUS WHERE id=?"
	Exec(sql, s.ID)
}
