# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from . import utilities, tables

class Provider(pulumi.ProviderResource):
    def __init__(__self__, resource_name, opts=None, client_certificate_password=None, client_certificate_path=None, client_id=None, client_secret=None, environment=None, msi_endpoint=None, partner_id=None, skip_credentials_validation=None, skip_provider_registration=None, subscription_id=None, tenant_id=None, use_msi=None, __name__=None, __opts__=None):
        """
        The provider type for the azurerm package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://pulumi.io/reference/programming-model.html#providers) for more information.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/index.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if client_certificate_password is None:
            client_certificate_password = (utilities.get_env('ARM_CLIENT_CERTIFICATE_PASSWORD') or '')
        __props__['client_certificate_password'] = client_certificate_password

        if client_certificate_path is None:
            client_certificate_path = (utilities.get_env('ARM_CLIENT_CERTIFICATE_PATH') or '')
        __props__['client_certificate_path'] = client_certificate_path

        if client_id is None:
            client_id = (utilities.get_env('ARM_CLIENT_ID') or '')
        __props__['client_id'] = client_id

        if client_secret is None:
            client_secret = (utilities.get_env('ARM_CLIENT_SECRET') or '')
        __props__['client_secret'] = client_secret

        if environment is None:
            environment = (utilities.get_env('ARM_ENVIRONMENT') or 'public')
        __props__['environment'] = environment

        if msi_endpoint is None:
            msi_endpoint = (utilities.get_env('ARM_MSI_ENDPOINT') or '')
        __props__['msi_endpoint'] = msi_endpoint

        if partner_id is None:
            partner_id = (utilities.get_env('ARM_PARTNER_ID') or '')
        __props__['partner_id'] = partner_id

        if skip_credentials_validation is None:
            skip_credentials_validation = (utilities.get_env_bool('ARM_SKIP_CREDENTIALS_VALIDATION') or False)
        __props__['skip_credentials_validation'] = pulumi.Output.from_input(skip_credentials_validation).apply(json.dumps) if skip_credentials_validation is not None else None

        if skip_provider_registration is None:
            skip_provider_registration = (utilities.get_env_bool('ARM_SKIP_PROVIDER_REGISTRATION') or False)
        __props__['skip_provider_registration'] = pulumi.Output.from_input(skip_provider_registration).apply(json.dumps) if skip_provider_registration is not None else None

        if subscription_id is None:
            subscription_id = (utilities.get_env('ARM_SUBSCRIPTION_ID') or '')
        __props__['subscription_id'] = subscription_id

        if tenant_id is None:
            tenant_id = (utilities.get_env('ARM_TENANT_ID') or '')
        __props__['tenant_id'] = tenant_id

        if use_msi is None:
            use_msi = (utilities.get_env_bool('ARM_USE_MSI') or False)
        __props__['use_msi'] = pulumi.Output.from_input(use_msi).apply(json.dumps) if use_msi is not None else None

        super(Provider, __self__).__init__(
            'azure',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

