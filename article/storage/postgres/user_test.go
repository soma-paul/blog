package postgres

import (
	"practice/blog/article/storage"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestCreateUser(t *testing.T) {

	s := newTestStorage(t)
	password := "1234"
	hash, er := bcrypt.GenerateFromPassword([]byte(password), 10)
	if er != nil {
		t.Error("error bycrypting password", er)
	}
	tests := []struct {
		name    string
		in      storage.SignupUser
		want    int32
		wantErr bool
	}{
		{
			name: "CREATE_User_SUCCESS",
			in: storage.SignupUser{
				FirstName: "Soma",
				LastName:  "Paul",
				Username:  "som",
				Email:     "somapaul1@gmail.com",
				Password:  string(hash),
			},
			want: 2,
		},
		{
			name: "CREATE_User_Failure_Duplicate_Username",
			in: storage.SignupUser{
				FirstName: "Soma",
				LastName:  "Paul",
				Username:  "soma",
				Email:     "somapaul1@gmail.com",
				Password:  string(hash),
			},
			wantErr: true,
		},
		{
			name: "CREATE_User_Failure_Duplicate_Email",
			in: storage.SignupUser{
				FirstName: "Soma",
				LastName:  "Paul",
				Username:  "soma2",
				Email:     "somapaul1@gmail.com",
				Password:  string(hash),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.CreateUser(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Storage.CreateUser() = %v, want %v", got, tt.want)
			}

		})
	}
}

/*
func TestGetTodo(t *testing.T) {

	s := newTestStorage(t)

	tests := []struct {
		name    string
		in      int64
		want    *storage.Articles
		wantErr bool
	}{
		{
			name: "GET_TODO_SUCCESS",
			in:   1,
			want: &storage.Articles{},
		},
		{
			name:    "FAILED_TODO_DOES_NOT_EXIST",
			in:      100,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.Get(context.TODO(), tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !cmp.Equal(got, tt.want) {
				t.Errorf("Diff: got -, want += %v", cmp.Diff(got, tt.want))
			}
		})
	}
}

func TestUpdateTodo(t *testing.T) {

	s := newTestStorage(t)

	tests := []struct {
		name    string
		in      storage.Todo
		want    *storage.Todo
		wantErr bool
	}{
		{
			name: "UPDATE_TODO_SUCCESS",
			in: storage.Todo{
				ID:          1,
				Title:       "this is title update",
				Description: "this is description update",
			},
			want: &storage.Todo{
				ID:          1,
				Title:       "this is title update",
				Description: "this is description update",
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.Update(context.TODO(), tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("Diff: got -, want += %v", cmp.Diff(got, tt.want))
			}
		})
	}
}

func TestDeleteTodo(t *testing.T) {

	s := newTestStorage(t)

	tests := []struct {
		name    string
		in      int64
		wantErr bool
	}{
		{
			name: "DELETE_TODO_SUCCESS",
			in:   1,
		},
		{
			name:    "FAILED_TO_DELETE_TODO_ID_DOES_NOT_EXIST",
			in:      100,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			err := s.Delete(context.TODO(), tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}

}
*/
