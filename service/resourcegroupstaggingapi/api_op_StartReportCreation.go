// Code generated by smithy-go-codegen DO NOT EDIT.

package resourcegroupstaggingapi

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Generates a report that lists all tagged resources in the accounts across your
// organization and tells whether each resource is compliant with the effective tag
// policy. Compliance data is refreshed daily. The report is generated
// asynchronously. The generated report is saved to the following location:
// s3://example-bucket/AwsTagPolicies/o-exampleorgid/YYYY-MM-ddTHH:mm:ssZ/report.csv
// You can call this operation only from the organization's management account and
// from the us-east-1 Region.
func (c *Client) StartReportCreation(ctx context.Context, params *StartReportCreationInput, optFns ...func(*Options)) (*StartReportCreationOutput, error) {
	if params == nil {
		params = &StartReportCreationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartReportCreation", params, optFns, c.addOperationStartReportCreationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartReportCreationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartReportCreationInput struct {

	// The name of the Amazon S3 bucket where the report will be stored; for example:
	// awsexamplebucket For more information on S3 bucket requirements, including an
	// example bucket policy, see the example S3 bucket policy on this page.
	//
	// This member is required.
	S3Bucket *string

	noSmithyDocumentSerde
}

type StartReportCreationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartReportCreationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartReportCreation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartReportCreation{}, middleware.After)
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
	if err = addOpStartReportCreationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartReportCreation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartReportCreation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "tagging",
		OperationName: "StartReportCreation",
	}
}