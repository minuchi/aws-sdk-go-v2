// Code generated by smithy-go-codegen DO NOT EDIT.

package migrationhuborchestrator

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Delete a step in a migration workflow. Pause the workflow to delete a running
// step.
func (c *Client) DeleteWorkflowStep(ctx context.Context, params *DeleteWorkflowStepInput, optFns ...func(*Options)) (*DeleteWorkflowStepOutput, error) {
	if params == nil {
		params = &DeleteWorkflowStepInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteWorkflowStep", params, optFns, c.addOperationDeleteWorkflowStepMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteWorkflowStepOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteWorkflowStepInput struct {

	// The ID of the step you want to delete.
	//
	// This member is required.
	Id *string

	// The ID of the step group that contains the step you want to delete.
	//
	// This member is required.
	StepGroupId *string

	// The ID of the migration workflow.
	//
	// This member is required.
	WorkflowId *string

	noSmithyDocumentSerde
}

type DeleteWorkflowStepOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteWorkflowStepMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteWorkflowStep{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteWorkflowStep{}, middleware.After)
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
	if err = addOpDeleteWorkflowStepValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteWorkflowStep(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteWorkflowStep(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "migrationhub-orchestrator",
		OperationName: "DeleteWorkflowStep",
	}
}