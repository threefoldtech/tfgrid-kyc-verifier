package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Verification struct {
	ID                    primitive.ObjectID `bson:"_id,omitempty"`
	Final                 *bool              `bson:"final"`
	Platform              string             `bson:"platform"`
	Status                Status             `bson:"status"`
	Data                  PersonData         `bson:"data"`
	FileUrls              map[string]string  `bson:"fileUrls"`
	AdditionalStepPdfUrls map[string]string  `bson:"additionalStepPdfUrls"`
	AML                   []AMLCheck         `bson:"AML"`
	LID                   interface{}        `bson:"LID"`
	ScanRef               string             `bson:"scanRef"`
	ExternalRef           string             `bson:"externalRef"`
	ClientID              string             `bson:"clientId"`
	StartTime             int64              `bson:"startTime"`
	FinishTime            int64              `bson:"finishTime"`
	ClientIP              string             `bson:"clientIp"`
	ClientIPCountry       string             `bson:"clientIpCountry"`
	ClientLocation        string             `bson:"clientLocation"`
	ManualAddress         interface{}        `bson:"manualAddress"`
	ManualAddressMatch    bool               `bson:"manualAddressMatch"`
	RegistryCenterCheck   interface{}        `bson:"registryCenterCheck"`
	AddressVerification   interface{}        `bson:"addressVerification"`
	QuestionnaireAnswers  interface{}        `bson:"questionnaireAnswers"`
	CompanyID             interface{}        `bson:"companyId"`
	BeneficiaryID         interface{}        `bson:"beneficiaryId"`
	AdditionalSteps       map[string]string  `bson:"additionalSteps"`
	CreatedAt             time.Time          `bson:"createdAt"`
}

type Overall string

const (
	OverallApproved  Overall = "APPROVED"
	OverallDenied    Overall = "DENIED"
	OverallSuspected Overall = "SUSPECTED"
	OverallExpired   Overall = "EXPIRED"
)

type Status struct {
	Overall            Overall  `bson:"overall"`
	SuspicionReasons   []string `bson:"suspicionReasons"`
	DenyReasons        []string `bson:"denyReasons"`
	FraudTags          []string `bson:"fraudTags"`
	MismatchTags       []string `bson:"mismatchTags"`
	AutoFace           string   `bson:"autoFace"`
	ManualFace         string   `bson:"manualFace"`
	AutoDocument       string   `bson:"autoDocument"`
	ManualDocument     string   `bson:"manualDocument"`
	AdditionalSteps    string   `bson:"additionalSteps"`
	AMLResultClass     string   `bson:"amlResultClass"`
	PEPSStatus         string   `bson:"pepsStatus"`
	SanctionsStatus    string   `bson:"sanctionsStatus"`
	AdverseMediaStatus string   `bson:"adverseMediaStatus"`
}

type DocumentType string

const (
	ID_CARD                    DocumentType = "ID_CARD"
	PASSPORT                   DocumentType = "PASSPORT"
	RESIDENCE_PERMIT           DocumentType = "RESIDENCE_PERMIT"
	DRIVER_LICENSE             DocumentType = "DRIVER_LICENSE"
	PAN_CARD                   DocumentType = "PAN_CARD"
	AADHAAR                    DocumentType = "AADHAAR"
	OTHER                      DocumentType = "OTHER"
	VISA                       DocumentType = "VISA"
	BORDER_CROSSING            DocumentType = "BORDER_CROSSING"
	ASYLUM                     DocumentType = "ASYLUM"
	NATIONAL_PASSPORT          DocumentType = "NATIONAL_PASSPORT"
	PROVISIONAL_DRIVER_LICENSE DocumentType = "PROVISIONAL_DRIVER_LICENSE"
	VOTER_CARD                 DocumentType = "VOTER_CARD"
	OLD_ID_CARD                DocumentType = "OLD_ID_CARD"
	TRAVEL_CARD                DocumentType = "TRAVEL_CARD"
	PHOTO_CARD                 DocumentType = "PHOTO_CARD"
	MILITARY_CARD              DocumentType = "MILITARY_CARD"
	PROOF_OF_AGE_CARD          DocumentType = "PROOF_OF_AGE_CARD"
	DIPLOMATIC_ID              DocumentType = "DIPLOMATIC_ID"
)

