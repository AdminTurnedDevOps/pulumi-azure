// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package compute

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/linux_virtual_machine_scale_set.html.markdown.
type LinuxVirtualMachineScaleSet struct {
	pulumi.CustomResourceState

	// A `additionalCapabilities` block as defined below.
	AdditionalCapabilities LinuxVirtualMachineScaleSetAdditionalCapabilitiesPtrOutput `pulumi:"additionalCapabilities"`
	// The Password which should be used for the local-administrator on this Virtual Machine. Changing this forces a new resource to be created.
	AdminPassword pulumi.StringPtrOutput `pulumi:"adminPassword"`
	// One or more `adminSshKey` blocks as defined below.
	AdminSshKeys LinuxVirtualMachineScaleSetAdminSshKeyArrayOutput `pulumi:"adminSshKeys"`
	// The username of the local administrator on each Virtual Machine Scale Set instance. Changing this forces a new resource to be created.
	AdminUsername pulumi.StringOutput `pulumi:"adminUsername"`
	// A `automaticOsUpgradePolicy` block as defined below. This is Required and can only be specified when `upgradeMode` is set to `Automatic`.
	AutomaticOsUpgradePolicy LinuxVirtualMachineScaleSetAutomaticOsUpgradePolicyPtrOutput `pulumi:"automaticOsUpgradePolicy"`
	// A `bootDiagnostics` block as defined below.
	BootDiagnostics LinuxVirtualMachineScaleSetBootDiagnosticsPtrOutput `pulumi:"bootDiagnostics"`
	// The prefix which should be used for the name of the Virtual Machines in this Scale Set. If unspecified this defaults to the value for the `name` field.
	ComputerNamePrefix pulumi.StringOutput `pulumi:"computerNamePrefix"`
	// The Base64-Encoded Custom Data which should be used for this Virtual Machine Scale Set.
	CustomData pulumi.StringPtrOutput `pulumi:"customData"`
	// One or more `dataDisk` blocks as defined below.
	DataDisks LinuxVirtualMachineScaleSetDataDiskArrayOutput `pulumi:"dataDisks"`
	// Should Password Authentication be disabled on this Virtual Machine Scale Set? Defaults to `true`.
	DisablePasswordAuthentication pulumi.BoolPtrOutput `pulumi:"disablePasswordAuthentication"`
	// Should Virtual Machine Extensions be run on Overprovisioned Virtual Machines in the Scale Set? Defaults to `false`.
	DoNotRunExtensionsOnOverprovisionedMachines pulumi.BoolPtrOutput `pulumi:"doNotRunExtensionsOnOverprovisionedMachines"`
	// The Policy which should be used Virtual Machines are Evicted from the Scale Set. Changing this forces a new resource to be created.
	EvictionPolicy pulumi.StringPtrOutput `pulumi:"evictionPolicy"`
	// The ID of a Load Balancer Probe which should be used to determine the health of an instance. Changing this forces a new resource to be created. This is Required and can only be specified when `upgradeMode` is set to `Automatic` or `Rolling`.
	HealthProbeId pulumi.StringPtrOutput `pulumi:"healthProbeId"`
	// A `identity` block as defined below.
	Identity LinuxVirtualMachineScaleSetIdentityPtrOutput `pulumi:"identity"`
	// The number of Virtual Machines in the Scale Set.
	Instances pulumi.IntOutput `pulumi:"instances"`
	// The Azure location where the Linux Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The maximum price you're willing to pay for each Virtual Machine in this Scale Set, in US Dollars; which must be greater than the current spot price. If this bid price falls below the current spot price the Virtual Machines in the Scale Set will be evicted using the `evictionPolicy`. Defaults to `-1`, which means that each Virtual Machine in this Scale Set should not be evicted for price reasons.
	MaxBidPrice pulumi.Float64PtrOutput `pulumi:"maxBidPrice"`
	// The name of the Linux Virtual Machine Scale Set. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// One or more `networkInterface` blocks as defined below.
	NetworkInterfaces LinuxVirtualMachineScaleSetNetworkInterfaceArrayOutput `pulumi:"networkInterfaces"`
	// An `osDisk` block as defined below.
	OsDisk LinuxVirtualMachineScaleSetOsDiskOutput `pulumi:"osDisk"`
	// Should Azure over-provision Virtual Machines in this Scale Set? This means that multiple Virtual Machines will be provisioned and Azure will keep the instances which become available first - which improves provisioning success rates and improves deployment time. You're not billed for these over-provisioned VM's and they don't count towards the Subscription Quota. Defaults to `false`.
	Overprovision pulumi.BoolPtrOutput `pulumi:"overprovision"`
	Plan LinuxVirtualMachineScaleSetPlanPtrOutput `pulumi:"plan"`
	// The Priority of this Virtual Machine Scale Set. Possible values are `Regular` and `Spot`. Defaults to `Regular`. Changing this value forces a new resource.
	Priority pulumi.StringPtrOutput `pulumi:"priority"`
	// Should the Azure VM Agent be provisioned on each Virtual Machine in the Scale Set? Defaults to `true`. Changing this value forces a new resource to be created.
	ProvisionVmAgent pulumi.BoolPtrOutput `pulumi:"provisionVmAgent"`
	// The ID of the Proximity Placement Group in which the Virtual Machine Scale Set should be assigned to. Changing this forces a new resource to be created.
	ProximityPlacementGroupId pulumi.StringPtrOutput `pulumi:"proximityPlacementGroupId"`
	// The name of the Resource Group in which the Linux Virtual Machine Scale Set should be exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A `rollingUpgradePolicy` block as defined below. This is Required and can only be specified when `upgradeMode` is set to `Automatic` or `Rolling`.
	RollingUpgradePolicy LinuxVirtualMachineScaleSetRollingUpgradePolicyPtrOutput `pulumi:"rollingUpgradePolicy"`
	// One or more `secret` blocks as defined below.
	Secrets LinuxVirtualMachineScaleSetSecretArrayOutput `pulumi:"secrets"`
	// Should this Virtual Machine Scale Set be limited to a Single Placement Group, which means the number of instances will be capped at 100 Virtual Machines. Defaults to `true`.
	SinglePlacementGroup pulumi.BoolPtrOutput `pulumi:"singlePlacementGroup"`
	// The Virtual Machine SKU for the Scale Set, such as `Standard_F2`.
	Sku pulumi.StringOutput `pulumi:"sku"`
	// The ID of an Image which each Virtual Machine in this Scale Set should be based on.
	SourceImageId pulumi.StringPtrOutput `pulumi:"sourceImageId"`
	// A `sourceImageReference` block as defined below.
	SourceImageReference LinuxVirtualMachineScaleSetSourceImageReferencePtrOutput `pulumi:"sourceImageReference"`
	// A mapping of tags which should be assigned to this Virtual Machine Scale Set.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The Unique ID for this Linux Virtual Machine Scale Set.
	UniqueId pulumi.StringOutput `pulumi:"uniqueId"`
	// Specifies how Upgrades (e.g. changing the Image/SKU) should be performed to Virtual Machine Instances. Possible values are `Automatic`, `Manual` and `Rolling`. Defaults to `Manual`.
	UpgradeMode pulumi.StringPtrOutput `pulumi:"upgradeMode"`
	// Should the Virtual Machines in this Scale Set be strictly evenly distributed across Availability Zones? Defaults to `false`. Changing this forces a new resource to be created.
	ZoneBalance pulumi.BoolPtrOutput `pulumi:"zoneBalance"`
	// A list of Availability Zones in which the Virtual Machines in this Scale Set should be created in. Changing this forces a new resource to be created.
	Zones pulumi.StringArrayOutput `pulumi:"zones"`
}

