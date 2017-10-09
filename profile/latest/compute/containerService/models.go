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

package containerservice

import original "github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2017-01-31/containerservice"

type ContainerServicesClient = original.ContainerServicesClient
type OrchestratorTypes = original.OrchestratorTypes

const (
	Custom		OrchestratorTypes	= original.Custom
	DCOS		OrchestratorTypes	= original.DCOS
	Kubernetes	OrchestratorTypes	= original.Kubernetes
	Swarm		OrchestratorTypes	= original.Swarm
)

type VMSizeTypes = original.VMSizeTypes

const (
	StandardA0	VMSizeTypes	= original.StandardA0
	StandardA1	VMSizeTypes	= original.StandardA1
	StandardA10	VMSizeTypes	= original.StandardA10
	StandardA11	VMSizeTypes	= original.StandardA11
	StandardA2	VMSizeTypes	= original.StandardA2
	StandardA3	VMSizeTypes	= original.StandardA3
	StandardA4	VMSizeTypes	= original.StandardA4
	StandardA5	VMSizeTypes	= original.StandardA5
	StandardA6	VMSizeTypes	= original.StandardA6
	StandardA7	VMSizeTypes	= original.StandardA7
	StandardA8	VMSizeTypes	= original.StandardA8
	StandardA9	VMSizeTypes	= original.StandardA9
	StandardD1	VMSizeTypes	= original.StandardD1
	StandardD11	VMSizeTypes	= original.StandardD11
	StandardD11V2	VMSizeTypes	= original.StandardD11V2
	StandardD12	VMSizeTypes	= original.StandardD12
	StandardD12V2	VMSizeTypes	= original.StandardD12V2
	StandardD13	VMSizeTypes	= original.StandardD13
	StandardD13V2	VMSizeTypes	= original.StandardD13V2
	StandardD14	VMSizeTypes	= original.StandardD14
	StandardD14V2	VMSizeTypes	= original.StandardD14V2
	StandardD1V2	VMSizeTypes	= original.StandardD1V2
	StandardD2	VMSizeTypes	= original.StandardD2
	StandardD2V2	VMSizeTypes	= original.StandardD2V2
	StandardD3	VMSizeTypes	= original.StandardD3
	StandardD3V2	VMSizeTypes	= original.StandardD3V2
	StandardD4	VMSizeTypes	= original.StandardD4
	StandardD4V2	VMSizeTypes	= original.StandardD4V2
	StandardD5V2	VMSizeTypes	= original.StandardD5V2
	StandardDS1	VMSizeTypes	= original.StandardDS1
	StandardDS11	VMSizeTypes	= original.StandardDS11
	StandardDS12	VMSizeTypes	= original.StandardDS12
	StandardDS13	VMSizeTypes	= original.StandardDS13
	StandardDS14	VMSizeTypes	= original.StandardDS14
	StandardDS2	VMSizeTypes	= original.StandardDS2
	StandardDS3	VMSizeTypes	= original.StandardDS3
	StandardDS4	VMSizeTypes	= original.StandardDS4
	StandardG1	VMSizeTypes	= original.StandardG1
	StandardG2	VMSizeTypes	= original.StandardG2
	StandardG3	VMSizeTypes	= original.StandardG3
	StandardG4	VMSizeTypes	= original.StandardG4
	StandardG5	VMSizeTypes	= original.StandardG5
	StandardGS1	VMSizeTypes	= original.StandardGS1
	StandardGS2	VMSizeTypes	= original.StandardGS2
	StandardGS3	VMSizeTypes	= original.StandardGS3
	StandardGS4	VMSizeTypes	= original.StandardGS4
	StandardGS5	VMSizeTypes	= original.StandardGS5
)

type AgentPoolProfile = original.AgentPoolProfile
type ContainerService = original.ContainerService
type CustomProfile = original.CustomProfile
type DiagnosticsProfile = original.DiagnosticsProfile
type LinuxProfile = original.LinuxProfile
type ListResult = original.ListResult
type MasterProfile = original.MasterProfile
type OrchestratorProfile = original.OrchestratorProfile
type Properties = original.Properties
type Resource = original.Resource
type ServicePrincipalProfile = original.ServicePrincipalProfile
type SSHConfiguration = original.SSHConfiguration
type SSHPublicKey = original.SSHPublicKey
type VMDiagnostics = original.VMDiagnostics
type WindowsProfile = original.WindowsProfile

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ManagementClient = original.ManagementClient

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
func NewContainerServicesClient(subscriptionID string) ContainerServicesClient {
	return original.NewContainerServicesClient(subscriptionID)
}
func NewContainerServicesClientWithBaseURI(baseURI string, subscriptionID string) ContainerServicesClient {
	return original.NewContainerServicesClientWithBaseURI(baseURI, subscriptionID)
}
