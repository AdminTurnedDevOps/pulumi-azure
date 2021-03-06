// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azure

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type ProviderFeatures struct {
	KeyVault               *ProviderFeaturesKeyVault               `pulumi:"keyVault"`
	VirtualMachine         *ProviderFeaturesVirtualMachine         `pulumi:"virtualMachine"`
	VirtualMachineScaleSet *ProviderFeaturesVirtualMachineScaleSet `pulumi:"virtualMachineScaleSet"`
}

// ProviderFeaturesInput is an input type that accepts ProviderFeaturesArgs and ProviderFeaturesOutput values.
// You can construct a concrete instance of `ProviderFeaturesInput` via:
//
// 		 ProviderFeaturesArgs{...}
//
type ProviderFeaturesInput interface {
	pulumi.Input

	ToProviderFeaturesOutput() ProviderFeaturesOutput
	ToProviderFeaturesOutputWithContext(context.Context) ProviderFeaturesOutput
}

type ProviderFeaturesArgs struct {
	KeyVault               ProviderFeaturesKeyVaultPtrInput               `pulumi:"keyVault"`
	VirtualMachine         ProviderFeaturesVirtualMachinePtrInput         `pulumi:"virtualMachine"`
	VirtualMachineScaleSet ProviderFeaturesVirtualMachineScaleSetPtrInput `pulumi:"virtualMachineScaleSet"`
}

func (ProviderFeaturesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProviderFeatures)(nil)).Elem()
}

func (i ProviderFeaturesArgs) ToProviderFeaturesOutput() ProviderFeaturesOutput {
	return i.ToProviderFeaturesOutputWithContext(context.Background())
}

func (i ProviderFeaturesArgs) ToProviderFeaturesOutputWithContext(ctx context.Context) ProviderFeaturesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderFeaturesOutput)
}

func (i ProviderFeaturesArgs) ToProviderFeaturesPtrOutput() ProviderFeaturesPtrOutput {
	return i.ToProviderFeaturesPtrOutputWithContext(context.Background())
}

func (i ProviderFeaturesArgs) ToProviderFeaturesPtrOutputWithContext(ctx context.Context) ProviderFeaturesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderFeaturesOutput).ToProviderFeaturesPtrOutputWithContext(ctx)
}

// ProviderFeaturesPtrInput is an input type that accepts ProviderFeaturesArgs, ProviderFeaturesPtr and ProviderFeaturesPtrOutput values.
// You can construct a concrete instance of `ProviderFeaturesPtrInput` via:
//
// 		 ProviderFeaturesArgs{...}
//
//  or:
//
// 		 nil
//
type ProviderFeaturesPtrInput interface {
	pulumi.Input

	ToProviderFeaturesPtrOutput() ProviderFeaturesPtrOutput
	ToProviderFeaturesPtrOutputWithContext(context.Context) ProviderFeaturesPtrOutput
}

type providerFeaturesPtrType ProviderFeaturesArgs

func ProviderFeaturesPtr(v *ProviderFeaturesArgs) ProviderFeaturesPtrInput {
	return (*providerFeaturesPtrType)(v)
}

func (*providerFeaturesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ProviderFeatures)(nil)).Elem()
}

func (i *providerFeaturesPtrType) ToProviderFeaturesPtrOutput() ProviderFeaturesPtrOutput {
	return i.ToProviderFeaturesPtrOutputWithContext(context.Background())
}

func (i *providerFeaturesPtrType) ToProviderFeaturesPtrOutputWithContext(ctx context.Context) ProviderFeaturesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderFeaturesPtrOutput)
}

type ProviderFeaturesOutput struct{ *pulumi.OutputState }

func (ProviderFeaturesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProviderFeatures)(nil)).Elem()
}

func (o ProviderFeaturesOutput) ToProviderFeaturesOutput() ProviderFeaturesOutput {
	return o
}

func (o ProviderFeaturesOutput) ToProviderFeaturesOutputWithContext(ctx context.Context) ProviderFeaturesOutput {
	return o
}

func (o ProviderFeaturesOutput) ToProviderFeaturesPtrOutput() ProviderFeaturesPtrOutput {
	return o.ToProviderFeaturesPtrOutputWithContext(context.Background())
}

