package attendanceController

import (
	"first-app-golang/entities"
	"first-app-golang/libraries"
	"first-app-golang/models"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

var validation = libraries.NewValidation()
var employeeModel = models.NewEmployeeModel()

func Index(response http.ResponseWriter, request *http.Request) {

	employee, _ := employeeModel.GetAll()

	data := map[string]interface{}{
		"employee": employee,
	}

	temp, err := template.ParseFiles("views/attendance/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(response, data)
}

func Add(response http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		temp, err := template.ParseFiles("views/attendance/add.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, nil)
	} else if request.Method == http.MethodPost {
		request.ParseForm()
		var employee entities.Employee

		employee.Name = request.Form.Get("name")
		employee.Address = request.Form.Get("address")
		employee.Position = request.Form.Get("position")
		employee.Company = request.Form.Get("company")
		employee.PresenceIn = request.Form.Get("presence_in")
		employee.PresenceOut = request.Form.Get("presence_out")
		employee.Gender = request.Form.Get("gender")
		employee.EmployeeNumber = request.Form.Get("employee_number")

		data := make(map[string]interface{})
		datasemployee, _ := employeeModel.GetAll()

		vErrors := validation.Struct(employee)

		if vErrors != nil {
			data["validation"] = vErrors
			data["employee"] = employee
			temp, _ := template.ParseFiles("views/attendance/add.html")
			temp.Execute(response, data)
			fmt.Println(vErrors)
		} else {
			employeeModel.Create(employee)
			data["message"] = "Data success added"
			data["employee"] = datasemployee

            // with message
			// temp, _ := template.ParseFiles("views/attendance/index.html")
			// temp.Execute(response, data)
            
            // without message
            http.Redirect(response, request, "/", http.StatusSeeOther)
            
            fmt.Println(employee)
		}

		// data := map[string]interface{}{
		//     "message": "Data success added",
		//     "employee" : datasemployee,
		// }

	}
}

func Edit(response http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {

		queryString := request.URL.Query()
		id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

		var employee entities.Employee
		employeeModel.Find(id, &employee)

		data := map[string]interface{}{
			"employee": employee,
		}

		fmt.Println(employee)

		temp, err := template.ParseFiles("views/attendance/edit.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, data)
	} else if request.Method == http.MethodPost {
		request.ParseForm()
		var employee entities.Employee

		employee.Id, _ = strconv.ParseInt(request.Form.Get("id"), 10, 64)
		employee.Name = request.Form.Get("name")
		employee.Address = request.Form.Get("address")
		employee.Position = request.Form.Get("position")
		employee.Company = request.Form.Get("company")
		employee.PresenceIn = request.Form.Get("presence_in")
		employee.PresenceOut = request.Form.Get("presence_out")
		employee.Gender = request.Form.Get("gender")
		employee.EmployeeNumber = request.Form.Get("employee_number")

		data := make(map[string]interface{})

		vErrors := validation.Struct(employee)

		if vErrors != nil {
			data["validation"] = vErrors
			data["employee"] = employee
			temp, _ := template.ParseFiles("views/attendance/edit.html")
			temp.Execute(response, data)
			fmt.Println(vErrors)
		} else {
            employeeModel.Update(employee)
			data["message"] = "Data success updated"
            
			datasemployee, _ := employeeModel.GetAll()
			data["employee"] = datasemployee

            // with message
			// temp, _ := template.ParseFiles("views/attendance/index.html")
			// temp.Execute(response, data)
            
            // without message
            http.Redirect(response, request, "/", http.StatusSeeOther)
          
            fmt.Println(employee)
		}
	}
}

func Delete(response http.ResponseWriter, request *http.Request) {
    queryString := request.URL.Query()
	id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

    employeeModel.Delete(id)

    http.Redirect(response, request, "/", http.StatusSeeOther)
}
