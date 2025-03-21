package validation

import (
	errors2 "errors"
	"reflect"
	"strings"

	"github.com/geekswamp/zen/internal/errors"
	"github.com/geekswamp/zen/internal/http"
	"github.com/geekswamp/zen/internal/logger"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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
	trans, _ = ut.New(en.New()).GetTranslator("en")

	if err := enTrans.RegisterDefaultTranslations(binding.Validator.Engine().(*validator.Validate), trans); err != nil {
		log.Error(errors.ErrValidatorTrans.Error(), logger.ErrDetails(err))
	}

	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	if err := enTrans.RegisterDefaultTranslations(binding.Validator.Engine().(*validator.Validate), trans); err != nil {
		log.Error(errors.ErrValidatorTrans.Error(), logger.ErrDetails(err))
	}
}

// ValidateQuery is a generic function that validates query parameters from a gin.Context.
func ValidateQuery[T any](c *gin.Context) (*T, any) {
	query := new(T)
	if err := c.ShouldBindQuery(query); err != nil {
		return nil, http.Error{Code: http.PANotValidQuery, Reason: err.Error()}
	}

	if err := validateStruct(query); err != nil {
		return nil, err
	}

	return query, nil
}

// ValidateBody is a generic function that validates and binds the request body to a given struct type T.
func ValidateBody[T any](c *gin.Context) (*T, any) {
	body := new(T)

	if err := c.ShouldBindJSON(body); err != nil {
		code := http.PANotValidJSONFormat
		errStr := err.Error()
		return nil, http.Error{Code: code, Reason: errStr}
	}

	if err := validateStruct(body); err != nil {
		return nil, err
	}

	return body, nil
}

func validateStruct(body any) *http.Error {
	var code, msg string

	if err := validate.Struct(body); err != nil {
		var validationErrs validator.ValidationErrors
		if errors2.As(err, &validationErrs) {
			firstErr := validationErrs[0]
			code = http.PAInputNotValid
			msg = firstErr.Translate(trans)

			return &http.Error{Code: code, Reason: msg}
		}

		code = http.PANotValidJSONFormat
		msg = err.Error()

		return &http.Error{Code: code, Reason: msg}
	}

	return nil
}
