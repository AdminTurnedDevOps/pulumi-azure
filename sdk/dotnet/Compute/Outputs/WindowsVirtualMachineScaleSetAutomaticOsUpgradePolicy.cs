// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Compute.Outputs
{

    [OutputType]
    public sealed class WindowsVirtualMachineScaleSetAutomaticOsUpgradePolicy
    {
        /// <summary>
        /// Should automatic rollbacks be disabled? Changing this forces a new resource to be created.
        /// </summary>
        public readonly bool DisableAutomaticRollback;
        /// <summary>
        /// Should OS Upgrades automatically be applied to Scale Set instances in a rolling fashion when a newer version of the OS Image becomes available? Changing this forces a new resource to be created.
        /// </summary>
        public readonly bool EnableAutomaticOsUpgrade;

        [OutputConstructor]
        private WindowsVirtualMachineScaleSetAutomaticOsUpgradePolicy(
            bool disableAutomaticRollback,

            bool enableAutomaticOsUpgrade)
        {
            DisableAutomaticRollback = disableAutomaticRollback;
            EnableAutomaticOsUpgrade = enableAutomaticOsUpgrade;
        }
    }
}
