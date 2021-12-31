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
