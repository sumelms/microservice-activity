package domain

import "github.com/google/uuid"

type ActivityRepository interface {
	Activity(activity_uuid uuid.UUID) (Activity, error)
	Activities() ([]Activity, error)
	CreateActivity(activity *Activity) error
	UpdateActivity(activity *Activity) error
	DeleteActivity(activity_uuid uuid.UUID) error
}
