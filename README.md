# Jedi Club - awesome Golang application

## What has been done:
* The application runs as a HTTP-server on your own port
* Jedi is described by his name, email, registration date and unique ID
* Jedi’s email is unique
* All data is stored in the application memory (no data base used)
* Golang has been used for the back-end /html+bootstrap for the front-end - no js was used for simplicity
* Back-end side form validation
* Front-end side form validation

## How to start:
* install git
* download the repository from github `git clone git@github.com:pmatiash/jediclub.git`
* install golang
* go to `/jediclub/`
* execute command `export PORT=%YOUR_FAVORITE_PORT_FOR_LOCAL_APP%` 
* execute command `go build`
* execute file `./jediclub` 
* open `localhost:%YOUR_FAVORITE_PORT_FOR_LOCAL_APP%` in your browser

## Dokerized image has been already deployed via Heroku:
https://frozen-temple-96054.herokuapp.com/

You may use following API calls as well:
* GET https://frozen-temple-96054.herokuapp.com/jedis - list of all available resources
* POST https://frozen-temple-96054.herokuapp.com/jedis - create new resource
