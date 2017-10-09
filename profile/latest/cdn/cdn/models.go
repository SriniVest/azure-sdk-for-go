// +build go1.9

// Copyright 2017 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder
// commit ID: 0557aed5f03338a57db6ea921d282d5a85622abf

package cdn

import original "github.com/Azure/azure-sdk-for-go/services/cdn/mgmt/2016-10-02/cdn"

type EdgeNodesClient = original.EdgeNodesClient
type EndpointsClient = original.EndpointsClient
type CustomDomainResourceState = original.CustomDomainResourceState

const (
	Active		CustomDomainResourceState	= original.Active
	Creating	CustomDomainResourceState	= original.Creating
	Deleting	CustomDomainResourceState	= original.Deleting
)

type CustomHTTPSProvisioningState = original.CustomHTTPSProvisioningState

const (
	Disabled	CustomHTTPSProvisioningState	= original.Disabled
	Disabling	CustomHTTPSProvisioningState	= original.Disabling
	Enabled		CustomHTTPSProvisioningState	= original.Enabled
	Enabling	CustomHTTPSProvisioningState	= original.Enabling
	Failed		CustomHTTPSProvisioningState	= original.Failed
)

type EndpointResourceState = original.EndpointResourceState

const (
	EndpointResourceStateCreating	EndpointResourceState	= original.EndpointResourceStateCreating
	EndpointResourceStateDeleting	EndpointResourceState	= original.EndpointResourceStateDeleting
	EndpointResourceStateRunning	EndpointResourceState	= original.EndpointResourceStateRunning
	EndpointResourceStateStarting	EndpointResourceState	= original.EndpointResourceStateStarting
	EndpointResourceStateStopped	EndpointResourceState	= original.EndpointResourceStateStopped
	EndpointResourceStateStopping	EndpointResourceState	= original.EndpointResourceStateStopping
)

type GeoFilterActions = original.GeoFilterActions

const (
	Allow	GeoFilterActions	= original.Allow
	Block	GeoFilterActions	= original.Block
)

type OptimizationType = original.OptimizationType

const (
	DynamicSiteAcceleration		OptimizationType	= original.DynamicSiteAcceleration
	GeneralMediaStreaming		OptimizationType	= original.GeneralMediaStreaming
	GeneralWebDelivery		OptimizationType	= original.GeneralWebDelivery
	LargeFileDownload		OptimizationType	= original.LargeFileDownload
	VideoOnDemandMediaStreaming	OptimizationType	= original.VideoOnDemandMediaStreaming
)

type OriginResourceState = original.OriginResourceState

const (
	OriginResourceStateActive	OriginResourceState	= original.OriginResourceStateActive
	OriginResourceStateCreating	OriginResourceState	= original.OriginResourceStateCreating
	OriginResourceStateDeleting	OriginResourceState	= original.OriginResourceStateDeleting
)

type ProfileResourceState = original.ProfileResourceState

const (
	ProfileResourceStateActive	ProfileResourceState	= original.ProfileResourceStateActive
	ProfileResourceStateCreating	ProfileResourceState	= original.ProfileResourceStateCreating
	ProfileResourceStateDeleting	ProfileResourceState	= original.ProfileResourceStateDeleting
	ProfileResourceStateDisabled	ProfileResourceState	= original.ProfileResourceStateDisabled
)

type QueryStringCachingBehavior = original.QueryStringCachingBehavior

const (
	BypassCaching		QueryStringCachingBehavior	= original.BypassCaching
	IgnoreQueryString	QueryStringCachingBehavior	= original.IgnoreQueryString
	NotSet			QueryStringCachingBehavior	= original.NotSet
	UseQueryString		QueryStringCachingBehavior	= original.UseQueryString
)

type ResourceType = original.ResourceType

const (
	MicrosoftCdnProfilesEndpoints ResourceType = original.MicrosoftCdnProfilesEndpoints
)

type SkuName = original.SkuName

const (
	CustomVerizon		SkuName	= original.CustomVerizon
	PremiumVerizon		SkuName	= original.PremiumVerizon
	StandardAkamai		SkuName	= original.StandardAkamai
	StandardChinaCdn	SkuName	= original.StandardChinaCdn
	StandardVerizon		SkuName	= original.StandardVerizon
)

