build:
	go build ./main.go
	go build ./internal/controllers/EventsControllers.go
    go build ./internal/cfg/Config.go
   	go build ./internal/db/EventsQueries.go
   	go build ./internal/models/Event.go
   	go build ./internal/service/eventsService.go

