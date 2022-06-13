// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type MediaCapabilities string

// Enum values for MediaCapabilities
const (
	MediaCapabilitiesSendReceive MediaCapabilities = "SendReceive"
	MediaCapabilitiesSend        MediaCapabilities = "Send"
	MediaCapabilitiesReceive     MediaCapabilities = "Receive"
	MediaCapabilitiesNone        MediaCapabilities = "None"
)

// Values returns all known values for MediaCapabilities. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (MediaCapabilities) Values() []MediaCapabilities {
	return []MediaCapabilities{
		"SendReceive",
		"Send",
		"Receive",
		"None",
	}
}

type MeetingFeatureStatus string

// Enum values for MeetingFeatureStatus
const (
	MeetingFeatureStatusAvailable   MeetingFeatureStatus = "AVAILABLE"
	MeetingFeatureStatusUnavailable MeetingFeatureStatus = "UNAVAILABLE"
)

// Values returns all known values for MeetingFeatureStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (MeetingFeatureStatus) Values() []MeetingFeatureStatus {
	return []MeetingFeatureStatus{
		"AVAILABLE",
		"UNAVAILABLE",
	}
}

type TranscribeContentIdentificationType string

// Enum values for TranscribeContentIdentificationType
const (
	TranscribeContentIdentificationTypePii TranscribeContentIdentificationType = "PII"
)

// Values returns all known values for TranscribeContentIdentificationType. Note
// that this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (TranscribeContentIdentificationType) Values() []TranscribeContentIdentificationType {
	return []TranscribeContentIdentificationType{
		"PII",
	}
}

type TranscribeContentRedactionType string

// Enum values for TranscribeContentRedactionType
const (
	TranscribeContentRedactionTypePii TranscribeContentRedactionType = "PII"
)

// Values returns all known values for TranscribeContentRedactionType. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (TranscribeContentRedactionType) Values() []TranscribeContentRedactionType {
	return []TranscribeContentRedactionType{
		"PII",
	}
}

type TranscribeLanguageCode string

// Enum values for TranscribeLanguageCode
const (
	TranscribeLanguageCodeEnUs TranscribeLanguageCode = "en-US"
	TranscribeLanguageCodeEnGb TranscribeLanguageCode = "en-GB"
	TranscribeLanguageCodeEsUs TranscribeLanguageCode = "es-US"
	TranscribeLanguageCodeFrCa TranscribeLanguageCode = "fr-CA"
	TranscribeLanguageCodeFrFr TranscribeLanguageCode = "fr-FR"
	TranscribeLanguageCodeEnAu TranscribeLanguageCode = "en-AU"
	TranscribeLanguageCodeItIt TranscribeLanguageCode = "it-IT"
	TranscribeLanguageCodeDeDe TranscribeLanguageCode = "de-DE"
	TranscribeLanguageCodePtBr TranscribeLanguageCode = "pt-BR"
	TranscribeLanguageCodeJaJp TranscribeLanguageCode = "ja-JP"
	TranscribeLanguageCodeKoKr TranscribeLanguageCode = "ko-KR"
	TranscribeLanguageCodeZhCn TranscribeLanguageCode = "zh-CN"
)

// Values returns all known values for TranscribeLanguageCode. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (TranscribeLanguageCode) Values() []TranscribeLanguageCode {
	return []TranscribeLanguageCode{
		"en-US",
		"en-GB",
		"es-US",
		"fr-CA",
		"fr-FR",
		"en-AU",
		"it-IT",
		"de-DE",
		"pt-BR",
		"ja-JP",
		"ko-KR",
		"zh-CN",
	}
}

type TranscribeMedicalContentIdentificationType string

// Enum values for TranscribeMedicalContentIdentificationType
const (
	TranscribeMedicalContentIdentificationTypePhi TranscribeMedicalContentIdentificationType = "PHI"
)

// Values returns all known values for TranscribeMedicalContentIdentificationType.
// Note that this can be expanded in the future, and so it is only as up to date as
// the client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (TranscribeMedicalContentIdentificationType) Values() []TranscribeMedicalContentIdentificationType {
	return []TranscribeMedicalContentIdentificationType{
		"PHI",
	}
}

type TranscribeMedicalLanguageCode string

// Enum values for TranscribeMedicalLanguageCode
const (
	TranscribeMedicalLanguageCodeEnUs TranscribeMedicalLanguageCode = "en-US"
)

// Values returns all known values for TranscribeMedicalLanguageCode. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (TranscribeMedicalLanguageCode) Values() []TranscribeMedicalLanguageCode {
	return []TranscribeMedicalLanguageCode{
		"en-US",
	}
}

type TranscribeMedicalRegion string

// Enum values for TranscribeMedicalRegion
const (
	TranscribeMedicalRegionUsEast1      TranscribeMedicalRegion = "us-east-1"
	TranscribeMedicalRegionUsEast2      TranscribeMedicalRegion = "us-east-2"
	TranscribeMedicalRegionUsWest2      TranscribeMedicalRegion = "us-west-2"
	TranscribeMedicalRegionApSoutheast2 TranscribeMedicalRegion = "ap-southeast-2"
	TranscribeMedicalRegionCaCentral1   TranscribeMedicalRegion = "ca-central-1"
	TranscribeMedicalRegionEuWest1      TranscribeMedicalRegion = "eu-west-1"
	TranscribeMedicalRegionAuto         TranscribeMedicalRegion = "auto"
)

