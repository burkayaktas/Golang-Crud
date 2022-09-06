package handlers

import (
	"net/http"
	"testing"

	"github.com/tutorials/go/crud/models"
	"gorm.io/gorm"
)

type testUsersRepository struct {
	users []models.User
}

func (r *testUsersRepository) Find(name string) (models.User, error) {
	for _, user := range r.users {
		if user.Name == name {
			return user, nil
		}
	}
	return models.User{}, nil
}

func (r *testUsersRepository) Insert(polygon models.User) error {
	r.users = append(r.users, r.users...)
	return nil
}
func Test_handler_GetAllUsers(t *testing.T) {
	repo := &testUsersRepository{}
	repo.users = []models.User{
		{
			Name:  "Burkay",
			Email: "mburkayaktas@gmail.com",
		},
		{
			Name:  "Deneme",
			Email: "example@gmail.com",
		},
	}
	h := handler{
		DB: &gorm.DB{},
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		h    handler
		args args
	}{
		{
			name: "test",
			h:    h,
			args: args{
				w: nil,
				r: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.GetAllUsers(tt.args.w, tt.args.r)
		})
	}
}

// unit test add user
func Test_handler_AddUser(t *testing.T) {
	repo := &testUsersRepository{}
	repo.users = []models.User{
		{
			Name:  "Burkay",
			Email: "",
		},
		{
			Name:  "Deneme",
			Email: "",
		},
	}
	h := handler{
		DB: &gorm.DB{},
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		h    handler
		args args
	}{
		{
			name: "test",
			h:    h,
			args: args{
				w: nil,
				r: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.AddUser(tt.args.w, tt.args.r)
		})
	}
}

