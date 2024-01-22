// This file contains the interfaces for the repository layer.
// The repository layer is responsible for interacting with the database.
// For testing purpose we will generate mock implementations of these
// interfaces using mockgen. See the Makefile for more information.
package repository

// RepositoryInterface ...
type RepositoryInterface interface {
	SaveUser(entity *Users) (err error)
	FindUserByPhoneNumber(phoneNumber string) (entity *Users, err error)
	FindUserById(userID string) (entity *Users, err error)
	FindUserByName(name string) (entity *Users, err error)
}
