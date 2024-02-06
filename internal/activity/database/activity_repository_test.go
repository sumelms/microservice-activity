package database

import (
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"github.com/sumelms/microservice-activity/internal/activity/domain"
	utils "github.com/sumelms/microservice-activity/tests"
)

var (
	activity = domain.Activity{
		ID:        1,
		UUID:      utils.ActivityUUID,
		ForkID:    utils.ForkID,
		ContentID: utils.ContentID,

		Name:        "name-text",
		Description: "description-text",
		ContentType: "type-text",
		Taxonomy:    "taxonomy-text",

		CreatedAt: utils.Now,
		UpdatedAt: utils.Now,
		DeletedAt: nil,
	}
)

func newActivityTestDB() (*sqlx.DB, sqlmock.Sqlmock, map[string]*sqlmock.ExpectedPrepare) {
	return utils.NewTestDB(queriesActivity())
}

func TestRepository_Activity(t *testing.T) {
	validRows := sqlmock.NewRows([]string{
		"id", "uuid", "fork_id", "content_id",
		"name", "description", "content_type", "taxonomy",
		"created_at", "updated_at", "deleted_at",
	}).AddRow(
		activity.ID, activity.UUID, activity.ForkID, activity.ContentID,
		activity.Name, activity.Description, activity.ContentType, activity.Taxonomy,
		activity.CreatedAt, activity.UpdatedAt, activity.DeletedAt,
	)

	type args struct {
		id uuid.UUID
	}

	tests := []struct {
		name    string
		args    args
		rows    *sqlmock.Rows
		want    domain.Activity
		wantErr bool
	}{
		{
			name:    "get activity",
			args:    args{id: utils.ActivityUUID},
			rows:    validRows,
			want:    activity,
			wantErr: false,
		},
		{
			name:    "activity not found error",
			args:    args{id: uuid.MustParse("8281f61e-956e-4f64-ac0e-860c444c5f86")},
			rows:    utils.EmptyRows,
			want:    domain.Activity{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			db, _, stmts := newActivityTestDB()
			r, err := NewActivityRepository(db)
			if err != nil {
				t.Fatalf("an error '%s' was not expected creating the repository", err)
			}
			prep, ok := stmts[getActivity]
			if !ok {
				t.Fatalf("prepared statement %s not found", getActivity)
			}

			prep.ExpectQuery().WithArgs(utils.ActivityUUID).WillReturnRows(tt.rows)

			got, err := r.Activity(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Activity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Activity() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_Activities(t *testing.T) {
	validRows := sqlmock.NewRows([]string{
		"id", "uuid", "fork_id", "content_id",
		"name", "description", "content_type", "taxonomy",
		"created_at", "updated_at", "deleted_at",
	}).AddRow(
		activity.ID, activity.UUID, activity.ForkID, activity.ContentID,
		activity.Name, activity.Description, activity.ContentType, activity.Taxonomy,
		activity.CreatedAt, activity.UpdatedAt, activity.DeletedAt,
	).AddRow(
		2, uuid.MustParse("19cea06b-d068-44e1-ac76-d4761dec2e5d"), activity.ForkID, activity.ContentID,
		activity.Name, activity.Description, activity.ContentType, activity.Taxonomy,
		activity.CreatedAt, activity.UpdatedAt, activity.DeletedAt,
	)

	tests := []struct {
		name    string
		rows    *sqlmock.Rows
		wantLen int
		wantErr bool
	}{
		{
			name:    "get all activities",
			rows:    validRows,
			wantLen: 2,
			wantErr: false,
		},
		{
			name:    "get no activities",
			rows:    utils.EmptyRows,
			wantLen: 0,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			db, _, stmts := newActivityTestDB()
			r, err := NewActivityRepository(db)
			if err != nil {
				t.Fatalf("an error '%s' was not expected creating the repository", err)
			}
			prep, ok := stmts[listActivity]
			if !ok {
				t.Fatalf("prepared statement %s not found", listActivity)
			}

			prep.ExpectQuery().WillReturnRows(tt.rows)

			got, err := r.Activities()

			if (err != nil) != tt.wantErr {
				t.Errorf("Activities() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.wantLen {
				t.Errorf("Activities() got = %v, want %v", got, tt.wantLen)
			}
		})
	}
}

func TestRepository_CreateActivity(t *testing.T) {
	validRows := sqlmock.NewRows([]string{
		"id", "uuid", "fork_id", "content_id",
		"name", "description", "content_type", "taxonomy",
		"created_at", "updated_at", "deleted_at",
	}).AddRow(
		activity.ID, activity.UUID, activity.ForkID, activity.ContentID,
		activity.Name, activity.Description, activity.ContentType, activity.Taxonomy,
		activity.CreatedAt, activity.UpdatedAt, activity.DeletedAt,
	)

	type args struct {
		a *domain.Activity
	}

	tests := []struct {
		name    string
		rows    *sqlmock.Rows
		args    args
		wantErr bool
	}{
		{
			name:    "create activity",
			rows:    validRows,
			args:    args{a: &activity},
			wantErr: false,
		},
		{
			name:    "empty fields",
			rows:    utils.EmptyRows,
			args:    args{a: &activity},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			db, _, stmts := newActivityTestDB()
			r, err := NewActivityRepository(db)
			if err != nil {
				t.Fatalf("an error '%s' was not expected when creating the repository", err)
			}
			prep, ok := stmts[createActivity]
			if !ok {
				t.Fatalf("prepared statement %s not found", createActivity)
			}

			prep.ExpectQuery().WillReturnRows(tt.rows)

			if err := r.CreateActivity(tt.args.a); (err != nil) != tt.wantErr {
				t.Errorf("CreateActivity() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepository_UpdateActivity(t *testing.T) {
	validRows := sqlmock.NewRows([]string{
		"id", "uuid", "fork_id", "content_id",
		"name", "description", "content_type", "taxonomy",
		"created_at", "updated_at", "deleted_at",
	}).AddRow(
		activity.ID, activity.UUID, activity.ForkID, activity.ContentID,
		activity.Name, activity.Description, activity.ContentType, activity.Taxonomy,
		activity.CreatedAt, activity.UpdatedAt, activity.DeletedAt,
	)

	type fields struct {
		DB *sqlx.DB
	}
	type args struct {
		a *domain.Activity
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		rows    *sqlmock.Rows
		wantErr bool
	}{
		{
			name:    "update activity",
			args:    args{a: &activity},
			rows:    validRows,
			wantErr: false,
		},
		{
			name:    "empty activity",
			args:    args{a: &domain.Activity{}},
			rows:    utils.EmptyRows,
			wantErr: true,
		},
	}
	for _, testCase := range tests {
		tt := testCase
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			db, _, stmts := newActivityTestDB()
			r, err := NewActivityRepository(db)
			if err != nil {
				t.Fatalf("an error '%s' was not expected when creating the repository", err)
			}
			prep, ok := stmts[updateActivity]
			if !ok {
				t.Fatalf("prepared statement %s not found", updateActivity)
			}

			prep.ExpectQuery().WillReturnRows(tt.rows)

			if err := r.UpdateActivity(tt.args.a); (err != nil) != tt.wantErr {
				t.Errorf("UpdateActivity() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
