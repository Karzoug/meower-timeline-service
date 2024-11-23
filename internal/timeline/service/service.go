package service

import "github.com/rs/zerolog"

type TimelineService struct {
	logger zerolog.Logger
}

func NewTimelineService(logger zerolog.Logger) TimelineService {
	logger = logger.With().
		Str("component", "user service").
		Logger()

	return TimelineService{
		logger: logger,
	}
}
