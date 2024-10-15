package responses

import "github.com/gofiber/fiber/v2"

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}

func SuccessResponse(c *fiber.Ctx, statusCode int, data interface{}, message string) error {
	return c.Status(statusCode).JSON(Response{
		Success: true,
		Data:    data,
		Message: message,
	})
}

type TokenResponse struct {
	Message       string `json:"message"`
	AuthToken     string `json:"authToken"`
	ScanRef       string `json:"scanRef"`
	ClientID      string `json:"clientId"`
	ExpiryTime    int    `json:"expiryTime"`
	SessionLength int    `json:"sessionLength"`
	DigitString   string `json:"digitString"`
	TokenType     string `json:"tokenType"`
}

type VerificationStatusResponse struct {
	FraudTags      []string `json:"fraudTags"`
	MismatchTags   []string `json:"mismatchTags"`
	AutoDocument   string   `json:"autoDocument"`
	AutoFace       string   `json:"autoFace"`
	ManualDocument string   `json:"manualDocument"`
	ManualFace     string   `json:"manualFace"`
	ScanRef        string   `json:"scanRef"`
	ClientID       string   `json:"clientId"`
	Status         string   `json:"status"`
}

type VerificationDataResponse struct {
	DocFirstName           string      `json:"docFirstName"`
	DocLastName            string      `json:"docLastName"`
	DocNumber              string      `json:"docNumber"`
	DocPersonalCode        string      `json:"docPersonalCode"`
	DocExpiry              string      `json:"docExpiry"`
	DocDob                 string      `json:"docDob"`
	DocDateOfIssue         string      `json:"docDateOfIssue"`
	DocType                string      `json:"docType"`
	DocSex                 string      `json:"docSex"`
	DocNationality         string      `json:"docNationality"`
	DocIssuingCountry      string      `json:"docIssuingCountry"`
	DocTemporaryAddress    string      `json:"docTemporaryAddress"`
	DocBirthName           string      `json:"docBirthName"`
	BirthPlace             string      `json:"birthPlace"`
	Authority              string      `json:"authority"`
	Address                string      `json:"address"`
	MotherMaidenName       string      `json:"mothersMaidenName"`
	DriverLicenseCategory  string      `json:"driverLicenseCategory"`
	ManuallyDataChanged    bool        `json:"manuallyDataChanged"`
	FullName               string      `json:"fullName"`
	OrgFirstName           string      `json:"orgFirstName"`
	OrgLastName            string      `json:"orgLastName"`
	OrgNationality         string      `json:"orgNationality"`
	OrgBirthPlace          string      `json:"orgBirthPlace"`
	OrgAuthority           string      `json:"orgAuthority"`
	OrgAddress             string      `json:"orgAddress"`
	OrgTemporaryAddress    string      `json:"orgTemporaryAddress"`
	OrgMothersMaidenName   string      `json:"orgMothersMaidenName"`
	OrgBirthName           string      `json:"orgBirthName"`
	SelectedCountry        string      `json:"selectedCountry"`
	AgeEstimate            string      `json:"ageEstimate"`
	ClientIpProxyRiskLevel string      `json:"clientIpProxyRiskLevel"`
	DuplicateFaces         []string    `json:"duplicateFaces"`
	DuplicateDocFaces      []string    `json:"duplicateDocFaces"`
	AddressVerification    interface{} `json:"addressVerification"`
	AdditionalData         interface{} `json:"additionalData"`
	ScanRef                string      `json:"scanRef"`
	ClientID               string      `json:"clientId"`
}
