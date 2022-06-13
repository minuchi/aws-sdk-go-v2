// Code generated by smithy-go-codegen DO NOT EDIT.

package mq

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns the specified configuration revision for the specified configuration.
func (c *Client) DescribeConfigurationRevision(ctx context.Context, params *DescribeConfigurationRevisionInput, optFns ...func(*Options)) (*DescribeConfigurationRevisionOutput, error) {
	if params == nil {
		params = &DescribeConfigurationRevisionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeConfigurationRevision", params, optFns, c.addOperationDescribeConfigurationRevisionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeConfigurationRevisionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeConfigurationRevisionInput struct {

	// The unique ID that Amazon MQ generates for the configuration.
	//
	// This member is required.
	ConfigurationId *string

	// The revision of the configuration.
	//
	// This member is required.
	ConfigurationRevision *string

	noSmithyDocumentSerde
}

type DescribeConfigurationRevisionOutput struct {

	// Required. The unique ID that Amazon MQ generates for the configuration.
	ConfigurationId *string

	// Required. The date and time of the configuration.
	Created *time.Time

	// Required. The base64-encoded XML configuration.
	Data *string

	// The description of the configuration.
	Description *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeConfigurationRevisionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeConfigurationRevision{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeConfigurationRevision{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDescribeConfigurationRevisionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeConfigurationRevision(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opDescribeConfigurationRevision(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mq",
		OperationName: "DescribeConfigurationRevision",
	}
}