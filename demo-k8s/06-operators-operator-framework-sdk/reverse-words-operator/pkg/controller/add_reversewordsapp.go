package controller

import (
	"github.com/golang-info/everse-words-operator/pkg/controller/reversewordsapp"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, reversewordsapp.Add)
}