func (o ProviderFeaturesOutput) ToProviderFeaturesPtrOutputWithContext(ctx context.Context) ProviderFeaturesPtrOutput {
	return o.ApplyT(func(v ProviderFeatures) *ProviderFeatures {
		return &v
	}).(ProviderFeaturesPtrOutput)
}
func (o ProviderFeaturesOutput) KeyVault() ProviderFeaturesKeyVaultPtrOutput {
	return o.ApplyT(func(v ProviderFeatures) *ProviderFeaturesKeyVault { return v.KeyVault }).(ProviderFeaturesKeyVaultPtrOutput)
}

func (o ProviderFeaturesOutput) VirtualMachine() ProviderFeaturesVirtualMachinePtrOutput {
	return o.ApplyT(func(v ProviderFeatures) *ProviderFeaturesVirtualMachine { return v.VirtualMachine }).(ProviderFeaturesVirtualMachinePtrOutput)
}

func (o ProviderFeaturesOutput) VirtualMachineScaleSet() ProviderFeaturesVirtualMachineScaleSetPtrOutput {
	return o.ApplyT(func(v ProviderFeatures) *ProviderFeaturesVirtualMachineScaleSet { return v.VirtualMachineScaleSet }).(ProviderFeaturesVirtualMachineScaleSetPtrOutput)
}

type ProviderFeaturesPtrOutput struct{ *pulumi.OutputState }

func (ProviderFeaturesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProviderFeatures)(nil)).Elem()
}

func (o ProviderFeaturesPtrOutput) ToProviderFeaturesPtrOutput() ProviderFeaturesPtrOutput {
	return o
}

func (o ProviderFeaturesPtrOutput) ToProviderFeaturesPtrOutputWithContext(ctx context.Context) ProviderFeaturesPtrOutput {
	return o
}

func (o ProviderFeaturesPtrOutput) Elem() ProviderFeaturesOutput {
	return o.ApplyT(func(v *ProviderFeatures) ProviderFeatures { return *v }).(ProviderFeaturesOutput)
}

func (o ProviderFeaturesPtrOutput) KeyVault() ProviderFeaturesKeyVaultPtrOutput {
	return o.ApplyT(func(v ProviderFeatures) *ProviderFeaturesKeyVault { return v.KeyVault }).(ProviderFeaturesKeyVaultPtrOutput)
}

func (o ProviderFeaturesPtrOutput) VirtualMachine() ProviderFeaturesVirtualMachinePtrOutput {
	return o.ApplyT(func(v ProviderFeatures) *ProviderFeaturesVirtualMachine { return v.VirtualMachine }).(ProviderFeaturesVirtualMachinePtrOutput)
}

func (o ProviderFeaturesPtrOutput) VirtualMachineScaleSet() ProviderFeaturesVirtualMachineScaleSetPtrOutput {
	return o.ApplyT(func(v ProviderFeatures) *ProviderFeaturesVirtualMachineScaleSet { return v.VirtualMachineScaleSet }).(ProviderFeaturesVirtualMachineScaleSetPtrOutput)
}

type ProviderFeaturesKeyVault struct {
	PurgeSoftDeleteOnDestroy    *bool `pulumi:"purgeSoftDeleteOnDestroy"`
	RecoverSoftDeletedKeyVaults *bool `pulumi:"recoverSoftDeletedKeyVaults"`
}

// ProviderFeaturesKeyVaultInput is an input type that accepts ProviderFeaturesKeyVaultArgs and ProviderFeaturesKeyVaultOutput values.
// You can construct a concrete instance of `ProviderFeaturesKeyVaultInput` via:
//
// 		 ProviderFeaturesKeyVaultArgs{...}
//
type ProviderFeaturesKeyVaultInput interface {
	pulumi.Input

	ToProviderFeaturesKeyVaultOutput() ProviderFeaturesKeyVaultOutput
	ToProviderFeaturesKeyVaultOutputWithContext(context.Context) ProviderFeaturesKeyVaultOutput
}

type ProviderFeaturesKeyVaultArgs struct {
	PurgeSoftDeleteOnDestroy    pulumi.BoolPtrInput `pulumi:"purgeSoftDeleteOnDestroy"`
	RecoverSoftDeletedKeyVaults pulumi.BoolPtrInput `pulumi:"recoverSoftDeletedKeyVaults"`
}

func (ProviderFeaturesKeyVaultArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProviderFeaturesKeyVault)(nil)).Elem()
}

