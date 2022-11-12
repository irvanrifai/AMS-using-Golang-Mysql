package models

import (
	"database/sql"
	"first-app-golang/config"
	"first-app-golang/entities"
	"fmt"
)

type EmployeeModel struct {
	conn *sql.DB
}

func NewEmployeeModel() *EmployeeModel {
	conn, err := config.DBconnection()
	if err != nil {
		panic(err)
	}
	return &EmployeeModel{
		conn: conn,
	}
}

func (e *EmployeeModel) GetAll() ([]entities.Employee, error) {
	rows, err := e.conn.Query("select * from employee order by id desc")
	if err != nil {
		return []entities.Employee{}, err
	}
	defer rows.Close()

	var dataEmployee []entities.Employee
	for rows.Next() {
		var employee entities.Employee
		rows.Scan(&employee.Id, &employee.Name, &employee.Address, &employee.Position, &employee.Company, &employee.PresenceIn, &employee.PresenceOut, &employee.Gender, &employee.EmployeeNumber)

		if employee.Gender == "1" {
			employee.Gender = "Male"
		} else if employee.Gender == "2" {
			employee.Gender = "Female"
		}

		dataEmployee = append(dataEmployee, employee)
	}

	return dataEmployee, nil
}

func (e *EmployeeModel) Create(employee entities.Employee) bool {
	result, err := e.conn.Exec("insert into employee (name, address, position, company, presence_in, presence_out, gender, employee_number) values(?,?,?,?,?,?,?,?)",
		employee.Name, employee.Address, employee.Position, employee.Company, employee.PresenceIn, employee.PresenceOut, employee.Gender, employee.EmployeeNumber)

	if err != nil {
		fmt.Println(err)
		return false
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId > 0
}

func (e* EmployeeModel) Find(id int64, employee* entities.Employee) error {
    return e.conn.QueryRow("select * from employee where id = ?", id).Scan(&employee.Id, &employee.Name, &employee.Address, &employee.Position, &employee.Company, &employee.PresenceIn, &employee.PresenceOut, &employee.Gender, &employee.EmployeeNumber)
}

func (e* EmployeeModel) Update(employee entities.Employee) error {
    _, err := e.conn.Exec(
        "update employee set name = ?, address = ?, position = ?, company = ?, presence_in = ?, presence_out = ?, gender = ?, employee_number = ? where id = ?",
        employee.Name, employee.Address, employee.Position, employee.Company, employee.PresenceIn, employee.PresenceOut, employee.Gender, employee.EmployeeNumber, employee.Id)
    if err != nil {
        return err
    }
    return nil
}

func (e* EmployeeModel) Delete(id int64) {
    e.conn.Exec("delete from employee where id = ?", id)
}
