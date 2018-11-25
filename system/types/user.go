package types

import (
	"time"

	"github.com/jmoiron/sqlx/types"
	"golang.org/x/crypto/bcrypt"
)

type (
	User struct {
		ID       uint64         `json:"id,string" db:"id"`
		Username string         `json:"username" db:"username"`
		Email    string         `json:"email" db:"email"`
		Name     string         `json:"name" db:"name"`
		Handle   string         `json:"handle" db:"handle"`
		Kind     string         `json:"kind" db:"kind"`
		SatosaID string         `json:"-" db:"satosa_id"`
		Meta     types.JSONText `json:"-" db:"meta"`

		OrganisationID uint64 `json:"organisationID,string" db:"rel_organisation"`
		UserID         uint64 `json:"userID,string" db:"rel_user_id"`
		User           *User  `json:"user" db:"-"`

		Password []byte `json:"-" db:"password"`

		CreatedAt   time.Time  `json:"createdAt,omitempty" db:"created_at"`
		UpdatedAt   *time.Time `json:"updatedAt,omitempty" db:"updated_at"`
		SuspendedAt *time.Time `json:"suspendedAt,omitempty" db:"suspended_at"`
		DeletedAt   *time.Time `json:"deletedAt,omitempty" db:"deleted_at"`

		Teams []*Team `json:"teams,omitempty" db:"-"`
	}

	UserFilter struct {
		Query string
	}
)

func (u *User) Valid() bool {
	return u.ID > 0 && u.SuspendedAt == nil && u.DeletedAt == nil
}

func (u *User) Identity() uint64 {
	return u.ID
}

func (u *User) ValidatePassword(password string) bool {
	return bcrypt.CompareHashAndPassword(u.Password, []byte(password)) == nil
}

func (u *User) GeneratePassword(password string) error {
	pwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = pwd
	return nil
}
