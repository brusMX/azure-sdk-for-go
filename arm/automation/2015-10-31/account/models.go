package account

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator 1.0.1.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
    "github.com/Azure/go-autorest/autorest"
    "github.com/Azure/go-autorest/autorest/date"
    "github.com/Azure/go-autorest/autorest/to"
    "net/http"
)

// AutomationAccountState enumerates the values for automation account state.
type AutomationAccountState string

const (
    // Ok specifies the ok state for automation account state.
    Ok AutomationAccountState = "Ok"
    // Suspended specifies the suspended state for automation account state.
    Suspended AutomationAccountState = "Suspended"
    // Unavailable specifies the unavailable state for automation account
    // state.
    Unavailable AutomationAccountState = "Unavailable"
)

// ContentSourceType enumerates the values for content source type.
type ContentSourceType string

const (
    // EmbeddedContent specifies the embedded content state for content source
    // type.
    EmbeddedContent ContentSourceType = "embeddedContent"
    // URI specifies the uri state for content source type.
    URI ContentSourceType = "uri"
)

// DscConfigurationProvisioningState enumerates the values for dsc
// configuration provisioning state.
type DscConfigurationProvisioningState string

const (
    // Succeeded specifies the succeeded state for dsc configuration
    // provisioning state.
    Succeeded DscConfigurationProvisioningState = "Succeeded"
)

// DscConfigurationState enumerates the values for dsc configuration state.
type DscConfigurationState string

const (
    // DscConfigurationStateEdit specifies the dsc configuration state edit
    // state for dsc configuration state.
    DscConfigurationStateEdit DscConfigurationState = "Edit"
    // DscConfigurationStateNew specifies the dsc configuration state new state
    // for dsc configuration state.
    DscConfigurationStateNew DscConfigurationState = "New"
    // DscConfigurationStatePublished specifies the dsc configuration state
    // published state for dsc configuration state.
    DscConfigurationStatePublished DscConfigurationState = "Published"
)

// ModuleProvisioningState enumerates the values for module provisioning state.
type ModuleProvisioningState string

const (
    // ModuleProvisioningStateActivitiesStored specifies the module
    // provisioning state activities stored state for module provisioning
    // state.
    ModuleProvisioningStateActivitiesStored ModuleProvisioningState = "ActivitiesStored"
    // ModuleProvisioningStateCancelled specifies the module provisioning state
    // cancelled state for module provisioning state.
    ModuleProvisioningStateCancelled ModuleProvisioningState = "Cancelled"
    // ModuleProvisioningStateConnectionTypeImported specifies the module
    // provisioning state connection type imported state for module
    // provisioning state.
    ModuleProvisioningStateConnectionTypeImported ModuleProvisioningState = "ConnectionTypeImported"
    // ModuleProvisioningStateContentDownloaded specifies the module
    // provisioning state content downloaded state for module provisioning
    // state.
    ModuleProvisioningStateContentDownloaded ModuleProvisioningState = "ContentDownloaded"
    // ModuleProvisioningStateContentRetrieved specifies the module
    // provisioning state content retrieved state for module provisioning
    // state.
    ModuleProvisioningStateContentRetrieved ModuleProvisioningState = "ContentRetrieved"
    // ModuleProvisioningStateContentStored specifies the module provisioning
    // state content stored state for module provisioning state.
    ModuleProvisioningStateContentStored ModuleProvisioningState = "ContentStored"
    // ModuleProvisioningStateContentValidated specifies the module
    // provisioning state content validated state for module provisioning
    // state.
    ModuleProvisioningStateContentValidated ModuleProvisioningState = "ContentValidated"
    // ModuleProvisioningStateCreated specifies the module provisioning state
    // created state for module provisioning state.
    ModuleProvisioningStateCreated ModuleProvisioningState = "Created"
    // ModuleProvisioningStateCreating specifies the module provisioning state
    // creating state for module provisioning state.
    ModuleProvisioningStateCreating ModuleProvisioningState = "Creating"
    // ModuleProvisioningStateFailed specifies the module provisioning state
    // failed state for module provisioning state.
    ModuleProvisioningStateFailed ModuleProvisioningState = "Failed"
    // ModuleProvisioningStateModuleDataStored specifies the module
    // provisioning state module data stored state for module provisioning
    // state.
    ModuleProvisioningStateModuleDataStored ModuleProvisioningState = "ModuleDataStored"
    // ModuleProvisioningStateModuleImportRunbookComplete specifies the module
    // provisioning state module import runbook complete state for module
    // provisioning state.
    ModuleProvisioningStateModuleImportRunbookComplete ModuleProvisioningState = "ModuleImportRunbookComplete"
    // ModuleProvisioningStateRunningImportModuleRunbook specifies the module
    // provisioning state running import module runbook state for module
    // provisioning state.
    ModuleProvisioningStateRunningImportModuleRunbook ModuleProvisioningState = "RunningImportModuleRunbook"
    // ModuleProvisioningStateStartingImportModuleRunbook specifies the module
    // provisioning state starting import module runbook state for module
    // provisioning state.
    ModuleProvisioningStateStartingImportModuleRunbook ModuleProvisioningState = "StartingImportModuleRunbook"
    // ModuleProvisioningStateSucceeded specifies the module provisioning state
    // succeeded state for module provisioning state.
    ModuleProvisioningStateSucceeded ModuleProvisioningState = "Succeeded"
    // ModuleProvisioningStateUpdating specifies the module provisioning state
    // updating state for module provisioning state.
    ModuleProvisioningStateUpdating ModuleProvisioningState = "Updating"
)

