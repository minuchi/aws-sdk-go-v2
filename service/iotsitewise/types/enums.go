// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AggregateType string

// Enum values for AggregateType
const (
	AggregateTypeAverage           AggregateType = "AVERAGE"
	AggregateTypeCount             AggregateType = "COUNT"
	AggregateTypeMaximum           AggregateType = "MAXIMUM"
	AggregateTypeMinimum           AggregateType = "MINIMUM"
	AggregateTypeSum               AggregateType = "SUM"
	AggregateTypeStandardDeviation AggregateType = "STANDARD_DEVIATION"
)

// Values returns all known values for AggregateType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AggregateType) Values() []AggregateType {
	return []AggregateType{
		"AVERAGE",
		"COUNT",
		"MAXIMUM",
		"MINIMUM",
		"SUM",
		"STANDARD_DEVIATION",
	}
}

type AssetErrorCode string

// Enum values for AssetErrorCode
const (
	AssetErrorCodeInternalFailure AssetErrorCode = "INTERNAL_FAILURE"
)

// Values returns all known values for AssetErrorCode. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AssetErrorCode) Values() []AssetErrorCode {
	return []AssetErrorCode{
		"INTERNAL_FAILURE",
	}
}

type AssetModelState string

// Enum values for AssetModelState
const (
	AssetModelStateCreating    AssetModelState = "CREATING"
	AssetModelStateActive      AssetModelState = "ACTIVE"
	AssetModelStateUpdating    AssetModelState = "UPDATING"
	AssetModelStatePropagating AssetModelState = "PROPAGATING"
	AssetModelStateDeleting    AssetModelState = "DELETING"
	AssetModelStateFailed      AssetModelState = "FAILED"
)

// Values returns all known values for AssetModelState. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AssetModelState) Values() []AssetModelState {
	return []AssetModelState{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"PROPAGATING",
		"DELETING",
		"FAILED",
	}
}

type AssetRelationshipType string

// Enum values for AssetRelationshipType
const (
	AssetRelationshipTypeHierarchy AssetRelationshipType = "HIERARCHY"
)

// Values returns all known values for AssetRelationshipType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AssetRelationshipType) Values() []AssetRelationshipType {
	return []AssetRelationshipType{
		"HIERARCHY",
	}
}

type AssetState string

// Enum values for AssetState
const (
	AssetStateCreating AssetState = "CREATING"
	AssetStateActive   AssetState = "ACTIVE"
	AssetStateUpdating AssetState = "UPDATING"
	AssetStateDeleting AssetState = "DELETING"
	AssetStateFailed   AssetState = "FAILED"
)

// Values returns all known values for AssetState. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (AssetState) Values() []AssetState {
	return []AssetState{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"FAILED",
	}
}

type AuthMode string

// Enum values for AuthMode
const (
	AuthModeIam AuthMode = "IAM"
	AuthModeSso AuthMode = "SSO"
)

// Values returns all known values for AuthMode. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (AuthMode) Values() []AuthMode {
	return []AuthMode{
		"IAM",
		"SSO",
	}
}

type BatchEntryCompletionStatus string

// Enum values for BatchEntryCompletionStatus
const (
	BatchEntryCompletionStatusSuccess BatchEntryCompletionStatus = "SUCCESS"
	BatchEntryCompletionStatusError   BatchEntryCompletionStatus = "ERROR"
)

// Values returns all known values for BatchEntryCompletionStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (BatchEntryCompletionStatus) Values() []BatchEntryCompletionStatus {
	return []BatchEntryCompletionStatus{
		"SUCCESS",
		"ERROR",
	}
}

type BatchGetAssetPropertyAggregatesErrorCode string

// Enum values for BatchGetAssetPropertyAggregatesErrorCode
const (
	BatchGetAssetPropertyAggregatesErrorCodeResourceNotFoundException BatchGetAssetPropertyAggregatesErrorCode = "ResourceNotFoundException"
	BatchGetAssetPropertyAggregatesErrorCodeInvalidRequestException   BatchGetAssetPropertyAggregatesErrorCode = "InvalidRequestException"
	BatchGetAssetPropertyAggregatesErrorCodeAccessDeniedException     BatchGetAssetPropertyAggregatesErrorCode = "AccessDeniedException"
)

