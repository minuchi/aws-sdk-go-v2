// Code generated by smithy-go-codegen DO NOT EDIT.

package budgets

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/budgets/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes a budget action history detail.
func (c *Client) DescribeBudgetActionHistories(ctx context.Context, params *DescribeBudgetActionHistoriesInput, optFns ...func(*Options)) (*DescribeBudgetActionHistoriesOutput, error) {
	if params == nil {
		params = &DescribeBudgetActionHistoriesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeBudgetActionHistories", params, optFns, c.addOperationDescribeBudgetActionHistoriesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeBudgetActionHistoriesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeBudgetActionHistoriesInput struct {

	// The account ID of the user. It's a 12-digit number.
	//
	// This member is required.
	AccountId *string

	// A system-generated universally unique identifier (UUID) for the action.
	//
	// This member is required.
	ActionId *string

	// A string that represents the budget name. The ":" and "\" characters aren't
	// allowed.
	//
	// This member is required.
	BudgetName *string

	// An integer that represents how many entries a paginated response contains. The
	// maximum is 100.
	MaxResults *int32

	// A generic string.
	NextToken *string

	// The period of time that's covered by a budget. The period has a start date and
	// an end date. The start date must come before the end date. There are no
	// restrictions on the end date.
	TimePeriod *types.TimePeriod

	noSmithyDocumentSerde
}

type DescribeBudgetActionHistoriesOutput struct {

	// The historical record of the budget action resource.
	//
	// This member is required.
	ActionHistories []types.ActionHistory

	// A generic string.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeBudgetActionHistoriesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeBudgetActionHistories{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeBudgetActionHistories{}, middleware.After)
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
	if err = addOpDescribeBudgetActionHistoriesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeBudgetActionHistories(options.Region), middleware.Before); err != nil {
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

// DescribeBudgetActionHistoriesAPIClient is a client that implements the
// DescribeBudgetActionHistories operation.
type DescribeBudgetActionHistoriesAPIClient interface {
	DescribeBudgetActionHistories(context.Context, *DescribeBudgetActionHistoriesInput, ...func(*Options)) (*DescribeBudgetActionHistoriesOutput, error)
}

var _ DescribeBudgetActionHistoriesAPIClient = (*Client)(nil)

// DescribeBudgetActionHistoriesPaginatorOptions is the paginator options for
// DescribeBudgetActionHistories
type DescribeBudgetActionHistoriesPaginatorOptions struct {
	// An integer that represents how many entries a paginated response contains. The
	// maximum is 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeBudgetActionHistoriesPaginator is a paginator for
// DescribeBudgetActionHistories
type DescribeBudgetActionHistoriesPaginator struct {
	options   DescribeBudgetActionHistoriesPaginatorOptions
	client    DescribeBudgetActionHistoriesAPIClient
	params    *DescribeBudgetActionHistoriesInput
	nextToken *string
	firstPage bool
}

// NewDescribeBudgetActionHistoriesPaginator returns a new
// DescribeBudgetActionHistoriesPaginator
func NewDescribeBudgetActionHistoriesPaginator(client DescribeBudgetActionHistoriesAPIClient, params *DescribeBudgetActionHistoriesInput, optFns ...func(*DescribeBudgetActionHistoriesPaginatorOptions)) *DescribeBudgetActionHistoriesPaginator {
	if params == nil {
		params = &DescribeBudgetActionHistoriesInput{}
	}

	options := DescribeBudgetActionHistoriesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeBudgetActionHistoriesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeBudgetActionHistoriesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeBudgetActionHistories page.
func (p *DescribeBudgetActionHistoriesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeBudgetActionHistoriesOutput, error) {
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

	result, err := p.client.DescribeBudgetActionHistories(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeBudgetActionHistories(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "budgets",
		OperationName: "DescribeBudgetActionHistories",
	}
}