module github.com/rjsikarwar/microservices-docker-go-mongodb/website

go 1.16

require (
	github.com/gorilla/mux v1.8.0
	github.com/rjsikarwar/microservices-docker-go-mongodb/bookings v0.0.0-00010101000000-000000000000
	github.com/rjsikarwar/microservices-docker-go-mongodb/showtimes v0.0.0-00010101000000-000000000000
	github.com/rjsikarwar/microservices-docker-go-mongodb/users v0.0.0-00010101000000-000000000000
)

replace github.com/rjsikarwar/microservices-docker-go-mongodb/movies => ../movies

replace github.com/rjsikarwar/microservices-docker-go-mongodb/bookings => ../bookings

replace github.com/rjsikarwar/microservices-docker-go-mongodb/showtimes => ../showtimes

replace github.com/rjsikarwar/microservices-docker-go-mongodb/users => ../users
