# Mattertime

Just a small plugin to

 - Record time spent by date/task/user from a mattermost hook (/spent e.g)
 - Learn Golang through writing a basic API with:
	 - Routing
	 - Logging
	 - Persistence
	 - Error handling

## Stack

 - Go 1.10
 - MongoDB 3.6.3
 - Gorilla Mux (router)

## Usage && dependencies

- files can be placed in an `/src/mattertime` directory
- `go get github.com/gorilla/mux` is the router
- `go get gopkg.in/mgo.v2` mongodb go driver
- `cd /src/mattertime && go install` will build the binary and place it in `/bin`
- `mattertime` runs it
- (i might have forgotten some pckgs...?)

## Todo

- ~~plug to mattermost of course~~
- ~~handle users (they should see only their documents)~~
- export env config (db name, collection name)
- secure mongodb with auth
- CI with build commands and dep installs
- secure routes with jwt
- Unit tests (oh really ?)

## Improvments

- dockerization ?
- Add js front-end with hooks to other platforms, for fun
- improve queries
- improve error handling
- better logging with levels, payloads and so on
- Db instance should be pass with dependency injection
- make a "commit" system to record and validate spent time ?

## Plugged commands: 

- `/mtm ls [yyyy-mm-dd]`
	all times recorded for today or specific day

- `/mtm add <duration> <task> [yyyy-mm-dd]`
	add new time

- `/mtm rm <ID>`
	remove specific time spent

- `/mtm clear`
	clear all times of the day

- `/mtm help`
	display table with commands, arguments, payloads and comments

## Commands to plug: 

- `/mtm tasks`
	all tasks recorded

- `/mtm stats`
	sends some stats

- `/mtm start <task>`
	start timer on this task

- `/mtm stop`
	stop timer on all tasks

- `/mtm poke <message> <time>`
	create a reminder that will pop in the channel with [message] in [time] minutes/hours

## Useful

- net/http/httputil to pretty dump requests


## See in action

![Alt Text](/demo_mtm3.gif)