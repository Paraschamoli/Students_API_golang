package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/Paraschamoli/students_API/internal/storage"
	"github.com/Paraschamoli/students_API/internal/types"
	"github.com/Paraschamoli/students_API/internal/utils/response"
	"github.com/go-playground/validator/v10"
)

func New(storage storage.Storage) http.HandlerFunc {
return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("creating a student")

		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty body")))
			return
		}

		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		if err := validator.New().Struct(student); err != nil {

			 validateErrs := err.(validator.ValidationErrors)
			 response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validateErrs))
			 return
		}
slog.Info("creating new student", "student", student)

		// TODO: Save the student to database (e.g., insert to MongoDB)
		lastId,err:=storage.CreateStudent(
			student.Name,
			student.Email,
			student.Age,
		)
		if err!=nil {
	response.WriteJson(w,http.StatusInternalServerError,err)
	return 
}
		response.WriteJson(w, http.StatusCreated, map[string]string{
    "success": "student created",
    "id":      fmt.Sprintf("%d", lastId),
})

	}
}

func GetById(storage storage.Storage) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		id:=r.PathValue("id")
		slog.Info("getting a student",slog.String("id",id))

		intId,err:=strconv.ParseInt(id,10,64)
		if err!=nil{
			response.WriteJson(w,http.StatusBadRequest,response.GeneralError(err))
			return 
		}
		student,err:=storage.GetStudentById(intId)
		if err!=nil{
			response.WriteJson(w,http.StatusInternalServerError,response.GeneralError(err))
			return 
		}
		response.WriteJson(w,http.StatusOK,student)
	}
}


		
	