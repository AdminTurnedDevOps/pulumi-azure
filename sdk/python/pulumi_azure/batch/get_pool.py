# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetPoolResult:
    """
    A collection of values returned by getPool.
    """
    def __init__(__self__, account_name=None, auto_scales=None, display_name=None, fixed_scales=None, max_tasks_per_node=None, name=None, node_agent_sku_id=None, resource_group_name=None, start_task=None, storage_image_references=None, vm_size=None, id=None):
        if account_name and not isinstance(account_name, str):
            raise TypeError('Expected argument account_name to be a str')
        __self__.account_name = account_name
        """
        The name of the Batch account.
        """
        if auto_scales and not isinstance(auto_scales, list):
            raise TypeError('Expected argument auto_scales to be a list')
        __self__.auto_scales = auto_scales
        """
        A `auto_scale` block that describes the scale settings when using auto scale.
        """
        if display_name and not isinstance(display_name, str):
            raise TypeError('Expected argument display_name to be a str')
        __self__.display_name = display_name
        if fixed_scales and not isinstance(fixed_scales, list):
            raise TypeError('Expected argument fixed_scales to be a list')
        __self__.fixed_scales = fixed_scales
        """
        A `fixed_scale` block that describes the scale settings when using fixed scale.
        """
        if max_tasks_per_node and not isinstance(max_tasks_per_node, float):
            raise TypeError('Expected argument max_tasks_per_node to be a float')
        __self__.max_tasks_per_node = max_tasks_per_node
        """
        The maximum number of tasks that can run concurrently on a single compute node in the pool.
        """
        if name and not isinstance(name, str):
            raise TypeError('Expected argument name to be a str')
        __self__.name = name
        """
        The name of the Batch pool.
        """
        if node_agent_sku_id and not isinstance(node_agent_sku_id, str):
            raise TypeError('Expected argument node_agent_sku_id to be a str')
        __self__.node_agent_sku_id = node_agent_sku_id
        """
        The Sku of the node agents in the Batch pool.
        """
        if resource_group_name and not isinstance(resource_group_name, str):
            raise TypeError('Expected argument resource_group_name to be a str')
        __self__.resource_group_name = resource_group_name
        if start_task and not isinstance(start_task, dict):
            raise TypeError('Expected argument start_task to be a dict')
        __self__.start_task = start_task
        """
        A `start_task` block that describes the start task settings for the Batch pool.
        """
        if storage_image_references and not isinstance(storage_image_references, list):
            raise TypeError('Expected argument storage_image_references to be a list')
        __self__.storage_image_references = storage_image_references
        """
        The reference of the storage image used by the nodes in the Batch pool.
        """
        if vm_size and not isinstance(vm_size, str):
            raise TypeError('Expected argument vm_size to be a str')
        __self__.vm_size = vm_size
        """
        The size of the VM created in the Batch pool.
        """
        if id and not isinstance(id, str):
            raise TypeError('Expected argument id to be a str')
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

async def get_pool(account_name=None,name=None,resource_group_name=None,start_task=None,opts=None):
    """
    Use this data source to access information about an existing Batch pool
    """
    __args__ = dict()

    __args__['accountName'] = account_name
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    __args__['startTask'] = start_task
    __ret__ = await pulumi.runtime.invoke('azure:batch/getPool:getPool', __args__, opts=opts)

    return GetPoolResult(
        account_name=__ret__.get('accountName'),
        auto_scales=__ret__.get('autoScales'),
        display_name=__ret__.get('displayName'),
        fixed_scales=__ret__.get('fixedScales'),
        max_tasks_per_node=__ret__.get('maxTasksPerNode'),
        name=__ret__.get('name'),
        node_agent_sku_id=__ret__.get('nodeAgentSkuId'),
        resource_group_name=__ret__.get('resourceGroupName'),
        start_task=__ret__.get('startTask'),
        storage_image_references=__ret__.get('storageImageReferences'),
        vm_size=__ret__.get('vmSize'),
        id=__ret__.get('id'))
