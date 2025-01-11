package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-clean-architecture-sample/api/models"
	"go-clean-architecture-sample/application/common/errors"
	"go-clean-architecture-sample/application/common/interfaces"
	"go-clean-architecture-sample/application/employee/commands"
	"go-clean-architecture-sample/application/employee/query"
	"io/ioutil"
	"net/http"
	"strconv"
)

type EmployeeController struct {
	employeeRepository interfaces.EmployeeRepository
	avatarProvider     interfaces.AvatarProvider
}

func NewEmployeeController(
	employeeRepository interfaces.EmployeeRepository,
	avatarProvider interfaces.AvatarProvider,
) *EmployeeController {
	return &EmployeeController{
		employeeRepository: employeeRepository,
		avatarProvider:     avatarProvider,
	}
}

func (c *EmployeeController) AddEmployee(httpResponseWriter http.ResponseWriter, httpRequest *http.Request) {
	body, err := ioutil.ReadAll(httpRequest.Body)
	if err != nil {
		writeJsonHttpResponse(httpResponseWriter, http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}

	var requestObj models.AddEmployeeRequestModel
	err = json.Unmarshal(body, &requestObj)
	if err != nil {
		writeJsonHttpResponse(httpResponseWriter, http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}

	addEmployeeHandler := commands.NewAddEmployeeCommandHandler(
		c.employeeRepository,
		c.avatarProvider,
	)

	createdEmployee, err := addEmployeeHandler.Handle(commands.AddEmployeeCommand{
		Name:  requestObj.Name,
		Email: requestObj.Email,
	})
	if err != nil {
		writeJsonHttpResponse(httpResponseWriter, http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}

	responseObj := models.EmployeeResponseModel{
		Id:        createdEmployee.Id,
		Name:      createdEmployee.Name,
		Email:     createdEmployee.Email,
		AvatarUrl: createdEmployee.AvatarUrl,
	}

	writeJsonHttpResponse(httpResponseWriter, http.StatusCreated, responseObj)
}

func (c *EmployeeController) GetAllEmployees(httpResponseWriter http.ResponseWriter, _ *http.Request) {
	getAllEmployeesHandler := query.NewGetAllEmployeesQueryHandler(c.employeeRepository, c.avatarProvider)

	employees, err := getAllEmployeesHandler.Handle(query.GetAllEmployeesQuery{})
	if err != nil {
		writeJsonHttpResponse(httpResponseWriter, http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}

	responseObj := make([]models.EmployeeResponseModel, 0)
	for _, employee := range employees {
		responseObj = append(responseObj, models.EmployeeResponseModel{
			Id:        employee.Id,
			Name:      employee.Name,
			Email:     employee.Email,
			AvatarUrl: employee.AvatarUrl,
		})
	}

	writeJsonHttpResponse(httpResponseWriter, http.StatusOK, responseObj)
}

func (c *EmployeeController) GetEmployeeById(httpResponseWriter http.ResponseWriter, httpRequest *http.Request) {
	strEmployeeId, employeeIdSpecified := mux.Vars(httpRequest)["employee_id"]
	if !employeeIdSpecified {
		writeJsonHttpResponse(httpResponseWriter, http.StatusBadRequest, models.ErrorResponse{Message: "employee_id not specified"})
		return
	}

	employeeId, err := strconv.Atoi(strEmployeeId)
	if err != nil {
		writeJsonHttpResponse(httpResponseWriter, http.StatusBadRequest, models.ErrorResponse{Message: "invalid employee_id format"})
		return
	}

	getEmployeeByIdHandler := query.NewGetEmployeeByIdQueryHandler(c.employeeRepository, c.avatarProvider)

	employee, err := getEmployeeByIdHandler.Handle(query.GetEmployeeByIdQuery{EmployeeId: employeeId})
	if err != nil {
		var statusCode int
		if err == errors.ErrNotFound {
			statusCode = http.StatusNotFound
		} else {
			statusCode = http.StatusInternalServerError
		}
		writeJsonHttpResponse(httpResponseWriter, statusCode, models.ErrorResponse{Message: err.Error()})
		return
	}

	responseObj := models.EmployeeResponseModel{
		Id:        employee.Id,
		Name:      employee.Name,
		Email:     employee.Email,
		AvatarUrl: employee.AvatarUrl,
	}

	writeJsonHttpResponse(httpResponseWriter, http.StatusOK, responseObj)
}

func (c *EmployeeController) DeleteEmployeeById(httpResponseWriter http.ResponseWriter, httpRequest *http.Request) {
	strEmployeeId, employeeIdSpecified := mux.Vars(httpRequest)["employee_id"]
	if !employeeIdSpecified {
		writeJsonHttpResponse(httpResponseWriter, http.StatusBadRequest, models.ErrorResponse{Message: "employee_id not specified"})
		return
	}

	employeeId, err := strconv.Atoi(strEmployeeId)
	if err != nil {
		writeJsonHttpResponse(httpResponseWriter, http.StatusBadRequest, models.ErrorResponse{Message: "invalid employee_id format"})
		return
	}

	deleteEmployeeByIdHandler := commands.NewDeleteEmployeeByIdCommandHandler(c.employeeRepository)

	err = deleteEmployeeByIdHandler.Handle(commands.DeleteEmployeeByIdCommand{EmployeeId: employeeId})
	if err != nil {
		var statusCode int
		if err == errors.ErrNotFound {
			statusCode = http.StatusNotFound
		} else {
			statusCode = http.StatusInternalServerError
		}
		writeJsonHttpResponse(httpResponseWriter, statusCode, models.ErrorResponse{Message: err.Error()})
		return
	}

	httpResponseWriter.WriteHeader(http.StatusNoContent)
}

func writeJsonHttpResponse(httpResponseWriter http.ResponseWriter, statusCode int, responseObj interface{}) {
	responseJson, _ := json.Marshal(responseObj)
	httpResponseWriter.Header().Set("Content-Type", "application/json")
	httpResponseWriter.WriteHeader(statusCode)
	httpResponseWriter.Write(responseJson)
}