// RunbookProvisioningState enumerates the values for runbook provisioning
// state.
type RunbookProvisioningState string

const (
    // RunbookProvisioningStateSucceeded specifies the runbook provisioning
    // state succeeded state for runbook provisioning state.
    RunbookProvisioningStateSucceeded RunbookProvisioningState = "Succeeded"
)

// RunbookState enumerates the values for runbook state.
type RunbookState string

const (
    // RunbookStateEdit specifies the runbook state edit state for runbook
    // state.
    RunbookStateEdit RunbookState = "Edit"
    // RunbookStateNew specifies the runbook state new state for runbook state.
    RunbookStateNew RunbookState = "New"
    // RunbookStatePublished specifies the runbook state published state for
    // runbook state.
    RunbookStatePublished RunbookState = "Published"
)

// RunbookTypeEnum enumerates the values for runbook type enum.
type RunbookTypeEnum string

const (
    // Graph specifies the graph state for runbook type enum.
    Graph RunbookTypeEnum = "Graph"
    // GraphPowerShell specifies the graph power shell state for runbook type
    // enum.
    GraphPowerShell RunbookTypeEnum = "GraphPowerShell"
    // GraphPowerShellWorkflow specifies the graph power shell workflow state
    // for runbook type enum.
    GraphPowerShellWorkflow RunbookTypeEnum = "GraphPowerShellWorkflow"
    // PowerShell specifies the power shell state for runbook type enum.
    PowerShell RunbookTypeEnum = "PowerShell"
    // PowerShellWorkflow specifies the power shell workflow state for runbook
    // type enum.
    PowerShellWorkflow RunbookTypeEnum = "PowerShellWorkflow"
    // Script specifies the script state for runbook type enum.
    Script RunbookTypeEnum = "Script"
)

// SkuNameEnum enumerates the values for sku name enum.
type SkuNameEnum string

const (
    // Basic specifies the basic state for sku name enum.
    Basic SkuNameEnum = "Basic"
    // Free specifies the free state for sku name enum.
    Free SkuNameEnum = "Free"
)

// AutomationAccount is definition of the automation account type.
type AutomationAccount struct {
    autorest.Response `json:"-"`
    ID *string `json:"id,omitempty"`
    Name *string `json:"name,omitempty"`
    Type *string `json:"type,omitempty"`
    Location *string `json:"location,omitempty"`
    Tags *map[string]*string `json:"tags,omitempty"`
    *AutomationAccountProperties `json:"properties,omitempty"`
    Etag *string `json:"etag,omitempty"`
}

