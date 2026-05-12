package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"slices"

	"github.com/google/uuid"
)

type Fighter struct {
	Id       string `json:"id,omitempty"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Nickname string `json:"nickname"`
}

type Fighters struct {
	Farray []Fighter
}

func (fighter *Fighter) updateFields(reqBody io.ReadCloser) bool {
	var newData map[string]interface{}

	err := json.NewDecoder(reqBody).Decode(&newData)
	if err != nil {
		return false
	}

	if v, ok := newData["name"]; ok {
		fighter.Name = v.(string)
	}

	if v, ok := newData["age"]; ok {
		fighter.Age = int(v.(float64))
	}

	if v, ok := newData["nickname"]; ok {
		fighter.Nickname = v.(string)
	}

	return true
}

func getFromFile() (*Fighters, error) {
	file, err := os.ReadFile("fighters.json")

	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return nil, err

	}

	var fighters Fighters
	if errors.Is(err, os.ErrNotExist) {
		return &fighters, nil
	}

	errFighter := json.Unmarshal(file, &fighters)

	if errFighter != nil {
		return nil, errFighter
	}

	return &fighters, nil
}

func writeToFile(fighters Fighters) bool {
	fileData, err := json.MarshalIndent(fighters, "", " ")

	if err != nil {
		return false
	}

	err = os.WriteFile("fighters.json", fileData, 0644)

	if err != nil {
		return false
	}

	return true
}

func getFigthers(w http.ResponseWriter, r *http.Request) {

	fighters, err := getFromFile()

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Fprint(w, "List is empty for now")
			return
		}
		fmt.Fprint(w, "Something went wrong getting fighters")
		return
	}
	jsonData, err := json.Marshal(fighters.Farray)

	if err != nil {
		fmt.Fprintf(w, "Can't be JSONed, here is the struct : %+v", fighters)
		return
	}
	fmt.Fprint(w, string(jsonData))

}

func getFighterById(w http.ResponseWriter, r *http.Request) {
	fighters, err := getFromFile()

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Fprint(w, "List is empty for now")
			return
		}
		fmt.Fprint(w, "Something went wrong getting fighters")
	}

	// id := r.URL.Query().Get("id")
	id := r.PathValue("id")

	if id == "" {
		http.Error(w, "Id is required path param", 400)
		return
	}

	for _, v := range fighters.Farray {
		if v.Id == id {
			jsonData, _ := json.Marshal(v)
			fmt.Fprint(w, string(jsonData))
			return
		}
	}

	fmt.Fprintf(w, "There is no such a fighter with id: %s", id)
}

func createFighter(w http.ResponseWriter, r *http.Request) {
	fighters, err := getFromFile()

	if err != nil {
		fmt.Fprint(w, "Something went wrong")
		return
	}

	var newFighter Fighter

	err = json.NewDecoder(r.Body).Decode(&newFighter)

	newFighter.Id = uuid.New().String()

	if err != nil {
		fmt.Fprint(w, "invalid keys or values")
		return
	}

	fighters.Farray = append(fighters.Farray, newFighter)

	done := writeToFile(*fighters)

	if !done {
		fmt.Fprint(w, "something went wrong while writing data to file")
		return
	}

	fmt.Fprintf(w, "fighter added id: %s", newFighter.Id)

}

func updateFighter(w http.ResponseWriter, r *http.Request) {
	fighters, err := getFromFile()

	if err != nil {
		fmt.Fprint(w, "Something went wrong")
	}

	id := r.PathValue("id")

	if id == "" {
		http.Error(w, "Id is required", 400)
		return
	}

	done := false
	for i := range fighters.Farray {
		if fighters.Farray[i].Id == id {
			fighters.Farray[i].updateFields(r.Body)
			done = true
			break
		}
	}

	if done {
		done = writeToFile(*fighters)
		if !done {
			fmt.Fprint(w, "Something went wrong writing data to file")
			return
		}
		fmt.Fprintf(w, "Fighter successfully updated %s", id)
		return
	}

	fmt.Fprintf(w, "There is no such a figher with id: %s", id)
}

func deleteFighter(w http.ResponseWriter, r *http.Request) {
	fighters, err := getFromFile()

	if err != nil {
		fmt.Fprint(w, "Something went wrong")
	}

	id := r.PathValue("id")

	ind := -1

	for i := range fighters.Farray {
		if fighters.Farray[i].Id == id {
			ind = i
			break
		}
	}

	if ind == -1 {
		fmt.Fprint(w, "There is no such fighter with this id")
		return
	}

	fighters.Farray = slices.Delete(fighters.Farray, ind, ind+1)

	done := writeToFile(*fighters)

	if !done {
		fmt.Fprint(w, "Something went wrong writing data to file")
	}

	fmt.Fprintf(w, "Fighter with id: %s successfully deleted", id)

}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /fighters", getFigthers)
	mux.HandleFunc("GET /fighters/{id}", getFighterById)
	mux.HandleFunc("POST /fighters", createFighter)
	mux.HandleFunc("PATCH /fighters/{id}", updateFighter)
	mux.HandleFunc("PUT /fighters/{id}", updateFighter)
	mux.HandleFunc("DELETE /fighters/{id}", deleteFighter)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println("Server didn't start")
	}
}
