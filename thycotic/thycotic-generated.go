package thycotic

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"time"
	"github.com/golang/glog"
)

//PE added
type String string

// against "unused imports"
var _ time.Time
var _ xml.Name

type SecretAccessRequestStatus string

const (
	SecretAccessRequestStatusPending SecretAccessRequestStatus = "Pending"

	SecretAccessRequestStatusApproved SecretAccessRequestStatus = "Approved"

	SecretAccessRequestStatusDenied SecretAccessRequestStatus = "Denied"

	SecretAccessRequestStatusCanceled SecretAccessRequestStatus = "Canceled"
)

type SshArgumentType string

const (
	SshArgumentTypeInterpreted SshArgumentType = "Interpreted"

	SshArgumentTypeLiteral SshArgumentType = "Literal"
)

type DbType string

const (
	DbTypeAnsiString DbType = "AnsiString"

	DbTypeBinary DbType = "Binary"

	DbTypeByte DbType = "Byte"

	DbTypeBoolean DbType = "Boolean"

	DbTypeCurrency DbType = "Currency"

	DbTypeDate DbType = "Date"

	DbTypeDateTime DbType = "DateTime"

	DbTypeDecimal DbType = "Decimal"

	DbTypeDouble DbType = "Double"

	DbTypeGuid DbType = "Guid"

	DbTypeInt16 DbType = "Int16"

	DbTypeInt32 DbType = "Int32"

	DbTypeInt64 DbType = "Int64"

	DbTypeObject DbType = "Object"

	DbTypeSByte DbType = "SByte"

	DbTypeSingle DbType = "Single"

	DbTypeString DbType = "String"

	DbTypeTime DbType = "Time"

	DbTypeUInt16 DbType = "UInt16"

	DbTypeUInt32 DbType = "UInt32"

	DbTypeUInt64 DbType = "UInt64"

	DbTypeVarNumeric DbType = "VarNumeric"

	DbTypeAnsiStringFixedLength DbType = "AnsiStringFixedLength"

	DbTypeStringFixedLength DbType = "StringFixedLength"

	DbTypeXml DbType = "Xml"

	DbTypeDateTime2 DbType = "DateTime2"

	DbTypeDateTimeOffset DbType = "DateTimeOffset"
)

type UserGroupMapType string

const (
	UserGroupMapTypeUser UserGroupMapType = "User"

	UserGroupMapTypeGroup UserGroupMapType = "Group"
)

type LineEnding string

const (
	LineEndingNewLine LineEnding = "NewLine"

	LineEndingCarriageReturn LineEnding = "CarriageReturn"

	LineEndingCarriageReturnNewLine LineEnding = "CarriageReturnNewLine"
)

type SshArgumentType2 string

const (
	SshArgumentType2Interpreted SshArgumentType2 = "Interpreted"

	SshArgumentType2Literal SshArgumentType2 = "Literal"
)

type ApproveSecretAccessRequest struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com ApproveSecretAccessRequest"`

	ApprovalId string `xml:"approvalId,omitempty"`

	Hours string `xml:"hours,omitempty"`

	UserOverride bool `xml:"userOverride,omitempty"`
}

type ApproveSecretAccessRequestResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com ApproveSecretAccessRequestResponse"`

	ApproveSecretAccessRequestResult *RequestApprovalResult `xml:"ApproveSecretAccessRequestResult,omitempty"`
}

type DenySecretAccessRequest struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com DenySecretAccessRequest"`

	ApprovalId string `xml:"approvalId,omitempty"`

	UserOverride bool `xml:"userOverride,omitempty"`
}

type DenySecretAccessRequestResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com DenySecretAccessRequestResponse"`

	DenySecretAccessRequestResult *RequestApprovalResult `xml:"DenySecretAccessRequestResult,omitempty"`
}

type Authenticate struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com Authenticate"`

	Username string `xml:"username,omitempty"`

	Password string `xml:"password,omitempty"`

	Organization string `xml:"organization,omitempty"`

	Domain string `xml:"domain,omitempty"`
}

type AuthenticateResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com AuthenticateResponse"`

	AuthenticateResult *AuthenticateResult `xml:"AuthenticateResult,omitempty"`
}

type ImpersonateUser struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com ImpersonateUser"`

	Token string `xml:"token,omitempty"`

	Username string `xml:"username,omitempty"`

	Organization string `xml:"organization,omitempty"`

	Domain string `xml:"domain,omitempty"`
}

type ImpersonateUserResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com ImpersonateUserResponse"`

	ImpersonateUserResult *ImpersonateResult `xml:"ImpersonateUserResult,omitempty"`
}

type AuthenticateRADIUS struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com AuthenticateRADIUS"`

	Username string `xml:"username,omitempty"`

	Password string `xml:"password,omitempty"`

	Organization string `xml:"organization,omitempty"`

	Domain string `xml:"domain,omitempty"`

	RadiusPassword string `xml:"radiusPassword,omitempty"`
}

type AuthenticateRADIUSResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com AuthenticateRADIUSResponse"`

	AuthenticateRADIUSResult *AuthenticateResult `xml:"AuthenticateRADIUSResult,omitempty"`
}

type GetTokenIsValid struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GetTokenIsValid"`

	Token string `xml:"token,omitempty"`
}

type GetTokenIsValidResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GetTokenIsValidResponse"`

	GetTokenIsValidResult *TokenIsValidResult `xml:"GetTokenIsValidResult,omitempty"`
}

type GetSecretLegacy struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GetSecretLegacy"`

	Token string `xml:"token,omitempty"`

	SecretId int32 `xml:"secretId,omitempty"`
}

type GetSecretLegacyResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GetSecretLegacyResponse"`

	GetSecretLegacyResult *GetSecretResult `xml:"GetSecretLegacyResult,omitempty"`
}

type GetSecret struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GetSecret"`

	Token string `xml:"token,omitempty"`

	SecretId int32 `xml:"secretId,omitempty"`

	LoadSettingsAndPermissions bool `xml:"loadSettingsAndPermissions,omitempty"`

	CodeResponses *ArrayOfCodeResponse `xml:"codeResponses,omitempty"`
}

type GetSecretResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GetSecretResponse"`

	GetSecretResult *GetSecretResult `xml:"GetSecretResult,omitempty"`
}

type GetCheckOutStatus struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GetCheckOutStatus"`

	Token string `xml:"token,omitempty"`

	SecretId int32 `xml:"secretId,omitempty"`
}

type GetCheckOutStatusResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GetCheckOutStatusResponse"`

	GetCheckOutStatusResult *GetCheckOutStatusResult `xml:"GetCheckOutStatusResult,omitempty"`
}

type ChangePassword struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com ChangePassword"`

	Token string `xml:"token,omitempty"`

	CurrentPassword string `xml:"currentPassword,omitempty"`

	NewPassword string `xml:"newPassword,omitempty"`
}

type ChangePasswordResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com ChangePasswordResponse"`

	ChangePasswordResult *WebServiceResult `xml:"ChangePasswordResult,omitempty"`
}

type GetSecretsByFieldValue struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GetSecretsByFieldValue"`

	Token string `xml:"token,omitempty"`

	FieldName string `xml:"fieldName,omitempty"`

	SearchTerm string `xml:"searchTerm,omitempty"`

	ShowDeleted bool `xml:"showDeleted,omitempty"`
}

type GetSecretsByFieldValueResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GetSecretsByFieldValueResponse"`

	GetSecretsByFieldValueResult *GetSecretsByFieldValueResult `xml:"GetSecretsByFieldValueResult,omitempty"`
}

type SearchSecretsByFieldValue struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com SearchSecretsByFieldValue"`

	Token string `xml:"token,omitempty"`

	FieldName string `xml:"fieldName,omitempty"`

	SearchTerm string `xml:"searchTerm,omitempty"`

	ShowDeleted bool `xml:"showDeleted,omitempty"`

	ShowRestricted bool `xml:"showRestricted,omitempty"`
}

type SearchSecretsByFieldValueResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com SearchSecretsByFieldValueResponse"`

	SearchSecretsByFieldValueResult *SearchSecretsResult `xml:"SearchSecretsByFieldValueResult,omitempty"`
}

type GetSecretsByExposedFieldValue struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GetSecretsByExposedFieldValue"`

	Token string `xml:"token,omitempty"`

	FieldName string `xml:"fieldName,omitempty"`

	SearchTerm string `xml:"searchTerm,omitempty"`

	ShowDeleted bool `xml:"showDeleted,omitempty"`

	ShowPartialMatches bool `xml:"showPartialMatches,omitempty"`
}

type GetSecretsByExposedFieldValueResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GetSecretsByExposedFieldValueResponse"`

	GetSecretsByExposedFieldValueResult *GetSecretsByFieldValueResult `xml:"GetSecretsByExposedFieldValueResult,omitempty"`
}

type SearchSecretsByExposedFieldValue struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com SearchSecretsByExposedFieldValue"`

	Token string `xml:"token,omitempty"`

	FieldName string `xml:"fieldName,omitempty"`

	SearchTerm string `xml:"searchTerm,omitempty"`

	ShowDeleted bool `xml:"showDeleted,omitempty"`

	ShowRestricted bool `xml:"showRestricted,omitempty"`

	ShowPartialMatches bool `xml:"showPartialMatches,omitempty"`
}

type SearchSecretsByExposedFieldValueResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com SearchSecretsByExposedFieldValueResponse"`

	SearchSecretsByExposedFieldValueResult *SearchSecretsResult `xml:"SearchSecretsByExposedFieldValueResult,omitempty"`
}

type SearchSecretsByExposedValues struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com SearchSecretsByExposedValues"`

	Token string `xml:"token,omitempty"`

	SearchTerm string `xml:"searchTerm,omitempty"`

	ShowDeleted bool `xml:"showDeleted,omitempty"`

	ShowRestricted bool `xml:"showRestricted,omitempty"`

	ShowPartialMatches bool `xml:"showPartialMatches,omitempty"`
}

type SearchSecretsByExposedValuesResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com SearchSecretsByExposedValuesResponse"`

	SearchSecretsByExposedValuesResult *SearchSecretsResult `xml:"SearchSecretsByExposedValuesResult,omitempty"`
}

type AddUser struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com AddUser"`

	Token string `xml:"token,omitempty"`

	NewUser *User `xml:"newUser,omitempty"`
}

type AddUserResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com AddUserResponse"`

	AddUserResult *WebServiceResult `xml:"AddUserResult,omitempty"`
}

type SearchSecrets struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com SearchSecrets"`

	Token string `xml:"token,omitempty"`

	SearchTerm string `xml:"searchTerm,omitempty"`

	IncludeDeleted bool `xml:"includeDeleted,omitempty"`

	IncludeRestricted bool `xml:"includeRestricted,omitempty"`
}

type SearchSecretsResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com SearchSecretsResponse"`

	SearchSecretsResult *SearchSecretsResult `xml:"SearchSecretsResult,omitempty"`
}

type SearchSecretsLegacy struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com SearchSecretsLegacy"`

	Token string `xml:"token,omitempty"`

	SearchTerm string `xml:"searchTerm,omitempty"`
}

type SearchSecretsLegacyResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com SearchSecretsLegacyResponse"`

	SearchSecretsLegacyResult *SearchSecretsResult `xml:"SearchSecretsLegacyResult,omitempty"`
}

type SearchSecretsByFolder struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com SearchSecretsByFolder"`

	Token string `xml:"token,omitempty"`

	SearchTerm string `xml:"searchTerm,omitempty"`

	FolderId int32 `xml:"folderId,omitempty"`

	IncludeSubFolders bool `xml:"includeSubFolders,omitempty"`

	IncludeDeleted bool `xml:"includeDeleted,omitempty"`

	IncludeRestricted bool `xml:"includeRestricted,omitempty"`
}

type SearchSecretsByFolderResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com SearchSecretsByFolderResponse"`

	SearchSecretsByFolderResult *SearchSecretsResult `xml:"SearchSecretsByFolderResult,omitempty"`
}

type SearchSecretsByFolderLegacy struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com SearchSecretsByFolderLegacy"`

	Token string `xml:"token,omitempty"`

	SearchTerm string `xml:"searchTerm,omitempty"`

	FolderId int32 `xml:"folderId,omitempty"`

	IncludeSubFolders bool `xml:"includeSubFolders,omitempty"`
}

type SearchSecretsByFolderLegacyResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com SearchSecretsByFolderLegacyResponse"`

	SearchSecretsByFolderLegacyResult *SearchSecretsResult `xml:"SearchSecretsByFolderLegacyResult,omitempty"`
}

type GetFavorites struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GetFavorites"`

	Token string `xml:"token,omitempty"`

	IncludeRestricted bool `xml:"includeRestricted,omitempty"`
}

type GetFavoritesResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GetFavoritesResponse"`

	GetFavoritesResult *GetFavoritesResult `xml:"GetFavoritesResult,omitempty"`
}

type UpdateIsFavorite struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com UpdateIsFavorite"`

	Token string `xml:"token,omitempty"`

	SecretId int32 `xml:"secretId,omitempty"`

	IsFavorite bool `xml:"isFavorite,omitempty"`
}

type UpdateIsFavoriteResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com UpdateIsFavoriteResponse"`

	UpdateIsFavoriteResult *WebServiceResult `xml:"UpdateIsFavoriteResult,omitempty"`
}

type AddSecret struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com AddSecret"`

	Token string `xml:"token,omitempty"`

	SecretTypeId int32 `xml:"secretTypeId,omitempty"`

	SecretName string `xml:"secretName,omitempty"`

	SecretFieldIds *ArrayOfInt `xml:"secretFieldIds,omitempty"`

	SecretItemValues *ArrayOfString `xml:"secretItemValues,omitempty"`

	FolderId int32 `xml:"folderId,omitempty"`
}

type AddSecretResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com AddSecretResponse"`

	AddSecretResult *AddSecretResult `xml:"AddSecretResult,omitempty"`
}

type AddNewSecret struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com AddNewSecret"`

	Token string `xml:"token,omitempty"`

	Secret *Secret `xml:"secret,omitempty"`
}

type AddNewSecretResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com AddNewSecretResponse"`

	AddNewSecretResult *AddSecretResult `xml:"AddNewSecretResult,omitempty"`
}

type GetNewSecret struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GetNewSecret"`

	Token string `xml:"token,omitempty"`

	SecretTypeId int32 `xml:"secretTypeId,omitempty"`

	FolderId int32 `xml:"folderId,omitempty"`
}

type GetNewSecretResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GetNewSecretResponse"`

	GetNewSecretResult *GetSecretResult `xml:"GetNewSecretResult,omitempty"`
}

type GetSecretTemplateFields struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GetSecretTemplateFields"`

	Token string `xml:"token,omitempty"`

	SecretTypeId int32 `xml:"secretTypeId,omitempty"`
}

type GetSecretTemplateFieldsResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GetSecretTemplateFieldsResponse"`

	GetSecretTemplateFieldsResult *GetSecretTemplateFieldsResult `xml:"GetSecretTemplateFieldsResult,omitempty"`
}

type UpdateSecret struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com UpdateSecret"`

	Token string `xml:"token,omitempty"`

	Secret *Secret `xml:"secret,omitempty"`
}

type UpdateSecretResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com UpdateSecretResponse"`

	UpdateSecretResult *WebServiceResult `xml:"UpdateSecretResult,omitempty"`
}

type GetSecretTemplates struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GetSecretTemplates"`

	Token string `xml:"token,omitempty"`
}

type GetSecretTemplatesResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GetSecretTemplatesResponse"`

	GetSecretTemplatesResult *GetSecretTemplatesResult `xml:"GetSecretTemplatesResult,omitempty"`
}

type GeneratePassword struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GeneratePassword"`

	Token string `xml:"token,omitempty"`

	SecretFieldId int32 `xml:"secretFieldId,omitempty"`
}

type GeneratePasswordResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GeneratePasswordResponse"`

	GeneratePasswordResult *GeneratePasswordResult `xml:"GeneratePasswordResult,omitempty"`
}

type DeactivateSecret struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com DeactivateSecret"`

	Token string `xml:"token,omitempty"`

	SecretId int32 `xml:"secretId,omitempty"`
}

type DeactivateSecretResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com DeactivateSecretResponse"`

	DeactivateSecretResult *WebServiceResult `xml:"DeactivateSecretResult,omitempty"`
}

type VersionGet struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com VersionGet"`
}

type VersionGetResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com VersionGetResponse"`

	VersionGetResult *VersionGetResult `xml:"VersionGetResult,omitempty"`
}

type FolderGet struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com FolderGet"`

	Token string `xml:"token,omitempty"`

	FolderId int32 `xml:"folderId,omitempty"`
}

type FolderGetResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com FolderGetResponse"`

	FolderGetResult *GetFolderResult `xml:"FolderGetResult,omitempty"`
}

type FolderUpdate struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com FolderUpdate"`

	Token string `xml:"token,omitempty"`

	ModifiedFolder *Folder `xml:"modifiedFolder,omitempty"`
}

type FolderUpdateResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com FolderUpdateResponse"`

	FolderUpdateResult *WebServiceResult `xml:"FolderUpdateResult,omitempty"`
}

type FolderGetAllChildren struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com FolderGetAllChildren"`

	Token string `xml:"token,omitempty"`

	ParentFolderId int32 `xml:"parentFolderId,omitempty"`
}

type FolderGetAllChildrenResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com FolderGetAllChildrenResponse"`

	FolderGetAllChildrenResult *GetFoldersResult `xml:"FolderGetAllChildrenResult,omitempty"`
}

type FolderCreate struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com FolderCreate"`

	Token string `xml:"token,omitempty"`

	FolderName string `xml:"folderName,omitempty"`

	ParentFolderId int32 `xml:"parentFolderId,omitempty"`

	FolderTypeId int32 `xml:"folderTypeId,omitempty"`
}

type FolderCreateResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com FolderCreateResponse"`

	FolderCreateResult *CreateFolderResult `xml:"FolderCreateResult,omitempty"`
}

type FolderExtendedCreate struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com FolderExtendedCreate"`

	Token string `xml:"token,omitempty"`

	Folder *FolderExtended `xml:"folder,omitempty"`
}

type FolderExtendedCreateResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com FolderExtendedCreateResponse"`

	FolderExtendedCreateResult *FolderExtendedCreateResult `xml:"FolderExtendedCreateResult,omitempty"`
}

type FolderExtendedGet struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com FolderExtendedGet"`

	Token string `xml:"token,omitempty"`

	FolderId int32 `xml:"folderId,omitempty"`
}

type FolderExtendedGetResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com FolderExtendedGetResponse"`

	FolderExtendedGetResult *FolderExtendedGetResult `xml:"FolderExtendedGetResult,omitempty"`
}

type FolderExtendedUpdate struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com FolderExtendedUpdate"`

	Token string `xml:"token,omitempty"`

	Folder *FolderExtended `xml:"folder,omitempty"`
}

type FolderExtendedUpdateResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com FolderExtendedUpdateResponse"`

	FolderExtendedUpdateResult *FolderExtendedUpdateResult `xml:"FolderExtendedUpdateResult,omitempty"`
}

type FolderExtendedGetNew struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com FolderExtendedGetNew"`

	Token string `xml:"token,omitempty"`

	FolderExtendedGetNewRequest *FolderExtendedGetNewRequest `xml:"folderExtendedGetNewRequest,omitempty"`
}

type FolderExtendedGetNewResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com FolderExtendedGetNewResponse"`

	FolderExtendedGetNewResult *FolderExtendedGetNewResult `xml:"FolderExtendedGetNewResult,omitempty"`
}

type SearchFolders struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com SearchFolders"`

	Token string `xml:"token,omitempty"`

	FolderName string `xml:"folderName,omitempty"`
}

type SearchFoldersResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com SearchFoldersResponse"`

	SearchFoldersResult *SearchFolderResult `xml:"SearchFoldersResult,omitempty"`
}

type DownloadFileAttachment struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com DownloadFileAttachment"`

	Token string `xml:"token,omitempty"`

	SecretId int32 `xml:"secretId,omitempty"`
}

type DownloadFileAttachmentResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com DownloadFileAttachmentResponse"`

	DownloadFileAttachmentResult *FileDownloadResult `xml:"DownloadFileAttachmentResult,omitempty"`
}

type DownloadFileAttachmentByItemId struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com DownloadFileAttachmentByItemId"`

	Token string `xml:"token,omitempty"`

	SecretId int32 `xml:"secretId,omitempty"`

	SecretItemId int32 `xml:"secretItemId,omitempty"`
}

type DownloadFileAttachmentByItemIdResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com DownloadFileAttachmentByItemIdResponse"`

	DownloadFileAttachmentByItemIdResult *FileDownloadResult `xml:"DownloadFileAttachmentByItemIdResult,omitempty"`
}

type UploadFileAttachment struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com UploadFileAttachment"`

	Token string `xml:"token,omitempty"`

	SecretId int32 `xml:"secretId,omitempty"`

	FileData []byte `xml:"fileData,omitempty"`

	FileName string `xml:"fileName,omitempty"`
}

type UploadFileAttachmentResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com UploadFileAttachmentResponse"`

	UploadFileAttachmentResult *WebServiceResult `xml:"UploadFileAttachmentResult,omitempty"`
}

type UploadFileAttachmentByItemId struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com UploadFileAttachmentByItemId"`

	Token string `xml:"token,omitempty"`

	SecretId int32 `xml:"secretId,omitempty"`

	SecretItemId int32 `xml:"secretItemId,omitempty"`

	FileData []byte `xml:"fileData,omitempty"`

	FileName string `xml:"fileName,omitempty"`
}

type UploadFileAttachmentByItemIdResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com UploadFileAttachmentByItemIdResponse"`

	UploadFileAttachmentByItemIdResult *WebServiceResult `xml:"UploadFileAttachmentByItemIdResult,omitempty"`
}

type ExpireSecret struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com ExpireSecret"`

	Token string `xml:"token,omitempty"`

	SecretId int32 `xml:"secretId,omitempty"`
}

type ExpireSecretResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com ExpireSecretResponse"`

	ExpireSecretResult *WebServiceResult `xml:"ExpireSecretResult,omitempty"`
}

type SetCheckOutEnabled struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com SetCheckOutEnabled"`

	Token string `xml:"token,omitempty"`

	SecretId int32 `xml:"secretId,omitempty"`

	SetCheckOut bool `xml:"setCheckOut,omitempty"`

	SetPasswordChangeOnCheckIn bool `xml:"setPasswordChangeOnCheckIn,omitempty"`

	CheckOutInterval int32 `xml:"checkOutInterval,omitempty"`
}

type SetCheckOutEnabledResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com SetCheckOutEnabledResponse"`

	SetCheckOutEnabledResult *WebServiceResult `xml:"SetCheckOutEnabledResult,omitempty"`
}

type ImportXML struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com ImportXML"`

	Token string `xml:"token,omitempty"`

	Xml string `xml:"xml,omitempty"`
}

type ImportXMLResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com ImportXMLResponse"`

	ImportXMLResult *WebServiceResult `xml:"ImportXMLResult,omitempty"`
}

type GetSecretAudit struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GetSecretAudit"`

	Token string `xml:"token,omitempty"`

	SecretId int32 `xml:"secretId,omitempty"`
}

type GetSecretAuditResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GetSecretAuditResponse"`

	GetSecretAuditResult *GetSecretAuditResult `xml:"GetSecretAuditResult,omitempty"`
}

type AddDependency struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com AddDependency"`

	Token string `xml:"token,omitempty"`

	Dependency *Dependency `xml:"dependency,omitempty"`
}

type AddDependencyResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com AddDependencyResponse"`

	AddDependencyResult *WebServiceResult `xml:"AddDependencyResult,omitempty"`
}

type RemoveDependency struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com RemoveDependency"`

	Token string `xml:"token,omitempty"`

	DependencyId int32 `xml:"dependencyId,omitempty"`

	SecretId int32 `xml:"secretId,omitempty"`
}

type RemoveDependencyResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com RemoveDependencyResponse"`

	RemoveDependencyResult *WebServiceResult `xml:"RemoveDependencyResult,omitempty"`
}

type GetDependencies struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GetDependencies"`

	Token string `xml:"token,omitempty"`

	SecretId int32 `xml:"secretId,omitempty"`
}

type GetDependenciesResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GetDependenciesResponse"`

	GetDependenciesResult *GetDependenciesResult `xml:"GetDependenciesResult,omitempty"`
}

type GetDistributedEngines struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GetDistributedEngines"`

	Token string `xml:"token,omitempty"`
}

type GetDistributedEnginesResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GetDistributedEnginesResponse"`

	GetDistributedEnginesResult *GetSitesResult `xml:"GetDistributedEnginesResult,omitempty"`
}

type GetTicketSystems struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GetTicketSystems"`

	Token string `xml:"token,omitempty"`
}

type GetTicketSystemsResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GetTicketSystemsResponse"`

	GetTicketSystemsResult *GetTicketSystemsResult `xml:"GetTicketSystemsResult,omitempty"`
}

type AssignSite struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com AssignSite"`

	Token string `xml:"token,omitempty"`

	SecretId int32 `xml:"secretId,omitempty"`

	SiteId int32 `xml:"siteId,omitempty"`
}

type AssignSiteResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com AssignSiteResponse"`

	AssignSiteResult *WebServiceResult `xml:"AssignSiteResult,omitempty"`
}

type CheckIn struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com CheckIn"`

	Token string `xml:"token,omitempty"`

	SecretId int32 `xml:"secretId,omitempty"`
}

type CheckInResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com CheckInResponse"`

	CheckInResult *WebServiceResult `xml:"CheckInResult,omitempty"`
}

type AddSecretCustomAudit struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com AddSecretCustomAudit"`

	Token string `xml:"token,omitempty"`

	SecretId int32 `xml:"secretId,omitempty"`

	Notes string `xml:"notes,omitempty"`

	IpAddress string `xml:"ipAddress,omitempty"`

	ReferenceId int32 `xml:"referenceId,omitempty"`

	TicketNumber string `xml:"ticketNumber,omitempty"`

	UserId int32 `xml:"userId,omitempty"`
}

type AddSecretCustomAuditResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com AddSecretCustomAuditResponse"`

	AddSecretCustomAuditResult *WebServiceResult `xml:"AddSecretCustomAuditResult,omitempty"`
}

type UpdateSecretPermission struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com UpdateSecretPermission"`

	Token string `xml:"token,omitempty"`

	SecretId int32 `xml:"secretId,omitempty"`

	GroupOrUserRecord *GroupOrUserRecord `xml:"groupOrUserRecord,omitempty"`

	View bool `xml:"view,omitempty"`

	Edit bool `xml:"edit,omitempty"`

	Owner bool `xml:"owner,omitempty"`
}

type UpdateSecretPermissionResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com UpdateSecretPermissionResponse"`

	UpdateSecretPermissionResult *WebServiceResult `xml:"UpdateSecretPermissionResult,omitempty"`
}

type CheckInByKey struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com CheckInByKey"`

	SessionKey string `xml:"sessionKey,omitempty"`
}

type CheckInByKeyResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com CheckInByKeyResponse"`

	CheckInByKeyResult *WebServiceResult `xml:"CheckInByKeyResult,omitempty"`
}

type WhoAmI struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com WhoAmI"`

	Token string `xml:"token,omitempty"`
}

type WhoAmIResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com WhoAmIResponse"`

	WhoAmIResult *UserInfoResult `xml:"WhoAmIResult,omitempty"`
}

type GetAllGroups struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GetAllGroups"`

	Token string `xml:"token,omitempty"`
}

type GetAllGroupsResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GetAllGroupsResponse"`

	GetAllGroupsResult *GetAllGroupsResult `xml:"GetAllGroupsResult,omitempty"`
}

type AssignUserToGroup struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com AssignUserToGroup"`

	Token string `xml:"token,omitempty"`

	UserId int32 `xml:"userId,omitempty"`

	GroupId int32 `xml:"groupId,omitempty"`
}

type AssignUserToGroupResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com AssignUserToGroupResponse"`

	AssignUserToGroupResult *WebServiceResult `xml:"AssignUserToGroupResult,omitempty"`
}

type GetSSHLoginCredentials struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GetSSHLoginCredentials"`

	Token string `xml:"token,omitempty"`

	SecretId int32 `xml:"secretId,omitempty"`
}

type GetSSHLoginCredentialsResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GetSSHLoginCredentialsResponse"`

	GetSSHLoginCredentialsResult *SSHCredentialsResult `xml:"GetSSHLoginCredentialsResult,omitempty"`
}

type GetSSHLoginCredentialsWithMachine struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GetSSHLoginCredentialsWithMachine"`

	Token string `xml:"token,omitempty"`

	SecretId int32 `xml:"secretId,omitempty"`

	Machine string `xml:"machine,omitempty"`
}

type GetSSHLoginCredentialsWithMachineResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GetSSHLoginCredentialsWithMachineResponse"`

	GetSSHLoginCredentialsWithMachineResult *SSHCredentialsResult `xml:"GetSSHLoginCredentialsWithMachineResult,omitempty"`
}

type SearchUsers struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com SearchUsers"`

	Token string `xml:"token,omitempty"`

	SearchTerm string `xml:"searchTerm,omitempty"`

	IncludeInactiveUsers bool `xml:"includeInactiveUsers,omitempty"`
}

type SearchUsersResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com SearchUsersResponse"`

	SearchUsersResult *GetUsersResult `xml:"SearchUsersResult,omitempty"`
}