// Values returns all known values for TranscribeMedicalRegion. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (TranscribeMedicalRegion) Values() []TranscribeMedicalRegion {
	return []TranscribeMedicalRegion{
		"us-east-1",
		"us-east-2",
		"us-west-2",
		"ap-southeast-2",
		"ca-central-1",
		"eu-west-1",
		"auto",
	}
}

type TranscribeMedicalSpecialty string

// Enum values for TranscribeMedicalSpecialty
const (
	TranscribeMedicalSpecialtyPrimarycare TranscribeMedicalSpecialty = "PRIMARYCARE"
	TranscribeMedicalSpecialtyCardiology  TranscribeMedicalSpecialty = "CARDIOLOGY"
	TranscribeMedicalSpecialtyNeurology   TranscribeMedicalSpecialty = "NEUROLOGY"
	TranscribeMedicalSpecialtyOncology    TranscribeMedicalSpecialty = "ONCOLOGY"
	TranscribeMedicalSpecialtyRadiology   TranscribeMedicalSpecialty = "RADIOLOGY"
	TranscribeMedicalSpecialtyUrology     TranscribeMedicalSpecialty = "UROLOGY"
)

// Values returns all known values for TranscribeMedicalSpecialty. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (TranscribeMedicalSpecialty) Values() []TranscribeMedicalSpecialty {
	return []TranscribeMedicalSpecialty{
		"PRIMARYCARE",
		"CARDIOLOGY",
		"NEUROLOGY",
		"ONCOLOGY",
		"RADIOLOGY",
		"UROLOGY",
	}
}

type TranscribeMedicalType string

// Enum values for TranscribeMedicalType
const (
	TranscribeMedicalTypeConversation TranscribeMedicalType = "CONVERSATION"
	TranscribeMedicalTypeDictation    TranscribeMedicalType = "DICTATION"
)

// Values returns all known values for TranscribeMedicalType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (TranscribeMedicalType) Values() []TranscribeMedicalType {
	return []TranscribeMedicalType{
		"CONVERSATION",
		"DICTATION",
	}
}

type TranscribePartialResultsStability string

// Enum values for TranscribePartialResultsStability
const (
	TranscribePartialResultsStabilityLow    TranscribePartialResultsStability = "low"
	TranscribePartialResultsStabilityMedium TranscribePartialResultsStability = "medium"
	TranscribePartialResultsStabilityHigh   TranscribePartialResultsStability = "high"
)

// Values returns all known values for TranscribePartialResultsStability. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (TranscribePartialResultsStability) Values() []TranscribePartialResultsStability {
	return []TranscribePartialResultsStability{
		"low",
		"medium",
		"high",
	}
}

type TranscribeRegion string

// Enum values for TranscribeRegion
const (
	TranscribeRegionUsEast2      TranscribeRegion = "us-east-2"
	TranscribeRegionUsEast1      TranscribeRegion = "us-east-1"
	TranscribeRegionUsWest2      TranscribeRegion = "us-west-2"
	TranscribeRegionApNortheast2 TranscribeRegion = "ap-northeast-2"
	TranscribeRegionApSoutheast2 TranscribeRegion = "ap-southeast-2"
	TranscribeRegionApNortheast1 TranscribeRegion = "ap-northeast-1"
	TranscribeRegionCaCentral1   TranscribeRegion = "ca-central-1"
	TranscribeRegionEuCentral1   TranscribeRegion = "eu-central-1"
	TranscribeRegionEuWest1      TranscribeRegion = "eu-west-1"
	TranscribeRegionEuWest2      TranscribeRegion = "eu-west-2"
	TranscribeRegionSaEast1      TranscribeRegion = "sa-east-1"
	TranscribeRegionAuto         TranscribeRegion = "auto"
	TranscribeRegionUsGovWest1   TranscribeRegion = "us-gov-west-1"
)

// Values returns all known values for TranscribeRegion. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (TranscribeRegion) Values() []TranscribeRegion {
	return []TranscribeRegion{
		"us-east-2",
		"us-east-1",
		"us-west-2",
		"ap-northeast-2",
		"ap-southeast-2",
		"ap-northeast-1",
		"ca-central-1",
		"eu-central-1",
		"eu-west-1",
		"eu-west-2",
		"sa-east-1",
		"auto",
		"us-gov-west-1",
	}
}

type TranscribeVocabularyFilterMethod string

// Enum values for TranscribeVocabularyFilterMethod
const (
	TranscribeVocabularyFilterMethodRemove TranscribeVocabularyFilterMethod = "remove"
	TranscribeVocabularyFilterMethodMask   TranscribeVocabularyFilterMethod = "mask"
	TranscribeVocabularyFilterMethodTag    TranscribeVocabularyFilterMethod = "tag"
)

// Values returns all known values for TranscribeVocabularyFilterMethod. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (TranscribeVocabularyFilterMethod) Values() []TranscribeVocabularyFilterMethod {
	return []TranscribeVocabularyFilterMethod{
		"remove",
		"mask",
		"tag",
	}
}