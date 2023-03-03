package account

import (
	"errors"
	"net/url"

	"github.com/gameloungee/server/internal"
	"github.com/gameloungee/server/pkg/mysql"
	"github.com/google/uuid"
)

type Account struct {
	Id       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	Login    string    `json:"login"`
	Password Password  `json:"password"`
	Nickname string    `json:"nickname"`
	Age      int       `json:"age"`
	Avatar   url.URL   `json:"avatar"`
	Gender   string    `json:"gender"`
	Country  string    `json:"country"`
	City     string    `json:"city"`
}

func (a *Account) ToDatabase() error {
	//_, err := mail.ParseAddress(a.Email)

	//if err != nil {
	//return err
	//}

	a.Id = uuid.New()

	pass, err := a.Password.String()

	if err != nil {
		return err
	}

	if !internal.HasNilFileds(a) {
		err = mysql.WithPrepare("INSERT INTO accounts (id, email, login, password, nickname, age, avatar, gender, country, city)"+
			"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)", a.Id, a.Email, a.Login, pass, a.Nickname, a.Age, a.Avatar.String(), a.Gender, a.Country, a.City)

		if err != nil {
			return err
		}
	} else {
		return errors.New("one of the fields of the structure is empty")
	}

	return nil
}