type Sex string

const (
	MALE      Sex = "MALE"
	FEMALE    Sex = "FEMALE"
	UNDEFINED Sex = "UNDEFINED"
)

type AgeEstimate string

const (
	UNDER_13 AgeEstimate = "UNDER_13"
	OVER_13  AgeEstimate = "OVER_13"
	OVER_18  AgeEstimate = "OVER_18"
	OVER_22  AgeEstimate = "OVER_22"
	OVER_25  AgeEstimate = "OVER_25"
	OVER_30  AgeEstimate = "OVER_30"
)

type PersonData struct {
	DocFirstName           string       `bson:"docFirstName"`
	DocLastName            string       `bson:"docLastName"`
	DocNumber              string       `bson:"docNumber"`
	DocPersonalCode        string       `bson:"docPersonalCode"`
	DocExpiry              string       `bson:"docExpiry"`
	DocDOB                 string       `bson:"docDob"`
	DocDateOfIssue         string       `bson:"docDateOfIssue"`
	DocType                DocumentType `bson:"docType"`
	DocSex                 Sex          `bson:"docSex"`
	DocNationality         string       `bson:"docNationality"`
	DocIssuingCountry      string       `bson:"docIssuingCountry"`
	BirthPlace             string       `bson:"birthPlace"`
	Authority              string       `bson:"authority"`
	Address                string       `bson:"address"`
	DocTemporaryAddress    string       `bson:"docTemporaryAddress"`
	MothersMaidenName      string       `bson:"mothersMaidenName"`
	DocBirthName           string       `bson:"docBirthName"`
	DriverLicenseCategory  string       `bson:"driverLicenseCategory"`
	ManuallyDataChanged    bool         `bson:"manuallyDataChanged"`
	FullName               string       `bson:"fullName"`
	SelectedCountry        string       `bson:"selectedCountry"`
	OrgFirstName           string       `bson:"orgFirstName"`
	OrgLastName            string       `bson:"orgLastName"`
	OrgNationality         string       `bson:"orgNationality"`
	OrgBirthPlace          string       `bson:"orgBirthPlace"`
	OrgAuthority           string       `bson:"orgAuthority"`
	OrgAddress             string       `bson:"orgAddress"`
	OrgTemporaryAddress    string       `bson:"orgTemporaryAddress"`
	OrgMothersMaidenName   string       `bson:"orgMothersMaidenName"`
	OrgBirthName           string       `bson:"orgBirthName"`
	AgeEstimate            AgeEstimate  `bson:"ageEstimate"`
	ClientIPProxyRiskLevel string       `bson:"clientIpProxyRiskLevel"`
	DuplicateFaces         []string     `bson:"duplicateFaces"`
	DuplicateDocFaces      []string     `bson:"duplicateDocFaces"`
	AdditionalData         interface{}  `bson:"additionalData"`
}

type AMLCheck struct {
	Status           AMLStatus `bson:"status"`
	Data             []AMLData `bson:"data"`
	ServiceName      string    `bson:"serviceName"`
	ServiceGroupType string    `bson:"serviceGroupType"`
	UID              string    `bson:"uid"`
	ErrorMessage     string    `bson:"errorMessage"`
}

type AMLStatus struct {
	ServiceSuspected bool   `bson:"serviceSuspected"`
	ServiceUsed      bool   `bson:"serviceUsed"`
	ServiceFound     bool   `bson:"serviceFound"`
	CheckSuccessful  bool   `bson:"checkSuccessful"`
	OverallStatus    string `bson:"overallStatus"`
}

type AMLData struct {
	Name        string   `bson:"name"`
	Surname     string   `bson:"surname"`
	Nationality string   `bson:"nationality"`
	DOB         string   `bson:"dob"`
	Suspicion   string   `bson:"suspicion"`
	Reason      string   `bson:"reason"`
	ListNumber  string   `bson:"listNumber"`
	ListName    string   `bson:"listName"`
	Score       *float64 `bson:"score"`
	LastUpdate  *string  `bson:"lastUpdate"`
	IsPerson    *bool    `bson:"isPerson"`
	IsActive    *bool    `bson:"isActive"`
	CheckDate   string   `bson:"checkDate"`
}