// AutomationAccountCreateOrUpdateParameters is the parameters supplied to the
// create or update automation account operation.
type AutomationAccountCreateOrUpdateParameters struct {
    *AutomationAccountCreateOrUpdateProperties `json:"properties,omitempty"`
    Name *string `json:"name,omitempty"`
    Location *string `json:"location,omitempty"`
    Tags *map[string]*string `json:"tags,omitempty"`
}

// AutomationAccountCreateOrUpdateProperties is the parameters supplied to the
// create or update account properties.
type AutomationAccountCreateOrUpdateProperties struct {
    Sku *Sku `json:"sku,omitempty"`
}

// AutomationAccountListResult is the response model for the list account
// operation.
type AutomationAccountListResult struct {
    autorest.Response `json:"-"`
    Value *[]AutomationAccount `json:"value,omitempty"`
    NextLink *string `json:"nextLink,omitempty"`
}

// AutomationAccountListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client AutomationAccountListResult) AutomationAccountListResultPreparer() (*http.Request, error) {
    if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
        return nil, nil
    }
    return autorest.Prepare(&http.Request{},
        autorest.AsJSON(),
        autorest.AsGet(),
        autorest.WithBaseURL(to.String(client.NextLink)));
}

// AutomationAccountProperties is definition of the account property.
type AutomationAccountProperties struct {
    Sku *Sku `json:"sku,omitempty"`
    LastModifiedBy *string `json:"lastModifiedBy,omitempty"`
    State AutomationAccountState `json:"state,omitempty"`
    CreationTime *date.Time `json:"creationTime,omitempty"`
    LastModifiedTime *date.Time `json:"lastModifiedTime,omitempty"`
    Description *string `json:"description,omitempty"`
}

// AutomationAccountUpdateParameters is the parameters supplied to the update
// automation account operation.
type AutomationAccountUpdateParameters struct {
    *AutomationAccountUpdateProperties `json:"properties,omitempty"`
    Name *string `json:"name,omitempty"`
    Location *string `json:"location,omitempty"`
    Tags *map[string]*string `json:"tags,omitempty"`
}

// AutomationAccountUpdateProperties is the parameters supplied to the update
// account properties.
type AutomationAccountUpdateProperties struct {
    Sku *Sku `json:"sku,omitempty"`
}

// ContentHash is definition of the runbook property type.
type ContentHash struct {
    Algorithm *string `json:"algorithm,omitempty"`
    Value *string `json:"value,omitempty"`
}

// ContentLink is definition of the content link.
type ContentLink struct {
    URI *string `json:"uri,omitempty"`
    ContentHash *ContentHash `json:"contentHash,omitempty"`
    Version *string `json:"version,omitempty"`
}

// ContentSource is definition of the content source.
type ContentSource struct {
    Hash *ContentHash `json:"hash,omitempty"`
    Type ContentSourceType `json:"type,omitempty"`
    Value *string `json:"value,omitempty"`
    Version *string `json:"version,omitempty"`
}

// DscConfiguration is definition of the configuration type.
type DscConfiguration struct {
    ID *string `json:"id,omitempty"`
    Name *string `json:"name,omitempty"`
    Type *string `json:"type,omitempty"`
    Location *string `json:"location,omitempty"`
    Tags *map[string]*string `json:"tags,omitempty"`
    *DscConfigurationProperties `json:"properties,omitempty"`
    Etag *string `json:"etag,omitempty"`
}

// DscConfigurationParameter is definition of the configuration parameter type.
type DscConfigurationParameter struct {
    Type *string `json:"type,omitempty"`
    IsMandatory *bool `json:"isMandatory,omitempty"`
    Position *int32 `json:"position,omitempty"`
    DefaultValue *string `json:"defaultValue,omitempty"`
}

