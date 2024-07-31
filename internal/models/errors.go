package models

import "errors"

var (
	ErrNoRecord           = errors.New("models: no macthing record found")
	ErrInvalidCredtionals = errors.New("models :invalid credtionals")

	ErrDuplicateEmail = errors.New("models :duplicate email")
)