type GetUser struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GetUser"`

	Token string `xml:"token,omitempty"`

	UserId int32 `xml:"userId,omitempty"`
}

type GetUserResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GetUserResponse"`

	GetUserResult *GetUserResult `xml:"GetUserResult,omitempty"`
}

type UpdateUser struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com UpdateUser"`

	Token string `xml:"token,omitempty"`

	User *User `xml:"user,omitempty"`
}

type UpdateUserResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com UpdateUserResponse"`

	UpdateUserResult *UpdateUserResult `xml:"UpdateUserResult,omitempty"`
}

type GetSecretItemHistoryByFieldName struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GetSecretItemHistoryByFieldName"`

	Token string `xml:"token,omitempty"`

	SecretId int32 `xml:"secretId,omitempty"`

	FieldDisplayName string `xml:"fieldDisplayName,omitempty"`
}

type GetSecretItemHistoryByFieldNameResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GetSecretItemHistoryByFieldNameResponse"`

	GetSecretItemHistoryByFieldNameResult *SecretItemHistoryResult `xml:"GetSecretItemHistoryByFieldNameResult,omitempty"`
}

type GetSecretPolicyForSecret struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GetSecretPolicyForSecret"`

	Token string `xml:"token,omitempty"`

	SecretId int32 `xml:"secretId,omitempty"`
}

type GetSecretPolicyForSecretResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GetSecretPolicyForSecretResponse"`

	GetSecretPolicyForSecretResult *SecretPolicyForSecretResult `xml:"GetSecretPolicyForSecretResult,omitempty"`
}

type AssignSecretPolicyForSecret struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com AssignSecretPolicyForSecret"`

	Token string `xml:"token,omitempty"`

	SecretPolicyForSecret *SecretPolicyForSecret `xml:"secretPolicyForSecret,omitempty"`
}

type AssignSecretPolicyForSecretResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com AssignSecretPolicyForSecretResponse"`

	AssignSecretPolicyForSecretResult *SecretPolicyForSecretResult `xml:"AssignSecretPolicyForSecretResult,omitempty"`
}

type SearchSecretPolicies struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com SearchSecretPolicies"`

	Token string `xml:"token,omitempty"`

	Term string `xml:"term,omitempty"`

	IncludeInactive bool `xml:"includeInactive,omitempty"`
}

type SearchSecretPoliciesResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com SearchSecretPoliciesResponse"`

	SearchSecretPoliciesResult *SearchSecretPoliciesResult `xml:"SearchSecretPoliciesResult,omitempty"`
}

type RunActiveDirectorySynchronization struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com RunActiveDirectorySynchronization"`

	Token string `xml:"token,omitempty"`
}

type RunActiveDirectorySynchronizationResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com RunActiveDirectorySynchronizationResponse"`

	RunActiveDirectorySynchronizationResult *WebServiceResult `xml:"RunActiveDirectorySynchronizationResult,omitempty"`
}

type AddGroupToActiveDirectorySynchronization struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com AddGroupToActiveDirectorySynchronization"`

	Token string `xml:"token,omitempty"`

	AddGroupRequestMessage *AddGroupRequestMessage `xml:"addGroupRequestMessage,omitempty"`
}

type AddGroupToActiveDirectorySynchronizationResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com AddGroupToActiveDirectorySynchronizationResponse"`

	AddGroupToActiveDirectorySynchronizationResult *WebServiceResult `xml:"AddGroupToActiveDirectorySynchronizationResult,omitempty"`
}

type AddSecretPolicy struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com AddSecretPolicy"`

	Token string `xml:"token,omitempty"`

	SecretPolicy *SecretPolicyDetail `xml:"secretPolicy,omitempty"`
}

type AddSecretPolicyResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com AddSecretPolicyResponse"`

	AddSecretPolicyResult *SecretPolicyResult `xml:"AddSecretPolicyResult,omitempty"`
}

type GetNewSecretPolicy struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GetNewSecretPolicy"`

	Token string `xml:"token,omitempty"`
}

type GetNewSecretPolicyResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GetNewSecretPolicyResponse"`

	GetNewSecretPolicyResult *SecretPolicyResult `xml:"GetNewSecretPolicyResult,omitempty"`
}

type GetSSHCommandMenu struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GetSSHCommandMenu"`

	Token string `xml:"token,omitempty"`

	SshCommandMenuId int32 `xml:"sshCommandMenuId,omitempty"`
}

type GetSSHCommandMenuResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GetSSHCommandMenuResponse"`

	GetSSHCommandMenuResult *GetSshCommandMenuResult `xml:"GetSSHCommandMenuResult,omitempty"`
}

type SaveSSHCommandMenu struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com SaveSSHCommandMenu"`

	Token string `xml:"token,omitempty"`

	SshCommandMenu *SshCommandMenu `xml:"sshCommandMenu,omitempty"`

	CommandsText string `xml:"commandsText,omitempty"`

	DeleteCommands bool `xml:"deleteCommands,omitempty"`
}

type SaveSSHCommandMenuResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com SaveSSHCommandMenuResponse"`

	SaveSSHCommandMenuResult *GetSshCommandMenuResult `xml:"SaveSSHCommandMenuResult,omitempty"`
}

type GetAllSSHCommandMenus struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GetAllSSHCommandMenus"`

	Token string `xml:"token,omitempty"`

	IncludeInactive bool `xml:"includeInactive,omitempty"`
}

type GetAllSSHCommandMenusResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GetAllSSHCommandMenusResponse"`

	GetAllSSHCommandMenusResult *GetSshCommandMenusResult `xml:"GetAllSSHCommandMenusResult,omitempty"`
}

type DeleteSSHCommandMenu struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com DeleteSSHCommandMenu"`

	Token string `xml:"token,omitempty"`

	SshCommandMenuId int32 `xml:"sshCommandMenuId,omitempty"`
}

type DeleteSSHCommandMenuResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com DeleteSSHCommandMenuResponse"`

	DeleteSSHCommandMenuResult *WebServiceResult `xml:"DeleteSSHCommandMenuResult,omitempty"`
}

type RestoreSSHCommandMenu struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com RestoreSSHCommandMenu"`

	Token string `xml:"token,omitempty"`

	SshCommandMenuId int32 `xml:"sshCommandMenuId,omitempty"`
}

type RestoreSSHCommandMenuResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com RestoreSSHCommandMenuResponse"`

	RestoreSSHCommandMenuResult *WebServiceResult `xml:"RestoreSSHCommandMenuResult,omitempty"`
}

type AddScript struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com AddScript"`

	Token string `xml:"token,omitempty"`

	NewUserScript *UserScript `xml:"newUserScript,omitempty"`
}

type AddScriptResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com AddScriptResponse"`

	AddScriptResult *WebServiceResult `xml:"AddScriptResult,omitempty"`
}

type GetAllScripts struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GetAllScripts"`

	Token string `xml:"token,omitempty"`

	IncludeInactiveUserScripts bool `xml:"includeInactiveUserScripts,omitempty"`
}

type GetAllScriptsResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GetAllScriptsResponse"`

	GetAllScriptsResult *GetUserScriptsResult `xml:"GetAllScriptsResult,omitempty"`
}

type GetScript struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GetScript"`

	Token string `xml:"token,omitempty"`

	UserScriptId int32 `xml:"userScriptId,omitempty"`
}

type GetScriptResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GetScriptResponse"`

	GetScriptResult *GetUserScriptResult `xml:"GetScriptResult,omitempty"`
}

type UpdateScript struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com UpdateScript"`

	Token string `xml:"token,omitempty"`

	UserScript *UserScript `xml:"userScript,omitempty"`
}

type UpdateScriptResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com UpdateScriptResponse"`

	UpdateScriptResult *UpdateUserScriptResult `xml:"UpdateScriptResult,omitempty"`
}

type RequestApprovalResult struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com RequestApprovalResult"`

	*GenericResult

	ApprovalInfo *ApprovalInfo `xml:"ApprovalInfo,omitempty"`
}

type GenericResult struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GenericResult"`

	ErrorMessage string `xml:"ErrorMessage,omitempty"`
}

type ApprovalInfo struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com ApprovalInfo"`

	Status *SecretAccessRequestStatus `xml:"Status,omitempty"`

	Responder string `xml:"Responder,omitempty"`

	ResponseDate time.Time `xml:"ResponseDate,omitempty"`

	ResponseComment string `xml:"ResponseComment,omitempty"`

	ExpirationDate time.Time `xml:"ExpirationDate,omitempty"`
}

type AuthenticateResult struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com AuthenticateResult"`

	Errors *ArrayOfString `xml:"Errors,omitempty"`

	Token string `xml:"Token,omitempty"`
}

type ArrayOfString struct {
	//REMOVED(xml: name "Errors" in tag of thycotic.GetSecretResult.Errors conflicts with name "ArrayOfString" in *thycotic.ArrayOfString.XMLName) XMLName xml.Name `xml:"urn:thesecretserver.com ArrayOfString"`

	String []string `xml:"string,omitempty"`
}

type ImpersonateResult struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com ImpersonateResult"`

	Errors *ArrayOfString `xml:"Errors,omitempty"`

	Token string `xml:"Token,omitempty"`

	AuthorizeURL string `xml:"AuthorizeURL,omitempty"`
}

type TokenIsValidResult struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com TokenIsValidResult"`

	Errors *ArrayOfString `xml:"Errors,omitempty"`

	MaxOfflineSeconds int32 `xml:"MaxOfflineSeconds,omitempty"`

	Version string `xml:"Version,omitempty"`
}

type GetSecretResult struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GetSecretResult"`

	Errors *ArrayOfString `xml:"Errors,omitempty"`

	SecretError *SecretError `xml:"SecretError,omitempty"`

	Secret *Secret `xml:"Secret,omitempty"`
}

type SecretError struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com SecretError"`

	ErrorCode string `xml:"ErrorCode,omitempty"`

	ErrorMessage string `xml:"ErrorMessage,omitempty"`

	AllowsResponse bool `xml:"AllowsResponse,omitempty"`

	CommentTitle string `xml:"CommentTitle,omitempty"`

	AdditionalCommentTitle string `xml:"AdditionalCommentTitle,omitempty"`
}

type Secret struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com Secret"`

	Name string `xml:"Name,omitempty"`

	Items *ArrayOfSecretItem `xml:"Items,omitempty"`

	Id int32 `xml:"Id,omitempty"`

	SecretTypeId int32 `xml:"SecretTypeId,omitempty"`

	FolderId int32 `xml:"FolderId,omitempty"`

	IsWebLauncher bool `xml:"IsWebLauncher,omitempty"`

	//REMOVED(strconv.ParseInt: parsing "": invalid syntax) CheckOutMinutesRemaining int32 `xml:"CheckOutMinutesRemaining,omitempty"`

	//REMOVED(strconv.ParseBool: parsing "": invalid syntax) IsCheckedOut bool `xml:"IsCheckedOut,omitempty"`

	CheckOutUserDisplayName string `xml:"CheckOutUserDisplayName,omitempty"`

	//REMOVED(strconv.ParseInt: parsing "": invalid syntax) CheckOutUserId int32 `xml:"CheckOutUserId,omitempty"`

	//REMOVED(strconv.ParseBool: parsing "": invalid syntax) IsOutOfSync bool `xml:"IsOutOfSync,omitempty"`

	IsRestricted bool `xml:"IsRestricted,omitempty"`

	OutOfSyncReason string `xml:"OutOfSyncReason,omitempty"`

	SecretSettings *SecretSettings `xml:"SecretSettings,omitempty"`

	SecretPermissions *SecretPermissions `xml:"SecretPermissions,omitempty"`

	Active bool `xml:"Active,omitempty"`
}

type ArrayOfSecretItem struct {
	//REMOVED(xml: name "Items" in tag of thycotic.Secret.Items conflicts with name "ArrayOfSecretItem" in *thycotic.ArrayOfSecretItem.XMLName) XMLName xml.Name `xml:"urn:thesecretserver.com ArrayOfSecretItem"`

	SecretItem []*SecretItem `xml:"SecretItem,omitempty"`
}

type SecretItem struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com SecretItem"`

	Value string `xml:"Value,omitempty"`

	Id int32 `xml:"Id,omitempty"`

	FieldId int32 `xml:"FieldId,omitempty"`

	FieldName string `xml:"FieldName,omitempty"`

	IsFile bool `xml:"IsFile,omitempty"`

	IsNotes bool `xml:"IsNotes,omitempty"`

	IsPassword bool `xml:"IsPassword,omitempty"`

	FieldDisplayName string `xml:"FieldDisplayName,omitempty"`
}

type SecretSettings struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com SecretSettings"`

	AutoChangeEnabled bool `xml:"AutoChangeEnabled,omitempty"`

	RequiresApprovalForAccess bool `xml:"RequiresApprovalForAccess,omitempty"`

	RequiresComment bool `xml:"RequiresComment,omitempty"`

	CheckOutEnabled bool `xml:"CheckOutEnabled,omitempty"`

	CheckOutChangePasswordEnabled bool `xml:"CheckOutChangePasswordEnabled,omitempty"`

	ProxyEnabled bool `xml:"ProxyEnabled,omitempty"`

	SessionRecordingEnabled bool `xml:"SessionRecordingEnabled,omitempty"`

	RestrictSshCommands bool `xml:"RestrictSshCommands,omitempty"`

	AllowOwnersUnrestrictedSshCommands bool `xml:"AllowOwnersUnrestrictedSshCommands,omitempty"`

	PrivilegedSecretId int32 `xml:"PrivilegedSecretId,omitempty"`

	AssociatedSecretIds *ArrayOfInt `xml:"AssociatedSecretIds,omitempty"`

	Approvers *ArrayOfGroupOrUserRecord `xml:"Approvers,omitempty"`

	SshCommandMenuAccessPermissions *ArrayOfSshCommandMenuAccessPermission `xml:"SshCommandMenuAccessPermissions,omitempty"`

	IsChangeToSettings bool `xml:"IsChangeToSettings,omitempty"`
}

type ArrayOfInt struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com ArrayOfInt"`

	Int []int32 `xml:"int,omitempty"`
}

type ArrayOfGroupOrUserRecord struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com ArrayOfGroupOrUserRecord"`

	GroupOrUserRecord []*GroupOrUserRecord `xml:"GroupOrUserRecord,omitempty"`
}

type GroupOrUserRecord struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GroupOrUserRecord"`

	Name string `xml:"Name,omitempty"`

	DomainName string `xml:"DomainName,omitempty"`

	IsUser bool `xml:"IsUser,omitempty"`

	GroupId int32 `xml:"GroupId,omitempty"`

	UserId int32 `xml:"UserId,omitempty"`
}

type ArrayOfSshCommandMenuAccessPermission struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com ArrayOfSshCommandMenuAccessPermission"`

	SshCommandMenuAccessPermission []*SshCommandMenuAccessPermission `xml:"SshCommandMenuAccessPermission,omitempty"`
}

