package personsController

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"go-api/app/types"
	"go-api/app/util"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(util.Persons)
}

func GetPersonByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	person := util.Person{}

	for _, currentPerson := range util.Persons {
		if currentPerson.Id == id {
			person = currentPerson
		}
	}

	json.NewEncoder(w).Encode(person)
}

func Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var person util.Person

	if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if person.Avatar == "" {
		http.Error(w, "Avatar is a required field", http.StatusBadRequest)
		return
	}

	if person.Firstname == "" {
		http.Error(w, "Firstname is a required field", http.StatusBadRequest)
		return
	}

	if person.Lastname == "" {
		http.Error(w, "Lastname is a required field", http.StatusBadRequest)
		return
	}

	if person.Email == "" {
		http.Error(w, "Email is a required field", http.StatusBadRequest)
		return
	}

	if !util.IsValidEmail(person.Email) {
		http.Error(w, "Email is invalid", http.StatusBadRequest)
		return
	}

	if person.Age <= 0 || person.Age > 120 {
		http.Error(w, "The Age field must be greater than 0 and less than or equal to 120", http.StatusBadRequest)
		return
	}

	person.Id = util.GenerateRandomID()
	util.Persons = append(util.Persons, person)

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, json.NewEncoder(w).Encode(types.SuccessMessage{Message: "Stored successfully!"}))
}

func Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	vars := mux.Vars(r)
	id := vars["id"]

	filtedPersons := []util.Person{}

	for _, currentPerson := range util.Persons {
		if currentPerson.Id != id {
			filtedPersons = append(filtedPersons, currentPerson)
		}
	}

	util.Persons = filtedPersons
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, json.NewEncoder(w).Encode(types.SuccessMessage{Message: "User deleted successfully!"}))
}

func Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	vars := mux.Vars(r)
	id := vars["id"]

	var newDataPerson util.Person

	if err := json.NewDecoder(r.Body).Decode(&newDataPerson); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	indexToUpdate := -1

	for index, currentPerson := range util.Persons {
		if currentPerson.Id == id {
			indexToUpdate = index
			break
		}
	}

	if indexToUpdate == -1 {
		http.Error(w, "Person not exist", http.StatusBadRequest)
		return
	}

	if newDataPerson.Avatar == "" {
		http.Error(w, "Avatar is a required field", http.StatusBadRequest)
		return
	}

	if newDataPerson.Firstname == "" {
		http.Error(w, "Firstname is a required field", http.StatusBadRequest)
		return
	}

	if newDataPerson.Lastname == "" {
		http.Error(w, "Lastname is a required field", http.StatusBadRequest)
		return
	}

	if newDataPerson.Email == "" {
		http.Error(w, "Email is a required field", http.StatusBadRequest)
		return
	}

	if !util.IsValidEmail(newDataPerson.Email) {
		http.Error(w, "Email is invalid", http.StatusBadRequest)
		return
	}

	if newDataPerson.Age <= 0 || newDataPerson.Age > 120 {
		http.Error(w, "The Age field must be greater than 0 and less than or equal to 120", http.StatusBadRequest)
		return
	}

	newDataPerson.Id = util.Persons[indexToUpdate].Id
	util.Persons[indexToUpdate] = newDataPerson
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, json.NewEncoder(w).Encode(types.SuccessMessage{Message: "User updated successfully!"}))
}
