# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class LocalNetworkGateway(pulumi.CustomResource):
    address_spaces: pulumi.Output[list]
    """
    The list of string CIDRs representing the
    address spaces the gateway exposes.
    """
    bgp_settings: pulumi.Output[dict]
    """
    A `bgp_settings` block as defined below containing the
    Local Network Gateway's BGP speaker settings.

      * `asn` (`float`) - The BGP speaker's ASN.
      * `bgpPeeringAddress` (`str`) - The BGP peering address and BGP identifier
        of this BGP speaker.
      * `peerWeight` (`float`) - The weight added to routes learned from this
        BGP speaker.
    """
    gateway_address: pulumi.Output[str]
    """
    The IP address of the gateway to which to
    connect.
    """
    location: pulumi.Output[str]
    """
    The location/region where the local network gateway is
    created. Changing this forces a new resource to be created.
    """
    name: pulumi.Output[str]
    """
    The name of the local network gateway. Changing this
    forces a new resource to be created.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the resource group in which to
    create the local network gateway.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    def __init__(__self__, resource_name, opts=None, address_spaces=None, bgp_settings=None, gateway_address=None, location=None, name=None, resource_group_name=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a local network gateway connection over which specific connections can be configured.



        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] address_spaces: The list of string CIDRs representing the
               address spaces the gateway exposes.
        :param pulumi.Input[dict] bgp_settings: A `bgp_settings` block as defined below containing the
               Local Network Gateway's BGP speaker settings.
        :param pulumi.Input[str] gateway_address: The IP address of the gateway to which to
               connect.
        :param pulumi.Input[str] location: The location/region where the local network gateway is
               created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the local network gateway. Changing this
               forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to
               create the local network gateway.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.

        The **bgp_settings** object supports the following:

          * `asn` (`pulumi.Input[float]`) - The BGP speaker's ASN.
          * `bgpPeeringAddress` (`pulumi.Input[str]`) - The BGP peering address and BGP identifier
            of this BGP speaker.
          * `peerWeight` (`pulumi.Input[float]`) - The weight added to routes learned from this
            BGP speaker.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if address_spaces is None:
                raise TypeError("Missing required property 'address_spaces'")
            __props__['address_spaces'] = address_spaces
            __props__['bgp_settings'] = bgp_settings
            if gateway_address is None:
                raise TypeError("Missing required property 'gateway_address'")
            __props__['gateway_address'] = gateway_address
            __props__['location'] = location
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
        super(LocalNetworkGateway, __self__).__init__(
            'azure:network/localNetworkGateway:LocalNetworkGateway',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, address_spaces=None, bgp_settings=None, gateway_address=None, location=None, name=None, resource_group_name=None, tags=None):
        """
        Get an existing LocalNetworkGateway resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] address_spaces: The list of string CIDRs representing the
               address spaces the gateway exposes.
        :param pulumi.Input[dict] bgp_settings: A `bgp_settings` block as defined below containing the
               Local Network Gateway's BGP speaker settings.
        :param pulumi.Input[str] gateway_address: The IP address of the gateway to which to
               connect.
        :param pulumi.Input[str] location: The location/region where the local network gateway is
               created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the local network gateway. Changing this
               forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to
               create the local network gateway.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.

        The **bgp_settings** object supports the following:

          * `asn` (`pulumi.Input[float]`) - The BGP speaker's ASN.
          * `bgpPeeringAddress` (`pulumi.Input[str]`) - The BGP peering address and BGP identifier
            of this BGP speaker.
          * `peerWeight` (`pulumi.Input[float]`) - The weight added to routes learned from this
            BGP speaker.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["address_spaces"] = address_spaces
        __props__["bgp_settings"] = bgp_settings
        __props__["gateway_address"] = gateway_address
        __props__["location"] = location
        __props__["name"] = name
        __props__["resource_group_name"] = resource_group_name
        __props__["tags"] = tags
        return LocalNetworkGateway(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

