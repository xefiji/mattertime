package main

import (
"net/http"
"encoding/json"
"errors"
"net/url"
"fmt"
"net/http/httputil"
"regexp"
"strings"
"strconv"
"time"
)

/*
	TODO: 
	- Regex: be less strict in task names: allow special chars
	- Regex: allow no minutes or no hours
*/

type Command struct {
	Action string
	Duration string
	Task string
	Date string
	Id string
	Message string
	User string	
}

type MattermostRet struct {
	ResponseType string `json:"response_type"`
	Text string `json:"text"`
}

//r.ParseForm() seems to be best to parse Content-Type: application/x-www-form-urlencoded
func MattermostMain(w http.ResponseWriter, r *http.Request) {
	
	// RequestDebug(r)

	if err := r.ParseForm(); err != nil {
		panic(err)
	}

	err, commande := ParseRequestToCommand(r.PostForm)
	if err != nil{
		fmt.Println(err)
	}

	
	// fmt.Printf("Commande:	%+v\n", commande)
	// fmt.Printf("TimeSpent:	%+v\n", timeSpent)

	var response MattermostRet
	switch commande.Action{

	case "add":
		response = Add(commande)

	case "ls":
		response = List(commande)

	case "rm":
		response = Remove(commande)

	case "clear":
		response = Clear(commande)
	}

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")	
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil{
		panic(err)
	}

}

func Clear(commande Command) MattermostRet {
	date := time.Now().Format("2006-01-02")
	err := ClearUserTimeSpentByDay(date, commande.User)
	if err != nil {
		panic(err)
	}

	response := FindTimeSpentFormatted(commande.User)
	return response
}

func Remove(commande Command) MattermostRet {
	err, _ := DestroyTimeSpent(commande.Id)
	if err != nil {
		panic(err)
	}

	response := FindTimeSpentFormatted(commande.User)
	return response
}

func List(commande Command) MattermostRet {
	response := FindTimeSpentFormatted(commande.User)
	return response
}

func Add(commande Command) MattermostRet {
	timeSpent := Convert(commande)

	err, _ := CreateTimeSpent(timeSpent)
	if err != nil {
		panic(err)
	}
	
	response := FindTimeSpentFormatted(commande.User)
	return response
}

func FindTimeSpentFormatted(user string) MattermostRet{
	
	date := time.Now().Format("2006-01-02") // today default
	_, t := FindTimeSpentByDateAndUser(date, user)

	var text string
	text += "#### Temps saisis par " + user + " pour la journée du " + date + "\n"
	text += "|ID|Date|Temps|Tâche|Crée le|\n"
	text += "|:-|:-|:-|:-|:-|:-|\n"

	for _, t := range t {
		created_at := t.CreatedAt.Format("02-01-2006 à 15:04")
		text += "|" + t.ID.Hex() + "|" + t.Date + "|" + strconv.FormatFloat(t.Spent, 'E', 2, 64) + "|" + t.Task + "|" + created_at + "|\n"
	}

	response := MattermostRet{ ResponseType: "in_channel", Text: text }

	return response

}

//all tests on strings are: https://regex101.com/r/QOcVZ7/4
func ParseRequestToCommand(requestDatas url.Values) (error, Command) {

	c := Command{}

	//regex powaaa
	pattern := `(?im)(?P<action>^\s*[a-z]+\t*)(?:(?P<id>\s+[\d\w]{24})|(?P<date>\s+[\d]{4}-[\d]{2}-[\d]{2})?|(?P<duration>\s+[0-9]{1,2}h[0-9]{1,2}m)?|(?P<task>\s+"[a-z0-9\s]+"))*(?:(?P<task2>\s+"[a-z0-9\s]+")(?P<date2>\s+[\d]{4}-[\d]{2}-[\d]{2}))?`
	var re = regexp.MustCompile(pattern)
	found := re.FindStringSubmatch(requestDatas.Get("text"))

	if len(found) == 0{
		return errors.New("Action vide ou requête invalide : " + requestDatas.Get("text")), c
	}

	//clean space from each captured string
	for i, f := range found {
		found[i] = strings.Trim(f, " ")
	}

	// fmt.Printf("PostForm:	%q\n", requestDatas.Get("text"))
	// for index, element := range found{
	// 	fmt.Printf("%v => %v\n", index, element)
	// }

	//fetch
	commandVars := map[string]int{}
	for k, v := range re.SubexpNames(){
		commandVars[v] = k
	}		

	// fmt.Printf("Found:	%v\n", found)
	// fmt.Printf("Subexp:	%v\n", re.SubexpNames())
	// fmt.Printf("CommandVars:	%v\n", commandVars)

	c.User = requestDatas.Get("user_name")
	c.Action = strings.ToLower(found[commandVars["action"]])
	switch c.Action {
	case "add":
		c.Task = found[commandVars["task"]]
		c.Duration = found[commandVars["duration"]]		
		c.Date = found[commandVars["date"]]
	case "ls":
		c.Date = found[commandVars["date"]]
	case "rm":
		c.Id = found[commandVars["id"]]
	case "start":
		c.Task = found[commandVars["task"]]
	case "poke":
		c.Message = ""
		c.Duration = ""
	case "tasks", "clear", "stats", "help", "stop":
		break
	}	

	return nil, c
}

func Convert(c Command) TimeSpent {

	t := TimeSpent{
		Date: c.Date,
		Spent: ConvertDuration(c.Duration),
		Task: strings.Replace(c.Task, "\"", "", -1),
		User: c.User,
	}

	return t
}

func ConvertDuration(s string) float64 {
	
	var re = regexp.MustCompile(`(?i)(?P<hour>[0-9]{1,2})h(?P<minute>[0-9]{1,2})m`)
	found := re.FindStringSubmatch(s)

	//clean leading zeros
	for i, f := range found {
		if i == 0 {
			continue
		}
		
		if f[:1] == "0"{
			found[i] = f[1:]
		}
	}

	//set parts
	parts := map[string]int{}
	for k, v := range re.SubexpNames(){
		parts[v] = k
	}

	hour, _ := strconv.ParseFloat(found[parts["hour"]], 64)
	minute, _ := strconv.ParseFloat(found[parts["minute"]], 64)
	var hourcents float64
	hourcents = ((hour * 60) + minute) / 60
	
	return hourcents
}

//should be moved somewhere more global
func RequestDebug(r *http.Request) {
	fmt.Printf("\n\n--------------------------\n\n")	
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(requestDump))
	fmt.Printf("\n\n--------------------------\n\n")	
}