// Values returns all known values for BatchGetAssetPropertyAggregatesErrorCode.
// Note that this can be expanded in the future, and so it is only as up to date as
// the client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (BatchGetAssetPropertyAggregatesErrorCode) Values() []BatchGetAssetPropertyAggregatesErrorCode {
	return []BatchGetAssetPropertyAggregatesErrorCode{
		"ResourceNotFoundException",
		"InvalidRequestException",
		"AccessDeniedException",
	}
}

type BatchGetAssetPropertyValueErrorCode string

// Enum values for BatchGetAssetPropertyValueErrorCode
const (
	BatchGetAssetPropertyValueErrorCodeResourceNotFoundException BatchGetAssetPropertyValueErrorCode = "ResourceNotFoundException"
	BatchGetAssetPropertyValueErrorCodeInvalidRequestException   BatchGetAssetPropertyValueErrorCode = "InvalidRequestException"
	BatchGetAssetPropertyValueErrorCodeAccessDeniedException     BatchGetAssetPropertyValueErrorCode = "AccessDeniedException"
)

// Values returns all known values for BatchGetAssetPropertyValueErrorCode. Note
// that this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (BatchGetAssetPropertyValueErrorCode) Values() []BatchGetAssetPropertyValueErrorCode {
	return []BatchGetAssetPropertyValueErrorCode{
		"ResourceNotFoundException",
		"InvalidRequestException",
		"AccessDeniedException",
	}
}

type BatchGetAssetPropertyValueHistoryErrorCode string

// Enum values for BatchGetAssetPropertyValueHistoryErrorCode
const (
	BatchGetAssetPropertyValueHistoryErrorCodeResourceNotFoundException BatchGetAssetPropertyValueHistoryErrorCode = "ResourceNotFoundException"
	BatchGetAssetPropertyValueHistoryErrorCodeInvalidRequestException   BatchGetAssetPropertyValueHistoryErrorCode = "InvalidRequestException"
	BatchGetAssetPropertyValueHistoryErrorCodeAccessDeniedException     BatchGetAssetPropertyValueHistoryErrorCode = "AccessDeniedException"
)

// Values returns all known values for BatchGetAssetPropertyValueHistoryErrorCode.
// Note that this can be expanded in the future, and so it is only as up to date as
// the client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (BatchGetAssetPropertyValueHistoryErrorCode) Values() []BatchGetAssetPropertyValueHistoryErrorCode {
	return []BatchGetAssetPropertyValueHistoryErrorCode{
		"ResourceNotFoundException",
		"InvalidRequestException",
		"AccessDeniedException",
	}
}

type BatchPutAssetPropertyValueErrorCode string

// Enum values for BatchPutAssetPropertyValueErrorCode
const (
	BatchPutAssetPropertyValueErrorCodeResourceNotFoundException     BatchPutAssetPropertyValueErrorCode = "ResourceNotFoundException"
	BatchPutAssetPropertyValueErrorCodeInvalidRequestException       BatchPutAssetPropertyValueErrorCode = "InvalidRequestException"
	BatchPutAssetPropertyValueErrorCodeInternalFailureException      BatchPutAssetPropertyValueErrorCode = "InternalFailureException"
	BatchPutAssetPropertyValueErrorCodeServiceUnavailableException   BatchPutAssetPropertyValueErrorCode = "ServiceUnavailableException"
	BatchPutAssetPropertyValueErrorCodeThrottlingException           BatchPutAssetPropertyValueErrorCode = "ThrottlingException"
	BatchPutAssetPropertyValueErrorCodeLimitExceededException        BatchPutAssetPropertyValueErrorCode = "LimitExceededException"
	BatchPutAssetPropertyValueErrorCodeConflictingOperationException BatchPutAssetPropertyValueErrorCode = "ConflictingOperationException"
	BatchPutAssetPropertyValueErrorCodeTimestampOutOfRangeException  BatchPutAssetPropertyValueErrorCode = "TimestampOutOfRangeException"
	BatchPutAssetPropertyValueErrorCodeAccessDeniedException         BatchPutAssetPropertyValueErrorCode = "AccessDeniedException"
)

