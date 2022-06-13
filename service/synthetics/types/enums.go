// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type CanaryRunState string

// Enum values for CanaryRunState
const (
	CanaryRunStateRunning CanaryRunState = "RUNNING"
	CanaryRunStatePassed  CanaryRunState = "PASSED"
	CanaryRunStateFailed  CanaryRunState = "FAILED"
)

// Values returns all known values for CanaryRunState. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (CanaryRunState) Values() []CanaryRunState {
	return []CanaryRunState{
		"RUNNING",
		"PASSED",
		"FAILED",
	}
}

type CanaryRunStateReasonCode string

// Enum values for CanaryRunStateReasonCode
const (
	CanaryRunStateReasonCodeCanaryFailure    CanaryRunStateReasonCode = "CANARY_FAILURE"
	CanaryRunStateReasonCodeExecutionFailure CanaryRunStateReasonCode = "EXECUTION_FAILURE"
)

// Values returns all known values for CanaryRunStateReasonCode. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (CanaryRunStateReasonCode) Values() []CanaryRunStateReasonCode {
	return []CanaryRunStateReasonCode{
		"CANARY_FAILURE",
		"EXECUTION_FAILURE",
	}
}

type CanaryState string

// Enum values for CanaryState
const (
	CanaryStateCreating CanaryState = "CREATING"
	CanaryStateReady    CanaryState = "READY"
	CanaryStateStarting CanaryState = "STARTING"
	CanaryStateRunning  CanaryState = "RUNNING"
	CanaryStateUpdating CanaryState = "UPDATING"
	CanaryStateStopping CanaryState = "STOPPING"
	CanaryStateStopped  CanaryState = "STOPPED"
	CanaryStateError    CanaryState = "ERROR"
	CanaryStateDeleting CanaryState = "DELETING"
)

// Values returns all known values for CanaryState. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (CanaryState) Values() []CanaryState {
	return []CanaryState{
		"CREATING",
		"READY",
		"STARTING",
		"RUNNING",
		"UPDATING",
		"STOPPING",
		"STOPPED",
		"ERROR",
		"DELETING",
	}
}

type CanaryStateReasonCode string

// Enum values for CanaryStateReasonCode
const (
	CanaryStateReasonCodeInvalidPermissions   CanaryStateReasonCode = "INVALID_PERMISSIONS"
	CanaryStateReasonCodeCreatePending        CanaryStateReasonCode = "CREATE_PENDING"
	CanaryStateReasonCodeCreateInProgress     CanaryStateReasonCode = "CREATE_IN_PROGRESS"
	CanaryStateReasonCodeCreateFailed         CanaryStateReasonCode = "CREATE_FAILED"
	CanaryStateReasonCodeUpdatePending        CanaryStateReasonCode = "UPDATE_PENDING"
	CanaryStateReasonCodeUpdateInProgress     CanaryStateReasonCode = "UPDATE_IN_PROGRESS"
	CanaryStateReasonCodeUpdateComplete       CanaryStateReasonCode = "UPDATE_COMPLETE"
	CanaryStateReasonCodeRollbackComplete     CanaryStateReasonCode = "ROLLBACK_COMPLETE"
	CanaryStateReasonCodeRollbackFailed       CanaryStateReasonCode = "ROLLBACK_FAILED"
	CanaryStateReasonCodeDeleteInProgress     CanaryStateReasonCode = "DELETE_IN_PROGRESS"
	CanaryStateReasonCodeDeleteFailed         CanaryStateReasonCode = "DELETE_FAILED"
	CanaryStateReasonCodeSyncDeleteInProgress CanaryStateReasonCode = "SYNC_DELETE_IN_PROGRESS"
)

// Values returns all known values for CanaryStateReasonCode. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (CanaryStateReasonCode) Values() []CanaryStateReasonCode {
	return []CanaryStateReasonCode{
		"INVALID_PERMISSIONS",
		"CREATE_PENDING",
		"CREATE_IN_PROGRESS",
		"CREATE_FAILED",
		"UPDATE_PENDING",
		"UPDATE_IN_PROGRESS",
		"UPDATE_COMPLETE",
		"ROLLBACK_COMPLETE",
		"ROLLBACK_FAILED",
		"DELETE_IN_PROGRESS",
		"DELETE_FAILED",
		"SYNC_DELETE_IN_PROGRESS",
	}
}

type EncryptionMode string

// Enum values for EncryptionMode
const (
	EncryptionModeSseS3  EncryptionMode = "SSE_S3"
	EncryptionModeSseKms EncryptionMode = "SSE_KMS"
)

// Values returns all known values for EncryptionMode. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (EncryptionMode) Values() []EncryptionMode {
	return []EncryptionMode{
		"SSE_S3",
		"SSE_KMS",
	}
}