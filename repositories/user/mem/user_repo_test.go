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
		data *models.User
		want int64
	}{
		{
			data: &models.User{
				Name:    "Name1",
				Surname: "Surname1",
				Age:     1,
			},
			want: 1,
		},
		{
			data: &models.User{
				Name:    "Name2",
				Surname: "Surname2",
				Age:     2,
			},
			want: 2,
		},
		{
			data: &models.User{
				Name:    "Name3",
				Surname: "Surname3",
				Age:     3,
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
	tests := []*models.User{
		&models.User{
			Name:    "Name1",
			Surname: "Surname1",
			Age:     1,
		},
		&models.User{
			Name:    "Name2",
			Surname: "Surname2",
			Age:     2,
		},
		&models.User{
			Name:    "Name3",
			Surname: "Surname3",
			Age:     3,
		},
	}

	repo := New()

	for _, test := range tests {
		repo.Add(test)
	}

	test, err := repo.Get(1)
	if err != nil {
		t.Errorf("Got err selecting user: %v", err)
	}
	if test.ID != 1 {
		t.Errorf("Got wrong id. Got %v, want %v", test.ID, 1)
	}

	test, err = repo.Get(2)
	if err != nil {
		t.Errorf("Got err selecting user: %v", err)
	}
	if test.ID != 2 {
		t.Errorf("Got wrong id. Got %v, want %v", test.ID, 2)
	}

	test, err = repo.Get(3)
	if err != nil {
		t.Errorf("Got err selecting user: %v", err)
	}
	if test.ID != 3 {
		t.Errorf("Got wrong id. Got %v, want %v", test.ID, 3)
	}
}

func TestList(t *testing.T) {
	tests := []*models.User{
		&models.User{
			Name:    "Name1",
			Surname: "Surname1",
			Age:     1,
		},
		&models.User{
			Name:    "Name2",
			Surname: "Surname2",
			Age:     2,
		},
		&models.User{
			Name:    "Name3",
			Surname: "Surname3",
			Age:     3,
		},
	}

	repo := New()

	for _, test := range tests {
		repo.Add(test)
	}

	users, err := repo.List()
	if err != nil {
		t.Errorf("Got err selecting users list: %v", err)
	}
	if len(users) != 3 {
		t.Errorf("Got wrong users count in %v", users)
	}
}

func TestUpdate(t *testing.T) {
	tests := []*models.User{
		&models.User{
			Name:    "Name1",
			Surname: "Surname1",
			Age:     1,
		},
		&models.User{
			Name:    "Name2",
			Surname: "Surname2",
			Age:     2,
		},
		&models.User{
			Name:    "Name3",
			Surname: "Surname3",
			Age:     3,
		},
	}

	repo := New()

	for _, test := range tests {
		test.ID, _ = repo.Add(test)
	}

	tests[1].Name = "test"
	tests[1].Surname = "test"
	tests[1].Age = 48

	err := repo.Update(tests[1])
	if err != nil {
		t.Errorf("Got err updating: %v", err)
	}

	test, _ := repo.Get(tests[1].ID)
	if test.Name != "test" {
		t.Errorf("Wrong name value after update. Got %v, want %v", test.Name,
			"test")
	}
	if test.Surname != "test" {
		t.Errorf("Wrong surname value after update. Got %v, want %v", test.Surname,
			"test")
	}
	if test.Age != 48 {
		t.Errorf("Wrong age value after update. Got %v, want %v", test.Age,
			48)
	}
}

func TestDelete(t *testing.T) {
	tests := []*models.User{
		&models.User{
			Name:    "Name1",
			Surname: "Surname1",
			Age:     1,
		},
		&models.User{
			Name:    "Name2",
			Surname: "Surname2",
			Age:     2,
		},
		&models.User{
			Name:    "Name3",
			Surname: "Surname3",
			Age:     3,
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
