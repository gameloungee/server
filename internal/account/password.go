package account

import "golang.org/x/crypto/bcrypt"

type Password struct {
	source string
}

func CreatePassword(pass string) *Password {
	return &Password{
		source: pass,
	}
}

func CreateFromHash(hash string) *Password {
	return &Password{
		source: string(hash),
	}
}

func (p *Password) String() (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(p.source), 10)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (p *Password) Compare(input string) bool {
	password, err := p.String()

	if err != nil {
		return false
	}

	return bcrypt.CompareHashAndPassword([]byte(password), []byte(input)) == nil
}
