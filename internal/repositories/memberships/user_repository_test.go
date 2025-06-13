package memberships

import (
	"catalog-music/internal/models/memberships"
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Test_repository_CreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create sqlmock: %v", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))
	if err != nil {
		t.Fatalf("failed to open gorm db: %v", err)
	}
	assert.NoError(t, err)

	type args struct {
		user *memberships.User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		mokcFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				user: &memberships.User{
					Email:     "test@test.com",
					Username:  "testusername",
					Password:  "testpassword",
					CreatedBy: "test@test.com",
					UpdatedBy: "test@test.com",
				},
			},
			wantErr: false,
			mokcFn: func(args args) {
				mock.ExpectBegin()

				mock.ExpectQuery(`INSERT INTO "users" (.+) VALUES (.+)`).
					WithArgs(
						sqlmock.AnyArg(),
						sqlmock.AnyArg(),
						sqlmock.AnyArg(),
						args.user.Email,
						args.user.Username,
						args.user.Password,
						args.user.CreatedBy,
						args.user.UpdatedBy,
					).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

				mock.ExpectCommit()
			},
		},
		{
			name: "failed",
			args: args{
				user: &memberships.User{
					Email:     "test@test.com",
					Username:  "testusername",
					Password:  "testpassword",
					CreatedBy: "test@test.com",
					UpdatedBy: "test@test.com",
				},
			},
			wantErr: true,
			mokcFn: func(args args) {
				mock.ExpectBegin()

				mock.ExpectQuery(`INSERT INTO "users" (.+) VALUES (.+)`).
					WithArgs(
						sqlmock.AnyArg(),
						sqlmock.AnyArg(),
						sqlmock.AnyArg(),
						args.user.Email,
						args.user.Username,
						args.user.Password,
						args.user.CreatedBy,
						args.user.UpdatedBy,
					).
					WillReturnError(assert.AnError)

				mock.ExpectRollback()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mokcFn(tt.args)

			r := &repository{
				db: gormDB,
			}

			if err := r.CreateUser(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("repository.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func Test_repository_GetUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create sqlmock: %v", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))
	if err != nil {
		t.Fatalf("failed to open gorm db: %v", err)
	}
	assert.NoError(t, err)

	now := time.Now().UTC()

	type args struct {
		id       uint
		email    string
		username string
	}
	tests := []struct {
		name    string
		args    args
		want    *memberships.User
		wantErr bool
		mokcFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				id:       1,
				email:    "test@test.com",
				username: "testusername",
			},
			want: &memberships.User{
				Model: gorm.Model{
					ID:        1,
					CreatedAt: now,
					UpdatedAt: now,
				},
				Email:     "test@test.com",
				Username:  "testusername",
				Password:  "testpassword",
				CreatedBy: "test@test.com",
				UpdatedBy: "test@test.com",
			},
			wantErr: false,
			mokcFn: func(args args) {
				mock.ExpectQuery(`SELECT \* FROM "users" .+`).
					WithArgs(args.id, args.email, args.username, 1).
					WillReturnRows(sqlmock.NewRows([]string{
						"id", "created_at", "updated_at",
						"email", "username", "password", "created_by", "updated_by",
					}).AddRow(
						1,
						now,
						now,
						"test@test.com",
						"testusername",
						"testpassword",
						"test@test.com",
						"test@test.com",
					))
			},
		},
		{
			name: "failed",
			args: args{
				id:       1,
				email:    "test@test.com",
				username: "testusername",
			},
			wantErr: true,
			want:    nil,
			mokcFn: func(args args) {
				mock.ExpectQuery(`SELECT \* FROM "users" .+`).
					WithArgs(args.id, args.email, args.username, 1).
					WillReturnError(assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mokcFn(tt.args)

			r := &repository{
				db: gormDB,
			}
			got, err := r.GetUser(tt.args.id, tt.args.email, tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.GetUser() = %v, want %v", got, tt.want)
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
