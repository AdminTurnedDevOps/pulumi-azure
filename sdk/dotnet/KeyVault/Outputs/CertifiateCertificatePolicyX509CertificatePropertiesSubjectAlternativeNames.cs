// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.KeyVault.Outputs
{

    [OutputType]
    public sealed class CertifiateCertificatePolicyX509CertificatePropertiesSubjectAlternativeNames
    {
        /// <summary>
        /// A list of alternative DNS names (FQDNs) identified by the Certificate. Changing this forces a new resource to be created.
        /// </summary>
        public readonly ImmutableArray<string> DnsNames;
        /// <summary>
        /// A list of email addresses identified by this Certificate. Changing this forces a new resource to be created.
        /// </summary>
        public readonly ImmutableArray<string> Emails;
        /// <summary>
        /// A list of User Principal Names identified by the Certificate. Changing this forces a new resource to be created.
        /// </summary>
        public readonly ImmutableArray<string> Upns;

        [OutputConstructor]
        private CertifiateCertificatePolicyX509CertificatePropertiesSubjectAlternativeNames(
            ImmutableArray<string> dnsNames,

            ImmutableArray<string> emails,

            ImmutableArray<string> upns)
        {
            DnsNames = dnsNames;
            Emails = emails;
            Upns = upns;
        }
    }
}