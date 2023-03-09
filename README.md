# shortener service #
 - - - -
 # contents #
 * #### [project definition](#about-the-project)
 * #### [run without docker](#run-project-without-docker)
 * #### [run project through docker](#run-project-through-docker)
## About the project ##
> A URL shortener service is a tool that takes a long, unwieldy web address and creates a much shorter, more manageable version of it. This shorter version of the URL is easier to share, particularly on social media sites like Twitter where character limits are in place.
- - - -
## run project without docker ##
if you want to run the project without using docker:
### requirements ###
* __Angular CLI: 12.2.16__  
* __go 1.16.6__
* __npm 8.12.1 &rarr; node v16.16.0__
### install dependencies ###
1. go to ___back-end___ directory and run
```
go get -u -v all
```
the output must be something as below
```
go: downloading github.com/golang-jwt/jwt/v4 v4.5.0
go: downloading github.com/labstack/gommon v0.4.0
go: downloading golang.org/x/crypto v0.7.0
go: downloading github.com/mattn/go-colorable v0.1.13
go: downloading github.com/valyala/fasttemplate v1.2.2
go: downloading github.com/mattn/go-isatty v0.0.17
...
```
2. go to ___front-end___ directory and run
```
npm install
```
### launch ###

1. inside ___back-end___ directory run
```
go run main.go
```
2. inside ___front-end___ directory run
```
ng serve
```
now you can use the web app. start your journey by sending request to __http://localhost:4200/__  
good luck :)
- - - -
## use project through docker ##
this section is in progress.  
stay tuned ;)
