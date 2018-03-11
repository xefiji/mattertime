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

## Todo

- export env config
- CI with build commands and dep installs
- plug to mattermost
- secure with jwt
- handle users

## Improvments
- Add js front-end with hooks to other platforms
- improve queries
- improve error handling
- better logging
- Db instance should be pass with dependency injection
- add a commit system to record and validate spent time

## COMMANDS ?

- /spent ls
	all times recorded for today

- /spent ls [date]
	all times recorded for specific day

- /spent ls [task]
	all times recorded for specific task

- /spent tasks
	all tasks recorded

- /spent add [datas]
	add new time

- /spent rm [ID]
	remove specific time spent
