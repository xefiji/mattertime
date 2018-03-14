package main

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"time"
)

type Ret struct {
	Message string `json:"message,omitempty"`
	Datas TimeSpents `json:"datas"`
}

func GetTimeToday(w http.ResponseWriter, r *http.Request) {		
	
	//find todays spent time
	today := time.Now().Format("2006-01-02")
	_, t := FindTimeSpent(today)

 	//prepare for return
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")	
	w.WriteHeader(http.StatusOK)
	e := Ret{"", t}
	if err := json.NewEncoder(w).Encode(e); err != nil{
		panic(err)
	}
}

func GetTimeDay(w http.ResponseWriter, r *http.Request) {	
	//find specific date's spent time
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")	
	vars := mux.Vars(r)
	dateString := vars["date"] + "T00:00:00.000Z"

	d, err := time.Parse(time.RFC3339,dateString);
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		e := Ret{err.Error(), TimeSpents{}}
		if err := json.NewEncoder(w).Encode(e); err != nil{
			panic(err)
		}
		return
	}

	_, t := FindTimeSpent(d.Format("2006-01-02"))
	w.WriteHeader(http.StatusOK)
	e := Ret{"", t}
	if err := json.NewEncoder(w).Encode(e); err != nil{
		panic(err)
	}
}

//ioutil.ReadAll() seems to be best to parse Content-Type: application/json
func RecordTime(w http.ResponseWriter, r *http.Request) {	
	
	//some checks on request	
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil{
		panic(err)
	}
	if err := r.Body.Close(); err != nil{
		panic(err)
	}	

	//populate spent time struc with json datas
	var timeSpent TimeSpent	
	if err := json.Unmarshal(body, &timeSpent); err != nil{				
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(422) //why ??
		e := Ret{err.Error(), TimeSpents{}}
		if err := json.NewEncoder(w).Encode(e); err != nil{
			panic(err)
		}
		return
	}
	
	//Creation
	errorInsert, t := CreateTimeSpent(timeSpent)
	if errorInsert != nil{
		w.WriteHeader(422) //why ??
		e := Ret{errorInsert.Error(), TimeSpents{}}
		if err := json.NewEncoder(w).Encode(e); err != nil{
			panic(err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	e := Ret{"Spent time created", append(TimeSpents{}, t)}
	if err := json.NewEncoder(w).Encode(e); err != nil{
		panic(err)
	}
}

func DeleteTime(w http.ResponseWriter, r *http.Request) {			
	vars := mux.Vars(r)
	id := vars["id"]
	
	//Delete
	errorDelete, _ := DestroyTimeSpent(id)
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")

	if errorDelete != nil{
		w.WriteHeader(422) //why ??
		e := Ret{errorDelete.Error(), TimeSpents{}}
		if err := json.NewEncoder(w).Encode(e); err != nil{
			panic(err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	e := Ret{id + " has been deleted", TimeSpents{}}
	if err := json.NewEncoder(w).Encode(e); err != nil{
		panic(err)
	}
}
