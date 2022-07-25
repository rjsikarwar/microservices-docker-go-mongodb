module github.com/mmorejon/microservices-docker-go-mongodb/website

go 1.16

require (
	github.com/gorilla/mux v1.8.0
	github.com/mmorejon/microservices-docker-go-mongodb/bookings v0.0.0-00010101000000-000000000000
	github.com/mmorejon/microservices-docker-go-mongodb/movies v0.0.0-00010101000000-000000000000
	github.com/mmorejon/microservices-docker-go-mongodb/showtimes v0.0.0-00010101000000-000000000000
	github.com/mmorejon/microservices-docker-go-mongodb/users v0.0.0-00010101000000-000000000000
)

replace github.com/mmorejon/microservices-docker-go-mongodb/movies => ../movies

replace github.com/mmorejon/microservices-docker-go-mongodb/bookings => ../bookings

replace github.com/mmorejon/microservices-docker-go-mongodb/showtimes => ../showtimes

replace github.com/mmorejon/microservices-docker-go-mongodb/users => ../users
