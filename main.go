package main

import (
	attendanceController "first-app-golang/controllers"
	"fmt"

	// "net"
	"net/http"
)

func main() {
	http.HandleFunc("/", attendanceController.Index)
	http.HandleFunc("/attendance", attendanceController.Index)
	http.HandleFunc("/attendance/index", attendanceController.Index)

	http.HandleFunc("/attendance/add", attendanceController.Add)
	http.HandleFunc("/attendance/edit", attendanceController.Edit)
	http.HandleFunc("/attendance/delete", attendanceController.Delete)

	// fmt.Println("listening on localhost:8080...")
	http.ListenAndServe(":3000", nil)

	fmt.Println("listening on your ip address")
	// listen, _ := net.Listen("tcp", "192.168.1.30")
	// http.Serve(listen, nil)
}
