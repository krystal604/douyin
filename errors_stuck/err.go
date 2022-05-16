package errors_stuck

import "errors"

var PassWordWrongs = errors.New("password is wrong")

var DoesNotExist = errors.New("doesn't exist")

var AlreadyExists = errors.New("already exists")

var NoAction = errors.New("no this action")
