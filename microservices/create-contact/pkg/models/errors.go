package models

import "encoding/json"

type InternalError struct {
	Op  string
	Err error
}

func (e *InternalError) Error() string {
	return e.Op + ": " + e.Err.Error()
}

func (e *InternalError) Unwrap() error {
	return e.Err
}

type MarshalError struct {
	In  interface{}
	Err error
}

func (e *MarshalError) Error() string {
	b, err := json.Marshal(e.In)
	if err != nil {
		return "Input: " + e.Err.Error()
	}
	return "Input: " + string(b) + "| Error: " + e.Err.Error()
}

func (e *MarshalError) Unwrap() error {
	return e.Err
}

type UnmarshalError struct {
	Out interface{}
	Err error
}

func (e *UnmarshalError) Error() string {
	b, err := json.Marshal(e.Out)
	if err != nil {
		return "Input: " + e.Err.Error()
	}
	return "Input: " + string(b) + "| Error: " + e.Err.Error()
}

func (e *UnmarshalError) Unwrap() error {
	return e.Err
}

type NotFoundError struct {
	Attribute string
	Err       error
}

func (e *NotFoundError) Error() string {
	return "Attribute " + e.Attribute + " " + e.Err.Error()
}

func (e *NotFoundError) Unwrap() error {
	return e.Err
}
