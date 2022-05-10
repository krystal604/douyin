package errors_stuck

import "errors"

var PassWordWrongs = errors.New("password is wrong")

var DoesNotExist = errors.New("doesn't exist")
