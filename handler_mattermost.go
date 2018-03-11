package main

import (
"net/http"
"fmt"
"net/http/httputil"
)

type Command struct {
	Action string
	Duration string
	Date string
	User string
	CreatedAt string
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

	if err := r.ParseForm(); err != nil {
		msg := fmt.Sprintf("Unable to parse form :: %v", err.Error())
		fmt.Printf(msg)
	}

	c := Command{"nil", "nil", "date concernée", r.PostForm["user_name"][0], "NOW"}
	final := ParseAction(r.PostForm["text"][0], c)
	fmt.Println(final)
}

func ParseAction(datas string, command Command) Command{
	command.Action = "Add"
	return command
}
