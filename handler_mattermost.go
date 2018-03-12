package main

import (
"net/http"
"fmt"
"net/http/httputil"
"time"
)

type Command struct {
	Action string
	Duration string
	Date string
	User string
	CreatedAt time.Time
}

type MattermostRet struct {
	ResponseType string `json:"response_type"`
	Text string `json:"text"`
}

func MattermostMain(w http.ResponseWriter, r *http.Request) {
	
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(requestDump))

	fmt.Printf("\n\n--------------------------\n\n")	

	ParseCommand(r)
}

//r.ParseForm() seems to be best to parse Content-Type: application/x-www-form-urlencoded
func ParseCommand(r *http.Request) Command{
	if err := r.ParseForm(); err != nil {
		panic(err)
	}

	requestDatas := r.PostForm

	c := Command{}
	c.User = requestDatas.Get("user_name")
	c.CreatedAt = time.Now()

	fmt.Printf("%+v\n", c)

	return c
}

func ParseAction(s string) string{
	return s
}

func ParseDuration(s string) string{
	return s
}

func ParseTask(s string) string{
	return s
}

func ParseDate(s string) string{
	return s
}


