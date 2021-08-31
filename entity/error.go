package entity

import "errors"

//ErrNotFound not found
var ErrNotFound = errors.New("Not found")

//ErrInvalidEntity invalid entity
var ErrInvalidEntity = errors.New("Invalid entity")

//ErrCannotBeDeleted cannot be deleted
var ErrCannotBeDeleted = errors.New("Cannot Be Deleted")

//ErrNotEnoughAnimes cannot borrow
var ErrNotEnoughAnimes = errors.New("Not enough animes")

//ErrBookAlreadyBorrowed cannot borrow
var ErrBookAlreadyBorrowed = errors.New("Book already borrowed")

//ErrBookNotBorrowed cannot return
var ErrBookNotBorrowed = errors.New("Book not borrowed")
