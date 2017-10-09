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

package search

import original "github.com/Azure/azure-sdk-for-go/services/search/mgmt/2015-08-19/search"

type AdminKeyKind = original.AdminKeyKind

const (
	Primary		AdminKeyKind	= original.Primary
	Secondary	AdminKeyKind	= original.Secondary
)

type HostingMode = original.HostingMode

const (
	Default		HostingMode	= original.Default
	HighDensity	HostingMode	= original.HighDensity
)

type ProvisioningState = original.ProvisioningState

const (
	Failed		ProvisioningState	= original.Failed
	Provisioning	ProvisioningState	= original.Provisioning
	Succeeded	ProvisioningState	= original.Succeeded
)

type ServiceStatus = original.ServiceStatus

const (
	ServiceStatusDegraded		ServiceStatus	= original.ServiceStatusDegraded
	ServiceStatusDeleting		ServiceStatus	= original.ServiceStatusDeleting
	ServiceStatusDisabled		ServiceStatus	= original.ServiceStatusDisabled
	ServiceStatusError		ServiceStatus	= original.ServiceStatusError
	ServiceStatusProvisioning	ServiceStatus	= original.ServiceStatusProvisioning
	ServiceStatusRunning		ServiceStatus	= original.ServiceStatusRunning
)

type SkuName = original.SkuName

const (
	Basic		SkuName	= original.Basic
	Free		SkuName	= original.Free
	Standard	SkuName	= original.Standard
	Standard2	SkuName	= original.Standard2
	Standard3	SkuName	= original.Standard3
)

type UnavailableNameReason = original.UnavailableNameReason

const (
	AlreadyExists	UnavailableNameReason	= original.AlreadyExists
	Invalid		UnavailableNameReason	= original.Invalid
)

type AdminKeyResult = original.AdminKeyResult
type CheckNameAvailabilityInput = original.CheckNameAvailabilityInput
type CheckNameAvailabilityOutput = original.CheckNameAvailabilityOutput
type CloudError = original.CloudError
type CloudErrorBody = original.CloudErrorBody
type ListQueryKeysResult = original.ListQueryKeysResult
type QueryKey = original.QueryKey
type Resource = original.Resource
type Service = original.Service
type ServiceListResult = original.ServiceListResult
type ServiceProperties = original.ServiceProperties
type Sku = original.Sku
type QueryKeysClient = original.QueryKeysClient
type ServicesClient = original.ServicesClient
type AdminKeysClient = original.AdminKeysClient

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ManagementClient = original.ManagementClient

func NewAdminKeysClient(subscriptionID string) AdminKeysClient {
	return original.NewAdminKeysClient(subscriptionID)
}
func NewAdminKeysClientWithBaseURI(baseURI string, subscriptionID string) AdminKeysClient {
	return original.NewAdminKeysClientWithBaseURI(baseURI, subscriptionID)
}
func New(subscriptionID string) ManagementClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) ManagementClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func NewQueryKeysClient(subscriptionID string) QueryKeysClient {
	return original.NewQueryKeysClient(subscriptionID)
}
func NewQueryKeysClientWithBaseURI(baseURI string, subscriptionID string) QueryKeysClient {
	return original.NewQueryKeysClientWithBaseURI(baseURI, subscriptionID)
}
func NewServicesClient(subscriptionID string) ServicesClient {
	return original.NewServicesClient(subscriptionID)
}
func NewServicesClientWithBaseURI(baseURI string, subscriptionID string) ServicesClient {
	return original.NewServicesClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent()
}
func Version() string {
	return original.Version()
}
