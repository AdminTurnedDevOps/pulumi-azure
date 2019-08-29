# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Schedule(pulumi.CustomResource):
    daily_recurrence: pulumi.Output[dict]
    hourly_recurrence: pulumi.Output[dict]
    lab_name: pulumi.Output[str]
    location: pulumi.Output[str]
    name: pulumi.Output[str]
    notification_settings: pulumi.Output[dict]
    resource_group_name: pulumi.Output[str]
    status: pulumi.Output[str]
    tags: pulumi.Output[dict]
    task_type: pulumi.Output[str]
    time_zone_id: pulumi.Output[str]
    weekly_recurrence: pulumi.Output[dict]
    def __init__(__self__, resource_name, opts=None, daily_recurrence=None, hourly_recurrence=None, lab_name=None, location=None, name=None, notification_settings=None, resource_group_name=None, status=None, tags=None, task_type=None, time_zone_id=None, weekly_recurrence=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a Schedule resource with the given unique name, props, and options.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        
        The **daily_recurrence** object supports the following:
        
          * `time` (`pulumi.Input[str]`)
        
        The **hourly_recurrence** object supports the following:
        
          * `minute` (`pulumi.Input[float]`)
        
        The **notification_settings** object supports the following:
        
          * `status` (`pulumi.Input[str]`)
          * `timeInMinutes` (`pulumi.Input[float]`)
          * `webhookUrl` (`pulumi.Input[str]`)
        
        The **weekly_recurrence** object supports the following:
        
          * `time` (`pulumi.Input[str]`)
          * `week_days` (`pulumi.Input[list]`)
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

            __props__['daily_recurrence'] = daily_recurrence
            __props__['hourly_recurrence'] = hourly_recurrence
            if lab_name is None:
                raise TypeError("Missing required property 'lab_name'")
            __props__['lab_name'] = lab_name
            __props__['location'] = location
            __props__['name'] = name
            if notification_settings is None:
                raise TypeError("Missing required property 'notification_settings'")
            __props__['notification_settings'] = notification_settings
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['status'] = status
            __props__['tags'] = tags
            if task_type is None:
                raise TypeError("Missing required property 'task_type'")
            __props__['task_type'] = task_type
            if time_zone_id is None:
                raise TypeError("Missing required property 'time_zone_id'")
            __props__['time_zone_id'] = time_zone_id
            __props__['weekly_recurrence'] = weekly_recurrence
        super(Schedule, __self__).__init__(
            'azure:devtest/schedule:Schedule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, daily_recurrence=None, hourly_recurrence=None, lab_name=None, location=None, name=None, notification_settings=None, resource_group_name=None, status=None, tags=None, task_type=None, time_zone_id=None, weekly_recurrence=None):
        """
        Get an existing Schedule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        
        The **daily_recurrence** object supports the following:
        
          * `time` (`pulumi.Input[str]`)
        
        The **hourly_recurrence** object supports the following:
        
          * `minute` (`pulumi.Input[float]`)
        
        The **notification_settings** object supports the following:
        
          * `status` (`pulumi.Input[str]`)
          * `timeInMinutes` (`pulumi.Input[float]`)
          * `webhookUrl` (`pulumi.Input[str]`)
        
        The **weekly_recurrence** object supports the following:
        
          * `time` (`pulumi.Input[str]`)
          * `week_days` (`pulumi.Input[list]`)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["daily_recurrence"] = daily_recurrence
        __props__["hourly_recurrence"] = hourly_recurrence
        __props__["lab_name"] = lab_name
        __props__["location"] = location
        __props__["name"] = name
        __props__["notification_settings"] = notification_settings
        __props__["resource_group_name"] = resource_group_name
        __props__["status"] = status
        __props__["tags"] = tags
        __props__["task_type"] = task_type
        __props__["time_zone_id"] = time_zone_id
        __props__["weekly_recurrence"] = weekly_recurrence
        return Schedule(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
