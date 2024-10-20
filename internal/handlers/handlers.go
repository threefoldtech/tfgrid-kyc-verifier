package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"

	"example.com/tfgrid-kyc-service/internal/models"
	"example.com/tfgrid-kyc-service/internal/responses"
	"example.com/tfgrid-kyc-service/internal/services"
)

type Handler struct {
	tokenService        services.TokenService
	verificationService services.VerificationService
	coordinatorService  services.CoordinatorService
}

func NewHandler(tokenService services.TokenService, verificationService services.VerificationService, coordinatorService services.CoordinatorService) *Handler {
	return &Handler{tokenService: tokenService, verificationService: verificationService, coordinatorService: coordinatorService}
}

// @Summary		Get or Generate iDenfy Verification Token
// @Description	Returns a token for a client
// @Tags			Token
// @Accept			json
// @Produce		json
// @Param			X-Client-ID	header		string	true	"TFChain SS58Address"								minlength(48)	maxlength(48)
// @Param			X-Challenge	header		string	true	"hex-encoded message `{api-domain}:{timestamp}`"
// @Param			X-Signature	header		string	true	"hex-encoded sr25519|ed25519 signature"				minlength(128)	maxlength(128)
// @Success		200			{object}	responses.TokenResponse "Existing token retrieved"
// @Success		201			{object}	responses.TokenResponse "New token created"
// @Router			/api/v1/token [post]
func (h *Handler) GetorCreateVerificationToken() fiber.Handler {
	return func(c *fiber.Ctx) error {
		clientID := c.Get("X-Client-ID")

		token, isNewToken, err := h.coordinatorService.GetorCreateVerificationToken(c.Context(), clientID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		response := responses.NewTokenResponseWithStatus(token, isNewToken)
		if isNewToken {
			return c.Status(fiber.StatusCreated).JSON(fiber.Map{"result": response})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"result": response})
	}
}

// @Summary		Get Verification Data
// @Description	Returns the verification data for a client
// @Tags			Verification
// @Accept			json
// @Produce		json
// @Param			X-Client-ID	header		string	true	"TFChain SS58Address"								minlength(48)	maxlength(48)
// @Param			X-Challenge	header		string	true	"hex-encoded message `{api-domain}:{timestamp}`"
// @Param			X-Signature	header		string	true	"hex-encoded sr25519|ed25519 signature"				minlength(128)	maxlength(128)
// @Success		200			{object}	responses.VerificationDataResponse
// @Router			/api/v1/data [get]
func (h *Handler) GetVerificationData() fiber.Handler {
	return func(c *fiber.Ctx) error {
		clientID := c.Get("X-Client-ID")
		verification, err := h.verificationService.GetVerification(c.Context(), clientID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		if verification == nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Verification not found"})
		}
		response := responses.NewVerificationDataResponse(verification)
		return c.JSON(fiber.Map{"result": response})
	}
}

// @Summary		Get Verification Status
// @Description	Returns the verification status for a client
// @Tags			Verification
// @Accept			json
// @Produce		json
// @Param			X-Client-ID	header		string	true	"TFChain SS58Address"								minlength(48)	maxlength(48)
// @Param			X-Challenge	header		string	true	"hex-encoded message `{api-domain}:{timestamp}`"
// @Param			X-Signature	header		string	true	"hex-encoded sr25519|ed25519 signature"				minlength(128)	maxlength(128)
// @Success		200			{object}	responses.VerificationStatusResponse
// @Router			/api/v1/status [get]
func (h *Handler) GetVerificationStatus() fiber.Handler {
	return func(c *fiber.Ctx) error {
		clientID := c.Get("X-Client-ID")
		verification, err := h.verificationService.GetVerification(c.Context(), clientID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		if verification == nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Verification not found"})
		}
		response := responses.NewVerificationStatusResponse(verification)
		return c.JSON(fiber.Map{"result": response})
	}
}

// @Summary		Process Verification Update
// @Description	Processes the verification update for a client
// @Tags			Webhooks
// @Accept			json
// @Produce		json
// @Success		200
// @Router			/webhooks/idenfy/verification-update [post]
func (h *Handler) ProcessVerificationResult() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// print request body and headers and return 200
		fmt.Printf("%+v", c.Body())
		fmt.Printf("%+v", &c.Request().Header)
		sigHeader := c.Get("Idenfy-Signature")
		if len(sigHeader) < 1 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "No signature provided"})
		}
		body := c.Body()
		// encode the body to json and save it to the database
		var result models.Verification
		decoder := json.NewDecoder(bytes.NewReader(body))
		err := decoder.Decode(&result)
		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		fmt.Printf("after decoding: %+v", result)
		err = h.verificationService.ProcessVerificationResult(c.Context(), body, sigHeader, result)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.SendStatus(fiber.StatusOK)
	}
}

// @Summary		Process Doc Expiration Notification
// @Description	Processes the doc expiration notification for a client
// @Tags			Webhooks
// @Accept			json
// @Produce		json
// @Success		200
// @Router			/webhooks/idenfy/id-expiration [post]
func (h *Handler) ProcessDocExpirationNotification() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return nil
	}
}