func (i ProviderFeaturesKeyVaultArgs) ToProviderFeaturesKeyVaultOutput() ProviderFeaturesKeyVaultOutput {
	return i.ToProviderFeaturesKeyVaultOutputWithContext(context.Background())
}

func (i ProviderFeaturesKeyVaultArgs) ToProviderFeaturesKeyVaultOutputWithContext(ctx context.Context) ProviderFeaturesKeyVaultOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderFeaturesKeyVaultOutput)
}

func (i ProviderFeaturesKeyVaultArgs) ToProviderFeaturesKeyVaultPtrOutput() ProviderFeaturesKeyVaultPtrOutput {
	return i.ToProviderFeaturesKeyVaultPtrOutputWithContext(context.Background())
}

func (i ProviderFeaturesKeyVaultArgs) ToProviderFeaturesKeyVaultPtrOutputWithContext(ctx context.Context) ProviderFeaturesKeyVaultPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderFeaturesKeyVaultOutput).ToProviderFeaturesKeyVaultPtrOutputWithContext(ctx)
}

// ProviderFeaturesKeyVaultPtrInput is an input type that accepts ProviderFeaturesKeyVaultArgs, ProviderFeaturesKeyVaultPtr and ProviderFeaturesKeyVaultPtrOutput values.
// You can construct a concrete instance of `ProviderFeaturesKeyVaultPtrInput` via:
//
// 		 ProviderFeaturesKeyVaultArgs{...}
//
//  or:
//
// 		 nil
//
type ProviderFeaturesKeyVaultPtrInput interface {
	pulumi.Input

	ToProviderFeaturesKeyVaultPtrOutput() ProviderFeaturesKeyVaultPtrOutput
	ToProviderFeaturesKeyVaultPtrOutputWithContext(context.Context) ProviderFeaturesKeyVaultPtrOutput
}

type providerFeaturesKeyVaultPtrType ProviderFeaturesKeyVaultArgs

func ProviderFeaturesKeyVaultPtr(v *ProviderFeaturesKeyVaultArgs) ProviderFeaturesKeyVaultPtrInput {
	return (*providerFeaturesKeyVaultPtrType)(v)
}

func (*providerFeaturesKeyVaultPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ProviderFeaturesKeyVault)(nil)).Elem()
}

func (i *providerFeaturesKeyVaultPtrType) ToProviderFeaturesKeyVaultPtrOutput() ProviderFeaturesKeyVaultPtrOutput {
	return i.ToProviderFeaturesKeyVaultPtrOutputWithContext(context.Background())
}

func (i *providerFeaturesKeyVaultPtrType) ToProviderFeaturesKeyVaultPtrOutputWithContext(ctx context.Context) ProviderFeaturesKeyVaultPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderFeaturesKeyVaultPtrOutput)
}

type ProviderFeaturesKeyVaultOutput struct{ *pulumi.OutputState }

func (ProviderFeaturesKeyVaultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProviderFeaturesKeyVault)(nil)).Elem()
}

func (o ProviderFeaturesKeyVaultOutput) ToProviderFeaturesKeyVaultOutput() ProviderFeaturesKeyVaultOutput {
	return o
}

func (o ProviderFeaturesKeyVaultOutput) ToProviderFeaturesKeyVaultOutputWithContext(ctx context.Context) ProviderFeaturesKeyVaultOutput {
	return o
}

func (o ProviderFeaturesKeyVaultOutput) ToProviderFeaturesKeyVaultPtrOutput() ProviderFeaturesKeyVaultPtrOutput {
	return o.ToProviderFeaturesKeyVaultPtrOutputWithContext(context.Background())
}

func (o ProviderFeaturesKeyVaultOutput) ToProviderFeaturesKeyVaultPtrOutputWithContext(ctx context.Context) ProviderFeaturesKeyVaultPtrOutput {
	return o.ApplyT(func(v ProviderFeaturesKeyVault) *ProviderFeaturesKeyVault {
		return &v
	}).(ProviderFeaturesKeyVaultPtrOutput)
}
func (o ProviderFeaturesKeyVaultOutput) PurgeSoftDeleteOnDestroy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ProviderFeaturesKeyVault) *bool { return v.PurgeSoftDeleteOnDestroy }).(pulumi.BoolPtrOutput)
}

