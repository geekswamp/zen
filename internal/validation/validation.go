package validation

import (
	"errors"
	"reflect"
	"strings"

	errs "github.com/geekswamp/zen/internal/errors"
	"github.com/geekswamp/zen/internal/http"
	"github.com/geekswamp/zen/internal/logger"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTrans "github.com/go-playground/validator/v10/translations/en"
)

var (
	trans    ut.Translator
	log      = logger.New()
	validate = validator.New()
)

func init() {
	translator, _ := ut.New(en.New()).GetTranslator("en")
	trans = translator

	if err := enTrans.RegisterDefaultTranslations(validate, trans); err != nil {
		log.Error(errs.ErrValidatorTrans.Error(), logger.ErrDetails(err))
	}

	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}

		return name
	})
}

// ValidateBody performs JSON binding and validation using custom validator
func ValidateBody[T any](c *gin.Context) (*T, *http.Error) {
	body := new(T)
	if err := c.ShouldBindJSON(body); err != nil {
		return nil, &http.Error{Code: http.NotValidJSONFormat.Code(), Reason: http.NotValidJSONFormat.Detail()}
	}

	if err := validateStruct(body); err != nil {
		return nil, err
	}

	return body, nil
}

// ValidateQuery performs query binding and validation using custom validator
func ValidateQuery[T any](c *gin.Context) (*T, *http.Error) {
	query := new(T)
	if err := c.ShouldBindQuery(query); err != nil {
		return nil, &http.Error{Code: http.NotValidQuery.Code(), Reason: err.Error()}
	}

	if err := validateStruct(query); err != nil {
		return nil, err
	}

	return query, nil
}

// validateStruct performs field-level validation and returns formatted error
func validateStruct(data any) *http.Error {
	if err := validate.Struct(data); err != nil {
		var validationErrs validator.ValidationErrors
		if errors.As(err, &validationErrs) {
			firstErr := validationErrs[0]
			return &http.Error{
				Code:   http.InputNotValid.Code(),
				Reason: firstErr.Translate(trans),
			}
		}
		return &http.Error{
			Code:   http.NotValidJSONFormat.Code(),
			Reason: err.Error(),
		}
	}

	return nil
}
