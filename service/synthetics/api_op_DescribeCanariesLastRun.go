// Code generated by smithy-go-codegen DO NOT EDIT.

package synthetics

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/synthetics/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Use this operation to see information from the most recent run of each canary
// that you have created. This operation supports resource-level authorization
// using an IAM policy and the Names parameter. If you specify the Names parameter,
// the operation is successful only if you have authorization to view all the
// canaries that you specify in your request. If you do not have permission to view
// any of the canaries, the request fails with a 403 response. You are required to
// use the Names parameter if you are logged on to a user or role that has an IAM
// policy that restricts which canaries that you are allowed to view. For more
// information, see  Limiting a user to viewing specific canaries
// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_Restricted.html).
func (c *Client) DescribeCanariesLastRun(ctx context.Context, params *DescribeCanariesLastRunInput, optFns ...func(*Options)) (*DescribeCanariesLastRunOutput, error) {
	if params == nil {
		params = &DescribeCanariesLastRunInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeCanariesLastRun", params, optFns, c.addOperationDescribeCanariesLastRunMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeCanariesLastRunOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeCanariesLastRunInput struct {

	// Specify this parameter to limit how many runs are returned each time you use the
	// DescribeLastRun operation. If you omit this parameter, the default of 100 is
	// used.
	MaxResults *int32

	// Use this parameter to return only canaries that match the names that you specify
	// here. You can specify as many as five canary names. If you specify this
	// parameter, the operation is successful only if you have authorization to view
	// all the canaries that you specify in your request. If you do not have permission
	// to view any of the canaries, the request fails with a 403 response. You are
	// required to use the Names parameter if you are logged on to a user or role that
	// has an IAM policy that restricts which canaries that you are allowed to view.
	// For more information, see  Limiting a user to viewing specific canaries
	// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_Restricted.html).
	Names []string

	// A token that indicates that there is more data available. You can use this token
	// in a subsequent DescribeCanaries operation to retrieve the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeCanariesLastRunOutput struct {

	// An array that contains the information from the most recent run of each canary.
	CanariesLastRun []types.CanaryLastRun

	// A token that indicates that there is more data available. You can use this token
	// in a subsequent DescribeCanariesLastRun operation to retrieve the next set of
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeCanariesLastRunMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeCanariesLastRun{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeCanariesLastRun{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeCanariesLastRun(options.Region), middleware.Before); err != nil {
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

// DescribeCanariesLastRunAPIClient is a client that implements the
// DescribeCanariesLastRun operation.
type DescribeCanariesLastRunAPIClient interface {
	DescribeCanariesLastRun(context.Context, *DescribeCanariesLastRunInput, ...func(*Options)) (*DescribeCanariesLastRunOutput, error)
}

var _ DescribeCanariesLastRunAPIClient = (*Client)(nil)

// DescribeCanariesLastRunPaginatorOptions is the paginator options for
// DescribeCanariesLastRun
type DescribeCanariesLastRunPaginatorOptions struct {
	// Specify this parameter to limit how many runs are returned each time you use the
	// DescribeLastRun operation. If you omit this parameter, the default of 100 is
	// used.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeCanariesLastRunPaginator is a paginator for DescribeCanariesLastRun
type DescribeCanariesLastRunPaginator struct {
	options   DescribeCanariesLastRunPaginatorOptions
	client    DescribeCanariesLastRunAPIClient
	params    *DescribeCanariesLastRunInput
	nextToken *string
	firstPage bool
}

// NewDescribeCanariesLastRunPaginator returns a new
// DescribeCanariesLastRunPaginator
func NewDescribeCanariesLastRunPaginator(client DescribeCanariesLastRunAPIClient, params *DescribeCanariesLastRunInput, optFns ...func(*DescribeCanariesLastRunPaginatorOptions)) *DescribeCanariesLastRunPaginator {
	if params == nil {
		params = &DescribeCanariesLastRunInput{}
	}

	options := DescribeCanariesLastRunPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeCanariesLastRunPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeCanariesLastRunPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeCanariesLastRun page.
func (p *DescribeCanariesLastRunPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeCanariesLastRunOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.DescribeCanariesLastRun(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeCanariesLastRun(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "synthetics",
		OperationName: "DescribeCanariesLastRun",
	}
}