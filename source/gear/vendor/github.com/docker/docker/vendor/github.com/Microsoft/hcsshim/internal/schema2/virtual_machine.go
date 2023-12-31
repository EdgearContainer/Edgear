/*
 * HCS API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package hcsschema

type VirtualMachine struct {
	Chipset *Chipset `json:"Chipset,omitempty"`

	ComputeTopology *Topology `json:"ComputeTopology,omitempty"`

	Devices *Devices `json:"Devices,omitempty"`

	GuestState *GuestState `json:"GuestState,omitempty"`

	RestoreState *RestoreState `json:"RestoreState,omitempty"`

	RegistryChanges *RegistryChanges `json:"RegistryChanges,omitempty"`

	StorageQoS *StorageQoS `json:"StorageQoS,omitempty"`

	GuestConnection *GuestConnection `json:"GuestConnection,omitempty"`
}