func (o ProviderFeaturesKeyVaultOutput) RecoverSoftDeletedKeyVaults() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ProviderFeaturesKeyVault) *bool { return v.RecoverSoftDeletedKeyVaults }).(pulumi.BoolPtrOutput)
}

type ProviderFeaturesKeyVaultPtrOutput struct{ *pulumi.OutputState }

func (ProviderFeaturesKeyVaultPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProviderFeaturesKeyVault)(nil)).Elem()
}

func (o ProviderFeaturesKeyVaultPtrOutput) ToProviderFeaturesKeyVaultPtrOutput() ProviderFeaturesKeyVaultPtrOutput {
	return o
}

func (o ProviderFeaturesKeyVaultPtrOutput) ToProviderFeaturesKeyVaultPtrOutputWithContext(ctx context.Context) ProviderFeaturesKeyVaultPtrOutput {
	return o
}

func (o ProviderFeaturesKeyVaultPtrOutput) Elem() ProviderFeaturesKeyVaultOutput {
	return o.ApplyT(func(v *ProviderFeaturesKeyVault) ProviderFeaturesKeyVault { return *v }).(ProviderFeaturesKeyVaultOutput)
}

func (o ProviderFeaturesKeyVaultPtrOutput) PurgeSoftDeleteOnDestroy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ProviderFeaturesKeyVault) *bool { return v.PurgeSoftDeleteOnDestroy }).(pulumi.BoolPtrOutput)
}

func (o ProviderFeaturesKeyVaultPtrOutput) RecoverSoftDeletedKeyVaults() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ProviderFeaturesKeyVault) *bool { return v.RecoverSoftDeletedKeyVaults }).(pulumi.BoolPtrOutput)
}

type ProviderFeaturesVirtualMachine struct {
	DeleteOsDiskOnDeletion bool `pulumi:"deleteOsDiskOnDeletion"`
}

// ProviderFeaturesVirtualMachineInput is an input type that accepts ProviderFeaturesVirtualMachineArgs and ProviderFeaturesVirtualMachineOutput values.
// You can construct a concrete instance of `ProviderFeaturesVirtualMachineInput` via:
//
// 		 ProviderFeaturesVirtualMachineArgs{...}
//
type ProviderFeaturesVirtualMachineInput interface {
	pulumi.Input

	ToProviderFeaturesVirtualMachineOutput() ProviderFeaturesVirtualMachineOutput
	ToProviderFeaturesVirtualMachineOutputWithContext(context.Context) ProviderFeaturesVirtualMachineOutput
}

type ProviderFeaturesVirtualMachineArgs struct {
	DeleteOsDiskOnDeletion pulumi.BoolInput `pulumi:"deleteOsDiskOnDeletion"`
}

func (ProviderFeaturesVirtualMachineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProviderFeaturesVirtualMachine)(nil)).Elem()
}

func (i ProviderFeaturesVirtualMachineArgs) ToProviderFeaturesVirtualMachineOutput() ProviderFeaturesVirtualMachineOutput {
	return i.ToProviderFeaturesVirtualMachineOutputWithContext(context.Background())
}

func (i ProviderFeaturesVirtualMachineArgs) ToProviderFeaturesVirtualMachineOutputWithContext(ctx context.Context) ProviderFeaturesVirtualMachineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderFeaturesVirtualMachineOutput)
}

func (i ProviderFeaturesVirtualMachineArgs) ToProviderFeaturesVirtualMachinePtrOutput() ProviderFeaturesVirtualMachinePtrOutput {
	return i.ToProviderFeaturesVirtualMachinePtrOutputWithContext(context.Background())
}

func (i ProviderFeaturesVirtualMachineArgs) ToProviderFeaturesVirtualMachinePtrOutputWithContext(ctx context.Context) ProviderFeaturesVirtualMachinePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderFeaturesVirtualMachineOutput).ToProviderFeaturesVirtualMachinePtrOutputWithContext(ctx)
}

// ProviderFeaturesVirtualMachinePtrInput is an input type that accepts ProviderFeaturesVirtualMachineArgs, ProviderFeaturesVirtualMachinePtr and ProviderFeaturesVirtualMachinePtrOutput values.
// You can construct a concrete instance of `ProviderFeaturesVirtualMachinePtrInput` via:
//
// 		 ProviderFeaturesVirtualMachineArgs{...}
//
//  or:
//
// 		 nil
//
type ProviderFeaturesVirtualMachinePtrInput interface {
	pulumi.Input

	ToProviderFeaturesVirtualMachinePtrOutput() ProviderFeaturesVirtualMachinePtrOutput
	ToProviderFeaturesVirtualMachinePtrOutputWithContext(context.Context) ProviderFeaturesVirtualMachinePtrOutput
}

