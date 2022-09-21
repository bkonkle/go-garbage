// Package example carries supporting code for the example cmd process
package example

import (
	"fmt"
	"strings"

	"github.com/ettle/strcase"
	"github.com/gin-gonic/gin"
	"github.com/juju/errors"
	"github.com/loopfz/gadgeto/tonic"
	"github.com/loopfz/gadgeto/tonic/utils/jujerr"
)

// ErrHook wraps the jujerr.ErrHook to apply custom logic, like validation error messages
func ErrHook(c *gin.Context, e error) (int, interface{}) {
	err := e

	if bindErr, ok := e.(tonic.BindError); ok {
		valErrs := bindErr.ValidationErrors()
		if valErrs != nil {
			errs := make([]string, 0, len(valErrs))

			for _, fieldErr := range valErrs {
				field := strcase.ToCamel(fieldErr.Field())
				tag := fieldErr.Tag()

				errs = append(errs, fmt.Sprintf("%s is %s", field, tag))
			}

			err = errors.BadRequestf("Validation error: %s", strings.Join(errs, ", "))
		} else {
			err = errors.BadRequestf(bindErr.Error())
		}
	}

	return jujerr.ErrHook(c, err)
}