type SshCommandMenuAccessPermission struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com SshCommandMenuAccessPermission"`

	GroupOrUserRecord *GroupOrUserRecord `xml:"GroupOrUserRecord,omitempty"`

	SecretId int32 `xml:"SecretId,omitempty"`

	ConcurrencyId string `xml:"ConcurrencyId,omitempty"`

	DisplayName string `xml:"DisplayName,omitempty"`

	SshCommandMenuName string `xml:"SshCommandMenuName,omitempty"`

	IsUnrestricted bool `xml:"IsUnrestricted,omitempty"`

	SshCommandMenuId int32 `xml:"SshCommandMenuId,omitempty"`
}

type SecretPermissions struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com SecretPermissions"`

	CurrentUserHasView bool `xml:"CurrentUserHasView,omitempty"`

	CurrentUserHasEdit bool `xml:"CurrentUserHasEdit,omitempty"`

	CurrentUserHasOwner bool `xml:"CurrentUserHasOwner,omitempty"`

	InheritPermissionsEnabled bool `xml:"InheritPermissionsEnabled,omitempty"`

	IsChangeToPermissions bool `xml:"IsChangeToPermissions,omitempty"`

	Permissions *ArrayOfPermission `xml:"Permissions,omitempty"`
}

type ArrayOfPermission struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com ArrayOfPermission"`

	Permission []*Permission `xml:"Permission,omitempty"`
}

type Permission struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com Permission"`

	UserOrGroup *GroupOrUserRecord `xml:"UserOrGroup,omitempty"`

	View bool `xml:"View,omitempty"`

	Edit bool `xml:"Edit,omitempty"`

	Owner bool `xml:"Owner,omitempty"`

	SecretAccessRoleName string `xml:"SecretAccessRoleName,omitempty"`

	SecretAccessRoleId int32 `xml:"SecretAccessRoleId,omitempty"`
}

type ArrayOfCodeResponse struct {
//	XMLName xml.Name `xml:"urn:thesecretserver.com ArrayOfCodeResponse"`

	CodeResponse []*CodeResponse `xml:"CodeResponse,omitempty"`
}

type CodeResponse struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com CodeResponse"`

	ErrorCode string `xml:"ErrorCode,omitempty"`

	Comment string `xml:"Comment,omitempty"`

	AdditionalComment string `xml:"AdditionalComment,omitempty"`

	TicketSystemId int32 `xml:"TicketSystemId,omitempty"`
}

type GetCheckOutStatusResult struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GetCheckOutStatusResult"`

	Errors *ArrayOfString `xml:"Errors,omitempty"`

	Secret *Secret `xml:"Secret,omitempty"`

	CheckOutMinutesRemaining int32 `xml:"CheckOutMinutesRemaining,omitempty"`

	IsCheckedOut bool `xml:"IsCheckedOut,omitempty"`

	CheckOutUserDisplayName string `xml:"CheckOutUserDisplayName,omitempty"`

	CheckOutUserId int32 `xml:"CheckOutUserId,omitempty"`
}

type WebServiceResult struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com WebServiceResult"`

	Errors *ArrayOfString `xml:"Errors,omitempty"`
}

type GetSecretsByFieldValueResult struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GetSecretsByFieldValueResult"`

	Errors *ArrayOfString `xml:"Errors,omitempty"`

	Secrets *ArrayOfSecret `xml:"Secrets,omitempty"`
}

type ArrayOfSecret struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com ArrayOfSecret"`

	Secret []*Secret `xml:"Secret,omitempty"`
}

type SearchSecretsResult struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com SearchSecretsResult"`

	Errors *ArrayOfString `xml:"Errors,omitempty"`

	SecretSummaries *ArrayOfSecretSummary `xml:"SecretSummaries,omitempty"`
}

type ArrayOfSecretSummary struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com ArrayOfSecretSummary"`

	SecretSummary []*SecretSummary `xml:"SecretSummary,omitempty"`
}

type SecretSummary struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com SecretSummary"`

	SecretId int32 `xml:"SecretId,omitempty"`

	SecretName string `xml:"SecretName,omitempty"`

	SecretTypeName string `xml:"SecretTypeName,omitempty"`

	SecretTypeId int32 `xml:"SecretTypeId,omitempty"`

	FolderId int32 `xml:"FolderId,omitempty"`

	IsRestricted bool `xml:"IsRestricted,omitempty"`
}

type User struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com User"`

	Id int32 `xml:"Id,omitempty"`

	UserName string `xml:"UserName,omitempty"`

	DisplayName string `xml:"DisplayName,omitempty"`

	DomainId int32 `xml:"DomainId,omitempty"`

	IsApplicationAccount bool `xml:"IsApplicationAccount,omitempty"`

	RadiusTwoFactor bool `xml:"RadiusTwoFactor,omitempty"`

	EmailTwoFactor bool `xml:"EmailTwoFactor,omitempty"`

	RadiusUserName string `xml:"RadiusUserName,omitempty"`

	EmailAddress string `xml:"EmailAddress,omitempty"`

	Password string `xml:"Password,omitempty"`

	Enabled bool `xml:"Enabled,omitempty"`

	DuoTwoFactor bool `xml:"DuoTwoFactor,omitempty"`

	OATHTwoFactor bool `xml:"OATHTwoFactor,omitempty"`
}

type GetFavoritesResult struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GetFavoritesResult"`

	Errors *ArrayOfString `xml:"Errors,omitempty"`

	SecretSummaries *ArrayOfSecretSummary `xml:"SecretSummaries,omitempty"`
}

type AddSecretResult struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com AddSecretResult"`

	Errors *ArrayOfString `xml:"Errors,omitempty"`

	Secret *Secret `xml:"Secret,omitempty"`
}

type GetSecretTemplateFieldsResult struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GetSecretTemplateFieldsResult"`

	Errors *ArrayOfString `xml:"Errors,omitempty"`

	Fields *ArrayOfSecretField `xml:"Fields,omitempty"`
}

type ArrayOfSecretField struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com ArrayOfSecretField"`

	SecretField []*SecretField `xml:"SecretField,omitempty"`
}

type SecretField struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com SecretField"`

	DisplayName string `xml:"DisplayName,omitempty"`

	Id int32 `xml:"Id,omitempty"`

	IsPassword bool `xml:"IsPassword,omitempty"`

	IsUrl bool `xml:"IsUrl,omitempty"`

	IsNotes bool `xml:"IsNotes,omitempty"`

	IsFile bool `xml:"IsFile,omitempty"`
}

type GetSecretTemplatesResult struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GetSecretTemplatesResult"`

	Errors *ArrayOfString `xml:"Errors,omitempty"`

	SecretTemplates *ArrayOfSecretTemplate `xml:"SecretTemplates,omitempty"`
}

type ArrayOfSecretTemplate struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com ArrayOfSecretTemplate"`

	SecretTemplate []*SecretTemplate `xml:"SecretTemplate,omitempty"`
}

type SecretTemplate struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com SecretTemplate"`

	Id int32 `xml:"Id,omitempty"`

	Name string `xml:"Name,omitempty"`

	Fields *ArrayOfSecretField `xml:"Fields,omitempty"`
}

type GeneratePasswordResult struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GeneratePasswordResult"`

	GeneratedPassword string `xml:"GeneratedPassword,omitempty"`

	Errors *ArrayOfString `xml:"Errors,omitempty"`
}

type VersionGetResult struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com VersionGetResult"`

	Errors *ArrayOfString `xml:"Errors,omitempty"`

	Version string `xml:"Version,omitempty"`
}

type GetFolderResult struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GetFolderResult"`

	Folder *Folder `xml:"Folder,omitempty"`

	Errors *ArrayOfString `xml:"Errors,omitempty"`

	Success bool `xml:"Success,omitempty"`
}

type Folder struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com Folder"`

	Id int32 `xml:"Id,omitempty"`

	Name string `xml:"Name,omitempty"`

	TypeId int32 `xml:"TypeId,omitempty"`

	ParentFolderId int32 `xml:"ParentFolderId,omitempty"`
}

type GetFoldersResult struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GetFoldersResult"`

	Folders *ArrayOfFolder `xml:"Folders,omitempty"`

	Errors *ArrayOfString `xml:"Errors,omitempty"`

	Success bool `xml:"Success,omitempty"`
}

type ArrayOfFolder struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com ArrayOfFolder"`

	Folder []*Folder `xml:"Folder,omitempty"`
}

type CreateFolderResult struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com CreateFolderResult"`

	*WebServiceResult

	FolderId int32 `xml:"FolderId,omitempty"`
}

type FolderExtended struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com FolderExtended"`

	*Folder

	PermissionSettings *FolderPermissions `xml:"PermissionSettings,omitempty"`

	Settings *FolderSettings `xml:"Settings,omitempty"`
}

type FolderPermissions struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com FolderPermissions"`

	IsChangeToPermissions bool `xml:"IsChangeToPermissions,omitempty"`

	InheritPermissionsEnabled bool `xml:"InheritPermissionsEnabled,omitempty"`

	Permissions *ArrayOfFolderPermission `xml:"Permissions,omitempty"`
}

type ArrayOfFolderPermission struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com ArrayOfFolderPermission"`

	FolderPermission []*FolderPermission `xml:"FolderPermission,omitempty"`
}

type FolderPermission struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com FolderPermission"`

	UserOrGroup *GroupOrUserRecord `xml:"UserOrGroup,omitempty"`

	FolderAccessRoleName string `xml:"FolderAccessRoleName,omitempty"`

	FolderAccessRoleId int32 `xml:"FolderAccessRoleId,omitempty"`

	SecretAccessRoleName string `xml:"SecretAccessRoleName,omitempty"`

	SecretAccessRoleId int32 `xml:"SecretAccessRoleId,omitempty"`
}

type FolderSettings struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com FolderSettings"`

	IsChangeToSettings bool `xml:"IsChangeToSettings,omitempty"`

	InheritSecretPolicy bool `xml:"InheritSecretPolicy,omitempty"`

	SecretPolicyId int32 `xml:"SecretPolicyId,omitempty"`
}

type FolderExtendedCreateResult struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com FolderExtendedCreateResult"`

	*FolderExtendedResultBase

	FolderId int32 `xml:"FolderId,omitempty"`
}

type FolderExtendedResultBase struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com FolderExtendedResultBase"`

	Success bool `xml:"Success,omitempty"`

	Errors *ArrayOfString `xml:"Errors,omitempty"`
}

type FolderExtendedGetResult struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com FolderExtendedGetResult"`

	*FolderExtendedResultBase

	Folder *FolderExtended `xml:"Folder,omitempty"`
}

type FolderExtendedUpdateResult struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com FolderExtendedUpdateResult"`

	*FolderExtendedResultBase
}

type FolderExtendedGetNewRequest struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com FolderExtendedGetNewRequest"`

	FolderName string `xml:"FolderName,omitempty"`

	ParentFolderId int32 `xml:"ParentFolderId,omitempty"`

	InheritPermissions bool `xml:"InheritPermissions,omitempty"`
}

type FolderExtendedGetNewResult struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com FolderExtendedGetNewResult"`

	*FolderExtendedResultBase

	Folder *FolderExtended `xml:"Folder,omitempty"`
}

type SearchFolderResult struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com SearchFolderResult"`

	Errors *ArrayOfString `xml:"Errors,omitempty"`

	Folders *ArrayOfFolder `xml:"Folders,omitempty"`
}

type FileDownloadResult struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com FileDownloadResult"`

	*WebServiceResult

	FileAttachment []byte `xml:"FileAttachment,omitempty"`

	FileName string `xml:"FileName,omitempty"`
}

type GetSecretAuditResult struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GetSecretAuditResult"`

	Errors *ArrayOfString `xml:"Errors,omitempty"`

	SecretAudits *ArrayOfAuditSecret `xml:"SecretAudits,omitempty"`
}

type ArrayOfAuditSecret struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com ArrayOfAuditSecret"`

	AuditSecret []*AuditSecret `xml:"AuditSecret,omitempty"`
}

type AuditSecret struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com AuditSecret"`

	AuditSecretId int32 `xml:"AuditSecretId,omitempty"`

	SecretId int32 `xml:"SecretId,omitempty"`

	DateRecorded time.Time `xml:"DateRecorded,omitempty"`

	Action string `xml:"Action,omitempty"`

	Notes string `xml:"Notes,omitempty"`

	UserId int32 `xml:"UserId,omitempty"`

	SecretName string `xml:"SecretName,omitempty"`

	IpAddress string `xml:"IpAddress,omitempty"`

	ReferenceId int32 `xml:"ReferenceId,omitempty"`

	ByUserDisplayName string `xml:"ByUserDisplayName,omitempty"`

	TicketNumber string `xml:"TicketNumber,omitempty"`
}

type Dependency struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com Dependency"`

	SecretId int32 `xml:"SecretId,omitempty"`

	SecretDependencyTypeId int32 `xml:"SecretDependencyTypeId,omitempty"`

	MachineName string `xml:"MachineName,omitempty"`

	ServiceName string `xml:"ServiceName,omitempty"`

	PrivilegedAccountSecretId int32 `xml:"PrivilegedAccountSecretId,omitempty"`

	Active bool `xml:"Active,omitempty"`

	RestartOnPasswordChange bool `xml:"RestartOnPasswordChange,omitempty"`

	WaitBeforeSeconds int32 `xml:"WaitBeforeSeconds,omitempty"`

	AdditionalInfo *AdditionalDependencyInfoJson `xml:"AdditionalInfo,omitempty"`

	Description string `xml:"Description,omitempty"`

	ScriptId int32 `xml:"ScriptId,omitempty"`

	SecretDependencyId int32 `xml:"SecretDependencyId,omitempty"`

	SSHKeySecretId int32 `xml:"SSHKeySecretId,omitempty"`
}

type AdditionalDependencyInfoJson struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com AdditionalDependencyInfoJson"`

	Regex string `xml:"Regex,omitempty"`

	PowershellArguments string `xml:"PowershellArguments,omitempty"`

	SshArguments *ArrayOfSshScriptArgument `xml:"SshArguments,omitempty"`

	SqlArguments *ArrayOfSqlScriptArgument `xml:"SqlArguments,omitempty"`

	OdbcConnectionArguments *ArrayOfOdbcConnectionArg `xml:"OdbcConnectionArguments,omitempty"`

	Port string `xml:"Port,omitempty"`

	Database string `xml:"Database,omitempty"`

	ServerKeyDigest string `xml:"ServerKeyDigest,omitempty"`
}

type ArrayOfSshScriptArgument struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com ArrayOfSshScriptArgument"`

	SshScriptArgument []*SshScriptArgument `xml:"SshScriptArgument,omitempty"`
}

type SshScriptArgument struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com SshScriptArgument"`

	Name string `xml:"Name,omitempty"`

	Value string `xml:"Value,omitempty"`

	SshType *SshArgumentType `xml:"SshType,omitempty"`
}

