package postgres

import (
	"practice/blog/article/storage"
	"sort"
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

func TestGetAllArticles(t *testing.T) {
	s := newTestStorage(t)

	tests := []struct {
		name    string
		want    []*storage.Articles
		wanterr bool
	}{
		{
			name: "Get_All_Article_SUCCESS",
			want: []*storage.Articles{
				{
					ID:          1,
					Title:       "this is title",
					Description: "this is description",
					Author:      "soma",
					UserID:      1,
				},
				{
					Title:       "this is title2",
					Description: "this is description2",
					Author:      "soma2",
					UserID:      1,
				},
				{
					Title:       "this is title3",
					Description: "this is description3",
					Author:      "soma3",
					UserID:      1,
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			res, _ := s.CreateArticle(storage.Articles{
				Title:       "this is title2",
				Description: "this is description2",
				Author:      "soma2",
				UserID:      1,
			})
			tt.want[1].ID = res

			res, _ = s.CreateArticle(storage.Articles{
				Title:       "this is title3",
				Description: "this is description3",
				Author:      "soma3",
				UserID:      1,
			})
			tt.want[2].ID = res

			got, err := s.ShowAllArticles()

			ignoreOpt := cmpopts.IgnoreFields(storage.Articles{}, "CreatedAt")
			if (err != nil) != tt.wanterr {
				t.Errorf("Storage.Get() error = %v, wantErr %v", err, tt.wanterr)
				return
			}
			sort.Slice(tt.want, func(i, j int) bool {
				return tt.want[i].ID < tt.want[j].ID
			})

			sort.Slice(got, func(i, j int) bool {
				return got[i].ID < got[j].ID
			})

			if !cmp.Equal(got, tt.want, ignoreOpt) {

				t.Errorf("Diff: got -, want += %v", cmp.Diff(got, tt.want, ignoreOpt))

			}

		})
	}
}

func TestUpdateArticle(t *testing.T) {

	s := newTestStorage(t)

	tests := []struct {
		name    string
		in      storage.Articles
		want    storage.Articles
		wantErr bool
	}{
		{
			name: "Update_Article_SUCCESS",
			in: storage.Articles{
				ID:          1,
				Title:       "this is title7",
				Description: "this is description7",
				UserID:      1,
			},
			want: storage.Articles{
				ID:          1,
				Title:       "this is title7",
				Description: "this is description7",
				UserID:      1,
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			got, err := s.UpdateIndexedArticle(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			ignoreOpt := cmpopts.IgnoreFields(storage.Articles{}, "CreatedAt", "UpdatedAt", "Author")

			if !cmp.Equal(got, tt.want, ignoreOpt) {

				t.Errorf("Diff: got -, want += %v", cmp.Diff(got, tt.want, ignoreOpt))

			}

		})
	}
}

func TestDeleteArticle(t *testing.T) {

	s := newTestStorage(t)

	tests := []struct {
		name    string
		in      int32
		want    storage.Articles
		wantErr bool
	}{
		{
			name: "Delete_Article_SUCCESS",
			in:   2,
			want: storage.Articles{
				ID:          2,
				Title:       "this is title2",
				Description: "this is description2",
				Author:      "soma2",
				UserID:      1,
			},
		},
		{
			name:    "Delete_Article_FAILED",
			in:      100,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			got, err := s.DeleteArticleByID(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.DeleteArticleByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			ignoreOpt := cmpopts.IgnoreFields(storage.Articles{}, "CreatedAt", "UpdatedAt")
			if !cmp.Equal(got, tt.want, ignoreOpt) {
				t.Errorf("Diff: got -, want += %v", cmp.Diff(got, tt.want, ignoreOpt))
			}

		})
	}
}