// Values returns all known values for BatchPutAssetPropertyValueErrorCode. Note
// that this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (BatchPutAssetPropertyValueErrorCode) Values() []BatchPutAssetPropertyValueErrorCode {
	return []BatchPutAssetPropertyValueErrorCode{
		"ResourceNotFoundException",
		"InvalidRequestException",
		"InternalFailureException",
		"ServiceUnavailableException",
		"ThrottlingException",
		"LimitExceededException",
		"ConflictingOperationException",
		"TimestampOutOfRangeException",
		"AccessDeniedException",
	}
}

type CapabilitySyncStatus string

// Enum values for CapabilitySyncStatus
const (
	CapabilitySyncStatusInSync     CapabilitySyncStatus = "IN_SYNC"
	CapabilitySyncStatusOutOfSync  CapabilitySyncStatus = "OUT_OF_SYNC"
	CapabilitySyncStatusSyncFailed CapabilitySyncStatus = "SYNC_FAILED"
	CapabilitySyncStatusUnknown    CapabilitySyncStatus = "UNKNOWN"
)

// Values returns all known values for CapabilitySyncStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (CapabilitySyncStatus) Values() []CapabilitySyncStatus {
	return []CapabilitySyncStatus{
		"IN_SYNC",
		"OUT_OF_SYNC",
		"SYNC_FAILED",
		"UNKNOWN",
	}
}

type ComputeLocation string

// Enum values for ComputeLocation
const (
	ComputeLocationEdge  ComputeLocation = "EDGE"
	ComputeLocationCloud ComputeLocation = "CLOUD"
)

// Values returns all known values for ComputeLocation. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ComputeLocation) Values() []ComputeLocation {
	return []ComputeLocation{
		"EDGE",
		"CLOUD",
	}
}

type ConfigurationState string

// Enum values for ConfigurationState
const (
	ConfigurationStateActive           ConfigurationState = "ACTIVE"
	ConfigurationStateUpdateInProgress ConfigurationState = "UPDATE_IN_PROGRESS"
	ConfigurationStateUpdateFailed     ConfigurationState = "UPDATE_FAILED"
)

// Values returns all known values for ConfigurationState. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ConfigurationState) Values() []ConfigurationState {
	return []ConfigurationState{
		"ACTIVE",
		"UPDATE_IN_PROGRESS",
		"UPDATE_FAILED",
	}
}

type DetailedErrorCode string

// Enum values for DetailedErrorCode
const (
	DetailedErrorCodeIncompatibleComputeLocation         DetailedErrorCode = "INCOMPATIBLE_COMPUTE_LOCATION"
	DetailedErrorCodeIncompatibleForwardingConfiguration DetailedErrorCode = "INCOMPATIBLE_FORWARDING_CONFIGURATION"
)

// Values returns all known values for DetailedErrorCode. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DetailedErrorCode) Values() []DetailedErrorCode {
	return []DetailedErrorCode{
		"INCOMPATIBLE_COMPUTE_LOCATION",
		"INCOMPATIBLE_FORWARDING_CONFIGURATION",
	}
}

type DisassociatedDataStorageState string

// Enum values for DisassociatedDataStorageState
const (
	DisassociatedDataStorageStateEnabled  DisassociatedDataStorageState = "ENABLED"
	DisassociatedDataStorageStateDisabled DisassociatedDataStorageState = "DISABLED"
)

// Values returns all known values for DisassociatedDataStorageState. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (DisassociatedDataStorageState) Values() []DisassociatedDataStorageState {
	return []DisassociatedDataStorageState{
		"ENABLED",
		"DISABLED",
	}
}

type EncryptionType string

// Enum values for EncryptionType
const (
	EncryptionTypeSitewiseDefaultEncryption EncryptionType = "SITEWISE_DEFAULT_ENCRYPTION"
	EncryptionTypeKmsBasedEncryption        EncryptionType = "KMS_BASED_ENCRYPTION"
)