type ArrayOfSqlScriptArgument struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com ArrayOfSqlScriptArgument"`

	SqlScriptArgument []*SqlScriptArgument `xml:"SqlScriptArgument,omitempty"`
}

type SqlScriptArgument struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com SqlScriptArgument"`

	Name string `xml:"Name,omitempty"`

	Value struct {
	} `xml:"Value,omitempty"`

	DbType *DbType `xml:"DbType,omitempty"`
}

type ArrayOfOdbcConnectionArg struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com ArrayOfOdbcConnectionArg"`

	OdbcConnectionArg []*OdbcConnectionArg `xml:"OdbcConnectionArg,omitempty"`
}

type OdbcConnectionArg struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com OdbcConnectionArg"`

	Name string `xml:"Name,omitempty"`

	Value string `xml:"Value,omitempty"`
}

type GetDependenciesResult struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GetDependenciesResult"`

	Dependencies *ArrayOfDependency `xml:"Dependencies,omitempty"`

	Errors *ArrayOfString `xml:"Errors,omitempty"`

	Success bool `xml:"Success,omitempty"`
}

type ArrayOfDependency struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com ArrayOfDependency"`

	Dependency []*Dependency `xml:"Dependency,omitempty"`
}

type GetSitesResult struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GetSitesResult"`

	Sites *ArrayOfSite `xml:"Sites,omitempty"`

	Errors *ArrayOfString `xml:"Errors,omitempty"`

	Success bool `xml:"Success,omitempty"`
}

type ArrayOfSite struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com ArrayOfSite"`

	Site []*Site `xml:"Site,omitempty"`
}

type Site struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com Site"`

	SiteId int32 `xml:"SiteId,omitempty"`

	OrganizationId int32 `xml:"OrganizationId,omitempty"`

	SymmetricKey string `xml:"SymmetricKey,omitempty"`

	SymmetricKeyIV []byte `xml:"SymmetricKeyIV,omitempty"`

	InitializationVector []byte `xml:"InitializationVector,omitempty"`

	SiteName string `xml:"SiteName,omitempty"`

	Active bool `xml:"Active,omitempty"`

	HeartbeatInterval int32 `xml:"HeartbeatInterval,omitempty"`

	UseWebSite bool `xml:"UseWebSite,omitempty"`

	SystemSite bool `xml:"SystemSite,omitempty"`

	LastModifiedDate time.Time `xml:"LastModifiedDate,omitempty"`

	WinRMEndpoint string `xml:"WinRMEndpoint,omitempty"`

	EnableCredSSPForWinRM bool `xml:"EnableCredSSPForWinRM,omitempty"`

	SiteConnectorId int32 `xml:"SiteConnectorId,omitempty"`

	SiteConnector *SiteConnector `xml:"SiteConnector,omitempty"`
}

type SiteConnector struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com SiteConnector"`

	SiteConnectorId int32 `xml:"SiteConnectorId,omitempty"`

	SiteConnectorName string `xml:"SiteConnectorName,omitempty"`

	QueueType string `xml:"QueueType,omitempty"`

	HostName string `xml:"HostName,omitempty"`

	Port int32 `xml:"Port,omitempty"`

	Active bool `xml:"Active,omitempty"`

	Validated bool `xml:"Validated,omitempty"`

	UseSsl bool `xml:"UseSsl,omitempty"`

	SslCertificateThumbprint string `xml:"SslCertificateThumbprint,omitempty"`

	LastModifiedDate time.Time `xml:"LastModifiedDate,omitempty"`

	UserName string `xml:"UserName,omitempty"`

	PasswordIV []byte `xml:"PasswordIV,omitempty"`

	Version string `xml:"Version,omitempty"`
}

type GetTicketSystemsResult struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GetTicketSystemsResult"`

	TicketSystems *ArrayOfTicketSystem `xml:"TicketSystems,omitempty"`

	Errors *ArrayOfString `xml:"Errors,omitempty"`

	Success bool `xml:"Success,omitempty"`
}

type ArrayOfTicketSystem struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com ArrayOfTicketSystem"`

	TicketSystem []*TicketSystem `xml:"TicketSystem,omitempty"`
}

type TicketSystem struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com TicketSystem"`

	TicketSystemId int32 `xml:"TicketSystemId,omitempty"`

	Name string `xml:"Name,omitempty"`

	Description string `xml:"Description,omitempty"`

	IsDefault bool `xml:"IsDefault,omitempty"`
}

type UserInfoResult struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com UserInfoResult"`

	Errors *ArrayOfString `xml:"Errors,omitempty"`

	DisplayName string `xml:"DisplayName,omitempty"`

	UserName string `xml:"UserName,omitempty"`

	KnownAs string `xml:"KnownAs,omitempty"`

	UserId int32 `xml:"UserId,omitempty"`

	DomainId int32 `xml:"DomainId,omitempty"`

	DomainName string `xml:"DomainName,omitempty"`
}

type GetAllGroupsResult struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GetAllGroupsResult"`

	Errors *ArrayOfString `xml:"Errors,omitempty"`

	Groups *ArrayOfGroup `xml:"Groups,omitempty"`
}

type ArrayOfGroup struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com ArrayOfGroup"`

	Group []*Group `xml:"Group,omitempty"`
}

type Group struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com Group"`

	Id int32 `xml:"Id,omitempty"`

	Name string `xml:"Name,omitempty"`

	DomainId int32 `xml:"DomainId,omitempty"`

	DomainName string `xml:"DomainName,omitempty"`
}

type SSHCredentialsResult struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com SSHCredentialsResult"`

	*WebServiceResult

	Username string `xml:"Username,omitempty"`

	Password string `xml:"Password,omitempty"`

	Host string `xml:"Host,omitempty"`

	Port string `xml:"Port,omitempty"`
}

type GetUsersResult struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GetUsersResult"`

	Users *ArrayOfUser `xml:"Users,omitempty"`

	Errors *ArrayOfString `xml:"Errors,omitempty"`
}

type ArrayOfUser struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com ArrayOfUser"`

	User []*User `xml:"User,omitempty"`
}

type GetUserResult struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GetUserResult"`

	User *User `xml:"User,omitempty"`

	Errors *ArrayOfString `xml:"Errors,omitempty"`
}

type UpdateUserResult struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com UpdateUserResult"`

	User *User `xml:"User,omitempty"`

	Errors *ArrayOfString `xml:"Errors,omitempty"`
}

type SecretItemHistoryResult struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com SecretItemHistoryResult"`

	Errors *ArrayOfString `xml:"Errors,omitempty"`

	Success bool `xml:"Success,omitempty"`

	SecretItemHistories *ArrayOfSecretItemHistoryWebServiceResult `xml:"SecretItemHistories,omitempty"`
}

type ArrayOfSecretItemHistoryWebServiceResult struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com ArrayOfSecretItemHistoryWebServiceResult"`

	SecretItemHistoryWebServiceResult []*SecretItemHistoryWebServiceResult `xml:"SecretItemHistoryWebServiceResult,omitempty"`
}

type SecretItemHistoryWebServiceResult struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com SecretItemHistoryWebServiceResult"`

	SecretItemHistoryId int32 `xml:"SecretItemHistoryId,omitempty"`

	UserId int32 `xml:"UserId,omitempty"`

	SecretItemId int32 `xml:"SecretItemId,omitempty"`

	SecretId int32 `xml:"SecretId,omitempty"`

	Date time.Time `xml:"Date,omitempty"`

	ItemValueNew string `xml:"ItemValueNew,omitempty"`

	ItemValueNew2 string `xml:"ItemValueNew2,omitempty"`
}

type SecretPolicyForSecret struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com SecretPolicyForSecret"`

	SecretId int32 `xml:"SecretId,omitempty"`

	SecretPolicyId int32 `xml:"SecretPolicyId,omitempty"`

	Inherit bool `xml:"Inherit,omitempty"`
}

type SecretPolicyForSecretResult struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com SecretPolicyForSecretResult"`

	*WebServiceResult

	SecretPolicyForSecret *SecretPolicyForSecret `xml:"SecretPolicyForSecret,omitempty"`
}

type SecretPolicySummary struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com SecretPolicySummary"`

	SecretPolicyId int32 `xml:"SecretPolicyId,omitempty"`

	SecretPolicyName string `xml:"SecretPolicyName,omitempty"`

	SecretPolicyDescription string `xml:"SecretPolicyDescription,omitempty"`

	Active bool `xml:"Active,omitempty"`
}

type ArrayOfSecretPolicySummary struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com ArrayOfSecretPolicySummary"`

	SecretPolicySummary []*SecretPolicySummary `xml:"SecretPolicySummary,omitempty"`
}

type SearchSecretPoliciesResult struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com SearchSecretPoliciesResult"`

	*WebServiceResult

	SecretPolicies *ArrayOfSecretPolicySummary `xml:"SecretPolicies,omitempty"`
}

type AddGroupRequestMessage struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com AddGroupRequestMessage"`

	GroupName string `xml:"GroupName,omitempty"`

	DomainId int32 `xml:"DomainId,omitempty"`

	DomainName string `xml:"DomainName,omitempty"`
}

type SecretPolicyDetail struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com SecretPolicyDetail"`

	*SecretPolicySummary

	SecretPolicyItems *ArrayOfSecretPolicyItem `xml:"SecretPolicyItems,omitempty"`
}

type ArrayOfSecretPolicyItem struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com ArrayOfSecretPolicyItem"`

	SecretPolicyItem []*SecretPolicyItem `xml:"SecretPolicyItem,omitempty"`
}

type SecretPolicyItem struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com SecretPolicyItem"`

	SecretPolicyItemMapId int32 `xml:"SecretPolicyItemMapId,omitempty"`

	SecretPolicyItemId int32 `xml:"SecretPolicyItemId,omitempty"`

	PolicyApplyCode string `xml:"PolicyApplyCode,omitempty"`

	EnabledValue bool `xml:"EnabledValue,omitempty"`

	IntegerValue int32 `xml:"IntegerValue,omitempty"`

	SecretId int32 `xml:"SecretId,omitempty"`

	StringValue string `xml:"StringValue,omitempty"`

	Name string `xml:"Name,omitempty"`

	Description string `xml:"Description,omitempty"`

	ValueType string `xml:"ValueType,omitempty"`

	ParentSecretPolicyItemId int32 `xml:"ParentSecretPolicyItemId,omitempty"`

	SectionName string `xml:"SectionName,omitempty"`

	UserGroupMaps *ArrayOfUserGroupMap `xml:"UserGroupMaps,omitempty"`

	SshCommandMenuGroupMaps *ArrayOfSshCommandMenuGroupMap `xml:"SshCommandMenuGroupMaps,omitempty"`
}

type ArrayOfUserGroupMap struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com ArrayOfUserGroupMap"`

	UserGroupMap []*UserGroupMap `xml:"UserGroupMap,omitempty"`
}

type UserGroupMap struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com UserGroupMap"`

	Id int32 `xml:"Id,omitempty"`

	UserGroupMapType *UserGroupMapType `xml:"UserGroupMapType,omitempty"`
}

type ArrayOfSshCommandMenuGroupMap struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com ArrayOfSshCommandMenuGroupMap"`

	SshCommandMenuGroupMap []*SshCommandMenuGroupMap `xml:"SshCommandMenuGroupMap,omitempty"`
}

type SshCommandMenuGroupMap struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com SshCommandMenuGroupMap"`

	SshCommandMenuId int32 `xml:"SshCommandMenuId,omitempty"`

	UserGroupMap *UserGroupMap `xml:"UserGroupMap,omitempty"`
}

type SecretPolicyResult struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com SecretPolicyResult"`

	*WebServiceResult

	SecretPolicy *SecretPolicyDetail `xml:"SecretPolicy,omitempty"`
}

type GetSshCommandMenuResult struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GetSshCommandMenuResult"`

	SshCommandMenu *SshCommandMenu `xml:"SshCommandMenu,omitempty"`

	SshCommands string `xml:"SshCommands,omitempty"`

	Errors *ArrayOfString `xml:"Errors,omitempty"`
}

type SshCommandMenu struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com SshCommandMenu"`

	SshCommandMenuId int32 `xml:"SshCommandMenuId,omitempty"`

	Name string `xml:"Name,omitempty"`

	Active bool `xml:"Active,omitempty"`

	Description string `xml:"Description,omitempty"`
}

type GetSshCommandMenusResult struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GetSshCommandMenusResult"`

	SshCommandMenus *ArrayOfSshCommandMenu `xml:"SshCommandMenus,omitempty"`

	Errors *ArrayOfString `xml:"Errors,omitempty"`
}

type ArrayOfSshCommandMenu struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com ArrayOfSshCommandMenu"`

	SshCommandMenu []*SshCommandMenu `xml:"SshCommandMenu,omitempty"`
}

type UserScript struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com UserScript"`

	ScriptId int32 `xml:"ScriptId,omitempty"`

	Name string `xml:"Name,omitempty"`

	Description string `xml:"Description,omitempty"`

	Script string `xml:"Script,omitempty"`

	Active bool `xml:"Active,omitempty"`
}

type SshUserScript struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com SshUserScript"`

	*UserScript

	AdditionalDataObject *AdditionalDataSshObject `xml:"AdditionalDataObject,omitempty"`
}

type AdditionalDataSshObject struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com AdditionalDataSshObject"`

	Port string `xml:"Port,omitempty"`

	LineEnding *LineEnding `xml:"LineEnding,omitempty"`

	Params *ArrayOfSshScriptArgument2 `xml:"Params,omitempty"`

	Version int32 `xml:"Version,omitempty"`
}

type ArrayOfSshScriptArgument2 struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com ArrayOfSshScriptArgument2"`

	SshScriptArgument2 []*SshScriptArgument2 `xml:"SshScriptArgument2,omitempty"`
}

type SshScriptArgument2 struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com SshScriptArgument2"`

	Name string `xml:"Name,omitempty"`

	Value string `xml:"Value,omitempty"`

	SshType *SshArgumentType2 `xml:"SshType,omitempty"`
}

type PowerShellUserScript struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com PowerShellUserScript"`

	*UserScript
}

type SqlUserScript struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com SqlUserScript"`

	*UserScript

	AdditionalDataObject *AdditionalDataSqlObject `xml:"AdditionalDataObject,omitempty"`
}

type AdditionalDataSqlObject struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com AdditionalDataSqlObject"`

	Params *ArrayOfSqlScriptArgument2 `xml:"Params,omitempty"`

	PasswordChangerId int32 `xml:"PasswordChangerId,omitempty"`

	Version int32 `xml:"Version,omitempty"`

	Database string `xml:"Database,omitempty"`

	Port string `xml:"Port,omitempty"`
}

type ArrayOfSqlScriptArgument2 struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com ArrayOfSqlScriptArgument2"`

	SqlScriptArgument2 []*SqlScriptArgument2 `xml:"SqlScriptArgument2,omitempty"`
}

