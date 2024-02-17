package domain

import "github.com/google/uuid"

type ActivityRepository interface {
	Activity(activityUUID uuid.UUID) (Activity, error)
	Activities() ([]Activity, error)
	CreateActivity(activity *Activity) error
	UpdateActivity(activity *Activity) error
	DeleteActivity(activityUUID uuid.UUID) error
}
