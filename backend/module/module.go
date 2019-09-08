package module

import "github.com/simple-event-app/backend/models/repository"

// Configure Inject
var Configure = struct {
	Repository struct {
		EventRepository repository.EventRepository
		TagRepository   repository.TagRepository
	}
}{}
