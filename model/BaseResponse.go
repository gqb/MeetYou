package model

type BaseResponse[T any] struct {
	Code int
	Message string
	Data T
}