type providerFeaturesVirtualMachinePtrType ProviderFeaturesVirtualMachineArgs

func ProviderFeaturesVirtualMachinePtr(v *ProviderFeaturesVirtualMachineArgs) ProviderFeaturesVirtualMachinePtrInput {
	return (*providerFeaturesVirtualMachinePtrType)(v)
}

func (*providerFeaturesVirtualMachinePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ProviderFeaturesVirtualMachine)(nil)).Elem()
}

func (i *providerFeaturesVirtualMachinePtrType) ToProviderFeaturesVirtualMachinePtrOutput() ProviderFeaturesVirtualMachinePtrOutput {
	return i.ToProviderFeaturesVirtualMachinePtrOutputWithContext(context.Background())
}

func (i *providerFeaturesVirtualMachinePtrType) ToProviderFeaturesVirtualMachinePtrOutputWithContext(ctx context.Context) ProviderFeaturesVirtualMachinePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderFeaturesVirtualMachinePtrOutput)
}

type ProviderFeaturesVirtualMachineOutput struct{ *pulumi.OutputState }

func (ProviderFeaturesVirtualMachineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProviderFeaturesVirtualMachine)(nil)).Elem()
}

func (o ProviderFeaturesVirtualMachineOutput) ToProviderFeaturesVirtualMachineOutput() ProviderFeaturesVirtualMachineOutput {
	return o
}

func (o ProviderFeaturesVirtualMachineOutput) ToProviderFeaturesVirtualMachineOutputWithContext(ctx context.Context) ProviderFeaturesVirtualMachineOutput {
	return o
}

func (o ProviderFeaturesVirtualMachineOutput) ToProviderFeaturesVirtualMachinePtrOutput() ProviderFeaturesVirtualMachinePtrOutput {
	return o.ToProviderFeaturesVirtualMachinePtrOutputWithContext(context.Background())
}

func (o ProviderFeaturesVirtualMachineOutput) ToProviderFeaturesVirtualMachinePtrOutputWithContext(ctx context.Context) ProviderFeaturesVirtualMachinePtrOutput {
	return o.ApplyT(func(v ProviderFeaturesVirtualMachine) *ProviderFeaturesVirtualMachine {
		return &v
	}).(ProviderFeaturesVirtualMachinePtrOutput)
}
func (o ProviderFeaturesVirtualMachineOutput) DeleteOsDiskOnDeletion() pulumi.BoolOutput {
	return o.ApplyT(func(v ProviderFeaturesVirtualMachine) bool { return v.DeleteOsDiskOnDeletion }).(pulumi.BoolOutput)
}

type ProviderFeaturesVirtualMachinePtrOutput struct{ *pulumi.OutputState }

func (ProviderFeaturesVirtualMachinePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProviderFeaturesVirtualMachine)(nil)).Elem()
}

func (o ProviderFeaturesVirtualMachinePtrOutput) ToProviderFeaturesVirtualMachinePtrOutput() ProviderFeaturesVirtualMachinePtrOutput {
	return o
}

func (o ProviderFeaturesVirtualMachinePtrOutput) ToProviderFeaturesVirtualMachinePtrOutputWithContext(ctx context.Context) ProviderFeaturesVirtualMachinePtrOutput {
	return o
}

func (o ProviderFeaturesVirtualMachinePtrOutput) Elem() ProviderFeaturesVirtualMachineOutput {
	return o.ApplyT(func(v *ProviderFeaturesVirtualMachine) ProviderFeaturesVirtualMachine { return *v }).(ProviderFeaturesVirtualMachineOutput)
}

func (o ProviderFeaturesVirtualMachinePtrOutput) DeleteOsDiskOnDeletion() pulumi.BoolOutput {
	return o.ApplyT(func(v ProviderFeaturesVirtualMachine) bool { return v.DeleteOsDiskOnDeletion }).(pulumi.BoolOutput)
}

