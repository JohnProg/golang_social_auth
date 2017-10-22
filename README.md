# golang_social_auth
The golang application with, with ability to 
    
    - signup
    - login 
    - updated profile
    - reset password by email
    - reset password by token
    
## Getting started
    Download go src from https://golang.org/dl/
    Download IDE editor from https://www.jetbrains.com/go/?fromMenu
    Download curl from https://curl.haxx.se/
    Install mysql db from https://www.mysql.com/
    Install mysql client from https://www.sequelpro.com/ for OSX 
    Instal iTerm from https://www.iterm2.com/ for OSX

## How to run app
    
    `go run main.go`
    
## How to run integration-tests

    `make integration-tests` 

## Testing application

    `curl -i -H "Content-Type: application/json" -X POST -d '{"type":"email","value":"siemenspromaster@gmail.com","password":"xyz"}' http://localhost:8080/signup`

    `curl -i -H "Content-Type: application/json" -X POST -d '{"type":"linkedin","value":"i_esVlB5FI"}' http://localhost:8080/signup`

    `curl -i -H "Content-Type: application/json" -X POST -d '{"type":"email","value":"siemenspromaster@gmail.com","password":"xyz"}' http://localhost:8080/login`

    `curl -i -H "Content-Type: application/json" -X POST -d '{"type":"linkedin","value":"i_esVlB5FI"}' http://localhost:8080/login`

    `curl -i -H "Authorization: token 060de32a-c3bb-497b-acf5-8910d6925c43" -H "Content-Type: application/json" -X PUT -d '{"first_name":"ievgen","last_name":"maksymenko","username":"skews"}' http://localhost:8080/user/me`

    `curl -i -H "Content-Type: application/json" -X POST -d '{"email":"siemenspromaster@gmail.com"}' http://localhost:8080/password_reset`

    `curl -i -H "Content-Type: application/json" -X POST -d '{"password":"xyz"}' http://localhost:8080/e96de8df-375b-4237-94b8-7320818ae388/password_reset`