type CheckNameAvailabilityInput = original.CheckNameAvailabilityInput
type CheckNameAvailabilityOutput = original.CheckNameAvailabilityOutput
type CidrIPAddress = original.CidrIPAddress
type CustomDomain = original.CustomDomain
type CustomDomainListResult = original.CustomDomainListResult
type CustomDomainParameters = original.CustomDomainParameters
type CustomDomainProperties = original.CustomDomainProperties
type CustomDomainPropertiesParameters = original.CustomDomainPropertiesParameters
type DeepCreatedOrigin = original.DeepCreatedOrigin
type DeepCreatedOriginProperties = original.DeepCreatedOriginProperties
type EdgeNode = original.EdgeNode
type EdgeNodeProperties = original.EdgeNodeProperties
type EdgenodeResult = original.EdgenodeResult
type Endpoint = original.Endpoint
type EndpointListResult = original.EndpointListResult
type EndpointProperties = original.EndpointProperties
type EndpointPropertiesUpdateParameters = original.EndpointPropertiesUpdateParameters
type EndpointUpdateParameters = original.EndpointUpdateParameters
type ErrorResponse = original.ErrorResponse
type GeoFilter = original.GeoFilter
type IPAddressGroup = original.IPAddressGroup
type LoadParameters = original.LoadParameters
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type Origin = original.Origin
type OriginListResult = original.OriginListResult
type OriginProperties = original.OriginProperties
type OriginPropertiesParameters = original.OriginPropertiesParameters
type OriginUpdateParameters = original.OriginUpdateParameters
type Profile = original.Profile
type ProfileListResult = original.ProfileListResult
type ProfileProperties = original.ProfileProperties
type ProfileUpdateParameters = original.ProfileUpdateParameters
type PurgeParameters = original.PurgeParameters
type Resource = original.Resource
type ResourceUsage = original.ResourceUsage
type ResourceUsageListResult = original.ResourceUsageListResult
type Sku = original.Sku
type SsoURI = original.SsoURI
type SupportedOptimizationTypesResult = original.SupportedOptimizationTypesResult
type ValidateCustomDomainInput = original.ValidateCustomDomainInput
type ValidateCustomDomainOutput = original.ValidateCustomDomainOutput
type OriginsClient = original.OriginsClient
type ProfilesClient = original.ProfilesClient

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ManagementClient = original.ManagementClient
type CustomDomainsClient = original.CustomDomainsClient

func NewProfilesClient(subscriptionID string) ProfilesClient {
	return original.NewProfilesClient(subscriptionID)
}
func NewProfilesClientWithBaseURI(baseURI string, subscriptionID string) ProfilesClient {
	return original.NewProfilesClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent()
}
func Version() string {
	return original.Version()
}
func New(subscriptionID string) ManagementClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) ManagementClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func NewCustomDomainsClient(subscriptionID string) CustomDomainsClient {
	return original.NewCustomDomainsClient(subscriptionID)
}
func NewCustomDomainsClientWithBaseURI(baseURI string, subscriptionID string) CustomDomainsClient {
	return original.NewCustomDomainsClientWithBaseURI(baseURI, subscriptionID)
}
func NewEdgeNodesClient(subscriptionID string) EdgeNodesClient {
	return original.NewEdgeNodesClient(subscriptionID)
}
func NewEdgeNodesClientWithBaseURI(baseURI string, subscriptionID string) EdgeNodesClient {
	return original.NewEdgeNodesClientWithBaseURI(baseURI, subscriptionID)
}
func NewEndpointsClient(subscriptionID string) EndpointsClient {
	return original.NewEndpointsClient(subscriptionID)
}
func NewEndpointsClientWithBaseURI(baseURI string, subscriptionID string) EndpointsClient {
	return original.NewEndpointsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOriginsClient(subscriptionID string) OriginsClient {
	return original.NewOriginsClient(subscriptionID)
}
func NewOriginsClientWithBaseURI(baseURI string, subscriptionID string) OriginsClient {
	return original.NewOriginsClientWithBaseURI(baseURI, subscriptionID)
}
