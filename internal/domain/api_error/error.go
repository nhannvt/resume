package api_error

import (
	"fmt"
	"net/http"
)

// TODO: remove http status code. domain shoudln't know the ui
type APIError interface {
	error
	HttpStatusCode() int
}

type TagGetFailedError struct {
	ID      int
	Context string
}

func (e *TagGetFailedError) Error() string {
	return fmt.Sprintf("Failed to get a tag(id=%d,context=%s)", e.ID, e.Context)
}

func (e *TagGetFailedError) HttpStatusCode() int {
	return http.StatusInternalServerError
}

type TagSearchFailedError struct {
	Query map[string]interface{}
}

func (e *TagSearchFailedError) Error() string {
	return fmt.Sprint("Failed to search tags")
}

func (e *TagSearchFailedError) HttpStatusCode() int {
	return http.StatusInternalServerError
}

type TagNotFoundError struct {
	ID      int
	Context string
}

func (e *TagNotFoundError) Error() string {
	return fmt.Sprintf("Tag(id=%d,context=%s) doesn't exist", e.ID, e.Context)
}

func (e *TagNotFoundError) HttpStatusCode() int {
	return http.StatusNotFound
}

type TagDeletionFailedError struct {
	ID      int
	Context string
}

func (e *TagDeletionFailedError) Error() string {
	return fmt.Sprintf("Failed to delete a tag(id=%d,context=%s)", e.ID, e.Context)
}

func (e *TagDeletionFailedError) HttpStatusCode() int {
	return http.StatusInternalServerError
}

type TagCreateFailedError struct {
}

func (e *TagCreateFailedError) Error() string {
	return fmt.Sprint("Failed to create a tag")
}

func (e *TagCreateFailedError) HttpStatusCode() int {
	return http.StatusInternalServerError
}

type TagUpdateFailedError struct {
}

func (e *TagUpdateFailedError) Error() string {
	return fmt.Sprint("Failed to update the tag")
}

func (e *TagUpdateFailedError) HttpStatusCode() int {
	return http.StatusInternalServerError
}

type AuthenticationError struct {
}

func (e *AuthenticationError) Error() string {
	return fmt.Sprint("Failed to authenticate client")
}

func (e *AuthenticationError) HttpStatusCode() int {
	return http.StatusUnauthorized
}

type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}

func (e *ValidationError) HttpStatusCode() int {
	return http.StatusBadRequest
}

type InconsistencyTagNameError struct {
}

func (e *InconsistencyTagNameError) Error() string {
	return "Found an inconsistency tag name. You can't register tag with different name with same id."
}

func (e *InconsistencyTagNameError) HttpStatusCode() int {
	return http.StatusConflict
}
