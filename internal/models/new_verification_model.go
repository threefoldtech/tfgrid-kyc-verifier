package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TODO: switch to this model
type Platform string

const (
	PlatformPC        Platform = "PC"
	PlatformMobile    Platform = "MOBILE"
	PlatformTablet    Platform = "TABLET"
	PlatformMobileApp Platform = "MOBILE_APP"
	PlatformMobileSDK Platform = "MOBILE_SDK"
	PlatformOther     Platform = "OTHER"
)

type OverallStatus string

const (
	StatusApproved  OverallStatus = "APPROVED"
	StatusDenied    OverallStatus = "DENIED"
	StatusSuspected OverallStatus = "SUSPECTED"
	StatusReviewing OverallStatus = "REVIEWING"
	StatusExpired   OverallStatus = "EXPIRED"
	StatusActive    OverallStatus = "ACTIVE"
	StatusDeleted   OverallStatus = "DELETED"
	StatusArchived  OverallStatus = "ARCHIVED"
)

type SuspicionReason string

// Define constants for SuspicionReason

type FraudTag string

// Define constants for FraudTag

type MismatchTag string

// Define constants for MismatchTag

type FaceStatus string

const (
	FaceMatch        FaceStatus = "FACE_MATCH"
	FaceNotChecked   FaceStatus = "FACE_NOT_CHECKED"
	FaceMismatch     FaceStatus = "FACE_MISMATCH"
	NoFaceFound      FaceStatus = "NO_FACE_FOUND"
	TooManyFaces     FaceStatus = "TOO_MANY_FACES"
	FaceTooBlurry    FaceStatus = "FACE_TOO_BLURRY"
	FaceGlared       FaceStatus = "FACE_GLARED"
	FaceUncertain    FaceStatus = "FACE_UNCERTAIN"
	FaceNotAnalysed  FaceStatus = "FACE_NOT_ANALYSED"
	FaceError        FaceStatus = "FACE_ERROR"
	AutoUnverifiable FaceStatus = "AUTO_UNVERIFIABLE"
	FakeFace         FaceStatus = "FAKE_FACE"
)

type DocumentStatus string

// Define constants for DocumentStatus

type VerificationStatus string

const (
	StatusVerified          VerificationStatus = "VERIFIED"
	StatusPartiallyVerified VerificationStatus = "PARTIALLY_VERIFIED"
	StatusUnverified        VerificationStatus = "UNVERIFIED"
)

type Quality string

const (
	QualityExcellent Quality = "EXCELLENT"
	QualityGood      Quality = "GOOD"
	QualityAverage   Quality = "AVERAGE"
	QualityPoor      Quality = "POOR"
	QualityBad       Quality = "BAD"
)

type QuestionType string

const (
	QuestionTypeCheckbox    QuestionType = "CHECKBOX"
	QuestionTypeColor       QuestionType = "COLOR"
	QuestionTypeCountry     QuestionType = "COUNTRY"
	QuestionTypeDate        QuestionType = "DATE"
	QuestionTypeDateTime    QuestionType = "DATETIME"
	QuestionTypeEmail       QuestionType = "EMAIL"
	QuestionTypeFile        QuestionType = "FILE"
	QuestionTypeFloat       QuestionType = "FLOAT"
	QuestionTypeList        QuestionType = "LIST"
	QuestionTypeInteger     QuestionType = "INTEGER"
	QuestionTypePassword    QuestionType = "PASSWORD"
	QuestionTypeRadio       QuestionType = "RADIO"
	QuestionTypeSelect      QuestionType = "SELECT"
	QuestionTypeSelectMulti QuestionType = "SELECT_MULTI"
	QuestionTypeTel         QuestionType = "TEL"
	QuestionTypeText        QuestionType = "TEXT"
	QuestionTypeTextArea    QuestionType = "TEXT_AREA"
	QuestionTypeTime        QuestionType = "TIME"
	QuestionTypeURL         QuestionType = "URL"
)

type Verification_ struct {
	ID                   primitive.ObjectID     `bson:"_id,omitempty"`
	Final                bool                   `bson:"final"`
	Platform             Platform               `bson:"platform"`
	Status               Status_                `bson:"status"`
	Data                 PersonData             `bson:"data"`
	FileUrls             map[string]string      `bson:"fileUrls"`
	ScanRef              string                 `bson:"scanRef"`
	ClientID             string                 `bson:"clientId"`
	CompanyID            string                 `bson:"companyId"`
	BeneficiaryID        string                 `bson:"beneficiaryId"`
	StartTime            int64                  `bson:"startTime"`
	FinishTime           int64                  `bson:"finishTime"`
	ClientIP             string                 `bson:"clientIp"`
	ClientIPCountry      string                 `bson:"clientIpCountry"`
	ClientLocation       string                 `bson:"clientLocation"`
	QuestionnaireAnswers *QuestionnaireAnswers  `bson:"questionnaireAnswers,omitempty"`
	AdditionalSteps      map[string]interface{} `bson:"additionalSteps,omitempty"`
	UtilityData          []string               `bson:"utilityData,omitempty"`
}

type Status_ struct {
	Overall          OverallStatus     `bson:"overall"`
	SuspicionReasons []SuspicionReason `bson:"suspicionReasons"`
	DenyReasons      []string          `bson:"denyReasons"`
	FraudTags        []FraudTag        `bson:"fraudTags"`
	MismatchTags     []MismatchTag     `bson:"mismatchTags"`
	AutoFace         FaceStatus        `bson:"autoFace"`
	ManualFace       FaceStatus        `bson:"manualFace"`
	AutoDocument     DocumentStatus    `bson:"autoDocument"`
	ManualDocument   DocumentStatus    `bson:"manualDocument"`
}

type PersonData_ struct {
	// Add fields based on the schema
	Address  string             `bson:"address"`
	Status   VerificationStatus `bson:"status"`
	Accuracy *int               `bson:"accuracy,omitempty"`
	Quality  *Quality           `bson:"quality,omitempty"`
}

type QuestionnaireAnswers struct {
	Title    string    `bson:"title"`
	Sections []Section `bson:"sections"`
}

type Section struct {
	Title     string     `bson:"title"`
	Questions []Question `bson:"questions"`
}

type Question struct {
	Key   string       `bson:"key"`
	Title string       `bson:"title"`
	Type  QuestionType `bson:"type"`
	Value string       `bson:"value"`
}
