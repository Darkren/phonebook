package mem

import (
	"testing"

	"github.com/Darkren/phonebook/models"
)

func TestNew(t *testing.T) {
	repo := New()
	if repo == nil {
		t.Error("Got err creating repo")
	}
}

func TestAdd(t *testing.T) {
	tests := []struct {
		data *models.Phone
		want int64
	}{
		{
			data: &models.Phone{
				Phone:  "123123123",
				UserID: 1,
			},
			want: 1,
		},
		{
			data: &models.Phone{
				Phone:  "234",
				UserID: 2,
			},
			want: 2,
		},
		{
			data: &models.Phone{
				Phone:  "567",
				UserID: 1,
			},
			want: 3,
		},
	}

	repo := New()

	for _, test := range tests {
		id, err := repo.Add(test.data)
		if err != nil {
			t.Errorf("Got err on inserting: %v", err)
		}

		if id != test.want {
			t.Errorf("Wron id. Expected: %v, got: %v", id, test.want)
		}
	}
}

func TestGet(t *testing.T) {
	tests := []*models.Phone{
		&models.Phone{
			Phone:  "123123123",
			UserID: 1,
		},
		&models.Phone{
			Phone:  "234",
			UserID: 2,
		},
		&models.Phone{
			Phone:  "567",
			UserID: 1,
		},
	}

	repo := New()

	for _, test := range tests {
		repo.Add(test)
	}

	test, err := repo.Get(1)
	if err != nil {
		t.Errorf("Got err selecting phone: %v", err)
	}
	if test.ID != 1 {
		t.Errorf("Got wrong id. Got %v, want %v", test.ID, 1)
	}

	test, err = repo.Get(2)
	if err != nil {
		t.Errorf("Got err selecting phone: %v", err)
	}
	if test.ID != 2 {
		t.Errorf("Got wrong id. Got %v, want %v", test.ID, 2)
	}

	test, err = repo.Get(3)
	if err != nil {
		t.Errorf("Got err selecting phone: %v", err)
	}
	if test.ID != 3 {
		t.Errorf("Got wrong id. Got %v, want %v", test.ID, 3)
	}
}

func TestListByUser(t *testing.T) {
	tests := []*models.Phone{
		&models.Phone{
			Phone:  "123123123",
			UserID: 1,
		},
		&models.Phone{
			Phone:  "234",
			UserID: 2,
		},
		&models.Phone{
			Phone:  "567",
			UserID: 1,
		},
	}

	repo := New()

	for _, test := range tests {
		repo.Add(test)
	}

	users, err := repo.ListByUser(2)
	if err != nil {
		t.Errorf("Got err selecting phones by user id: %v", err)
	}
	if len(users) != 1 {
		t.Errorf("Got wrong users count in %v", users)
	}
	if users[0].UserID != 2 {
		t.Errorf("Got wrong user id. Got %v, want %v", users[0].UserID, 2)
	}

	users, err = repo.ListByUser(1)
	if err != nil {
		t.Errorf("Got err selecting phones by user id: %v", err)
	}
	if len(users) != 2 {
		t.Errorf("Got wrong users count in %v", users)
	}
	if users[0].UserID != 1 {
		t.Errorf("Got wrong user id. Got %v, want %v", users[0].UserID, 2)
	}
	if users[1].UserID != 1 {
		t.Errorf("Got wrong user id. Got %v, want %v", users[0].UserID, 2)
	}
}

func TestUpdate(t *testing.T) {
	tests := []*models.Phone{
		&models.Phone{
			Phone:  "123123123",
			UserID: 1,
		},
		&models.Phone{
			Phone:  "234",
			UserID: 2,
		},
		&models.Phone{
			Phone:  "567",
			UserID: 1,
		},
	}

	repo := New()

	for _, test := range tests {
		test.ID, _ = repo.Add(test)
	}

	tests[1].Phone = "test"

	err := repo.Update(tests[1])
	if err != nil {
		t.Errorf("Got err updating: %v", err)
	}

	test, _ := repo.Get(tests[1].ID)
	if test.Phone != "test" {
		t.Errorf("Wrong phone values after update. Got %v, want %v", test.Phone,
			tests[1].Phone)
	}
}

func TestDelete(t *testing.T) {
	tests := []*models.Phone{
		&models.Phone{
			Phone:  "123123123",
			UserID: 1,
		},
		&models.Phone{
			Phone:  "234",
			UserID: 2,
		},
		&models.Phone{
			Phone:  "567",
			UserID: 1,
		},
	}

	repo := New()

	for _, test := range tests {
		test.ID, _ = repo.Add(test)
	}

	err := repo.Delete(tests[1].ID)
	if err != nil {
		t.Errorf("Got err deleting: %v", err)
	}

	test, _ := repo.Get(tests[1].ID)
	if test != nil {
		t.Errorf("Expected nothing, got: %v", test)
	}
}