// Values returns all known values for EncryptionType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (EncryptionType) Values() []EncryptionType {
	return []EncryptionType{
		"SITEWISE_DEFAULT_ENCRYPTION",
		"KMS_BASED_ENCRYPTION",
	}
}

type ErrorCode string

// Enum values for ErrorCode
const (
	ErrorCodeValidationError ErrorCode = "VALIDATION_ERROR"
	ErrorCodeInternalFailure ErrorCode = "INTERNAL_FAILURE"
)

// Values returns all known values for ErrorCode. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (ErrorCode) Values() []ErrorCode {
	return []ErrorCode{
		"VALIDATION_ERROR",
		"INTERNAL_FAILURE",
	}
}

type ForwardingConfigState string

// Enum values for ForwardingConfigState
const (
	ForwardingConfigStateDisabled ForwardingConfigState = "DISABLED"
	ForwardingConfigStateEnabled  ForwardingConfigState = "ENABLED"
)

// Values returns all known values for ForwardingConfigState. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ForwardingConfigState) Values() []ForwardingConfigState {
	return []ForwardingConfigState{
		"DISABLED",
		"ENABLED",
	}
}

type IdentityType string

// Enum values for IdentityType
const (
	IdentityTypeUser  IdentityType = "USER"
	IdentityTypeGroup IdentityType = "GROUP"
	IdentityTypeIam   IdentityType = "IAM"
)

// Values returns all known values for IdentityType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (IdentityType) Values() []IdentityType {
	return []IdentityType{
		"USER",
		"GROUP",
		"IAM",
	}
}

type ImageFileType string

// Enum values for ImageFileType
const (
	ImageFileTypePng ImageFileType = "PNG"
)

// Values returns all known values for ImageFileType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ImageFileType) Values() []ImageFileType {
	return []ImageFileType{
		"PNG",
	}
}

type ListAssetsFilter string

// Enum values for ListAssetsFilter
const (
	ListAssetsFilterAll      ListAssetsFilter = "ALL"
	ListAssetsFilterTopLevel ListAssetsFilter = "TOP_LEVEL"
)

// Values returns all known values for ListAssetsFilter. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ListAssetsFilter) Values() []ListAssetsFilter {
	return []ListAssetsFilter{
		"ALL",
		"TOP_LEVEL",
	}
}

type ListTimeSeriesType string

// Enum values for ListTimeSeriesType
const (
	ListTimeSeriesTypeAssociated    ListTimeSeriesType = "ASSOCIATED"
	ListTimeSeriesTypeDisassociated ListTimeSeriesType = "DISASSOCIATED"
)

// Values returns all known values for ListTimeSeriesType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ListTimeSeriesType) Values() []ListTimeSeriesType {
	return []ListTimeSeriesType{
		"ASSOCIATED",
		"DISASSOCIATED",
	}
}

type LoggingLevel string

// Enum values for LoggingLevel
const (
	LoggingLevelError LoggingLevel = "ERROR"
	LoggingLevelInfo  LoggingLevel = "INFO"
	LoggingLevelOff   LoggingLevel = "OFF"
)

// Values returns all known values for LoggingLevel. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (LoggingLevel) Values() []LoggingLevel {
	return []LoggingLevel{
		"ERROR",
		"INFO",
		"OFF",
	}
}

type MonitorErrorCode string

// Enum values for MonitorErrorCode
const (
	MonitorErrorCodeInternalFailure MonitorErrorCode = "INTERNAL_FAILURE"
	MonitorErrorCodeValidationError MonitorErrorCode = "VALIDATION_ERROR"
	MonitorErrorCodeLimitExceeded   MonitorErrorCode = "LIMIT_EXCEEDED"
)

// Values returns all known values for MonitorErrorCode. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (MonitorErrorCode) Values() []MonitorErrorCode {
	return []MonitorErrorCode{
		"INTERNAL_FAILURE",
		"VALIDATION_ERROR",
		"LIMIT_EXCEEDED",
	}
}

type Permission string

// Enum values for Permission
const (
	PermissionAdministrator Permission = "ADMINISTRATOR"
	PermissionViewer        Permission = "VIEWER"
)

