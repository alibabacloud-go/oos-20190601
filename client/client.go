// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CancelExecutionRequest struct {
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ExecutionId *string `json:"ExecutionId,omitempty" xml:"ExecutionId,omitempty"`
}

func (s CancelExecutionRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelExecutionRequest) GoString() string {
	return s.String()
}

func (s *CancelExecutionRequest) SetRegionId(v string) *CancelExecutionRequest {
	s.RegionId = &v
	return s
}

func (s *CancelExecutionRequest) SetExecutionId(v string) *CancelExecutionRequest {
	s.ExecutionId = &v
	return s
}

type CancelExecutionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelExecutionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelExecutionResponseBody) GoString() string {
	return s.String()
}

func (s *CancelExecutionResponseBody) SetRequestId(v string) *CancelExecutionResponseBody {
	s.RequestId = &v
	return s
}

type CancelExecutionResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CancelExecutionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CancelExecutionResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelExecutionResponse) GoString() string {
	return s.String()
}

func (s *CancelExecutionResponse) SetHeaders(v map[string]*string) *CancelExecutionResponse {
	s.Headers = v
	return s
}

func (s *CancelExecutionResponse) SetBody(v *CancelExecutionResponseBody) *CancelExecutionResponse {
	s.Body = v
	return s
}

type ChangeResourceGroupRequest struct {
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId         *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	NewResourceGroupId *string `json:"NewResourceGroupId,omitempty" xml:"NewResourceGroupId,omitempty"`
	ResourceType       *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ChangeResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupRequest) SetRegionId(v string) *ChangeResourceGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetResourceId(v string) *ChangeResourceGroupRequest {
	s.ResourceId = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetNewResourceGroupId(v string) *ChangeResourceGroupRequest {
	s.NewResourceGroupId = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetResourceType(v string) *ChangeResourceGroupRequest {
	s.ResourceType = &v
	return s
}

type ChangeResourceGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ChangeResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChangeResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupResponseBody) SetRequestId(v string) *ChangeResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

type ChangeResourceGroupResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ChangeResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ChangeResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ChangeResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupResponse) SetHeaders(v map[string]*string) *ChangeResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *ChangeResourceGroupResponse) SetBody(v *ChangeResourceGroupResponseBody) *ChangeResourceGroupResponse {
	s.Body = v
	return s
}

type CreateApplicationRequest struct {
	RegionId                         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Name                             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Description                      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Type                             *string `json:"Type,omitempty" xml:"Type,omitempty"`
	ResourceGroupId                  *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	CloudMonitorContactGroupList     *string `json:"CloudMonitorContactGroupList,omitempty" xml:"CloudMonitorContactGroupList,omitempty"`
	CloudMonitorEnableSubscribeEvent *bool   `json:"CloudMonitorEnableSubscribeEvent,omitempty" xml:"CloudMonitorEnableSubscribeEvent,omitempty"`
	CloudMonitorEnableInstallAgent   *bool   `json:"CloudMonitorEnableInstallAgent,omitempty" xml:"CloudMonitorEnableInstallAgent,omitempty"`
	CloudMonitorTemplateIdList       *string `json:"CloudMonitorTemplateIdList,omitempty" xml:"CloudMonitorTemplateIdList,omitempty"`
	CloudMonitorRuleEnabled          *bool   `json:"CloudMonitorRuleEnabled,omitempty" xml:"CloudMonitorRuleEnabled,omitempty"`
}

func (s CreateApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationRequest) GoString() string {
	return s.String()
}

func (s *CreateApplicationRequest) SetRegionId(v string) *CreateApplicationRequest {
	s.RegionId = &v
	return s
}

func (s *CreateApplicationRequest) SetName(v string) *CreateApplicationRequest {
	s.Name = &v
	return s
}

func (s *CreateApplicationRequest) SetDescription(v string) *CreateApplicationRequest {
	s.Description = &v
	return s
}

func (s *CreateApplicationRequest) SetType(v string) *CreateApplicationRequest {
	s.Type = &v
	return s
}

func (s *CreateApplicationRequest) SetResourceGroupId(v string) *CreateApplicationRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateApplicationRequest) SetCloudMonitorContactGroupList(v string) *CreateApplicationRequest {
	s.CloudMonitorContactGroupList = &v
	return s
}

func (s *CreateApplicationRequest) SetCloudMonitorEnableSubscribeEvent(v bool) *CreateApplicationRequest {
	s.CloudMonitorEnableSubscribeEvent = &v
	return s
}

func (s *CreateApplicationRequest) SetCloudMonitorEnableInstallAgent(v bool) *CreateApplicationRequest {
	s.CloudMonitorEnableInstallAgent = &v
	return s
}

func (s *CreateApplicationRequest) SetCloudMonitorTemplateIdList(v string) *CreateApplicationRequest {
	s.CloudMonitorTemplateIdList = &v
	return s
}

func (s *CreateApplicationRequest) SetCloudMonitorRuleEnabled(v bool) *CreateApplicationRequest {
	s.CloudMonitorRuleEnabled = &v
	return s
}

type CreateApplicationResponseBody struct {
	RequestId   *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Application *CreateApplicationResponseBodyApplication `json:"Application,omitempty" xml:"Application,omitempty" type:"Struct"`
}

func (s CreateApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateApplicationResponseBody) SetRequestId(v string) *CreateApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateApplicationResponseBody) SetApplication(v *CreateApplicationResponseBodyApplication) *CreateApplicationResponseBody {
	s.Application = v
	return s
}

type CreateApplicationResponseBodyApplication struct {
	Type             *string                                                   `json:"Type,omitempty" xml:"Type,omitempty"`
	IsSystem         *bool                                                     `json:"IsSystem,omitempty" xml:"IsSystem,omitempty"`
	Description      *string                                                   `json:"Description,omitempty" xml:"Description,omitempty"`
	UpdateDate       *string                                                   `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
	Name             *string                                                   `json:"Name,omitempty" xml:"Name,omitempty"`
	CreateDate       *string                                                   `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	CloudMonitorRule *CreateApplicationResponseBodyApplicationCloudMonitorRule `json:"CloudMonitorRule,omitempty" xml:"CloudMonitorRule,omitempty" type:"Struct"`
}

func (s CreateApplicationResponseBodyApplication) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationResponseBodyApplication) GoString() string {
	return s.String()
}

func (s *CreateApplicationResponseBodyApplication) SetType(v string) *CreateApplicationResponseBodyApplication {
	s.Type = &v
	return s
}

func (s *CreateApplicationResponseBodyApplication) SetIsSystem(v bool) *CreateApplicationResponseBodyApplication {
	s.IsSystem = &v
	return s
}

func (s *CreateApplicationResponseBodyApplication) SetDescription(v string) *CreateApplicationResponseBodyApplication {
	s.Description = &v
	return s
}

func (s *CreateApplicationResponseBodyApplication) SetUpdateDate(v string) *CreateApplicationResponseBodyApplication {
	s.UpdateDate = &v
	return s
}

func (s *CreateApplicationResponseBodyApplication) SetName(v string) *CreateApplicationResponseBodyApplication {
	s.Name = &v
	return s
}

func (s *CreateApplicationResponseBodyApplication) SetCreateDate(v string) *CreateApplicationResponseBodyApplication {
	s.CreateDate = &v
	return s
}

func (s *CreateApplicationResponseBodyApplication) SetCloudMonitorRule(v *CreateApplicationResponseBodyApplicationCloudMonitorRule) *CreateApplicationResponseBodyApplication {
	s.CloudMonitorRule = v
	return s
}

type CreateApplicationResponseBodyApplicationCloudMonitorRule struct {
	EnableSubscribeEvent *bool     `json:"EnableSubscribeEvent,omitempty" xml:"EnableSubscribeEvent,omitempty"`
	EnableInstallAgent   *bool     `json:"EnableInstallAgent,omitempty" xml:"EnableInstallAgent,omitempty"`
	Enabled              *bool     `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	ContactGroupList     []*string `json:"ContactGroupList,omitempty" xml:"ContactGroupList,omitempty" type:"Repeated"`
	TemplateIdList       []*int32  `json:"TemplateIdList,omitempty" xml:"TemplateIdList,omitempty" type:"Repeated"`
}

func (s CreateApplicationResponseBodyApplicationCloudMonitorRule) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationResponseBodyApplicationCloudMonitorRule) GoString() string {
	return s.String()
}

func (s *CreateApplicationResponseBodyApplicationCloudMonitorRule) SetEnableSubscribeEvent(v bool) *CreateApplicationResponseBodyApplicationCloudMonitorRule {
	s.EnableSubscribeEvent = &v
	return s
}

func (s *CreateApplicationResponseBodyApplicationCloudMonitorRule) SetEnableInstallAgent(v bool) *CreateApplicationResponseBodyApplicationCloudMonitorRule {
	s.EnableInstallAgent = &v
	return s
}

func (s *CreateApplicationResponseBodyApplicationCloudMonitorRule) SetEnabled(v bool) *CreateApplicationResponseBodyApplicationCloudMonitorRule {
	s.Enabled = &v
	return s
}

func (s *CreateApplicationResponseBodyApplicationCloudMonitorRule) SetContactGroupList(v []*string) *CreateApplicationResponseBodyApplicationCloudMonitorRule {
	s.ContactGroupList = v
	return s
}

func (s *CreateApplicationResponseBodyApplicationCloudMonitorRule) SetTemplateIdList(v []*int32) *CreateApplicationResponseBodyApplicationCloudMonitorRule {
	s.TemplateIdList = v
	return s
}

type CreateApplicationResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationResponse) GoString() string {
	return s.String()
}

func (s *CreateApplicationResponse) SetHeaders(v map[string]*string) *CreateApplicationResponse {
	s.Headers = v
	return s
}

func (s *CreateApplicationResponse) SetBody(v *CreateApplicationResponseBody) *CreateApplicationResponse {
	s.Body = v
	return s
}

type CreateApplicationGroupRequest struct {
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DeployRegionId  *string `json:"DeployRegionId,omitempty" xml:"DeployRegionId,omitempty"`
	Environment     *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	CreateType      *string `json:"CreateType,omitempty" xml:"CreateType,omitempty"`
	ImportClusterId *string `json:"ImportClusterId,omitempty" xml:"ImportClusterId,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateApplicationGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateApplicationGroupRequest) SetRegionId(v string) *CreateApplicationGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateApplicationGroupRequest) SetApplicationName(v string) *CreateApplicationGroupRequest {
	s.ApplicationName = &v
	return s
}

func (s *CreateApplicationGroupRequest) SetDescription(v string) *CreateApplicationGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateApplicationGroupRequest) SetDeployRegionId(v string) *CreateApplicationGroupRequest {
	s.DeployRegionId = &v
	return s
}

func (s *CreateApplicationGroupRequest) SetEnvironment(v string) *CreateApplicationGroupRequest {
	s.Environment = &v
	return s
}

func (s *CreateApplicationGroupRequest) SetCreateType(v string) *CreateApplicationGroupRequest {
	s.CreateType = &v
	return s
}

func (s *CreateApplicationGroupRequest) SetImportClusterId(v string) *CreateApplicationGroupRequest {
	s.ImportClusterId = &v
	return s
}

func (s *CreateApplicationGroupRequest) SetName(v string) *CreateApplicationGroupRequest {
	s.Name = &v
	return s
}

type CreateApplicationGroupResponseBody struct {
	RequestId        *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ApplicationGroup *CreateApplicationGroupResponseBodyApplicationGroup `json:"ApplicationGroup,omitempty" xml:"ApplicationGroup,omitempty" type:"Struct"`
}

func (s CreateApplicationGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateApplicationGroupResponseBody) SetRequestId(v string) *CreateApplicationGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateApplicationGroupResponseBody) SetApplicationGroup(v *CreateApplicationGroupResponseBodyApplicationGroup) *CreateApplicationGroupResponseBody {
	s.ApplicationGroup = v
	return s
}

type CreateApplicationGroupResponseBodyApplicationGroup struct {
	DeployRegionId  *string `json:"DeployRegionId,omitempty" xml:"DeployRegionId,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	UpdatedDate     *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
	CreatedDate     *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Environment     *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	CreateType      *string `json:"CreateType,omitempty" xml:"CreateType,omitempty"`
	ScalingGroupId  *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	ImportClusterId *string `json:"ImportClusterId,omitempty" xml:"ImportClusterId,omitempty"`
}

func (s CreateApplicationGroupResponseBodyApplicationGroup) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationGroupResponseBodyApplicationGroup) GoString() string {
	return s.String()
}

func (s *CreateApplicationGroupResponseBodyApplicationGroup) SetDeployRegionId(v string) *CreateApplicationGroupResponseBodyApplicationGroup {
	s.DeployRegionId = &v
	return s
}

func (s *CreateApplicationGroupResponseBodyApplicationGroup) SetDescription(v string) *CreateApplicationGroupResponseBodyApplicationGroup {
	s.Description = &v
	return s
}

func (s *CreateApplicationGroupResponseBodyApplicationGroup) SetUpdatedDate(v string) *CreateApplicationGroupResponseBodyApplicationGroup {
	s.UpdatedDate = &v
	return s
}

func (s *CreateApplicationGroupResponseBodyApplicationGroup) SetCreatedDate(v string) *CreateApplicationGroupResponseBodyApplicationGroup {
	s.CreatedDate = &v
	return s
}

func (s *CreateApplicationGroupResponseBodyApplicationGroup) SetApplicationName(v string) *CreateApplicationGroupResponseBodyApplicationGroup {
	s.ApplicationName = &v
	return s
}

func (s *CreateApplicationGroupResponseBodyApplicationGroup) SetName(v string) *CreateApplicationGroupResponseBodyApplicationGroup {
	s.Name = &v
	return s
}

func (s *CreateApplicationGroupResponseBodyApplicationGroup) SetEnvironment(v string) *CreateApplicationGroupResponseBodyApplicationGroup {
	s.Environment = &v
	return s
}

func (s *CreateApplicationGroupResponseBodyApplicationGroup) SetCreateType(v string) *CreateApplicationGroupResponseBodyApplicationGroup {
	s.CreateType = &v
	return s
}

func (s *CreateApplicationGroupResponseBodyApplicationGroup) SetScalingGroupId(v string) *CreateApplicationGroupResponseBodyApplicationGroup {
	s.ScalingGroupId = &v
	return s
}

func (s *CreateApplicationGroupResponseBodyApplicationGroup) SetImportClusterId(v string) *CreateApplicationGroupResponseBodyApplicationGroup {
	s.ImportClusterId = &v
	return s
}

type CreateApplicationGroupResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateApplicationGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateApplicationGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateApplicationGroupResponse) SetHeaders(v map[string]*string) *CreateApplicationGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateApplicationGroupResponse) SetBody(v *CreateApplicationGroupResponseBody) *CreateApplicationGroupResponse {
	s.Body = v
	return s
}

type CreateParameterRequest struct {
	RegionId        *string                `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Name            *string                `json:"Name,omitempty" xml:"Name,omitempty"`
	Type            *string                `json:"Type,omitempty" xml:"Type,omitempty"`
	Value           *string                `json:"Value,omitempty" xml:"Value,omitempty"`
	Description     *string                `json:"Description,omitempty" xml:"Description,omitempty"`
	ClientToken     *string                `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Constraints     *string                `json:"Constraints,omitempty" xml:"Constraints,omitempty"`
	Tags            map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	ResourceGroupId *string                `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s CreateParameterRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateParameterRequest) GoString() string {
	return s.String()
}

func (s *CreateParameterRequest) SetRegionId(v string) *CreateParameterRequest {
	s.RegionId = &v
	return s
}

func (s *CreateParameterRequest) SetName(v string) *CreateParameterRequest {
	s.Name = &v
	return s
}

func (s *CreateParameterRequest) SetType(v string) *CreateParameterRequest {
	s.Type = &v
	return s
}

func (s *CreateParameterRequest) SetValue(v string) *CreateParameterRequest {
	s.Value = &v
	return s
}

func (s *CreateParameterRequest) SetDescription(v string) *CreateParameterRequest {
	s.Description = &v
	return s
}

func (s *CreateParameterRequest) SetClientToken(v string) *CreateParameterRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateParameterRequest) SetConstraints(v string) *CreateParameterRequest {
	s.Constraints = &v
	return s
}

func (s *CreateParameterRequest) SetTags(v map[string]interface{}) *CreateParameterRequest {
	s.Tags = v
	return s
}

func (s *CreateParameterRequest) SetResourceGroupId(v string) *CreateParameterRequest {
	s.ResourceGroupId = &v
	return s
}

type CreateParameterShrinkRequest struct {
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Type            *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value           *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ClientToken     *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Constraints     *string `json:"Constraints,omitempty" xml:"Constraints,omitempty"`
	TagsShrink      *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s CreateParameterShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateParameterShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateParameterShrinkRequest) SetRegionId(v string) *CreateParameterShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateParameterShrinkRequest) SetName(v string) *CreateParameterShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateParameterShrinkRequest) SetType(v string) *CreateParameterShrinkRequest {
	s.Type = &v
	return s
}

func (s *CreateParameterShrinkRequest) SetValue(v string) *CreateParameterShrinkRequest {
	s.Value = &v
	return s
}

func (s *CreateParameterShrinkRequest) SetDescription(v string) *CreateParameterShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateParameterShrinkRequest) SetClientToken(v string) *CreateParameterShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateParameterShrinkRequest) SetConstraints(v string) *CreateParameterShrinkRequest {
	s.Constraints = &v
	return s
}

func (s *CreateParameterShrinkRequest) SetTagsShrink(v string) *CreateParameterShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *CreateParameterShrinkRequest) SetResourceGroupId(v string) *CreateParameterShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

type CreateParameterResponseBody struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Parameter *CreateParameterResponseBodyParameter `json:"Parameter,omitempty" xml:"Parameter,omitempty" type:"Struct"`
}

func (s CreateParameterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateParameterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateParameterResponseBody) SetRequestId(v string) *CreateParameterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateParameterResponseBody) SetParameter(v *CreateParameterResponseBodyParameter) *CreateParameterResponseBody {
	s.Parameter = v
	return s
}

type CreateParameterResponseBodyParameter struct {
	Type             *string                `json:"Type,omitempty" xml:"Type,omitempty"`
	UpdatedDate      *string                `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
	UpdatedBy        *string                `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	Tags             map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	Description      *string                `json:"Description,omitempty" xml:"Description,omitempty"`
	Constraints      *string                `json:"Constraints,omitempty" xml:"Constraints,omitempty"`
	ResourceGroupId  *string                `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	CreatedBy        *string                `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	CreatedDate      *string                `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	ParameterVersion *int32                 `json:"ParameterVersion,omitempty" xml:"ParameterVersion,omitempty"`
	Name             *string                `json:"Name,omitempty" xml:"Name,omitempty"`
	Id               *string                `json:"Id,omitempty" xml:"Id,omitempty"`
	ShareType        *string                `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
}

func (s CreateParameterResponseBodyParameter) String() string {
	return tea.Prettify(s)
}

func (s CreateParameterResponseBodyParameter) GoString() string {
	return s.String()
}

func (s *CreateParameterResponseBodyParameter) SetType(v string) *CreateParameterResponseBodyParameter {
	s.Type = &v
	return s
}

func (s *CreateParameterResponseBodyParameter) SetUpdatedDate(v string) *CreateParameterResponseBodyParameter {
	s.UpdatedDate = &v
	return s
}

func (s *CreateParameterResponseBodyParameter) SetUpdatedBy(v string) *CreateParameterResponseBodyParameter {
	s.UpdatedBy = &v
	return s
}

func (s *CreateParameterResponseBodyParameter) SetTags(v map[string]interface{}) *CreateParameterResponseBodyParameter {
	s.Tags = v
	return s
}

func (s *CreateParameterResponseBodyParameter) SetDescription(v string) *CreateParameterResponseBodyParameter {
	s.Description = &v
	return s
}

func (s *CreateParameterResponseBodyParameter) SetConstraints(v string) *CreateParameterResponseBodyParameter {
	s.Constraints = &v
	return s
}

func (s *CreateParameterResponseBodyParameter) SetResourceGroupId(v string) *CreateParameterResponseBodyParameter {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateParameterResponseBodyParameter) SetCreatedBy(v string) *CreateParameterResponseBodyParameter {
	s.CreatedBy = &v
	return s
}

func (s *CreateParameterResponseBodyParameter) SetCreatedDate(v string) *CreateParameterResponseBodyParameter {
	s.CreatedDate = &v
	return s
}

func (s *CreateParameterResponseBodyParameter) SetParameterVersion(v int32) *CreateParameterResponseBodyParameter {
	s.ParameterVersion = &v
	return s
}

func (s *CreateParameterResponseBodyParameter) SetName(v string) *CreateParameterResponseBodyParameter {
	s.Name = &v
	return s
}

func (s *CreateParameterResponseBodyParameter) SetId(v string) *CreateParameterResponseBodyParameter {
	s.Id = &v
	return s
}

func (s *CreateParameterResponseBodyParameter) SetShareType(v string) *CreateParameterResponseBodyParameter {
	s.ShareType = &v
	return s
}

type CreateParameterResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateParameterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateParameterResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateParameterResponse) GoString() string {
	return s.String()
}

func (s *CreateParameterResponse) SetHeaders(v map[string]*string) *CreateParameterResponse {
	s.Headers = v
	return s
}

func (s *CreateParameterResponse) SetBody(v *CreateParameterResponseBody) *CreateParameterResponse {
	s.Body = v
	return s
}

type CreatePatchBaselineRequest struct {
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ClientToken     *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OperationSystem *string `json:"OperationSystem,omitempty" xml:"OperationSystem,omitempty"`
	ApprovalRules   *string `json:"ApprovalRules,omitempty" xml:"ApprovalRules,omitempty"`
}

func (s CreatePatchBaselineRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePatchBaselineRequest) GoString() string {
	return s.String()
}

func (s *CreatePatchBaselineRequest) SetRegionId(v string) *CreatePatchBaselineRequest {
	s.RegionId = &v
	return s
}

func (s *CreatePatchBaselineRequest) SetName(v string) *CreatePatchBaselineRequest {
	s.Name = &v
	return s
}

func (s *CreatePatchBaselineRequest) SetDescription(v string) *CreatePatchBaselineRequest {
	s.Description = &v
	return s
}

func (s *CreatePatchBaselineRequest) SetClientToken(v string) *CreatePatchBaselineRequest {
	s.ClientToken = &v
	return s
}

func (s *CreatePatchBaselineRequest) SetOperationSystem(v string) *CreatePatchBaselineRequest {
	s.OperationSystem = &v
	return s
}

func (s *CreatePatchBaselineRequest) SetApprovalRules(v string) *CreatePatchBaselineRequest {
	s.ApprovalRules = &v
	return s
}

type CreatePatchBaselineResponseBody struct {
	RequestId     *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PatchBaseline *CreatePatchBaselineResponseBodyPatchBaseline `json:"PatchBaseline,omitempty" xml:"PatchBaseline,omitempty" type:"Struct"`
}

func (s CreatePatchBaselineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePatchBaselineResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePatchBaselineResponseBody) SetRequestId(v string) *CreatePatchBaselineResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePatchBaselineResponseBody) SetPatchBaseline(v *CreatePatchBaselineResponseBodyPatchBaseline) *CreatePatchBaselineResponseBody {
	s.PatchBaseline = v
	return s
}

type CreatePatchBaselineResponseBodyPatchBaseline struct {
	OperationSystem *string `json:"OperationSystem,omitempty" xml:"OperationSystem,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	UpdatedDate     *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
	UpdatedBy       *string `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	CreatedBy       *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	CreatedDate     *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ApprovalRules   *string `json:"ApprovalRules,omitempty" xml:"ApprovalRules,omitempty"`
	Id              *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ShareType       *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
}

func (s CreatePatchBaselineResponseBodyPatchBaseline) String() string {
	return tea.Prettify(s)
}

func (s CreatePatchBaselineResponseBodyPatchBaseline) GoString() string {
	return s.String()
}

func (s *CreatePatchBaselineResponseBodyPatchBaseline) SetOperationSystem(v string) *CreatePatchBaselineResponseBodyPatchBaseline {
	s.OperationSystem = &v
	return s
}

func (s *CreatePatchBaselineResponseBodyPatchBaseline) SetDescription(v string) *CreatePatchBaselineResponseBodyPatchBaseline {
	s.Description = &v
	return s
}

func (s *CreatePatchBaselineResponseBodyPatchBaseline) SetUpdatedDate(v string) *CreatePatchBaselineResponseBodyPatchBaseline {
	s.UpdatedDate = &v
	return s
}

func (s *CreatePatchBaselineResponseBodyPatchBaseline) SetUpdatedBy(v string) *CreatePatchBaselineResponseBodyPatchBaseline {
	s.UpdatedBy = &v
	return s
}

func (s *CreatePatchBaselineResponseBodyPatchBaseline) SetCreatedBy(v string) *CreatePatchBaselineResponseBodyPatchBaseline {
	s.CreatedBy = &v
	return s
}

func (s *CreatePatchBaselineResponseBodyPatchBaseline) SetCreatedDate(v string) *CreatePatchBaselineResponseBodyPatchBaseline {
	s.CreatedDate = &v
	return s
}

func (s *CreatePatchBaselineResponseBodyPatchBaseline) SetName(v string) *CreatePatchBaselineResponseBodyPatchBaseline {
	s.Name = &v
	return s
}

func (s *CreatePatchBaselineResponseBodyPatchBaseline) SetApprovalRules(v string) *CreatePatchBaselineResponseBodyPatchBaseline {
	s.ApprovalRules = &v
	return s
}

func (s *CreatePatchBaselineResponseBodyPatchBaseline) SetId(v string) *CreatePatchBaselineResponseBodyPatchBaseline {
	s.Id = &v
	return s
}

func (s *CreatePatchBaselineResponseBodyPatchBaseline) SetShareType(v string) *CreatePatchBaselineResponseBodyPatchBaseline {
	s.ShareType = &v
	return s
}

type CreatePatchBaselineResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreatePatchBaselineResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreatePatchBaselineResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePatchBaselineResponse) GoString() string {
	return s.String()
}

func (s *CreatePatchBaselineResponse) SetHeaders(v map[string]*string) *CreatePatchBaselineResponse {
	s.Headers = v
	return s
}

func (s *CreatePatchBaselineResponse) SetBody(v *CreatePatchBaselineResponseBody) *CreatePatchBaselineResponse {
	s.Body = v
	return s
}

type CreateSecretParameterRequest struct {
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Type            *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value           *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	KeyId           *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	ClientToken     *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Constraints     *string `json:"Constraints,omitempty" xml:"Constraints,omitempty"`
	Tags            *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s CreateSecretParameterRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSecretParameterRequest) GoString() string {
	return s.String()
}

func (s *CreateSecretParameterRequest) SetRegionId(v string) *CreateSecretParameterRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSecretParameterRequest) SetName(v string) *CreateSecretParameterRequest {
	s.Name = &v
	return s
}

func (s *CreateSecretParameterRequest) SetType(v string) *CreateSecretParameterRequest {
	s.Type = &v
	return s
}

func (s *CreateSecretParameterRequest) SetValue(v string) *CreateSecretParameterRequest {
	s.Value = &v
	return s
}

func (s *CreateSecretParameterRequest) SetDescription(v string) *CreateSecretParameterRequest {
	s.Description = &v
	return s
}

func (s *CreateSecretParameterRequest) SetKeyId(v string) *CreateSecretParameterRequest {
	s.KeyId = &v
	return s
}

func (s *CreateSecretParameterRequest) SetClientToken(v string) *CreateSecretParameterRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateSecretParameterRequest) SetConstraints(v string) *CreateSecretParameterRequest {
	s.Constraints = &v
	return s
}

func (s *CreateSecretParameterRequest) SetTags(v string) *CreateSecretParameterRequest {
	s.Tags = &v
	return s
}

func (s *CreateSecretParameterRequest) SetResourceGroupId(v string) *CreateSecretParameterRequest {
	s.ResourceGroupId = &v
	return s
}

type CreateSecretParameterResponseBody struct {
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Parameter *CreateSecretParameterResponseBodyParameter `json:"Parameter,omitempty" xml:"Parameter,omitempty" type:"Struct"`
}

func (s CreateSecretParameterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSecretParameterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSecretParameterResponseBody) SetRequestId(v string) *CreateSecretParameterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSecretParameterResponseBody) SetParameter(v *CreateSecretParameterResponseBodyParameter) *CreateSecretParameterResponseBody {
	s.Parameter = v
	return s
}

type CreateSecretParameterResponseBodyParameter struct {
	Type             *string `json:"Type,omitempty" xml:"Type,omitempty"`
	UpdatedDate      *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
	UpdatedBy        *string `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	KeyId            *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	Tags             *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Constraints      *string `json:"Constraints,omitempty" xml:"Constraints,omitempty"`
	ResourceGroupId  *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	CreatedBy        *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	CreatedDate      *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	ParameterVersion *int32  `json:"ParameterVersion,omitempty" xml:"ParameterVersion,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id               *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ShareType        *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
}

func (s CreateSecretParameterResponseBodyParameter) String() string {
	return tea.Prettify(s)
}

func (s CreateSecretParameterResponseBodyParameter) GoString() string {
	return s.String()
}

func (s *CreateSecretParameterResponseBodyParameter) SetType(v string) *CreateSecretParameterResponseBodyParameter {
	s.Type = &v
	return s
}

func (s *CreateSecretParameterResponseBodyParameter) SetUpdatedDate(v string) *CreateSecretParameterResponseBodyParameter {
	s.UpdatedDate = &v
	return s
}

func (s *CreateSecretParameterResponseBodyParameter) SetUpdatedBy(v string) *CreateSecretParameterResponseBodyParameter {
	s.UpdatedBy = &v
	return s
}

func (s *CreateSecretParameterResponseBodyParameter) SetKeyId(v string) *CreateSecretParameterResponseBodyParameter {
	s.KeyId = &v
	return s
}

func (s *CreateSecretParameterResponseBodyParameter) SetTags(v string) *CreateSecretParameterResponseBodyParameter {
	s.Tags = &v
	return s
}

func (s *CreateSecretParameterResponseBodyParameter) SetDescription(v string) *CreateSecretParameterResponseBodyParameter {
	s.Description = &v
	return s
}

func (s *CreateSecretParameterResponseBodyParameter) SetConstraints(v string) *CreateSecretParameterResponseBodyParameter {
	s.Constraints = &v
	return s
}

func (s *CreateSecretParameterResponseBodyParameter) SetResourceGroupId(v string) *CreateSecretParameterResponseBodyParameter {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateSecretParameterResponseBodyParameter) SetCreatedBy(v string) *CreateSecretParameterResponseBodyParameter {
	s.CreatedBy = &v
	return s
}

func (s *CreateSecretParameterResponseBodyParameter) SetCreatedDate(v string) *CreateSecretParameterResponseBodyParameter {
	s.CreatedDate = &v
	return s
}

func (s *CreateSecretParameterResponseBodyParameter) SetParameterVersion(v int32) *CreateSecretParameterResponseBodyParameter {
	s.ParameterVersion = &v
	return s
}

func (s *CreateSecretParameterResponseBodyParameter) SetName(v string) *CreateSecretParameterResponseBodyParameter {
	s.Name = &v
	return s
}

func (s *CreateSecretParameterResponseBodyParameter) SetId(v string) *CreateSecretParameterResponseBodyParameter {
	s.Id = &v
	return s
}

func (s *CreateSecretParameterResponseBodyParameter) SetShareType(v string) *CreateSecretParameterResponseBodyParameter {
	s.ShareType = &v
	return s
}

type CreateSecretParameterResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateSecretParameterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSecretParameterResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSecretParameterResponse) GoString() string {
	return s.String()
}

func (s *CreateSecretParameterResponse) SetHeaders(v map[string]*string) *CreateSecretParameterResponse {
	s.Headers = v
	return s
}

func (s *CreateSecretParameterResponse) SetBody(v *CreateSecretParameterResponseBody) *CreateSecretParameterResponse {
	s.Body = v
	return s
}

type CreateStateConfigurationRequest struct {
	RegionId           *string                `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Description        *string                `json:"Description,omitempty" xml:"Description,omitempty"`
	TemplateName       *string                `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateVersion    *string                `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	Parameters         *string                `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	ConfigureMode      *string                `json:"ConfigureMode,omitempty" xml:"ConfigureMode,omitempty"`
	ScheduleType       *string                `json:"ScheduleType,omitempty" xml:"ScheduleType,omitempty"`
	ScheduleExpression *string                `json:"ScheduleExpression,omitempty" xml:"ScheduleExpression,omitempty"`
	Targets            *string                `json:"Targets,omitempty" xml:"Targets,omitempty"`
	ClientToken        *string                `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Tags               map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	ResourceGroupId    *string                `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s CreateStateConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateStateConfigurationRequest) GoString() string {
	return s.String()
}

func (s *CreateStateConfigurationRequest) SetRegionId(v string) *CreateStateConfigurationRequest {
	s.RegionId = &v
	return s
}

func (s *CreateStateConfigurationRequest) SetDescription(v string) *CreateStateConfigurationRequest {
	s.Description = &v
	return s
}

func (s *CreateStateConfigurationRequest) SetTemplateName(v string) *CreateStateConfigurationRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateStateConfigurationRequest) SetTemplateVersion(v string) *CreateStateConfigurationRequest {
	s.TemplateVersion = &v
	return s
}

func (s *CreateStateConfigurationRequest) SetParameters(v string) *CreateStateConfigurationRequest {
	s.Parameters = &v
	return s
}

func (s *CreateStateConfigurationRequest) SetConfigureMode(v string) *CreateStateConfigurationRequest {
	s.ConfigureMode = &v
	return s
}

func (s *CreateStateConfigurationRequest) SetScheduleType(v string) *CreateStateConfigurationRequest {
	s.ScheduleType = &v
	return s
}

func (s *CreateStateConfigurationRequest) SetScheduleExpression(v string) *CreateStateConfigurationRequest {
	s.ScheduleExpression = &v
	return s
}

func (s *CreateStateConfigurationRequest) SetTargets(v string) *CreateStateConfigurationRequest {
	s.Targets = &v
	return s
}

func (s *CreateStateConfigurationRequest) SetClientToken(v string) *CreateStateConfigurationRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateStateConfigurationRequest) SetTags(v map[string]interface{}) *CreateStateConfigurationRequest {
	s.Tags = v
	return s
}

func (s *CreateStateConfigurationRequest) SetResourceGroupId(v string) *CreateStateConfigurationRequest {
	s.ResourceGroupId = &v
	return s
}

type CreateStateConfigurationShrinkRequest struct {
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Description        *string `json:"Description,omitempty" xml:"Description,omitempty"`
	TemplateName       *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateVersion    *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	Parameters         *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	ConfigureMode      *string `json:"ConfigureMode,omitempty" xml:"ConfigureMode,omitempty"`
	ScheduleType       *string `json:"ScheduleType,omitempty" xml:"ScheduleType,omitempty"`
	ScheduleExpression *string `json:"ScheduleExpression,omitempty" xml:"ScheduleExpression,omitempty"`
	Targets            *string `json:"Targets,omitempty" xml:"Targets,omitempty"`
	ClientToken        *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	TagsShrink         *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	ResourceGroupId    *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s CreateStateConfigurationShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateStateConfigurationShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateStateConfigurationShrinkRequest) SetRegionId(v string) *CreateStateConfigurationShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateStateConfigurationShrinkRequest) SetDescription(v string) *CreateStateConfigurationShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateStateConfigurationShrinkRequest) SetTemplateName(v string) *CreateStateConfigurationShrinkRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateStateConfigurationShrinkRequest) SetTemplateVersion(v string) *CreateStateConfigurationShrinkRequest {
	s.TemplateVersion = &v
	return s
}

func (s *CreateStateConfigurationShrinkRequest) SetParameters(v string) *CreateStateConfigurationShrinkRequest {
	s.Parameters = &v
	return s
}

func (s *CreateStateConfigurationShrinkRequest) SetConfigureMode(v string) *CreateStateConfigurationShrinkRequest {
	s.ConfigureMode = &v
	return s
}

func (s *CreateStateConfigurationShrinkRequest) SetScheduleType(v string) *CreateStateConfigurationShrinkRequest {
	s.ScheduleType = &v
	return s
}

func (s *CreateStateConfigurationShrinkRequest) SetScheduleExpression(v string) *CreateStateConfigurationShrinkRequest {
	s.ScheduleExpression = &v
	return s
}

func (s *CreateStateConfigurationShrinkRequest) SetTargets(v string) *CreateStateConfigurationShrinkRequest {
	s.Targets = &v
	return s
}

func (s *CreateStateConfigurationShrinkRequest) SetClientToken(v string) *CreateStateConfigurationShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateStateConfigurationShrinkRequest) SetTagsShrink(v string) *CreateStateConfigurationShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *CreateStateConfigurationShrinkRequest) SetResourceGroupId(v string) *CreateStateConfigurationShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

type CreateStateConfigurationResponseBody struct {
	RequestId          *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StateConfiguration *CreateStateConfigurationResponseBodyStateConfiguration `json:"StateConfiguration,omitempty" xml:"StateConfiguration,omitempty" type:"Struct"`
}

func (s CreateStateConfigurationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateStateConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateStateConfigurationResponseBody) SetRequestId(v string) *CreateStateConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateStateConfigurationResponseBody) SetStateConfiguration(v *CreateStateConfigurationResponseBodyStateConfiguration) *CreateStateConfigurationResponseBody {
	s.StateConfiguration = v
	return s
}

type CreateStateConfigurationResponseBodyStateConfiguration struct {
	CreateTime           *string                `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Targets              *string                `json:"Targets,omitempty" xml:"Targets,omitempty"`
	Tags                 map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	StateConfigurationId *string                `json:"StateConfigurationId,omitempty" xml:"StateConfigurationId,omitempty"`
	ScheduleExpression   *string                `json:"ScheduleExpression,omitempty" xml:"ScheduleExpression,omitempty"`
	TemplateName         *string                `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateVersion      *string                `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	ConfigureMode        *string                `json:"ConfigureMode,omitempty" xml:"ConfigureMode,omitempty"`
	ScheduleType         *string                `json:"ScheduleType,omitempty" xml:"ScheduleType,omitempty"`
	Parameters           map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	Description          *string                `json:"Description,omitempty" xml:"Description,omitempty"`
	ResourceGroupId      *string                `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	TemplateId           *string                `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s CreateStateConfigurationResponseBodyStateConfiguration) String() string {
	return tea.Prettify(s)
}

func (s CreateStateConfigurationResponseBodyStateConfiguration) GoString() string {
	return s.String()
}

func (s *CreateStateConfigurationResponseBodyStateConfiguration) SetCreateTime(v string) *CreateStateConfigurationResponseBodyStateConfiguration {
	s.CreateTime = &v
	return s
}

func (s *CreateStateConfigurationResponseBodyStateConfiguration) SetTargets(v string) *CreateStateConfigurationResponseBodyStateConfiguration {
	s.Targets = &v
	return s
}

func (s *CreateStateConfigurationResponseBodyStateConfiguration) SetTags(v map[string]interface{}) *CreateStateConfigurationResponseBodyStateConfiguration {
	s.Tags = v
	return s
}

func (s *CreateStateConfigurationResponseBodyStateConfiguration) SetStateConfigurationId(v string) *CreateStateConfigurationResponseBodyStateConfiguration {
	s.StateConfigurationId = &v
	return s
}

func (s *CreateStateConfigurationResponseBodyStateConfiguration) SetScheduleExpression(v string) *CreateStateConfigurationResponseBodyStateConfiguration {
	s.ScheduleExpression = &v
	return s
}

func (s *CreateStateConfigurationResponseBodyStateConfiguration) SetTemplateName(v string) *CreateStateConfigurationResponseBodyStateConfiguration {
	s.TemplateName = &v
	return s
}

func (s *CreateStateConfigurationResponseBodyStateConfiguration) SetTemplateVersion(v string) *CreateStateConfigurationResponseBodyStateConfiguration {
	s.TemplateVersion = &v
	return s
}

func (s *CreateStateConfigurationResponseBodyStateConfiguration) SetConfigureMode(v string) *CreateStateConfigurationResponseBodyStateConfiguration {
	s.ConfigureMode = &v
	return s
}

func (s *CreateStateConfigurationResponseBodyStateConfiguration) SetScheduleType(v string) *CreateStateConfigurationResponseBodyStateConfiguration {
	s.ScheduleType = &v
	return s
}

func (s *CreateStateConfigurationResponseBodyStateConfiguration) SetParameters(v map[string]interface{}) *CreateStateConfigurationResponseBodyStateConfiguration {
	s.Parameters = v
	return s
}

func (s *CreateStateConfigurationResponseBodyStateConfiguration) SetDescription(v string) *CreateStateConfigurationResponseBodyStateConfiguration {
	s.Description = &v
	return s
}

func (s *CreateStateConfigurationResponseBodyStateConfiguration) SetResourceGroupId(v string) *CreateStateConfigurationResponseBodyStateConfiguration {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateStateConfigurationResponseBodyStateConfiguration) SetTemplateId(v string) *CreateStateConfigurationResponseBodyStateConfiguration {
	s.TemplateId = &v
	return s
}

type CreateStateConfigurationResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateStateConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateStateConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateStateConfigurationResponse) GoString() string {
	return s.String()
}

func (s *CreateStateConfigurationResponse) SetHeaders(v map[string]*string) *CreateStateConfigurationResponse {
	s.Headers = v
	return s
}

func (s *CreateStateConfigurationResponse) SetBody(v *CreateStateConfigurationResponseBody) *CreateStateConfigurationResponse {
	s.Body = v
	return s
}

type CreateTemplateRequest struct {
	RegionId        *string                `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TemplateName    *string                `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	Content         *string                `json:"Content,omitempty" xml:"Content,omitempty"`
	Tags            map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	VersionName     *string                `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
	ResourceGroupId *string                `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s CreateTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateTemplateRequest) SetRegionId(v string) *CreateTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *CreateTemplateRequest) SetTemplateName(v string) *CreateTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateTemplateRequest) SetContent(v string) *CreateTemplateRequest {
	s.Content = &v
	return s
}

func (s *CreateTemplateRequest) SetTags(v map[string]interface{}) *CreateTemplateRequest {
	s.Tags = v
	return s
}

func (s *CreateTemplateRequest) SetVersionName(v string) *CreateTemplateRequest {
	s.VersionName = &v
	return s
}

func (s *CreateTemplateRequest) SetResourceGroupId(v string) *CreateTemplateRequest {
	s.ResourceGroupId = &v
	return s
}

type CreateTemplateShrinkRequest struct {
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TemplateName    *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	Content         *string `json:"Content,omitempty" xml:"Content,omitempty"`
	TagsShrink      *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	VersionName     *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s CreateTemplateShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateTemplateShrinkRequest) SetRegionId(v string) *CreateTemplateShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateTemplateShrinkRequest) SetTemplateName(v string) *CreateTemplateShrinkRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateTemplateShrinkRequest) SetContent(v string) *CreateTemplateShrinkRequest {
	s.Content = &v
	return s
}

func (s *CreateTemplateShrinkRequest) SetTagsShrink(v string) *CreateTemplateShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *CreateTemplateShrinkRequest) SetVersionName(v string) *CreateTemplateShrinkRequest {
	s.VersionName = &v
	return s
}

func (s *CreateTemplateShrinkRequest) SetResourceGroupId(v string) *CreateTemplateShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

type CreateTemplateResponseBody struct {
	TemplateType *string                             `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	RequestId    *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Template     *CreateTemplateResponseBodyTemplate `json:"Template,omitempty" xml:"Template,omitempty" type:"Struct"`
}

func (s CreateTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTemplateResponseBody) SetTemplateType(v string) *CreateTemplateResponseBody {
	s.TemplateType = &v
	return s
}

func (s *CreateTemplateResponseBody) SetRequestId(v string) *CreateTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTemplateResponseBody) SetTemplate(v *CreateTemplateResponseBodyTemplate) *CreateTemplateResponseBody {
	s.Template = v
	return s
}

type CreateTemplateResponseBodyTemplate struct {
	Hash            *string                `json:"Hash,omitempty" xml:"Hash,omitempty"`
	UpdatedDate     *string                `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
	UpdatedBy       *string                `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	Tags            map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	TemplateName    *string                `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateVersion *string                `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	TemplateFormat  *string                `json:"TemplateFormat,omitempty" xml:"TemplateFormat,omitempty"`
	Description     *string                `json:"Description,omitempty" xml:"Description,omitempty"`
	ResourceGroupId *string                `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	CreatedBy       *string                `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	CreatedDate     *string                `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	TemplateId      *string                `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	HasTrigger      *bool                  `json:"HasTrigger,omitempty" xml:"HasTrigger,omitempty"`
	ShareType       *string                `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
}

func (s CreateTemplateResponseBodyTemplate) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplateResponseBodyTemplate) GoString() string {
	return s.String()
}

func (s *CreateTemplateResponseBodyTemplate) SetHash(v string) *CreateTemplateResponseBodyTemplate {
	s.Hash = &v
	return s
}

func (s *CreateTemplateResponseBodyTemplate) SetUpdatedDate(v string) *CreateTemplateResponseBodyTemplate {
	s.UpdatedDate = &v
	return s
}

func (s *CreateTemplateResponseBodyTemplate) SetUpdatedBy(v string) *CreateTemplateResponseBodyTemplate {
	s.UpdatedBy = &v
	return s
}

func (s *CreateTemplateResponseBodyTemplate) SetTags(v map[string]interface{}) *CreateTemplateResponseBodyTemplate {
	s.Tags = v
	return s
}

func (s *CreateTemplateResponseBodyTemplate) SetTemplateName(v string) *CreateTemplateResponseBodyTemplate {
	s.TemplateName = &v
	return s
}

func (s *CreateTemplateResponseBodyTemplate) SetTemplateVersion(v string) *CreateTemplateResponseBodyTemplate {
	s.TemplateVersion = &v
	return s
}

func (s *CreateTemplateResponseBodyTemplate) SetTemplateFormat(v string) *CreateTemplateResponseBodyTemplate {
	s.TemplateFormat = &v
	return s
}

func (s *CreateTemplateResponseBodyTemplate) SetDescription(v string) *CreateTemplateResponseBodyTemplate {
	s.Description = &v
	return s
}

func (s *CreateTemplateResponseBodyTemplate) SetResourceGroupId(v string) *CreateTemplateResponseBodyTemplate {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateTemplateResponseBodyTemplate) SetCreatedBy(v string) *CreateTemplateResponseBodyTemplate {
	s.CreatedBy = &v
	return s
}

func (s *CreateTemplateResponseBodyTemplate) SetCreatedDate(v string) *CreateTemplateResponseBodyTemplate {
	s.CreatedDate = &v
	return s
}

func (s *CreateTemplateResponseBodyTemplate) SetTemplateId(v string) *CreateTemplateResponseBodyTemplate {
	s.TemplateId = &v
	return s
}

func (s *CreateTemplateResponseBodyTemplate) SetHasTrigger(v bool) *CreateTemplateResponseBodyTemplate {
	s.HasTrigger = &v
	return s
}

func (s *CreateTemplateResponseBodyTemplate) SetShareType(v string) *CreateTemplateResponseBodyTemplate {
	s.ShareType = &v
	return s
}

type CreateTemplateResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateTemplateResponse) SetHeaders(v map[string]*string) *CreateTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateTemplateResponse) SetBody(v *CreateTemplateResponseBody) *CreateTemplateResponse {
	s.Body = v
	return s
}

type DeleteApplicationRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DeleteApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteApplicationRequest) GoString() string {
	return s.String()
}

func (s *DeleteApplicationRequest) SetRegionId(v string) *DeleteApplicationRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteApplicationRequest) SetName(v string) *DeleteApplicationRequest {
	s.Name = &v
	return s
}

type DeleteApplicationResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteApplicationResponseBody) SetRequestId(v string) *DeleteApplicationResponseBody {
	s.RequestId = &v
	return s
}

type DeleteApplicationResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteApplicationResponse) GoString() string {
	return s.String()
}

func (s *DeleteApplicationResponse) SetHeaders(v map[string]*string) *DeleteApplicationResponse {
	s.Headers = v
	return s
}

func (s *DeleteApplicationResponse) SetBody(v *DeleteApplicationResponseBody) *DeleteApplicationResponse {
	s.Body = v
	return s
}

type DeleteApplicationGroupRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DeleteApplicationGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteApplicationGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteApplicationGroupRequest) SetRegionId(v string) *DeleteApplicationGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteApplicationGroupRequest) SetName(v string) *DeleteApplicationGroupRequest {
	s.Name = &v
	return s
}

type DeleteApplicationGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteApplicationGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteApplicationGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteApplicationGroupResponseBody) SetRequestId(v string) *DeleteApplicationGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteApplicationGroupResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteApplicationGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteApplicationGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteApplicationGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteApplicationGroupResponse) SetHeaders(v map[string]*string) *DeleteApplicationGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteApplicationGroupResponse) SetBody(v *DeleteApplicationGroupResponseBody) *DeleteApplicationGroupResponse {
	s.Body = v
	return s
}

type DeleteExecutionsRequest struct {
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ExecutionIds *string `json:"ExecutionIds,omitempty" xml:"ExecutionIds,omitempty"`
}

func (s DeleteExecutionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteExecutionsRequest) GoString() string {
	return s.String()
}

func (s *DeleteExecutionsRequest) SetRegionId(v string) *DeleteExecutionsRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteExecutionsRequest) SetExecutionIds(v string) *DeleteExecutionsRequest {
	s.ExecutionIds = &v
	return s
}

type DeleteExecutionsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteExecutionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteExecutionsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteExecutionsResponseBody) SetRequestId(v string) *DeleteExecutionsResponseBody {
	s.RequestId = &v
	return s
}

type DeleteExecutionsResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteExecutionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteExecutionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteExecutionsResponse) GoString() string {
	return s.String()
}

func (s *DeleteExecutionsResponse) SetHeaders(v map[string]*string) *DeleteExecutionsResponse {
	s.Headers = v
	return s
}

func (s *DeleteExecutionsResponse) SetBody(v *DeleteExecutionsResponseBody) *DeleteExecutionsResponse {
	s.Body = v
	return s
}

type DeleteParameterRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DeleteParameterRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteParameterRequest) GoString() string {
	return s.String()
}

func (s *DeleteParameterRequest) SetRegionId(v string) *DeleteParameterRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteParameterRequest) SetName(v string) *DeleteParameterRequest {
	s.Name = &v
	return s
}

type DeleteParameterResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteParameterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteParameterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteParameterResponseBody) SetRequestId(v string) *DeleteParameterResponseBody {
	s.RequestId = &v
	return s
}

type DeleteParameterResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteParameterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteParameterResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteParameterResponse) GoString() string {
	return s.String()
}

func (s *DeleteParameterResponse) SetHeaders(v map[string]*string) *DeleteParameterResponse {
	s.Headers = v
	return s
}

func (s *DeleteParameterResponse) SetBody(v *DeleteParameterResponseBody) *DeleteParameterResponse {
	s.Body = v
	return s
}

type DeletePatchBaselineRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DeletePatchBaselineRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePatchBaselineRequest) GoString() string {
	return s.String()
}

func (s *DeletePatchBaselineRequest) SetRegionId(v string) *DeletePatchBaselineRequest {
	s.RegionId = &v
	return s
}

func (s *DeletePatchBaselineRequest) SetName(v string) *DeletePatchBaselineRequest {
	s.Name = &v
	return s
}

type DeletePatchBaselineResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePatchBaselineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeletePatchBaselineResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePatchBaselineResponseBody) SetRequestId(v string) *DeletePatchBaselineResponseBody {
	s.RequestId = &v
	return s
}

type DeletePatchBaselineResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeletePatchBaselineResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeletePatchBaselineResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePatchBaselineResponse) GoString() string {
	return s.String()
}

func (s *DeletePatchBaselineResponse) SetHeaders(v map[string]*string) *DeletePatchBaselineResponse {
	s.Headers = v
	return s
}

func (s *DeletePatchBaselineResponse) SetBody(v *DeletePatchBaselineResponseBody) *DeletePatchBaselineResponse {
	s.Body = v
	return s
}

type DeleteSecretParameterRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DeleteSecretParameterRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSecretParameterRequest) GoString() string {
	return s.String()
}

func (s *DeleteSecretParameterRequest) SetRegionId(v string) *DeleteSecretParameterRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteSecretParameterRequest) SetName(v string) *DeleteSecretParameterRequest {
	s.Name = &v
	return s
}

type DeleteSecretParameterResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSecretParameterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSecretParameterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSecretParameterResponseBody) SetRequestId(v string) *DeleteSecretParameterResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSecretParameterResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteSecretParameterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSecretParameterResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSecretParameterResponse) GoString() string {
	return s.String()
}

func (s *DeleteSecretParameterResponse) SetHeaders(v map[string]*string) *DeleteSecretParameterResponse {
	s.Headers = v
	return s
}

func (s *DeleteSecretParameterResponse) SetBody(v *DeleteSecretParameterResponseBody) *DeleteSecretParameterResponse {
	s.Body = v
	return s
}

type DeleteStateConfigurationsRequest struct {
	RegionId              *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StateConfigurationIds *string `json:"StateConfigurationIds,omitempty" xml:"StateConfigurationIds,omitempty"`
	ClientToken           *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s DeleteStateConfigurationsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteStateConfigurationsRequest) GoString() string {
	return s.String()
}

func (s *DeleteStateConfigurationsRequest) SetRegionId(v string) *DeleteStateConfigurationsRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteStateConfigurationsRequest) SetStateConfigurationIds(v string) *DeleteStateConfigurationsRequest {
	s.StateConfigurationIds = &v
	return s
}

func (s *DeleteStateConfigurationsRequest) SetClientToken(v string) *DeleteStateConfigurationsRequest {
	s.ClientToken = &v
	return s
}

type DeleteStateConfigurationsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteStateConfigurationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteStateConfigurationsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteStateConfigurationsResponseBody) SetRequestId(v string) *DeleteStateConfigurationsResponseBody {
	s.RequestId = &v
	return s
}

type DeleteStateConfigurationsResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteStateConfigurationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteStateConfigurationsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteStateConfigurationsResponse) GoString() string {
	return s.String()
}

func (s *DeleteStateConfigurationsResponse) SetHeaders(v map[string]*string) *DeleteStateConfigurationsResponse {
	s.Headers = v
	return s
}

func (s *DeleteStateConfigurationsResponse) SetBody(v *DeleteStateConfigurationsResponseBody) *DeleteStateConfigurationsResponse {
	s.Body = v
	return s
}

type DeleteTemplateRequest struct {
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TemplateName         *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	AutoDeleteExecutions *bool   `json:"AutoDeleteExecutions,omitempty" xml:"AutoDeleteExecutions,omitempty"`
}

func (s DeleteTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteTemplateRequest) SetRegionId(v string) *DeleteTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteTemplateRequest) SetTemplateName(v string) *DeleteTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *DeleteTemplateRequest) SetAutoDeleteExecutions(v bool) *DeleteTemplateRequest {
	s.AutoDeleteExecutions = &v
	return s
}

type DeleteTemplateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTemplateResponseBody) SetRequestId(v string) *DeleteTemplateResponseBody {
	s.RequestId = &v
	return s
}

type DeleteTemplateResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteTemplateResponse) SetHeaders(v map[string]*string) *DeleteTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteTemplateResponse) SetBody(v *DeleteTemplateResponseBody) *DeleteTemplateResponse {
	s.Body = v
	return s
}

type DeleteTemplatesRequest struct {
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TemplateNames        *string `json:"TemplateNames,omitempty" xml:"TemplateNames,omitempty"`
	AutoDeleteExecutions *bool   `json:"AutoDeleteExecutions,omitempty" xml:"AutoDeleteExecutions,omitempty"`
}

func (s DeleteTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTemplatesRequest) GoString() string {
	return s.String()
}

func (s *DeleteTemplatesRequest) SetRegionId(v string) *DeleteTemplatesRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteTemplatesRequest) SetTemplateNames(v string) *DeleteTemplatesRequest {
	s.TemplateNames = &v
	return s
}

func (s *DeleteTemplatesRequest) SetAutoDeleteExecutions(v bool) *DeleteTemplatesRequest {
	s.AutoDeleteExecutions = &v
	return s
}

type DeleteTemplatesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTemplatesResponseBody) SetRequestId(v string) *DeleteTemplatesResponseBody {
	s.RequestId = &v
	return s
}

type DeleteTemplatesResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTemplatesResponse) GoString() string {
	return s.String()
}

func (s *DeleteTemplatesResponse) SetHeaders(v map[string]*string) *DeleteTemplatesResponse {
	s.Headers = v
	return s
}

func (s *DeleteTemplatesResponse) SetBody(v *DeleteTemplatesResponseBody) *DeleteTemplatesResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetRegionId(v string) *DescribeRegionsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionsRequest) SetAcceptLanguage(v string) *DescribeRegionsRequest {
	s.AcceptLanguage = &v
	return s
}

type DescribeRegionsResponseBody struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Regions   []*DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetRegions(v []*DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

type DescribeRegionsResponseBodyRegions struct {
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	LocalName      *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionEndpoint(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionEndpoint = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetLocalName(v string) *DescribeRegionsResponseBodyRegions {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionId(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionId = &v
	return s
}

type DescribeRegionsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponse) SetHeaders(v map[string]*string) *DescribeRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type GenerateExecutionPolicyRequest struct {
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TemplateName    *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
}

func (s GenerateExecutionPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateExecutionPolicyRequest) GoString() string {
	return s.String()
}

func (s *GenerateExecutionPolicyRequest) SetRegionId(v string) *GenerateExecutionPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *GenerateExecutionPolicyRequest) SetTemplateName(v string) *GenerateExecutionPolicyRequest {
	s.TemplateName = &v
	return s
}

func (s *GenerateExecutionPolicyRequest) SetTemplateVersion(v string) *GenerateExecutionPolicyRequest {
	s.TemplateVersion = &v
	return s
}

type GenerateExecutionPolicyResponseBody struct {
	Policy    *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateExecutionPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateExecutionPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateExecutionPolicyResponseBody) SetPolicy(v string) *GenerateExecutionPolicyResponseBody {
	s.Policy = &v
	return s
}

func (s *GenerateExecutionPolicyResponseBody) SetRequestId(v string) *GenerateExecutionPolicyResponseBody {
	s.RequestId = &v
	return s
}

type GenerateExecutionPolicyResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GenerateExecutionPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GenerateExecutionPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateExecutionPolicyResponse) GoString() string {
	return s.String()
}

func (s *GenerateExecutionPolicyResponse) SetHeaders(v map[string]*string) *GenerateExecutionPolicyResponse {
	s.Headers = v
	return s
}

func (s *GenerateExecutionPolicyResponse) SetBody(v *GenerateExecutionPolicyResponseBody) *GenerateExecutionPolicyResponse {
	s.Body = v
	return s
}

type GetApplicationRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationRequest) GoString() string {
	return s.String()
}

func (s *GetApplicationRequest) SetRegionId(v string) *GetApplicationRequest {
	s.RegionId = &v
	return s
}

func (s *GetApplicationRequest) SetName(v string) *GetApplicationRequest {
	s.Name = &v
	return s
}

type GetApplicationResponseBody struct {
	RequestId   *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Application *GetApplicationResponseBodyApplication `json:"Application,omitempty" xml:"Application,omitempty" type:"Struct"`
}

func (s GetApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBody) SetRequestId(v string) *GetApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetApplicationResponseBody) SetApplication(v *GetApplicationResponseBodyApplication) *GetApplicationResponseBody {
	s.Application = v
	return s
}

type GetApplicationResponseBodyApplication struct {
	IsSystem         *string                                                `json:"IsSystem,omitempty" xml:"IsSystem,omitempty"`
	Description      *string                                                `json:"Description,omitempty" xml:"Description,omitempty"`
	UpdatedDate      *string                                                `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
	ResourceGroupId  *string                                                `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	CreatedDate      *string                                                `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	Name             *string                                                `json:"Name,omitempty" xml:"Name,omitempty"`
	CloudMonitorRule *GetApplicationResponseBodyApplicationCloudMonitorRule `json:"CloudMonitorRule,omitempty" xml:"CloudMonitorRule,omitempty" type:"Struct"`
}

func (s GetApplicationResponseBodyApplication) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationResponseBodyApplication) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyApplication) SetIsSystem(v string) *GetApplicationResponseBodyApplication {
	s.IsSystem = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetDescription(v string) *GetApplicationResponseBodyApplication {
	s.Description = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetUpdatedDate(v string) *GetApplicationResponseBodyApplication {
	s.UpdatedDate = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetResourceGroupId(v string) *GetApplicationResponseBodyApplication {
	s.ResourceGroupId = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetCreatedDate(v string) *GetApplicationResponseBodyApplication {
	s.CreatedDate = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetName(v string) *GetApplicationResponseBodyApplication {
	s.Name = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetCloudMonitorRule(v *GetApplicationResponseBodyApplicationCloudMonitorRule) *GetApplicationResponseBodyApplication {
	s.CloudMonitorRule = v
	return s
}

type GetApplicationResponseBodyApplicationCloudMonitorRule struct {
	EnableSubscribeEvent *bool     `json:"EnableSubscribeEvent,omitempty" xml:"EnableSubscribeEvent,omitempty"`
	EnableInstallAgent   *bool     `json:"EnableInstallAgent,omitempty" xml:"EnableInstallAgent,omitempty"`
	Enabled              *bool     `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	ContactGroupList     []*string `json:"ContactGroupList,omitempty" xml:"ContactGroupList,omitempty" type:"Repeated"`
	TemplateIdList       []*int32  `json:"TemplateIdList,omitempty" xml:"TemplateIdList,omitempty" type:"Repeated"`
}

func (s GetApplicationResponseBodyApplicationCloudMonitorRule) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationResponseBodyApplicationCloudMonitorRule) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyApplicationCloudMonitorRule) SetEnableSubscribeEvent(v bool) *GetApplicationResponseBodyApplicationCloudMonitorRule {
	s.EnableSubscribeEvent = &v
	return s
}

func (s *GetApplicationResponseBodyApplicationCloudMonitorRule) SetEnableInstallAgent(v bool) *GetApplicationResponseBodyApplicationCloudMonitorRule {
	s.EnableInstallAgent = &v
	return s
}

func (s *GetApplicationResponseBodyApplicationCloudMonitorRule) SetEnabled(v bool) *GetApplicationResponseBodyApplicationCloudMonitorRule {
	s.Enabled = &v
	return s
}

func (s *GetApplicationResponseBodyApplicationCloudMonitorRule) SetContactGroupList(v []*string) *GetApplicationResponseBodyApplicationCloudMonitorRule {
	s.ContactGroupList = v
	return s
}

func (s *GetApplicationResponseBodyApplicationCloudMonitorRule) SetTemplateIdList(v []*int32) *GetApplicationResponseBodyApplicationCloudMonitorRule {
	s.TemplateIdList = v
	return s
}

type GetApplicationResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationResponse) GoString() string {
	return s.String()
}

func (s *GetApplicationResponse) SetHeaders(v map[string]*string) *GetApplicationResponse {
	s.Headers = v
	return s
}

func (s *GetApplicationResponse) SetBody(v *GetApplicationResponseBody) *GetApplicationResponse {
	s.Body = v
	return s
}

type GetApplicationGroupRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetApplicationGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationGroupRequest) GoString() string {
	return s.String()
}

func (s *GetApplicationGroupRequest) SetRegionId(v string) *GetApplicationGroupRequest {
	s.RegionId = &v
	return s
}

func (s *GetApplicationGroupRequest) SetName(v string) *GetApplicationGroupRequest {
	s.Name = &v
	return s
}

type GetApplicationGroupResponseBody struct {
	RequestId        *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ApplicationGroup *GetApplicationGroupResponseBodyApplicationGroup `json:"ApplicationGroup,omitempty" xml:"ApplicationGroup,omitempty" type:"Struct"`
}

func (s GetApplicationGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetApplicationGroupResponseBody) SetRequestId(v string) *GetApplicationGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetApplicationGroupResponseBody) SetApplicationGroup(v *GetApplicationGroupResponseBodyApplicationGroup) *GetApplicationGroupResponseBody {
	s.ApplicationGroup = v
	return s
}

type GetApplicationGroupResponseBodyApplicationGroup struct {
	DeployRegionId  *string `json:"DeployRegionId,omitempty" xml:"DeployRegionId,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	UpdatedDate     *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
	CreatedDate     *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Environment     *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	CreateType      *string `json:"CreateType,omitempty" xml:"CreateType,omitempty"`
	ScalingGroupId  *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	ImportClusterId *string `json:"ImportClusterId,omitempty" xml:"ImportClusterId,omitempty"`
}

func (s GetApplicationGroupResponseBodyApplicationGroup) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationGroupResponseBodyApplicationGroup) GoString() string {
	return s.String()
}

func (s *GetApplicationGroupResponseBodyApplicationGroup) SetDeployRegionId(v string) *GetApplicationGroupResponseBodyApplicationGroup {
	s.DeployRegionId = &v
	return s
}

func (s *GetApplicationGroupResponseBodyApplicationGroup) SetDescription(v string) *GetApplicationGroupResponseBodyApplicationGroup {
	s.Description = &v
	return s
}

func (s *GetApplicationGroupResponseBodyApplicationGroup) SetUpdatedDate(v string) *GetApplicationGroupResponseBodyApplicationGroup {
	s.UpdatedDate = &v
	return s
}

func (s *GetApplicationGroupResponseBodyApplicationGroup) SetCreatedDate(v string) *GetApplicationGroupResponseBodyApplicationGroup {
	s.CreatedDate = &v
	return s
}

func (s *GetApplicationGroupResponseBodyApplicationGroup) SetApplicationName(v string) *GetApplicationGroupResponseBodyApplicationGroup {
	s.ApplicationName = &v
	return s
}

func (s *GetApplicationGroupResponseBodyApplicationGroup) SetName(v string) *GetApplicationGroupResponseBodyApplicationGroup {
	s.Name = &v
	return s
}

func (s *GetApplicationGroupResponseBodyApplicationGroup) SetEnvironment(v string) *GetApplicationGroupResponseBodyApplicationGroup {
	s.Environment = &v
	return s
}

func (s *GetApplicationGroupResponseBodyApplicationGroup) SetCreateType(v string) *GetApplicationGroupResponseBodyApplicationGroup {
	s.CreateType = &v
	return s
}

func (s *GetApplicationGroupResponseBodyApplicationGroup) SetScalingGroupId(v string) *GetApplicationGroupResponseBodyApplicationGroup {
	s.ScalingGroupId = &v
	return s
}

func (s *GetApplicationGroupResponseBodyApplicationGroup) SetImportClusterId(v string) *GetApplicationGroupResponseBodyApplicationGroup {
	s.ImportClusterId = &v
	return s
}

type GetApplicationGroupResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetApplicationGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetApplicationGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationGroupResponse) GoString() string {
	return s.String()
}

func (s *GetApplicationGroupResponse) SetHeaders(v map[string]*string) *GetApplicationGroupResponse {
	s.Headers = v
	return s
}

func (s *GetApplicationGroupResponse) SetBody(v *GetApplicationGroupResponseBody) *GetApplicationGroupResponse {
	s.Body = v
	return s
}

type GetExecutionTemplateRequest struct {
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ExecutionId *string `json:"ExecutionId,omitempty" xml:"ExecutionId,omitempty"`
}

func (s GetExecutionTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s GetExecutionTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetExecutionTemplateRequest) SetRegionId(v string) *GetExecutionTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *GetExecutionTemplateRequest) SetExecutionId(v string) *GetExecutionTemplateRequest {
	s.ExecutionId = &v
	return s
}

type GetExecutionTemplateResponseBody struct {
	Content   *string                                   `json:"Content,omitempty" xml:"Content,omitempty"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Template  *GetExecutionTemplateResponseBodyTemplate `json:"Template,omitempty" xml:"Template,omitempty" type:"Struct"`
}

func (s GetExecutionTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetExecutionTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetExecutionTemplateResponseBody) SetContent(v string) *GetExecutionTemplateResponseBody {
	s.Content = &v
	return s
}

func (s *GetExecutionTemplateResponseBody) SetRequestId(v string) *GetExecutionTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetExecutionTemplateResponseBody) SetTemplate(v *GetExecutionTemplateResponseBodyTemplate) *GetExecutionTemplateResponseBody {
	s.Template = v
	return s
}

type GetExecutionTemplateResponseBodyTemplate struct {
	Hash            *string                `json:"Hash,omitempty" xml:"Hash,omitempty"`
	UpdatedDate     *string                `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
	UpdatedBy       *string                `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	Tags            map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	TemplateName    *string                `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateVersion *string                `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	TemplateFormat  *string                `json:"TemplateFormat,omitempty" xml:"TemplateFormat,omitempty"`
	Description     *string                `json:"Description,omitempty" xml:"Description,omitempty"`
	CreatedBy       *string                `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	CreatedDate     *string                `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	TemplateId      *string                `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	ShareType       *string                `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
}

func (s GetExecutionTemplateResponseBodyTemplate) String() string {
	return tea.Prettify(s)
}

func (s GetExecutionTemplateResponseBodyTemplate) GoString() string {
	return s.String()
}

func (s *GetExecutionTemplateResponseBodyTemplate) SetHash(v string) *GetExecutionTemplateResponseBodyTemplate {
	s.Hash = &v
	return s
}

func (s *GetExecutionTemplateResponseBodyTemplate) SetUpdatedDate(v string) *GetExecutionTemplateResponseBodyTemplate {
	s.UpdatedDate = &v
	return s
}

func (s *GetExecutionTemplateResponseBodyTemplate) SetUpdatedBy(v string) *GetExecutionTemplateResponseBodyTemplate {
	s.UpdatedBy = &v
	return s
}

func (s *GetExecutionTemplateResponseBodyTemplate) SetTags(v map[string]interface{}) *GetExecutionTemplateResponseBodyTemplate {
	s.Tags = v
	return s
}

func (s *GetExecutionTemplateResponseBodyTemplate) SetTemplateName(v string) *GetExecutionTemplateResponseBodyTemplate {
	s.TemplateName = &v
	return s
}

func (s *GetExecutionTemplateResponseBodyTemplate) SetTemplateVersion(v string) *GetExecutionTemplateResponseBodyTemplate {
	s.TemplateVersion = &v
	return s
}

func (s *GetExecutionTemplateResponseBodyTemplate) SetTemplateFormat(v string) *GetExecutionTemplateResponseBodyTemplate {
	s.TemplateFormat = &v
	return s
}

func (s *GetExecutionTemplateResponseBodyTemplate) SetDescription(v string) *GetExecutionTemplateResponseBodyTemplate {
	s.Description = &v
	return s
}

func (s *GetExecutionTemplateResponseBodyTemplate) SetCreatedBy(v string) *GetExecutionTemplateResponseBodyTemplate {
	s.CreatedBy = &v
	return s
}

func (s *GetExecutionTemplateResponseBodyTemplate) SetCreatedDate(v string) *GetExecutionTemplateResponseBodyTemplate {
	s.CreatedDate = &v
	return s
}

func (s *GetExecutionTemplateResponseBodyTemplate) SetTemplateId(v string) *GetExecutionTemplateResponseBodyTemplate {
	s.TemplateId = &v
	return s
}

func (s *GetExecutionTemplateResponseBodyTemplate) SetShareType(v string) *GetExecutionTemplateResponseBodyTemplate {
	s.ShareType = &v
	return s
}

type GetExecutionTemplateResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetExecutionTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetExecutionTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s GetExecutionTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetExecutionTemplateResponse) SetHeaders(v map[string]*string) *GetExecutionTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetExecutionTemplateResponse) SetBody(v *GetExecutionTemplateResponseBody) *GetExecutionTemplateResponse {
	s.Body = v
	return s
}

type GetInventorySchemaRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Aggregator *bool   `json:"Aggregator,omitempty" xml:"Aggregator,omitempty"`
	TypeName   *string `json:"TypeName,omitempty" xml:"TypeName,omitempty"`
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s GetInventorySchemaRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInventorySchemaRequest) GoString() string {
	return s.String()
}

func (s *GetInventorySchemaRequest) SetRegionId(v string) *GetInventorySchemaRequest {
	s.RegionId = &v
	return s
}

func (s *GetInventorySchemaRequest) SetAggregator(v bool) *GetInventorySchemaRequest {
	s.Aggregator = &v
	return s
}

func (s *GetInventorySchemaRequest) SetTypeName(v string) *GetInventorySchemaRequest {
	s.TypeName = &v
	return s
}

func (s *GetInventorySchemaRequest) SetMaxResults(v int32) *GetInventorySchemaRequest {
	s.MaxResults = &v
	return s
}

func (s *GetInventorySchemaRequest) SetNextToken(v string) *GetInventorySchemaRequest {
	s.NextToken = &v
	return s
}

type GetInventorySchemaResponseBody struct {
	NextToken  *string                                  `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId  *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	MaxResults *string                                  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Schemas    []*GetInventorySchemaResponseBodySchemas `json:"Schemas,omitempty" xml:"Schemas,omitempty" type:"Repeated"`
}

func (s GetInventorySchemaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInventorySchemaResponseBody) GoString() string {
	return s.String()
}

func (s *GetInventorySchemaResponseBody) SetNextToken(v string) *GetInventorySchemaResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetInventorySchemaResponseBody) SetRequestId(v string) *GetInventorySchemaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInventorySchemaResponseBody) SetMaxResults(v string) *GetInventorySchemaResponseBody {
	s.MaxResults = &v
	return s
}

func (s *GetInventorySchemaResponseBody) SetSchemas(v []*GetInventorySchemaResponseBodySchemas) *GetInventorySchemaResponseBody {
	s.Schemas = v
	return s
}

type GetInventorySchemaResponseBodySchemas struct {
	Version    *string                                            `json:"Version,omitempty" xml:"Version,omitempty"`
	TypeName   *string                                            `json:"TypeName,omitempty" xml:"TypeName,omitempty"`
	Attributes []*GetInventorySchemaResponseBodySchemasAttributes `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
}

func (s GetInventorySchemaResponseBodySchemas) String() string {
	return tea.Prettify(s)
}

func (s GetInventorySchemaResponseBodySchemas) GoString() string {
	return s.String()
}

func (s *GetInventorySchemaResponseBodySchemas) SetVersion(v string) *GetInventorySchemaResponseBodySchemas {
	s.Version = &v
	return s
}

func (s *GetInventorySchemaResponseBodySchemas) SetTypeName(v string) *GetInventorySchemaResponseBodySchemas {
	s.TypeName = &v
	return s
}

func (s *GetInventorySchemaResponseBodySchemas) SetAttributes(v []*GetInventorySchemaResponseBodySchemasAttributes) *GetInventorySchemaResponseBodySchemas {
	s.Attributes = v
	return s
}

type GetInventorySchemaResponseBodySchemasAttributes struct {
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
}

func (s GetInventorySchemaResponseBodySchemasAttributes) String() string {
	return tea.Prettify(s)
}

func (s GetInventorySchemaResponseBodySchemasAttributes) GoString() string {
	return s.String()
}

func (s *GetInventorySchemaResponseBodySchemasAttributes) SetName(v string) *GetInventorySchemaResponseBodySchemasAttributes {
	s.Name = &v
	return s
}

func (s *GetInventorySchemaResponseBodySchemasAttributes) SetDataType(v string) *GetInventorySchemaResponseBodySchemasAttributes {
	s.DataType = &v
	return s
}

type GetInventorySchemaResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetInventorySchemaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetInventorySchemaResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInventorySchemaResponse) GoString() string {
	return s.String()
}

func (s *GetInventorySchemaResponse) SetHeaders(v map[string]*string) *GetInventorySchemaResponse {
	s.Headers = v
	return s
}

func (s *GetInventorySchemaResponse) SetBody(v *GetInventorySchemaResponseBody) *GetInventorySchemaResponse {
	s.Body = v
	return s
}

type GetParameterRequest struct {
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ParameterVersion *int32  `json:"ParameterVersion,omitempty" xml:"ParameterVersion,omitempty"`
	ResourceGroupId  *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s GetParameterRequest) String() string {
	return tea.Prettify(s)
}

func (s GetParameterRequest) GoString() string {
	return s.String()
}

func (s *GetParameterRequest) SetRegionId(v string) *GetParameterRequest {
	s.RegionId = &v
	return s
}

func (s *GetParameterRequest) SetName(v string) *GetParameterRequest {
	s.Name = &v
	return s
}

func (s *GetParameterRequest) SetParameterVersion(v int32) *GetParameterRequest {
	s.ParameterVersion = &v
	return s
}

func (s *GetParameterRequest) SetResourceGroupId(v string) *GetParameterRequest {
	s.ResourceGroupId = &v
	return s
}

type GetParameterResponseBody struct {
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Parameter *GetParameterResponseBodyParameter `json:"Parameter,omitempty" xml:"Parameter,omitempty" type:"Struct"`
}

func (s GetParameterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetParameterResponseBody) GoString() string {
	return s.String()
}

func (s *GetParameterResponseBody) SetRequestId(v string) *GetParameterResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetParameterResponseBody) SetParameter(v *GetParameterResponseBodyParameter) *GetParameterResponseBody {
	s.Parameter = v
	return s
}

type GetParameterResponseBodyParameter struct {
	Type             *string                `json:"Type,omitempty" xml:"Type,omitempty"`
	UpdatedDate      *string                `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
	UpdatedBy        *string                `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	Tags             map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	Value            *string                `json:"Value,omitempty" xml:"Value,omitempty"`
	Description      *string                `json:"Description,omitempty" xml:"Description,omitempty"`
	Constraints      *string                `json:"Constraints,omitempty" xml:"Constraints,omitempty"`
	ResourceGroupId  *string                `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	CreatedBy        *string                `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	CreatedDate      *string                `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	ParameterVersion *int32                 `json:"ParameterVersion,omitempty" xml:"ParameterVersion,omitempty"`
	Name             *string                `json:"Name,omitempty" xml:"Name,omitempty"`
	Id               *string                `json:"Id,omitempty" xml:"Id,omitempty"`
	ShareType        *string                `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
}

func (s GetParameterResponseBodyParameter) String() string {
	return tea.Prettify(s)
}

func (s GetParameterResponseBodyParameter) GoString() string {
	return s.String()
}

func (s *GetParameterResponseBodyParameter) SetType(v string) *GetParameterResponseBodyParameter {
	s.Type = &v
	return s
}

func (s *GetParameterResponseBodyParameter) SetUpdatedDate(v string) *GetParameterResponseBodyParameter {
	s.UpdatedDate = &v
	return s
}

func (s *GetParameterResponseBodyParameter) SetUpdatedBy(v string) *GetParameterResponseBodyParameter {
	s.UpdatedBy = &v
	return s
}

func (s *GetParameterResponseBodyParameter) SetTags(v map[string]interface{}) *GetParameterResponseBodyParameter {
	s.Tags = v
	return s
}

func (s *GetParameterResponseBodyParameter) SetValue(v string) *GetParameterResponseBodyParameter {
	s.Value = &v
	return s
}

func (s *GetParameterResponseBodyParameter) SetDescription(v string) *GetParameterResponseBodyParameter {
	s.Description = &v
	return s
}

func (s *GetParameterResponseBodyParameter) SetConstraints(v string) *GetParameterResponseBodyParameter {
	s.Constraints = &v
	return s
}

func (s *GetParameterResponseBodyParameter) SetResourceGroupId(v string) *GetParameterResponseBodyParameter {
	s.ResourceGroupId = &v
	return s
}

func (s *GetParameterResponseBodyParameter) SetCreatedBy(v string) *GetParameterResponseBodyParameter {
	s.CreatedBy = &v
	return s
}

func (s *GetParameterResponseBodyParameter) SetCreatedDate(v string) *GetParameterResponseBodyParameter {
	s.CreatedDate = &v
	return s
}

func (s *GetParameterResponseBodyParameter) SetParameterVersion(v int32) *GetParameterResponseBodyParameter {
	s.ParameterVersion = &v
	return s
}

func (s *GetParameterResponseBodyParameter) SetName(v string) *GetParameterResponseBodyParameter {
	s.Name = &v
	return s
}

func (s *GetParameterResponseBodyParameter) SetId(v string) *GetParameterResponseBodyParameter {
	s.Id = &v
	return s
}

func (s *GetParameterResponseBodyParameter) SetShareType(v string) *GetParameterResponseBodyParameter {
	s.ShareType = &v
	return s
}

type GetParameterResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetParameterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetParameterResponse) String() string {
	return tea.Prettify(s)
}

func (s GetParameterResponse) GoString() string {
	return s.String()
}

func (s *GetParameterResponse) SetHeaders(v map[string]*string) *GetParameterResponse {
	s.Headers = v
	return s
}

func (s *GetParameterResponse) SetBody(v *GetParameterResponseBody) *GetParameterResponse {
	s.Body = v
	return s
}

type GetParametersRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Names    *string `json:"Names,omitempty" xml:"Names,omitempty"`
}

func (s GetParametersRequest) String() string {
	return tea.Prettify(s)
}

func (s GetParametersRequest) GoString() string {
	return s.String()
}

func (s *GetParametersRequest) SetRegionId(v string) *GetParametersRequest {
	s.RegionId = &v
	return s
}

func (s *GetParametersRequest) SetNames(v string) *GetParametersRequest {
	s.Names = &v
	return s
}

type GetParametersResponseBody struct {
	RequestId         *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	InvalidParameters []*string                              `json:"InvalidParameters,omitempty" xml:"InvalidParameters,omitempty" type:"Repeated"`
	Parameters        []*GetParametersResponseBodyParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
}

func (s GetParametersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetParametersResponseBody) GoString() string {
	return s.String()
}

func (s *GetParametersResponseBody) SetRequestId(v string) *GetParametersResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetParametersResponseBody) SetInvalidParameters(v []*string) *GetParametersResponseBody {
	s.InvalidParameters = v
	return s
}

func (s *GetParametersResponseBody) SetParameters(v []*GetParametersResponseBodyParameters) *GetParametersResponseBody {
	s.Parameters = v
	return s
}

type GetParametersResponseBodyParameters struct {
	Type             *string                `json:"Type,omitempty" xml:"Type,omitempty"`
	UpdatedDate      *string                `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
	UpdatedBy        *string                `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	Tags             map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	Value            *string                `json:"Value,omitempty" xml:"Value,omitempty"`
	Description      *string                `json:"Description,omitempty" xml:"Description,omitempty"`
	Constraints      *string                `json:"Constraints,omitempty" xml:"Constraints,omitempty"`
	ResourceGroupId  *string                `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	CreatedBy        *string                `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	CreatedDate      *string                `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	ParameterVersion *int32                 `json:"ParameterVersion,omitempty" xml:"ParameterVersion,omitempty"`
	Name             *string                `json:"Name,omitempty" xml:"Name,omitempty"`
	Id               *string                `json:"Id,omitempty" xml:"Id,omitempty"`
	ShareType        *string                `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
}

func (s GetParametersResponseBodyParameters) String() string {
	return tea.Prettify(s)
}

func (s GetParametersResponseBodyParameters) GoString() string {
	return s.String()
}

func (s *GetParametersResponseBodyParameters) SetType(v string) *GetParametersResponseBodyParameters {
	s.Type = &v
	return s
}

func (s *GetParametersResponseBodyParameters) SetUpdatedDate(v string) *GetParametersResponseBodyParameters {
	s.UpdatedDate = &v
	return s
}

func (s *GetParametersResponseBodyParameters) SetUpdatedBy(v string) *GetParametersResponseBodyParameters {
	s.UpdatedBy = &v
	return s
}

func (s *GetParametersResponseBodyParameters) SetTags(v map[string]interface{}) *GetParametersResponseBodyParameters {
	s.Tags = v
	return s
}

func (s *GetParametersResponseBodyParameters) SetValue(v string) *GetParametersResponseBodyParameters {
	s.Value = &v
	return s
}

func (s *GetParametersResponseBodyParameters) SetDescription(v string) *GetParametersResponseBodyParameters {
	s.Description = &v
	return s
}

func (s *GetParametersResponseBodyParameters) SetConstraints(v string) *GetParametersResponseBodyParameters {
	s.Constraints = &v
	return s
}

func (s *GetParametersResponseBodyParameters) SetResourceGroupId(v string) *GetParametersResponseBodyParameters {
	s.ResourceGroupId = &v
	return s
}

func (s *GetParametersResponseBodyParameters) SetCreatedBy(v string) *GetParametersResponseBodyParameters {
	s.CreatedBy = &v
	return s
}

func (s *GetParametersResponseBodyParameters) SetCreatedDate(v string) *GetParametersResponseBodyParameters {
	s.CreatedDate = &v
	return s
}

func (s *GetParametersResponseBodyParameters) SetParameterVersion(v int32) *GetParametersResponseBodyParameters {
	s.ParameterVersion = &v
	return s
}

func (s *GetParametersResponseBodyParameters) SetName(v string) *GetParametersResponseBodyParameters {
	s.Name = &v
	return s
}

func (s *GetParametersResponseBodyParameters) SetId(v string) *GetParametersResponseBodyParameters {
	s.Id = &v
	return s
}

func (s *GetParametersResponseBodyParameters) SetShareType(v string) *GetParametersResponseBodyParameters {
	s.ShareType = &v
	return s
}

type GetParametersResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetParametersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetParametersResponse) String() string {
	return tea.Prettify(s)
}

func (s GetParametersResponse) GoString() string {
	return s.String()
}

func (s *GetParametersResponse) SetHeaders(v map[string]*string) *GetParametersResponse {
	s.Headers = v
	return s
}

func (s *GetParametersResponse) SetBody(v *GetParametersResponseBody) *GetParametersResponse {
	s.Body = v
	return s
}

type GetParametersByPathRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Path       *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Recursive  *bool   `json:"Recursive,omitempty" xml:"Recursive,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
}

func (s GetParametersByPathRequest) String() string {
	return tea.Prettify(s)
}

func (s GetParametersByPathRequest) GoString() string {
	return s.String()
}

func (s *GetParametersByPathRequest) SetRegionId(v string) *GetParametersByPathRequest {
	s.RegionId = &v
	return s
}

func (s *GetParametersByPathRequest) SetPath(v string) *GetParametersByPathRequest {
	s.Path = &v
	return s
}

func (s *GetParametersByPathRequest) SetRecursive(v bool) *GetParametersByPathRequest {
	s.Recursive = &v
	return s
}

func (s *GetParametersByPathRequest) SetNextToken(v string) *GetParametersByPathRequest {
	s.NextToken = &v
	return s
}

func (s *GetParametersByPathRequest) SetMaxResults(v int32) *GetParametersByPathRequest {
	s.MaxResults = &v
	return s
}

type GetParametersByPathResponseBody struct {
	NextToken  *string                                      `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId  *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                       `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	MaxResults *int32                                       `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Parameters []*GetParametersByPathResponseBodyParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
}

func (s GetParametersByPathResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetParametersByPathResponseBody) GoString() string {
	return s.String()
}

func (s *GetParametersByPathResponseBody) SetNextToken(v string) *GetParametersByPathResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetParametersByPathResponseBody) SetRequestId(v string) *GetParametersByPathResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetParametersByPathResponseBody) SetTotalCount(v int32) *GetParametersByPathResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetParametersByPathResponseBody) SetMaxResults(v int32) *GetParametersByPathResponseBody {
	s.MaxResults = &v
	return s
}

func (s *GetParametersByPathResponseBody) SetParameters(v []*GetParametersByPathResponseBodyParameters) *GetParametersByPathResponseBody {
	s.Parameters = v
	return s
}

type GetParametersByPathResponseBodyParameters struct {
	Type             *string `json:"Type,omitempty" xml:"Type,omitempty"`
	UpdatedDate      *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
	UpdatedBy        *string `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	Value            *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Constraints      *string `json:"Constraints,omitempty" xml:"Constraints,omitempty"`
	CreatedBy        *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	CreatedDate      *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	ParameterVersion *int32  `json:"ParameterVersion,omitempty" xml:"ParameterVersion,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id               *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ShareType        *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
}

func (s GetParametersByPathResponseBodyParameters) String() string {
	return tea.Prettify(s)
}

func (s GetParametersByPathResponseBodyParameters) GoString() string {
	return s.String()
}

func (s *GetParametersByPathResponseBodyParameters) SetType(v string) *GetParametersByPathResponseBodyParameters {
	s.Type = &v
	return s
}

func (s *GetParametersByPathResponseBodyParameters) SetUpdatedDate(v string) *GetParametersByPathResponseBodyParameters {
	s.UpdatedDate = &v
	return s
}

func (s *GetParametersByPathResponseBodyParameters) SetUpdatedBy(v string) *GetParametersByPathResponseBodyParameters {
	s.UpdatedBy = &v
	return s
}

func (s *GetParametersByPathResponseBodyParameters) SetValue(v string) *GetParametersByPathResponseBodyParameters {
	s.Value = &v
	return s
}

func (s *GetParametersByPathResponseBodyParameters) SetDescription(v string) *GetParametersByPathResponseBodyParameters {
	s.Description = &v
	return s
}

func (s *GetParametersByPathResponseBodyParameters) SetConstraints(v string) *GetParametersByPathResponseBodyParameters {
	s.Constraints = &v
	return s
}

func (s *GetParametersByPathResponseBodyParameters) SetCreatedBy(v string) *GetParametersByPathResponseBodyParameters {
	s.CreatedBy = &v
	return s
}

func (s *GetParametersByPathResponseBodyParameters) SetCreatedDate(v string) *GetParametersByPathResponseBodyParameters {
	s.CreatedDate = &v
	return s
}

func (s *GetParametersByPathResponseBodyParameters) SetParameterVersion(v int32) *GetParametersByPathResponseBodyParameters {
	s.ParameterVersion = &v
	return s
}

func (s *GetParametersByPathResponseBodyParameters) SetName(v string) *GetParametersByPathResponseBodyParameters {
	s.Name = &v
	return s
}

func (s *GetParametersByPathResponseBodyParameters) SetId(v string) *GetParametersByPathResponseBodyParameters {
	s.Id = &v
	return s
}

func (s *GetParametersByPathResponseBodyParameters) SetShareType(v string) *GetParametersByPathResponseBodyParameters {
	s.ShareType = &v
	return s
}

type GetParametersByPathResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetParametersByPathResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetParametersByPathResponse) String() string {
	return tea.Prettify(s)
}

func (s GetParametersByPathResponse) GoString() string {
	return s.String()
}

func (s *GetParametersByPathResponse) SetHeaders(v map[string]*string) *GetParametersByPathResponse {
	s.Headers = v
	return s
}

func (s *GetParametersByPathResponse) SetBody(v *GetParametersByPathResponseBody) *GetParametersByPathResponse {
	s.Body = v
	return s
}

type GetPatchBaselineRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetPatchBaselineRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPatchBaselineRequest) GoString() string {
	return s.String()
}

func (s *GetPatchBaselineRequest) SetRegionId(v string) *GetPatchBaselineRequest {
	s.RegionId = &v
	return s
}

func (s *GetPatchBaselineRequest) SetName(v string) *GetPatchBaselineRequest {
	s.Name = &v
	return s
}

type GetPatchBaselineResponseBody struct {
	RequestId     *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PatchBaseline *GetPatchBaselineResponseBodyPatchBaseline `json:"PatchBaseline,omitempty" xml:"PatchBaseline,omitempty" type:"Struct"`
}

func (s GetPatchBaselineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPatchBaselineResponseBody) GoString() string {
	return s.String()
}

func (s *GetPatchBaselineResponseBody) SetRequestId(v string) *GetPatchBaselineResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPatchBaselineResponseBody) SetPatchBaseline(v *GetPatchBaselineResponseBodyPatchBaseline) *GetPatchBaselineResponseBody {
	s.PatchBaseline = v
	return s
}

type GetPatchBaselineResponseBodyPatchBaseline struct {
	OperationSystem *string `json:"OperationSystem,omitempty" xml:"OperationSystem,omitempty"`
	IsDefault       *bool   `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	UpdatedDate     *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
	UpdatedBy       *string `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	CreatedBy       *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	CreatedDate     *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ApprovalRules   *string `json:"ApprovalRules,omitempty" xml:"ApprovalRules,omitempty"`
	Id              *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ShareType       *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
}

func (s GetPatchBaselineResponseBodyPatchBaseline) String() string {
	return tea.Prettify(s)
}

func (s GetPatchBaselineResponseBodyPatchBaseline) GoString() string {
	return s.String()
}

func (s *GetPatchBaselineResponseBodyPatchBaseline) SetOperationSystem(v string) *GetPatchBaselineResponseBodyPatchBaseline {
	s.OperationSystem = &v
	return s
}

func (s *GetPatchBaselineResponseBodyPatchBaseline) SetIsDefault(v bool) *GetPatchBaselineResponseBodyPatchBaseline {
	s.IsDefault = &v
	return s
}

func (s *GetPatchBaselineResponseBodyPatchBaseline) SetDescription(v string) *GetPatchBaselineResponseBodyPatchBaseline {
	s.Description = &v
	return s
}

func (s *GetPatchBaselineResponseBodyPatchBaseline) SetUpdatedDate(v string) *GetPatchBaselineResponseBodyPatchBaseline {
	s.UpdatedDate = &v
	return s
}

func (s *GetPatchBaselineResponseBodyPatchBaseline) SetUpdatedBy(v string) *GetPatchBaselineResponseBodyPatchBaseline {
	s.UpdatedBy = &v
	return s
}

func (s *GetPatchBaselineResponseBodyPatchBaseline) SetCreatedBy(v string) *GetPatchBaselineResponseBodyPatchBaseline {
	s.CreatedBy = &v
	return s
}

func (s *GetPatchBaselineResponseBodyPatchBaseline) SetCreatedDate(v string) *GetPatchBaselineResponseBodyPatchBaseline {
	s.CreatedDate = &v
	return s
}

func (s *GetPatchBaselineResponseBodyPatchBaseline) SetName(v string) *GetPatchBaselineResponseBodyPatchBaseline {
	s.Name = &v
	return s
}

func (s *GetPatchBaselineResponseBodyPatchBaseline) SetApprovalRules(v string) *GetPatchBaselineResponseBodyPatchBaseline {
	s.ApprovalRules = &v
	return s
}

func (s *GetPatchBaselineResponseBodyPatchBaseline) SetId(v string) *GetPatchBaselineResponseBodyPatchBaseline {
	s.Id = &v
	return s
}

func (s *GetPatchBaselineResponseBodyPatchBaseline) SetShareType(v string) *GetPatchBaselineResponseBodyPatchBaseline {
	s.ShareType = &v
	return s
}

type GetPatchBaselineResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetPatchBaselineResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPatchBaselineResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPatchBaselineResponse) GoString() string {
	return s.String()
}

func (s *GetPatchBaselineResponse) SetHeaders(v map[string]*string) *GetPatchBaselineResponse {
	s.Headers = v
	return s
}

func (s *GetPatchBaselineResponse) SetBody(v *GetPatchBaselineResponseBody) *GetPatchBaselineResponse {
	s.Body = v
	return s
}

type GetSecretParameterRequest struct {
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ParameterVersion *int32  `json:"ParameterVersion,omitempty" xml:"ParameterVersion,omitempty"`
	WithDecryption   *bool   `json:"WithDecryption,omitempty" xml:"WithDecryption,omitempty"`
}

func (s GetSecretParameterRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSecretParameterRequest) GoString() string {
	return s.String()
}

func (s *GetSecretParameterRequest) SetRegionId(v string) *GetSecretParameterRequest {
	s.RegionId = &v
	return s
}

func (s *GetSecretParameterRequest) SetName(v string) *GetSecretParameterRequest {
	s.Name = &v
	return s
}

func (s *GetSecretParameterRequest) SetParameterVersion(v int32) *GetSecretParameterRequest {
	s.ParameterVersion = &v
	return s
}

func (s *GetSecretParameterRequest) SetWithDecryption(v bool) *GetSecretParameterRequest {
	s.WithDecryption = &v
	return s
}

type GetSecretParameterResponseBody struct {
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Parameter *GetSecretParameterResponseBodyParameter `json:"Parameter,omitempty" xml:"Parameter,omitempty" type:"Struct"`
}

func (s GetSecretParameterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSecretParameterResponseBody) GoString() string {
	return s.String()
}

func (s *GetSecretParameterResponseBody) SetRequestId(v string) *GetSecretParameterResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSecretParameterResponseBody) SetParameter(v *GetSecretParameterResponseBodyParameter) *GetSecretParameterResponseBody {
	s.Parameter = v
	return s
}

type GetSecretParameterResponseBodyParameter struct {
	Type             *string                `json:"Type,omitempty" xml:"Type,omitempty"`
	UpdatedDate      *string                `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
	UpdatedBy        *string                `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	KeyId            *string                `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	Tags             map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	Value            *string                `json:"Value,omitempty" xml:"Value,omitempty"`
	Description      *string                `json:"Description,omitempty" xml:"Description,omitempty"`
	Constraints      *string                `json:"Constraints,omitempty" xml:"Constraints,omitempty"`
	ResourceGroupId  *string                `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	CreatedBy        *string                `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	CreatedDate      *string                `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	ParameterVersion *int32                 `json:"ParameterVersion,omitempty" xml:"ParameterVersion,omitempty"`
	Name             *string                `json:"Name,omitempty" xml:"Name,omitempty"`
	Id               *string                `json:"Id,omitempty" xml:"Id,omitempty"`
	ShareType        *string                `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
}

func (s GetSecretParameterResponseBodyParameter) String() string {
	return tea.Prettify(s)
}

func (s GetSecretParameterResponseBodyParameter) GoString() string {
	return s.String()
}

func (s *GetSecretParameterResponseBodyParameter) SetType(v string) *GetSecretParameterResponseBodyParameter {
	s.Type = &v
	return s
}

func (s *GetSecretParameterResponseBodyParameter) SetUpdatedDate(v string) *GetSecretParameterResponseBodyParameter {
	s.UpdatedDate = &v
	return s
}

func (s *GetSecretParameterResponseBodyParameter) SetUpdatedBy(v string) *GetSecretParameterResponseBodyParameter {
	s.UpdatedBy = &v
	return s
}

func (s *GetSecretParameterResponseBodyParameter) SetKeyId(v string) *GetSecretParameterResponseBodyParameter {
	s.KeyId = &v
	return s
}

func (s *GetSecretParameterResponseBodyParameter) SetTags(v map[string]interface{}) *GetSecretParameterResponseBodyParameter {
	s.Tags = v
	return s
}

func (s *GetSecretParameterResponseBodyParameter) SetValue(v string) *GetSecretParameterResponseBodyParameter {
	s.Value = &v
	return s
}

func (s *GetSecretParameterResponseBodyParameter) SetDescription(v string) *GetSecretParameterResponseBodyParameter {
	s.Description = &v
	return s
}

func (s *GetSecretParameterResponseBodyParameter) SetConstraints(v string) *GetSecretParameterResponseBodyParameter {
	s.Constraints = &v
	return s
}

func (s *GetSecretParameterResponseBodyParameter) SetResourceGroupId(v string) *GetSecretParameterResponseBodyParameter {
	s.ResourceGroupId = &v
	return s
}

func (s *GetSecretParameterResponseBodyParameter) SetCreatedBy(v string) *GetSecretParameterResponseBodyParameter {
	s.CreatedBy = &v
	return s
}

func (s *GetSecretParameterResponseBodyParameter) SetCreatedDate(v string) *GetSecretParameterResponseBodyParameter {
	s.CreatedDate = &v
	return s
}

func (s *GetSecretParameterResponseBodyParameter) SetParameterVersion(v int32) *GetSecretParameterResponseBodyParameter {
	s.ParameterVersion = &v
	return s
}

func (s *GetSecretParameterResponseBodyParameter) SetName(v string) *GetSecretParameterResponseBodyParameter {
	s.Name = &v
	return s
}

func (s *GetSecretParameterResponseBodyParameter) SetId(v string) *GetSecretParameterResponseBodyParameter {
	s.Id = &v
	return s
}

func (s *GetSecretParameterResponseBodyParameter) SetShareType(v string) *GetSecretParameterResponseBodyParameter {
	s.ShareType = &v
	return s
}

type GetSecretParameterResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSecretParameterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSecretParameterResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSecretParameterResponse) GoString() string {
	return s.String()
}

func (s *GetSecretParameterResponse) SetHeaders(v map[string]*string) *GetSecretParameterResponse {
	s.Headers = v
	return s
}

func (s *GetSecretParameterResponse) SetBody(v *GetSecretParameterResponseBody) *GetSecretParameterResponse {
	s.Body = v
	return s
}

type GetSecretParametersRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Names          *string `json:"Names,omitempty" xml:"Names,omitempty"`
	WithDecryption *bool   `json:"WithDecryption,omitempty" xml:"WithDecryption,omitempty"`
}

func (s GetSecretParametersRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSecretParametersRequest) GoString() string {
	return s.String()
}

func (s *GetSecretParametersRequest) SetRegionId(v string) *GetSecretParametersRequest {
	s.RegionId = &v
	return s
}

func (s *GetSecretParametersRequest) SetNames(v string) *GetSecretParametersRequest {
	s.Names = &v
	return s
}

func (s *GetSecretParametersRequest) SetWithDecryption(v bool) *GetSecretParametersRequest {
	s.WithDecryption = &v
	return s
}

type GetSecretParametersResponseBody struct {
	RequestId         *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	InvalidParameters []*string                                    `json:"InvalidParameters,omitempty" xml:"InvalidParameters,omitempty" type:"Repeated"`
	Parameters        []*GetSecretParametersResponseBodyParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
}

func (s GetSecretParametersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSecretParametersResponseBody) GoString() string {
	return s.String()
}

func (s *GetSecretParametersResponseBody) SetRequestId(v string) *GetSecretParametersResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSecretParametersResponseBody) SetInvalidParameters(v []*string) *GetSecretParametersResponseBody {
	s.InvalidParameters = v
	return s
}

func (s *GetSecretParametersResponseBody) SetParameters(v []*GetSecretParametersResponseBodyParameters) *GetSecretParametersResponseBody {
	s.Parameters = v
	return s
}

type GetSecretParametersResponseBodyParameters struct {
	Type             *string                `json:"Type,omitempty" xml:"Type,omitempty"`
	UpdatedDate      *string                `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
	UpdatedBy        *string                `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	KeyId            *string                `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	Tags             map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	Value            *string                `json:"Value,omitempty" xml:"Value,omitempty"`
	Description      *string                `json:"Description,omitempty" xml:"Description,omitempty"`
	Constraints      *string                `json:"Constraints,omitempty" xml:"Constraints,omitempty"`
	ResourceGroupId  *string                `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	CreatedBy        *string                `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	CreatedDate      *string                `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	ParameterVersion *int32                 `json:"ParameterVersion,omitempty" xml:"ParameterVersion,omitempty"`
	Name             *string                `json:"Name,omitempty" xml:"Name,omitempty"`
	Id               *string                `json:"Id,omitempty" xml:"Id,omitempty"`
	ShareType        *string                `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
}

func (s GetSecretParametersResponseBodyParameters) String() string {
	return tea.Prettify(s)
}

func (s GetSecretParametersResponseBodyParameters) GoString() string {
	return s.String()
}

func (s *GetSecretParametersResponseBodyParameters) SetType(v string) *GetSecretParametersResponseBodyParameters {
	s.Type = &v
	return s
}

func (s *GetSecretParametersResponseBodyParameters) SetUpdatedDate(v string) *GetSecretParametersResponseBodyParameters {
	s.UpdatedDate = &v
	return s
}

func (s *GetSecretParametersResponseBodyParameters) SetUpdatedBy(v string) *GetSecretParametersResponseBodyParameters {
	s.UpdatedBy = &v
	return s
}

func (s *GetSecretParametersResponseBodyParameters) SetKeyId(v string) *GetSecretParametersResponseBodyParameters {
	s.KeyId = &v
	return s
}

func (s *GetSecretParametersResponseBodyParameters) SetTags(v map[string]interface{}) *GetSecretParametersResponseBodyParameters {
	s.Tags = v
	return s
}

func (s *GetSecretParametersResponseBodyParameters) SetValue(v string) *GetSecretParametersResponseBodyParameters {
	s.Value = &v
	return s
}

func (s *GetSecretParametersResponseBodyParameters) SetDescription(v string) *GetSecretParametersResponseBodyParameters {
	s.Description = &v
	return s
}

func (s *GetSecretParametersResponseBodyParameters) SetConstraints(v string) *GetSecretParametersResponseBodyParameters {
	s.Constraints = &v
	return s
}

func (s *GetSecretParametersResponseBodyParameters) SetResourceGroupId(v string) *GetSecretParametersResponseBodyParameters {
	s.ResourceGroupId = &v
	return s
}

func (s *GetSecretParametersResponseBodyParameters) SetCreatedBy(v string) *GetSecretParametersResponseBodyParameters {
	s.CreatedBy = &v
	return s
}

func (s *GetSecretParametersResponseBodyParameters) SetCreatedDate(v string) *GetSecretParametersResponseBodyParameters {
	s.CreatedDate = &v
	return s
}

func (s *GetSecretParametersResponseBodyParameters) SetParameterVersion(v int32) *GetSecretParametersResponseBodyParameters {
	s.ParameterVersion = &v
	return s
}

func (s *GetSecretParametersResponseBodyParameters) SetName(v string) *GetSecretParametersResponseBodyParameters {
	s.Name = &v
	return s
}

func (s *GetSecretParametersResponseBodyParameters) SetId(v string) *GetSecretParametersResponseBodyParameters {
	s.Id = &v
	return s
}

func (s *GetSecretParametersResponseBodyParameters) SetShareType(v string) *GetSecretParametersResponseBodyParameters {
	s.ShareType = &v
	return s
}

type GetSecretParametersResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSecretParametersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSecretParametersResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSecretParametersResponse) GoString() string {
	return s.String()
}

func (s *GetSecretParametersResponse) SetHeaders(v map[string]*string) *GetSecretParametersResponse {
	s.Headers = v
	return s
}

func (s *GetSecretParametersResponse) SetBody(v *GetSecretParametersResponseBody) *GetSecretParametersResponse {
	s.Body = v
	return s
}

type GetSecretParametersByPathRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Path           *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Recursive      *bool   `json:"Recursive,omitempty" xml:"Recursive,omitempty"`
	NextToken      *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	MaxResults     *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	WithDecryption *bool   `json:"WithDecryption,omitempty" xml:"WithDecryption,omitempty"`
}

func (s GetSecretParametersByPathRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSecretParametersByPathRequest) GoString() string {
	return s.String()
}

func (s *GetSecretParametersByPathRequest) SetRegionId(v string) *GetSecretParametersByPathRequest {
	s.RegionId = &v
	return s
}

func (s *GetSecretParametersByPathRequest) SetPath(v string) *GetSecretParametersByPathRequest {
	s.Path = &v
	return s
}

func (s *GetSecretParametersByPathRequest) SetRecursive(v bool) *GetSecretParametersByPathRequest {
	s.Recursive = &v
	return s
}

func (s *GetSecretParametersByPathRequest) SetNextToken(v string) *GetSecretParametersByPathRequest {
	s.NextToken = &v
	return s
}

func (s *GetSecretParametersByPathRequest) SetMaxResults(v int32) *GetSecretParametersByPathRequest {
	s.MaxResults = &v
	return s
}

func (s *GetSecretParametersByPathRequest) SetWithDecryption(v bool) *GetSecretParametersByPathRequest {
	s.WithDecryption = &v
	return s
}

type GetSecretParametersByPathResponseBody struct {
	NextToken  *string                                            `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId  *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	MaxResults *int32                                             `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Parameters []*GetSecretParametersByPathResponseBodyParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
}

func (s GetSecretParametersByPathResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSecretParametersByPathResponseBody) GoString() string {
	return s.String()
}

func (s *GetSecretParametersByPathResponseBody) SetNextToken(v string) *GetSecretParametersByPathResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetSecretParametersByPathResponseBody) SetRequestId(v string) *GetSecretParametersByPathResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSecretParametersByPathResponseBody) SetTotalCount(v int32) *GetSecretParametersByPathResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetSecretParametersByPathResponseBody) SetMaxResults(v int32) *GetSecretParametersByPathResponseBody {
	s.MaxResults = &v
	return s
}

func (s *GetSecretParametersByPathResponseBody) SetParameters(v []*GetSecretParametersByPathResponseBodyParameters) *GetSecretParametersByPathResponseBody {
	s.Parameters = v
	return s
}

type GetSecretParametersByPathResponseBodyParameters struct {
	Type             *string `json:"Type,omitempty" xml:"Type,omitempty"`
	UpdatedDate      *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
	UpdatedBy        *string `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	KeyId            *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	Value            *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Constraints      *string `json:"Constraints,omitempty" xml:"Constraints,omitempty"`
	CreatedBy        *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	CreatedDate      *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	ParameterVersion *int32  `json:"ParameterVersion,omitempty" xml:"ParameterVersion,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id               *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ShareType        *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
}

func (s GetSecretParametersByPathResponseBodyParameters) String() string {
	return tea.Prettify(s)
}

func (s GetSecretParametersByPathResponseBodyParameters) GoString() string {
	return s.String()
}

func (s *GetSecretParametersByPathResponseBodyParameters) SetType(v string) *GetSecretParametersByPathResponseBodyParameters {
	s.Type = &v
	return s
}

func (s *GetSecretParametersByPathResponseBodyParameters) SetUpdatedDate(v string) *GetSecretParametersByPathResponseBodyParameters {
	s.UpdatedDate = &v
	return s
}

func (s *GetSecretParametersByPathResponseBodyParameters) SetUpdatedBy(v string) *GetSecretParametersByPathResponseBodyParameters {
	s.UpdatedBy = &v
	return s
}

func (s *GetSecretParametersByPathResponseBodyParameters) SetKeyId(v string) *GetSecretParametersByPathResponseBodyParameters {
	s.KeyId = &v
	return s
}

func (s *GetSecretParametersByPathResponseBodyParameters) SetValue(v string) *GetSecretParametersByPathResponseBodyParameters {
	s.Value = &v
	return s
}

func (s *GetSecretParametersByPathResponseBodyParameters) SetDescription(v string) *GetSecretParametersByPathResponseBodyParameters {
	s.Description = &v
	return s
}

func (s *GetSecretParametersByPathResponseBodyParameters) SetConstraints(v string) *GetSecretParametersByPathResponseBodyParameters {
	s.Constraints = &v
	return s
}

func (s *GetSecretParametersByPathResponseBodyParameters) SetCreatedBy(v string) *GetSecretParametersByPathResponseBodyParameters {
	s.CreatedBy = &v
	return s
}

func (s *GetSecretParametersByPathResponseBodyParameters) SetCreatedDate(v string) *GetSecretParametersByPathResponseBodyParameters {
	s.CreatedDate = &v
	return s
}

func (s *GetSecretParametersByPathResponseBodyParameters) SetParameterVersion(v int32) *GetSecretParametersByPathResponseBodyParameters {
	s.ParameterVersion = &v
	return s
}

func (s *GetSecretParametersByPathResponseBodyParameters) SetName(v string) *GetSecretParametersByPathResponseBodyParameters {
	s.Name = &v
	return s
}

func (s *GetSecretParametersByPathResponseBodyParameters) SetId(v string) *GetSecretParametersByPathResponseBodyParameters {
	s.Id = &v
	return s
}

func (s *GetSecretParametersByPathResponseBodyParameters) SetShareType(v string) *GetSecretParametersByPathResponseBodyParameters {
	s.ShareType = &v
	return s
}

type GetSecretParametersByPathResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSecretParametersByPathResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSecretParametersByPathResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSecretParametersByPathResponse) GoString() string {
	return s.String()
}

func (s *GetSecretParametersByPathResponse) SetHeaders(v map[string]*string) *GetSecretParametersByPathResponse {
	s.Headers = v
	return s
}

func (s *GetSecretParametersByPathResponse) SetBody(v *GetSecretParametersByPathResponseBody) *GetSecretParametersByPathResponse {
	s.Body = v
	return s
}

type GetServiceSettingsRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetServiceSettingsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetServiceSettingsRequest) GoString() string {
	return s.String()
}

func (s *GetServiceSettingsRequest) SetRegionId(v string) *GetServiceSettingsRequest {
	s.RegionId = &v
	return s
}

type GetServiceSettingsResponseBody struct {
	RequestId       *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServiceSettings []*GetServiceSettingsResponseBodyServiceSettings `json:"ServiceSettings,omitempty" xml:"ServiceSettings,omitempty" type:"Repeated"`
}

func (s GetServiceSettingsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetServiceSettingsResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceSettingsResponseBody) SetRequestId(v string) *GetServiceSettingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServiceSettingsResponseBody) SetServiceSettings(v []*GetServiceSettingsResponseBodyServiceSettings) *GetServiceSettingsResponseBody {
	s.ServiceSettings = v
	return s
}

type GetServiceSettingsResponseBodyServiceSettings struct {
	DeliveryOssBucketName  *string `json:"DeliveryOssBucketName,omitempty" xml:"DeliveryOssBucketName,omitempty"`
	DeliveryOssKeyPrefix   *string `json:"DeliveryOssKeyPrefix,omitempty" xml:"DeliveryOssKeyPrefix,omitempty"`
	DeliveryOssEnabled     *bool   `json:"DeliveryOssEnabled,omitempty" xml:"DeliveryOssEnabled,omitempty"`
	DeliverySlsEnabled     *bool   `json:"DeliverySlsEnabled,omitempty" xml:"DeliverySlsEnabled,omitempty"`
	DeliverySlsProjectName *string `json:"DeliverySlsProjectName,omitempty" xml:"DeliverySlsProjectName,omitempty"`
	RdcEnterpriseId        *string `json:"RdcEnterpriseId,omitempty" xml:"RdcEnterpriseId,omitempty"`
}

func (s GetServiceSettingsResponseBodyServiceSettings) String() string {
	return tea.Prettify(s)
}

func (s GetServiceSettingsResponseBodyServiceSettings) GoString() string {
	return s.String()
}

func (s *GetServiceSettingsResponseBodyServiceSettings) SetDeliveryOssBucketName(v string) *GetServiceSettingsResponseBodyServiceSettings {
	s.DeliveryOssBucketName = &v
	return s
}

func (s *GetServiceSettingsResponseBodyServiceSettings) SetDeliveryOssKeyPrefix(v string) *GetServiceSettingsResponseBodyServiceSettings {
	s.DeliveryOssKeyPrefix = &v
	return s
}

func (s *GetServiceSettingsResponseBodyServiceSettings) SetDeliveryOssEnabled(v bool) *GetServiceSettingsResponseBodyServiceSettings {
	s.DeliveryOssEnabled = &v
	return s
}

func (s *GetServiceSettingsResponseBodyServiceSettings) SetDeliverySlsEnabled(v bool) *GetServiceSettingsResponseBodyServiceSettings {
	s.DeliverySlsEnabled = &v
	return s
}

func (s *GetServiceSettingsResponseBodyServiceSettings) SetDeliverySlsProjectName(v string) *GetServiceSettingsResponseBodyServiceSettings {
	s.DeliverySlsProjectName = &v
	return s
}

func (s *GetServiceSettingsResponseBodyServiceSettings) SetRdcEnterpriseId(v string) *GetServiceSettingsResponseBodyServiceSettings {
	s.RdcEnterpriseId = &v
	return s
}

type GetServiceSettingsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetServiceSettingsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetServiceSettingsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetServiceSettingsResponse) GoString() string {
	return s.String()
}

func (s *GetServiceSettingsResponse) SetHeaders(v map[string]*string) *GetServiceSettingsResponse {
	s.Headers = v
	return s
}

func (s *GetServiceSettingsResponse) SetBody(v *GetServiceSettingsResponseBody) *GetServiceSettingsResponse {
	s.Body = v
	return s
}

type GetTemplateRequest struct {
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TemplateName    *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
}

func (s GetTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetTemplateRequest) SetRegionId(v string) *GetTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *GetTemplateRequest) SetTemplateName(v string) *GetTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *GetTemplateRequest) SetTemplateVersion(v string) *GetTemplateRequest {
	s.TemplateVersion = &v
	return s
}

type GetTemplateResponseBody struct {
	Content   *string                          `json:"Content,omitempty" xml:"Content,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Template  *GetTemplateResponseBodyTemplate `json:"Template,omitempty" xml:"Template,omitempty" type:"Struct"`
}

func (s GetTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetTemplateResponseBody) SetContent(v string) *GetTemplateResponseBody {
	s.Content = &v
	return s
}

func (s *GetTemplateResponseBody) SetRequestId(v string) *GetTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTemplateResponseBody) SetTemplate(v *GetTemplateResponseBodyTemplate) *GetTemplateResponseBody {
	s.Template = v
	return s
}

type GetTemplateResponseBodyTemplate struct {
	Hash            *string                `json:"Hash,omitempty" xml:"Hash,omitempty"`
	UpdatedDate     *string                `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
	UpdatedBy       *string                `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	TemplateType    *string                `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	Tags            map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	TemplateName    *string                `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateVersion *string                `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	TemplateFormat  *string                `json:"TemplateFormat,omitempty" xml:"TemplateFormat,omitempty"`
	Description     *string                `json:"Description,omitempty" xml:"Description,omitempty"`
	ResourceGroupId *string                `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	CreatedBy       *string                `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	CreatedDate     *string                `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	VersionName     *string                `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
	TemplateId      *string                `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	HasTrigger      *bool                  `json:"HasTrigger,omitempty" xml:"HasTrigger,omitempty"`
	ShareType       *string                `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
}

func (s GetTemplateResponseBodyTemplate) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateResponseBodyTemplate) GoString() string {
	return s.String()
}

func (s *GetTemplateResponseBodyTemplate) SetHash(v string) *GetTemplateResponseBodyTemplate {
	s.Hash = &v
	return s
}

func (s *GetTemplateResponseBodyTemplate) SetUpdatedDate(v string) *GetTemplateResponseBodyTemplate {
	s.UpdatedDate = &v
	return s
}

func (s *GetTemplateResponseBodyTemplate) SetUpdatedBy(v string) *GetTemplateResponseBodyTemplate {
	s.UpdatedBy = &v
	return s
}

func (s *GetTemplateResponseBodyTemplate) SetTemplateType(v string) *GetTemplateResponseBodyTemplate {
	s.TemplateType = &v
	return s
}

func (s *GetTemplateResponseBodyTemplate) SetTags(v map[string]interface{}) *GetTemplateResponseBodyTemplate {
	s.Tags = v
	return s
}

func (s *GetTemplateResponseBodyTemplate) SetTemplateName(v string) *GetTemplateResponseBodyTemplate {
	s.TemplateName = &v
	return s
}

func (s *GetTemplateResponseBodyTemplate) SetTemplateVersion(v string) *GetTemplateResponseBodyTemplate {
	s.TemplateVersion = &v
	return s
}

func (s *GetTemplateResponseBodyTemplate) SetTemplateFormat(v string) *GetTemplateResponseBodyTemplate {
	s.TemplateFormat = &v
	return s
}

func (s *GetTemplateResponseBodyTemplate) SetDescription(v string) *GetTemplateResponseBodyTemplate {
	s.Description = &v
	return s
}

func (s *GetTemplateResponseBodyTemplate) SetResourceGroupId(v string) *GetTemplateResponseBodyTemplate {
	s.ResourceGroupId = &v
	return s
}

func (s *GetTemplateResponseBodyTemplate) SetCreatedBy(v string) *GetTemplateResponseBodyTemplate {
	s.CreatedBy = &v
	return s
}

func (s *GetTemplateResponseBodyTemplate) SetCreatedDate(v string) *GetTemplateResponseBodyTemplate {
	s.CreatedDate = &v
	return s
}

func (s *GetTemplateResponseBodyTemplate) SetVersionName(v string) *GetTemplateResponseBodyTemplate {
	s.VersionName = &v
	return s
}

func (s *GetTemplateResponseBodyTemplate) SetTemplateId(v string) *GetTemplateResponseBodyTemplate {
	s.TemplateId = &v
	return s
}

func (s *GetTemplateResponseBodyTemplate) SetHasTrigger(v bool) *GetTemplateResponseBodyTemplate {
	s.HasTrigger = &v
	return s
}

func (s *GetTemplateResponseBodyTemplate) SetShareType(v string) *GetTemplateResponseBodyTemplate {
	s.ShareType = &v
	return s
}

type GetTemplateResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetTemplateResponse) SetHeaders(v map[string]*string) *GetTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetTemplateResponse) SetBody(v *GetTemplateResponseBody) *GetTemplateResponse {
	s.Body = v
	return s
}

type ListActionsRequest struct {
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	OOSActionName *string `json:"OOSActionName,omitempty" xml:"OOSActionName,omitempty"`
	MaxResults    *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken     *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListActionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListActionsRequest) GoString() string {
	return s.String()
}

func (s *ListActionsRequest) SetRegionId(v string) *ListActionsRequest {
	s.RegionId = &v
	return s
}

func (s *ListActionsRequest) SetOOSActionName(v string) *ListActionsRequest {
	s.OOSActionName = &v
	return s
}

func (s *ListActionsRequest) SetMaxResults(v int32) *ListActionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListActionsRequest) SetNextToken(v string) *ListActionsRequest {
	s.NextToken = &v
	return s
}

type ListActionsResponseBody struct {
	NextToken  *string                           `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId  *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	MaxResults *int32                            `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Actions    []*ListActionsResponseBodyActions `json:"Actions,omitempty" xml:"Actions,omitempty" type:"Repeated"`
}

func (s ListActionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListActionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListActionsResponseBody) SetNextToken(v string) *ListActionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListActionsResponseBody) SetRequestId(v string) *ListActionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListActionsResponseBody) SetMaxResults(v int32) *ListActionsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListActionsResponseBody) SetActions(v []*ListActionsResponseBodyActions) *ListActionsResponseBody {
	s.Actions = v
	return s
}

type ListActionsResponseBodyActions struct {
	Popularity      *int32  `json:"Popularity,omitempty" xml:"Popularity,omitempty"`
	ActionType      *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	CreatedDate     *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	OOSActionName   *string `json:"OOSActionName,omitempty" xml:"OOSActionName,omitempty"`
	Properties      *string `json:"Properties,omitempty" xml:"Properties,omitempty"`
}

func (s ListActionsResponseBodyActions) String() string {
	return tea.Prettify(s)
}

func (s ListActionsResponseBodyActions) GoString() string {
	return s.String()
}

func (s *ListActionsResponseBodyActions) SetPopularity(v int32) *ListActionsResponseBodyActions {
	s.Popularity = &v
	return s
}

func (s *ListActionsResponseBodyActions) SetActionType(v string) *ListActionsResponseBodyActions {
	s.ActionType = &v
	return s
}

func (s *ListActionsResponseBodyActions) SetDescription(v string) *ListActionsResponseBodyActions {
	s.Description = &v
	return s
}

func (s *ListActionsResponseBodyActions) SetCreatedDate(v string) *ListActionsResponseBodyActions {
	s.CreatedDate = &v
	return s
}

func (s *ListActionsResponseBodyActions) SetTemplateVersion(v string) *ListActionsResponseBodyActions {
	s.TemplateVersion = &v
	return s
}

func (s *ListActionsResponseBodyActions) SetOOSActionName(v string) *ListActionsResponseBodyActions {
	s.OOSActionName = &v
	return s
}

func (s *ListActionsResponseBodyActions) SetProperties(v string) *ListActionsResponseBodyActions {
	s.Properties = &v
	return s
}

type ListActionsResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListActionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListActionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListActionsResponse) GoString() string {
	return s.String()
}

func (s *ListActionsResponse) SetHeaders(v map[string]*string) *ListActionsResponse {
	s.Headers = v
	return s
}

func (s *ListActionsResponse) SetBody(v *ListActionsResponseBody) *ListActionsResponse {
	s.Body = v
	return s
}

type ListApplicationGroupsRequest struct {
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	DeployRegionId  *bool   `json:"DeployRegionId,omitempty" xml:"DeployRegionId,omitempty"`
	Environment     *bool   `json:"Environment,omitempty" xml:"Environment,omitempty"`
	MaxResults      *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken       *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListApplicationGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListApplicationGroupsRequest) SetRegionId(v string) *ListApplicationGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *ListApplicationGroupsRequest) SetApplicationName(v string) *ListApplicationGroupsRequest {
	s.ApplicationName = &v
	return s
}

func (s *ListApplicationGroupsRequest) SetDeployRegionId(v bool) *ListApplicationGroupsRequest {
	s.DeployRegionId = &v
	return s
}

func (s *ListApplicationGroupsRequest) SetEnvironment(v bool) *ListApplicationGroupsRequest {
	s.Environment = &v
	return s
}

func (s *ListApplicationGroupsRequest) SetMaxResults(v int32) *ListApplicationGroupsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListApplicationGroupsRequest) SetNextToken(v string) *ListApplicationGroupsRequest {
	s.NextToken = &v
	return s
}

type ListApplicationGroupsResponseBody struct {
	NextToken         *string                                               `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId         *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	MaxResults        *int32                                                `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	ApplicationGroups []*ListApplicationGroupsResponseBodyApplicationGroups `json:"ApplicationGroups,omitempty" xml:"ApplicationGroups,omitempty" type:"Repeated"`
}

func (s ListApplicationGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicationGroupsResponseBody) SetNextToken(v string) *ListApplicationGroupsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListApplicationGroupsResponseBody) SetRequestId(v string) *ListApplicationGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApplicationGroupsResponseBody) SetMaxResults(v int32) *ListApplicationGroupsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListApplicationGroupsResponseBody) SetApplicationGroups(v []*ListApplicationGroupsResponseBodyApplicationGroups) *ListApplicationGroupsResponseBody {
	s.ApplicationGroups = v
	return s
}

type ListApplicationGroupsResponseBodyApplicationGroups struct {
	DeployRegionId  *string `json:"DeployRegionId,omitempty" xml:"DeployRegionId,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	UpdatedDate     *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
	CreatedDate     *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Environment     *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	CreateType      *string `json:"CreateType,omitempty" xml:"CreateType,omitempty"`
	ScalingGroupId  *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	ImportClusterId *string `json:"ImportClusterId,omitempty" xml:"ImportClusterId,omitempty"`
}

func (s ListApplicationGroupsResponseBodyApplicationGroups) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationGroupsResponseBodyApplicationGroups) GoString() string {
	return s.String()
}

func (s *ListApplicationGroupsResponseBodyApplicationGroups) SetDeployRegionId(v string) *ListApplicationGroupsResponseBodyApplicationGroups {
	s.DeployRegionId = &v
	return s
}

func (s *ListApplicationGroupsResponseBodyApplicationGroups) SetDescription(v string) *ListApplicationGroupsResponseBodyApplicationGroups {
	s.Description = &v
	return s
}

func (s *ListApplicationGroupsResponseBodyApplicationGroups) SetUpdatedDate(v string) *ListApplicationGroupsResponseBodyApplicationGroups {
	s.UpdatedDate = &v
	return s
}

func (s *ListApplicationGroupsResponseBodyApplicationGroups) SetCreatedDate(v string) *ListApplicationGroupsResponseBodyApplicationGroups {
	s.CreatedDate = &v
	return s
}

func (s *ListApplicationGroupsResponseBodyApplicationGroups) SetApplicationName(v string) *ListApplicationGroupsResponseBodyApplicationGroups {
	s.ApplicationName = &v
	return s
}

func (s *ListApplicationGroupsResponseBodyApplicationGroups) SetName(v string) *ListApplicationGroupsResponseBodyApplicationGroups {
	s.Name = &v
	return s
}

func (s *ListApplicationGroupsResponseBodyApplicationGroups) SetEnvironment(v string) *ListApplicationGroupsResponseBodyApplicationGroups {
	s.Environment = &v
	return s
}

func (s *ListApplicationGroupsResponseBodyApplicationGroups) SetCreateType(v string) *ListApplicationGroupsResponseBodyApplicationGroups {
	s.CreateType = &v
	return s
}

func (s *ListApplicationGroupsResponseBodyApplicationGroups) SetScalingGroupId(v string) *ListApplicationGroupsResponseBodyApplicationGroups {
	s.ScalingGroupId = &v
	return s
}

func (s *ListApplicationGroupsResponseBodyApplicationGroups) SetImportClusterId(v string) *ListApplicationGroupsResponseBodyApplicationGroups {
	s.ImportClusterId = &v
	return s
}

type ListApplicationGroupsResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListApplicationGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListApplicationGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListApplicationGroupsResponse) SetHeaders(v map[string]*string) *ListApplicationGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListApplicationGroupsResponse) SetBody(v *ListApplicationGroupsResponseBody) *ListApplicationGroupsResponse {
	s.Body = v
	return s
}

type ListApplicationsRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	IsSystem   *bool   `json:"IsSystem,omitempty" xml:"IsSystem,omitempty"`
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListApplicationsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationsRequest) GoString() string {
	return s.String()
}

func (s *ListApplicationsRequest) SetRegionId(v string) *ListApplicationsRequest {
	s.RegionId = &v
	return s
}

func (s *ListApplicationsRequest) SetIsSystem(v bool) *ListApplicationsRequest {
	s.IsSystem = &v
	return s
}

func (s *ListApplicationsRequest) SetMaxResults(v int32) *ListApplicationsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListApplicationsRequest) SetNextToken(v string) *ListApplicationsRequest {
	s.NextToken = &v
	return s
}

type ListApplicationsResponseBody struct {
	NextToken    *string                                     `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	MaxResults   *int32                                      `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Applications []*ListApplicationsResponseBodyApplications `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Repeated"`
}

func (s ListApplicationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponseBody) SetNextToken(v string) *ListApplicationsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListApplicationsResponseBody) SetRequestId(v string) *ListApplicationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApplicationsResponseBody) SetMaxResults(v int32) *ListApplicationsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListApplicationsResponseBody) SetApplications(v []*ListApplicationsResponseBodyApplications) *ListApplicationsResponseBody {
	s.Applications = v
	return s
}

type ListApplicationsResponseBodyApplications struct {
	Type             *string                                                   `json:"Type,omitempty" xml:"Type,omitempty"`
	IsSystem         *bool                                                     `json:"IsSystem,omitempty" xml:"IsSystem,omitempty"`
	Description      *string                                                   `json:"Description,omitempty" xml:"Description,omitempty"`
	UpdateDate       *string                                                   `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
	ResourceGroupId  *string                                                   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	CreatedDate      *string                                                   `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	Name             *string                                                   `json:"Name,omitempty" xml:"Name,omitempty"`
	CloudMonitorRule *ListApplicationsResponseBodyApplicationsCloudMonitorRule `json:"CloudMonitorRule,omitempty" xml:"CloudMonitorRule,omitempty" type:"Struct"`
}

func (s ListApplicationsResponseBodyApplications) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationsResponseBodyApplications) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponseBodyApplications) SetType(v string) *ListApplicationsResponseBodyApplications {
	s.Type = &v
	return s
}

func (s *ListApplicationsResponseBodyApplications) SetIsSystem(v bool) *ListApplicationsResponseBodyApplications {
	s.IsSystem = &v
	return s
}

func (s *ListApplicationsResponseBodyApplications) SetDescription(v string) *ListApplicationsResponseBodyApplications {
	s.Description = &v
	return s
}

func (s *ListApplicationsResponseBodyApplications) SetUpdateDate(v string) *ListApplicationsResponseBodyApplications {
	s.UpdateDate = &v
	return s
}

func (s *ListApplicationsResponseBodyApplications) SetResourceGroupId(v string) *ListApplicationsResponseBodyApplications {
	s.ResourceGroupId = &v
	return s
}

func (s *ListApplicationsResponseBodyApplications) SetCreatedDate(v string) *ListApplicationsResponseBodyApplications {
	s.CreatedDate = &v
	return s
}

func (s *ListApplicationsResponseBodyApplications) SetName(v string) *ListApplicationsResponseBodyApplications {
	s.Name = &v
	return s
}

func (s *ListApplicationsResponseBodyApplications) SetCloudMonitorRule(v *ListApplicationsResponseBodyApplicationsCloudMonitorRule) *ListApplicationsResponseBodyApplications {
	s.CloudMonitorRule = v
	return s
}

type ListApplicationsResponseBodyApplicationsCloudMonitorRule struct {
	EnableSubscribeEvent *bool     `json:"EnableSubscribeEvent,omitempty" xml:"EnableSubscribeEvent,omitempty"`
	EnableInstallAgent   *bool     `json:"EnableInstallAgent,omitempty" xml:"EnableInstallAgent,omitempty"`
	Enabled              *bool     `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	ContactGroupList     []*string `json:"ContactGroupList,omitempty" xml:"ContactGroupList,omitempty" type:"Repeated"`
	TemplateIdList       []*int32  `json:"TemplateIdList,omitempty" xml:"TemplateIdList,omitempty" type:"Repeated"`
}

func (s ListApplicationsResponseBodyApplicationsCloudMonitorRule) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationsResponseBodyApplicationsCloudMonitorRule) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponseBodyApplicationsCloudMonitorRule) SetEnableSubscribeEvent(v bool) *ListApplicationsResponseBodyApplicationsCloudMonitorRule {
	s.EnableSubscribeEvent = &v
	return s
}

func (s *ListApplicationsResponseBodyApplicationsCloudMonitorRule) SetEnableInstallAgent(v bool) *ListApplicationsResponseBodyApplicationsCloudMonitorRule {
	s.EnableInstallAgent = &v
	return s
}

func (s *ListApplicationsResponseBodyApplicationsCloudMonitorRule) SetEnabled(v bool) *ListApplicationsResponseBodyApplicationsCloudMonitorRule {
	s.Enabled = &v
	return s
}

func (s *ListApplicationsResponseBodyApplicationsCloudMonitorRule) SetContactGroupList(v []*string) *ListApplicationsResponseBodyApplicationsCloudMonitorRule {
	s.ContactGroupList = v
	return s
}

func (s *ListApplicationsResponseBodyApplicationsCloudMonitorRule) SetTemplateIdList(v []*int32) *ListApplicationsResponseBodyApplicationsCloudMonitorRule {
	s.TemplateIdList = v
	return s
}

type ListApplicationsResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListApplicationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListApplicationsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationsResponse) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponse) SetHeaders(v map[string]*string) *ListApplicationsResponse {
	s.Headers = v
	return s
}

func (s *ListApplicationsResponse) SetBody(v *ListApplicationsResponseBody) *ListApplicationsResponse {
	s.Body = v
	return s
}

type ListExecutionLogsRequest struct {
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ExecutionId     *string `json:"ExecutionId,omitempty" xml:"ExecutionId,omitempty"`
	TaskExecutionId *string `json:"TaskExecutionId,omitempty" xml:"TaskExecutionId,omitempty"`
	LogType         *string `json:"LogType,omitempty" xml:"LogType,omitempty"`
	MaxResults      *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken       *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListExecutionLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListExecutionLogsRequest) GoString() string {
	return s.String()
}

func (s *ListExecutionLogsRequest) SetRegionId(v string) *ListExecutionLogsRequest {
	s.RegionId = &v
	return s
}

func (s *ListExecutionLogsRequest) SetExecutionId(v string) *ListExecutionLogsRequest {
	s.ExecutionId = &v
	return s
}

func (s *ListExecutionLogsRequest) SetTaskExecutionId(v string) *ListExecutionLogsRequest {
	s.TaskExecutionId = &v
	return s
}

func (s *ListExecutionLogsRequest) SetLogType(v string) *ListExecutionLogsRequest {
	s.LogType = &v
	return s
}

func (s *ListExecutionLogsRequest) SetMaxResults(v int32) *ListExecutionLogsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListExecutionLogsRequest) SetNextToken(v string) *ListExecutionLogsRequest {
	s.NextToken = &v
	return s
}

type ListExecutionLogsResponseBody struct {
	NextToken     *string                                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId     *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	IsTruncated   *bool                                         `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	MaxResults    *int32                                        `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	ExecutionLogs []*ListExecutionLogsResponseBodyExecutionLogs `json:"ExecutionLogs,omitempty" xml:"ExecutionLogs,omitempty" type:"Repeated"`
}

func (s ListExecutionLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListExecutionLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListExecutionLogsResponseBody) SetNextToken(v string) *ListExecutionLogsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListExecutionLogsResponseBody) SetRequestId(v string) *ListExecutionLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListExecutionLogsResponseBody) SetIsTruncated(v bool) *ListExecutionLogsResponseBody {
	s.IsTruncated = &v
	return s
}

func (s *ListExecutionLogsResponseBody) SetMaxResults(v int32) *ListExecutionLogsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListExecutionLogsResponseBody) SetExecutionLogs(v []*ListExecutionLogsResponseBodyExecutionLogs) *ListExecutionLogsResponseBody {
	s.ExecutionLogs = v
	return s
}

type ListExecutionLogsResponseBodyExecutionLogs struct {
	TaskExecutionId *string `json:"TaskExecutionId,omitempty" xml:"TaskExecutionId,omitempty"`
	Message         *string `json:"Message,omitempty" xml:"Message,omitempty"`
	LogType         *string `json:"LogType,omitempty" xml:"LogType,omitempty"`
	Timestamp       *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s ListExecutionLogsResponseBodyExecutionLogs) String() string {
	return tea.Prettify(s)
}

func (s ListExecutionLogsResponseBodyExecutionLogs) GoString() string {
	return s.String()
}

func (s *ListExecutionLogsResponseBodyExecutionLogs) SetTaskExecutionId(v string) *ListExecutionLogsResponseBodyExecutionLogs {
	s.TaskExecutionId = &v
	return s
}

func (s *ListExecutionLogsResponseBodyExecutionLogs) SetMessage(v string) *ListExecutionLogsResponseBodyExecutionLogs {
	s.Message = &v
	return s
}

func (s *ListExecutionLogsResponseBodyExecutionLogs) SetLogType(v string) *ListExecutionLogsResponseBodyExecutionLogs {
	s.LogType = &v
	return s
}

func (s *ListExecutionLogsResponseBodyExecutionLogs) SetTimestamp(v string) *ListExecutionLogsResponseBodyExecutionLogs {
	s.Timestamp = &v
	return s
}

type ListExecutionLogsResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListExecutionLogsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListExecutionLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListExecutionLogsResponse) GoString() string {
	return s.String()
}

func (s *ListExecutionLogsResponse) SetHeaders(v map[string]*string) *ListExecutionLogsResponse {
	s.Headers = v
	return s
}

func (s *ListExecutionLogsResponse) SetBody(v *ListExecutionLogsResponseBody) *ListExecutionLogsResponse {
	s.Body = v
	return s
}

type ListExecutionRiskyTasksRequest struct {
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s ListExecutionRiskyTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListExecutionRiskyTasksRequest) GoString() string {
	return s.String()
}

func (s *ListExecutionRiskyTasksRequest) SetRegionId(v string) *ListExecutionRiskyTasksRequest {
	s.RegionId = &v
	return s
}

func (s *ListExecutionRiskyTasksRequest) SetTemplateName(v string) *ListExecutionRiskyTasksRequest {
	s.TemplateName = &v
	return s
}

type ListExecutionRiskyTasksResponseBody struct {
	RequestId  *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RiskyTasks []*ListExecutionRiskyTasksResponseBodyRiskyTasks `json:"RiskyTasks,omitempty" xml:"RiskyTasks,omitempty" type:"Repeated"`
}

func (s ListExecutionRiskyTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListExecutionRiskyTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListExecutionRiskyTasksResponseBody) SetRequestId(v string) *ListExecutionRiskyTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListExecutionRiskyTasksResponseBody) SetRiskyTasks(v []*ListExecutionRiskyTasksResponseBodyRiskyTasks) *ListExecutionRiskyTasksResponseBody {
	s.RiskyTasks = v
	return s
}

type ListExecutionRiskyTasksResponseBodyRiskyTasks struct {
	Service  *string   `json:"Service,omitempty" xml:"Service,omitempty"`
	API      *string   `json:"API,omitempty" xml:"API,omitempty"`
	Task     []*string `json:"Task,omitempty" xml:"Task,omitempty" type:"Repeated"`
	Template []*string `json:"Template,omitempty" xml:"Template,omitempty" type:"Repeated"`
}

func (s ListExecutionRiskyTasksResponseBodyRiskyTasks) String() string {
	return tea.Prettify(s)
}

func (s ListExecutionRiskyTasksResponseBodyRiskyTasks) GoString() string {
	return s.String()
}

func (s *ListExecutionRiskyTasksResponseBodyRiskyTasks) SetService(v string) *ListExecutionRiskyTasksResponseBodyRiskyTasks {
	s.Service = &v
	return s
}

func (s *ListExecutionRiskyTasksResponseBodyRiskyTasks) SetAPI(v string) *ListExecutionRiskyTasksResponseBodyRiskyTasks {
	s.API = &v
	return s
}

func (s *ListExecutionRiskyTasksResponseBodyRiskyTasks) SetTask(v []*string) *ListExecutionRiskyTasksResponseBodyRiskyTasks {
	s.Task = v
	return s
}

func (s *ListExecutionRiskyTasksResponseBodyRiskyTasks) SetTemplate(v []*string) *ListExecutionRiskyTasksResponseBodyRiskyTasks {
	s.Template = v
	return s
}

type ListExecutionRiskyTasksResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListExecutionRiskyTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListExecutionRiskyTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListExecutionRiskyTasksResponse) GoString() string {
	return s.String()
}

func (s *ListExecutionRiskyTasksResponse) SetHeaders(v map[string]*string) *ListExecutionRiskyTasksResponse {
	s.Headers = v
	return s
}

func (s *ListExecutionRiskyTasksResponse) SetBody(v *ListExecutionRiskyTasksResponseBody) *ListExecutionRiskyTasksResponse {
	s.Body = v
	return s
}

type ListExecutionsRequest struct {
	RegionId              *string                `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TemplateName          *string                `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	Status                *string                `json:"Status,omitempty" xml:"Status,omitempty"`
	ExecutionId           *string                `json:"ExecutionId,omitempty" xml:"ExecutionId,omitempty"`
	StartDateBefore       *string                `json:"StartDateBefore,omitempty" xml:"StartDateBefore,omitempty"`
	StartDateAfter        *string                `json:"StartDateAfter,omitempty" xml:"StartDateAfter,omitempty"`
	EndDateBefore         *string                `json:"EndDateBefore,omitempty" xml:"EndDateBefore,omitempty"`
	EndDateAfter          *string                `json:"EndDateAfter,omitempty" xml:"EndDateAfter,omitempty"`
	Mode                  *string                `json:"Mode,omitempty" xml:"Mode,omitempty"`
	ExecutedBy            *string                `json:"ExecutedBy,omitempty" xml:"ExecutedBy,omitempty"`
	ParentExecutionId     *string                `json:"ParentExecutionId,omitempty" xml:"ParentExecutionId,omitempty"`
	RamRole               *string                `json:"RamRole,omitempty" xml:"RamRole,omitempty"`
	IncludeChildExecution *bool                  `json:"IncludeChildExecution,omitempty" xml:"IncludeChildExecution,omitempty"`
	Category              *string                `json:"Category,omitempty" xml:"Category,omitempty"`
	Tags                  map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	MaxResults            *int32                 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken             *string                `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	SortField             *string                `json:"SortField,omitempty" xml:"SortField,omitempty"`
	SortOrder             *string                `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	ResourceId            *string                `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceTemplateName  *string                `json:"ResourceTemplateName,omitempty" xml:"ResourceTemplateName,omitempty"`
	ResourceGroupId       *string                `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListExecutionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListExecutionsRequest) GoString() string {
	return s.String()
}

func (s *ListExecutionsRequest) SetRegionId(v string) *ListExecutionsRequest {
	s.RegionId = &v
	return s
}

func (s *ListExecutionsRequest) SetTemplateName(v string) *ListExecutionsRequest {
	s.TemplateName = &v
	return s
}

func (s *ListExecutionsRequest) SetStatus(v string) *ListExecutionsRequest {
	s.Status = &v
	return s
}

func (s *ListExecutionsRequest) SetExecutionId(v string) *ListExecutionsRequest {
	s.ExecutionId = &v
	return s
}

func (s *ListExecutionsRequest) SetStartDateBefore(v string) *ListExecutionsRequest {
	s.StartDateBefore = &v
	return s
}

func (s *ListExecutionsRequest) SetStartDateAfter(v string) *ListExecutionsRequest {
	s.StartDateAfter = &v
	return s
}

func (s *ListExecutionsRequest) SetEndDateBefore(v string) *ListExecutionsRequest {
	s.EndDateBefore = &v
	return s
}

func (s *ListExecutionsRequest) SetEndDateAfter(v string) *ListExecutionsRequest {
	s.EndDateAfter = &v
	return s
}

func (s *ListExecutionsRequest) SetMode(v string) *ListExecutionsRequest {
	s.Mode = &v
	return s
}

func (s *ListExecutionsRequest) SetExecutedBy(v string) *ListExecutionsRequest {
	s.ExecutedBy = &v
	return s
}

func (s *ListExecutionsRequest) SetParentExecutionId(v string) *ListExecutionsRequest {
	s.ParentExecutionId = &v
	return s
}

func (s *ListExecutionsRequest) SetRamRole(v string) *ListExecutionsRequest {
	s.RamRole = &v
	return s
}

func (s *ListExecutionsRequest) SetIncludeChildExecution(v bool) *ListExecutionsRequest {
	s.IncludeChildExecution = &v
	return s
}

func (s *ListExecutionsRequest) SetCategory(v string) *ListExecutionsRequest {
	s.Category = &v
	return s
}

func (s *ListExecutionsRequest) SetTags(v map[string]interface{}) *ListExecutionsRequest {
	s.Tags = v
	return s
}

func (s *ListExecutionsRequest) SetMaxResults(v int32) *ListExecutionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListExecutionsRequest) SetNextToken(v string) *ListExecutionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListExecutionsRequest) SetSortField(v string) *ListExecutionsRequest {
	s.SortField = &v
	return s
}

func (s *ListExecutionsRequest) SetSortOrder(v string) *ListExecutionsRequest {
	s.SortOrder = &v
	return s
}

func (s *ListExecutionsRequest) SetResourceId(v string) *ListExecutionsRequest {
	s.ResourceId = &v
	return s
}

func (s *ListExecutionsRequest) SetResourceTemplateName(v string) *ListExecutionsRequest {
	s.ResourceTemplateName = &v
	return s
}

func (s *ListExecutionsRequest) SetResourceGroupId(v string) *ListExecutionsRequest {
	s.ResourceGroupId = &v
	return s
}

type ListExecutionsShrinkRequest struct {
	RegionId              *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TemplateName          *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	Status                *string `json:"Status,omitempty" xml:"Status,omitempty"`
	ExecutionId           *string `json:"ExecutionId,omitempty" xml:"ExecutionId,omitempty"`
	StartDateBefore       *string `json:"StartDateBefore,omitempty" xml:"StartDateBefore,omitempty"`
	StartDateAfter        *string `json:"StartDateAfter,omitempty" xml:"StartDateAfter,omitempty"`
	EndDateBefore         *string `json:"EndDateBefore,omitempty" xml:"EndDateBefore,omitempty"`
	EndDateAfter          *string `json:"EndDateAfter,omitempty" xml:"EndDateAfter,omitempty"`
	Mode                  *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	ExecutedBy            *string `json:"ExecutedBy,omitempty" xml:"ExecutedBy,omitempty"`
	ParentExecutionId     *string `json:"ParentExecutionId,omitempty" xml:"ParentExecutionId,omitempty"`
	RamRole               *string `json:"RamRole,omitempty" xml:"RamRole,omitempty"`
	IncludeChildExecution *bool   `json:"IncludeChildExecution,omitempty" xml:"IncludeChildExecution,omitempty"`
	Category              *string `json:"Category,omitempty" xml:"Category,omitempty"`
	TagsShrink            *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	MaxResults            *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken             *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	SortField             *string `json:"SortField,omitempty" xml:"SortField,omitempty"`
	SortOrder             *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	ResourceId            *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceTemplateName  *string `json:"ResourceTemplateName,omitempty" xml:"ResourceTemplateName,omitempty"`
	ResourceGroupId       *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListExecutionsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListExecutionsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListExecutionsShrinkRequest) SetRegionId(v string) *ListExecutionsShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ListExecutionsShrinkRequest) SetTemplateName(v string) *ListExecutionsShrinkRequest {
	s.TemplateName = &v
	return s
}

func (s *ListExecutionsShrinkRequest) SetStatus(v string) *ListExecutionsShrinkRequest {
	s.Status = &v
	return s
}

func (s *ListExecutionsShrinkRequest) SetExecutionId(v string) *ListExecutionsShrinkRequest {
	s.ExecutionId = &v
	return s
}

func (s *ListExecutionsShrinkRequest) SetStartDateBefore(v string) *ListExecutionsShrinkRequest {
	s.StartDateBefore = &v
	return s
}

func (s *ListExecutionsShrinkRequest) SetStartDateAfter(v string) *ListExecutionsShrinkRequest {
	s.StartDateAfter = &v
	return s
}

func (s *ListExecutionsShrinkRequest) SetEndDateBefore(v string) *ListExecutionsShrinkRequest {
	s.EndDateBefore = &v
	return s
}

func (s *ListExecutionsShrinkRequest) SetEndDateAfter(v string) *ListExecutionsShrinkRequest {
	s.EndDateAfter = &v
	return s
}

func (s *ListExecutionsShrinkRequest) SetMode(v string) *ListExecutionsShrinkRequest {
	s.Mode = &v
	return s
}

func (s *ListExecutionsShrinkRequest) SetExecutedBy(v string) *ListExecutionsShrinkRequest {
	s.ExecutedBy = &v
	return s
}

func (s *ListExecutionsShrinkRequest) SetParentExecutionId(v string) *ListExecutionsShrinkRequest {
	s.ParentExecutionId = &v
	return s
}

func (s *ListExecutionsShrinkRequest) SetRamRole(v string) *ListExecutionsShrinkRequest {
	s.RamRole = &v
	return s
}

func (s *ListExecutionsShrinkRequest) SetIncludeChildExecution(v bool) *ListExecutionsShrinkRequest {
	s.IncludeChildExecution = &v
	return s
}

func (s *ListExecutionsShrinkRequest) SetCategory(v string) *ListExecutionsShrinkRequest {
	s.Category = &v
	return s
}

func (s *ListExecutionsShrinkRequest) SetTagsShrink(v string) *ListExecutionsShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *ListExecutionsShrinkRequest) SetMaxResults(v int32) *ListExecutionsShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListExecutionsShrinkRequest) SetNextToken(v string) *ListExecutionsShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListExecutionsShrinkRequest) SetSortField(v string) *ListExecutionsShrinkRequest {
	s.SortField = &v
	return s
}

func (s *ListExecutionsShrinkRequest) SetSortOrder(v string) *ListExecutionsShrinkRequest {
	s.SortOrder = &v
	return s
}

func (s *ListExecutionsShrinkRequest) SetResourceId(v string) *ListExecutionsShrinkRequest {
	s.ResourceId = &v
	return s
}

func (s *ListExecutionsShrinkRequest) SetResourceTemplateName(v string) *ListExecutionsShrinkRequest {
	s.ResourceTemplateName = &v
	return s
}

func (s *ListExecutionsShrinkRequest) SetResourceGroupId(v string) *ListExecutionsShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

type ListExecutionsResponseBody struct {
	NextToken  *string                                 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	MaxResults *int32                                  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Executions []*ListExecutionsResponseBodyExecutions `json:"Executions,omitempty" xml:"Executions,omitempty" type:"Repeated"`
}

func (s ListExecutionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListExecutionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListExecutionsResponseBody) SetNextToken(v string) *ListExecutionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListExecutionsResponseBody) SetRequestId(v string) *ListExecutionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListExecutionsResponseBody) SetMaxResults(v int32) *ListExecutionsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListExecutionsResponseBody) SetExecutions(v []*ListExecutionsResponseBodyExecutions) *ListExecutionsResponseBody {
	s.Executions = v
	return s
}

type ListExecutionsResponseBodyExecutions struct {
	Status                    *string                                             `json:"Status,omitempty" xml:"Status,omitempty"`
	WaitingStatus             *string                                             `json:"WaitingStatus,omitempty" xml:"WaitingStatus,omitempty"`
	Targets                   *string                                             `json:"Targets,omitempty" xml:"Targets,omitempty"`
	StatusReason              *string                                             `json:"StatusReason,omitempty" xml:"StatusReason,omitempty"`
	Tags                      map[string]interface{}                              `json:"Tags,omitempty" xml:"Tags,omitempty"`
	LastSuccessfulTriggerTime *string                                             `json:"LastSuccessfulTriggerTime,omitempty" xml:"LastSuccessfulTriggerTime,omitempty"`
	Mode                      *string                                             `json:"Mode,omitempty" xml:"Mode,omitempty"`
	SafetyCheck               *string                                             `json:"SafetyCheck,omitempty" xml:"SafetyCheck,omitempty"`
	TemplateName              *string                                             `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateVersion           *string                                             `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	CreateDate                *string                                             `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	UpdateDate                *string                                             `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
	Description               *string                                             `json:"Description,omitempty" xml:"Description,omitempty"`
	LastTriggerTime           *string                                             `json:"LastTriggerTime,omitempty" xml:"LastTriggerTime,omitempty"`
	ParentExecutionId         *string                                             `json:"ParentExecutionId,omitempty" xml:"ParentExecutionId,omitempty"`
	LastTriggerStatus         *string                                             `json:"LastTriggerStatus,omitempty" xml:"LastTriggerStatus,omitempty"`
	StatusMessage             *string                                             `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	Outputs                   *string                                             `json:"Outputs,omitempty" xml:"Outputs,omitempty"`
	ExecutedBy                *string                                             `json:"ExecutedBy,omitempty" xml:"ExecutedBy,omitempty"`
	EndDate                   *string                                             `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	IsParent                  *bool                                               `json:"IsParent,omitempty" xml:"IsParent,omitempty"`
	StartDate                 *string                                             `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	ExecutionId               *string                                             `json:"ExecutionId,omitempty" xml:"ExecutionId,omitempty"`
	Parameters                map[string]interface{}                              `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	Counters                  map[string]interface{}                              `json:"Counters,omitempty" xml:"Counters,omitempty"`
	ResourceGroupId           *string                                             `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Category                  *string                                             `json:"Category,omitempty" xml:"Category,omitempty"`
	TemplateId                *string                                             `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	RamRole                   *string                                             `json:"RamRole,omitempty" xml:"RamRole,omitempty"`
	ResourceStatus            *string                                             `json:"ResourceStatus,omitempty" xml:"ResourceStatus,omitempty"`
	CurrentTasks              []*ListExecutionsResponseBodyExecutionsCurrentTasks `json:"CurrentTasks,omitempty" xml:"CurrentTasks,omitempty" type:"Repeated"`
}

func (s ListExecutionsResponseBodyExecutions) String() string {
	return tea.Prettify(s)
}

func (s ListExecutionsResponseBodyExecutions) GoString() string {
	return s.String()
}

func (s *ListExecutionsResponseBodyExecutions) SetStatus(v string) *ListExecutionsResponseBodyExecutions {
	s.Status = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetWaitingStatus(v string) *ListExecutionsResponseBodyExecutions {
	s.WaitingStatus = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetTargets(v string) *ListExecutionsResponseBodyExecutions {
	s.Targets = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetStatusReason(v string) *ListExecutionsResponseBodyExecutions {
	s.StatusReason = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetTags(v map[string]interface{}) *ListExecutionsResponseBodyExecutions {
	s.Tags = v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetLastSuccessfulTriggerTime(v string) *ListExecutionsResponseBodyExecutions {
	s.LastSuccessfulTriggerTime = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetMode(v string) *ListExecutionsResponseBodyExecutions {
	s.Mode = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetSafetyCheck(v string) *ListExecutionsResponseBodyExecutions {
	s.SafetyCheck = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetTemplateName(v string) *ListExecutionsResponseBodyExecutions {
	s.TemplateName = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetTemplateVersion(v string) *ListExecutionsResponseBodyExecutions {
	s.TemplateVersion = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetCreateDate(v string) *ListExecutionsResponseBodyExecutions {
	s.CreateDate = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetUpdateDate(v string) *ListExecutionsResponseBodyExecutions {
	s.UpdateDate = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetDescription(v string) *ListExecutionsResponseBodyExecutions {
	s.Description = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetLastTriggerTime(v string) *ListExecutionsResponseBodyExecutions {
	s.LastTriggerTime = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetParentExecutionId(v string) *ListExecutionsResponseBodyExecutions {
	s.ParentExecutionId = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetLastTriggerStatus(v string) *ListExecutionsResponseBodyExecutions {
	s.LastTriggerStatus = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetStatusMessage(v string) *ListExecutionsResponseBodyExecutions {
	s.StatusMessage = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetOutputs(v string) *ListExecutionsResponseBodyExecutions {
	s.Outputs = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetExecutedBy(v string) *ListExecutionsResponseBodyExecutions {
	s.ExecutedBy = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetEndDate(v string) *ListExecutionsResponseBodyExecutions {
	s.EndDate = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetIsParent(v bool) *ListExecutionsResponseBodyExecutions {
	s.IsParent = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetStartDate(v string) *ListExecutionsResponseBodyExecutions {
	s.StartDate = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetExecutionId(v string) *ListExecutionsResponseBodyExecutions {
	s.ExecutionId = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetParameters(v map[string]interface{}) *ListExecutionsResponseBodyExecutions {
	s.Parameters = v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetCounters(v map[string]interface{}) *ListExecutionsResponseBodyExecutions {
	s.Counters = v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetResourceGroupId(v string) *ListExecutionsResponseBodyExecutions {
	s.ResourceGroupId = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetCategory(v string) *ListExecutionsResponseBodyExecutions {
	s.Category = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetTemplateId(v string) *ListExecutionsResponseBodyExecutions {
	s.TemplateId = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetRamRole(v string) *ListExecutionsResponseBodyExecutions {
	s.RamRole = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetResourceStatus(v string) *ListExecutionsResponseBodyExecutions {
	s.ResourceStatus = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetCurrentTasks(v []*ListExecutionsResponseBodyExecutionsCurrentTasks) *ListExecutionsResponseBodyExecutions {
	s.CurrentTasks = v
	return s
}

type ListExecutionsResponseBodyExecutionsCurrentTasks struct {
	TaskExecutionId *string `json:"TaskExecutionId,omitempty" xml:"TaskExecutionId,omitempty"`
	TaskName        *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TaskAction      *string `json:"TaskAction,omitempty" xml:"TaskAction,omitempty"`
}

func (s ListExecutionsResponseBodyExecutionsCurrentTasks) String() string {
	return tea.Prettify(s)
}

func (s ListExecutionsResponseBodyExecutionsCurrentTasks) GoString() string {
	return s.String()
}

func (s *ListExecutionsResponseBodyExecutionsCurrentTasks) SetTaskExecutionId(v string) *ListExecutionsResponseBodyExecutionsCurrentTasks {
	s.TaskExecutionId = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutionsCurrentTasks) SetTaskName(v string) *ListExecutionsResponseBodyExecutionsCurrentTasks {
	s.TaskName = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutionsCurrentTasks) SetTaskAction(v string) *ListExecutionsResponseBodyExecutionsCurrentTasks {
	s.TaskAction = &v
	return s
}

type ListExecutionsResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListExecutionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListExecutionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListExecutionsResponse) GoString() string {
	return s.String()
}

func (s *ListExecutionsResponse) SetHeaders(v map[string]*string) *ListExecutionsResponse {
	s.Headers = v
	return s
}

func (s *ListExecutionsResponse) SetBody(v *ListExecutionsResponseBody) *ListExecutionsResponse {
	s.Body = v
	return s
}

type ListInstancePatchesRequest struct {
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MaxResults    *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken     *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	PatchStatuses *string `json:"PatchStatuses,omitempty" xml:"PatchStatuses,omitempty"`
}

func (s ListInstancePatchesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstancePatchesRequest) GoString() string {
	return s.String()
}

func (s *ListInstancePatchesRequest) SetRegionId(v string) *ListInstancePatchesRequest {
	s.RegionId = &v
	return s
}

func (s *ListInstancePatchesRequest) SetInstanceId(v string) *ListInstancePatchesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListInstancePatchesRequest) SetMaxResults(v int32) *ListInstancePatchesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListInstancePatchesRequest) SetNextToken(v string) *ListInstancePatchesRequest {
	s.NextToken = &v
	return s
}

func (s *ListInstancePatchesRequest) SetPatchStatuses(v string) *ListInstancePatchesRequest {
	s.PatchStatuses = &v
	return s
}

type ListInstancePatchesResponseBody struct {
	NextToken  *string                                   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId  *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	MaxResults *int32                                    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Patches    []*ListInstancePatchesResponseBodyPatches `json:"Patches,omitempty" xml:"Patches,omitempty" type:"Repeated"`
}

func (s ListInstancePatchesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstancePatchesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancePatchesResponseBody) SetNextToken(v string) *ListInstancePatchesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListInstancePatchesResponseBody) SetRequestId(v string) *ListInstancePatchesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstancePatchesResponseBody) SetMaxResults(v int32) *ListInstancePatchesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListInstancePatchesResponseBody) SetPatches(v []*ListInstancePatchesResponseBodyPatches) *ListInstancePatchesResponseBody {
	s.Patches = v
	return s
}

type ListInstancePatchesResponseBodyPatches struct {
	Severity       *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	InstalledTime  *string `json:"InstalledTime,omitempty" xml:"InstalledTime,omitempty"`
	KBId           *string `json:"KBId,omitempty" xml:"KBId,omitempty"`
	Title          *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Classification *string `json:"Classification,omitempty" xml:"Classification,omitempty"`
}

func (s ListInstancePatchesResponseBodyPatches) String() string {
	return tea.Prettify(s)
}

func (s ListInstancePatchesResponseBodyPatches) GoString() string {
	return s.String()
}

func (s *ListInstancePatchesResponseBodyPatches) SetSeverity(v string) *ListInstancePatchesResponseBodyPatches {
	s.Severity = &v
	return s
}

func (s *ListInstancePatchesResponseBodyPatches) SetStatus(v string) *ListInstancePatchesResponseBodyPatches {
	s.Status = &v
	return s
}

func (s *ListInstancePatchesResponseBodyPatches) SetInstalledTime(v string) *ListInstancePatchesResponseBodyPatches {
	s.InstalledTime = &v
	return s
}

func (s *ListInstancePatchesResponseBodyPatches) SetKBId(v string) *ListInstancePatchesResponseBodyPatches {
	s.KBId = &v
	return s
}

func (s *ListInstancePatchesResponseBodyPatches) SetTitle(v string) *ListInstancePatchesResponseBodyPatches {
	s.Title = &v
	return s
}

func (s *ListInstancePatchesResponseBodyPatches) SetClassification(v string) *ListInstancePatchesResponseBodyPatches {
	s.Classification = &v
	return s
}

type ListInstancePatchesResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListInstancePatchesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListInstancePatchesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstancePatchesResponse) GoString() string {
	return s.String()
}

func (s *ListInstancePatchesResponse) SetHeaders(v map[string]*string) *ListInstancePatchesResponse {
	s.Headers = v
	return s
}

func (s *ListInstancePatchesResponse) SetBody(v *ListInstancePatchesResponseBody) *ListInstancePatchesResponse {
	s.Body = v
	return s
}

type ListInstancePatchStatesRequest struct {
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	MaxResults  *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken   *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
}

func (s ListInstancePatchStatesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstancePatchStatesRequest) GoString() string {
	return s.String()
}

func (s *ListInstancePatchStatesRequest) SetRegionId(v string) *ListInstancePatchStatesRequest {
	s.RegionId = &v
	return s
}

func (s *ListInstancePatchStatesRequest) SetMaxResults(v int32) *ListInstancePatchStatesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListInstancePatchStatesRequest) SetNextToken(v string) *ListInstancePatchStatesRequest {
	s.NextToken = &v
	return s
}

func (s *ListInstancePatchStatesRequest) SetInstanceIds(v string) *ListInstancePatchStatesRequest {
	s.InstanceIds = &v
	return s
}

type ListInstancePatchStatesResponseBody struct {
	NextToken           *string                                                   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId           *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	MaxResults          *int32                                                    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	InstancePatchStates []*ListInstancePatchStatesResponseBodyInstancePatchStates `json:"InstancePatchStates,omitempty" xml:"InstancePatchStates,omitempty" type:"Repeated"`
}

func (s ListInstancePatchStatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstancePatchStatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancePatchStatesResponseBody) SetNextToken(v string) *ListInstancePatchStatesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListInstancePatchStatesResponseBody) SetRequestId(v string) *ListInstancePatchStatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstancePatchStatesResponseBody) SetMaxResults(v int32) *ListInstancePatchStatesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListInstancePatchStatesResponseBody) SetInstancePatchStates(v []*ListInstancePatchStatesResponseBodyInstancePatchStates) *ListInstancePatchStatesResponseBody {
	s.InstancePatchStates = v
	return s
}

type ListInstancePatchStatesResponseBodyInstancePatchStates struct {
	MissingCount                *string `json:"MissingCount,omitempty" xml:"MissingCount,omitempty"`
	OperationEndTime            *string `json:"OperationEndTime,omitempty" xml:"OperationEndTime,omitempty"`
	OwnerInformation            *string `json:"OwnerInformation,omitempty" xml:"OwnerInformation,omitempty"`
	InstalledOtherCount         *string `json:"InstalledOtherCount,omitempty" xml:"InstalledOtherCount,omitempty"`
	InstanceId                  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OperationType               *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	OperationStartTime          *string `json:"OperationStartTime,omitempty" xml:"OperationStartTime,omitempty"`
	FailedCount                 *string `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	BaselineId                  *string `json:"BaselineId,omitempty" xml:"BaselineId,omitempty"`
	InstalledPendingRebootCount *string `json:"InstalledPendingRebootCount,omitempty" xml:"InstalledPendingRebootCount,omitempty"`
	InstalledRejectedCount      *string `json:"InstalledRejectedCount,omitempty" xml:"InstalledRejectedCount,omitempty"`
	PatchGroup                  *string `json:"PatchGroup,omitempty" xml:"PatchGroup,omitempty"`
	InstalledCount              *string `json:"InstalledCount,omitempty" xml:"InstalledCount,omitempty"`
}

func (s ListInstancePatchStatesResponseBodyInstancePatchStates) String() string {
	return tea.Prettify(s)
}

func (s ListInstancePatchStatesResponseBodyInstancePatchStates) GoString() string {
	return s.String()
}

func (s *ListInstancePatchStatesResponseBodyInstancePatchStates) SetMissingCount(v string) *ListInstancePatchStatesResponseBodyInstancePatchStates {
	s.MissingCount = &v
	return s
}

func (s *ListInstancePatchStatesResponseBodyInstancePatchStates) SetOperationEndTime(v string) *ListInstancePatchStatesResponseBodyInstancePatchStates {
	s.OperationEndTime = &v
	return s
}

func (s *ListInstancePatchStatesResponseBodyInstancePatchStates) SetOwnerInformation(v string) *ListInstancePatchStatesResponseBodyInstancePatchStates {
	s.OwnerInformation = &v
	return s
}

func (s *ListInstancePatchStatesResponseBodyInstancePatchStates) SetInstalledOtherCount(v string) *ListInstancePatchStatesResponseBodyInstancePatchStates {
	s.InstalledOtherCount = &v
	return s
}

func (s *ListInstancePatchStatesResponseBodyInstancePatchStates) SetInstanceId(v string) *ListInstancePatchStatesResponseBodyInstancePatchStates {
	s.InstanceId = &v
	return s
}

func (s *ListInstancePatchStatesResponseBodyInstancePatchStates) SetOperationType(v string) *ListInstancePatchStatesResponseBodyInstancePatchStates {
	s.OperationType = &v
	return s
}

func (s *ListInstancePatchStatesResponseBodyInstancePatchStates) SetOperationStartTime(v string) *ListInstancePatchStatesResponseBodyInstancePatchStates {
	s.OperationStartTime = &v
	return s
}

func (s *ListInstancePatchStatesResponseBodyInstancePatchStates) SetFailedCount(v string) *ListInstancePatchStatesResponseBodyInstancePatchStates {
	s.FailedCount = &v
	return s
}

func (s *ListInstancePatchStatesResponseBodyInstancePatchStates) SetBaselineId(v string) *ListInstancePatchStatesResponseBodyInstancePatchStates {
	s.BaselineId = &v
	return s
}

func (s *ListInstancePatchStatesResponseBodyInstancePatchStates) SetInstalledPendingRebootCount(v string) *ListInstancePatchStatesResponseBodyInstancePatchStates {
	s.InstalledPendingRebootCount = &v
	return s
}

func (s *ListInstancePatchStatesResponseBodyInstancePatchStates) SetInstalledRejectedCount(v string) *ListInstancePatchStatesResponseBodyInstancePatchStates {
	s.InstalledRejectedCount = &v
	return s
}

func (s *ListInstancePatchStatesResponseBodyInstancePatchStates) SetPatchGroup(v string) *ListInstancePatchStatesResponseBodyInstancePatchStates {
	s.PatchGroup = &v
	return s
}

func (s *ListInstancePatchStatesResponseBodyInstancePatchStates) SetInstalledCount(v string) *ListInstancePatchStatesResponseBodyInstancePatchStates {
	s.InstalledCount = &v
	return s
}

type ListInstancePatchStatesResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListInstancePatchStatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListInstancePatchStatesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstancePatchStatesResponse) GoString() string {
	return s.String()
}

func (s *ListInstancePatchStatesResponse) SetHeaders(v map[string]*string) *ListInstancePatchStatesResponse {
	s.Headers = v
	return s
}

func (s *ListInstancePatchStatesResponse) SetBody(v *ListInstancePatchStatesResponseBody) *ListInstancePatchStatesResponse {
	s.Body = v
	return s
}

type ListInstanceStateReportsRequest struct {
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	StateConfigurationId *string `json:"StateConfigurationId,omitempty" xml:"StateConfigurationId,omitempty"`
	MaxResults           *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken            *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListInstanceStateReportsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceStateReportsRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceStateReportsRequest) SetRegionId(v string) *ListInstanceStateReportsRequest {
	s.RegionId = &v
	return s
}

func (s *ListInstanceStateReportsRequest) SetInstanceId(v string) *ListInstanceStateReportsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListInstanceStateReportsRequest) SetStateConfigurationId(v string) *ListInstanceStateReportsRequest {
	s.StateConfigurationId = &v
	return s
}

func (s *ListInstanceStateReportsRequest) SetMaxResults(v int32) *ListInstanceStateReportsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListInstanceStateReportsRequest) SetNextToken(v string) *ListInstanceStateReportsRequest {
	s.NextToken = &v
	return s
}

type ListInstanceStateReportsResponseBody struct {
	NextToken    *string                                             `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId    *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	MaxResults   *int32                                              `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	StateReports []*ListInstanceStateReportsResponseBodyStateReports `json:"StateReports,omitempty" xml:"StateReports,omitempty" type:"Repeated"`
}

func (s ListInstanceStateReportsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceStateReportsResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceStateReportsResponseBody) SetNextToken(v string) *ListInstanceStateReportsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListInstanceStateReportsResponseBody) SetRequestId(v string) *ListInstanceStateReportsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceStateReportsResponseBody) SetMaxResults(v int32) *ListInstanceStateReportsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListInstanceStateReportsResponseBody) SetStateReports(v []*ListInstanceStateReportsResponseBodyStateReports) *ListInstanceStateReportsResponseBody {
	s.StateReports = v
	return s
}

type ListInstanceStateReportsResponseBodyStateReports struct {
	ReportStatus         *string `json:"ReportStatus,omitempty" xml:"ReportStatus,omitempty"`
	ReportInfo           *string `json:"ReportInfo,omitempty" xml:"ReportInfo,omitempty"`
	SuccessApplyTime     *string `json:"SuccessApplyTime,omitempty" xml:"SuccessApplyTime,omitempty"`
	StateConfigurationId *string `json:"StateConfigurationId,omitempty" xml:"StateConfigurationId,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Mode                 *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	ReportTime           *string `json:"ReportTime,omitempty" xml:"ReportTime,omitempty"`
}

func (s ListInstanceStateReportsResponseBodyStateReports) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceStateReportsResponseBodyStateReports) GoString() string {
	return s.String()
}

func (s *ListInstanceStateReportsResponseBodyStateReports) SetReportStatus(v string) *ListInstanceStateReportsResponseBodyStateReports {
	s.ReportStatus = &v
	return s
}

func (s *ListInstanceStateReportsResponseBodyStateReports) SetReportInfo(v string) *ListInstanceStateReportsResponseBodyStateReports {
	s.ReportInfo = &v
	return s
}

func (s *ListInstanceStateReportsResponseBodyStateReports) SetSuccessApplyTime(v string) *ListInstanceStateReportsResponseBodyStateReports {
	s.SuccessApplyTime = &v
	return s
}

func (s *ListInstanceStateReportsResponseBodyStateReports) SetStateConfigurationId(v string) *ListInstanceStateReportsResponseBodyStateReports {
	s.StateConfigurationId = &v
	return s
}

func (s *ListInstanceStateReportsResponseBodyStateReports) SetInstanceId(v string) *ListInstanceStateReportsResponseBodyStateReports {
	s.InstanceId = &v
	return s
}

func (s *ListInstanceStateReportsResponseBodyStateReports) SetMode(v string) *ListInstanceStateReportsResponseBodyStateReports {
	s.Mode = &v
	return s
}

func (s *ListInstanceStateReportsResponseBodyStateReports) SetReportTime(v string) *ListInstanceStateReportsResponseBodyStateReports {
	s.ReportTime = &v
	return s
}

type ListInstanceStateReportsResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListInstanceStateReportsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListInstanceStateReportsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceStateReportsResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceStateReportsResponse) SetHeaders(v map[string]*string) *ListInstanceStateReportsResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceStateReportsResponse) SetBody(v *ListInstanceStateReportsResponseBody) *ListInstanceStateReportsResponse {
	s.Body = v
	return s
}

type ListInventoryEntriesRequest struct {
	InstanceId *string                              `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	TypeName   *string                              `json:"TypeName,omitempty" xml:"TypeName,omitempty"`
	MaxResults *int32                               `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string                              `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Filter     []*ListInventoryEntriesRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
}

func (s ListInventoryEntriesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInventoryEntriesRequest) GoString() string {
	return s.String()
}

func (s *ListInventoryEntriesRequest) SetInstanceId(v string) *ListInventoryEntriesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListInventoryEntriesRequest) SetTypeName(v string) *ListInventoryEntriesRequest {
	s.TypeName = &v
	return s
}

func (s *ListInventoryEntriesRequest) SetMaxResults(v int32) *ListInventoryEntriesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListInventoryEntriesRequest) SetNextToken(v string) *ListInventoryEntriesRequest {
	s.NextToken = &v
	return s
}

func (s *ListInventoryEntriesRequest) SetFilter(v []*ListInventoryEntriesRequestFilter) *ListInventoryEntriesRequest {
	s.Filter = v
	return s
}

type ListInventoryEntriesRequestFilter struct {
	Value    []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
	Operator *string   `json:"Operator,omitempty" xml:"Operator,omitempty"`
	Name     *string   `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListInventoryEntriesRequestFilter) String() string {
	return tea.Prettify(s)
}

func (s ListInventoryEntriesRequestFilter) GoString() string {
	return s.String()
}

func (s *ListInventoryEntriesRequestFilter) SetValue(v []*string) *ListInventoryEntriesRequestFilter {
	s.Value = v
	return s
}

func (s *ListInventoryEntriesRequestFilter) SetOperator(v string) *ListInventoryEntriesRequestFilter {
	s.Operator = &v
	return s
}

func (s *ListInventoryEntriesRequestFilter) SetName(v string) *ListInventoryEntriesRequestFilter {
	s.Name = &v
	return s
}

type ListInventoryEntriesResponseBody struct {
	NextToken     *string                  `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId     *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SchemaVersion *string                  `json:"SchemaVersion,omitempty" xml:"SchemaVersion,omitempty"`
	MaxResults    *int32                   `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	CaptureTime   *string                  `json:"CaptureTime,omitempty" xml:"CaptureTime,omitempty"`
	TypeName      *string                  `json:"TypeName,omitempty" xml:"TypeName,omitempty"`
	InstanceId    *string                  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Entries       []map[string]interface{} `json:"Entries,omitempty" xml:"Entries,omitempty" type:"Repeated"`
}

func (s ListInventoryEntriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInventoryEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInventoryEntriesResponseBody) SetNextToken(v string) *ListInventoryEntriesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListInventoryEntriesResponseBody) SetRequestId(v string) *ListInventoryEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInventoryEntriesResponseBody) SetSchemaVersion(v string) *ListInventoryEntriesResponseBody {
	s.SchemaVersion = &v
	return s
}

func (s *ListInventoryEntriesResponseBody) SetMaxResults(v int32) *ListInventoryEntriesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListInventoryEntriesResponseBody) SetCaptureTime(v string) *ListInventoryEntriesResponseBody {
	s.CaptureTime = &v
	return s
}

func (s *ListInventoryEntriesResponseBody) SetTypeName(v string) *ListInventoryEntriesResponseBody {
	s.TypeName = &v
	return s
}

func (s *ListInventoryEntriesResponseBody) SetInstanceId(v string) *ListInventoryEntriesResponseBody {
	s.InstanceId = &v
	return s
}

func (s *ListInventoryEntriesResponseBody) SetEntries(v []map[string]interface{}) *ListInventoryEntriesResponseBody {
	s.Entries = v
	return s
}

type ListInventoryEntriesResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListInventoryEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListInventoryEntriesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInventoryEntriesResponse) GoString() string {
	return s.String()
}

func (s *ListInventoryEntriesResponse) SetHeaders(v map[string]*string) *ListInventoryEntriesResponse {
	s.Headers = v
	return s
}

func (s *ListInventoryEntriesResponse) SetBody(v *ListInventoryEntriesResponseBody) *ListInventoryEntriesResponse {
	s.Body = v
	return s
}

type ListParametersRequest struct {
	RegionId        *string                `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Name            *string                `json:"Name,omitempty" xml:"Name,omitempty"`
	MaxResults      *int32                 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken       *string                `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	SortField       *string                `json:"SortField,omitempty" xml:"SortField,omitempty"`
	SortOrder       *string                `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	Type            *string                `json:"Type,omitempty" xml:"Type,omitempty"`
	Path            *string                `json:"Path,omitempty" xml:"Path,omitempty"`
	Recursive       *bool                  `json:"Recursive,omitempty" xml:"Recursive,omitempty"`
	Tags            map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	ResourceGroupId *string                `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListParametersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListParametersRequest) GoString() string {
	return s.String()
}

func (s *ListParametersRequest) SetRegionId(v string) *ListParametersRequest {
	s.RegionId = &v
	return s
}

func (s *ListParametersRequest) SetName(v string) *ListParametersRequest {
	s.Name = &v
	return s
}

func (s *ListParametersRequest) SetMaxResults(v int32) *ListParametersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListParametersRequest) SetNextToken(v string) *ListParametersRequest {
	s.NextToken = &v
	return s
}

func (s *ListParametersRequest) SetSortField(v string) *ListParametersRequest {
	s.SortField = &v
	return s
}

func (s *ListParametersRequest) SetSortOrder(v string) *ListParametersRequest {
	s.SortOrder = &v
	return s
}

func (s *ListParametersRequest) SetType(v string) *ListParametersRequest {
	s.Type = &v
	return s
}

func (s *ListParametersRequest) SetPath(v string) *ListParametersRequest {
	s.Path = &v
	return s
}

func (s *ListParametersRequest) SetRecursive(v bool) *ListParametersRequest {
	s.Recursive = &v
	return s
}

func (s *ListParametersRequest) SetTags(v map[string]interface{}) *ListParametersRequest {
	s.Tags = v
	return s
}

func (s *ListParametersRequest) SetResourceGroupId(v string) *ListParametersRequest {
	s.ResourceGroupId = &v
	return s
}

type ListParametersShrinkRequest struct {
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	MaxResults      *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken       *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	SortField       *string `json:"SortField,omitempty" xml:"SortField,omitempty"`
	SortOrder       *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	Type            *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Path            *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Recursive       *bool   `json:"Recursive,omitempty" xml:"Recursive,omitempty"`
	TagsShrink      *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListParametersShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListParametersShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListParametersShrinkRequest) SetRegionId(v string) *ListParametersShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ListParametersShrinkRequest) SetName(v string) *ListParametersShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListParametersShrinkRequest) SetMaxResults(v int32) *ListParametersShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListParametersShrinkRequest) SetNextToken(v string) *ListParametersShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListParametersShrinkRequest) SetSortField(v string) *ListParametersShrinkRequest {
	s.SortField = &v
	return s
}

func (s *ListParametersShrinkRequest) SetSortOrder(v string) *ListParametersShrinkRequest {
	s.SortOrder = &v
	return s
}

func (s *ListParametersShrinkRequest) SetType(v string) *ListParametersShrinkRequest {
	s.Type = &v
	return s
}

func (s *ListParametersShrinkRequest) SetPath(v string) *ListParametersShrinkRequest {
	s.Path = &v
	return s
}

func (s *ListParametersShrinkRequest) SetRecursive(v bool) *ListParametersShrinkRequest {
	s.Recursive = &v
	return s
}

func (s *ListParametersShrinkRequest) SetTagsShrink(v string) *ListParametersShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *ListParametersShrinkRequest) SetResourceGroupId(v string) *ListParametersShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

type ListParametersResponseBody struct {
	NextToken  *string                                 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	MaxResults *int32                                  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Parameters []*ListParametersResponseBodyParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
}

func (s ListParametersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListParametersResponseBody) GoString() string {
	return s.String()
}

func (s *ListParametersResponseBody) SetNextToken(v string) *ListParametersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListParametersResponseBody) SetRequestId(v string) *ListParametersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListParametersResponseBody) SetTotalCount(v int32) *ListParametersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListParametersResponseBody) SetMaxResults(v int32) *ListParametersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListParametersResponseBody) SetParameters(v []*ListParametersResponseBodyParameters) *ListParametersResponseBody {
	s.Parameters = v
	return s
}

type ListParametersResponseBodyParameters struct {
	Type             *string                `json:"Type,omitempty" xml:"Type,omitempty"`
	UpdatedDate      *string                `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
	UpdatedBy        *string                `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	Tags             map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	Description      *string                `json:"Description,omitempty" xml:"Description,omitempty"`
	CreatedBy        *string                `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	ResourceGroupId  *string                `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	CreatedDate      *string                `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	ParameterVersion *string                `json:"ParameterVersion,omitempty" xml:"ParameterVersion,omitempty"`
	Name             *string                `json:"Name,omitempty" xml:"Name,omitempty"`
	Id               *string                `json:"Id,omitempty" xml:"Id,omitempty"`
	ShareType        *string                `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
}

func (s ListParametersResponseBodyParameters) String() string {
	return tea.Prettify(s)
}

func (s ListParametersResponseBodyParameters) GoString() string {
	return s.String()
}

func (s *ListParametersResponseBodyParameters) SetType(v string) *ListParametersResponseBodyParameters {
	s.Type = &v
	return s
}

func (s *ListParametersResponseBodyParameters) SetUpdatedDate(v string) *ListParametersResponseBodyParameters {
	s.UpdatedDate = &v
	return s
}

func (s *ListParametersResponseBodyParameters) SetUpdatedBy(v string) *ListParametersResponseBodyParameters {
	s.UpdatedBy = &v
	return s
}

func (s *ListParametersResponseBodyParameters) SetTags(v map[string]interface{}) *ListParametersResponseBodyParameters {
	s.Tags = v
	return s
}

func (s *ListParametersResponseBodyParameters) SetDescription(v string) *ListParametersResponseBodyParameters {
	s.Description = &v
	return s
}

func (s *ListParametersResponseBodyParameters) SetCreatedBy(v string) *ListParametersResponseBodyParameters {
	s.CreatedBy = &v
	return s
}

func (s *ListParametersResponseBodyParameters) SetResourceGroupId(v string) *ListParametersResponseBodyParameters {
	s.ResourceGroupId = &v
	return s
}

func (s *ListParametersResponseBodyParameters) SetCreatedDate(v string) *ListParametersResponseBodyParameters {
	s.CreatedDate = &v
	return s
}

func (s *ListParametersResponseBodyParameters) SetParameterVersion(v string) *ListParametersResponseBodyParameters {
	s.ParameterVersion = &v
	return s
}

func (s *ListParametersResponseBodyParameters) SetName(v string) *ListParametersResponseBodyParameters {
	s.Name = &v
	return s
}

func (s *ListParametersResponseBodyParameters) SetId(v string) *ListParametersResponseBodyParameters {
	s.Id = &v
	return s
}

func (s *ListParametersResponseBodyParameters) SetShareType(v string) *ListParametersResponseBodyParameters {
	s.ShareType = &v
	return s
}

type ListParametersResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListParametersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListParametersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListParametersResponse) GoString() string {
	return s.String()
}

func (s *ListParametersResponse) SetHeaders(v map[string]*string) *ListParametersResponse {
	s.Headers = v
	return s
}

func (s *ListParametersResponse) SetBody(v *ListParametersResponseBody) *ListParametersResponse {
	s.Body = v
	return s
}

type ListParameterVersionsRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ShareType  *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
}

func (s ListParameterVersionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListParameterVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListParameterVersionsRequest) SetRegionId(v string) *ListParameterVersionsRequest {
	s.RegionId = &v
	return s
}

func (s *ListParameterVersionsRequest) SetName(v string) *ListParameterVersionsRequest {
	s.Name = &v
	return s
}

func (s *ListParameterVersionsRequest) SetMaxResults(v int32) *ListParameterVersionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListParameterVersionsRequest) SetNextToken(v string) *ListParameterVersionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListParameterVersionsRequest) SetShareType(v string) *ListParameterVersionsRequest {
	s.ShareType = &v
	return s
}

type ListParameterVersionsResponseBody struct {
	Type              *string                                               `json:"Type,omitempty" xml:"Type,omitempty"`
	NextToken         *string                                               `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId         *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Description       *string                                               `json:"Description,omitempty" xml:"Description,omitempty"`
	MaxResults        *int32                                                `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	CreatedBy         *string                                               `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	CreatedDate       *string                                               `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	Name              *string                                               `json:"Name,omitempty" xml:"Name,omitempty"`
	TotalCount        *int32                                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Id                *string                                               `json:"Id,omitempty" xml:"Id,omitempty"`
	ParameterVersions []*ListParameterVersionsResponseBodyParameterVersions `json:"ParameterVersions,omitempty" xml:"ParameterVersions,omitempty" type:"Repeated"`
}

func (s ListParameterVersionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListParameterVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListParameterVersionsResponseBody) SetType(v string) *ListParameterVersionsResponseBody {
	s.Type = &v
	return s
}

func (s *ListParameterVersionsResponseBody) SetNextToken(v string) *ListParameterVersionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListParameterVersionsResponseBody) SetRequestId(v string) *ListParameterVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListParameterVersionsResponseBody) SetDescription(v string) *ListParameterVersionsResponseBody {
	s.Description = &v
	return s
}

func (s *ListParameterVersionsResponseBody) SetMaxResults(v int32) *ListParameterVersionsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListParameterVersionsResponseBody) SetCreatedBy(v string) *ListParameterVersionsResponseBody {
	s.CreatedBy = &v
	return s
}

func (s *ListParameterVersionsResponseBody) SetCreatedDate(v string) *ListParameterVersionsResponseBody {
	s.CreatedDate = &v
	return s
}

func (s *ListParameterVersionsResponseBody) SetName(v string) *ListParameterVersionsResponseBody {
	s.Name = &v
	return s
}

func (s *ListParameterVersionsResponseBody) SetTotalCount(v int32) *ListParameterVersionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListParameterVersionsResponseBody) SetId(v string) *ListParameterVersionsResponseBody {
	s.Id = &v
	return s
}

func (s *ListParameterVersionsResponseBody) SetParameterVersions(v []*ListParameterVersionsResponseBodyParameterVersions) *ListParameterVersionsResponseBody {
	s.ParameterVersions = v
	return s
}

type ListParameterVersionsResponseBodyParameterVersions struct {
	ParameterVersion *int32  `json:"ParameterVersion,omitempty" xml:"ParameterVersion,omitempty"`
	Value            *string `json:"Value,omitempty" xml:"Value,omitempty"`
	UpdatedDate      *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
	UpdatedBy        *string `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
}

func (s ListParameterVersionsResponseBodyParameterVersions) String() string {
	return tea.Prettify(s)
}

func (s ListParameterVersionsResponseBodyParameterVersions) GoString() string {
	return s.String()
}

func (s *ListParameterVersionsResponseBodyParameterVersions) SetParameterVersion(v int32) *ListParameterVersionsResponseBodyParameterVersions {
	s.ParameterVersion = &v
	return s
}

func (s *ListParameterVersionsResponseBodyParameterVersions) SetValue(v string) *ListParameterVersionsResponseBodyParameterVersions {
	s.Value = &v
	return s
}

func (s *ListParameterVersionsResponseBodyParameterVersions) SetUpdatedDate(v string) *ListParameterVersionsResponseBodyParameterVersions {
	s.UpdatedDate = &v
	return s
}

func (s *ListParameterVersionsResponseBodyParameterVersions) SetUpdatedBy(v string) *ListParameterVersionsResponseBodyParameterVersions {
	s.UpdatedBy = &v
	return s
}

type ListParameterVersionsResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListParameterVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListParameterVersionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListParameterVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListParameterVersionsResponse) SetHeaders(v map[string]*string) *ListParameterVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListParameterVersionsResponse) SetBody(v *ListParameterVersionsResponseBody) *ListParameterVersionsResponse {
	s.Body = v
	return s
}

type ListPatchBaselinesRequest struct {
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OperationSystem *string `json:"OperationSystem,omitempty" xml:"OperationSystem,omitempty"`
	ShareType       *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	MaxResults      *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken       *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListPatchBaselinesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPatchBaselinesRequest) GoString() string {
	return s.String()
}

func (s *ListPatchBaselinesRequest) SetRegionId(v string) *ListPatchBaselinesRequest {
	s.RegionId = &v
	return s
}

func (s *ListPatchBaselinesRequest) SetName(v string) *ListPatchBaselinesRequest {
	s.Name = &v
	return s
}

func (s *ListPatchBaselinesRequest) SetOperationSystem(v string) *ListPatchBaselinesRequest {
	s.OperationSystem = &v
	return s
}

func (s *ListPatchBaselinesRequest) SetShareType(v string) *ListPatchBaselinesRequest {
	s.ShareType = &v
	return s
}

func (s *ListPatchBaselinesRequest) SetMaxResults(v int32) *ListPatchBaselinesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPatchBaselinesRequest) SetNextToken(v string) *ListPatchBaselinesRequest {
	s.NextToken = &v
	return s
}

type ListPatchBaselinesResponseBody struct {
	NextToken      *string                                         `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId      *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	MaxResults     *int32                                          `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	PatchBaselines []*ListPatchBaselinesResponseBodyPatchBaselines `json:"PatchBaselines,omitempty" xml:"PatchBaselines,omitempty" type:"Repeated"`
}

func (s ListPatchBaselinesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPatchBaselinesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPatchBaselinesResponseBody) SetNextToken(v string) *ListPatchBaselinesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPatchBaselinesResponseBody) SetRequestId(v string) *ListPatchBaselinesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPatchBaselinesResponseBody) SetMaxResults(v int32) *ListPatchBaselinesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListPatchBaselinesResponseBody) SetPatchBaselines(v []*ListPatchBaselinesResponseBodyPatchBaselines) *ListPatchBaselinesResponseBody {
	s.PatchBaselines = v
	return s
}

type ListPatchBaselinesResponseBodyPatchBaselines struct {
	OperationSystem *string `json:"OperationSystem,omitempty" xml:"OperationSystem,omitempty"`
	IsDefault       *bool   `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	UpdatedDate     *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
	UpdatedBy       *string `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	CreatedBy       *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	CreatedDate     *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id              *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ShareType       *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
}

func (s ListPatchBaselinesResponseBodyPatchBaselines) String() string {
	return tea.Prettify(s)
}

func (s ListPatchBaselinesResponseBodyPatchBaselines) GoString() string {
	return s.String()
}

func (s *ListPatchBaselinesResponseBodyPatchBaselines) SetOperationSystem(v string) *ListPatchBaselinesResponseBodyPatchBaselines {
	s.OperationSystem = &v
	return s
}

func (s *ListPatchBaselinesResponseBodyPatchBaselines) SetIsDefault(v bool) *ListPatchBaselinesResponseBodyPatchBaselines {
	s.IsDefault = &v
	return s
}

func (s *ListPatchBaselinesResponseBodyPatchBaselines) SetDescription(v string) *ListPatchBaselinesResponseBodyPatchBaselines {
	s.Description = &v
	return s
}

func (s *ListPatchBaselinesResponseBodyPatchBaselines) SetUpdatedDate(v string) *ListPatchBaselinesResponseBodyPatchBaselines {
	s.UpdatedDate = &v
	return s
}

func (s *ListPatchBaselinesResponseBodyPatchBaselines) SetUpdatedBy(v string) *ListPatchBaselinesResponseBodyPatchBaselines {
	s.UpdatedBy = &v
	return s
}

func (s *ListPatchBaselinesResponseBodyPatchBaselines) SetCreatedBy(v string) *ListPatchBaselinesResponseBodyPatchBaselines {
	s.CreatedBy = &v
	return s
}

func (s *ListPatchBaselinesResponseBodyPatchBaselines) SetCreatedDate(v string) *ListPatchBaselinesResponseBodyPatchBaselines {
	s.CreatedDate = &v
	return s
}

func (s *ListPatchBaselinesResponseBodyPatchBaselines) SetName(v string) *ListPatchBaselinesResponseBodyPatchBaselines {
	s.Name = &v
	return s
}

func (s *ListPatchBaselinesResponseBodyPatchBaselines) SetId(v string) *ListPatchBaselinesResponseBodyPatchBaselines {
	s.Id = &v
	return s
}

func (s *ListPatchBaselinesResponseBodyPatchBaselines) SetShareType(v string) *ListPatchBaselinesResponseBodyPatchBaselines {
	s.ShareType = &v
	return s
}

type ListPatchBaselinesResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListPatchBaselinesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPatchBaselinesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPatchBaselinesResponse) GoString() string {
	return s.String()
}

func (s *ListPatchBaselinesResponse) SetHeaders(v map[string]*string) *ListPatchBaselinesResponse {
	s.Headers = v
	return s
}

func (s *ListPatchBaselinesResponse) SetBody(v *ListPatchBaselinesResponseBody) *ListPatchBaselinesResponse {
	s.Body = v
	return s
}

type ListResourceExecutionStatusRequest struct {
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ExecutionId *string `json:"ExecutionId,omitempty" xml:"ExecutionId,omitempty"`
	MaxResults  *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken   *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListResourceExecutionStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResourceExecutionStatusRequest) GoString() string {
	return s.String()
}

func (s *ListResourceExecutionStatusRequest) SetRegionId(v string) *ListResourceExecutionStatusRequest {
	s.RegionId = &v
	return s
}

func (s *ListResourceExecutionStatusRequest) SetExecutionId(v string) *ListResourceExecutionStatusRequest {
	s.ExecutionId = &v
	return s
}

func (s *ListResourceExecutionStatusRequest) SetMaxResults(v int32) *ListResourceExecutionStatusRequest {
	s.MaxResults = &v
	return s
}

func (s *ListResourceExecutionStatusRequest) SetNextToken(v string) *ListResourceExecutionStatusRequest {
	s.NextToken = &v
	return s
}

type ListResourceExecutionStatusResponseBody struct {
	NextToken               *string                                                           `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId               *string                                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	MaxResults              *int32                                                            `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	ResourceExecutionStatus []*ListResourceExecutionStatusResponseBodyResourceExecutionStatus `json:"ResourceExecutionStatus,omitempty" xml:"ResourceExecutionStatus,omitempty" type:"Repeated"`
}

func (s ListResourceExecutionStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListResourceExecutionStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceExecutionStatusResponseBody) SetNextToken(v string) *ListResourceExecutionStatusResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListResourceExecutionStatusResponseBody) SetRequestId(v string) *ListResourceExecutionStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourceExecutionStatusResponseBody) SetMaxResults(v int32) *ListResourceExecutionStatusResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListResourceExecutionStatusResponseBody) SetResourceExecutionStatus(v []*ListResourceExecutionStatusResponseBodyResourceExecutionStatus) *ListResourceExecutionStatusResponseBody {
	s.ResourceExecutionStatus = v
	return s
}

type ListResourceExecutionStatusResponseBodyResourceExecutionStatus struct {
	Outputs       *string `json:"Outputs,omitempty" xml:"Outputs,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	ExecutionTime *string `json:"ExecutionTime,omitempty" xml:"ExecutionTime,omitempty"`
	ResourceId    *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ExecutionId   *string `json:"ExecutionId,omitempty" xml:"ExecutionId,omitempty"`
}

func (s ListResourceExecutionStatusResponseBodyResourceExecutionStatus) String() string {
	return tea.Prettify(s)
}

func (s ListResourceExecutionStatusResponseBodyResourceExecutionStatus) GoString() string {
	return s.String()
}

func (s *ListResourceExecutionStatusResponseBodyResourceExecutionStatus) SetOutputs(v string) *ListResourceExecutionStatusResponseBodyResourceExecutionStatus {
	s.Outputs = &v
	return s
}

func (s *ListResourceExecutionStatusResponseBodyResourceExecutionStatus) SetStatus(v string) *ListResourceExecutionStatusResponseBodyResourceExecutionStatus {
	s.Status = &v
	return s
}

func (s *ListResourceExecutionStatusResponseBodyResourceExecutionStatus) SetExecutionTime(v string) *ListResourceExecutionStatusResponseBodyResourceExecutionStatus {
	s.ExecutionTime = &v
	return s
}

func (s *ListResourceExecutionStatusResponseBodyResourceExecutionStatus) SetResourceId(v string) *ListResourceExecutionStatusResponseBodyResourceExecutionStatus {
	s.ResourceId = &v
	return s
}

func (s *ListResourceExecutionStatusResponseBodyResourceExecutionStatus) SetExecutionId(v string) *ListResourceExecutionStatusResponseBodyResourceExecutionStatus {
	s.ExecutionId = &v
	return s
}

type ListResourceExecutionStatusResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListResourceExecutionStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListResourceExecutionStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ListResourceExecutionStatusResponse) GoString() string {
	return s.String()
}

func (s *ListResourceExecutionStatusResponse) SetHeaders(v map[string]*string) *ListResourceExecutionStatusResponse {
	s.Headers = v
	return s
}

func (s *ListResourceExecutionStatusResponse) SetBody(v *ListResourceExecutionStatusResponseBody) *ListResourceExecutionStatusResponse {
	s.Body = v
	return s
}

type ListSecretParametersRequest struct {
	RegionId        *string                `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Name            *string                `json:"Name,omitempty" xml:"Name,omitempty"`
	MaxResults      *int32                 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken       *string                `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	SortField       *string                `json:"SortField,omitempty" xml:"SortField,omitempty"`
	SortOrder       *string                `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	Path            *string                `json:"Path,omitempty" xml:"Path,omitempty"`
	Recursive       *bool                  `json:"Recursive,omitempty" xml:"Recursive,omitempty"`
	Tags            map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	ResourceGroupId *string                `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListSecretParametersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSecretParametersRequest) GoString() string {
	return s.String()
}

func (s *ListSecretParametersRequest) SetRegionId(v string) *ListSecretParametersRequest {
	s.RegionId = &v
	return s
}

func (s *ListSecretParametersRequest) SetName(v string) *ListSecretParametersRequest {
	s.Name = &v
	return s
}

func (s *ListSecretParametersRequest) SetMaxResults(v int32) *ListSecretParametersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSecretParametersRequest) SetNextToken(v string) *ListSecretParametersRequest {
	s.NextToken = &v
	return s
}

func (s *ListSecretParametersRequest) SetSortField(v string) *ListSecretParametersRequest {
	s.SortField = &v
	return s
}

func (s *ListSecretParametersRequest) SetSortOrder(v string) *ListSecretParametersRequest {
	s.SortOrder = &v
	return s
}

func (s *ListSecretParametersRequest) SetPath(v string) *ListSecretParametersRequest {
	s.Path = &v
	return s
}

func (s *ListSecretParametersRequest) SetRecursive(v bool) *ListSecretParametersRequest {
	s.Recursive = &v
	return s
}

func (s *ListSecretParametersRequest) SetTags(v map[string]interface{}) *ListSecretParametersRequest {
	s.Tags = v
	return s
}

func (s *ListSecretParametersRequest) SetResourceGroupId(v string) *ListSecretParametersRequest {
	s.ResourceGroupId = &v
	return s
}

type ListSecretParametersShrinkRequest struct {
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	MaxResults      *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken       *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	SortField       *string `json:"SortField,omitempty" xml:"SortField,omitempty"`
	SortOrder       *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	Path            *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Recursive       *bool   `json:"Recursive,omitempty" xml:"Recursive,omitempty"`
	TagsShrink      *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListSecretParametersShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSecretParametersShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListSecretParametersShrinkRequest) SetRegionId(v string) *ListSecretParametersShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ListSecretParametersShrinkRequest) SetName(v string) *ListSecretParametersShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListSecretParametersShrinkRequest) SetMaxResults(v int32) *ListSecretParametersShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSecretParametersShrinkRequest) SetNextToken(v string) *ListSecretParametersShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListSecretParametersShrinkRequest) SetSortField(v string) *ListSecretParametersShrinkRequest {
	s.SortField = &v
	return s
}

func (s *ListSecretParametersShrinkRequest) SetSortOrder(v string) *ListSecretParametersShrinkRequest {
	s.SortOrder = &v
	return s
}

func (s *ListSecretParametersShrinkRequest) SetPath(v string) *ListSecretParametersShrinkRequest {
	s.Path = &v
	return s
}

func (s *ListSecretParametersShrinkRequest) SetRecursive(v bool) *ListSecretParametersShrinkRequest {
	s.Recursive = &v
	return s
}

func (s *ListSecretParametersShrinkRequest) SetTagsShrink(v string) *ListSecretParametersShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *ListSecretParametersShrinkRequest) SetResourceGroupId(v string) *ListSecretParametersShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

type ListSecretParametersResponseBody struct {
	NextToken  *string                                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId  *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	MaxResults *int32                                        `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Parameters []*ListSecretParametersResponseBodyParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
}

func (s ListSecretParametersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSecretParametersResponseBody) GoString() string {
	return s.String()
}

func (s *ListSecretParametersResponseBody) SetNextToken(v string) *ListSecretParametersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListSecretParametersResponseBody) SetRequestId(v string) *ListSecretParametersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSecretParametersResponseBody) SetMaxResults(v int32) *ListSecretParametersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListSecretParametersResponseBody) SetParameters(v []*ListSecretParametersResponseBodyParameters) *ListSecretParametersResponseBody {
	s.Parameters = v
	return s
}

type ListSecretParametersResponseBodyParameters struct {
	Type             *string                `json:"Type,omitempty" xml:"Type,omitempty"`
	UpdatedDate      *string                `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
	UpdatedBy        *string                `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	KeyId            *string                `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	Tags             map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	Description      *string                `json:"Description,omitempty" xml:"Description,omitempty"`
	ResourceGroupId  *string                `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	CreatedBy        *string                `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	CreatedDate      *string                `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	ParameterVersion *string                `json:"ParameterVersion,omitempty" xml:"ParameterVersion,omitempty"`
	Name             *string                `json:"Name,omitempty" xml:"Name,omitempty"`
	Id               *string                `json:"Id,omitempty" xml:"Id,omitempty"`
	ShareType        *string                `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
}

func (s ListSecretParametersResponseBodyParameters) String() string {
	return tea.Prettify(s)
}

func (s ListSecretParametersResponseBodyParameters) GoString() string {
	return s.String()
}

func (s *ListSecretParametersResponseBodyParameters) SetType(v string) *ListSecretParametersResponseBodyParameters {
	s.Type = &v
	return s
}

func (s *ListSecretParametersResponseBodyParameters) SetUpdatedDate(v string) *ListSecretParametersResponseBodyParameters {
	s.UpdatedDate = &v
	return s
}

func (s *ListSecretParametersResponseBodyParameters) SetUpdatedBy(v string) *ListSecretParametersResponseBodyParameters {
	s.UpdatedBy = &v
	return s
}

func (s *ListSecretParametersResponseBodyParameters) SetKeyId(v string) *ListSecretParametersResponseBodyParameters {
	s.KeyId = &v
	return s
}

func (s *ListSecretParametersResponseBodyParameters) SetTags(v map[string]interface{}) *ListSecretParametersResponseBodyParameters {
	s.Tags = v
	return s
}

func (s *ListSecretParametersResponseBodyParameters) SetDescription(v string) *ListSecretParametersResponseBodyParameters {
	s.Description = &v
	return s
}

func (s *ListSecretParametersResponseBodyParameters) SetResourceGroupId(v string) *ListSecretParametersResponseBodyParameters {
	s.ResourceGroupId = &v
	return s
}

func (s *ListSecretParametersResponseBodyParameters) SetCreatedBy(v string) *ListSecretParametersResponseBodyParameters {
	s.CreatedBy = &v
	return s
}

func (s *ListSecretParametersResponseBodyParameters) SetCreatedDate(v string) *ListSecretParametersResponseBodyParameters {
	s.CreatedDate = &v
	return s
}

func (s *ListSecretParametersResponseBodyParameters) SetParameterVersion(v string) *ListSecretParametersResponseBodyParameters {
	s.ParameterVersion = &v
	return s
}

func (s *ListSecretParametersResponseBodyParameters) SetName(v string) *ListSecretParametersResponseBodyParameters {
	s.Name = &v
	return s
}

func (s *ListSecretParametersResponseBodyParameters) SetId(v string) *ListSecretParametersResponseBodyParameters {
	s.Id = &v
	return s
}

func (s *ListSecretParametersResponseBodyParameters) SetShareType(v string) *ListSecretParametersResponseBodyParameters {
	s.ShareType = &v
	return s
}

type ListSecretParametersResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSecretParametersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSecretParametersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSecretParametersResponse) GoString() string {
	return s.String()
}

func (s *ListSecretParametersResponse) SetHeaders(v map[string]*string) *ListSecretParametersResponse {
	s.Headers = v
	return s
}

func (s *ListSecretParametersResponse) SetBody(v *ListSecretParametersResponseBody) *ListSecretParametersResponse {
	s.Body = v
	return s
}

type ListSecretParameterVersionsRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	MaxResults     *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken      *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ShareType      *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	WithDecryption *bool   `json:"WithDecryption,omitempty" xml:"WithDecryption,omitempty"`
}

func (s ListSecretParameterVersionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSecretParameterVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListSecretParameterVersionsRequest) SetRegionId(v string) *ListSecretParameterVersionsRequest {
	s.RegionId = &v
	return s
}

func (s *ListSecretParameterVersionsRequest) SetName(v string) *ListSecretParameterVersionsRequest {
	s.Name = &v
	return s
}

func (s *ListSecretParameterVersionsRequest) SetMaxResults(v int32) *ListSecretParameterVersionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSecretParameterVersionsRequest) SetNextToken(v string) *ListSecretParameterVersionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListSecretParameterVersionsRequest) SetShareType(v string) *ListSecretParameterVersionsRequest {
	s.ShareType = &v
	return s
}

func (s *ListSecretParameterVersionsRequest) SetWithDecryption(v bool) *ListSecretParameterVersionsRequest {
	s.WithDecryption = &v
	return s
}

type ListSecretParameterVersionsResponseBody struct {
	Type              *string                                                     `json:"Type,omitempty" xml:"Type,omitempty"`
	NextToken         *string                                                     `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId         *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Description       *string                                                     `json:"Description,omitempty" xml:"Description,omitempty"`
	MaxResults        *int32                                                      `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	CreatedBy         *string                                                     `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	CreatedDate       *string                                                     `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	Name              *string                                                     `json:"Name,omitempty" xml:"Name,omitempty"`
	TotalCount        *int32                                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Id                *string                                                     `json:"Id,omitempty" xml:"Id,omitempty"`
	ParameterVersions []*ListSecretParameterVersionsResponseBodyParameterVersions `json:"ParameterVersions,omitempty" xml:"ParameterVersions,omitempty" type:"Repeated"`
}

func (s ListSecretParameterVersionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSecretParameterVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSecretParameterVersionsResponseBody) SetType(v string) *ListSecretParameterVersionsResponseBody {
	s.Type = &v
	return s
}

func (s *ListSecretParameterVersionsResponseBody) SetNextToken(v string) *ListSecretParameterVersionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListSecretParameterVersionsResponseBody) SetRequestId(v string) *ListSecretParameterVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSecretParameterVersionsResponseBody) SetDescription(v string) *ListSecretParameterVersionsResponseBody {
	s.Description = &v
	return s
}

func (s *ListSecretParameterVersionsResponseBody) SetMaxResults(v int32) *ListSecretParameterVersionsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListSecretParameterVersionsResponseBody) SetCreatedBy(v string) *ListSecretParameterVersionsResponseBody {
	s.CreatedBy = &v
	return s
}

func (s *ListSecretParameterVersionsResponseBody) SetCreatedDate(v string) *ListSecretParameterVersionsResponseBody {
	s.CreatedDate = &v
	return s
}

func (s *ListSecretParameterVersionsResponseBody) SetName(v string) *ListSecretParameterVersionsResponseBody {
	s.Name = &v
	return s
}

func (s *ListSecretParameterVersionsResponseBody) SetTotalCount(v int32) *ListSecretParameterVersionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSecretParameterVersionsResponseBody) SetId(v string) *ListSecretParameterVersionsResponseBody {
	s.Id = &v
	return s
}

func (s *ListSecretParameterVersionsResponseBody) SetParameterVersions(v []*ListSecretParameterVersionsResponseBodyParameterVersions) *ListSecretParameterVersionsResponseBody {
	s.ParameterVersions = v
	return s
}

type ListSecretParameterVersionsResponseBodyParameterVersions struct {
	ParameterVersion *int32  `json:"ParameterVersion,omitempty" xml:"ParameterVersion,omitempty"`
	Value            *string `json:"Value,omitempty" xml:"Value,omitempty"`
	UpdatedDate      *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
	UpdatedBy        *string `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
}

func (s ListSecretParameterVersionsResponseBodyParameterVersions) String() string {
	return tea.Prettify(s)
}

func (s ListSecretParameterVersionsResponseBodyParameterVersions) GoString() string {
	return s.String()
}

func (s *ListSecretParameterVersionsResponseBodyParameterVersions) SetParameterVersion(v int32) *ListSecretParameterVersionsResponseBodyParameterVersions {
	s.ParameterVersion = &v
	return s
}

func (s *ListSecretParameterVersionsResponseBodyParameterVersions) SetValue(v string) *ListSecretParameterVersionsResponseBodyParameterVersions {
	s.Value = &v
	return s
}

func (s *ListSecretParameterVersionsResponseBodyParameterVersions) SetUpdatedDate(v string) *ListSecretParameterVersionsResponseBodyParameterVersions {
	s.UpdatedDate = &v
	return s
}

func (s *ListSecretParameterVersionsResponseBodyParameterVersions) SetUpdatedBy(v string) *ListSecretParameterVersionsResponseBodyParameterVersions {
	s.UpdatedBy = &v
	return s
}

type ListSecretParameterVersionsResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSecretParameterVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSecretParameterVersionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSecretParameterVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListSecretParameterVersionsResponse) SetHeaders(v map[string]*string) *ListSecretParameterVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListSecretParameterVersionsResponse) SetBody(v *ListSecretParameterVersionsResponseBody) *ListSecretParameterVersionsResponse {
	s.Body = v
	return s
}

type ListStateConfigurationsRequest struct {
	RegionId              *string                `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TemplateName          *string                `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateVersion       *string                `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	StateConfigurationIds *string                `json:"StateConfigurationIds,omitempty" xml:"StateConfigurationIds,omitempty"`
	Tags                  map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	MaxResults            *int32                 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken             *string                `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ResourceGroupId       *string                `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListStateConfigurationsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListStateConfigurationsRequest) GoString() string {
	return s.String()
}

func (s *ListStateConfigurationsRequest) SetRegionId(v string) *ListStateConfigurationsRequest {
	s.RegionId = &v
	return s
}

func (s *ListStateConfigurationsRequest) SetTemplateName(v string) *ListStateConfigurationsRequest {
	s.TemplateName = &v
	return s
}

func (s *ListStateConfigurationsRequest) SetTemplateVersion(v string) *ListStateConfigurationsRequest {
	s.TemplateVersion = &v
	return s
}

func (s *ListStateConfigurationsRequest) SetStateConfigurationIds(v string) *ListStateConfigurationsRequest {
	s.StateConfigurationIds = &v
	return s
}

func (s *ListStateConfigurationsRequest) SetTags(v map[string]interface{}) *ListStateConfigurationsRequest {
	s.Tags = v
	return s
}

func (s *ListStateConfigurationsRequest) SetMaxResults(v int32) *ListStateConfigurationsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListStateConfigurationsRequest) SetNextToken(v string) *ListStateConfigurationsRequest {
	s.NextToken = &v
	return s
}

func (s *ListStateConfigurationsRequest) SetResourceGroupId(v string) *ListStateConfigurationsRequest {
	s.ResourceGroupId = &v
	return s
}

type ListStateConfigurationsShrinkRequest struct {
	RegionId              *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TemplateName          *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateVersion       *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	StateConfigurationIds *string `json:"StateConfigurationIds,omitempty" xml:"StateConfigurationIds,omitempty"`
	TagsShrink            *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	MaxResults            *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken             *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ResourceGroupId       *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListStateConfigurationsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListStateConfigurationsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListStateConfigurationsShrinkRequest) SetRegionId(v string) *ListStateConfigurationsShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ListStateConfigurationsShrinkRequest) SetTemplateName(v string) *ListStateConfigurationsShrinkRequest {
	s.TemplateName = &v
	return s
}

func (s *ListStateConfigurationsShrinkRequest) SetTemplateVersion(v string) *ListStateConfigurationsShrinkRequest {
	s.TemplateVersion = &v
	return s
}

func (s *ListStateConfigurationsShrinkRequest) SetStateConfigurationIds(v string) *ListStateConfigurationsShrinkRequest {
	s.StateConfigurationIds = &v
	return s
}

func (s *ListStateConfigurationsShrinkRequest) SetTagsShrink(v string) *ListStateConfigurationsShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *ListStateConfigurationsShrinkRequest) SetMaxResults(v int32) *ListStateConfigurationsShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListStateConfigurationsShrinkRequest) SetNextToken(v string) *ListStateConfigurationsShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListStateConfigurationsShrinkRequest) SetResourceGroupId(v string) *ListStateConfigurationsShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

type ListStateConfigurationsResponseBody struct {
	RequestId           *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StateConfigurations []*ListStateConfigurationsResponseBodyStateConfigurations `json:"StateConfigurations,omitempty" xml:"StateConfigurations,omitempty" type:"Repeated"`
}

func (s ListStateConfigurationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListStateConfigurationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListStateConfigurationsResponseBody) SetRequestId(v string) *ListStateConfigurationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListStateConfigurationsResponseBody) SetStateConfigurations(v []*ListStateConfigurationsResponseBodyStateConfigurations) *ListStateConfigurationsResponseBody {
	s.StateConfigurations = v
	return s
}

type ListStateConfigurationsResponseBodyStateConfigurations struct {
	UpdateTime           *string                `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	CreateTime           *string                `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Targets              *string                `json:"Targets,omitempty" xml:"Targets,omitempty"`
	Tags                 map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	StateConfigurationId *string                `json:"StateConfigurationId,omitempty" xml:"StateConfigurationId,omitempty"`
	ScheduleExpression   *string                `json:"ScheduleExpression,omitempty" xml:"ScheduleExpression,omitempty"`
	TemplateName         *string                `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateVersion      *string                `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	ConfigureMode        *string                `json:"ConfigureMode,omitempty" xml:"ConfigureMode,omitempty"`
	ScheduleType         *string                `json:"ScheduleType,omitempty" xml:"ScheduleType,omitempty"`
	Parameters           *string                `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	Description          *string                `json:"Description,omitempty" xml:"Description,omitempty"`
	ResourceGroupId      *string                `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	TemplateId           *string                `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s ListStateConfigurationsResponseBodyStateConfigurations) String() string {
	return tea.Prettify(s)
}

func (s ListStateConfigurationsResponseBodyStateConfigurations) GoString() string {
	return s.String()
}

func (s *ListStateConfigurationsResponseBodyStateConfigurations) SetUpdateTime(v string) *ListStateConfigurationsResponseBodyStateConfigurations {
	s.UpdateTime = &v
	return s
}

func (s *ListStateConfigurationsResponseBodyStateConfigurations) SetCreateTime(v string) *ListStateConfigurationsResponseBodyStateConfigurations {
	s.CreateTime = &v
	return s
}

func (s *ListStateConfigurationsResponseBodyStateConfigurations) SetTargets(v string) *ListStateConfigurationsResponseBodyStateConfigurations {
	s.Targets = &v
	return s
}

func (s *ListStateConfigurationsResponseBodyStateConfigurations) SetTags(v map[string]interface{}) *ListStateConfigurationsResponseBodyStateConfigurations {
	s.Tags = v
	return s
}

func (s *ListStateConfigurationsResponseBodyStateConfigurations) SetStateConfigurationId(v string) *ListStateConfigurationsResponseBodyStateConfigurations {
	s.StateConfigurationId = &v
	return s
}

func (s *ListStateConfigurationsResponseBodyStateConfigurations) SetScheduleExpression(v string) *ListStateConfigurationsResponseBodyStateConfigurations {
	s.ScheduleExpression = &v
	return s
}

func (s *ListStateConfigurationsResponseBodyStateConfigurations) SetTemplateName(v string) *ListStateConfigurationsResponseBodyStateConfigurations {
	s.TemplateName = &v
	return s
}

func (s *ListStateConfigurationsResponseBodyStateConfigurations) SetTemplateVersion(v string) *ListStateConfigurationsResponseBodyStateConfigurations {
	s.TemplateVersion = &v
	return s
}

func (s *ListStateConfigurationsResponseBodyStateConfigurations) SetConfigureMode(v string) *ListStateConfigurationsResponseBodyStateConfigurations {
	s.ConfigureMode = &v
	return s
}

func (s *ListStateConfigurationsResponseBodyStateConfigurations) SetScheduleType(v string) *ListStateConfigurationsResponseBodyStateConfigurations {
	s.ScheduleType = &v
	return s
}

func (s *ListStateConfigurationsResponseBodyStateConfigurations) SetParameters(v string) *ListStateConfigurationsResponseBodyStateConfigurations {
	s.Parameters = &v
	return s
}

func (s *ListStateConfigurationsResponseBodyStateConfigurations) SetDescription(v string) *ListStateConfigurationsResponseBodyStateConfigurations {
	s.Description = &v
	return s
}

func (s *ListStateConfigurationsResponseBodyStateConfigurations) SetResourceGroupId(v string) *ListStateConfigurationsResponseBodyStateConfigurations {
	s.ResourceGroupId = &v
	return s
}

func (s *ListStateConfigurationsResponseBodyStateConfigurations) SetTemplateId(v string) *ListStateConfigurationsResponseBodyStateConfigurations {
	s.TemplateId = &v
	return s
}

type ListStateConfigurationsResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListStateConfigurationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListStateConfigurationsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListStateConfigurationsResponse) GoString() string {
	return s.String()
}

func (s *ListStateConfigurationsResponse) SetHeaders(v map[string]*string) *ListStateConfigurationsResponse {
	s.Headers = v
	return s
}

func (s *ListStateConfigurationsResponse) SetBody(v *ListStateConfigurationsResponseBody) *ListStateConfigurationsResponse {
	s.Body = v
	return s
}

type ListTagKeysRequest struct {
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	MaxResults   *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListTagKeysRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysRequest) GoString() string {
	return s.String()
}

func (s *ListTagKeysRequest) SetRegionId(v string) *ListTagKeysRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagKeysRequest) SetMaxResults(v int32) *ListTagKeysRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTagKeysRequest) SetNextToken(v string) *ListTagKeysRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagKeysRequest) SetResourceType(v string) *ListTagKeysRequest {
	s.ResourceType = &v
	return s
}

type ListTagKeysResponseBody struct {
	NextToken  *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId  *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	MaxResults *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Keys       []*string `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Repeated"`
}

func (s ListTagKeysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponseBody) SetNextToken(v string) *ListTagKeysResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagKeysResponseBody) SetRequestId(v string) *ListTagKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagKeysResponseBody) SetMaxResults(v int32) *ListTagKeysResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTagKeysResponseBody) SetKeys(v []*string) *ListTagKeysResponseBody {
	s.Keys = v
	return s
}

type ListTagKeysResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTagKeysResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTagKeysResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponse) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponse) SetHeaders(v map[string]*string) *ListTagKeysResponse {
	s.Headers = v
	return s
}

func (s *ListTagKeysResponse) SetBody(v *ListTagKeysResponseBody) *ListTagKeysResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	RegionId     *string                `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceIds  map[string]interface{} `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty"`
	ResourceType *string                `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tags         map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	NextToken    *string                `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetRegionId(v string) *ListTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceIds(v map[string]interface{}) *ListTagResourcesRequest {
	s.ResourceIds = v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetTags(v map[string]interface{}) *ListTagResourcesRequest {
	s.Tags = v
	return s
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

type ListTagResourcesShrinkRequest struct {
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceIdsShrink *string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty"`
	ResourceType      *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagsShrink        *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	NextToken         *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListTagResourcesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesShrinkRequest) SetRegionId(v string) *ListTagResourcesShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesShrinkRequest) SetResourceIdsShrink(v string) *ListTagResourcesShrinkRequest {
	s.ResourceIdsShrink = &v
	return s
}

func (s *ListTagResourcesShrinkRequest) SetResourceType(v string) *ListTagResourcesShrinkRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesShrinkRequest) SetTagsShrink(v string) *ListTagResourcesShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *ListTagResourcesShrinkRequest) SetNextToken(v string) *ListTagResourcesShrinkRequest {
	s.NextToken = &v
	return s
}

type ListTagResourcesResponseBody struct {
	NextToken    *string                                   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId    *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagResources *ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Struct"`
}

func (s ListTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBody) SetNextToken(v string) *ListTagResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetRequestId(v string) *ListTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetTagResources(v *ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

type ListTagResourcesResponseBodyTagResources struct {
	TagResource []*ListTagResourcesResponseBodyTagResourcesTagResource `json:"TagResource,omitempty" xml:"TagResource,omitempty" type:"Repeated"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagResource(v []*ListTagResourcesResponseBodyTagResourcesTagResource) *ListTagResourcesResponseBodyTagResources {
	s.TagResource = v
	return s
}

type ListTagResourcesResponseBodyTagResourcesTagResource struct {
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagValue     *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceType(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagValue(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagValue = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceId(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagKey(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagKey = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponse) SetHeaders(v map[string]*string) *ListTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type ListTagValuesRequest struct {
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	MaxResults   *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Key          *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s ListTagValuesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesRequest) GoString() string {
	return s.String()
}

func (s *ListTagValuesRequest) SetRegionId(v string) *ListTagValuesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagValuesRequest) SetMaxResults(v int32) *ListTagValuesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTagValuesRequest) SetNextToken(v string) *ListTagValuesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagValuesRequest) SetResourceType(v string) *ListTagValuesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagValuesRequest) SetKey(v string) *ListTagValuesRequest {
	s.Key = &v
	return s
}

type ListTagValuesResponseBody struct {
	NextToken  *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId  *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	MaxResults *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Values     []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ListTagValuesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagValuesResponseBody) SetNextToken(v string) *ListTagValuesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagValuesResponseBody) SetRequestId(v string) *ListTagValuesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagValuesResponseBody) SetMaxResults(v int32) *ListTagValuesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTagValuesResponseBody) SetValues(v []*string) *ListTagValuesResponseBody {
	s.Values = v
	return s
}

type ListTagValuesResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTagValuesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTagValuesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesResponse) GoString() string {
	return s.String()
}

func (s *ListTagValuesResponse) SetHeaders(v map[string]*string) *ListTagValuesResponse {
	s.Headers = v
	return s
}

func (s *ListTagValuesResponse) SetBody(v *ListTagValuesResponseBody) *ListTagValuesResponse {
	s.Body = v
	return s
}

type ListTaskExecutionsRequest struct {
	RegionId                  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ExecutionId               *string `json:"ExecutionId,omitempty" xml:"ExecutionId,omitempty"`
	Status                    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StartDateBefore           *string `json:"StartDateBefore,omitempty" xml:"StartDateBefore,omitempty"`
	StartDateAfter            *string `json:"StartDateAfter,omitempty" xml:"StartDateAfter,omitempty"`
	EndDateBefore             *string `json:"EndDateBefore,omitempty" xml:"EndDateBefore,omitempty"`
	EndDateAfter              *string `json:"EndDateAfter,omitempty" xml:"EndDateAfter,omitempty"`
	TaskExecutionId           *string `json:"TaskExecutionId,omitempty" xml:"TaskExecutionId,omitempty"`
	TaskName                  *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TaskAction                *string `json:"TaskAction,omitempty" xml:"TaskAction,omitempty"`
	ParentTaskExecutionId     *string `json:"ParentTaskExecutionId,omitempty" xml:"ParentTaskExecutionId,omitempty"`
	IncludeChildTaskExecution *bool   `json:"IncludeChildTaskExecution,omitempty" xml:"IncludeChildTaskExecution,omitempty"`
	MaxResults                *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken                 *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	SortField                 *string `json:"SortField,omitempty" xml:"SortField,omitempty"`
	SortOrder                 *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
}

func (s ListTaskExecutionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTaskExecutionsRequest) GoString() string {
	return s.String()
}

func (s *ListTaskExecutionsRequest) SetRegionId(v string) *ListTaskExecutionsRequest {
	s.RegionId = &v
	return s
}

func (s *ListTaskExecutionsRequest) SetExecutionId(v string) *ListTaskExecutionsRequest {
	s.ExecutionId = &v
	return s
}

func (s *ListTaskExecutionsRequest) SetStatus(v string) *ListTaskExecutionsRequest {
	s.Status = &v
	return s
}

func (s *ListTaskExecutionsRequest) SetStartDateBefore(v string) *ListTaskExecutionsRequest {
	s.StartDateBefore = &v
	return s
}

func (s *ListTaskExecutionsRequest) SetStartDateAfter(v string) *ListTaskExecutionsRequest {
	s.StartDateAfter = &v
	return s
}

func (s *ListTaskExecutionsRequest) SetEndDateBefore(v string) *ListTaskExecutionsRequest {
	s.EndDateBefore = &v
	return s
}

func (s *ListTaskExecutionsRequest) SetEndDateAfter(v string) *ListTaskExecutionsRequest {
	s.EndDateAfter = &v
	return s
}

func (s *ListTaskExecutionsRequest) SetTaskExecutionId(v string) *ListTaskExecutionsRequest {
	s.TaskExecutionId = &v
	return s
}

func (s *ListTaskExecutionsRequest) SetTaskName(v string) *ListTaskExecutionsRequest {
	s.TaskName = &v
	return s
}

func (s *ListTaskExecutionsRequest) SetTaskAction(v string) *ListTaskExecutionsRequest {
	s.TaskAction = &v
	return s
}

func (s *ListTaskExecutionsRequest) SetParentTaskExecutionId(v string) *ListTaskExecutionsRequest {
	s.ParentTaskExecutionId = &v
	return s
}

func (s *ListTaskExecutionsRequest) SetIncludeChildTaskExecution(v bool) *ListTaskExecutionsRequest {
	s.IncludeChildTaskExecution = &v
	return s
}

func (s *ListTaskExecutionsRequest) SetMaxResults(v int32) *ListTaskExecutionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTaskExecutionsRequest) SetNextToken(v string) *ListTaskExecutionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListTaskExecutionsRequest) SetSortField(v string) *ListTaskExecutionsRequest {
	s.SortField = &v
	return s
}

func (s *ListTaskExecutionsRequest) SetSortOrder(v string) *ListTaskExecutionsRequest {
	s.SortOrder = &v
	return s
}

type ListTaskExecutionsResponseBody struct {
	NextToken      *string                                         `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId      *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	MaxResults     *int32                                          `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	TaskExecutions []*ListTaskExecutionsResponseBodyTaskExecutions `json:"TaskExecutions,omitempty" xml:"TaskExecutions,omitempty" type:"Repeated"`
}

func (s ListTaskExecutionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTaskExecutionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTaskExecutionsResponseBody) SetNextToken(v string) *ListTaskExecutionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTaskExecutionsResponseBody) SetRequestId(v string) *ListTaskExecutionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTaskExecutionsResponseBody) SetMaxResults(v int32) *ListTaskExecutionsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTaskExecutionsResponseBody) SetTaskExecutions(v []*ListTaskExecutionsResponseBodyTaskExecutions) *ListTaskExecutionsResponseBody {
	s.TaskExecutions = v
	return s
}

type ListTaskExecutionsResponseBodyTaskExecutions struct {
	ChildExecutionId      *string                `json:"ChildExecutionId,omitempty" xml:"ChildExecutionId,omitempty"`
	Outputs               *string                `json:"Outputs,omitempty" xml:"Outputs,omitempty"`
	Status                *string                `json:"Status,omitempty" xml:"Status,omitempty"`
	EndDate               *string                `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	ParentTaskExecutionId *string                `json:"ParentTaskExecutionId,omitempty" xml:"ParentTaskExecutionId,omitempty"`
	TaskName              *string                `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	StartDate             *string                `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	LoopItem              *string                `json:"LoopItem,omitempty" xml:"LoopItem,omitempty"`
	CreateDate            *string                `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	ExecutionId           *string                `json:"ExecutionId,omitempty" xml:"ExecutionId,omitempty"`
	TaskAction            *string                `json:"TaskAction,omitempty" xml:"TaskAction,omitempty"`
	TaskExecutionId       *string                `json:"TaskExecutionId,omitempty" xml:"TaskExecutionId,omitempty"`
	UpdateDate            *string                `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
	Loop                  map[string]interface{} `json:"Loop,omitempty" xml:"Loop,omitempty"`
	TemplateId            *string                `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	LoopBatchNumber       *int32                 `json:"LoopBatchNumber,omitempty" xml:"LoopBatchNumber,omitempty"`
	StatusMessage         *string                `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	ExtraData             map[string]interface{} `json:"ExtraData,omitempty" xml:"ExtraData,omitempty"`
	Properties            *string                `json:"Properties,omitempty" xml:"Properties,omitempty"`
}

func (s ListTaskExecutionsResponseBodyTaskExecutions) String() string {
	return tea.Prettify(s)
}

func (s ListTaskExecutionsResponseBodyTaskExecutions) GoString() string {
	return s.String()
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) SetChildExecutionId(v string) *ListTaskExecutionsResponseBodyTaskExecutions {
	s.ChildExecutionId = &v
	return s
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) SetOutputs(v string) *ListTaskExecutionsResponseBodyTaskExecutions {
	s.Outputs = &v
	return s
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) SetStatus(v string) *ListTaskExecutionsResponseBodyTaskExecutions {
	s.Status = &v
	return s
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) SetEndDate(v string) *ListTaskExecutionsResponseBodyTaskExecutions {
	s.EndDate = &v
	return s
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) SetParentTaskExecutionId(v string) *ListTaskExecutionsResponseBodyTaskExecutions {
	s.ParentTaskExecutionId = &v
	return s
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) SetTaskName(v string) *ListTaskExecutionsResponseBodyTaskExecutions {
	s.TaskName = &v
	return s
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) SetStartDate(v string) *ListTaskExecutionsResponseBodyTaskExecutions {
	s.StartDate = &v
	return s
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) SetLoopItem(v string) *ListTaskExecutionsResponseBodyTaskExecutions {
	s.LoopItem = &v
	return s
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) SetCreateDate(v string) *ListTaskExecutionsResponseBodyTaskExecutions {
	s.CreateDate = &v
	return s
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) SetExecutionId(v string) *ListTaskExecutionsResponseBodyTaskExecutions {
	s.ExecutionId = &v
	return s
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) SetTaskAction(v string) *ListTaskExecutionsResponseBodyTaskExecutions {
	s.TaskAction = &v
	return s
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) SetTaskExecutionId(v string) *ListTaskExecutionsResponseBodyTaskExecutions {
	s.TaskExecutionId = &v
	return s
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) SetUpdateDate(v string) *ListTaskExecutionsResponseBodyTaskExecutions {
	s.UpdateDate = &v
	return s
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) SetLoop(v map[string]interface{}) *ListTaskExecutionsResponseBodyTaskExecutions {
	s.Loop = v
	return s
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) SetTemplateId(v string) *ListTaskExecutionsResponseBodyTaskExecutions {
	s.TemplateId = &v
	return s
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) SetLoopBatchNumber(v int32) *ListTaskExecutionsResponseBodyTaskExecutions {
	s.LoopBatchNumber = &v
	return s
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) SetStatusMessage(v string) *ListTaskExecutionsResponseBodyTaskExecutions {
	s.StatusMessage = &v
	return s
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) SetExtraData(v map[string]interface{}) *ListTaskExecutionsResponseBodyTaskExecutions {
	s.ExtraData = v
	return s
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) SetProperties(v string) *ListTaskExecutionsResponseBodyTaskExecutions {
	s.Properties = &v
	return s
}

type ListTaskExecutionsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTaskExecutionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTaskExecutionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTaskExecutionsResponse) GoString() string {
	return s.String()
}

func (s *ListTaskExecutionsResponse) SetHeaders(v map[string]*string) *ListTaskExecutionsResponse {
	s.Headers = v
	return s
}

func (s *ListTaskExecutionsResponse) SetBody(v *ListTaskExecutionsResponseBody) *ListTaskExecutionsResponse {
	s.Body = v
	return s
}

type ListTemplatesRequest struct {
	RegionId          *string                `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TemplateName      *string                `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateFormat    *string                `json:"TemplateFormat,omitempty" xml:"TemplateFormat,omitempty"`
	ShareType         *string                `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	CreatedBy         *string                `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	CreatedDateBefore *string                `json:"CreatedDateBefore,omitempty" xml:"CreatedDateBefore,omitempty"`
	CreatedDateAfter  *string                `json:"CreatedDateAfter,omitempty" xml:"CreatedDateAfter,omitempty"`
	Tags              map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	Category          *string                `json:"Category,omitempty" xml:"Category,omitempty"`
	MaxResults        *int32                 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken         *string                `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	SortField         *string                `json:"SortField,omitempty" xml:"SortField,omitempty"`
	SortOrder         *string                `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	HasTrigger        *bool                  `json:"HasTrigger,omitempty" xml:"HasTrigger,omitempty"`
	TemplateType      *string                `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	ResourceGroupId   *string                `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListTemplatesRequest) SetRegionId(v string) *ListTemplatesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTemplatesRequest) SetTemplateName(v string) *ListTemplatesRequest {
	s.TemplateName = &v
	return s
}

func (s *ListTemplatesRequest) SetTemplateFormat(v string) *ListTemplatesRequest {
	s.TemplateFormat = &v
	return s
}

func (s *ListTemplatesRequest) SetShareType(v string) *ListTemplatesRequest {
	s.ShareType = &v
	return s
}

func (s *ListTemplatesRequest) SetCreatedBy(v string) *ListTemplatesRequest {
	s.CreatedBy = &v
	return s
}

func (s *ListTemplatesRequest) SetCreatedDateBefore(v string) *ListTemplatesRequest {
	s.CreatedDateBefore = &v
	return s
}

func (s *ListTemplatesRequest) SetCreatedDateAfter(v string) *ListTemplatesRequest {
	s.CreatedDateAfter = &v
	return s
}

func (s *ListTemplatesRequest) SetTags(v map[string]interface{}) *ListTemplatesRequest {
	s.Tags = v
	return s
}

func (s *ListTemplatesRequest) SetCategory(v string) *ListTemplatesRequest {
	s.Category = &v
	return s
}

func (s *ListTemplatesRequest) SetMaxResults(v int32) *ListTemplatesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTemplatesRequest) SetNextToken(v string) *ListTemplatesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTemplatesRequest) SetSortField(v string) *ListTemplatesRequest {
	s.SortField = &v
	return s
}

func (s *ListTemplatesRequest) SetSortOrder(v string) *ListTemplatesRequest {
	s.SortOrder = &v
	return s
}

func (s *ListTemplatesRequest) SetHasTrigger(v bool) *ListTemplatesRequest {
	s.HasTrigger = &v
	return s
}

func (s *ListTemplatesRequest) SetTemplateType(v string) *ListTemplatesRequest {
	s.TemplateType = &v
	return s
}

func (s *ListTemplatesRequest) SetResourceGroupId(v string) *ListTemplatesRequest {
	s.ResourceGroupId = &v
	return s
}

type ListTemplatesShrinkRequest struct {
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TemplateName      *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateFormat    *string `json:"TemplateFormat,omitempty" xml:"TemplateFormat,omitempty"`
	ShareType         *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	CreatedBy         *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	CreatedDateBefore *string `json:"CreatedDateBefore,omitempty" xml:"CreatedDateBefore,omitempty"`
	CreatedDateAfter  *string `json:"CreatedDateAfter,omitempty" xml:"CreatedDateAfter,omitempty"`
	TagsShrink        *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	Category          *string `json:"Category,omitempty" xml:"Category,omitempty"`
	MaxResults        *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken         *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	SortField         *string `json:"SortField,omitempty" xml:"SortField,omitempty"`
	SortOrder         *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	HasTrigger        *bool   `json:"HasTrigger,omitempty" xml:"HasTrigger,omitempty"`
	TemplateType      *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	ResourceGroupId   *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListTemplatesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListTemplatesShrinkRequest) SetRegionId(v string) *ListTemplatesShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ListTemplatesShrinkRequest) SetTemplateName(v string) *ListTemplatesShrinkRequest {
	s.TemplateName = &v
	return s
}

func (s *ListTemplatesShrinkRequest) SetTemplateFormat(v string) *ListTemplatesShrinkRequest {
	s.TemplateFormat = &v
	return s
}

func (s *ListTemplatesShrinkRequest) SetShareType(v string) *ListTemplatesShrinkRequest {
	s.ShareType = &v
	return s
}

func (s *ListTemplatesShrinkRequest) SetCreatedBy(v string) *ListTemplatesShrinkRequest {
	s.CreatedBy = &v
	return s
}

func (s *ListTemplatesShrinkRequest) SetCreatedDateBefore(v string) *ListTemplatesShrinkRequest {
	s.CreatedDateBefore = &v
	return s
}

func (s *ListTemplatesShrinkRequest) SetCreatedDateAfter(v string) *ListTemplatesShrinkRequest {
	s.CreatedDateAfter = &v
	return s
}

func (s *ListTemplatesShrinkRequest) SetTagsShrink(v string) *ListTemplatesShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *ListTemplatesShrinkRequest) SetCategory(v string) *ListTemplatesShrinkRequest {
	s.Category = &v
	return s
}

func (s *ListTemplatesShrinkRequest) SetMaxResults(v int32) *ListTemplatesShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTemplatesShrinkRequest) SetNextToken(v string) *ListTemplatesShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListTemplatesShrinkRequest) SetSortField(v string) *ListTemplatesShrinkRequest {
	s.SortField = &v
	return s
}

func (s *ListTemplatesShrinkRequest) SetSortOrder(v string) *ListTemplatesShrinkRequest {
	s.SortOrder = &v
	return s
}

func (s *ListTemplatesShrinkRequest) SetHasTrigger(v bool) *ListTemplatesShrinkRequest {
	s.HasTrigger = &v
	return s
}

func (s *ListTemplatesShrinkRequest) SetTemplateType(v string) *ListTemplatesShrinkRequest {
	s.TemplateType = &v
	return s
}

func (s *ListTemplatesShrinkRequest) SetResourceGroupId(v string) *ListTemplatesShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

type ListTemplatesResponseBody struct {
	NextToken  *string                               `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	MaxResults *int32                                `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Templates  []*ListTemplatesResponseBodyTemplates `json:"Templates,omitempty" xml:"Templates,omitempty" type:"Repeated"`
}

func (s ListTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTemplatesResponseBody) SetNextToken(v string) *ListTemplatesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTemplatesResponseBody) SetRequestId(v string) *ListTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTemplatesResponseBody) SetMaxResults(v int32) *ListTemplatesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTemplatesResponseBody) SetTemplates(v []*ListTemplatesResponseBodyTemplates) *ListTemplatesResponseBody {
	s.Templates = v
	return s
}

type ListTemplatesResponseBodyTemplates struct {
	Hash                *string                `json:"Hash,omitempty" xml:"Hash,omitempty"`
	UpdatedDate         *string                `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
	UpdatedBy           *string                `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	TemplateType        *string                `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	Tags                map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	TemplateName        *string                `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateVersion     *string                `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	TemplateFormat      *string                `json:"TemplateFormat,omitempty" xml:"TemplateFormat,omitempty"`
	Popularity          *int32                 `json:"Popularity,omitempty" xml:"Popularity,omitempty"`
	TotalExecutionCount *int32                 `json:"TotalExecutionCount,omitempty" xml:"TotalExecutionCount,omitempty"`
	Description         *string                `json:"Description,omitempty" xml:"Description,omitempty"`
	ResourceGroupId     *string                `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	CreatedBy           *string                `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	CreatedDate         *string                `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	Category            *string                `json:"Category,omitempty" xml:"Category,omitempty"`
	HasTrigger          *bool                  `json:"HasTrigger,omitempty" xml:"HasTrigger,omitempty"`
	TemplateId          *string                `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	ShareType           *string                `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
}

func (s ListTemplatesResponseBodyTemplates) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesResponseBodyTemplates) GoString() string {
	return s.String()
}

func (s *ListTemplatesResponseBodyTemplates) SetHash(v string) *ListTemplatesResponseBodyTemplates {
	s.Hash = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetUpdatedDate(v string) *ListTemplatesResponseBodyTemplates {
	s.UpdatedDate = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetUpdatedBy(v string) *ListTemplatesResponseBodyTemplates {
	s.UpdatedBy = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetTemplateType(v string) *ListTemplatesResponseBodyTemplates {
	s.TemplateType = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetTags(v map[string]interface{}) *ListTemplatesResponseBodyTemplates {
	s.Tags = v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetTemplateName(v string) *ListTemplatesResponseBodyTemplates {
	s.TemplateName = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetTemplateVersion(v string) *ListTemplatesResponseBodyTemplates {
	s.TemplateVersion = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetTemplateFormat(v string) *ListTemplatesResponseBodyTemplates {
	s.TemplateFormat = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetPopularity(v int32) *ListTemplatesResponseBodyTemplates {
	s.Popularity = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetTotalExecutionCount(v int32) *ListTemplatesResponseBodyTemplates {
	s.TotalExecutionCount = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetDescription(v string) *ListTemplatesResponseBodyTemplates {
	s.Description = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetResourceGroupId(v string) *ListTemplatesResponseBodyTemplates {
	s.ResourceGroupId = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetCreatedBy(v string) *ListTemplatesResponseBodyTemplates {
	s.CreatedBy = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetCreatedDate(v string) *ListTemplatesResponseBodyTemplates {
	s.CreatedDate = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetCategory(v string) *ListTemplatesResponseBodyTemplates {
	s.Category = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetHasTrigger(v bool) *ListTemplatesResponseBodyTemplates {
	s.HasTrigger = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetTemplateId(v string) *ListTemplatesResponseBodyTemplates {
	s.TemplateId = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetShareType(v string) *ListTemplatesResponseBodyTemplates {
	s.ShareType = &v
	return s
}

type ListTemplatesResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListTemplatesResponse) SetHeaders(v map[string]*string) *ListTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListTemplatesResponse) SetBody(v *ListTemplatesResponseBody) *ListTemplatesResponse {
	s.Body = v
	return s
}

type ListTemplateVersionsRequest struct {
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	MaxResults   *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ShareType    *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
}

func (s ListTemplateVersionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTemplateVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListTemplateVersionsRequest) SetRegionId(v string) *ListTemplateVersionsRequest {
	s.RegionId = &v
	return s
}

func (s *ListTemplateVersionsRequest) SetTemplateName(v string) *ListTemplateVersionsRequest {
	s.TemplateName = &v
	return s
}

func (s *ListTemplateVersionsRequest) SetMaxResults(v int32) *ListTemplateVersionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTemplateVersionsRequest) SetNextToken(v string) *ListTemplateVersionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListTemplateVersionsRequest) SetShareType(v string) *ListTemplateVersionsRequest {
	s.ShareType = &v
	return s
}

type ListTemplateVersionsResponseBody struct {
	NextToken        *string                                             `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId        *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	MaxResults       *int32                                              `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	TemplateVersions []*ListTemplateVersionsResponseBodyTemplateVersions `json:"TemplateVersions,omitempty" xml:"TemplateVersions,omitempty" type:"Repeated"`
}

func (s ListTemplateVersionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTemplateVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTemplateVersionsResponseBody) SetNextToken(v string) *ListTemplateVersionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTemplateVersionsResponseBody) SetRequestId(v string) *ListTemplateVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTemplateVersionsResponseBody) SetMaxResults(v int32) *ListTemplateVersionsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTemplateVersionsResponseBody) SetTemplateVersions(v []*ListTemplateVersionsResponseBodyTemplateVersions) *ListTemplateVersionsResponseBody {
	s.TemplateVersions = v
	return s
}

type ListTemplateVersionsResponseBodyTemplateVersions struct {
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	UpdatedDate     *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
	UpdatedBy       *string `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	VersionName     *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	TemplateFormat  *string `json:"TemplateFormat,omitempty" xml:"TemplateFormat,omitempty"`
}

func (s ListTemplateVersionsResponseBodyTemplateVersions) String() string {
	return tea.Prettify(s)
}

func (s ListTemplateVersionsResponseBodyTemplateVersions) GoString() string {
	return s.String()
}

func (s *ListTemplateVersionsResponseBodyTemplateVersions) SetDescription(v string) *ListTemplateVersionsResponseBodyTemplateVersions {
	s.Description = &v
	return s
}

func (s *ListTemplateVersionsResponseBodyTemplateVersions) SetUpdatedDate(v string) *ListTemplateVersionsResponseBodyTemplateVersions {
	s.UpdatedDate = &v
	return s
}

func (s *ListTemplateVersionsResponseBodyTemplateVersions) SetUpdatedBy(v string) *ListTemplateVersionsResponseBodyTemplateVersions {
	s.UpdatedBy = &v
	return s
}

func (s *ListTemplateVersionsResponseBodyTemplateVersions) SetVersionName(v string) *ListTemplateVersionsResponseBodyTemplateVersions {
	s.VersionName = &v
	return s
}

func (s *ListTemplateVersionsResponseBodyTemplateVersions) SetTemplateVersion(v string) *ListTemplateVersionsResponseBodyTemplateVersions {
	s.TemplateVersion = &v
	return s
}

func (s *ListTemplateVersionsResponseBodyTemplateVersions) SetTemplateFormat(v string) *ListTemplateVersionsResponseBodyTemplateVersions {
	s.TemplateFormat = &v
	return s
}

type ListTemplateVersionsResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTemplateVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTemplateVersionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTemplateVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListTemplateVersionsResponse) SetHeaders(v map[string]*string) *ListTemplateVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListTemplateVersionsResponse) SetBody(v *ListTemplateVersionsResponseBody) *ListTemplateVersionsResponse {
	s.Body = v
	return s
}

type NotifyExecutionRequest struct {
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ExecutionId      *string `json:"ExecutionId,omitempty" xml:"ExecutionId,omitempty"`
	NotifyType       *string `json:"NotifyType,omitempty" xml:"NotifyType,omitempty"`
	NotifyNote       *string `json:"NotifyNote,omitempty" xml:"NotifyNote,omitempty"`
	TaskName         *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TaskExecutionId  *string `json:"TaskExecutionId,omitempty" xml:"TaskExecutionId,omitempty"`
	ExecutionStatus  *string `json:"ExecutionStatus,omitempty" xml:"ExecutionStatus,omitempty"`
	Parameters       *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	LoopItem         *string `json:"LoopItem,omitempty" xml:"LoopItem,omitempty"`
	TaskExecutionIds *string `json:"TaskExecutionIds,omitempty" xml:"TaskExecutionIds,omitempty"`
}

func (s NotifyExecutionRequest) String() string {
	return tea.Prettify(s)
}

func (s NotifyExecutionRequest) GoString() string {
	return s.String()
}

func (s *NotifyExecutionRequest) SetRegionId(v string) *NotifyExecutionRequest {
	s.RegionId = &v
	return s
}

func (s *NotifyExecutionRequest) SetExecutionId(v string) *NotifyExecutionRequest {
	s.ExecutionId = &v
	return s
}

func (s *NotifyExecutionRequest) SetNotifyType(v string) *NotifyExecutionRequest {
	s.NotifyType = &v
	return s
}

func (s *NotifyExecutionRequest) SetNotifyNote(v string) *NotifyExecutionRequest {
	s.NotifyNote = &v
	return s
}

func (s *NotifyExecutionRequest) SetTaskName(v string) *NotifyExecutionRequest {
	s.TaskName = &v
	return s
}

func (s *NotifyExecutionRequest) SetTaskExecutionId(v string) *NotifyExecutionRequest {
	s.TaskExecutionId = &v
	return s
}

func (s *NotifyExecutionRequest) SetExecutionStatus(v string) *NotifyExecutionRequest {
	s.ExecutionStatus = &v
	return s
}

func (s *NotifyExecutionRequest) SetParameters(v string) *NotifyExecutionRequest {
	s.Parameters = &v
	return s
}

func (s *NotifyExecutionRequest) SetLoopItem(v string) *NotifyExecutionRequest {
	s.LoopItem = &v
	return s
}

func (s *NotifyExecutionRequest) SetTaskExecutionIds(v string) *NotifyExecutionRequest {
	s.TaskExecutionIds = &v
	return s
}

type NotifyExecutionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s NotifyExecutionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s NotifyExecutionResponseBody) GoString() string {
	return s.String()
}

func (s *NotifyExecutionResponseBody) SetRequestId(v string) *NotifyExecutionResponseBody {
	s.RequestId = &v
	return s
}

type NotifyExecutionResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *NotifyExecutionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s NotifyExecutionResponse) String() string {
	return tea.Prettify(s)
}

func (s NotifyExecutionResponse) GoString() string {
	return s.String()
}

func (s *NotifyExecutionResponse) SetHeaders(v map[string]*string) *NotifyExecutionResponse {
	s.Headers = v
	return s
}

func (s *NotifyExecutionResponse) SetBody(v *NotifyExecutionResponseBody) *NotifyExecutionResponse {
	s.Body = v
	return s
}

type RegisterDefaultPatchBaselineRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s RegisterDefaultPatchBaselineRequest) String() string {
	return tea.Prettify(s)
}

func (s RegisterDefaultPatchBaselineRequest) GoString() string {
	return s.String()
}

func (s *RegisterDefaultPatchBaselineRequest) SetRegionId(v string) *RegisterDefaultPatchBaselineRequest {
	s.RegionId = &v
	return s
}

func (s *RegisterDefaultPatchBaselineRequest) SetName(v string) *RegisterDefaultPatchBaselineRequest {
	s.Name = &v
	return s
}

type RegisterDefaultPatchBaselineResponseBody struct {
	RequestId     *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PatchBaseline *RegisterDefaultPatchBaselineResponseBodyPatchBaseline `json:"PatchBaseline,omitempty" xml:"PatchBaseline,omitempty" type:"Struct"`
}

func (s RegisterDefaultPatchBaselineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RegisterDefaultPatchBaselineResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterDefaultPatchBaselineResponseBody) SetRequestId(v string) *RegisterDefaultPatchBaselineResponseBody {
	s.RequestId = &v
	return s
}

func (s *RegisterDefaultPatchBaselineResponseBody) SetPatchBaseline(v *RegisterDefaultPatchBaselineResponseBodyPatchBaseline) *RegisterDefaultPatchBaselineResponseBody {
	s.PatchBaseline = v
	return s
}

type RegisterDefaultPatchBaselineResponseBodyPatchBaseline struct {
	OperationSystem *string `json:"OperationSystem,omitempty" xml:"OperationSystem,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	UpdatedDate     *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
	UpdatedBy       *string `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	CreatedBy       *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	CreatedDate     *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ApprovalRules   *string `json:"ApprovalRules,omitempty" xml:"ApprovalRules,omitempty"`
	Id              *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ShareType       *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
}

func (s RegisterDefaultPatchBaselineResponseBodyPatchBaseline) String() string {
	return tea.Prettify(s)
}

func (s RegisterDefaultPatchBaselineResponseBodyPatchBaseline) GoString() string {
	return s.String()
}

func (s *RegisterDefaultPatchBaselineResponseBodyPatchBaseline) SetOperationSystem(v string) *RegisterDefaultPatchBaselineResponseBodyPatchBaseline {
	s.OperationSystem = &v
	return s
}

func (s *RegisterDefaultPatchBaselineResponseBodyPatchBaseline) SetDescription(v string) *RegisterDefaultPatchBaselineResponseBodyPatchBaseline {
	s.Description = &v
	return s
}

func (s *RegisterDefaultPatchBaselineResponseBodyPatchBaseline) SetUpdatedDate(v string) *RegisterDefaultPatchBaselineResponseBodyPatchBaseline {
	s.UpdatedDate = &v
	return s
}

func (s *RegisterDefaultPatchBaselineResponseBodyPatchBaseline) SetUpdatedBy(v string) *RegisterDefaultPatchBaselineResponseBodyPatchBaseline {
	s.UpdatedBy = &v
	return s
}

func (s *RegisterDefaultPatchBaselineResponseBodyPatchBaseline) SetCreatedBy(v string) *RegisterDefaultPatchBaselineResponseBodyPatchBaseline {
	s.CreatedBy = &v
	return s
}

func (s *RegisterDefaultPatchBaselineResponseBodyPatchBaseline) SetCreatedDate(v string) *RegisterDefaultPatchBaselineResponseBodyPatchBaseline {
	s.CreatedDate = &v
	return s
}

func (s *RegisterDefaultPatchBaselineResponseBodyPatchBaseline) SetName(v string) *RegisterDefaultPatchBaselineResponseBodyPatchBaseline {
	s.Name = &v
	return s
}

func (s *RegisterDefaultPatchBaselineResponseBodyPatchBaseline) SetApprovalRules(v string) *RegisterDefaultPatchBaselineResponseBodyPatchBaseline {
	s.ApprovalRules = &v
	return s
}

func (s *RegisterDefaultPatchBaselineResponseBodyPatchBaseline) SetId(v string) *RegisterDefaultPatchBaselineResponseBodyPatchBaseline {
	s.Id = &v
	return s
}

func (s *RegisterDefaultPatchBaselineResponseBodyPatchBaseline) SetShareType(v string) *RegisterDefaultPatchBaselineResponseBodyPatchBaseline {
	s.ShareType = &v
	return s
}

type RegisterDefaultPatchBaselineResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RegisterDefaultPatchBaselineResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RegisterDefaultPatchBaselineResponse) String() string {
	return tea.Prettify(s)
}

func (s RegisterDefaultPatchBaselineResponse) GoString() string {
	return s.String()
}

func (s *RegisterDefaultPatchBaselineResponse) SetHeaders(v map[string]*string) *RegisterDefaultPatchBaselineResponse {
	s.Headers = v
	return s
}

func (s *RegisterDefaultPatchBaselineResponse) SetBody(v *RegisterDefaultPatchBaselineResponseBody) *RegisterDefaultPatchBaselineResponse {
	s.Body = v
	return s
}

type SearchInventoryRequest struct {
	RegionId   *string                         `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	NextToken  *string                         `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	MaxResults *int32                          `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Filter     []*SearchInventoryRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	Aggregator []*string                       `json:"Aggregator,omitempty" xml:"Aggregator,omitempty" type:"Repeated"`
}

func (s SearchInventoryRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchInventoryRequest) GoString() string {
	return s.String()
}

func (s *SearchInventoryRequest) SetRegionId(v string) *SearchInventoryRequest {
	s.RegionId = &v
	return s
}

func (s *SearchInventoryRequest) SetNextToken(v string) *SearchInventoryRequest {
	s.NextToken = &v
	return s
}

func (s *SearchInventoryRequest) SetMaxResults(v int32) *SearchInventoryRequest {
	s.MaxResults = &v
	return s
}

func (s *SearchInventoryRequest) SetFilter(v []*SearchInventoryRequestFilter) *SearchInventoryRequest {
	s.Filter = v
	return s
}

func (s *SearchInventoryRequest) SetAggregator(v []*string) *SearchInventoryRequest {
	s.Aggregator = v
	return s
}

type SearchInventoryRequestFilter struct {
	Value    []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
	Operator *string   `json:"Operator,omitempty" xml:"Operator,omitempty"`
	Name     *string   `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s SearchInventoryRequestFilter) String() string {
	return tea.Prettify(s)
}

func (s SearchInventoryRequestFilter) GoString() string {
	return s.String()
}

func (s *SearchInventoryRequestFilter) SetValue(v []*string) *SearchInventoryRequestFilter {
	s.Value = v
	return s
}

func (s *SearchInventoryRequestFilter) SetOperator(v string) *SearchInventoryRequestFilter {
	s.Operator = &v
	return s
}

func (s *SearchInventoryRequestFilter) SetName(v string) *SearchInventoryRequestFilter {
	s.Name = &v
	return s
}

type SearchInventoryResponseBody struct {
	NextToken  *string                  `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId  *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	MaxResults *int32                   `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Entities   []map[string]interface{} `json:"Entities,omitempty" xml:"Entities,omitempty" type:"Repeated"`
}

func (s SearchInventoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchInventoryResponseBody) GoString() string {
	return s.String()
}

func (s *SearchInventoryResponseBody) SetNextToken(v string) *SearchInventoryResponseBody {
	s.NextToken = &v
	return s
}

func (s *SearchInventoryResponseBody) SetRequestId(v string) *SearchInventoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchInventoryResponseBody) SetMaxResults(v int32) *SearchInventoryResponseBody {
	s.MaxResults = &v
	return s
}

func (s *SearchInventoryResponseBody) SetEntities(v []map[string]interface{}) *SearchInventoryResponseBody {
	s.Entities = v
	return s
}

type SearchInventoryResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SearchInventoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchInventoryResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchInventoryResponse) GoString() string {
	return s.String()
}

func (s *SearchInventoryResponse) SetHeaders(v map[string]*string) *SearchInventoryResponse {
	s.Headers = v
	return s
}

func (s *SearchInventoryResponse) SetBody(v *SearchInventoryResponseBody) *SearchInventoryResponse {
	s.Body = v
	return s
}

type SetServiceSettingsRequest struct {
	DeliveryOssEnabled     *bool   `json:"DeliveryOssEnabled,omitempty" xml:"DeliveryOssEnabled,omitempty"`
	DeliveryOssBucketName  *string `json:"DeliveryOssBucketName,omitempty" xml:"DeliveryOssBucketName,omitempty"`
	DeliveryOssKeyPrefix   *string `json:"DeliveryOssKeyPrefix,omitempty" xml:"DeliveryOssKeyPrefix,omitempty"`
	DeliverySlsProjectName *string `json:"DeliverySlsProjectName,omitempty" xml:"DeliverySlsProjectName,omitempty"`
	DeliverySlsEnabled     *bool   `json:"DeliverySlsEnabled,omitempty" xml:"DeliverySlsEnabled,omitempty"`
	RegionId               *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RdcEnterpriseId        *string `json:"RdcEnterpriseId,omitempty" xml:"RdcEnterpriseId,omitempty"`
}

func (s SetServiceSettingsRequest) String() string {
	return tea.Prettify(s)
}

func (s SetServiceSettingsRequest) GoString() string {
	return s.String()
}

func (s *SetServiceSettingsRequest) SetDeliveryOssEnabled(v bool) *SetServiceSettingsRequest {
	s.DeliveryOssEnabled = &v
	return s
}

func (s *SetServiceSettingsRequest) SetDeliveryOssBucketName(v string) *SetServiceSettingsRequest {
	s.DeliveryOssBucketName = &v
	return s
}

func (s *SetServiceSettingsRequest) SetDeliveryOssKeyPrefix(v string) *SetServiceSettingsRequest {
	s.DeliveryOssKeyPrefix = &v
	return s
}

func (s *SetServiceSettingsRequest) SetDeliverySlsProjectName(v string) *SetServiceSettingsRequest {
	s.DeliverySlsProjectName = &v
	return s
}

func (s *SetServiceSettingsRequest) SetDeliverySlsEnabled(v bool) *SetServiceSettingsRequest {
	s.DeliverySlsEnabled = &v
	return s
}

func (s *SetServiceSettingsRequest) SetRegionId(v string) *SetServiceSettingsRequest {
	s.RegionId = &v
	return s
}

func (s *SetServiceSettingsRequest) SetRdcEnterpriseId(v string) *SetServiceSettingsRequest {
	s.RdcEnterpriseId = &v
	return s
}

type SetServiceSettingsResponseBody struct {
	RequestId       *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServiceSettings []*SetServiceSettingsResponseBodyServiceSettings `json:"ServiceSettings,omitempty" xml:"ServiceSettings,omitempty" type:"Repeated"`
}

func (s SetServiceSettingsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetServiceSettingsResponseBody) GoString() string {
	return s.String()
}

func (s *SetServiceSettingsResponseBody) SetRequestId(v string) *SetServiceSettingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetServiceSettingsResponseBody) SetServiceSettings(v []*SetServiceSettingsResponseBodyServiceSettings) *SetServiceSettingsResponseBody {
	s.ServiceSettings = v
	return s
}

type SetServiceSettingsResponseBodyServiceSettings struct {
	DeliveryOssBucketName  *string `json:"DeliveryOssBucketName,omitempty" xml:"DeliveryOssBucketName,omitempty"`
	DeliveryOssKeyPrefix   *string `json:"DeliveryOssKeyPrefix,omitempty" xml:"DeliveryOssKeyPrefix,omitempty"`
	DeliveryOssEnabled     *bool   `json:"DeliveryOssEnabled,omitempty" xml:"DeliveryOssEnabled,omitempty"`
	DeliverySlsEnabled     *bool   `json:"DeliverySlsEnabled,omitempty" xml:"DeliverySlsEnabled,omitempty"`
	DeliverySlsProjectName *string `json:"DeliverySlsProjectName,omitempty" xml:"DeliverySlsProjectName,omitempty"`
	RdcEnterpriseId        *string `json:"RdcEnterpriseId,omitempty" xml:"RdcEnterpriseId,omitempty"`
}

func (s SetServiceSettingsResponseBodyServiceSettings) String() string {
	return tea.Prettify(s)
}

func (s SetServiceSettingsResponseBodyServiceSettings) GoString() string {
	return s.String()
}

func (s *SetServiceSettingsResponseBodyServiceSettings) SetDeliveryOssBucketName(v string) *SetServiceSettingsResponseBodyServiceSettings {
	s.DeliveryOssBucketName = &v
	return s
}

func (s *SetServiceSettingsResponseBodyServiceSettings) SetDeliveryOssKeyPrefix(v string) *SetServiceSettingsResponseBodyServiceSettings {
	s.DeliveryOssKeyPrefix = &v
	return s
}

func (s *SetServiceSettingsResponseBodyServiceSettings) SetDeliveryOssEnabled(v bool) *SetServiceSettingsResponseBodyServiceSettings {
	s.DeliveryOssEnabled = &v
	return s
}

func (s *SetServiceSettingsResponseBodyServiceSettings) SetDeliverySlsEnabled(v bool) *SetServiceSettingsResponseBodyServiceSettings {
	s.DeliverySlsEnabled = &v
	return s
}

func (s *SetServiceSettingsResponseBodyServiceSettings) SetDeliverySlsProjectName(v string) *SetServiceSettingsResponseBodyServiceSettings {
	s.DeliverySlsProjectName = &v
	return s
}

func (s *SetServiceSettingsResponseBodyServiceSettings) SetRdcEnterpriseId(v string) *SetServiceSettingsResponseBodyServiceSettings {
	s.RdcEnterpriseId = &v
	return s
}

type SetServiceSettingsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetServiceSettingsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetServiceSettingsResponse) String() string {
	return tea.Prettify(s)
}

func (s SetServiceSettingsResponse) GoString() string {
	return s.String()
}

func (s *SetServiceSettingsResponse) SetHeaders(v map[string]*string) *SetServiceSettingsResponse {
	s.Headers = v
	return s
}

func (s *SetServiceSettingsResponse) SetBody(v *SetServiceSettingsResponseBody) *SetServiceSettingsResponse {
	s.Body = v
	return s
}

type StartExecutionRequest struct {
	RegionId          *string                `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TemplateName      *string                `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateVersion   *string                `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	Mode              *string                `json:"Mode,omitempty" xml:"Mode,omitempty"`
	LoopMode          *string                `json:"LoopMode,omitempty" xml:"LoopMode,omitempty"`
	ParentExecutionId *string                `json:"ParentExecutionId,omitempty" xml:"ParentExecutionId,omitempty"`
	SafetyCheck       *string                `json:"SafetyCheck,omitempty" xml:"SafetyCheck,omitempty"`
	Parameters        *string                `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	ClientToken       *string                `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Tags              map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	Description       *string                `json:"Description,omitempty" xml:"Description,omitempty"`
	TemplateContent   *string                `json:"TemplateContent,omitempty" xml:"TemplateContent,omitempty"`
	ResourceGroupId   *string                `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s StartExecutionRequest) String() string {
	return tea.Prettify(s)
}

func (s StartExecutionRequest) GoString() string {
	return s.String()
}

func (s *StartExecutionRequest) SetRegionId(v string) *StartExecutionRequest {
	s.RegionId = &v
	return s
}

func (s *StartExecutionRequest) SetTemplateName(v string) *StartExecutionRequest {
	s.TemplateName = &v
	return s
}

func (s *StartExecutionRequest) SetTemplateVersion(v string) *StartExecutionRequest {
	s.TemplateVersion = &v
	return s
}

func (s *StartExecutionRequest) SetMode(v string) *StartExecutionRequest {
	s.Mode = &v
	return s
}

func (s *StartExecutionRequest) SetLoopMode(v string) *StartExecutionRequest {
	s.LoopMode = &v
	return s
}

func (s *StartExecutionRequest) SetParentExecutionId(v string) *StartExecutionRequest {
	s.ParentExecutionId = &v
	return s
}

func (s *StartExecutionRequest) SetSafetyCheck(v string) *StartExecutionRequest {
	s.SafetyCheck = &v
	return s
}

func (s *StartExecutionRequest) SetParameters(v string) *StartExecutionRequest {
	s.Parameters = &v
	return s
}

func (s *StartExecutionRequest) SetClientToken(v string) *StartExecutionRequest {
	s.ClientToken = &v
	return s
}

func (s *StartExecutionRequest) SetTags(v map[string]interface{}) *StartExecutionRequest {
	s.Tags = v
	return s
}

func (s *StartExecutionRequest) SetDescription(v string) *StartExecutionRequest {
	s.Description = &v
	return s
}

func (s *StartExecutionRequest) SetTemplateContent(v string) *StartExecutionRequest {
	s.TemplateContent = &v
	return s
}

func (s *StartExecutionRequest) SetResourceGroupId(v string) *StartExecutionRequest {
	s.ResourceGroupId = &v
	return s
}

type StartExecutionShrinkRequest struct {
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TemplateName      *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateVersion   *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	Mode              *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	LoopMode          *string `json:"LoopMode,omitempty" xml:"LoopMode,omitempty"`
	ParentExecutionId *string `json:"ParentExecutionId,omitempty" xml:"ParentExecutionId,omitempty"`
	SafetyCheck       *string `json:"SafetyCheck,omitempty" xml:"SafetyCheck,omitempty"`
	Parameters        *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	ClientToken       *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	TagsShrink        *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	TemplateContent   *string `json:"TemplateContent,omitempty" xml:"TemplateContent,omitempty"`
	ResourceGroupId   *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s StartExecutionShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s StartExecutionShrinkRequest) GoString() string {
	return s.String()
}

func (s *StartExecutionShrinkRequest) SetRegionId(v string) *StartExecutionShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *StartExecutionShrinkRequest) SetTemplateName(v string) *StartExecutionShrinkRequest {
	s.TemplateName = &v
	return s
}

func (s *StartExecutionShrinkRequest) SetTemplateVersion(v string) *StartExecutionShrinkRequest {
	s.TemplateVersion = &v
	return s
}

func (s *StartExecutionShrinkRequest) SetMode(v string) *StartExecutionShrinkRequest {
	s.Mode = &v
	return s
}

func (s *StartExecutionShrinkRequest) SetLoopMode(v string) *StartExecutionShrinkRequest {
	s.LoopMode = &v
	return s
}

func (s *StartExecutionShrinkRequest) SetParentExecutionId(v string) *StartExecutionShrinkRequest {
	s.ParentExecutionId = &v
	return s
}

func (s *StartExecutionShrinkRequest) SetSafetyCheck(v string) *StartExecutionShrinkRequest {
	s.SafetyCheck = &v
	return s
}

func (s *StartExecutionShrinkRequest) SetParameters(v string) *StartExecutionShrinkRequest {
	s.Parameters = &v
	return s
}

func (s *StartExecutionShrinkRequest) SetClientToken(v string) *StartExecutionShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *StartExecutionShrinkRequest) SetTagsShrink(v string) *StartExecutionShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *StartExecutionShrinkRequest) SetDescription(v string) *StartExecutionShrinkRequest {
	s.Description = &v
	return s
}

func (s *StartExecutionShrinkRequest) SetTemplateContent(v string) *StartExecutionShrinkRequest {
	s.TemplateContent = &v
	return s
}

func (s *StartExecutionShrinkRequest) SetResourceGroupId(v string) *StartExecutionShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

type StartExecutionResponseBody struct {
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Execution *StartExecutionResponseBodyExecution `json:"Execution,omitempty" xml:"Execution,omitempty" type:"Struct"`
}

func (s StartExecutionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartExecutionResponseBody) GoString() string {
	return s.String()
}

func (s *StartExecutionResponseBody) SetRequestId(v string) *StartExecutionResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartExecutionResponseBody) SetExecution(v *StartExecutionResponseBodyExecution) *StartExecutionResponseBody {
	s.Execution = v
	return s
}

type StartExecutionResponseBodyExecution struct {
	Outputs           *string                                            `json:"Outputs,omitempty" xml:"Outputs,omitempty"`
	Status            *string                                            `json:"Status,omitempty" xml:"Status,omitempty"`
	EndDate           *string                                            `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	ExecutedBy        *string                                            `json:"ExecutedBy,omitempty" xml:"ExecutedBy,omitempty"`
	IsParent          *bool                                              `json:"IsParent,omitempty" xml:"IsParent,omitempty"`
	Tags              map[string]interface{}                             `json:"Tags,omitempty" xml:"Tags,omitempty"`
	StartDate         *string                                            `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	SafetyCheck       *string                                            `json:"SafetyCheck,omitempty" xml:"SafetyCheck,omitempty"`
	Mode              *string                                            `json:"Mode,omitempty" xml:"Mode,omitempty"`
	TemplateName      *string                                            `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	CreateDate        *string                                            `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	TemplateVersion   *string                                            `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	ExecutionId       *string                                            `json:"ExecutionId,omitempty" xml:"ExecutionId,omitempty"`
	Parameters        *string                                            `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	Description       *string                                            `json:"Description,omitempty" xml:"Description,omitempty"`
	Counters          map[string]interface{}                             `json:"Counters,omitempty" xml:"Counters,omitempty"`
	UpdateDate        *string                                            `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
	ResourceGroupId   *string                                            `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ParentExecutionId *string                                            `json:"ParentExecutionId,omitempty" xml:"ParentExecutionId,omitempty"`
	RamRole           *string                                            `json:"RamRole,omitempty" xml:"RamRole,omitempty"`
	TemplateId        *string                                            `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	StatusMessage     *string                                            `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	LoopMode          *string                                            `json:"LoopMode,omitempty" xml:"LoopMode,omitempty"`
	CurrentTasks      []*StartExecutionResponseBodyExecutionCurrentTasks `json:"CurrentTasks,omitempty" xml:"CurrentTasks,omitempty" type:"Repeated"`
}

func (s StartExecutionResponseBodyExecution) String() string {
	return tea.Prettify(s)
}

func (s StartExecutionResponseBodyExecution) GoString() string {
	return s.String()
}

func (s *StartExecutionResponseBodyExecution) SetOutputs(v string) *StartExecutionResponseBodyExecution {
	s.Outputs = &v
	return s
}

func (s *StartExecutionResponseBodyExecution) SetStatus(v string) *StartExecutionResponseBodyExecution {
	s.Status = &v
	return s
}

func (s *StartExecutionResponseBodyExecution) SetEndDate(v string) *StartExecutionResponseBodyExecution {
	s.EndDate = &v
	return s
}

func (s *StartExecutionResponseBodyExecution) SetExecutedBy(v string) *StartExecutionResponseBodyExecution {
	s.ExecutedBy = &v
	return s
}

func (s *StartExecutionResponseBodyExecution) SetIsParent(v bool) *StartExecutionResponseBodyExecution {
	s.IsParent = &v
	return s
}

func (s *StartExecutionResponseBodyExecution) SetTags(v map[string]interface{}) *StartExecutionResponseBodyExecution {
	s.Tags = v
	return s
}

func (s *StartExecutionResponseBodyExecution) SetStartDate(v string) *StartExecutionResponseBodyExecution {
	s.StartDate = &v
	return s
}

func (s *StartExecutionResponseBodyExecution) SetSafetyCheck(v string) *StartExecutionResponseBodyExecution {
	s.SafetyCheck = &v
	return s
}

func (s *StartExecutionResponseBodyExecution) SetMode(v string) *StartExecutionResponseBodyExecution {
	s.Mode = &v
	return s
}

func (s *StartExecutionResponseBodyExecution) SetTemplateName(v string) *StartExecutionResponseBodyExecution {
	s.TemplateName = &v
	return s
}

func (s *StartExecutionResponseBodyExecution) SetCreateDate(v string) *StartExecutionResponseBodyExecution {
	s.CreateDate = &v
	return s
}

func (s *StartExecutionResponseBodyExecution) SetTemplateVersion(v string) *StartExecutionResponseBodyExecution {
	s.TemplateVersion = &v
	return s
}

func (s *StartExecutionResponseBodyExecution) SetExecutionId(v string) *StartExecutionResponseBodyExecution {
	s.ExecutionId = &v
	return s
}

func (s *StartExecutionResponseBodyExecution) SetParameters(v string) *StartExecutionResponseBodyExecution {
	s.Parameters = &v
	return s
}

func (s *StartExecutionResponseBodyExecution) SetDescription(v string) *StartExecutionResponseBodyExecution {
	s.Description = &v
	return s
}

func (s *StartExecutionResponseBodyExecution) SetCounters(v map[string]interface{}) *StartExecutionResponseBodyExecution {
	s.Counters = v
	return s
}

func (s *StartExecutionResponseBodyExecution) SetUpdateDate(v string) *StartExecutionResponseBodyExecution {
	s.UpdateDate = &v
	return s
}

func (s *StartExecutionResponseBodyExecution) SetResourceGroupId(v string) *StartExecutionResponseBodyExecution {
	s.ResourceGroupId = &v
	return s
}

func (s *StartExecutionResponseBodyExecution) SetParentExecutionId(v string) *StartExecutionResponseBodyExecution {
	s.ParentExecutionId = &v
	return s
}

func (s *StartExecutionResponseBodyExecution) SetRamRole(v string) *StartExecutionResponseBodyExecution {
	s.RamRole = &v
	return s
}

func (s *StartExecutionResponseBodyExecution) SetTemplateId(v string) *StartExecutionResponseBodyExecution {
	s.TemplateId = &v
	return s
}

func (s *StartExecutionResponseBodyExecution) SetStatusMessage(v string) *StartExecutionResponseBodyExecution {
	s.StatusMessage = &v
	return s
}

func (s *StartExecutionResponseBodyExecution) SetLoopMode(v string) *StartExecutionResponseBodyExecution {
	s.LoopMode = &v
	return s
}

func (s *StartExecutionResponseBodyExecution) SetCurrentTasks(v []*StartExecutionResponseBodyExecutionCurrentTasks) *StartExecutionResponseBodyExecution {
	s.CurrentTasks = v
	return s
}

type StartExecutionResponseBodyExecutionCurrentTasks struct {
	TaskExecutionId *string `json:"TaskExecutionId,omitempty" xml:"TaskExecutionId,omitempty"`
	TaskName        *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TaskAction      *string `json:"TaskAction,omitempty" xml:"TaskAction,omitempty"`
}

func (s StartExecutionResponseBodyExecutionCurrentTasks) String() string {
	return tea.Prettify(s)
}

func (s StartExecutionResponseBodyExecutionCurrentTasks) GoString() string {
	return s.String()
}

func (s *StartExecutionResponseBodyExecutionCurrentTasks) SetTaskExecutionId(v string) *StartExecutionResponseBodyExecutionCurrentTasks {
	s.TaskExecutionId = &v
	return s
}

func (s *StartExecutionResponseBodyExecutionCurrentTasks) SetTaskName(v string) *StartExecutionResponseBodyExecutionCurrentTasks {
	s.TaskName = &v
	return s
}

func (s *StartExecutionResponseBodyExecutionCurrentTasks) SetTaskAction(v string) *StartExecutionResponseBodyExecutionCurrentTasks {
	s.TaskAction = &v
	return s
}

type StartExecutionResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartExecutionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartExecutionResponse) String() string {
	return tea.Prettify(s)
}

func (s StartExecutionResponse) GoString() string {
	return s.String()
}

func (s *StartExecutionResponse) SetHeaders(v map[string]*string) *StartExecutionResponse {
	s.Headers = v
	return s
}

func (s *StartExecutionResponse) SetBody(v *StartExecutionResponseBody) *StartExecutionResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	RegionId     *string                `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceIds  map[string]interface{} `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty"`
	ResourceType *string                `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tags         map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetRegionId(v string) *TagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceIds(v map[string]interface{}) *TagResourcesRequest {
	s.ResourceIds = v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetTags(v map[string]interface{}) *TagResourcesRequest {
	s.Tags = v
	return s
}

type TagResourcesShrinkRequest struct {
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceIdsShrink *string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty"`
	ResourceType      *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagsShrink        *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s TagResourcesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesShrinkRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesShrinkRequest) SetRegionId(v string) *TagResourcesShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesShrinkRequest) SetResourceIdsShrink(v string) *TagResourcesShrinkRequest {
	s.ResourceIdsShrink = &v
	return s
}

func (s *TagResourcesShrinkRequest) SetResourceType(v string) *TagResourcesShrinkRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesShrinkRequest) SetTagsShrink(v string) *TagResourcesShrinkRequest {
	s.TagsShrink = &v
	return s
}

type TagResourcesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBody) SetRequestId(v string) *TagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type TagResourcesResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponse) GoString() string {
	return s.String()
}

func (s *TagResourcesResponse) SetHeaders(v map[string]*string) *TagResourcesResponse {
	s.Headers = v
	return s
}

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type TriggerExecutionRequest struct {
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ExecutionId *string `json:"ExecutionId,omitempty" xml:"ExecutionId,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Content     *string `json:"Content,omitempty" xml:"Content,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s TriggerExecutionRequest) String() string {
	return tea.Prettify(s)
}

func (s TriggerExecutionRequest) GoString() string {
	return s.String()
}

func (s *TriggerExecutionRequest) SetRegionId(v string) *TriggerExecutionRequest {
	s.RegionId = &v
	return s
}

func (s *TriggerExecutionRequest) SetExecutionId(v string) *TriggerExecutionRequest {
	s.ExecutionId = &v
	return s
}

func (s *TriggerExecutionRequest) SetType(v string) *TriggerExecutionRequest {
	s.Type = &v
	return s
}

func (s *TriggerExecutionRequest) SetContent(v string) *TriggerExecutionRequest {
	s.Content = &v
	return s
}

func (s *TriggerExecutionRequest) SetClientToken(v string) *TriggerExecutionRequest {
	s.ClientToken = &v
	return s
}

type TriggerExecutionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TriggerExecutionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TriggerExecutionResponseBody) GoString() string {
	return s.String()
}

func (s *TriggerExecutionResponseBody) SetRequestId(v string) *TriggerExecutionResponseBody {
	s.RequestId = &v
	return s
}

type TriggerExecutionResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TriggerExecutionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TriggerExecutionResponse) String() string {
	return tea.Prettify(s)
}

func (s TriggerExecutionResponse) GoString() string {
	return s.String()
}

func (s *TriggerExecutionResponse) SetHeaders(v map[string]*string) *TriggerExecutionResponse {
	s.Headers = v
	return s
}

func (s *TriggerExecutionResponse) SetBody(v *TriggerExecutionResponseBody) *TriggerExecutionResponse {
	s.Body = v
	return s
}

type UntagResourcesRequest struct {
	RegionId     *string                `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceIds  map[string]interface{} `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty"`
	ResourceType *string                `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKeys      map[string]interface{} `json:"TagKeys,omitempty" xml:"TagKeys,omitempty"`
	All          *bool                  `json:"All,omitempty" xml:"All,omitempty"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetRegionId(v string) *UntagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceIds(v map[string]interface{}) *UntagResourcesRequest {
	s.ResourceIds = v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetTagKeys(v map[string]interface{}) *UntagResourcesRequest {
	s.TagKeys = v
	return s
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

type UntagResourcesShrinkRequest struct {
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceIdsShrink *string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty"`
	ResourceType      *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKeysShrink     *string `json:"TagKeys,omitempty" xml:"TagKeys,omitempty"`
	All               *bool   `json:"All,omitempty" xml:"All,omitempty"`
}

func (s UntagResourcesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesShrinkRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesShrinkRequest) SetRegionId(v string) *UntagResourcesShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UntagResourcesShrinkRequest) SetResourceIdsShrink(v string) *UntagResourcesShrinkRequest {
	s.ResourceIdsShrink = &v
	return s
}

func (s *UntagResourcesShrinkRequest) SetResourceType(v string) *UntagResourcesShrinkRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesShrinkRequest) SetTagKeysShrink(v string) *UntagResourcesShrinkRequest {
	s.TagKeysShrink = &v
	return s
}

func (s *UntagResourcesShrinkRequest) SetAll(v bool) *UntagResourcesShrinkRequest {
	s.All = &v
	return s
}

type UntagResourcesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UntagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBody) SetRequestId(v string) *UntagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type UntagResourcesResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UntagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponse) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponse) SetHeaders(v map[string]*string) *UntagResourcesResponse {
	s.Headers = v
	return s
}

func (s *UntagResourcesResponse) SetBody(v *UntagResourcesResponseBody) *UntagResourcesResponse {
	s.Body = v
	return s
}

type UpdateApplicationGroupRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s UpdateApplicationGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateApplicationGroupRequest) SetRegionId(v string) *UpdateApplicationGroupRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateApplicationGroupRequest) SetName(v string) *UpdateApplicationGroupRequest {
	s.Name = &v
	return s
}

func (s *UpdateApplicationGroupRequest) SetScalingGroupId(v string) *UpdateApplicationGroupRequest {
	s.ScalingGroupId = &v
	return s
}

type UpdateApplicationGroupResponseBody struct {
	RequestId        *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ApplicationGroup *UpdateApplicationGroupResponseBodyApplicationGroup `json:"ApplicationGroup,omitempty" xml:"ApplicationGroup,omitempty" type:"Struct"`
}

func (s UpdateApplicationGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateApplicationGroupResponseBody) SetRequestId(v string) *UpdateApplicationGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateApplicationGroupResponseBody) SetApplicationGroup(v *UpdateApplicationGroupResponseBodyApplicationGroup) *UpdateApplicationGroupResponseBody {
	s.ApplicationGroup = v
	return s
}

type UpdateApplicationGroupResponseBodyApplicationGroup struct {
	DeployRegionId  *string `json:"DeployRegionId,omitempty" xml:"DeployRegionId,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	UpdatedDate     *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
	CreatedDate     *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Environment     *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	CreateType      *string `json:"CreateType,omitempty" xml:"CreateType,omitempty"`
	ScalingGroupId  *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	ImportClusterId *string `json:"ImportClusterId,omitempty" xml:"ImportClusterId,omitempty"`
}

func (s UpdateApplicationGroupResponseBodyApplicationGroup) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationGroupResponseBodyApplicationGroup) GoString() string {
	return s.String()
}

func (s *UpdateApplicationGroupResponseBodyApplicationGroup) SetDeployRegionId(v string) *UpdateApplicationGroupResponseBodyApplicationGroup {
	s.DeployRegionId = &v
	return s
}

func (s *UpdateApplicationGroupResponseBodyApplicationGroup) SetDescription(v string) *UpdateApplicationGroupResponseBodyApplicationGroup {
	s.Description = &v
	return s
}

func (s *UpdateApplicationGroupResponseBodyApplicationGroup) SetUpdatedDate(v string) *UpdateApplicationGroupResponseBodyApplicationGroup {
	s.UpdatedDate = &v
	return s
}

func (s *UpdateApplicationGroupResponseBodyApplicationGroup) SetCreatedDate(v string) *UpdateApplicationGroupResponseBodyApplicationGroup {
	s.CreatedDate = &v
	return s
}

func (s *UpdateApplicationGroupResponseBodyApplicationGroup) SetApplicationName(v string) *UpdateApplicationGroupResponseBodyApplicationGroup {
	s.ApplicationName = &v
	return s
}

func (s *UpdateApplicationGroupResponseBodyApplicationGroup) SetName(v string) *UpdateApplicationGroupResponseBodyApplicationGroup {
	s.Name = &v
	return s
}

func (s *UpdateApplicationGroupResponseBodyApplicationGroup) SetEnvironment(v string) *UpdateApplicationGroupResponseBodyApplicationGroup {
	s.Environment = &v
	return s
}

func (s *UpdateApplicationGroupResponseBodyApplicationGroup) SetCreateType(v string) *UpdateApplicationGroupResponseBodyApplicationGroup {
	s.CreateType = &v
	return s
}

func (s *UpdateApplicationGroupResponseBodyApplicationGroup) SetScalingGroupId(v string) *UpdateApplicationGroupResponseBodyApplicationGroup {
	s.ScalingGroupId = &v
	return s
}

func (s *UpdateApplicationGroupResponseBodyApplicationGroup) SetImportClusterId(v string) *UpdateApplicationGroupResponseBodyApplicationGroup {
	s.ImportClusterId = &v
	return s
}

type UpdateApplicationGroupResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateApplicationGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateApplicationGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateApplicationGroupResponse) SetHeaders(v map[string]*string) *UpdateApplicationGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateApplicationGroupResponse) SetBody(v *UpdateApplicationGroupResponseBody) *UpdateApplicationGroupResponse {
	s.Body = v
	return s
}

type UpdateExecutionRequest struct {
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ExecutionId *string `json:"ExecutionId,omitempty" xml:"ExecutionId,omitempty"`
	Parameters  *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s UpdateExecutionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateExecutionRequest) GoString() string {
	return s.String()
}

func (s *UpdateExecutionRequest) SetRegionId(v string) *UpdateExecutionRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateExecutionRequest) SetExecutionId(v string) *UpdateExecutionRequest {
	s.ExecutionId = &v
	return s
}

func (s *UpdateExecutionRequest) SetParameters(v string) *UpdateExecutionRequest {
	s.Parameters = &v
	return s
}

func (s *UpdateExecutionRequest) SetClientToken(v string) *UpdateExecutionRequest {
	s.ClientToken = &v
	return s
}

type UpdateExecutionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateExecutionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateExecutionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateExecutionResponseBody) SetRequestId(v string) *UpdateExecutionResponseBody {
	s.RequestId = &v
	return s
}

type UpdateExecutionResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateExecutionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateExecutionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateExecutionResponse) GoString() string {
	return s.String()
}

func (s *UpdateExecutionResponse) SetHeaders(v map[string]*string) *UpdateExecutionResponse {
	s.Headers = v
	return s
}

func (s *UpdateExecutionResponse) SetBody(v *UpdateExecutionResponseBody) *UpdateExecutionResponse {
	s.Body = v
	return s
}

type UpdateInstanceInformationRequest struct {
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AgentVersion    *string `json:"AgentVersion,omitempty" xml:"AgentVersion,omitempty"`
	PlatformType    *string `json:"PlatformType,omitempty" xml:"PlatformType,omitempty"`
	PlatformName    *string `json:"PlatformName,omitempty" xml:"PlatformName,omitempty"`
	PlatformVersion *string `json:"PlatformVersion,omitempty" xml:"PlatformVersion,omitempty"`
	IpAddress       *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	ComputerName    *string `json:"ComputerName,omitempty" xml:"ComputerName,omitempty"`
	AgentName       *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
}

func (s UpdateInstanceInformationRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceInformationRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceInformationRequest) SetRegionId(v string) *UpdateInstanceInformationRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateInstanceInformationRequest) SetInstanceId(v string) *UpdateInstanceInformationRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateInstanceInformationRequest) SetAgentVersion(v string) *UpdateInstanceInformationRequest {
	s.AgentVersion = &v
	return s
}

func (s *UpdateInstanceInformationRequest) SetPlatformType(v string) *UpdateInstanceInformationRequest {
	s.PlatformType = &v
	return s
}

func (s *UpdateInstanceInformationRequest) SetPlatformName(v string) *UpdateInstanceInformationRequest {
	s.PlatformName = &v
	return s
}

func (s *UpdateInstanceInformationRequest) SetPlatformVersion(v string) *UpdateInstanceInformationRequest {
	s.PlatformVersion = &v
	return s
}

func (s *UpdateInstanceInformationRequest) SetIpAddress(v string) *UpdateInstanceInformationRequest {
	s.IpAddress = &v
	return s
}

func (s *UpdateInstanceInformationRequest) SetComputerName(v string) *UpdateInstanceInformationRequest {
	s.ComputerName = &v
	return s
}

func (s *UpdateInstanceInformationRequest) SetAgentName(v string) *UpdateInstanceInformationRequest {
	s.AgentName = &v
	return s
}

type UpdateInstanceInformationResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateInstanceInformationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceInformationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceInformationResponseBody) SetRequestId(v string) *UpdateInstanceInformationResponseBody {
	s.RequestId = &v
	return s
}

type UpdateInstanceInformationResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateInstanceInformationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateInstanceInformationResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceInformationResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstanceInformationResponse) SetHeaders(v map[string]*string) *UpdateInstanceInformationResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstanceInformationResponse) SetBody(v *UpdateInstanceInformationResponseBody) *UpdateInstanceInformationResponse {
	s.Body = v
	return s
}

type UpdateParameterRequest struct {
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value           *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Tags            *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s UpdateParameterRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateParameterRequest) GoString() string {
	return s.String()
}

func (s *UpdateParameterRequest) SetRegionId(v string) *UpdateParameterRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateParameterRequest) SetName(v string) *UpdateParameterRequest {
	s.Name = &v
	return s
}

func (s *UpdateParameterRequest) SetValue(v string) *UpdateParameterRequest {
	s.Value = &v
	return s
}

func (s *UpdateParameterRequest) SetDescription(v string) *UpdateParameterRequest {
	s.Description = &v
	return s
}

func (s *UpdateParameterRequest) SetTags(v string) *UpdateParameterRequest {
	s.Tags = &v
	return s
}

func (s *UpdateParameterRequest) SetResourceGroupId(v string) *UpdateParameterRequest {
	s.ResourceGroupId = &v
	return s
}

type UpdateParameterResponseBody struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Parameter *UpdateParameterResponseBodyParameter `json:"Parameter,omitempty" xml:"Parameter,omitempty" type:"Struct"`
}

func (s UpdateParameterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateParameterResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateParameterResponseBody) SetRequestId(v string) *UpdateParameterResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateParameterResponseBody) SetParameter(v *UpdateParameterResponseBodyParameter) *UpdateParameterResponseBody {
	s.Parameter = v
	return s
}

type UpdateParameterResponseBodyParameter struct {
	Type             *string `json:"Type,omitempty" xml:"Type,omitempty"`
	UpdatedDate      *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
	UpdatedBy        *string `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	Tags             *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Constraints      *string `json:"Constraints,omitempty" xml:"Constraints,omitempty"`
	ResourceGroupId  *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	CreatedBy        *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	CreatedDate      *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	ParameterVersion *int32  `json:"ParameterVersion,omitempty" xml:"ParameterVersion,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id               *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ShareType        *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
}

func (s UpdateParameterResponseBodyParameter) String() string {
	return tea.Prettify(s)
}

func (s UpdateParameterResponseBodyParameter) GoString() string {
	return s.String()
}

func (s *UpdateParameterResponseBodyParameter) SetType(v string) *UpdateParameterResponseBodyParameter {
	s.Type = &v
	return s
}

func (s *UpdateParameterResponseBodyParameter) SetUpdatedDate(v string) *UpdateParameterResponseBodyParameter {
	s.UpdatedDate = &v
	return s
}

func (s *UpdateParameterResponseBodyParameter) SetUpdatedBy(v string) *UpdateParameterResponseBodyParameter {
	s.UpdatedBy = &v
	return s
}

func (s *UpdateParameterResponseBodyParameter) SetTags(v string) *UpdateParameterResponseBodyParameter {
	s.Tags = &v
	return s
}

func (s *UpdateParameterResponseBodyParameter) SetDescription(v string) *UpdateParameterResponseBodyParameter {
	s.Description = &v
	return s
}

func (s *UpdateParameterResponseBodyParameter) SetConstraints(v string) *UpdateParameterResponseBodyParameter {
	s.Constraints = &v
	return s
}

func (s *UpdateParameterResponseBodyParameter) SetResourceGroupId(v string) *UpdateParameterResponseBodyParameter {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateParameterResponseBodyParameter) SetCreatedBy(v string) *UpdateParameterResponseBodyParameter {
	s.CreatedBy = &v
	return s
}

func (s *UpdateParameterResponseBodyParameter) SetCreatedDate(v string) *UpdateParameterResponseBodyParameter {
	s.CreatedDate = &v
	return s
}

func (s *UpdateParameterResponseBodyParameter) SetParameterVersion(v int32) *UpdateParameterResponseBodyParameter {
	s.ParameterVersion = &v
	return s
}

func (s *UpdateParameterResponseBodyParameter) SetName(v string) *UpdateParameterResponseBodyParameter {
	s.Name = &v
	return s
}

func (s *UpdateParameterResponseBodyParameter) SetId(v string) *UpdateParameterResponseBodyParameter {
	s.Id = &v
	return s
}

func (s *UpdateParameterResponseBodyParameter) SetShareType(v string) *UpdateParameterResponseBodyParameter {
	s.ShareType = &v
	return s
}

type UpdateParameterResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateParameterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateParameterResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateParameterResponse) GoString() string {
	return s.String()
}

func (s *UpdateParameterResponse) SetHeaders(v map[string]*string) *UpdateParameterResponse {
	s.Headers = v
	return s
}

func (s *UpdateParameterResponse) SetBody(v *UpdateParameterResponseBody) *UpdateParameterResponse {
	s.Body = v
	return s
}

type UpdatePatchBaselineRequest struct {
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ClientToken   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ApprovalRules *string `json:"ApprovalRules,omitempty" xml:"ApprovalRules,omitempty"`
}

func (s UpdatePatchBaselineRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePatchBaselineRequest) GoString() string {
	return s.String()
}

func (s *UpdatePatchBaselineRequest) SetRegionId(v string) *UpdatePatchBaselineRequest {
	s.RegionId = &v
	return s
}

func (s *UpdatePatchBaselineRequest) SetName(v string) *UpdatePatchBaselineRequest {
	s.Name = &v
	return s
}

func (s *UpdatePatchBaselineRequest) SetDescription(v string) *UpdatePatchBaselineRequest {
	s.Description = &v
	return s
}

func (s *UpdatePatchBaselineRequest) SetClientToken(v string) *UpdatePatchBaselineRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdatePatchBaselineRequest) SetApprovalRules(v string) *UpdatePatchBaselineRequest {
	s.ApprovalRules = &v
	return s
}

type UpdatePatchBaselineResponseBody struct {
	RequestId     *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PatchBaseline *UpdatePatchBaselineResponseBodyPatchBaseline `json:"PatchBaseline,omitempty" xml:"PatchBaseline,omitempty" type:"Struct"`
}

func (s UpdatePatchBaselineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdatePatchBaselineResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePatchBaselineResponseBody) SetRequestId(v string) *UpdatePatchBaselineResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePatchBaselineResponseBody) SetPatchBaseline(v *UpdatePatchBaselineResponseBodyPatchBaseline) *UpdatePatchBaselineResponseBody {
	s.PatchBaseline = v
	return s
}

type UpdatePatchBaselineResponseBodyPatchBaseline struct {
	OperationSystem *string `json:"OperationSystem,omitempty" xml:"OperationSystem,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	UpdatedDate     *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
	UpdatedBy       *string `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	CreatedBy       *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	CreatedDate     *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ApprovalRules   *string `json:"ApprovalRules,omitempty" xml:"ApprovalRules,omitempty"`
	Id              *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ShareType       *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
}

func (s UpdatePatchBaselineResponseBodyPatchBaseline) String() string {
	return tea.Prettify(s)
}

func (s UpdatePatchBaselineResponseBodyPatchBaseline) GoString() string {
	return s.String()
}

func (s *UpdatePatchBaselineResponseBodyPatchBaseline) SetOperationSystem(v string) *UpdatePatchBaselineResponseBodyPatchBaseline {
	s.OperationSystem = &v
	return s
}

func (s *UpdatePatchBaselineResponseBodyPatchBaseline) SetDescription(v string) *UpdatePatchBaselineResponseBodyPatchBaseline {
	s.Description = &v
	return s
}

func (s *UpdatePatchBaselineResponseBodyPatchBaseline) SetUpdatedDate(v string) *UpdatePatchBaselineResponseBodyPatchBaseline {
	s.UpdatedDate = &v
	return s
}

func (s *UpdatePatchBaselineResponseBodyPatchBaseline) SetUpdatedBy(v string) *UpdatePatchBaselineResponseBodyPatchBaseline {
	s.UpdatedBy = &v
	return s
}

func (s *UpdatePatchBaselineResponseBodyPatchBaseline) SetCreatedBy(v string) *UpdatePatchBaselineResponseBodyPatchBaseline {
	s.CreatedBy = &v
	return s
}

func (s *UpdatePatchBaselineResponseBodyPatchBaseline) SetCreatedDate(v string) *UpdatePatchBaselineResponseBodyPatchBaseline {
	s.CreatedDate = &v
	return s
}

func (s *UpdatePatchBaselineResponseBodyPatchBaseline) SetName(v string) *UpdatePatchBaselineResponseBodyPatchBaseline {
	s.Name = &v
	return s
}

func (s *UpdatePatchBaselineResponseBodyPatchBaseline) SetApprovalRules(v string) *UpdatePatchBaselineResponseBodyPatchBaseline {
	s.ApprovalRules = &v
	return s
}

func (s *UpdatePatchBaselineResponseBodyPatchBaseline) SetId(v string) *UpdatePatchBaselineResponseBodyPatchBaseline {
	s.Id = &v
	return s
}

func (s *UpdatePatchBaselineResponseBodyPatchBaseline) SetShareType(v string) *UpdatePatchBaselineResponseBodyPatchBaseline {
	s.ShareType = &v
	return s
}

type UpdatePatchBaselineResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdatePatchBaselineResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdatePatchBaselineResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdatePatchBaselineResponse) GoString() string {
	return s.String()
}

func (s *UpdatePatchBaselineResponse) SetHeaders(v map[string]*string) *UpdatePatchBaselineResponse {
	s.Headers = v
	return s
}

func (s *UpdatePatchBaselineResponse) SetBody(v *UpdatePatchBaselineResponseBody) *UpdatePatchBaselineResponse {
	s.Body = v
	return s
}

type UpdateSecretParameterRequest struct {
	RegionId        *string                `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Name            *string                `json:"Name,omitempty" xml:"Name,omitempty"`
	Value           *string                `json:"Value,omitempty" xml:"Value,omitempty"`
	Description     *string                `json:"Description,omitempty" xml:"Description,omitempty"`
	Tags            map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	ResourceGroupId *string                `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s UpdateSecretParameterRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSecretParameterRequest) GoString() string {
	return s.String()
}

func (s *UpdateSecretParameterRequest) SetRegionId(v string) *UpdateSecretParameterRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateSecretParameterRequest) SetName(v string) *UpdateSecretParameterRequest {
	s.Name = &v
	return s
}

func (s *UpdateSecretParameterRequest) SetValue(v string) *UpdateSecretParameterRequest {
	s.Value = &v
	return s
}

func (s *UpdateSecretParameterRequest) SetDescription(v string) *UpdateSecretParameterRequest {
	s.Description = &v
	return s
}

func (s *UpdateSecretParameterRequest) SetTags(v map[string]interface{}) *UpdateSecretParameterRequest {
	s.Tags = v
	return s
}

func (s *UpdateSecretParameterRequest) SetResourceGroupId(v string) *UpdateSecretParameterRequest {
	s.ResourceGroupId = &v
	return s
}

type UpdateSecretParameterShrinkRequest struct {
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value           *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	TagsShrink      *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s UpdateSecretParameterShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSecretParameterShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateSecretParameterShrinkRequest) SetRegionId(v string) *UpdateSecretParameterShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateSecretParameterShrinkRequest) SetName(v string) *UpdateSecretParameterShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateSecretParameterShrinkRequest) SetValue(v string) *UpdateSecretParameterShrinkRequest {
	s.Value = &v
	return s
}

func (s *UpdateSecretParameterShrinkRequest) SetDescription(v string) *UpdateSecretParameterShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateSecretParameterShrinkRequest) SetTagsShrink(v string) *UpdateSecretParameterShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *UpdateSecretParameterShrinkRequest) SetResourceGroupId(v string) *UpdateSecretParameterShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

type UpdateSecretParameterResponseBody struct {
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Parameter *UpdateSecretParameterResponseBodyParameter `json:"Parameter,omitempty" xml:"Parameter,omitempty" type:"Struct"`
}

func (s UpdateSecretParameterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSecretParameterResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSecretParameterResponseBody) SetRequestId(v string) *UpdateSecretParameterResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSecretParameterResponseBody) SetParameter(v *UpdateSecretParameterResponseBodyParameter) *UpdateSecretParameterResponseBody {
	s.Parameter = v
	return s
}

type UpdateSecretParameterResponseBodyParameter struct {
	Type             *string `json:"Type,omitempty" xml:"Type,omitempty"`
	UpdatedDate      *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
	UpdatedBy        *string `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	KeyId            *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	Tags             *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Constraints      *string `json:"Constraints,omitempty" xml:"Constraints,omitempty"`
	ResourceGroupId  *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	CreatedBy        *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	CreatedDate      *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	ParameterVersion *int32  `json:"ParameterVersion,omitempty" xml:"ParameterVersion,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id               *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ShareType        *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
}

func (s UpdateSecretParameterResponseBodyParameter) String() string {
	return tea.Prettify(s)
}

func (s UpdateSecretParameterResponseBodyParameter) GoString() string {
	return s.String()
}

func (s *UpdateSecretParameterResponseBodyParameter) SetType(v string) *UpdateSecretParameterResponseBodyParameter {
	s.Type = &v
	return s
}

func (s *UpdateSecretParameterResponseBodyParameter) SetUpdatedDate(v string) *UpdateSecretParameterResponseBodyParameter {
	s.UpdatedDate = &v
	return s
}

func (s *UpdateSecretParameterResponseBodyParameter) SetUpdatedBy(v string) *UpdateSecretParameterResponseBodyParameter {
	s.UpdatedBy = &v
	return s
}

func (s *UpdateSecretParameterResponseBodyParameter) SetKeyId(v string) *UpdateSecretParameterResponseBodyParameter {
	s.KeyId = &v
	return s
}

func (s *UpdateSecretParameterResponseBodyParameter) SetTags(v string) *UpdateSecretParameterResponseBodyParameter {
	s.Tags = &v
	return s
}

func (s *UpdateSecretParameterResponseBodyParameter) SetDescription(v string) *UpdateSecretParameterResponseBodyParameter {
	s.Description = &v
	return s
}

func (s *UpdateSecretParameterResponseBodyParameter) SetConstraints(v string) *UpdateSecretParameterResponseBodyParameter {
	s.Constraints = &v
	return s
}

func (s *UpdateSecretParameterResponseBodyParameter) SetResourceGroupId(v string) *UpdateSecretParameterResponseBodyParameter {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateSecretParameterResponseBodyParameter) SetCreatedBy(v string) *UpdateSecretParameterResponseBodyParameter {
	s.CreatedBy = &v
	return s
}

func (s *UpdateSecretParameterResponseBodyParameter) SetCreatedDate(v string) *UpdateSecretParameterResponseBodyParameter {
	s.CreatedDate = &v
	return s
}

func (s *UpdateSecretParameterResponseBodyParameter) SetParameterVersion(v int32) *UpdateSecretParameterResponseBodyParameter {
	s.ParameterVersion = &v
	return s
}

func (s *UpdateSecretParameterResponseBodyParameter) SetName(v string) *UpdateSecretParameterResponseBodyParameter {
	s.Name = &v
	return s
}

func (s *UpdateSecretParameterResponseBodyParameter) SetId(v string) *UpdateSecretParameterResponseBodyParameter {
	s.Id = &v
	return s
}

func (s *UpdateSecretParameterResponseBodyParameter) SetShareType(v string) *UpdateSecretParameterResponseBodyParameter {
	s.ShareType = &v
	return s
}

type UpdateSecretParameterResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateSecretParameterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateSecretParameterResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSecretParameterResponse) GoString() string {
	return s.String()
}

func (s *UpdateSecretParameterResponse) SetHeaders(v map[string]*string) *UpdateSecretParameterResponse {
	s.Headers = v
	return s
}

func (s *UpdateSecretParameterResponse) SetBody(v *UpdateSecretParameterResponseBody) *UpdateSecretParameterResponse {
	s.Body = v
	return s
}

type UpdateStateConfigurationRequest struct {
	RegionId             *string                `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Description          *string                `json:"Description,omitempty" xml:"Description,omitempty"`
	Parameters           map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	ConfigureMode        *string                `json:"ConfigureMode,omitempty" xml:"ConfigureMode,omitempty"`
	ScheduleType         *string                `json:"ScheduleType,omitempty" xml:"ScheduleType,omitempty"`
	ScheduleExpression   *string                `json:"ScheduleExpression,omitempty" xml:"ScheduleExpression,omitempty"`
	Targets              *string                `json:"Targets,omitempty" xml:"Targets,omitempty"`
	ClientToken          *string                `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	StateConfigurationId *string                `json:"StateConfigurationId,omitempty" xml:"StateConfigurationId,omitempty"`
	Tags                 map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	ResourceGroupId      *string                `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s UpdateStateConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateStateConfigurationRequest) GoString() string {
	return s.String()
}

func (s *UpdateStateConfigurationRequest) SetRegionId(v string) *UpdateStateConfigurationRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateStateConfigurationRequest) SetDescription(v string) *UpdateStateConfigurationRequest {
	s.Description = &v
	return s
}

func (s *UpdateStateConfigurationRequest) SetParameters(v map[string]interface{}) *UpdateStateConfigurationRequest {
	s.Parameters = v
	return s
}

func (s *UpdateStateConfigurationRequest) SetConfigureMode(v string) *UpdateStateConfigurationRequest {
	s.ConfigureMode = &v
	return s
}

func (s *UpdateStateConfigurationRequest) SetScheduleType(v string) *UpdateStateConfigurationRequest {
	s.ScheduleType = &v
	return s
}

func (s *UpdateStateConfigurationRequest) SetScheduleExpression(v string) *UpdateStateConfigurationRequest {
	s.ScheduleExpression = &v
	return s
}

func (s *UpdateStateConfigurationRequest) SetTargets(v string) *UpdateStateConfigurationRequest {
	s.Targets = &v
	return s
}

func (s *UpdateStateConfigurationRequest) SetClientToken(v string) *UpdateStateConfigurationRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateStateConfigurationRequest) SetStateConfigurationId(v string) *UpdateStateConfigurationRequest {
	s.StateConfigurationId = &v
	return s
}

func (s *UpdateStateConfigurationRequest) SetTags(v map[string]interface{}) *UpdateStateConfigurationRequest {
	s.Tags = v
	return s
}

func (s *UpdateStateConfigurationRequest) SetResourceGroupId(v string) *UpdateStateConfigurationRequest {
	s.ResourceGroupId = &v
	return s
}

type UpdateStateConfigurationShrinkRequest struct {
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ParametersShrink     *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	ConfigureMode        *string `json:"ConfigureMode,omitempty" xml:"ConfigureMode,omitempty"`
	ScheduleType         *string `json:"ScheduleType,omitempty" xml:"ScheduleType,omitempty"`
	ScheduleExpression   *string `json:"ScheduleExpression,omitempty" xml:"ScheduleExpression,omitempty"`
	Targets              *string `json:"Targets,omitempty" xml:"Targets,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	StateConfigurationId *string `json:"StateConfigurationId,omitempty" xml:"StateConfigurationId,omitempty"`
	TagsShrink           *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s UpdateStateConfigurationShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateStateConfigurationShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateStateConfigurationShrinkRequest) SetRegionId(v string) *UpdateStateConfigurationShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateStateConfigurationShrinkRequest) SetDescription(v string) *UpdateStateConfigurationShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateStateConfigurationShrinkRequest) SetParametersShrink(v string) *UpdateStateConfigurationShrinkRequest {
	s.ParametersShrink = &v
	return s
}

func (s *UpdateStateConfigurationShrinkRequest) SetConfigureMode(v string) *UpdateStateConfigurationShrinkRequest {
	s.ConfigureMode = &v
	return s
}

func (s *UpdateStateConfigurationShrinkRequest) SetScheduleType(v string) *UpdateStateConfigurationShrinkRequest {
	s.ScheduleType = &v
	return s
}

func (s *UpdateStateConfigurationShrinkRequest) SetScheduleExpression(v string) *UpdateStateConfigurationShrinkRequest {
	s.ScheduleExpression = &v
	return s
}

func (s *UpdateStateConfigurationShrinkRequest) SetTargets(v string) *UpdateStateConfigurationShrinkRequest {
	s.Targets = &v
	return s
}

func (s *UpdateStateConfigurationShrinkRequest) SetClientToken(v string) *UpdateStateConfigurationShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateStateConfigurationShrinkRequest) SetStateConfigurationId(v string) *UpdateStateConfigurationShrinkRequest {
	s.StateConfigurationId = &v
	return s
}

func (s *UpdateStateConfigurationShrinkRequest) SetTagsShrink(v string) *UpdateStateConfigurationShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *UpdateStateConfigurationShrinkRequest) SetResourceGroupId(v string) *UpdateStateConfigurationShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

type UpdateStateConfigurationResponseBody struct {
	RequestId          *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StateConfiguration []*UpdateStateConfigurationResponseBodyStateConfiguration `json:"StateConfiguration,omitempty" xml:"StateConfiguration,omitempty" type:"Repeated"`
}

func (s UpdateStateConfigurationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateStateConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateStateConfigurationResponseBody) SetRequestId(v string) *UpdateStateConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateStateConfigurationResponseBody) SetStateConfiguration(v []*UpdateStateConfigurationResponseBodyStateConfiguration) *UpdateStateConfigurationResponseBody {
	s.StateConfiguration = v
	return s
}

type UpdateStateConfigurationResponseBodyStateConfiguration struct {
	UpdateTime           *string                `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	CreateTime           *string                `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Targets              *string                `json:"Targets,omitempty" xml:"Targets,omitempty"`
	Tags                 map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	StateConfigurationId *string                `json:"StateConfigurationId,omitempty" xml:"StateConfigurationId,omitempty"`
	ScheduleExpression   *string                `json:"ScheduleExpression,omitempty" xml:"ScheduleExpression,omitempty"`
	TemplateName         *string                `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateVersion      *string                `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	ConfigureMode        *string                `json:"ConfigureMode,omitempty" xml:"ConfigureMode,omitempty"`
	ScheduleType         *string                `json:"ScheduleType,omitempty" xml:"ScheduleType,omitempty"`
	Parameters           *string                `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	Description          *string                `json:"Description,omitempty" xml:"Description,omitempty"`
	ResourceGroupId      *string                `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	TemplateId           *string                `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s UpdateStateConfigurationResponseBodyStateConfiguration) String() string {
	return tea.Prettify(s)
}

func (s UpdateStateConfigurationResponseBodyStateConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateStateConfigurationResponseBodyStateConfiguration) SetUpdateTime(v string) *UpdateStateConfigurationResponseBodyStateConfiguration {
	s.UpdateTime = &v
	return s
}

func (s *UpdateStateConfigurationResponseBodyStateConfiguration) SetCreateTime(v string) *UpdateStateConfigurationResponseBodyStateConfiguration {
	s.CreateTime = &v
	return s
}

func (s *UpdateStateConfigurationResponseBodyStateConfiguration) SetTargets(v string) *UpdateStateConfigurationResponseBodyStateConfiguration {
	s.Targets = &v
	return s
}

func (s *UpdateStateConfigurationResponseBodyStateConfiguration) SetTags(v map[string]interface{}) *UpdateStateConfigurationResponseBodyStateConfiguration {
	s.Tags = v
	return s
}

func (s *UpdateStateConfigurationResponseBodyStateConfiguration) SetStateConfigurationId(v string) *UpdateStateConfigurationResponseBodyStateConfiguration {
	s.StateConfigurationId = &v
	return s
}

func (s *UpdateStateConfigurationResponseBodyStateConfiguration) SetScheduleExpression(v string) *UpdateStateConfigurationResponseBodyStateConfiguration {
	s.ScheduleExpression = &v
	return s
}

func (s *UpdateStateConfigurationResponseBodyStateConfiguration) SetTemplateName(v string) *UpdateStateConfigurationResponseBodyStateConfiguration {
	s.TemplateName = &v
	return s
}

func (s *UpdateStateConfigurationResponseBodyStateConfiguration) SetTemplateVersion(v string) *UpdateStateConfigurationResponseBodyStateConfiguration {
	s.TemplateVersion = &v
	return s
}

func (s *UpdateStateConfigurationResponseBodyStateConfiguration) SetConfigureMode(v string) *UpdateStateConfigurationResponseBodyStateConfiguration {
	s.ConfigureMode = &v
	return s
}

func (s *UpdateStateConfigurationResponseBodyStateConfiguration) SetScheduleType(v string) *UpdateStateConfigurationResponseBodyStateConfiguration {
	s.ScheduleType = &v
	return s
}

func (s *UpdateStateConfigurationResponseBodyStateConfiguration) SetParameters(v string) *UpdateStateConfigurationResponseBodyStateConfiguration {
	s.Parameters = &v
	return s
}

func (s *UpdateStateConfigurationResponseBodyStateConfiguration) SetDescription(v string) *UpdateStateConfigurationResponseBodyStateConfiguration {
	s.Description = &v
	return s
}

func (s *UpdateStateConfigurationResponseBodyStateConfiguration) SetResourceGroupId(v string) *UpdateStateConfigurationResponseBodyStateConfiguration {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateStateConfigurationResponseBodyStateConfiguration) SetTemplateId(v string) *UpdateStateConfigurationResponseBodyStateConfiguration {
	s.TemplateId = &v
	return s
}

type UpdateStateConfigurationResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateStateConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateStateConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateStateConfigurationResponse) GoString() string {
	return s.String()
}

func (s *UpdateStateConfigurationResponse) SetHeaders(v map[string]*string) *UpdateStateConfigurationResponse {
	s.Headers = v
	return s
}

func (s *UpdateStateConfigurationResponse) SetBody(v *UpdateStateConfigurationResponseBody) *UpdateStateConfigurationResponse {
	s.Body = v
	return s
}

type UpdateTemplateRequest struct {
	RegionId        *string                `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TemplateName    *string                `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	Content         *string                `json:"Content,omitempty" xml:"Content,omitempty"`
	Tags            map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	VersionName     *string                `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
	ResourceGroupId *string                `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s UpdateTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateTemplateRequest) SetRegionId(v string) *UpdateTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateTemplateRequest) SetTemplateName(v string) *UpdateTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *UpdateTemplateRequest) SetContent(v string) *UpdateTemplateRequest {
	s.Content = &v
	return s
}

func (s *UpdateTemplateRequest) SetTags(v map[string]interface{}) *UpdateTemplateRequest {
	s.Tags = v
	return s
}

func (s *UpdateTemplateRequest) SetVersionName(v string) *UpdateTemplateRequest {
	s.VersionName = &v
	return s
}

func (s *UpdateTemplateRequest) SetResourceGroupId(v string) *UpdateTemplateRequest {
	s.ResourceGroupId = &v
	return s
}

type UpdateTemplateShrinkRequest struct {
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TemplateName    *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	Content         *string `json:"Content,omitempty" xml:"Content,omitempty"`
	TagsShrink      *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	VersionName     *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s UpdateTemplateShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateTemplateShrinkRequest) SetRegionId(v string) *UpdateTemplateShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateTemplateShrinkRequest) SetTemplateName(v string) *UpdateTemplateShrinkRequest {
	s.TemplateName = &v
	return s
}

func (s *UpdateTemplateShrinkRequest) SetContent(v string) *UpdateTemplateShrinkRequest {
	s.Content = &v
	return s
}

func (s *UpdateTemplateShrinkRequest) SetTagsShrink(v string) *UpdateTemplateShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *UpdateTemplateShrinkRequest) SetVersionName(v string) *UpdateTemplateShrinkRequest {
	s.VersionName = &v
	return s
}

func (s *UpdateTemplateShrinkRequest) SetResourceGroupId(v string) *UpdateTemplateShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

type UpdateTemplateResponseBody struct {
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Template  *UpdateTemplateResponseBodyTemplate `json:"Template,omitempty" xml:"Template,omitempty" type:"Struct"`
}

func (s UpdateTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTemplateResponseBody) SetRequestId(v string) *UpdateTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTemplateResponseBody) SetTemplate(v *UpdateTemplateResponseBodyTemplate) *UpdateTemplateResponseBody {
	s.Template = v
	return s
}

type UpdateTemplateResponseBodyTemplate struct {
	Hash            *string                `json:"Hash,omitempty" xml:"Hash,omitempty"`
	UpdatedDate     *string                `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
	UpdatedBy       *string                `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	Tags            map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	TemplateName    *string                `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateVersion *string                `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	TemplateFormat  *string                `json:"TemplateFormat,omitempty" xml:"TemplateFormat,omitempty"`
	Description     *string                `json:"Description,omitempty" xml:"Description,omitempty"`
	ResourceGroupId *string                `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	CreatedBy       *string                `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	CreatedDate     *string                `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	TemplateId      *string                `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	HasTrigger      *bool                  `json:"HasTrigger,omitempty" xml:"HasTrigger,omitempty"`
	ShareType       *string                `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
}

func (s UpdateTemplateResponseBodyTemplate) String() string {
	return tea.Prettify(s)
}

func (s UpdateTemplateResponseBodyTemplate) GoString() string {
	return s.String()
}

func (s *UpdateTemplateResponseBodyTemplate) SetHash(v string) *UpdateTemplateResponseBodyTemplate {
	s.Hash = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplate) SetUpdatedDate(v string) *UpdateTemplateResponseBodyTemplate {
	s.UpdatedDate = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplate) SetUpdatedBy(v string) *UpdateTemplateResponseBodyTemplate {
	s.UpdatedBy = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplate) SetTags(v map[string]interface{}) *UpdateTemplateResponseBodyTemplate {
	s.Tags = v
	return s
}

func (s *UpdateTemplateResponseBodyTemplate) SetTemplateName(v string) *UpdateTemplateResponseBodyTemplate {
	s.TemplateName = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplate) SetTemplateVersion(v string) *UpdateTemplateResponseBodyTemplate {
	s.TemplateVersion = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplate) SetTemplateFormat(v string) *UpdateTemplateResponseBodyTemplate {
	s.TemplateFormat = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplate) SetDescription(v string) *UpdateTemplateResponseBodyTemplate {
	s.Description = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplate) SetResourceGroupId(v string) *UpdateTemplateResponseBodyTemplate {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplate) SetCreatedBy(v string) *UpdateTemplateResponseBodyTemplate {
	s.CreatedBy = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplate) SetCreatedDate(v string) *UpdateTemplateResponseBodyTemplate {
	s.CreatedDate = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplate) SetTemplateId(v string) *UpdateTemplateResponseBodyTemplate {
	s.TemplateId = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplate) SetHasTrigger(v bool) *UpdateTemplateResponseBodyTemplate {
	s.HasTrigger = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplate) SetShareType(v string) *UpdateTemplateResponseBodyTemplate {
	s.ShareType = &v
	return s
}

type UpdateTemplateResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTemplateResponse) GoString() string {
	return s.String()
}

func (s *UpdateTemplateResponse) SetHeaders(v map[string]*string) *UpdateTemplateResponse {
	s.Headers = v
	return s
}

func (s *UpdateTemplateResponse) SetBody(v *UpdateTemplateResponseBody) *UpdateTemplateResponse {
	s.Body = v
	return s
}

type ValidateTemplateContentRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Content  *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s ValidateTemplateContentRequest) String() string {
	return tea.Prettify(s)
}

func (s ValidateTemplateContentRequest) GoString() string {
	return s.String()
}

func (s *ValidateTemplateContentRequest) SetRegionId(v string) *ValidateTemplateContentRequest {
	s.RegionId = &v
	return s
}

func (s *ValidateTemplateContentRequest) SetContent(v string) *ValidateTemplateContentRequest {
	s.Content = &v
	return s
}

type ValidateTemplateContentResponseBody struct {
	Outputs    *string                                     `json:"Outputs,omitempty" xml:"Outputs,omitempty"`
	RequestId  *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Parameters *string                                     `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	RamRole    *string                                     `json:"RamRole,omitempty" xml:"RamRole,omitempty"`
	Tasks      []*ValidateTemplateContentResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
}

func (s ValidateTemplateContentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ValidateTemplateContentResponseBody) GoString() string {
	return s.String()
}

func (s *ValidateTemplateContentResponseBody) SetOutputs(v string) *ValidateTemplateContentResponseBody {
	s.Outputs = &v
	return s
}

func (s *ValidateTemplateContentResponseBody) SetRequestId(v string) *ValidateTemplateContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *ValidateTemplateContentResponseBody) SetParameters(v string) *ValidateTemplateContentResponseBody {
	s.Parameters = &v
	return s
}

func (s *ValidateTemplateContentResponseBody) SetRamRole(v string) *ValidateTemplateContentResponseBody {
	s.RamRole = &v
	return s
}

func (s *ValidateTemplateContentResponseBody) SetTasks(v []*ValidateTemplateContentResponseBodyTasks) *ValidateTemplateContentResponseBody {
	s.Tasks = v
	return s
}

type ValidateTemplateContentResponseBodyTasks struct {
	Outputs     *string `json:"Outputs,omitempty" xml:"Outputs,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Properties  *string `json:"Properties,omitempty" xml:"Properties,omitempty"`
}

func (s ValidateTemplateContentResponseBodyTasks) String() string {
	return tea.Prettify(s)
}

func (s ValidateTemplateContentResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *ValidateTemplateContentResponseBodyTasks) SetOutputs(v string) *ValidateTemplateContentResponseBodyTasks {
	s.Outputs = &v
	return s
}

func (s *ValidateTemplateContentResponseBodyTasks) SetType(v string) *ValidateTemplateContentResponseBodyTasks {
	s.Type = &v
	return s
}

func (s *ValidateTemplateContentResponseBodyTasks) SetDescription(v string) *ValidateTemplateContentResponseBodyTasks {
	s.Description = &v
	return s
}

func (s *ValidateTemplateContentResponseBodyTasks) SetName(v string) *ValidateTemplateContentResponseBodyTasks {
	s.Name = &v
	return s
}

func (s *ValidateTemplateContentResponseBodyTasks) SetProperties(v string) *ValidateTemplateContentResponseBodyTasks {
	s.Properties = &v
	return s
}

type ValidateTemplateContentResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ValidateTemplateContentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ValidateTemplateContentResponse) String() string {
	return tea.Prettify(s)
}

func (s ValidateTemplateContentResponse) GoString() string {
	return s.String()
}

func (s *ValidateTemplateContentResponse) SetHeaders(v map[string]*string) *ValidateTemplateContentResponse {
	s.Headers = v
	return s
}

func (s *ValidateTemplateContentResponse) SetBody(v *ValidateTemplateContentResponseBody) *ValidateTemplateContentResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("oos"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CancelExecutionWithOptions(request *CancelExecutionRequest, runtime *util.RuntimeOptions) (_result *CancelExecutionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CancelExecutionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CancelExecution"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelExecution(request *CancelExecutionRequest) (_result *CancelExecutionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelExecutionResponse{}
	_body, _err := client.CancelExecutionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ChangeResourceGroupWithOptions(request *ChangeResourceGroupRequest, runtime *util.RuntimeOptions) (_result *ChangeResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ChangeResourceGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ChangeResourceGroup"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ChangeResourceGroup(request *ChangeResourceGroupRequest) (_result *ChangeResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ChangeResourceGroupResponse{}
	_body, _err := client.ChangeResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateApplicationWithOptions(request *CreateApplicationRequest, runtime *util.RuntimeOptions) (_result *CreateApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateApplicationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateApplication"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateApplication(request *CreateApplicationRequest) (_result *CreateApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateApplicationResponse{}
	_body, _err := client.CreateApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateApplicationGroupWithOptions(request *CreateApplicationGroupRequest, runtime *util.RuntimeOptions) (_result *CreateApplicationGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateApplicationGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateApplicationGroup"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateApplicationGroup(request *CreateApplicationGroupRequest) (_result *CreateApplicationGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateApplicationGroupResponse{}
	_body, _err := client.CreateApplicationGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateParameterWithOptions(tmpReq *CreateParameterRequest, runtime *util.RuntimeOptions) (_result *CreateParameterResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateParameterShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("Tags"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateParameterResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateParameter"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateParameter(request *CreateParameterRequest) (_result *CreateParameterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateParameterResponse{}
	_body, _err := client.CreateParameterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreatePatchBaselineWithOptions(request *CreatePatchBaselineRequest, runtime *util.RuntimeOptions) (_result *CreatePatchBaselineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreatePatchBaselineResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreatePatchBaseline"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreatePatchBaseline(request *CreatePatchBaselineRequest) (_result *CreatePatchBaselineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreatePatchBaselineResponse{}
	_body, _err := client.CreatePatchBaselineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSecretParameterWithOptions(request *CreateSecretParameterRequest, runtime *util.RuntimeOptions) (_result *CreateSecretParameterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateSecretParameterResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateSecretParameter"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSecretParameter(request *CreateSecretParameterRequest) (_result *CreateSecretParameterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSecretParameterResponse{}
	_body, _err := client.CreateSecretParameterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateStateConfigurationWithOptions(tmpReq *CreateStateConfigurationRequest, runtime *util.RuntimeOptions) (_result *CreateStateConfigurationResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateStateConfigurationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("Tags"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateStateConfigurationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateStateConfiguration"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateStateConfiguration(request *CreateStateConfigurationRequest) (_result *CreateStateConfigurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateStateConfigurationResponse{}
	_body, _err := client.CreateStateConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTemplateWithOptions(tmpReq *CreateTemplateRequest, runtime *util.RuntimeOptions) (_result *CreateTemplateResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("Tags"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateTemplate"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTemplate(request *CreateTemplateRequest) (_result *CreateTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTemplateResponse{}
	_body, _err := client.CreateTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteApplicationWithOptions(request *DeleteApplicationRequest, runtime *util.RuntimeOptions) (_result *DeleteApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteApplicationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteApplication"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteApplication(request *DeleteApplicationRequest) (_result *DeleteApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteApplicationResponse{}
	_body, _err := client.DeleteApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteApplicationGroupWithOptions(request *DeleteApplicationGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteApplicationGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteApplicationGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteApplicationGroup"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteApplicationGroup(request *DeleteApplicationGroupRequest) (_result *DeleteApplicationGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteApplicationGroupResponse{}
	_body, _err := client.DeleteApplicationGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteExecutionsWithOptions(request *DeleteExecutionsRequest, runtime *util.RuntimeOptions) (_result *DeleteExecutionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteExecutionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteExecutions"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteExecutions(request *DeleteExecutionsRequest) (_result *DeleteExecutionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteExecutionsResponse{}
	_body, _err := client.DeleteExecutionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteParameterWithOptions(request *DeleteParameterRequest, runtime *util.RuntimeOptions) (_result *DeleteParameterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteParameterResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteParameter"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteParameter(request *DeleteParameterRequest) (_result *DeleteParameterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteParameterResponse{}
	_body, _err := client.DeleteParameterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeletePatchBaselineWithOptions(request *DeletePatchBaselineRequest, runtime *util.RuntimeOptions) (_result *DeletePatchBaselineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeletePatchBaselineResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeletePatchBaseline"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeletePatchBaseline(request *DeletePatchBaselineRequest) (_result *DeletePatchBaselineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeletePatchBaselineResponse{}
	_body, _err := client.DeletePatchBaselineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSecretParameterWithOptions(request *DeleteSecretParameterRequest, runtime *util.RuntimeOptions) (_result *DeleteSecretParameterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteSecretParameterResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteSecretParameter"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSecretParameter(request *DeleteSecretParameterRequest) (_result *DeleteSecretParameterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSecretParameterResponse{}
	_body, _err := client.DeleteSecretParameterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteStateConfigurationsWithOptions(request *DeleteStateConfigurationsRequest, runtime *util.RuntimeOptions) (_result *DeleteStateConfigurationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteStateConfigurationsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteStateConfigurations"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteStateConfigurations(request *DeleteStateConfigurationsRequest) (_result *DeleteStateConfigurationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteStateConfigurationsResponse{}
	_body, _err := client.DeleteStateConfigurationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteTemplateWithOptions(request *DeleteTemplateRequest, runtime *util.RuntimeOptions) (_result *DeleteTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteTemplate"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteTemplate(request *DeleteTemplateRequest) (_result *DeleteTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteTemplateResponse{}
	_body, _err := client.DeleteTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteTemplatesWithOptions(request *DeleteTemplatesRequest, runtime *util.RuntimeOptions) (_result *DeleteTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteTemplatesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteTemplates"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteTemplates(request *DeleteTemplatesRequest) (_result *DeleteTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteTemplatesResponse{}
	_body, _err := client.DeleteTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRegions"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GenerateExecutionPolicyWithOptions(request *GenerateExecutionPolicyRequest, runtime *util.RuntimeOptions) (_result *GenerateExecutionPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GenerateExecutionPolicyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GenerateExecutionPolicy"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GenerateExecutionPolicy(request *GenerateExecutionPolicyRequest) (_result *GenerateExecutionPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateExecutionPolicyResponse{}
	_body, _err := client.GenerateExecutionPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetApplicationWithOptions(request *GetApplicationRequest, runtime *util.RuntimeOptions) (_result *GetApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetApplicationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetApplication"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetApplication(request *GetApplicationRequest) (_result *GetApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetApplicationResponse{}
	_body, _err := client.GetApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetApplicationGroupWithOptions(request *GetApplicationGroupRequest, runtime *util.RuntimeOptions) (_result *GetApplicationGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetApplicationGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetApplicationGroup"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetApplicationGroup(request *GetApplicationGroupRequest) (_result *GetApplicationGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetApplicationGroupResponse{}
	_body, _err := client.GetApplicationGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetExecutionTemplateWithOptions(request *GetExecutionTemplateRequest, runtime *util.RuntimeOptions) (_result *GetExecutionTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetExecutionTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetExecutionTemplate"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetExecutionTemplate(request *GetExecutionTemplateRequest) (_result *GetExecutionTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetExecutionTemplateResponse{}
	_body, _err := client.GetExecutionTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetInventorySchemaWithOptions(request *GetInventorySchemaRequest, runtime *util.RuntimeOptions) (_result *GetInventorySchemaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetInventorySchemaResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetInventorySchema"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetInventorySchema(request *GetInventorySchemaRequest) (_result *GetInventorySchemaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetInventorySchemaResponse{}
	_body, _err := client.GetInventorySchemaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetParameterWithOptions(request *GetParameterRequest, runtime *util.RuntimeOptions) (_result *GetParameterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetParameterResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetParameter"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetParameter(request *GetParameterRequest) (_result *GetParameterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetParameterResponse{}
	_body, _err := client.GetParameterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetParametersWithOptions(request *GetParametersRequest, runtime *util.RuntimeOptions) (_result *GetParametersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetParametersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetParameters"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetParameters(request *GetParametersRequest) (_result *GetParametersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetParametersResponse{}
	_body, _err := client.GetParametersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetParametersByPathWithOptions(request *GetParametersByPathRequest, runtime *util.RuntimeOptions) (_result *GetParametersByPathResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetParametersByPathResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetParametersByPath"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetParametersByPath(request *GetParametersByPathRequest) (_result *GetParametersByPathResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetParametersByPathResponse{}
	_body, _err := client.GetParametersByPathWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPatchBaselineWithOptions(request *GetPatchBaselineRequest, runtime *util.RuntimeOptions) (_result *GetPatchBaselineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetPatchBaselineResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetPatchBaseline"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPatchBaseline(request *GetPatchBaselineRequest) (_result *GetPatchBaselineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPatchBaselineResponse{}
	_body, _err := client.GetPatchBaselineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSecretParameterWithOptions(request *GetSecretParameterRequest, runtime *util.RuntimeOptions) (_result *GetSecretParameterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetSecretParameterResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetSecretParameter"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSecretParameter(request *GetSecretParameterRequest) (_result *GetSecretParameterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSecretParameterResponse{}
	_body, _err := client.GetSecretParameterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSecretParametersWithOptions(request *GetSecretParametersRequest, runtime *util.RuntimeOptions) (_result *GetSecretParametersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetSecretParametersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetSecretParameters"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSecretParameters(request *GetSecretParametersRequest) (_result *GetSecretParametersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSecretParametersResponse{}
	_body, _err := client.GetSecretParametersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSecretParametersByPathWithOptions(request *GetSecretParametersByPathRequest, runtime *util.RuntimeOptions) (_result *GetSecretParametersByPathResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetSecretParametersByPathResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetSecretParametersByPath"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSecretParametersByPath(request *GetSecretParametersByPathRequest) (_result *GetSecretParametersByPathResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSecretParametersByPathResponse{}
	_body, _err := client.GetSecretParametersByPathWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetServiceSettingsWithOptions(request *GetServiceSettingsRequest, runtime *util.RuntimeOptions) (_result *GetServiceSettingsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetServiceSettingsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetServiceSettings"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetServiceSettings(request *GetServiceSettingsRequest) (_result *GetServiceSettingsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetServiceSettingsResponse{}
	_body, _err := client.GetServiceSettingsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTemplateWithOptions(request *GetTemplateRequest, runtime *util.RuntimeOptions) (_result *GetTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetTemplate"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTemplate(request *GetTemplateRequest) (_result *GetTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTemplateResponse{}
	_body, _err := client.GetTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListActionsWithOptions(request *ListActionsRequest, runtime *util.RuntimeOptions) (_result *ListActionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListActionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListActions"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListActions(request *ListActionsRequest) (_result *ListActionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListActionsResponse{}
	_body, _err := client.ListActionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListApplicationGroupsWithOptions(request *ListApplicationGroupsRequest, runtime *util.RuntimeOptions) (_result *ListApplicationGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListApplicationGroupsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListApplicationGroups"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListApplicationGroups(request *ListApplicationGroupsRequest) (_result *ListApplicationGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListApplicationGroupsResponse{}
	_body, _err := client.ListApplicationGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListApplicationsWithOptions(request *ListApplicationsRequest, runtime *util.RuntimeOptions) (_result *ListApplicationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListApplicationsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListApplications"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListApplications(request *ListApplicationsRequest) (_result *ListApplicationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListApplicationsResponse{}
	_body, _err := client.ListApplicationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListExecutionLogsWithOptions(request *ListExecutionLogsRequest, runtime *util.RuntimeOptions) (_result *ListExecutionLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListExecutionLogsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListExecutionLogs"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListExecutionLogs(request *ListExecutionLogsRequest) (_result *ListExecutionLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListExecutionLogsResponse{}
	_body, _err := client.ListExecutionLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListExecutionRiskyTasksWithOptions(request *ListExecutionRiskyTasksRequest, runtime *util.RuntimeOptions) (_result *ListExecutionRiskyTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListExecutionRiskyTasksResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListExecutionRiskyTasks"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListExecutionRiskyTasks(request *ListExecutionRiskyTasksRequest) (_result *ListExecutionRiskyTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListExecutionRiskyTasksResponse{}
	_body, _err := client.ListExecutionRiskyTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListExecutionsWithOptions(tmpReq *ListExecutionsRequest, runtime *util.RuntimeOptions) (_result *ListExecutionsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListExecutionsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("Tags"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListExecutionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListExecutions"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListExecutions(request *ListExecutionsRequest) (_result *ListExecutionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListExecutionsResponse{}
	_body, _err := client.ListExecutionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListInstancePatchesWithOptions(request *ListInstancePatchesRequest, runtime *util.RuntimeOptions) (_result *ListInstancePatchesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListInstancePatchesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListInstancePatches"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListInstancePatches(request *ListInstancePatchesRequest) (_result *ListInstancePatchesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListInstancePatchesResponse{}
	_body, _err := client.ListInstancePatchesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListInstancePatchStatesWithOptions(request *ListInstancePatchStatesRequest, runtime *util.RuntimeOptions) (_result *ListInstancePatchStatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListInstancePatchStatesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListInstancePatchStates"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListInstancePatchStates(request *ListInstancePatchStatesRequest) (_result *ListInstancePatchStatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListInstancePatchStatesResponse{}
	_body, _err := client.ListInstancePatchStatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListInstanceStateReportsWithOptions(request *ListInstanceStateReportsRequest, runtime *util.RuntimeOptions) (_result *ListInstanceStateReportsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListInstanceStateReportsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListInstanceStateReports"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListInstanceStateReports(request *ListInstanceStateReportsRequest) (_result *ListInstanceStateReportsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListInstanceStateReportsResponse{}
	_body, _err := client.ListInstanceStateReportsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListInventoryEntriesWithOptions(request *ListInventoryEntriesRequest, runtime *util.RuntimeOptions) (_result *ListInventoryEntriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListInventoryEntriesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListInventoryEntries"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListInventoryEntries(request *ListInventoryEntriesRequest) (_result *ListInventoryEntriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListInventoryEntriesResponse{}
	_body, _err := client.ListInventoryEntriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListParametersWithOptions(tmpReq *ListParametersRequest, runtime *util.RuntimeOptions) (_result *ListParametersResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListParametersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("Tags"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListParametersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListParameters"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListParameters(request *ListParametersRequest) (_result *ListParametersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListParametersResponse{}
	_body, _err := client.ListParametersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListParameterVersionsWithOptions(request *ListParameterVersionsRequest, runtime *util.RuntimeOptions) (_result *ListParameterVersionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListParameterVersionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListParameterVersions"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListParameterVersions(request *ListParameterVersionsRequest) (_result *ListParameterVersionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListParameterVersionsResponse{}
	_body, _err := client.ListParameterVersionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPatchBaselinesWithOptions(request *ListPatchBaselinesRequest, runtime *util.RuntimeOptions) (_result *ListPatchBaselinesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListPatchBaselinesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListPatchBaselines"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPatchBaselines(request *ListPatchBaselinesRequest) (_result *ListPatchBaselinesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPatchBaselinesResponse{}
	_body, _err := client.ListPatchBaselinesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListResourceExecutionStatusWithOptions(request *ListResourceExecutionStatusRequest, runtime *util.RuntimeOptions) (_result *ListResourceExecutionStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListResourceExecutionStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListResourceExecutionStatus"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListResourceExecutionStatus(request *ListResourceExecutionStatusRequest) (_result *ListResourceExecutionStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListResourceExecutionStatusResponse{}
	_body, _err := client.ListResourceExecutionStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSecretParametersWithOptions(tmpReq *ListSecretParametersRequest, runtime *util.RuntimeOptions) (_result *ListSecretParametersResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListSecretParametersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("Tags"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListSecretParametersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListSecretParameters"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSecretParameters(request *ListSecretParametersRequest) (_result *ListSecretParametersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSecretParametersResponse{}
	_body, _err := client.ListSecretParametersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSecretParameterVersionsWithOptions(request *ListSecretParameterVersionsRequest, runtime *util.RuntimeOptions) (_result *ListSecretParameterVersionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListSecretParameterVersionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListSecretParameterVersions"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSecretParameterVersions(request *ListSecretParameterVersionsRequest) (_result *ListSecretParameterVersionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSecretParameterVersionsResponse{}
	_body, _err := client.ListSecretParameterVersionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListStateConfigurationsWithOptions(tmpReq *ListStateConfigurationsRequest, runtime *util.RuntimeOptions) (_result *ListStateConfigurationsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListStateConfigurationsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("Tags"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListStateConfigurationsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListStateConfigurations"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListStateConfigurations(request *ListStateConfigurationsRequest) (_result *ListStateConfigurationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListStateConfigurationsResponse{}
	_body, _err := client.ListStateConfigurationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagKeysWithOptions(request *ListTagKeysRequest, runtime *util.RuntimeOptions) (_result *ListTagKeysResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListTagKeysResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTagKeys"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagKeys(request *ListTagKeysRequest) (_result *ListTagKeysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagKeysResponse{}
	_body, _err := client.ListTagKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagResourcesWithOptions(tmpReq *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListTagResourcesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ResourceIds)) {
		request.ResourceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceIds, tea.String("ResourceIds"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("Tags"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTagResources"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagValuesWithOptions(request *ListTagValuesRequest, runtime *util.RuntimeOptions) (_result *ListTagValuesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListTagValuesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTagValues"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagValues(request *ListTagValuesRequest) (_result *ListTagValuesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagValuesResponse{}
	_body, _err := client.ListTagValuesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTaskExecutionsWithOptions(request *ListTaskExecutionsRequest, runtime *util.RuntimeOptions) (_result *ListTaskExecutionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListTaskExecutionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTaskExecutions"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTaskExecutions(request *ListTaskExecutionsRequest) (_result *ListTaskExecutionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTaskExecutionsResponse{}
	_body, _err := client.ListTaskExecutionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTemplatesWithOptions(tmpReq *ListTemplatesRequest, runtime *util.RuntimeOptions) (_result *ListTemplatesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListTemplatesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("Tags"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListTemplatesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTemplates"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTemplates(request *ListTemplatesRequest) (_result *ListTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTemplatesResponse{}
	_body, _err := client.ListTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTemplateVersionsWithOptions(request *ListTemplateVersionsRequest, runtime *util.RuntimeOptions) (_result *ListTemplateVersionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListTemplateVersionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTemplateVersions"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTemplateVersions(request *ListTemplateVersionsRequest) (_result *ListTemplateVersionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTemplateVersionsResponse{}
	_body, _err := client.ListTemplateVersionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) NotifyExecutionWithOptions(request *NotifyExecutionRequest, runtime *util.RuntimeOptions) (_result *NotifyExecutionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &NotifyExecutionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("NotifyExecution"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) NotifyExecution(request *NotifyExecutionRequest) (_result *NotifyExecutionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &NotifyExecutionResponse{}
	_body, _err := client.NotifyExecutionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RegisterDefaultPatchBaselineWithOptions(request *RegisterDefaultPatchBaselineRequest, runtime *util.RuntimeOptions) (_result *RegisterDefaultPatchBaselineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RegisterDefaultPatchBaselineResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RegisterDefaultPatchBaseline"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RegisterDefaultPatchBaseline(request *RegisterDefaultPatchBaselineRequest) (_result *RegisterDefaultPatchBaselineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RegisterDefaultPatchBaselineResponse{}
	_body, _err := client.RegisterDefaultPatchBaselineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchInventoryWithOptions(request *SearchInventoryRequest, runtime *util.RuntimeOptions) (_result *SearchInventoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SearchInventoryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SearchInventory"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchInventory(request *SearchInventoryRequest) (_result *SearchInventoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchInventoryResponse{}
	_body, _err := client.SearchInventoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetServiceSettingsWithOptions(request *SetServiceSettingsRequest, runtime *util.RuntimeOptions) (_result *SetServiceSettingsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetServiceSettingsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetServiceSettings"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetServiceSettings(request *SetServiceSettingsRequest) (_result *SetServiceSettingsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetServiceSettingsResponse{}
	_body, _err := client.SetServiceSettingsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartExecutionWithOptions(tmpReq *StartExecutionRequest, runtime *util.RuntimeOptions) (_result *StartExecutionResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &StartExecutionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("Tags"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StartExecutionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StartExecution"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartExecution(request *StartExecutionRequest) (_result *StartExecutionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartExecutionResponse{}
	_body, _err := client.StartExecutionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TagResourcesWithOptions(tmpReq *TagResourcesRequest, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &TagResourcesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ResourceIds)) {
		request.ResourceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceIds, tea.String("ResourceIds"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("Tags"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("TagResources"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TriggerExecutionWithOptions(request *TriggerExecutionRequest, runtime *util.RuntimeOptions) (_result *TriggerExecutionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &TriggerExecutionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("TriggerExecution"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TriggerExecution(request *TriggerExecutionRequest) (_result *TriggerExecutionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TriggerExecutionResponse{}
	_body, _err := client.TriggerExecutionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UntagResourcesWithOptions(tmpReq *UntagResourcesRequest, runtime *util.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UntagResourcesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ResourceIds)) {
		request.ResourceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceIds, tea.String("ResourceIds"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TagKeys)) {
		request.TagKeysShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TagKeys, tea.String("TagKeys"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UntagResources"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UntagResourcesResponse{}
	_body, _err := client.UntagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateApplicationGroupWithOptions(request *UpdateApplicationGroupRequest, runtime *util.RuntimeOptions) (_result *UpdateApplicationGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateApplicationGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateApplicationGroup"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateApplicationGroup(request *UpdateApplicationGroupRequest) (_result *UpdateApplicationGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateApplicationGroupResponse{}
	_body, _err := client.UpdateApplicationGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateExecutionWithOptions(request *UpdateExecutionRequest, runtime *util.RuntimeOptions) (_result *UpdateExecutionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateExecutionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateExecution"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateExecution(request *UpdateExecutionRequest) (_result *UpdateExecutionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateExecutionResponse{}
	_body, _err := client.UpdateExecutionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateInstanceInformationWithOptions(request *UpdateInstanceInformationRequest, runtime *util.RuntimeOptions) (_result *UpdateInstanceInformationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateInstanceInformationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateInstanceInformation"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateInstanceInformation(request *UpdateInstanceInformationRequest) (_result *UpdateInstanceInformationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateInstanceInformationResponse{}
	_body, _err := client.UpdateInstanceInformationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateParameterWithOptions(request *UpdateParameterRequest, runtime *util.RuntimeOptions) (_result *UpdateParameterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateParameterResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateParameter"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateParameter(request *UpdateParameterRequest) (_result *UpdateParameterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateParameterResponse{}
	_body, _err := client.UpdateParameterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdatePatchBaselineWithOptions(request *UpdatePatchBaselineRequest, runtime *util.RuntimeOptions) (_result *UpdatePatchBaselineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdatePatchBaselineResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdatePatchBaseline"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdatePatchBaseline(request *UpdatePatchBaselineRequest) (_result *UpdatePatchBaselineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdatePatchBaselineResponse{}
	_body, _err := client.UpdatePatchBaselineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateSecretParameterWithOptions(tmpReq *UpdateSecretParameterRequest, runtime *util.RuntimeOptions) (_result *UpdateSecretParameterResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateSecretParameterShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("Tags"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateSecretParameterResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateSecretParameter"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateSecretParameter(request *UpdateSecretParameterRequest) (_result *UpdateSecretParameterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateSecretParameterResponse{}
	_body, _err := client.UpdateSecretParameterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateStateConfigurationWithOptions(tmpReq *UpdateStateConfigurationRequest, runtime *util.RuntimeOptions) (_result *UpdateStateConfigurationResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateStateConfigurationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Parameters)) {
		request.ParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Parameters, tea.String("Parameters"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("Tags"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateStateConfigurationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateStateConfiguration"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateStateConfiguration(request *UpdateStateConfigurationRequest) (_result *UpdateStateConfigurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateStateConfigurationResponse{}
	_body, _err := client.UpdateStateConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateTemplateWithOptions(tmpReq *UpdateTemplateRequest, runtime *util.RuntimeOptions) (_result *UpdateTemplateResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("Tags"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateTemplate"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateTemplate(request *UpdateTemplateRequest) (_result *UpdateTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateTemplateResponse{}
	_body, _err := client.UpdateTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ValidateTemplateContentWithOptions(request *ValidateTemplateContentRequest, runtime *util.RuntimeOptions) (_result *ValidateTemplateContentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ValidateTemplateContentResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ValidateTemplateContent"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ValidateTemplateContent(request *ValidateTemplateContentRequest) (_result *ValidateTemplateContentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ValidateTemplateContentResponse{}
	_body, _err := client.ValidateTemplateContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