// DscConfigurationProperties is definition of the configuration property type.
type DscConfigurationProperties struct {
    ProvisioningState DscConfigurationProvisioningState `json:"provisioningState,omitempty"`
    JobCount *int32 `json:"jobCount,omitempty"`
    Parameters *map[string]*DscConfigurationParameter `json:"parameters,omitempty"`
    Source *ContentSource `json:"source,omitempty"`
    State DscConfigurationState `json:"state,omitempty"`
    LogVerbose *bool `json:"logVerbose,omitempty"`
    CreationTime *date.Time `json:"creationTime,omitempty"`
    LastModifiedTime *date.Time `json:"lastModifiedTime,omitempty"`
    Description *string `json:"description,omitempty"`
}

// DscNode is definition of the dsc node type.
type DscNode struct {
    ID *string `json:"id,omitempty"`
    Name *string `json:"name,omitempty"`
    Type *string `json:"type,omitempty"`
    Location *string `json:"location,omitempty"`
    Tags *map[string]*string `json:"tags,omitempty"`
    LastSeen *date.Time `json:"lastSeen,omitempty"`
    RegistrationTime *date.Time `json:"registrationTime,omitempty"`
    IP *string `json:"ip,omitempty"`
    AccountID *string `json:"accountId,omitempty"`
    NodeConfiguration *DscNodeConfigurationAssociationProperty `json:"nodeConfiguration,omitempty"`
    Status *string `json:"status,omitempty"`
    NodeID *string `json:"nodeId,omitempty"`
    Etag *string `json:"etag,omitempty"`
}

// DscNodeConfigurationAssociationProperty is the dsc nodeconfiguration
// property associated with the entity.
type DscNodeConfigurationAssociationProperty struct {
    Name *string `json:"name,omitempty"`
}

// ErrorResponse is error response of an operation failure
type ErrorResponse struct {
    Code *string `json:"code,omitempty"`
    Message *string `json:"message,omitempty"`
}

// Module is definition of the module type.
type Module struct {
    ID *string `json:"id,omitempty"`
    Name *string `json:"name,omitempty"`
    Type *string `json:"type,omitempty"`
    Location *string `json:"location,omitempty"`
    Tags *map[string]*string `json:"tags,omitempty"`
    *ModuleProperties `json:"properties,omitempty"`
    Etag *string `json:"etag,omitempty"`
}

// ModuleErrorInfo is definition of the module error info type.
type ModuleErrorInfo struct {
    Code *string `json:"code,omitempty"`
    Message *string `json:"message,omitempty"`
}

// ModuleProperties is definition of the module property type.
type ModuleProperties struct {
    IsGlobal *bool `json:"isGlobal,omitempty"`
    Version *string `json:"version,omitempty"`
    SizeInBytes *int64 `json:"sizeInBytes,omitempty"`
    ActivityCount *int32 `json:"activityCount,omitempty"`
    ProvisioningState ModuleProvisioningState `json:"provisioningState,omitempty"`
    ContentLink *ContentLink `json:"contentLink,omitempty"`
    Error *ModuleErrorInfo `json:"error,omitempty"`
    CreationTime *date.Time `json:"creationTime,omitempty"`
    LastModifiedTime *date.Time `json:"lastModifiedTime,omitempty"`
    Description *string `json:"description,omitempty"`
}

// Operation is automation REST API operation
type Operation struct {
    Name *string `json:"name,omitempty"`
    Display *OperationDisplay `json:"display,omitempty"`
}

// OperationDisplay is provider, Resource and Operation values
type OperationDisplay struct {
    Provider *string `json:"provider,omitempty"`
    Resource *string `json:"resource,omitempty"`
    Operation *string `json:"operation,omitempty"`
}

// OperationListResult is the response model for the list of Automation
// operations
type OperationListResult struct {
    autorest.Response `json:"-"`
    Value *[]Operation `json:"value,omitempty"`
}

// Resource is the Resource definition.
type Resource struct {
    ID *string `json:"id,omitempty"`
    Name *string `json:"name,omitempty"`
    Type *string `json:"type,omitempty"`
    Location *string `json:"location,omitempty"`
    Tags *map[string]*string `json:"tags,omitempty"`
}