type ProviderFeaturesVirtualMachineScaleSet struct {
	RollInstancesWhenRequired bool `pulumi:"rollInstancesWhenRequired"`
}

// ProviderFeaturesVirtualMachineScaleSetInput is an input type that accepts ProviderFeaturesVirtualMachineScaleSetArgs and ProviderFeaturesVirtualMachineScaleSetOutput values.
// You can construct a concrete instance of `ProviderFeaturesVirtualMachineScaleSetInput` via:
//
// 		 ProviderFeaturesVirtualMachineScaleSetArgs{...}
//
type ProviderFeaturesVirtualMachineScaleSetInput interface {
	pulumi.Input

	ToProviderFeaturesVirtualMachineScaleSetOutput() ProviderFeaturesVirtualMachineScaleSetOutput
	ToProviderFeaturesVirtualMachineScaleSetOutputWithContext(context.Context) ProviderFeaturesVirtualMachineScaleSetOutput
}

type ProviderFeaturesVirtualMachineScaleSetArgs struct {
	RollInstancesWhenRequired pulumi.BoolInput `pulumi:"rollInstancesWhenRequired"`
}

func (ProviderFeaturesVirtualMachineScaleSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProviderFeaturesVirtualMachineScaleSet)(nil)).Elem()
}

func (i ProviderFeaturesVirtualMachineScaleSetArgs) ToProviderFeaturesVirtualMachineScaleSetOutput() ProviderFeaturesVirtualMachineScaleSetOutput {
	return i.ToProviderFeaturesVirtualMachineScaleSetOutputWithContext(context.Background())
}

func (i ProviderFeaturesVirtualMachineScaleSetArgs) ToProviderFeaturesVirtualMachineScaleSetOutputWithContext(ctx context.Context) ProviderFeaturesVirtualMachineScaleSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderFeaturesVirtualMachineScaleSetOutput)
}

func (i ProviderFeaturesVirtualMachineScaleSetArgs) ToProviderFeaturesVirtualMachineScaleSetPtrOutput() ProviderFeaturesVirtualMachineScaleSetPtrOutput {
	return i.ToProviderFeaturesVirtualMachineScaleSetPtrOutputWithContext(context.Background())
}

func (i ProviderFeaturesVirtualMachineScaleSetArgs) ToProviderFeaturesVirtualMachineScaleSetPtrOutputWithContext(ctx context.Context) ProviderFeaturesVirtualMachineScaleSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderFeaturesVirtualMachineScaleSetOutput).ToProviderFeaturesVirtualMachineScaleSetPtrOutputWithContext(ctx)
}

// ProviderFeaturesVirtualMachineScaleSetPtrInput is an input type that accepts ProviderFeaturesVirtualMachineScaleSetArgs, ProviderFeaturesVirtualMachineScaleSetPtr and ProviderFeaturesVirtualMachineScaleSetPtrOutput values.
// You can construct a concrete instance of `ProviderFeaturesVirtualMachineScaleSetPtrInput` via:
//
// 		 ProviderFeaturesVirtualMachineScaleSetArgs{...}
//
//  or:
//
// 		 nil
//
type ProviderFeaturesVirtualMachineScaleSetPtrInput interface {
	pulumi.Input

	ToProviderFeaturesVirtualMachineScaleSetPtrOutput() ProviderFeaturesVirtualMachineScaleSetPtrOutput
	ToProviderFeaturesVirtualMachineScaleSetPtrOutputWithContext(context.Context) ProviderFeaturesVirtualMachineScaleSetPtrOutput
}

type providerFeaturesVirtualMachineScaleSetPtrType ProviderFeaturesVirtualMachineScaleSetArgs

func ProviderFeaturesVirtualMachineScaleSetPtr(v *ProviderFeaturesVirtualMachineScaleSetArgs) ProviderFeaturesVirtualMachineScaleSetPtrInput {
	return (*providerFeaturesVirtualMachineScaleSetPtrType)(v)
}

func (*providerFeaturesVirtualMachineScaleSetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ProviderFeaturesVirtualMachineScaleSet)(nil)).Elem()
}

func (i *providerFeaturesVirtualMachineScaleSetPtrType) ToProviderFeaturesVirtualMachineScaleSetPtrOutput() ProviderFeaturesVirtualMachineScaleSetPtrOutput {
	return i.ToProviderFeaturesVirtualMachineScaleSetPtrOutputWithContext(context.Background())
}

