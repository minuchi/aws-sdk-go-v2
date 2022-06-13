// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
)

// Information about a single retained message.
type RetainedMessageSummary struct {

	// The Epoch date and time, in milliseconds, when the retained message was stored
	// by IoT.
	LastModifiedTime int64

	// The size of the retained message's payload in bytes.
	PayloadSize int64

	// The quality of service (QoS) level used to publish the retained message.
	Qos int32

	// The topic name to which the retained message was published.
	Topic *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde