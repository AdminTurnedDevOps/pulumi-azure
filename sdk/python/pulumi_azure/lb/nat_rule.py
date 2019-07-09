# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class NatRule(pulumi.CustomResource):
    backend_ip_configuration_id: pulumi.Output[str]
    backend_port: pulumi.Output[float]
    """
    The port used for internal connections on the endpoint. Possible values range between 1 and 65535, inclusive.
    """
    enable_floating_ip: pulumi.Output[bool]
    """
    Enables the Floating IP Capacity, required to configure a SQL AlwaysOn Availability Group.
    """
    frontend_ip_configuration_id: pulumi.Output[str]
    frontend_ip_configuration_name: pulumi.Output[str]
    """
    The name of the frontend IP configuration exposing this rule.
    """
    frontend_port: pulumi.Output[float]
    """
    The port for the external endpoint. Port numbers for each Rule must be unique within the Load Balancer. Possible values range between 1 and 65534, inclusive.
    """
    loadbalancer_id: pulumi.Output[str]
    """
    The ID of the Load Balancer in which to create the NAT Rule.
    """
    location: pulumi.Output[str]
    name: pulumi.Output[str]
    """
    Specifies the name of the NAT Rule.
    """
    protocol: pulumi.Output[str]
    """
    The transport protocol for the external endpoint. Possible values are `Udp`, `Tcp` or `All`.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the resource group in which to create the resource.
    """
    def __init__(__self__, resource_name, opts=None, backend_port=None, enable_floating_ip=None, frontend_ip_configuration_name=None, frontend_port=None, loadbalancer_id=None, location=None, name=None, protocol=None, resource_group_name=None, __name__=None, __opts__=None):
        """
        Manages a Load Balancer NAT Rule.
        
        > **NOTE** When using this resource, the Load Balancer needs to have a FrontEnd IP Configuration Attached
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] backend_port: The port used for internal connections on the endpoint. Possible values range between 1 and 65535, inclusive.
        :param pulumi.Input[bool] enable_floating_ip: Enables the Floating IP Capacity, required to configure a SQL AlwaysOn Availability Group.
        :param pulumi.Input[str] frontend_ip_configuration_name: The name of the frontend IP configuration exposing this rule.
        :param pulumi.Input[float] frontend_port: The port for the external endpoint. Port numbers for each Rule must be unique within the Load Balancer. Possible values range between 1 and 65534, inclusive.
        :param pulumi.Input[str] loadbalancer_id: The ID of the Load Balancer in which to create the NAT Rule.
        :param pulumi.Input[str] name: Specifies the name of the NAT Rule.
        :param pulumi.Input[str] protocol: The transport protocol for the external endpoint. Possible values are `Udp`, `Tcp` or `All`.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the resource.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/lb_nat_rule.html.markdown.
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

        if backend_port is None:
            raise TypeError("Missing required property 'backend_port'")
        __props__['backend_port'] = backend_port

        __props__['enable_floating_ip'] = enable_floating_ip

        if frontend_ip_configuration_name is None:
            raise TypeError("Missing required property 'frontend_ip_configuration_name'")
        __props__['frontend_ip_configuration_name'] = frontend_ip_configuration_name

        if frontend_port is None:
            raise TypeError("Missing required property 'frontend_port'")
        __props__['frontend_port'] = frontend_port

        if loadbalancer_id is None:
            raise TypeError("Missing required property 'loadbalancer_id'")
        __props__['loadbalancer_id'] = loadbalancer_id

        __props__['location'] = location

        __props__['name'] = name

        if protocol is None:
            raise TypeError("Missing required property 'protocol'")
        __props__['protocol'] = protocol

        if resource_group_name is None:
            raise TypeError("Missing required property 'resource_group_name'")
        __props__['resource_group_name'] = resource_group_name

        __props__['backend_ip_configuration_id'] = None
        __props__['frontend_ip_configuration_id'] = None

        super(NatRule, __self__).__init__(
            'azure:lb/natRule:NatRule',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