// Runbook is definition of the runbook type.
type Runbook struct {
    ID *string `json:"id,omitempty"`
    Name *string `json:"name,omitempty"`
    Type *string `json:"type,omitempty"`
    Location *string `json:"location,omitempty"`
    Tags *map[string]*string `json:"tags,omitempty"`
    *RunbookProperties `json:"properties,omitempty"`
    Etag *string `json:"etag,omitempty"`
}

// RunbookDraft is definition of the runbook type.
type RunbookDraft struct {
    InEdit *bool `json:"inEdit,omitempty"`
    DraftContentLink *ContentLink `json:"draftContentLink,omitempty"`
    CreationTime *date.Time `json:"creationTime,omitempty"`
    LastModifiedTime *date.Time `json:"lastModifiedTime,omitempty"`
    Parameters *map[string]*RunbookParameter `json:"parameters,omitempty"`
    OutputTypes *[]string `json:"outputTypes,omitempty"`
}

// RunbookParameter is definition of the runbook parameter type.
type RunbookParameter struct {
    Type *string `json:"type,omitempty"`
    IsMandatory *bool `json:"isMandatory,omitempty"`
    Position *int32 `json:"position,omitempty"`
    DefaultValue *string `json:"defaultValue,omitempty"`
}

// RunbookProperties is definition of the runbook property type.
type RunbookProperties struct {
    RunbookType RunbookTypeEnum `json:"runbookType,omitempty"`
    PublishContentLink *ContentLink `json:"publishContentLink,omitempty"`
    State RunbookState `json:"state,omitempty"`
    LogVerbose *bool `json:"logVerbose,omitempty"`
    LogProgress *bool `json:"logProgress,omitempty"`
    LogActivityTrace *int32 `json:"logActivityTrace,omitempty"`
    JobCount *int32 `json:"jobCount,omitempty"`
    Parameters *map[string]*RunbookParameter `json:"parameters,omitempty"`
    OutputTypes *[]string `json:"outputTypes,omitempty"`
    Draft *RunbookDraft `json:"draft,omitempty"`
    ProvisioningState RunbookProvisioningState `json:"provisioningState,omitempty"`
    LastModifiedBy *string `json:"lastModifiedBy,omitempty"`
    CreationTime *date.Time `json:"creationTime,omitempty"`
    LastModifiedTime *date.Time `json:"lastModifiedTime,omitempty"`
    Description *string `json:"description,omitempty"`
}

// Sku is the account SKU.
type Sku struct {
    Name SkuNameEnum `json:"name,omitempty"`
    Family *string `json:"family,omitempty"`
    Capacity *int32 `json:"capacity,omitempty"`
}

// Statistics is definition of the statistic.
type Statistics struct {
    CounterProperty *string `json:"counterProperty,omitempty"`
    CounterValue *int64 `json:"counterValue,omitempty"`
    StartTime *date.Time `json:"startTime,omitempty"`
    EndTime *date.Time `json:"endTime,omitempty"`
    ID *string `json:"id,omitempty"`
}

// StatisticsListResult is the response model for the list statistics
// operation.
type StatisticsListResult struct {
    autorest.Response `json:"-"`
    Value *[]Statistics `json:"value,omitempty"`
}

// Usage is definition of Usage.
type Usage struct {
    ID *string `json:"id,omitempty"`
    Name *UsageCounterName `json:"name,omitempty"`
    Unit *string `json:"unit,omitempty"`
    CurrentValue *float64 `json:"currentValue,omitempty"`
    Limit *int64 `json:"limit,omitempty"`
    ThrottleStatus *string `json:"throttleStatus,omitempty"`
}

// UsageCounterName is definition of usage counter name.
type UsageCounterName struct {
    Value *string `json:"value,omitempty"`
    LocalizedValue *string `json:"localizedValue,omitempty"`
}

// UsageListResult is the response model for the get usage operation.
type UsageListResult struct {
    autorest.Response `json:"-"`
    Value *[]Usage `json:"value,omitempty"`
}

