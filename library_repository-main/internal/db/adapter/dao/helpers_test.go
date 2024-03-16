package dao

import (
	"github.com/geejjoo/library_repository/internal/models"
	"reflect"
	"testing"
)

func TestGetStructInfo(t *testing.T) {
	user := models.UserDTO{
		ID:      0,
		Name:    "",
		Surname: "",
		Age:     0,
	}
	info := GetStructInfo(&user)
	expected := []string{"id", "name", "surname", "age"}
	var res []string
	for _, v := range info.Fields {
		res = append(res, v)
	}
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("Результат неверный, ожидалось: %v, получил: %v", expected, res)
	}

}
