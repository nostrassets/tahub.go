package v2controllers

import (
	"net/http"

	"github.com/getAlby/lndhub.go/lib/responses"
	"github.com/getAlby/lndhub.go/lib/service"
	"github.com/labstack/echo/v4"
)

// CreateUserController : Create user controller struct
type CreateUserController struct {
	svc *service.LndhubService
	responder responses.RelayResponder
}

func NewCreateUserController(svc *service.LndhubService) *CreateUserController {
	return &CreateUserController{svc: svc, responder: responses.RelayResponder{}}
}

type CreateUserResponseBody struct {
	ID       int64  `json:"id"`
	Pubkey   string `json:"pubkey"`
}
type CreateUserRequestBody struct {
	Pubkey   string `json:"pubkey"`
}

// CreateUser godoc
// @Summary      Create an account
// @Description  Create a new account with a pubkey and password
// @Accept       json
// @Produce      json
// @Tags         Account
// @Param        account  body      CreateUserRequestBody  false  "Create User"
// @Success      200      {object}  CreateUserResponseBody
// @Failure      400      {object}  responses.ErrorResponse
// @Failure      500      {object}  responses.ErrorResponse
// @Router       /v2/users [post]
func (controller *CreateUserController) CreateUser(c echo.Context) error {

	var body CreateUserRequestBody

	if err := c.Bind(&body); err != nil {
		c.Logger().Errorf("Failed to load create user request body: %v", err)
		return c.JSON(http.StatusBadRequest, responses.BadArgumentsError)
	}
	if err := c.Validate(&body); err != nil {
		c.Logger().Errorf("Invalid Nostr Event request body: %v", err)
		// TODO this is not a nostr error responder
		return controller.responder.NostrErrorJson(c, responses.BadArgumentsError.Message)
	}
	user, err := controller.svc.CreateUser(c.Request().Context(), body.Pubkey)
	if err != nil {
		c.Logger().Errorf("Failed to create user: %v", err)
		return c.JSON(http.StatusBadRequest, responses.BadArgumentsError)
	}

	var ResponseBody CreateUserResponseBody
	ResponseBody.Pubkey = user.Pubkey
	ResponseBody.ID = user.ID

	return c.JSON(http.StatusOK, &ResponseBody)
}