// NewLinuxVirtualMachineScaleSet registers a new resource with the given unique name, arguments, and options.
func NewLinuxVirtualMachineScaleSet(ctx *pulumi.Context,
	name string, args *LinuxVirtualMachineScaleSetArgs, opts ...pulumi.ResourceOption) (*LinuxVirtualMachineScaleSet, error) {
	if args == nil || args.AdminUsername == nil {
		return nil, errors.New("missing required argument 'AdminUsername'")
	}
	if args == nil || args.Instances == nil {
		return nil, errors.New("missing required argument 'Instances'")
	}
	if args == nil || args.NetworkInterfaces == nil {
		return nil, errors.New("missing required argument 'NetworkInterfaces'")
	}
	if args == nil || args.OsDisk == nil {
		return nil, errors.New("missing required argument 'OsDisk'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Sku == nil {
		return nil, errors.New("missing required argument 'Sku'")
	}
	if args == nil {
		args = &LinuxVirtualMachineScaleSetArgs{}
	}
	var resource LinuxVirtualMachineScaleSet
	err := ctx.RegisterResource("azure:compute/linuxVirtualMachineScaleSet:LinuxVirtualMachineScaleSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLinuxVirtualMachineScaleSet gets an existing LinuxVirtualMachineScaleSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLinuxVirtualMachineScaleSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LinuxVirtualMachineScaleSetState, opts ...pulumi.ResourceOption) (*LinuxVirtualMachineScaleSet, error) {
	var resource LinuxVirtualMachineScaleSet
	err := ctx.ReadResource("azure:compute/linuxVirtualMachineScaleSet:LinuxVirtualMachineScaleSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LinuxVirtualMachineScaleSet resources.
type linuxVirtualMachineScaleSetState struct {
	// A `additionalCapabilities` block as defined below.
	AdditionalCapabilities *LinuxVirtualMachineScaleSetAdditionalCapabilities `pulumi:"additionalCapabilities"`
	// The Password which should be used for the local-administrator on this Virtual Machine. Changing this forces a new resource to be created.
	AdminPassword *string `pulumi:"adminPassword"`
	// One or more `adminSshKey` blocks as defined below.
	AdminSshKeys []LinuxVirtualMachineScaleSetAdminSshKey `pulumi:"adminSshKeys"`
	// The username of the local administrator on each Virtual Machine Scale Set instance. Changing this forces a new resource to be created.
	AdminUsername *string `pulumi:"adminUsername"`
	// A `automaticOsUpgradePolicy` block as defined below. This is Required and can only be specified when `upgradeMode` is set to `Automatic`.
	AutomaticOsUpgradePolicy *LinuxVirtualMachineScaleSetAutomaticOsUpgradePolicy `pulumi:"automaticOsUpgradePolicy"`
	// A `bootDiagnostics` block as defined below.
	BootDiagnostics *LinuxVirtualMachineScaleSetBootDiagnostics `pulumi:"bootDiagnostics"`
	// The prefix which should be used for the name of the Virtual Machines in this Scale Set. If unspecified this defaults to the value for the `name` field.
	ComputerNamePrefix *string `pulumi:"computerNamePrefix"`
	// The Base64-Encoded Custom Data which should be used for this Virtual Machine Scale Set.
	CustomData *string `pulumi:"customData"`
	// One or more `dataDisk` blocks as defined below.
	DataDisks []LinuxVirtualMachineScaleSetDataDisk `pulumi:"dataDisks"`
	// Should Password Authentication be disabled on this Virtual Machine Scale Set? Defaults to `true`.
	DisablePasswordAuthentication *bool `pulumi:"disablePasswordAuthentication"`
	// Should Virtual Machine Extensions be run on Overprovisioned Virtual Machines in the Scale Set? Defaults to `false`.
	DoNotRunExtensionsOnOverprovisionedMachines *bool `pulumi:"doNotRunExtensionsOnOverprovisionedMachines"`
	// The Policy which should be used Virtual Machines are Evicted from the Scale Set. Changing this forces a new resource to be created.
	EvictionPolicy *string `pulumi:"evictionPolicy"`
	// The ID of a Load Balancer Probe which should be used to determine the health of an instance. Changing this forces a new resource to be created. This is Required and can only be specified when `upgradeMode` is set to `Automatic` or `Rolling`.
	HealthProbeId *string `pulumi:"healthProbeId"`
	// A `identity` block as defined below.
	Identity *LinuxVirtualMachineScaleSetIdentity `pulumi:"identity"`
	// The number of Virtual Machines in the Scale Set.
	Instances *int `pulumi:"instances"`
	// The Azure location where the Linux Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The maximum price you're willing to pay for each Virtual Machine in this Scale Set, in US Dollars; which must be greater than the current spot price. If this bid price falls below the current spot price the Virtual Machines in the Scale Set will be evicted using the `evictionPolicy`. Defaults to `-1`, which means that each Virtual Machine in this Scale Set should not be evicted for price reasons.
	MaxBidPrice *float64 `pulumi:"maxBidPrice"`
	// The name of the Linux Virtual Machine Scale Set. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// One or more `networkInterface` blocks as defined below.
	NetworkInterfaces []LinuxVirtualMachineScaleSetNetworkInterface `pulumi:"networkInterfaces"`
	// An `osDisk` block as defined below.
	OsDisk *LinuxVirtualMachineScaleSetOsDisk `pulumi:"osDisk"`
	// Should Azure over-provision Virtual Machines in this Scale Set? This means that multiple Virtual Machines will be provisioned and Azure will keep the instances which become available first - which improves provisioning success rates and improves deployment time. You're not billed for these over-provisioned VM's and they don't count towards the Subscription Quota. Defaults to `false`.
	Overprovision *bool `pulumi:"overprovision"`
	Plan *LinuxVirtualMachineScaleSetPlan `pulumi:"plan"`
	// The Priority of this Virtual Machine Scale Set. Possible values are `Regular` and `Spot`. Defaults to `Regular`. Changing this value forces a new resource.
	Priority *string `pulumi:"priority"`
	// Should the Azure VM Agent be provisioned on each Virtual Machine in the Scale Set? Defaults to `true`. Changing this value forces a new resource to be created.
	ProvisionVmAgent *bool `pulumi:"provisionVmAgent"`
	// The ID of the Proximity Placement Group in which the Virtual Machine Scale Set should be assigned to. Changing this forces a new resource to be created.
	ProximityPlacementGroupId *string `pulumi:"proximityPlacementGroupId"`
	// The name of the Resource Group in which the Linux Virtual Machine Scale Set should be exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A `rollingUpgradePolicy` block as defined below. This is Required and can only be specified when `upgradeMode` is set to `Automatic` or `Rolling`.
	RollingUpgradePolicy *LinuxVirtualMachineScaleSetRollingUpgradePolicy `pulumi:"rollingUpgradePolicy"`
	// One or more `secret` blocks as defined below.
	Secrets []LinuxVirtualMachineScaleSetSecret `pulumi:"secrets"`
	// Should this Virtual Machine Scale Set be limited to a Single Placement Group, which means the number of instances will be capped at 100 Virtual Machines. Defaults to `true`.
	SinglePlacementGroup *bool `pulumi:"singlePlacementGroup"`
	// The Virtual Machine SKU for the Scale Set, such as `Standard_F2`.
	Sku *string `pulumi:"sku"`
	// The ID of an Image which each Virtual Machine in this Scale Set should be based on.
	SourceImageId *string `pulumi:"sourceImageId"`
	// A `sourceImageReference` block as defined below.
	SourceImageReference *LinuxVirtualMachineScaleSetSourceImageReference `pulumi:"sourceImageReference"`
	// A mapping of tags which should be assigned to this Virtual Machine Scale Set.
	Tags map[string]string `pulumi:"tags"`
	// The Unique ID for this Linux Virtual Machine Scale Set.
	UniqueId *string `pulumi:"uniqueId"`
	// Specifies how Upgrades (e.g. changing the Image/SKU) should be performed to Virtual Machine Instances. Possible values are `Automatic`, `Manual` and `Rolling`. Defaults to `Manual`.
	UpgradeMode *string `pulumi:"upgradeMode"`
	// Should the Virtual Machines in this Scale Set be strictly evenly distributed across Availability Zones? Defaults to `false`. Changing this forces a new resource to be created.
	ZoneBalance *bool `pulumi:"zoneBalance"`
	// A list of Availability Zones in which the Virtual Machines in this Scale Set should be created in. Changing this forces a new resource to be created.
	Zones []string `pulumi:"zones"`
}

type LinuxVirtualMachineScaleSetState struct {
	// A `additionalCapabilities` block as defined below.
	AdditionalCapabilities LinuxVirtualMachineScaleSetAdditionalCapabilitiesPtrInput
	// The Password which should be used for the local-administrator on this Virtual Machine. Changing this forces a new resource to be created.
	AdminPassword pulumi.StringPtrInput
	// One or more `adminSshKey` blocks as defined below.
	AdminSshKeys LinuxVirtualMachineScaleSetAdminSshKeyArrayInput
	// The username of the local administrator on each Virtual Machine Scale Set instance. Changing this forces a new resource to be created.
	AdminUsername pulumi.StringPtrInput
	// A `automaticOsUpgradePolicy` block as defined below. This is Required and can only be specified when `upgradeMode` is set to `Automatic`.
	AutomaticOsUpgradePolicy LinuxVirtualMachineScaleSetAutomaticOsUpgradePolicyPtrInput
	// A `bootDiagnostics` block as defined below.
	BootDiagnostics LinuxVirtualMachineScaleSetBootDiagnosticsPtrInput
	// The prefix which should be used for the name of the Virtual Machines in this Scale Set. If unspecified this defaults to the value for the `name` field.
	ComputerNamePrefix pulumi.StringPtrInput
	// The Base64-Encoded Custom Data which should be used for this Virtual Machine Scale Set.
	CustomData pulumi.StringPtrInput
	// One or more `dataDisk` blocks as defined below.
	DataDisks LinuxVirtualMachineScaleSetDataDiskArrayInput
	// Should Password Authentication be disabled on this Virtual Machine Scale Set? Defaults to `true`.
	DisablePasswordAuthentication pulumi.BoolPtrInput
	// Should Virtual Machine Extensions be run on Overprovisioned Virtual Machines in the Scale Set? Defaults to `false`.
	DoNotRunExtensionsOnOverprovisionedMachines pulumi.BoolPtrInput
	// The Policy which should be used Virtual Machines are Evicted from the Scale Set. Changing this forces a new resource to be created.
	EvictionPolicy pulumi.StringPtrInput
	// The ID of a Load Balancer Probe which should be used to determine the health of an instance. Changing this forces a new resource to be created. This is Required and can only be specified when `upgradeMode` is set to `Automatic` or `Rolling`.
	HealthProbeId pulumi.StringPtrInput
	// A `identity` block as defined below.
	Identity LinuxVirtualMachineScaleSetIdentityPtrInput
	// The number of Virtual Machines in the Scale Set.
	Instances pulumi.IntPtrInput
	// The Azure location where the Linux Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The maximum price you're willing to pay for each Virtual Machine in this Scale Set, in US Dollars; which must be greater than the current spot price. If this bid price falls below the current spot price the Virtual Machines in the Scale Set will be evicted using the `evictionPolicy`. Defaults to `-1`, which means that each Virtual Machine in this Scale Set should not be evicted for price reasons.
	MaxBidPrice pulumi.Float64PtrInput
	// The name of the Linux Virtual Machine Scale Set. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// One or more `networkInterface` blocks as defined below.
	NetworkInterfaces LinuxVirtualMachineScaleSetNetworkInterfaceArrayInput
	// An `osDisk` block as defined below.
	OsDisk LinuxVirtualMachineScaleSetOsDiskPtrInput
	// Should Azure over-provision Virtual Machines in this Scale Set? This means that multiple Virtual Machines will be provisioned and Azure will keep the instances which become available first - which improves provisioning success rates and improves deployment time. You're not billed for these over-provisioned VM's and they don't count towards the Subscription Quota. Defaults to `false`.
	Overprovision pulumi.BoolPtrInput
	Plan LinuxVirtualMachineScaleSetPlanPtrInput
	// The Priority of this Virtual Machine Scale Set. Possible values are `Regular` and `Spot`. Defaults to `Regular`. Changing this value forces a new resource.
	Priority pulumi.StringPtrInput
	// Should the Azure VM Agent be provisioned on each Virtual Machine in the Scale Set? Defaults to `true`. Changing this value forces a new resource to be created.
	ProvisionVmAgent pulumi.BoolPtrInput
	// The ID of the Proximity Placement Group in which the Virtual Machine Scale Set should be assigned to. Changing this forces a new resource to be created.
	ProximityPlacementGroupId pulumi.StringPtrInput
	// The name of the Resource Group in which the Linux Virtual Machine Scale Set should be exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A `rollingUpgradePolicy` block as defined below. This is Required and can only be specified when `upgradeMode` is set to `Automatic` or `Rolling`.
	RollingUpgradePolicy LinuxVirtualMachineScaleSetRollingUpgradePolicyPtrInput
	// One or more `secret` blocks as defined below.
	Secrets LinuxVirtualMachineScaleSetSecretArrayInput
	// Should this Virtual Machine Scale Set be limited to a Single Placement Group, which means the number of instances will be capped at 100 Virtual Machines. Defaults to `true`.
	SinglePlacementGroup pulumi.BoolPtrInput
	// The Virtual Machine SKU for the Scale Set, such as `Standard_F2`.
	Sku pulumi.StringPtrInput
	// The ID of an Image which each Virtual Machine in this Scale Set should be based on.
	SourceImageId pulumi.StringPtrInput
	// A `sourceImageReference` block as defined below.
	SourceImageReference LinuxVirtualMachineScaleSetSourceImageReferencePtrInput
	// A mapping of tags which should be assigned to this Virtual Machine Scale Set.
	Tags pulumi.StringMapInput
	// The Unique ID for this Linux Virtual Machine Scale Set.
	UniqueId pulumi.StringPtrInput
	// Specifies how Upgrades (e.g. changing the Image/SKU) should be performed to Virtual Machine Instances. Possible values are `Automatic`, `Manual` and `Rolling`. Defaults to `Manual`.
	UpgradeMode pulumi.StringPtrInput
	// Should the Virtual Machines in this Scale Set be strictly evenly distributed across Availability Zones? Defaults to `false`. Changing this forces a new resource to be created.
	ZoneBalance pulumi.BoolPtrInput
	// A list of Availability Zones in which the Virtual Machines in this Scale Set should be created in. Changing this forces a new resource to be created.
	Zones pulumi.StringArrayInput
}

func (LinuxVirtualMachineScaleSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*linuxVirtualMachineScaleSetState)(nil)).Elem()
}

type linuxVirtualMachineScaleSetArgs struct {
	// A `additionalCapabilities` block as defined below.
	AdditionalCapabilities *LinuxVirtualMachineScaleSetAdditionalCapabilities `pulumi:"additionalCapabilities"`
	// The Password which should be used for the local-administrator on this Virtual Machine. Changing this forces a new resource to be created.
	AdminPassword *string `pulumi:"adminPassword"`
	// One or more `adminSshKey` blocks as defined below.
	AdminSshKeys []LinuxVirtualMachineScaleSetAdminSshKey `pulumi:"adminSshKeys"`
	// The username of the local administrator on each Virtual Machine Scale Set instance. Changing this forces a new resource to be created.
	AdminUsername string `pulumi:"adminUsername"`
	// A `automaticOsUpgradePolicy` block as defined below. This is Required and can only be specified when `upgradeMode` is set to `Automatic`.
	AutomaticOsUpgradePolicy *LinuxVirtualMachineScaleSetAutomaticOsUpgradePolicy `pulumi:"automaticOsUpgradePolicy"`
	// A `bootDiagnostics` block as defined below.
	BootDiagnostics *LinuxVirtualMachineScaleSetBootDiagnostics `pulumi:"bootDiagnostics"`
	// The prefix which should be used for the name of the Virtual Machines in this Scale Set. If unspecified this defaults to the value for the `name` field.
	ComputerNamePrefix *string `pulumi:"computerNamePrefix"`
	// The Base64-Encoded Custom Data which should be used for this Virtual Machine Scale Set.
	CustomData *string `pulumi:"customData"`
	// One or more `dataDisk` blocks as defined below.
	DataDisks []LinuxVirtualMachineScaleSetDataDisk `pulumi:"dataDisks"`
	// Should Password Authentication be disabled on this Virtual Machine Scale Set? Defaults to `true`.
	DisablePasswordAuthentication *bool `pulumi:"disablePasswordAuthentication"`
	// Should Virtual Machine Extensions be run on Overprovisioned Virtual Machines in the Scale Set? Defaults to `false`.
	DoNotRunExtensionsOnOverprovisionedMachines *bool `pulumi:"doNotRunExtensionsOnOverprovisionedMachines"`
	// The Policy which should be used Virtual Machines are Evicted from the Scale Set. Changing this forces a new resource to be created.
	EvictionPolicy *string `pulumi:"evictionPolicy"`
	// The ID of a Load Balancer Probe which should be used to determine the health of an instance. Changing this forces a new resource to be created. This is Required and can only be specified when `upgradeMode` is set to `Automatic` or `Rolling`.
	HealthProbeId *string `pulumi:"healthProbeId"`
	// A `identity` block as defined below.
	Identity *LinuxVirtualMachineScaleSetIdentity `pulumi:"identity"`
	// The number of Virtual Machines in the Scale Set.
	Instances int `pulumi:"instances"`
	// The Azure location where the Linux Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The maximum price you're willing to pay for each Virtual Machine in this Scale Set, in US Dollars; which must be greater than the current spot price. If this bid price falls below the current spot price the Virtual Machines in the Scale Set will be evicted using the `evictionPolicy`. Defaults to `-1`, which means that each Virtual Machine in this Scale Set should not be evicted for price reasons.
	MaxBidPrice *float64 `pulumi:"maxBidPrice"`
	// The name of the Linux Virtual Machine Scale Set. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// One or more `networkInterface` blocks as defined below.
	NetworkInterfaces []LinuxVirtualMachineScaleSetNetworkInterface `pulumi:"networkInterfaces"`
	// An `osDisk` block as defined below.
	OsDisk LinuxVirtualMachineScaleSetOsDisk `pulumi:"osDisk"`
	// Should Azure over-provision Virtual Machines in this Scale Set? This means that multiple Virtual Machines will be provisioned and Azure will keep the instances which become available first - which improves provisioning success rates and improves deployment time. You're not billed for these over-provisioned VM's and they don't count towards the Subscription Quota. Defaults to `false`.
	Overprovision *bool `pulumi:"overprovision"`
	Plan *LinuxVirtualMachineScaleSetPlan `pulumi:"plan"`
	// The Priority of this Virtual Machine Scale Set. Possible values are `Regular` and `Spot`. Defaults to `Regular`. Changing this value forces a new resource.
	Priority *string `pulumi:"priority"`
	// Should the Azure VM Agent be provisioned on each Virtual Machine in the Scale Set? Defaults to `true`. Changing this value forces a new resource to be created.
	ProvisionVmAgent *bool `pulumi:"provisionVmAgent"`
	// The ID of the Proximity Placement Group in which the Virtual Machine Scale Set should be assigned to. Changing this forces a new resource to be created.
	ProximityPlacementGroupId *string `pulumi:"proximityPlacementGroupId"`
	// The name of the Resource Group in which the Linux Virtual Machine Scale Set should be exist. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A `rollingUpgradePolicy` block as defined below. This is Required and can only be specified when `upgradeMode` is set to `Automatic` or `Rolling`.
	RollingUpgradePolicy *LinuxVirtualMachineScaleSetRollingUpgradePolicy `pulumi:"rollingUpgradePolicy"`
	// One or more `secret` blocks as defined below.
	Secrets []LinuxVirtualMachineScaleSetSecret `pulumi:"secrets"`
	// Should this Virtual Machine Scale Set be limited to a Single Placement Group, which means the number of instances will be capped at 100 Virtual Machines. Defaults to `true`.
	SinglePlacementGroup *bool `pulumi:"singlePlacementGroup"`
	// The Virtual Machine SKU for the Scale Set, such as `Standard_F2`.
	Sku string `pulumi:"sku"`
	// The ID of an Image which each Virtual Machine in this Scale Set should be based on.
	SourceImageId *string `pulumi:"sourceImageId"`
	// A `sourceImageReference` block as defined below.
	SourceImageReference *LinuxVirtualMachineScaleSetSourceImageReference `pulumi:"sourceImageReference"`
	// A mapping of tags which should be assigned to this Virtual Machine Scale Set.
	Tags map[string]string `pulumi:"tags"`
	// Specifies how Upgrades (e.g. changing the Image/SKU) should be performed to Virtual Machine Instances. Possible values are `Automatic`, `Manual` and `Rolling`. Defaults to `Manual`.
	UpgradeMode *string `pulumi:"upgradeMode"`
	// Should the Virtual Machines in this Scale Set be strictly evenly distributed across Availability Zones? Defaults to `false`. Changing this forces a new resource to be created.
	ZoneBalance *bool `pulumi:"zoneBalance"`
	// A list of Availability Zones in which the Virtual Machines in this Scale Set should be created in. Changing this forces a new resource to be created.
	Zones []string `pulumi:"zones"`
}

// The set of arguments for constructing a LinuxVirtualMachineScaleSet resource.
type LinuxVirtualMachineScaleSetArgs struct {
	// A `additionalCapabilities` block as defined below.
	AdditionalCapabilities LinuxVirtualMachineScaleSetAdditionalCapabilitiesPtrInput
	// The Password which should be used for the local-administrator on this Virtual Machine. Changing this forces a new resource to be created.
	AdminPassword pulumi.StringPtrInput
	// One or more `adminSshKey` blocks as defined below.
	AdminSshKeys LinuxVirtualMachineScaleSetAdminSshKeyArrayInput
	// The username of the local administrator on each Virtual Machine Scale Set instance. Changing this forces a new resource to be created.
	AdminUsername pulumi.StringInput
	// A `automaticOsUpgradePolicy` block as defined below. This is Required and can only be specified when `upgradeMode` is set to `Automatic`.
	AutomaticOsUpgradePolicy LinuxVirtualMachineScaleSetAutomaticOsUpgradePolicyPtrInput
	// A `bootDiagnostics` block as defined below.
	BootDiagnostics LinuxVirtualMachineScaleSetBootDiagnosticsPtrInput
	// The prefix which should be used for the name of the Virtual Machines in this Scale Set. If unspecified this defaults to the value for the `name` field.
	ComputerNamePrefix pulumi.StringPtrInput
	// The Base64-Encoded Custom Data which should be used for this Virtual Machine Scale Set.
	CustomData pulumi.StringPtrInput
	// One or more `dataDisk` blocks as defined below.
	DataDisks LinuxVirtualMachineScaleSetDataDiskArrayInput
	// Should Password Authentication be disabled on this Virtual Machine Scale Set? Defaults to `true`.
	DisablePasswordAuthentication pulumi.BoolPtrInput
	// Should Virtual Machine Extensions be run on Overprovisioned Virtual Machines in the Scale Set? Defaults to `false`.
	DoNotRunExtensionsOnOverprovisionedMachines pulumi.BoolPtrInput
	// The Policy which should be used Virtual Machines are Evicted from the Scale Set. Changing this forces a new resource to be created.
	EvictionPolicy pulumi.StringPtrInput
	// The ID of a Load Balancer Probe which should be used to determine the health of an instance. Changing this forces a new resource to be created. This is Required and can only be specified when `upgradeMode` is set to `Automatic` or `Rolling`.
	HealthProbeId pulumi.StringPtrInput
	// A `identity` block as defined below.
	Identity LinuxVirtualMachineScaleSetIdentityPtrInput
	// The number of Virtual Machines in the Scale Set.
	Instances pulumi.IntInput
	// The Azure location where the Linux Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The maximum price you're willing to pay for each Virtual Machine in this Scale Set, in US Dollars; which must be greater than the current spot price. If this bid price falls below the current spot price the Virtual Machines in the Scale Set will be evicted using the `evictionPolicy`. Defaults to `-1`, which means that each Virtual Machine in this Scale Set should not be evicted for price reasons.
	MaxBidPrice pulumi.Float64PtrInput
	// The name of the Linux Virtual Machine Scale Set. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// One or more `networkInterface` blocks as defined below.
	NetworkInterfaces LinuxVirtualMachineScaleSetNetworkInterfaceArrayInput
	// An `osDisk` block as defined below.
	OsDisk LinuxVirtualMachineScaleSetOsDiskInput
	// Should Azure over-provision Virtual Machines in this Scale Set? This means that multiple Virtual Machines will be provisioned and Azure will keep the instances which become available first - which improves provisioning success rates and improves deployment time. You're not billed for these over-provisioned VM's and they don't count towards the Subscription Quota. Defaults to `false`.
	Overprovision pulumi.BoolPtrInput
	Plan LinuxVirtualMachineScaleSetPlanPtrInput
	// The Priority of this Virtual Machine Scale Set. Possible values are `Regular` and `Spot`. Defaults to `Regular`. Changing this value forces a new resource.
	Priority pulumi.StringPtrInput
	// Should the Azure VM Agent be provisioned on each Virtual Machine in the Scale Set? Defaults to `true`. Changing this value forces a new resource to be created.
	ProvisionVmAgent pulumi.BoolPtrInput
	// The ID of the Proximity Placement Group in which the Virtual Machine Scale Set should be assigned to. Changing this forces a new resource to be created.
	ProximityPlacementGroupId pulumi.StringPtrInput
	// The name of the Resource Group in which the Linux Virtual Machine Scale Set should be exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A `rollingUpgradePolicy` block as defined below. This is Required and can only be specified when `upgradeMode` is set to `Automatic` or `Rolling`.
	RollingUpgradePolicy LinuxVirtualMachineScaleSetRollingUpgradePolicyPtrInput
	// One or more `secret` blocks as defined below.
	Secrets LinuxVirtualMachineScaleSetSecretArrayInput
	// Should this Virtual Machine Scale Set be limited to a Single Placement Group, which means the number of instances will be capped at 100 Virtual Machines. Defaults to `true`.
	SinglePlacementGroup pulumi.BoolPtrInput
	// The Virtual Machine SKU for the Scale Set, such as `Standard_F2`.
	Sku pulumi.StringInput
	// The ID of an Image which each Virtual Machine in this Scale Set should be based on.
	SourceImageId pulumi.StringPtrInput
	// A `sourceImageReference` block as defined below.
	SourceImageReference LinuxVirtualMachineScaleSetSourceImageReferencePtrInput
	// A mapping of tags which should be assigned to this Virtual Machine Scale Set.
	Tags pulumi.StringMapInput
	// Specifies how Upgrades (e.g. changing the Image/SKU) should be performed to Virtual Machine Instances. Possible values are `Automatic`, `Manual` and `Rolling`. Defaults to `Manual`.
	UpgradeMode pulumi.StringPtrInput
	// Should the Virtual Machines in this Scale Set be strictly evenly distributed across Availability Zones? Defaults to `false`. Changing this forces a new resource to be created.
	ZoneBalance pulumi.BoolPtrInput
	// A list of Availability Zones in which the Virtual Machines in this Scale Set should be created in. Changing this forces a new resource to be created.
	Zones pulumi.StringArrayInput
}

func (LinuxVirtualMachineScaleSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*linuxVirtualMachineScaleSetArgs)(nil)).Elem()
}

