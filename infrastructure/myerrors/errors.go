package myerrors

import "errors"

var (
	ErrExists    = errors.New("source exists")
	ErrNotExists = errors.New("source not exists")

	ErrWorkerBusy = errors.New("worker is busy")
	ErrNoInfo     = errors.New("no info need to persist")
)