type SqlScriptArgument2 struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com SqlScriptArgument2"`

	Name string `xml:"Name,omitempty"`

	Value struct {
	} `xml:"Value,omitempty"`

	DbType *DbType `xml:"DbType,omitempty"`
}

type ArrayOfUserScript struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com ArrayOfUserScript"`

	UserScript []*UserScript `xml:"UserScript,omitempty"`
}

type GetUserScriptsResult struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GetUserScriptsResult"`

	*WebServiceResult

	UserScripts *ArrayOfUserScript `xml:"UserScripts,omitempty"`
}

type GetUserScriptResult struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com GetUserScriptResult"`

	*WebServiceResult

	UserScript *UserScript `xml:"UserScript,omitempty"`
}

type UpdateUserScriptResult struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com UpdateUserScriptResult"`

	*WebServiceResult

	UserScript *UserScript `xml:"UserScript,omitempty"`
}

type StringArray struct {
	XMLName xml.Name `xml:"urn:thesecretserver.com/AbstractTypes StringArray"`
}

type SSWebServiceSoap struct {
	client *SOAPClient
}

func NewSSWebServiceSoap(url string, tls bool, auth *BasicAuth) *SSWebServiceSoap {
	if url == "" {
		url = "https://secret.example.com/SecretServer/webservices/sswebservice.asmx"
	}
	client := NewSOAPClient(url, tls, auth)

	return &SSWebServiceSoap{
		client: client,
	}
}

func NewSSWebServiceSoapWithTLSConfig(url string, tlsCfg *tls.Config, auth *BasicAuth) *SSWebServiceSoap {
	if url == "" {
		url = "https://secret.example.com/SecretServer/webservices/sswebservice.asmx"
	}
	client := NewSOAPClientWithTLSConfig(url, tlsCfg, auth)

	return &SSWebServiceSoap{
		client: client,
	}
}

func (service *SSWebServiceSoap) AddHeader(header interface{}) {
	service.client.AddHeader(header)
}

// Backwards-compatible function: use AddHeader instead
func (service *SSWebServiceSoap) SetHeader(header interface{}) {
	service.client.AddHeader(header)
}

