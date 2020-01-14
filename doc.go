/*

Example of GET weather request:

http://localhost:50300/weather?city=paris&city=warszawa&city=london

Files description

main.go - entry file with http server initialization
api.go - api endpoints and http server
cache.go - cache for requests
config.go - reading config file
config.json - configuration file
contracts.go -  business logic models
service.go - buisness logic
openweather.go - external open weather endpoint service

*/

package main