func (i *providerFeaturesVirtualMachineScaleSetPtrType) ToProviderFeaturesVirtualMachineScaleSetPtrOutputWithContext(ctx context.Context) ProviderFeaturesVirtualMachineScaleSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderFeaturesVirtualMachineScaleSetPtrOutput)
}

type ProviderFeaturesVirtualMachineScaleSetOutput struct{ *pulumi.OutputState }

func (ProviderFeaturesVirtualMachineScaleSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProviderFeaturesVirtualMachineScaleSet)(nil)).Elem()
}

func (o ProviderFeaturesVirtualMachineScaleSetOutput) ToProviderFeaturesVirtualMachineScaleSetOutput() ProviderFeaturesVirtualMachineScaleSetOutput {
	return o
}

func (o ProviderFeaturesVirtualMachineScaleSetOutput) ToProviderFeaturesVirtualMachineScaleSetOutputWithContext(ctx context.Context) ProviderFeaturesVirtualMachineScaleSetOutput {
	return o
}

func (o ProviderFeaturesVirtualMachineScaleSetOutput) ToProviderFeaturesVirtualMachineScaleSetPtrOutput() ProviderFeaturesVirtualMachineScaleSetPtrOutput {
	return o.ToProviderFeaturesVirtualMachineScaleSetPtrOutputWithContext(context.Background())
}

func (o ProviderFeaturesVirtualMachineScaleSetOutput) ToProviderFeaturesVirtualMachineScaleSetPtrOutputWithContext(ctx context.Context) ProviderFeaturesVirtualMachineScaleSetPtrOutput {
	return o.ApplyT(func(v ProviderFeaturesVirtualMachineScaleSet) *ProviderFeaturesVirtualMachineScaleSet {
		return &v
	}).(ProviderFeaturesVirtualMachineScaleSetPtrOutput)
}
func (o ProviderFeaturesVirtualMachineScaleSetOutput) RollInstancesWhenRequired() pulumi.BoolOutput {
	return o.ApplyT(func(v ProviderFeaturesVirtualMachineScaleSet) bool { return v.RollInstancesWhenRequired }).(pulumi.BoolOutput)
}

type ProviderFeaturesVirtualMachineScaleSetPtrOutput struct{ *pulumi.OutputState }

func (ProviderFeaturesVirtualMachineScaleSetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProviderFeaturesVirtualMachineScaleSet)(nil)).Elem()
}

func (o ProviderFeaturesVirtualMachineScaleSetPtrOutput) ToProviderFeaturesVirtualMachineScaleSetPtrOutput() ProviderFeaturesVirtualMachineScaleSetPtrOutput {
	return o
}

func (o ProviderFeaturesVirtualMachineScaleSetPtrOutput) ToProviderFeaturesVirtualMachineScaleSetPtrOutputWithContext(ctx context.Context) ProviderFeaturesVirtualMachineScaleSetPtrOutput {
	return o
}

func (o ProviderFeaturesVirtualMachineScaleSetPtrOutput) Elem() ProviderFeaturesVirtualMachineScaleSetOutput {
	return o.ApplyT(func(v *ProviderFeaturesVirtualMachineScaleSet) ProviderFeaturesVirtualMachineScaleSet { return *v }).(ProviderFeaturesVirtualMachineScaleSetOutput)
}

func (o ProviderFeaturesVirtualMachineScaleSetPtrOutput) RollInstancesWhenRequired() pulumi.BoolOutput {
	return o.ApplyT(func(v ProviderFeaturesVirtualMachineScaleSet) bool { return v.RollInstancesWhenRequired }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(ProviderFeaturesOutput{})
	pulumi.RegisterOutputType(ProviderFeaturesPtrOutput{})
	pulumi.RegisterOutputType(ProviderFeaturesKeyVaultOutput{})
	pulumi.RegisterOutputType(ProviderFeaturesKeyVaultPtrOutput{})
	pulumi.RegisterOutputType(ProviderFeaturesVirtualMachineOutput{})
	pulumi.RegisterOutputType(ProviderFeaturesVirtualMachinePtrOutput{})
	pulumi.RegisterOutputType(ProviderFeaturesVirtualMachineScaleSetOutput{})
	pulumi.RegisterOutputType(ProviderFeaturesVirtualMachineScaleSetPtrOutput{})
}