// Values returns all known values for Permission. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (Permission) Values() []Permission {
	return []Permission{
		"ADMINISTRATOR",
		"VIEWER",
	}
}

type PortalState string

// Enum values for PortalState
const (
	PortalStateCreating PortalState = "CREATING"
	PortalStateUpdating PortalState = "UPDATING"
	PortalStateDeleting PortalState = "DELETING"
	PortalStateActive   PortalState = "ACTIVE"
	PortalStateFailed   PortalState = "FAILED"
)

// Values returns all known values for PortalState. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (PortalState) Values() []PortalState {
	return []PortalState{
		"CREATING",
		"UPDATING",
		"DELETING",
		"ACTIVE",
		"FAILED",
	}
}

type PropertyDataType string

// Enum values for PropertyDataType
const (
	PropertyDataTypeString  PropertyDataType = "STRING"
	PropertyDataTypeInteger PropertyDataType = "INTEGER"
	PropertyDataTypeDouble  PropertyDataType = "DOUBLE"
	PropertyDataTypeBoolean PropertyDataType = "BOOLEAN"
	PropertyDataTypeStruct  PropertyDataType = "STRUCT"
)

// Values returns all known values for PropertyDataType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PropertyDataType) Values() []PropertyDataType {
	return []PropertyDataType{
		"STRING",
		"INTEGER",
		"DOUBLE",
		"BOOLEAN",
		"STRUCT",
	}
}

type PropertyNotificationState string

// Enum values for PropertyNotificationState
const (
	PropertyNotificationStateEnabled  PropertyNotificationState = "ENABLED"
	PropertyNotificationStateDisabled PropertyNotificationState = "DISABLED"
)

// Values returns all known values for PropertyNotificationState. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (PropertyNotificationState) Values() []PropertyNotificationState {
	return []PropertyNotificationState{
		"ENABLED",
		"DISABLED",
	}
}

type Quality string

// Enum values for Quality
const (
	QualityGood      Quality = "GOOD"
	QualityBad       Quality = "BAD"
	QualityUncertain Quality = "UNCERTAIN"
)

// Values returns all known values for Quality. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Quality) Values() []Quality {
	return []Quality{
		"GOOD",
		"BAD",
		"UNCERTAIN",
	}
}

type ResourceType string

// Enum values for ResourceType
const (
	ResourceTypePortal  ResourceType = "PORTAL"
	ResourceTypeProject ResourceType = "PROJECT"
)

// Values returns all known values for ResourceType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ResourceType) Values() []ResourceType {
	return []ResourceType{
		"PORTAL",
		"PROJECT",
	}
}

type StorageType string

// Enum values for StorageType
const (
	StorageTypeSitewiseDefaultStorage StorageType = "SITEWISE_DEFAULT_STORAGE"
	StorageTypeMultiLayerStorage      StorageType = "MULTI_LAYER_STORAGE"
)

// Values returns all known values for StorageType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (StorageType) Values() []StorageType {
	return []StorageType{
		"SITEWISE_DEFAULT_STORAGE",
		"MULTI_LAYER_STORAGE",
	}
}

type TimeOrdering string

// Enum values for TimeOrdering
const (
	TimeOrderingAscending  TimeOrdering = "ASCENDING"
	TimeOrderingDescending TimeOrdering = "DESCENDING"
)

// Values returns all known values for TimeOrdering. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (TimeOrdering) Values() []TimeOrdering {
	return []TimeOrdering{
		"ASCENDING",
		"DESCENDING",
	}
}

type TraversalDirection string

// Enum values for TraversalDirection
const (
	TraversalDirectionParent TraversalDirection = "PARENT"
	TraversalDirectionChild  TraversalDirection = "CHILD"
)

// Values returns all known values for TraversalDirection. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (TraversalDirection) Values() []TraversalDirection {
	return []TraversalDirection{
		"PARENT",
		"CHILD",
	}
}

type TraversalType string

// Enum values for TraversalType
const (
	TraversalTypePathToRoot TraversalType = "PATH_TO_ROOT"
)

// Values returns all known values for TraversalType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (TraversalType) Values() []TraversalType {
	return []TraversalType{
		"PATH_TO_ROOT",
	}
}