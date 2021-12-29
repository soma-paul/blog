package postgres

import (
	"practice/blog/article/storage"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestCreateArticle(t *testing.T) {

	s := newTestStorage(t)

	tests := []struct {
		name    string
		in      storage.Articles
		want    int32
		wantErr bool
	}{
		{
			name: "CREATE_Article_SUCCESS",
			in: storage.Articles{
				Title:       "this is title",
				Description: "this is description",
				Author:      "soma",
				UserID:      1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			uid, err := s.CreateUser(storage.SignupUser{
				FirstName: "Soma",
				LastName:  "Paul",
				Username:  "soma",
				Email:     "somapaul@gmail.com",
				Password:  "password",
			})
			tt.in.UserID = uid
			got, err := s.CreateArticle(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Storage.Create() = %v, want %v", got, tt.want)
			}

		})
	}
}

func TestGetArticle(t *testing.T) {

	s := newTestStorage(t)

	tests := []struct {
		name    string
		in      int32
		want    *storage.Articles
		wantErr bool
	}{
		{
			name: "GET_Article_SUCCESS",
			in:   1,
			want: &storage.Articles{
				ID:          1,
				Title:       "this is title",
				Description: "this is description",
				Author:      "soma",
				UserID:      1,
			},
		},
		{
			name:    "FAILED_Article_DOES_NOT_EXIST",
			in:      100,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.GetIndexedArticle(tt.in)
			ignoreOpt := cmpopts.IgnoreFields(storage.Articles{}, "CreatedAt")
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !cmp.Equal(got, tt.want, ignoreOpt) {

				t.Errorf("Diff: got -, want += %v", cmp.Diff(got, tt.want, ignoreOpt))

			}
		})
	}
}
