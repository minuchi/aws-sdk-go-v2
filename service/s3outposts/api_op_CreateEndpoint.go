// Code generated by smithy-go-codegen DO NOT EDIT.

package s3outposts

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3outposts/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an endpoint and associates it with the specified Outpost. It can take up
// to 5 minutes for this action to finish. Related actions include:
//
// *
// DeleteEndpoint
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_s3outposts_DeleteEndpoint.html)
//
// *
// ListEndpoints
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_s3outposts_ListEndpoints.html)
func (c *Client) CreateEndpoint(ctx context.Context, params *CreateEndpointInput, optFns ...func(*Options)) (*CreateEndpointOutput, error) {
	if params == nil {
		params = &CreateEndpointInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateEndpoint", params, optFns, c.addOperationCreateEndpointMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateEndpointInput struct {

	// The ID of the Outposts.
	//
	// This member is required.
	OutpostId *string

	// The ID of the security group to use with the endpoint.
	//
	// This member is required.
	SecurityGroupId *string

	// The ID of the subnet in the selected VPC. The endpoint subnet must belong to the
	// Outpost that has Amazon S3 on Outposts provisioned.
	//
	// This member is required.
	SubnetId *string

	// The type of access for the network connectivity for the Amazon S3 on Outposts
	// endpoint. To use the Amazon Web Services VPC, choose Private. To use the
	// endpoint with an on-premises network, choose CustomerOwnedIp. If you choose
	// CustomerOwnedIp, you must also provide the customer-owned IP address pool (CoIP
	// pool). Private is the default access type value.
	AccessType types.EndpointAccessType

	// The ID of the customer-owned IPv4 address pool (CoIP pool) for the endpoint. IP
	// addresses are allocated from this pool for the endpoint.
	CustomerOwnedIpv4Pool *string

	noSmithyDocumentSerde
}

type CreateEndpointOutput struct {

	// The Amazon Resource Name (ARN) of the endpoint.
	EndpointArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateEndpointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateEndpoint{}, middleware.After)
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
	if err = addOpCreateEndpointValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateEndpoint(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateEndpoint(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3-outposts",
		OperationName: "CreateEndpoint",
	}
}