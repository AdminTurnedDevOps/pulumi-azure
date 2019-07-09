# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class AvailabilitySet(pulumi.CustomResource):
    location: pulumi.Output[str]
    """
    Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
    """
    managed: pulumi.Output[bool]
    """
    Specifies whether the availability set is managed or not. Possible values are `true` (to specify aligned) or `false` (to specify classic). Default is `false`.
    """
    name: pulumi.Output[str]
    """
    Specifies the name of the availability set. Changing this forces a new resource to be created.
    """
    platform_fault_domain_count: pulumi.Output[float]
    """
    Specifies the number of fault domains that are used. Defaults to 3.
    """
    platform_update_domain_count: pulumi.Output[float]
    """
    Specifies the number of update domains that are used. Defaults to 5.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the resource group in which to create the availability set. Changing this forces a new resource to be created.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    def __init__(__self__, resource_name, opts=None, location=None, managed=None, name=None, platform_fault_domain_count=None, platform_update_domain_count=None, resource_group_name=None, tags=None, __name__=None, __opts__=None):
        """
        Manages an availability set for virtual machines.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] managed: Specifies whether the availability set is managed or not. Possible values are `true` (to specify aligned) or `false` (to specify classic). Default is `false`.
        :param pulumi.Input[str] name: Specifies the name of the availability set. Changing this forces a new resource to be created.
        :param pulumi.Input[float] platform_fault_domain_count: Specifies the number of fault domains that are used. Defaults to 3.
        :param pulumi.Input[float] platform_update_domain_count: Specifies the number of update domains that are used. Defaults to 5.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the availability set. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/availability_set.html.markdown.
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

        __props__['location'] = location

        __props__['managed'] = managed

        __props__['name'] = name

        __props__['platform_fault_domain_count'] = platform_fault_domain_count

        __props__['platform_update_domain_count'] = platform_update_domain_count

        if resource_group_name is None:
            raise TypeError("Missing required property 'resource_group_name'")
        __props__['resource_group_name'] = resource_group_name

        __props__['tags'] = tags

        super(AvailabilitySet, __self__).__init__(
            'azure:compute/availabilitySet:AvailabilitySet',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

