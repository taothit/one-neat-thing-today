package discover

import (
	"context"
	"log"

	neatthing "github.com/taothit/one-neat-thing-today/gen/neat_thing"
)

// neatThing service example implementation.
// The example methods log the requests and return zero values.
type neatThingsrvc struct {
	logger *log.Logger
}

// NewNeatThing returns the neatThing service implementation.
func NewNeatThing(logger *log.Logger) neatthing.Service {
	return &neatThingsrvc{logger}
}

// NeatThingToday implements neatThingToday.
func (s *neatThingsrvc) NeatThingToday(ctx context.Context) (res *neatthing.NeatThing, err error) {
	res = &neatthing.NeatThing{}
	s.logger.Print("neatThing.neatThingToday")
	return
}
