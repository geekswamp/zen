package validation

import (
	"github.com/geekswamp/zen/internal/errors"
	"github.com/geekswamp/zen/internal/logger"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTrans "github.com/go-playground/validator/v10/translations/en"
)

var (
	trans ut.Translator
	log   = logger.New()
)

func init() {
	trans, _ = ut.New(en.New()).GetTranslator("en")
	if err := enTrans.RegisterDefaultTranslations(binding.Validator.Engine().(*validator.Validate), trans); err != nil {
		log.Error(errors.ErrValidatorTrans.Error(), logger.ErrDetails(err))
	}
}

func Error(err error) (msg string) {
	if validationErrs, ok := err.(validator.ValidationErrors); !ok {
		return err.Error()
	} else {
		for _, e := range validationErrs {
			msg += e.Translate(trans) + ";"
		}
	}

	return msg
}
