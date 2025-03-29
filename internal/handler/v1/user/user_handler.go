package user

import (
	"github.com/geekswamp/zen/internal/base"
	"github.com/geekswamp/zen/internal/core"
	"github.com/geekswamp/zen/internal/http"
	"github.com/geekswamp/zen/internal/model"
	"github.com/geekswamp/zen/internal/service"
	"github.com/geekswamp/zen/internal/validation"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	resp    http.BaseResponse
	service service.UserService
	context core.Context
}

func NewUserHandler(resp http.BaseResponse, service service.UserService) UserHandler {
	return UserHandler{resp: resp, service: service}
}

func (h UserHandler) Register(ctx *gin.Context) {
	body, err := validation.ValidateBody[UserCreateRequest](ctx)
	if err != nil {
		h.resp.Error(ctx, err)
		return
	}

	if err := h.service.Create(body.FullName, body.Email, body.Password, body.Phone, model.Gender(body.Gender)); err != nil {
		h.resp.Error(ctx, err)
		return
	}

	h.resp.Created(ctx, nil)
}

func (h UserHandler) GetCurrent(ctx *gin.Context) {
	user, err := h.service.Get(h.context.GetUserSession().ID)
	if err != nil {
		h.resp.Error(ctx, err)
		return
	}

	u := UserInfoResponse{
		ID:            user.ID,
		FullName:      user.FullName,
		Email:         user.Email,
		Phone:         user.Phone,
		Active:        user.Active,
		Gender:        int(user.Gender),
		ActivatedTime: user.ActivatedTime,
		CreatedTime:   user.CreatedTime,
		UpdateTime:    user.UpdatedTime,
		DeletedTime:   user.DeletedTime,
	}

	h.resp.Success(ctx, u)
}

func (h UserHandler) GetDetail(ctx *gin.Context) {
	ID, err := h.context.ParseIDParam()
	if err != nil {
		h.resp.Error(ctx, err)
		return
	}

	user, err := h.service.Get(ID)
	if err != nil {
		h.resp.Error(ctx, err)
		return
	}

	u := UserInfoResponse{
		ID:            user.ID,
		FullName:      user.FullName,
		Email:         user.Email,
		Phone:         user.Phone,
		Active:        user.Active,
		Gender:        int(user.Gender),
		ActivatedTime: user.ActivatedTime,
		CreatedTime:   user.CreatedTime,
		UpdateTime:    user.UpdatedTime,
		DeletedTime:   user.DeletedTime,
	}

	h.resp.Success(ctx, u)
}

func (h UserHandler) Update(ctx *gin.Context) {
	body, err := validation.ValidateBody[UserUpdateInfoRequest](ctx)
	if err != nil {
		h.resp.Error(ctx, err)
		return
	}

	if err := h.service.Update(h.context.GetUserSession().ID, base.UpdateMap{
		"FullName": body.FullName,
		"Email":    body.Email,
		"Phone":    body.Phone,
		"Gender":   body.Gender,
	}); err != nil {
		h.resp.Error(ctx, err)
		return
	}

	h.resp.Success(ctx, nil)
}

func (h UserHandler) HardDelete(ctx *gin.Context) {
	ID, err := h.context.ParseIDParam()
	if err != nil {
		h.resp.Error(ctx, err)
		return
	}

	if err := h.service.Delete(ID); err != nil {
		h.resp.Error(ctx, err)
		return
	}

	h.resp.Success(ctx, nil)
}

func (h UserHandler) SoftDelete(ctx *gin.Context) {
	ID, err := h.context.ParseIDParam()
	if err != nil {
		h.resp.Error(ctx, err)
		return
	}

	if err := h.service.SoftDelete(ID); err != nil {
		h.resp.Error(ctx, err)
		return
	}

	h.resp.Success(ctx, nil)
}

func (h UserHandler) SetToActive(ctx *gin.Context) {
	ID, err := h.context.ParseIDParam()
	if err != nil {
		h.resp.Error(ctx, err)
		return
	}

	if err := h.service.SetToActive(ID); err != nil {
		h.resp.Error(ctx, err)
		return
	}

	h.resp.Success(ctx, nil)
}

func (h UserHandler) SetToInactive(ctx *gin.Context) {
	ID, err := h.context.ParseIDParam()
	if err != nil {
		h.resp.Error(ctx, err)
		return
	}

	if err := h.service.SetToInactive(ID); err != nil {
		h.resp.Error(ctx, err)
		return
	}

	h.resp.Success(ctx, nil)
}
