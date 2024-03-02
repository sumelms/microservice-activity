package database

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/sumelms/microservice-activity/internal/activity/domain"
	"github.com/sumelms/microservice-activity/pkg/errors"
)

// NewActivityRepository creates the activity ActivityRepository.
func NewActivityRepository(db *sqlx.DB) (ActivityRepository, error) { //nolint: revive
	sqlStatements := make(map[string]*sqlx.Stmt)

	for queryName, query := range queriesActivity() {
		stmt, err := db.Preparex(query)
		if err != nil {
			return ActivityRepository{},
				errors.WrapErrorf(err, errors.ErrCodeUnknown, "error preparing statement %s", queryName)
		}
		sqlStatements[queryName] = stmt
	}

	return ActivityRepository{
		statements: sqlStatements,
	}, nil
}

type ActivityRepository struct {
	statements map[string]*sqlx.Stmt
}

func (r ActivityRepository) statement(s string) (*sqlx.Stmt, error) {
	stmt, ok := r.statements[s]
	if !ok {
		return nil, errors.NewErrorf(errors.ErrCodeUnknown, "prepared statement %s not found", s)
	}
	return stmt, nil
}

func (r ActivityRepository) Activity(activityUUID uuid.UUID) (domain.Activity, error) {
	stmt, err := r.statement(getActivity)
	if err != nil {
		return domain.Activity{}, err
	}

	var a domain.Activity
	if err := stmt.Get(&a, activityUUID); err != nil {
		return domain.Activity{},
			errors.WrapErrorf(err, errors.ErrCodeUnknown, "error getting activity")
	}
	return a, nil
}

func (r ActivityRepository) Activities() ([]domain.Activity, error) {
	stmt, err := r.statement(listActivity)
	if err != nil {
		return []domain.Activity{}, err
	}

	var as []domain.Activity
	if err := stmt.Select(&as); err != nil {
		return []domain.Activity{},
			errors.WrapErrorf(err, errors.ErrCodeUnknown, "error getting activities")
	}
	return as, nil
}

func (r ActivityRepository) CreateActivity(a *domain.Activity) error {
	stmt, err := r.statement(createActivity)
	if err != nil {
		return err
	}

	args := []interface{}{
		a.ContentUUID,
		a.Name,
		a.Description,
		a.ContentType,
		a.Taxonomy,
	}
	if err := stmt.QueryRowx(args...).StructScan(a); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error creating activity")
	}
	return nil
}

func (r ActivityRepository) UpdateActivity(a *domain.Activity) error {
	stmt, err := r.statement(updateActivity)
	if err != nil {
		return err
	}

	args := []interface{}{
		// set
		a.ContentUUID,
		a.Name,
		a.Description,
		a.ContentType,
		a.Taxonomy,
		// where
		a.UUID,
	}
	if err := stmt.QueryRowx(args...).StructScan(a); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error updating activity")
	}
	return nil
}

func (r ActivityRepository) DeleteActivity(activityUUID uuid.UUID) error {
	stmt, err := r.statement(deleteActivity)
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(activityUUID); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error deleting activity")
	}
	return nil
}
