// Code generated by smithy-go-codegen DO NOT EDIT.

package fsx

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/fsx/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new Amazon FSx for Lustre, Amazon FSx for Windows File Server, or
// Amazon FSx for OpenZFS file system from an existing Amazon FSx backup. If a file
// system with the specified client request token exists and the parameters match,
// this operation returns the description of the file system. If a file system with
// the specified client request token exists but the parameters don't match, this
// call returns IncompatibleParameterError. If a file system with the specified
// client request token doesn't exist, this operation does the following:
//
// *
// Creates a new Amazon FSx file system from backup with an assigned ID, and an
// initial lifecycle state of CREATING.
//
// * Returns the description of the file
// system.
//
// Parameters like the Active Directory, default share name, automatic
// backup, and backup settings default to the parameters of the file system that
// was backed up, unless overridden. You can explicitly supply other settings. By
// using the idempotent operation, you can retry a CreateFileSystemFromBackup call
// without the risk of creating an extra file system. This approach can be useful
// when an initial call fails in a way that makes it unclear whether a file system
// was created. Examples are if a transport level timeout occurred, or your
// connection was reset. If you use the same client request token and the initial
// call created a file system, the client receives a success message as long as the
// parameters are the same. The CreateFileSystemFromBackup call returns while the
// file system's lifecycle state is still CREATING. You can check the file-system
// creation status by calling the  DescribeFileSystems
// (https://docs.aws.amazon.com/fsx/latest/APIReference/API_DescribeFileSystems.html)
// operation, which returns the file system state along with other information.
func (c *Client) CreateFileSystemFromBackup(ctx context.Context, params *CreateFileSystemFromBackupInput, optFns ...func(*Options)) (*CreateFileSystemFromBackupOutput, error) {
	if params == nil {
		params = &CreateFileSystemFromBackupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateFileSystemFromBackup", params, optFns, c.addOperationCreateFileSystemFromBackupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateFileSystemFromBackupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request object for the CreateFileSystemFromBackup operation.
type CreateFileSystemFromBackupInput struct {

	// The ID of the source backup. Specifies the backup that you are copying.
	//
	// This member is required.
	BackupId *string

	// Specifies the IDs of the subnets that the file system will be accessible from.
	// For Windows MULTI_AZ_1 file system deployment types, provide exactly two subnet
	// IDs, one for the preferred file server and one for the standby file server. You
	// specify one of these subnets as the preferred subnet using the
	// WindowsConfiguration > PreferredSubnetID property. Windows SINGLE_AZ_1 and
	// SINGLE_AZ_2 file system deployment types, Lustre file systems, and OpenZFS file
	// systems provide exactly one subnet ID. The file server is launched in that
	// subnet's Availability Zone.
	//
	// This member is required.
	SubnetIds []string

	// A string of up to 64 ASCII characters that Amazon FSx uses to ensure idempotent
	// creation. This string is automatically filled on your behalf when you use the
	// Command Line Interface (CLI) or an Amazon Web Services SDK.
	ClientRequestToken *string

	// Sets the version for the Amazon FSx for Lustre file system that you're creating
	// from a backup. Valid values are 2.10 and 2.12. You don't need to specify
	// FileSystemTypeVersion because it will be applied using the backup's
	// FileSystemTypeVersion setting. If you choose to specify FileSystemTypeVersion
	// when creating from backup, the value must match the backup's
	// FileSystemTypeVersion setting.
	FileSystemTypeVersion *string

	// Specifies the ID of the Key Management Service (KMS) key to use for encrypting
	// data on Amazon FSx file systems, as follows:
	//
	// * Amazon FSx for Lustre
	// PERSISTENT_1 and PERSISTENT_2 deployment types only. SCRATCH_1 and SCRATCH_2
	// types are encrypted using the Amazon FSx service KMS key for your account.
	//
	// *
	// Amazon FSx for NetApp ONTAP
	//
	// * Amazon FSx for OpenZFS
	//
	// * Amazon FSx for Windows
	// File Server
	//
	// If a KmsKeyId isn't specified, the Amazon FSx-managed KMS key for
	// your account is used. For more information, see Encrypt
	// (https://docs.aws.amazon.com/kms/latest/APIReference/API_Encrypt.html) in the
	// Key Management Service API Reference.
	KmsKeyId *string

	// The Lustre configuration for the file system being created. The following
	// parameters are not supported for file systems with the Persistent_2 deployment
	// type. Instead, use CreateDataRepositoryAssociation to create a data repository
	// association to link your Lustre file system to a data repository.
	//
	// *
	// AutoImportPolicy
	//
	// * ExportPath
	//
	// * ImportedChunkSize
	//
	// * ImportPath
	LustreConfiguration *types.CreateFileSystemLustreConfiguration

	// The OpenZFS configuration for the file system that's being created.
	OpenZFSConfiguration *types.CreateFileSystemOpenZFSConfiguration

	// A list of IDs for the security groups that apply to the specified network
	// interfaces created for file system access. These security groups apply to all
	// network interfaces. This value isn't returned in later DescribeFileSystem
	// requests.
	SecurityGroupIds []string

	// Sets the storage type for the Windows or OpenZFS file system that you're
	// creating from a backup. Valid values are SSD and HDD.
	//
	// * Set to SSD to use solid
	// state drive storage. SSD is supported on all Windows and OpenZFS deployment
	// types.
	//
	// * Set to HDD to use hard disk drive storage. HDD is supported on
	// SINGLE_AZ_2 and MULTI_AZ_1 FSx for Windows File Server file system deployment
	// types.
	//
	// The default value is SSD. HDD and SSD storage types have different
	// minimum storage capacity requirements. A restored file system's storage capacity
	// is tied to the file system that was backed up. You can create a file system that
	// uses HDD storage from a backup of a file system that used SSD storage if the
	// original SSD file system had a storage capacity of at least 2000 GiB.
	StorageType types.StorageType

	// The tags to be applied to the file system at file system creation. The key value
	// of the Name tag appears in the console as the file system name.
	Tags []types.Tag

	// The configuration for this Microsoft Windows file system.
	WindowsConfiguration *types.CreateFileSystemWindowsConfiguration

	noSmithyDocumentSerde
}

// The response object for the CreateFileSystemFromBackup operation.
type CreateFileSystemFromBackupOutput struct {

	// A description of the file system.
	FileSystem *types.FileSystem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateFileSystemFromBackupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateFileSystemFromBackup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateFileSystemFromBackup{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateFileSystemFromBackupMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateFileSystemFromBackupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateFileSystemFromBackup(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateFileSystemFromBackup struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateFileSystemFromBackup) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateFileSystemFromBackup) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateFileSystemFromBackupInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateFileSystemFromBackupInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateFileSystemFromBackupMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateFileSystemFromBackup{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateFileSystemFromBackup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "fsx",
		OperationName: "CreateFileSystemFromBackup",
	}
}