func (service *SSWebServiceSoap) ApproveSecretAccessRequest(request *ApproveSecretAccessRequest) (*ApproveSecretAccessRequestResponse, error) {
	response := new(ApproveSecretAccessRequestResponse)
	err := service.client.Call("urn:thesecretserver.com/ApproveSecretAccessRequest", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) DenySecretAccessRequest(request *DenySecretAccessRequest) (*DenySecretAccessRequestResponse, error) {
	response := new(DenySecretAccessRequestResponse)
	err := service.client.Call("urn:thesecretserver.com/DenySecretAccessRequest", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) Authenticate(request *Authenticate) (*AuthenticateResponse, error) {
	response := new(AuthenticateResponse)
	err := service.client.Call("urn:thesecretserver.com/Authenticate", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) ImpersonateUser(request *ImpersonateUser) (*ImpersonateUserResponse, error) {
	response := new(ImpersonateUserResponse)
	err := service.client.Call("urn:thesecretserver.com/ImpersonateUser", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) AuthenticateRADIUS(request *AuthenticateRADIUS) (*AuthenticateRADIUSResponse, error) {
	response := new(AuthenticateRADIUSResponse)
	err := service.client.Call("urn:thesecretserver.com/AuthenticateRADIUS", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) GetTokenIsValid(request *GetTokenIsValid) (*GetTokenIsValidResponse, error) {
	response := new(GetTokenIsValidResponse)
	err := service.client.Call("urn:thesecretserver.com/GetTokenIsValid", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) GetSecretLegacy(request *GetSecretLegacy) (*GetSecretLegacyResponse, error) {
	response := new(GetSecretLegacyResponse)
	err := service.client.Call("urn:thesecretserver.com/GetSecretLegacy", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) GetSecret(request *GetSecret) (*GetSecretResponse, error) {
	response := new(GetSecretResponse)
	err := service.client.Call("urn:thesecretserver.com/GetSecret", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) GetCheckOutStatus(request *GetCheckOutStatus) (*GetCheckOutStatusResponse, error) {
	response := new(GetCheckOutStatusResponse)
	err := service.client.Call("urn:thesecretserver.com/GetCheckOutStatus", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) ChangePassword(request *ChangePassword) (*ChangePasswordResponse, error) {
	response := new(ChangePasswordResponse)
	err := service.client.Call("urn:thesecretserver.com/ChangePassword", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) GetSecretsByFieldValue(request *GetSecretsByFieldValue) (*GetSecretsByFieldValueResponse, error) {
	response := new(GetSecretsByFieldValueResponse)
	err := service.client.Call("urn:thesecretserver.com/GetSecretsByFieldValue", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) SearchSecretsByFieldValue(request *SearchSecretsByFieldValue) (*SearchSecretsByFieldValueResponse, error) {
	response := new(SearchSecretsByFieldValueResponse)
	err := service.client.Call("urn:thesecretserver.com/SearchSecretsByFieldValue", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) GetSecretsByExposedFieldValue(request *GetSecretsByExposedFieldValue) (*GetSecretsByExposedFieldValueResponse, error) {
	response := new(GetSecretsByExposedFieldValueResponse)
	err := service.client.Call("urn:thesecretserver.com/GetSecretsByExposedFieldValue", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) SearchSecretsByExposedFieldValue(request *SearchSecretsByExposedFieldValue) (*SearchSecretsByExposedFieldValueResponse, error) {
	response := new(SearchSecretsByExposedFieldValueResponse)
	err := service.client.Call("urn:thesecretserver.com/SearchSecretsByExposedFieldValue", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) SearchSecretsByExposedValues(request *SearchSecretsByExposedValues) (*SearchSecretsByExposedValuesResponse, error) {
	response := new(SearchSecretsByExposedValuesResponse)
	err := service.client.Call("urn:thesecretserver.com/SearchSecretsByExposedValues", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) AddUser(request *AddUser) (*AddUserResponse, error) {
	response := new(AddUserResponse)
	err := service.client.Call("urn:thesecretserver.com/AddUser", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) SearchSecrets(request *SearchSecrets) (*SearchSecretsResponse, error) {
	response := new(SearchSecretsResponse)
	err := service.client.Call("urn:thesecretserver.com/SearchSecrets", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) SearchSecretsLegacy(request *SearchSecretsLegacy) (*SearchSecretsLegacyResponse, error) {
	response := new(SearchSecretsLegacyResponse)
	err := service.client.Call("urn:thesecretserver.com/SearchSecretsLegacy", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) SearchSecretsByFolder(request *SearchSecretsByFolder) (*SearchSecretsByFolderResponse, error) {
	response := new(SearchSecretsByFolderResponse)
	err := service.client.Call("urn:thesecretserver.com/SearchSecretsByFolder", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) SearchSecretsByFolderLegacy(request *SearchSecretsByFolderLegacy) (*SearchSecretsByFolderLegacyResponse, error) {
	response := new(SearchSecretsByFolderLegacyResponse)
	err := service.client.Call("urn:thesecretserver.com/SearchSecretsByFolderLegacy", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) GetFavorites(request *GetFavorites) (*GetFavoritesResponse, error) {
	response := new(GetFavoritesResponse)
	err := service.client.Call("urn:thesecretserver.com/GetFavorites", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) UpdateIsFavorite(request *UpdateIsFavorite) (*UpdateIsFavoriteResponse, error) {
	response := new(UpdateIsFavoriteResponse)
	err := service.client.Call("urn:thesecretserver.com/UpdateIsFavorite", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) AddSecret(request *AddSecret) (*AddSecretResponse, error) {
	response := new(AddSecretResponse)
	err := service.client.Call("urn:thesecretserver.com/AddSecret", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) AddNewSecret(request *AddNewSecret) (*AddNewSecretResponse, error) {
	response := new(AddNewSecretResponse)
	err := service.client.Call("urn:thesecretserver.com/AddNewSecret", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) GetNewSecret(request *GetNewSecret) (*GetNewSecretResponse, error) {
	response := new(GetNewSecretResponse)
	err := service.client.Call("urn:thesecretserver.com/GetNewSecret", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) GetSecretTemplateFields(request *GetSecretTemplateFields) (*GetSecretTemplateFieldsResponse, error) {
	response := new(GetSecretTemplateFieldsResponse)
	err := service.client.Call("urn:thesecretserver.com/GetSecretTemplateFields", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) UpdateSecret(request *UpdateSecret) (*UpdateSecretResponse, error) {
	response := new(UpdateSecretResponse)
	err := service.client.Call("urn:thesecretserver.com/UpdateSecret", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) GetSecretTemplates(request *GetSecretTemplates) (*GetSecretTemplatesResponse, error) {
	response := new(GetSecretTemplatesResponse)
	err := service.client.Call("urn:thesecretserver.com/GetSecretTemplates", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) GeneratePassword(request *GeneratePassword) (*GeneratePasswordResponse, error) {
	response := new(GeneratePasswordResponse)
	err := service.client.Call("urn:thesecretserver.com/GeneratePassword", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) DeactivateSecret(request *DeactivateSecret) (*DeactivateSecretResponse, error) {
	response := new(DeactivateSecretResponse)
	err := service.client.Call("urn:thesecretserver.com/DeactivateSecret", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) VersionGet(request *VersionGet) (*VersionGetResponse, error) {
	response := new(VersionGetResponse)
	err := service.client.Call("urn:thesecretserver.com/VersionGet", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) FolderGet(request *FolderGet) (*FolderGetResponse, error) {
	response := new(FolderGetResponse)
	err := service.client.Call("urn:thesecretserver.com/FolderGet", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) FolderUpdate(request *FolderUpdate) (*FolderUpdateResponse, error) {
	response := new(FolderUpdateResponse)
	err := service.client.Call("urn:thesecretserver.com/FolderUpdate", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) FolderGetAllChildren(request *FolderGetAllChildren) (*FolderGetAllChildrenResponse, error) {
	response := new(FolderGetAllChildrenResponse)
	err := service.client.Call("urn:thesecretserver.com/FolderGetAllChildren", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) FolderCreate(request *FolderCreate) (*FolderCreateResponse, error) {
	response := new(FolderCreateResponse)
	err := service.client.Call("urn:thesecretserver.com/FolderCreate", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) FolderExtendedCreate(request *FolderExtendedCreate) (*FolderExtendedCreateResponse, error) {
	response := new(FolderExtendedCreateResponse)
	err := service.client.Call("urn:thesecretserver.com/FolderExtendedCreate", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) FolderExtendedGet(request *FolderExtendedGet) (*FolderExtendedGetResponse, error) {
	response := new(FolderExtendedGetResponse)
	err := service.client.Call("urn:thesecretserver.com/FolderExtendedGet", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) FolderExtendedUpdate(request *FolderExtendedUpdate) (*FolderExtendedUpdateResponse, error) {
	response := new(FolderExtendedUpdateResponse)
	err := service.client.Call("urn:thesecretserver.com/FolderExtendedUpdate", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) FolderExtendedGetNew(request *FolderExtendedGetNew) (*FolderExtendedGetNewResponse, error) {
	response := new(FolderExtendedGetNewResponse)
	err := service.client.Call("urn:thesecretserver.com/FolderExtendedGetNew", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) SearchFolders(request *SearchFolders) (*SearchFoldersResponse, error) {
	response := new(SearchFoldersResponse)
	err := service.client.Call("urn:thesecretserver.com/SearchFolders", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) DownloadFileAttachment(request *DownloadFileAttachment) (*DownloadFileAttachmentResponse, error) {
	response := new(DownloadFileAttachmentResponse)
	err := service.client.Call("urn:thesecretserver.com/DownloadFileAttachment", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) DownloadFileAttachmentByItemId(request *DownloadFileAttachmentByItemId) (*DownloadFileAttachmentByItemIdResponse, error) {
	response := new(DownloadFileAttachmentByItemIdResponse)
	err := service.client.Call("urn:thesecretserver.com/DownloadFileAttachmentByItemId", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) UploadFileAttachment(request *UploadFileAttachment) (*UploadFileAttachmentResponse, error) {
	response := new(UploadFileAttachmentResponse)
	err := service.client.Call("urn:thesecretserver.com/UploadFileAttachment", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) UploadFileAttachmentByItemId(request *UploadFileAttachmentByItemId) (*UploadFileAttachmentByItemIdResponse, error) {
	response := new(UploadFileAttachmentByItemIdResponse)
	err := service.client.Call("urn:thesecretserver.com/UploadFileAttachmentByItemId", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) ExpireSecret(request *ExpireSecret) (*ExpireSecretResponse, error) {
	response := new(ExpireSecretResponse)
	err := service.client.Call("urn:thesecretserver.com/ExpireSecret", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) SetCheckOutEnabled(request *SetCheckOutEnabled) (*SetCheckOutEnabledResponse, error) {
	response := new(SetCheckOutEnabledResponse)
	err := service.client.Call("urn:thesecretserver.com/SetCheckOutEnabled", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) ImportXML(request *ImportXML) (*ImportXMLResponse, error) {
	response := new(ImportXMLResponse)
	err := service.client.Call("urn:thesecretserver.com/ImportXML", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) GetSecretAudit(request *GetSecretAudit) (*GetSecretAuditResponse, error) {
	response := new(GetSecretAuditResponse)
	err := service.client.Call("urn:thesecretserver.com/GetSecretAudit", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) AddDependency(request *AddDependency) (*AddDependencyResponse, error) {
	response := new(AddDependencyResponse)
	err := service.client.Call("urn:thesecretserver.com/AddDependency", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) RemoveDependency(request *RemoveDependency) (*RemoveDependencyResponse, error) {
	response := new(RemoveDependencyResponse)
	err := service.client.Call("urn:thesecretserver.com/RemoveDependency", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) GetDependencies(request *GetDependencies) (*GetDependenciesResponse, error) {
	response := new(GetDependenciesResponse)
	err := service.client.Call("urn:thesecretserver.com/GetDependencies", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) GetDistributedEngines(request *GetDistributedEngines) (*GetDistributedEnginesResponse, error) {
	response := new(GetDistributedEnginesResponse)
	err := service.client.Call("urn:thesecretserver.com/GetDistributedEngines", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) GetTicketSystems(request *GetTicketSystems) (*GetTicketSystemsResponse, error) {
	response := new(GetTicketSystemsResponse)
	err := service.client.Call("urn:thesecretserver.com/GetTicketSystems", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) AssignSite(request *AssignSite) (*AssignSiteResponse, error) {
	response := new(AssignSiteResponse)
	err := service.client.Call("urn:thesecretserver.com/AssignSite", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) CheckIn(request *CheckIn) (*CheckInResponse, error) {
	response := new(CheckInResponse)
	err := service.client.Call("urn:thesecretserver.com/CheckIn", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) AddSecretCustomAudit(request *AddSecretCustomAudit) (*AddSecretCustomAuditResponse, error) {
	response := new(AddSecretCustomAuditResponse)
	err := service.client.Call("urn:thesecretserver.com/AddSecretCustomAudit", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) UpdateSecretPermission(request *UpdateSecretPermission) (*UpdateSecretPermissionResponse, error) {
	response := new(UpdateSecretPermissionResponse)
	err := service.client.Call("urn:thesecretserver.com/UpdateSecretPermission", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) CheckInByKey(request *CheckInByKey) (*CheckInByKeyResponse, error) {
	response := new(CheckInByKeyResponse)
	err := service.client.Call("urn:thesecretserver.com/CheckInByKey", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) WhoAmI(request *WhoAmI) (*WhoAmIResponse, error) {
	response := new(WhoAmIResponse)
	err := service.client.Call("urn:thesecretserver.com/WhoAmI", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) GetAllGroups(request *GetAllGroups) (*GetAllGroupsResponse, error) {
	response := new(GetAllGroupsResponse)
	err := service.client.Call("urn:thesecretserver.com/GetAllGroups", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) AssignUserToGroup(request *AssignUserToGroup) (*AssignUserToGroupResponse, error) {
	response := new(AssignUserToGroupResponse)
	err := service.client.Call("urn:thesecretserver.com/AssignUserToGroup", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) GetSSHLoginCredentials(request *GetSSHLoginCredentials) (*GetSSHLoginCredentialsResponse, error) {
	response := new(GetSSHLoginCredentialsResponse)
	err := service.client.Call("urn:thesecretserver.com/GetSSHLoginCredentials", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) GetSSHLoginCredentialsWithMachine(request *GetSSHLoginCredentialsWithMachine) (*GetSSHLoginCredentialsWithMachineResponse, error) {
	response := new(GetSSHLoginCredentialsWithMachineResponse)
	err := service.client.Call("urn:thesecretserver.com/GetSSHLoginCredentialsWithMachine", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) SearchUsers(request *SearchUsers) (*SearchUsersResponse, error) {
	response := new(SearchUsersResponse)
	err := service.client.Call("urn:thesecretserver.com/SearchUsers", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) GetUser(request *GetUser) (*GetUserResponse, error) {
	response := new(GetUserResponse)
	err := service.client.Call("urn:thesecretserver.com/GetUser", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) UpdateUser(request *UpdateUser) (*UpdateUserResponse, error) {
	response := new(UpdateUserResponse)
	err := service.client.Call("urn:thesecretserver.com/UpdateUser", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) GetSecretItemHistoryByFieldName(request *GetSecretItemHistoryByFieldName) (*GetSecretItemHistoryByFieldNameResponse, error) {
	response := new(GetSecretItemHistoryByFieldNameResponse)
	err := service.client.Call("urn:thesecretserver.com/GetSecretItemHistoryByFieldName", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) GetSecretPolicyForSecret(request *GetSecretPolicyForSecret) (*GetSecretPolicyForSecretResponse, error) {
	response := new(GetSecretPolicyForSecretResponse)
	err := service.client.Call("urn:thesecretserver.com/GetSecretPolicyForSecret", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) AssignSecretPolicyForSecret(request *AssignSecretPolicyForSecret) (*AssignSecretPolicyForSecretResponse, error) {
	response := new(AssignSecretPolicyForSecretResponse)
	err := service.client.Call("urn:thesecretserver.com/AssignSecretPolicyForSecret", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) SearchSecretPolicies(request *SearchSecretPolicies) (*SearchSecretPoliciesResponse, error) {
	response := new(SearchSecretPoliciesResponse)
	err := service.client.Call("urn:thesecretserver.com/SearchSecretPolicies", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) RunActiveDirectorySynchronization(request *RunActiveDirectorySynchronization) (*RunActiveDirectorySynchronizationResponse, error) {
	response := new(RunActiveDirectorySynchronizationResponse)
	err := service.client.Call("urn:thesecretserver.com/RunActiveDirectorySynchronization", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) AddGroupToActiveDirectorySynchronization(request *AddGroupToActiveDirectorySynchronization) (*AddGroupToActiveDirectorySynchronizationResponse, error) {
	response := new(AddGroupToActiveDirectorySynchronizationResponse)
	err := service.client.Call("urn:thesecretserver.com/AddGroupToActiveDirectorySynchronization", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) AddSecretPolicy(request *AddSecretPolicy) (*AddSecretPolicyResponse, error) {
	response := new(AddSecretPolicyResponse)
	err := service.client.Call("urn:thesecretserver.com/AddSecretPolicy", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) GetNewSecretPolicy(request *GetNewSecretPolicy) (*GetNewSecretPolicyResponse, error) {
	response := new(GetNewSecretPolicyResponse)
	err := service.client.Call("urn:thesecretserver.com/GetNewSecretPolicy", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) GetSSHCommandMenu(request *GetSSHCommandMenu) (*GetSSHCommandMenuResponse, error) {
	response := new(GetSSHCommandMenuResponse)
	err := service.client.Call("urn:thesecretserver.com/GetSSHCommandMenu", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) SaveSSHCommandMenu(request *SaveSSHCommandMenu) (*SaveSSHCommandMenuResponse, error) {
	response := new(SaveSSHCommandMenuResponse)
	err := service.client.Call("urn:thesecretserver.com/SaveSSHCommandMenu", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) GetAllSSHCommandMenus(request *GetAllSSHCommandMenus) (*GetAllSSHCommandMenusResponse, error) {
	response := new(GetAllSSHCommandMenusResponse)
	err := service.client.Call("urn:thesecretserver.com/GetAllSSHCommandMenus", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) DeleteSSHCommandMenu(request *DeleteSSHCommandMenu) (*DeleteSSHCommandMenuResponse, error) {
	response := new(DeleteSSHCommandMenuResponse)
	err := service.client.Call("urn:thesecretserver.com/DeleteSSHCommandMenu", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) RestoreSSHCommandMenu(request *RestoreSSHCommandMenu) (*RestoreSSHCommandMenuResponse, error) {
	response := new(RestoreSSHCommandMenuResponse)
	err := service.client.Call("urn:thesecretserver.com/RestoreSSHCommandMenu", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) AddScript(request *AddScript) (*AddScriptResponse, error) {
	response := new(AddScriptResponse)
	err := service.client.Call("urn:thesecretserver.com/AddScript", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) GetAllScripts(request *GetAllScripts) (*GetAllScriptsResponse, error) {
	response := new(GetAllScriptsResponse)
	err := service.client.Call("urn:thesecretserver.com/GetAllScripts", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) GetScript(request *GetScript) (*GetScriptResponse, error) {
	response := new(GetScriptResponse)
	err := service.client.Call("urn:thesecretserver.com/GetScript", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceSoap) UpdateScript(request *UpdateScript) (*UpdateScriptResponse, error) {
	response := new(UpdateScriptResponse)
	err := service.client.Call("urn:thesecretserver.com/UpdateScript", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

type SSWebServiceHttpGet struct {
	client *SOAPClient
}

func NewSSWebServiceHttpGet(url string, tls bool, auth *BasicAuth) *SSWebServiceHttpGet {
	if url == "" {
		url = ""
	}
	client := NewSOAPClient(url, tls, auth)

	return &SSWebServiceHttpGet{
		client: client,
	}
}

func NewSSWebServiceHttpGetWithTLSConfig(url string, tlsCfg *tls.Config, auth *BasicAuth) *SSWebServiceHttpGet {
	if url == "" {
		url = ""
	}
	client := NewSOAPClientWithTLSConfig(url, tlsCfg, auth)

	return &SSWebServiceHttpGet{
		client: client,
	}
}

func (service *SSWebServiceHttpGet) AddHeader(header interface{}) {
	service.client.AddHeader(header)
}

// Backwards-compatible function: use AddHeader instead
func (service *SSWebServiceHttpGet) SetHeader(header interface{}) {
	service.client.AddHeader(header)
}

func (service *SSWebServiceHttpGet) ApproveSecretAccessRequest(request *String) (*RequestApprovalResult, error) {
	response := new(RequestApprovalResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpGet) DenySecretAccessRequest(request *String) (*RequestApprovalResult, error) {
	response := new(RequestApprovalResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpGet) Authenticate(request *String) (*AuthenticateResult, error) {
	response := new(AuthenticateResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpGet) ImpersonateUser(request *String) (*ImpersonateResult, error) {
	response := new(ImpersonateResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpGet) AuthenticateRADIUS(request *String) (*AuthenticateResult, error) {
	response := new(AuthenticateResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpGet) GetTokenIsValid(request *String) (*TokenIsValidResult, error) {
	response := new(TokenIsValidResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpGet) GetSecretLegacy(request *String) (*GetSecretResult, error) {
	response := new(GetSecretResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpGet) GetCheckOutStatus(request *String) (*GetCheckOutStatusResult, error) {
	response := new(GetCheckOutStatusResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpGet) ChangePassword(request *String) (*WebServiceResult, error) {
	response := new(WebServiceResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpGet) GetSecretsByFieldValue(request *String) (*GetSecretsByFieldValueResult, error) {
	response := new(GetSecretsByFieldValueResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpGet) SearchSecretsByFieldValue(request *String) (*SearchSecretsResult, error) {
	response := new(SearchSecretsResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpGet) GetSecretsByExposedFieldValue(request *String) (*GetSecretsByFieldValueResult, error) {
	response := new(GetSecretsByFieldValueResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpGet) SearchSecretsByExposedFieldValue(request *String) (*SearchSecretsResult, error) {
	response := new(SearchSecretsResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpGet) SearchSecretsByExposedValues(request *String) (*SearchSecretsResult, error) {
	response := new(SearchSecretsResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpGet) SearchSecretsLegacy(request *String) (*SearchSecretsResult, error) {
	response := new(SearchSecretsResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpGet) GetFavorites(request *String) (*GetFavoritesResult, error) {
	response := new(GetFavoritesResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpGet) UpdateIsFavorite(request *String) (*WebServiceResult, error) {
	response := new(WebServiceResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpGet) AddSecret(request *String) (*AddSecretResult, error) {
	response := new(AddSecretResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpGet) GetNewSecret(request *String) (*GetSecretResult, error) {
	response := new(GetSecretResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpGet) GetSecretTemplateFields(request *String) (*GetSecretTemplateFieldsResult, error) {
	response := new(GetSecretTemplateFieldsResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpGet) GetSecretTemplates(request *String) (*GetSecretTemplatesResult, error) {
	response := new(GetSecretTemplatesResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpGet) GeneratePassword(request *String) (*GeneratePasswordResult, error) {
	response := new(GeneratePasswordResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpGet) DeactivateSecret(request *String) (*WebServiceResult, error) {
	response := new(WebServiceResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpGet) VersionGet() (*VersionGetResult, error) {
	response := new(VersionGetResult)
	err := service.client.Call("", nil, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpGet) FolderGet(request *String) (*GetFolderResult, error) {
	response := new(GetFolderResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpGet) FolderGetAllChildren(request *String) (*GetFoldersResult, error) {
	response := new(GetFoldersResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpGet) FolderCreate(request *String) (*CreateFolderResult, error) {
	response := new(CreateFolderResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpGet) FolderExtendedGet(request *String) (*FolderExtendedGetResult, error) {
	response := new(FolderExtendedGetResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpGet) SearchFolders(request *String) (*SearchFolderResult, error) {
	response := new(SearchFolderResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpGet) DownloadFileAttachment(request *String) (*FileDownloadResult, error) {
	response := new(FileDownloadResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpGet) DownloadFileAttachmentByItemId(request *String) (*FileDownloadResult, error) {
	response := new(FileDownloadResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpGet) UploadFileAttachment(request *String) (*WebServiceResult, error) {
	response := new(WebServiceResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpGet) UploadFileAttachmentByItemId(request *String) (*WebServiceResult, error) {
	response := new(WebServiceResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpGet) ExpireSecret(request *String) (*WebServiceResult, error) {
	response := new(WebServiceResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpGet) ImportXML(request *String) (*WebServiceResult, error) {
	response := new(WebServiceResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpGet) GetSecretAudit(request *String) (*GetSecretAuditResult, error) {
	response := new(GetSecretAuditResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpGet) RemoveDependency(request *String) (*WebServiceResult, error) {
	response := new(WebServiceResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpGet) GetDependencies(request *String) (*GetDependenciesResult, error) {
	response := new(GetDependenciesResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpGet) GetDistributedEngines(request *String) (*GetSitesResult, error) {
	response := new(GetSitesResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpGet) GetTicketSystems(request *String) (*GetTicketSystemsResult, error) {
	response := new(GetTicketSystemsResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpGet) AssignSite(request *String) (*WebServiceResult, error) {
	response := new(WebServiceResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpGet) CheckIn(request *String) (*WebServiceResult, error) {
	response := new(WebServiceResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpGet) CheckInByKey(request *String) (*WebServiceResult, error) {
	response := new(WebServiceResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpGet) WhoAmI(request *String) (*UserInfoResult, error) {
	response := new(UserInfoResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpGet) GetAllGroups(request *String) (*GetAllGroupsResult, error) {
	response := new(GetAllGroupsResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpGet) AssignUserToGroup(request *String) (*WebServiceResult, error) {
	response := new(WebServiceResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpGet) GetSSHLoginCredentials(request *String) (*SSHCredentialsResult, error) {
	response := new(SSHCredentialsResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpGet) GetSSHLoginCredentialsWithMachine(request *String) (*SSHCredentialsResult, error) {
	response := new(SSHCredentialsResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpGet) SearchUsers(request *String) (*GetUsersResult, error) {
	response := new(GetUsersResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpGet) GetUser(request *String) (*GetUserResult, error) {
	response := new(GetUserResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpGet) GetSecretItemHistoryByFieldName(request *String) (*SecretItemHistoryResult, error) {
	response := new(SecretItemHistoryResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpGet) GetSecretPolicyForSecret(request *String) (*SecretPolicyForSecretResult, error) {
	response := new(SecretPolicyForSecretResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpGet) SearchSecretPolicies(request *String) (*SearchSecretPoliciesResult, error) {
	response := new(SearchSecretPoliciesResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpGet) RunActiveDirectorySynchronization(request *String) (*WebServiceResult, error) {
	response := new(WebServiceResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpGet) GetNewSecretPolicy(request *String) (*SecretPolicyResult, error) {
	response := new(SecretPolicyResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpGet) GetSSHCommandMenu(request *String) (*GetSshCommandMenuResult, error) {
	response := new(GetSshCommandMenuResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpGet) DeleteSSHCommandMenu(request *String) (*WebServiceResult, error) {
	response := new(WebServiceResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpGet) RestoreSSHCommandMenu(request *String) (*WebServiceResult, error) {
	response := new(WebServiceResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpGet) GetAllScripts(request *String) (*GetUserScriptsResult, error) {
	response := new(GetUserScriptsResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpGet) GetScript(request *String) (*GetUserScriptResult, error) {
	response := new(GetUserScriptResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

type SSWebServiceHttpPost struct {
	client *SOAPClient
}

func NewSSWebServiceHttpPost(url string, tls bool, auth *BasicAuth) *SSWebServiceHttpPost {
	if url == "" {
		url = ""
	}
	client := NewSOAPClient(url, tls, auth)

	return &SSWebServiceHttpPost{
		client: client,
	}
}

func NewSSWebServiceHttpPostWithTLSConfig(url string, tlsCfg *tls.Config, auth *BasicAuth) *SSWebServiceHttpPost {
	if url == "" {
		url = ""
	}
	client := NewSOAPClientWithTLSConfig(url, tlsCfg, auth)

	return &SSWebServiceHttpPost{
		client: client,
	}
}

func (service *SSWebServiceHttpPost) AddHeader(header interface{}) {
	service.client.AddHeader(header)
}

// Backwards-compatible function: use AddHeader instead
func (service *SSWebServiceHttpPost) SetHeader(header interface{}) {
	service.client.AddHeader(header)
}

func (service *SSWebServiceHttpPost) ApproveSecretAccessRequest(request *String) (*RequestApprovalResult, error) {
	response := new(RequestApprovalResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpPost) DenySecretAccessRequest(request *String) (*RequestApprovalResult, error) {
	response := new(RequestApprovalResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpPost) Authenticate(request *String) (*AuthenticateResult, error) {
	response := new(AuthenticateResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpPost) ImpersonateUser(request *String) (*ImpersonateResult, error) {
	response := new(ImpersonateResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpPost) AuthenticateRADIUS(request *String) (*AuthenticateResult, error) {
	response := new(AuthenticateResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpPost) GetTokenIsValid(request *String) (*TokenIsValidResult, error) {
	response := new(TokenIsValidResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpPost) GetSecretLegacy(request *String) (*GetSecretResult, error) {
	response := new(GetSecretResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpPost) GetCheckOutStatus(request *String) (*GetCheckOutStatusResult, error) {
	response := new(GetCheckOutStatusResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpPost) ChangePassword(request *String) (*WebServiceResult, error) {
	response := new(WebServiceResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpPost) GetSecretsByFieldValue(request *String) (*GetSecretsByFieldValueResult, error) {
	response := new(GetSecretsByFieldValueResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpPost) SearchSecretsByFieldValue(request *String) (*SearchSecretsResult, error) {
	response := new(SearchSecretsResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpPost) GetSecretsByExposedFieldValue(request *String) (*GetSecretsByFieldValueResult, error) {
	response := new(GetSecretsByFieldValueResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpPost) SearchSecretsByExposedFieldValue(request *String) (*SearchSecretsResult, error) {
	response := new(SearchSecretsResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpPost) SearchSecretsByExposedValues(request *String) (*SearchSecretsResult, error) {
	response := new(SearchSecretsResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpPost) SearchSecretsLegacy(request *String) (*SearchSecretsResult, error) {
	response := new(SearchSecretsResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpPost) GetFavorites(request *String) (*GetFavoritesResult, error) {
	response := new(GetFavoritesResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpPost) UpdateIsFavorite(request *String) (*WebServiceResult, error) {
	response := new(WebServiceResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpPost) AddSecret(request *String) (*AddSecretResult, error) {
	response := new(AddSecretResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpPost) GetNewSecret(request *String) (*GetSecretResult, error) {
	response := new(GetSecretResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpPost) GetSecretTemplateFields(request *String) (*GetSecretTemplateFieldsResult, error) {
	response := new(GetSecretTemplateFieldsResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpPost) GetSecretTemplates(request *String) (*GetSecretTemplatesResult, error) {
	response := new(GetSecretTemplatesResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpPost) GeneratePassword(request *String) (*GeneratePasswordResult, error) {
	response := new(GeneratePasswordResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpPost) DeactivateSecret(request *String) (*WebServiceResult, error) {
	response := new(WebServiceResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpPost) VersionGet() (*VersionGetResult, error) {
	response := new(VersionGetResult)
	err := service.client.Call("", nil, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpPost) FolderGet(request *String) (*GetFolderResult, error) {
	response := new(GetFolderResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpPost) FolderGetAllChildren(request *String) (*GetFoldersResult, error) {
	response := new(GetFoldersResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpPost) FolderCreate(request *String) (*CreateFolderResult, error) {
	response := new(CreateFolderResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpPost) FolderExtendedGet(request *String) (*FolderExtendedGetResult, error) {
	response := new(FolderExtendedGetResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpPost) SearchFolders(request *String) (*SearchFolderResult, error) {
	response := new(SearchFolderResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpPost) DownloadFileAttachment(request *String) (*FileDownloadResult, error) {
	response := new(FileDownloadResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpPost) DownloadFileAttachmentByItemId(request *String) (*FileDownloadResult, error) {
	response := new(FileDownloadResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpPost) UploadFileAttachment(request *String) (*WebServiceResult, error) {
	response := new(WebServiceResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpPost) UploadFileAttachmentByItemId(request *String) (*WebServiceResult, error) {
	response := new(WebServiceResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpPost) ExpireSecret(request *String) (*WebServiceResult, error) {
	response := new(WebServiceResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpPost) ImportXML(request *String) (*WebServiceResult, error) {
	response := new(WebServiceResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpPost) GetSecretAudit(request *String) (*GetSecretAuditResult, error) {
	response := new(GetSecretAuditResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpPost) RemoveDependency(request *String) (*WebServiceResult, error) {
	response := new(WebServiceResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpPost) GetDependencies(request *String) (*GetDependenciesResult, error) {
	response := new(GetDependenciesResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpPost) GetDistributedEngines(request *String) (*GetSitesResult, error) {
	response := new(GetSitesResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpPost) GetTicketSystems(request *String) (*GetTicketSystemsResult, error) {
	response := new(GetTicketSystemsResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpPost) AssignSite(request *String) (*WebServiceResult, error) {
	response := new(WebServiceResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpPost) CheckIn(request *String) (*WebServiceResult, error) {
	response := new(WebServiceResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpPost) CheckInByKey(request *String) (*WebServiceResult, error) {
	response := new(WebServiceResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpPost) WhoAmI(request *String) (*UserInfoResult, error) {
	response := new(UserInfoResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpPost) GetAllGroups(request *String) (*GetAllGroupsResult, error) {
	response := new(GetAllGroupsResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpPost) AssignUserToGroup(request *String) (*WebServiceResult, error) {
	response := new(WebServiceResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpPost) GetSSHLoginCredentials(request *String) (*SSHCredentialsResult, error) {
	response := new(SSHCredentialsResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpPost) GetSSHLoginCredentialsWithMachine(request *String) (*SSHCredentialsResult, error) {
	response := new(SSHCredentialsResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpPost) SearchUsers(request *String) (*GetUsersResult, error) {
	response := new(GetUsersResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpPost) GetUser(request *String) (*GetUserResult, error) {
	response := new(GetUserResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpPost) GetSecretItemHistoryByFieldName(request *String) (*SecretItemHistoryResult, error) {
	response := new(SecretItemHistoryResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpPost) GetSecretPolicyForSecret(request *String) (*SecretPolicyForSecretResult, error) {
	response := new(SecretPolicyForSecretResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpPost) SearchSecretPolicies(request *String) (*SearchSecretPoliciesResult, error) {
	response := new(SearchSecretPoliciesResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpPost) RunActiveDirectorySynchronization(request *String) (*WebServiceResult, error) {
	response := new(WebServiceResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpPost) GetNewSecretPolicy(request *String) (*SecretPolicyResult, error) {
	response := new(SecretPolicyResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpPost) GetSSHCommandMenu(request *String) (*GetSshCommandMenuResult, error) {
	response := new(GetSshCommandMenuResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpPost) DeleteSSHCommandMenu(request *String) (*WebServiceResult, error) {
	response := new(WebServiceResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpPost) RestoreSSHCommandMenu(request *String) (*WebServiceResult, error) {
	response := new(WebServiceResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpPost) GetAllScripts(request *String) (*GetUserScriptsResult, error) {
	response := new(GetUserScriptsResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SSWebServiceHttpPost) GetScript(request *String) (*GetUserScriptResult, error) {
	response := new(GetUserScriptResult)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

var timeout = time.Duration(30 * time.Second)

func dialTimeout(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, timeout)
}

type SOAPEnvelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Header  *SOAPHeader
	Body    SOAPBody
}

type SOAPHeader struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header"`

	Items []interface{} `xml:",omitempty"`
}

type SOAPBody struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`

	Fault   *SOAPFault  `xml:",omitempty"`
	Content interface{} `xml:",omitempty"`
}

type SOAPFault struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault"`

	Code   string `xml:"faultcode,omitempty"`
	String string `xml:"faultstring,omitempty"`
	Actor  string `xml:"faultactor,omitempty"`
	Detail string `xml:"detail,omitempty"`
}

const (
	// Predefined WSS namespaces to be used in
	WssNsWSSE string = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd"
	WssNsWSU  string = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd"
	WssNsType string = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-username-token-profile-1.0#PasswordText"
)

type WSSSecurityHeader struct {
	XMLName   xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ wsse:Security"`
	XmlNSWsse string   `xml:"xmlns:wsse,attr"`

	MustUnderstand string `xml:"mustUnderstand,attr,omitempty"`

	Token *WSSUsernameToken `xml:",omitempty"`
}

type WSSUsernameToken struct {
	XMLName   xml.Name `xml:"wsse:UsernameToken"`
	XmlNSWsu  string   `xml:"xmlns:wsu,attr"`
	XmlNSWsse string   `xml:"xmlns:wsse,attr"`

	Id string `xml:"wsu:Id,attr,omitempty"`

	Username *WSSUsername `xml:",omitempty"`
	Password *WSSPassword `xml:",omitempty"`
}

type WSSUsername struct {
	XMLName   xml.Name `xml:"wsse:Username"`
	XmlNSWsse string   `xml:"xmlns:wsse,attr"`

	Data string `xml:",chardata"`
}

type WSSPassword struct {
	XMLName   xml.Name `xml:"wsse:Password"`
	XmlNSWsse string   `xml:"xmlns:wsse,attr"`
	XmlNSType string   `xml:"Type,attr"`

	Data string `xml:",chardata"`
}

type BasicAuth struct {
	Login    string
	Password string
}

type SOAPClient struct {
	url     string
	tlsCfg  *tls.Config
	auth    *BasicAuth
	headers []interface{}
}

// **********
// Accepted solution from http://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang
// Author: Icza - http://stackoverflow.com/users/1705598/icza

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func randStringBytesMaskImprSrc(n int) string {
	src := rand.NewSource(time.Now().UnixNano())
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}

// **********

func NewWSSSecurityHeader(user, pass, mustUnderstand string) *WSSSecurityHeader {
	hdr := &WSSSecurityHeader{XmlNSWsse: WssNsWSSE, MustUnderstand: mustUnderstand}
	hdr.Token = &WSSUsernameToken{XmlNSWsu: WssNsWSU, XmlNSWsse: WssNsWSSE, Id: "UsernameToken-" + randStringBytesMaskImprSrc(9)}
	hdr.Token.Username = &WSSUsername{XmlNSWsse: WssNsWSSE, Data: user}
	hdr.Token.Password = &WSSPassword{XmlNSWsse: WssNsWSSE, XmlNSType: WssNsType, Data: pass}
	return hdr
}

func (b *SOAPBody) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if b.Content == nil {
		return xml.UnmarshalError("Content must be a pointer to a struct")
	}

	var (
		token    xml.Token
		err      error
		consumed bool
	)

Loop:
	for {
		if token, err = d.Token(); err != nil {
			return err
		}

		if token == nil {
			break
		}

		switch se := token.(type) {
		case xml.StartElement:
			if consumed {
				return xml.UnmarshalError("Found multiple elements inside SOAP body; not wrapped-document/literal WS-I compliant")
			} else if se.Name.Space == "http://schemas.xmlsoap.org/soap/envelope/" && se.Name.Local == "Fault" {
				b.Fault = &SOAPFault{}
				b.Content = nil

				err = d.DecodeElement(b.Fault, &se)
				if err != nil {
					return err
				}

				consumed = true
			} else {
				if err = d.DecodeElement(b.Content, &se); err != nil {
					return err
				}

				consumed = true
			}
		case xml.EndElement:
			break Loop
		}
	}

	return nil
}

func (f *SOAPFault) Error() string {
	return f.String
}

func NewSOAPClient(url string, insecureSkipVerify bool, auth *BasicAuth) *SOAPClient {
	tlsCfg := &tls.Config{
		InsecureSkipVerify: insecureSkipVerify,
	}
	return NewSOAPClientWithTLSConfig(url, tlsCfg, auth)
}

func NewSOAPClientWithTLSConfig(url string, tlsCfg *tls.Config, auth *BasicAuth) *SOAPClient {
	return &SOAPClient{
		url:    url,
		tlsCfg: tlsCfg,
		auth:   auth,
	}
}

func (s *SOAPClient) AddHeader(header interface{}) {
	s.headers = append(s.headers, header)
}

func (s *SOAPClient) Call(soapAction string, request, response interface{}) error {
	envelope := SOAPEnvelope{}

	if s.headers != nil && len(s.headers) > 0 {
		soapHeader := &SOAPHeader{Items: make([]interface{}, len(s.headers))}
		copy(soapHeader.Items, s.headers)
		envelope.Header = soapHeader
	}

	envelope.Body.Content = request
	buffer := new(bytes.Buffer)

	encoder := xml.NewEncoder(buffer)
	//encoder.Indent("  ", "    ")

	if err := encoder.Encode(envelope); err != nil {
		return err
	}

	if err := encoder.Flush(); err != nil {
		return err
	}

	glog.V(3).Info(buffer.String())

	req, err := http.NewRequest("POST", s.url, buffer)
	if err != nil {
		return err
	}
	if s.auth != nil {
		req.SetBasicAuth(s.auth.Login, s.auth.Password)
	}

	req.Header.Add("Content-Type", "text/xml; charset=\"utf-8\"")
	req.Header.Add("SOAPAction", soapAction)

	req.Header.Set("User-Agent", "gowsdl/0.1")
	req.Close = true

	tr := &http.Transport{
		TLSClientConfig: s.tlsCfg,
		Dial:            dialTimeout,
	}

	client := &http.Client{Transport: tr}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	rawbody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if len(rawbody) == 0 {
		glog.V(3).Info("empty response")
		return nil
	}

	glog.V(3).Info(string(rawbody))
	respEnvelope := new(SOAPEnvelope)
	respEnvelope.Body = SOAPBody{Content: response}
	err = xml.Unmarshal(rawbody, respEnvelope)
	if err != nil {
		return err
	}

	fault := respEnvelope.Body.Fault
	if fault != nil {
		return fault
	}

	return nil
}
