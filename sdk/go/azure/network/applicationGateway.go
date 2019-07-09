// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an Application Gateway.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/application_gateway.html.markdown.
type ApplicationGateway struct {
	s *pulumi.ResourceState
}

// NewApplicationGateway registers a new resource with the given unique name, arguments, and options.
func NewApplicationGateway(ctx *pulumi.Context,
	name string, args *ApplicationGatewayArgs, opts ...pulumi.ResourceOpt) (*ApplicationGateway, error) {
	if args == nil || args.BackendAddressPools == nil {
		return nil, errors.New("missing required argument 'BackendAddressPools'")
	}
	if args == nil || args.BackendHttpSettings == nil {
		return nil, errors.New("missing required argument 'BackendHttpSettings'")
	}
	if args == nil || args.FrontendIpConfigurations == nil {
		return nil, errors.New("missing required argument 'FrontendIpConfigurations'")
	}
	if args == nil || args.FrontendPorts == nil {
		return nil, errors.New("missing required argument 'FrontendPorts'")
	}
	if args == nil || args.GatewayIpConfigurations == nil {
		return nil, errors.New("missing required argument 'GatewayIpConfigurations'")
	}
	if args == nil || args.HttpListeners == nil {
		return nil, errors.New("missing required argument 'HttpListeners'")
	}
	if args == nil || args.RequestRoutingRules == nil {
		return nil, errors.New("missing required argument 'RequestRoutingRules'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Sku == nil {
		return nil, errors.New("missing required argument 'Sku'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["authenticationCertificates"] = nil
		inputs["autoscaleConfiguration"] = nil
		inputs["backendAddressPools"] = nil
		inputs["backendHttpSettings"] = nil
		inputs["customErrorConfigurations"] = nil
		inputs["disabledSslProtocols"] = nil
		inputs["enableHttp2"] = nil
		inputs["frontendIpConfigurations"] = nil
		inputs["frontendPorts"] = nil
		inputs["gatewayIpConfigurations"] = nil
		inputs["httpListeners"] = nil
		inputs["location"] = nil
		inputs["name"] = nil
		inputs["probes"] = nil
		inputs["redirectConfigurations"] = nil
		inputs["requestRoutingRules"] = nil
		inputs["resourceGroupName"] = nil
		inputs["rewriteRuleSets"] = nil
		inputs["sku"] = nil
		inputs["sslCertificates"] = nil
		inputs["sslPolicies"] = nil
		inputs["tags"] = nil
		inputs["urlPathMaps"] = nil
		inputs["wafConfiguration"] = nil
		inputs["zones"] = nil
	} else {
		inputs["authenticationCertificates"] = args.AuthenticationCertificates
		inputs["autoscaleConfiguration"] = args.AutoscaleConfiguration
		inputs["backendAddressPools"] = args.BackendAddressPools
		inputs["backendHttpSettings"] = args.BackendHttpSettings
		inputs["customErrorConfigurations"] = args.CustomErrorConfigurations
		inputs["disabledSslProtocols"] = args.DisabledSslProtocols
		inputs["enableHttp2"] = args.EnableHttp2
		inputs["frontendIpConfigurations"] = args.FrontendIpConfigurations
		inputs["frontendPorts"] = args.FrontendPorts
		inputs["gatewayIpConfigurations"] = args.GatewayIpConfigurations
		inputs["httpListeners"] = args.HttpListeners
		inputs["location"] = args.Location
		inputs["name"] = args.Name
		inputs["probes"] = args.Probes
		inputs["redirectConfigurations"] = args.RedirectConfigurations
		inputs["requestRoutingRules"] = args.RequestRoutingRules
		inputs["resourceGroupName"] = args.ResourceGroupName
		inputs["rewriteRuleSets"] = args.RewriteRuleSets
		inputs["sku"] = args.Sku
		inputs["sslCertificates"] = args.SslCertificates
		inputs["sslPolicies"] = args.SslPolicies
		inputs["tags"] = args.Tags
		inputs["urlPathMaps"] = args.UrlPathMaps
		inputs["wafConfiguration"] = args.WafConfiguration
		inputs["zones"] = args.Zones
	}
	s, err := ctx.RegisterResource("azure:network/applicationGateway:ApplicationGateway", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ApplicationGateway{s: s}, nil
}

// GetApplicationGateway gets an existing ApplicationGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplicationGateway(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ApplicationGatewayState, opts ...pulumi.ResourceOpt) (*ApplicationGateway, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["authenticationCertificates"] = state.AuthenticationCertificates
		inputs["autoscaleConfiguration"] = state.AutoscaleConfiguration
		inputs["backendAddressPools"] = state.BackendAddressPools
		inputs["backendHttpSettings"] = state.BackendHttpSettings
		inputs["customErrorConfigurations"] = state.CustomErrorConfigurations
		inputs["disabledSslProtocols"] = state.DisabledSslProtocols
		inputs["enableHttp2"] = state.EnableHttp2
		inputs["frontendIpConfigurations"] = state.FrontendIpConfigurations
		inputs["frontendPorts"] = state.FrontendPorts
		inputs["gatewayIpConfigurations"] = state.GatewayIpConfigurations
		inputs["httpListeners"] = state.HttpListeners
		inputs["location"] = state.Location
		inputs["name"] = state.Name
		inputs["probes"] = state.Probes
		inputs["redirectConfigurations"] = state.RedirectConfigurations
		inputs["requestRoutingRules"] = state.RequestRoutingRules
		inputs["resourceGroupName"] = state.ResourceGroupName
		inputs["rewriteRuleSets"] = state.RewriteRuleSets
		inputs["sku"] = state.Sku
		inputs["sslCertificates"] = state.SslCertificates
		inputs["sslPolicies"] = state.SslPolicies
		inputs["tags"] = state.Tags
		inputs["urlPathMaps"] = state.UrlPathMaps
		inputs["wafConfiguration"] = state.WafConfiguration
		inputs["zones"] = state.Zones
	}
	s, err := ctx.ReadResource("azure:network/applicationGateway:ApplicationGateway", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ApplicationGateway{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *ApplicationGateway) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *ApplicationGateway) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// One or more `authentication_certificate` blocks as defined below.
func (r *ApplicationGateway) AuthenticationCertificates() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["authenticationCertificates"])
}

// A `autoscale_configuration` block as defined below.
func (r *ApplicationGateway) AutoscaleConfiguration() *pulumi.Output {
	return r.s.State["autoscaleConfiguration"]
}

// One or more `backend_address_pool` blocks as defined below.
func (r *ApplicationGateway) BackendAddressPools() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["backendAddressPools"])
}

// One or more `backend_http_settings` blocks as defined below.
func (r *ApplicationGateway) BackendHttpSettings() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["backendHttpSettings"])
}

// One or more `custom_error_configuration` blocks as defined below.
func (r *ApplicationGateway) CustomErrorConfigurations() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["customErrorConfigurations"])
}

// A list of SSL Protocols which should be disabled on this Application Gateway. Possible values are `TLSv1_0`, `TLSv1_1` and `TLSv1_2`.
// > **NOTE:** `disabled_ssl_protocols ` has been deprecated in favour of `disabled_protocols` in the `ssl_policy` block.
func (r *ApplicationGateway) DisabledSslProtocols() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["disabledSslProtocols"])
}

// Is HTTP2 enabled on the application gateway resource? Defaults to `false`.
func (r *ApplicationGateway) EnableHttp2() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["enableHttp2"])
}

// One or more `frontend_ip_configuration` blocks as defined below.
func (r *ApplicationGateway) FrontendIpConfigurations() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["frontendIpConfigurations"])
}

// One or more `frontend_port` blocks as defined below.
func (r *ApplicationGateway) FrontendPorts() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["frontendPorts"])
}

// One or more `gateway_ip_configuration` blocks as defined below.
func (r *ApplicationGateway) GatewayIpConfigurations() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["gatewayIpConfigurations"])
}

// One or more `http_listener` blocks as defined below.
func (r *ApplicationGateway) HttpListeners() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["httpListeners"])
}

// The Azure region where the Application Gateway should exist. Changing this forces a new resource to be created.
func (r *ApplicationGateway) Location() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["location"])
}

// The name of the Application Gateway. Changing this forces a new resource to be created.
func (r *ApplicationGateway) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// One or more `probe` blocks as defined below.
func (r *ApplicationGateway) Probes() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["probes"])
}

// A `redirect_configuration` block as defined below.
func (r *ApplicationGateway) RedirectConfigurations() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["redirectConfigurations"])
}

// One or more `request_routing_rule` blocks as defined below.
func (r *ApplicationGateway) RequestRoutingRules() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["requestRoutingRules"])
}

// The name of the resource group in which to the Application Gateway should exist. Changing this forces a new resource to be created.
func (r *ApplicationGateway) ResourceGroupName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// One or more `rewrite_rule_set` blocks as defined below. Only valid for v2 SKUs.
func (r *ApplicationGateway) RewriteRuleSets() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["rewriteRuleSets"])
}

// A `sku` block as defined below.
func (r *ApplicationGateway) Sku() *pulumi.Output {
	return r.s.State["sku"]
}

// One or more `ssl_certificate` blocks as defined below.
func (r *ApplicationGateway) SslCertificates() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["sslCertificates"])
}

// a `ssl policy` block as defined below.
func (r *ApplicationGateway) SslPolicies() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["sslPolicies"])
}

// A mapping of tags to assign to the resource.
func (r *ApplicationGateway) Tags() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["tags"])
}

// One or more `url_path_map` blocks as defined below.
func (r *ApplicationGateway) UrlPathMaps() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["urlPathMaps"])
}

// A `waf_configuration` block as defined below.
func (r *ApplicationGateway) WafConfiguration() *pulumi.Output {
	return r.s.State["wafConfiguration"]
}

// A collection of availability zones to spread the Application Gateway over.
func (r *ApplicationGateway) Zones() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["zones"])
}

// Input properties used for looking up and filtering ApplicationGateway resources.
type ApplicationGatewayState struct {
	// One or more `authentication_certificate` blocks as defined below.
	AuthenticationCertificates interface{}
	// A `autoscale_configuration` block as defined below.
	AutoscaleConfiguration interface{}
	// One or more `backend_address_pool` blocks as defined below.
	BackendAddressPools interface{}
	// One or more `backend_http_settings` blocks as defined below.
	BackendHttpSettings interface{}
	// One or more `custom_error_configuration` blocks as defined below.
	CustomErrorConfigurations interface{}
	// A list of SSL Protocols which should be disabled on this Application Gateway. Possible values are `TLSv1_0`, `TLSv1_1` and `TLSv1_2`.
	// > **NOTE:** `disabled_ssl_protocols ` has been deprecated in favour of `disabled_protocols` in the `ssl_policy` block.
	DisabledSslProtocols interface{}
	// Is HTTP2 enabled on the application gateway resource? Defaults to `false`.
	EnableHttp2 interface{}
	// One or more `frontend_ip_configuration` blocks as defined below.
	FrontendIpConfigurations interface{}
	// One or more `frontend_port` blocks as defined below.
	FrontendPorts interface{}
	// One or more `gateway_ip_configuration` blocks as defined below.
	GatewayIpConfigurations interface{}
	// One or more `http_listener` blocks as defined below.
	HttpListeners interface{}
	// The Azure region where the Application Gateway should exist. Changing this forces a new resource to be created.
	Location interface{}
	// The name of the Application Gateway. Changing this forces a new resource to be created.
	Name interface{}
	// One or more `probe` blocks as defined below.
	Probes interface{}
	// A `redirect_configuration` block as defined below.
	RedirectConfigurations interface{}
	// One or more `request_routing_rule` blocks as defined below.
	RequestRoutingRules interface{}
	// The name of the resource group in which to the Application Gateway should exist. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// One or more `rewrite_rule_set` blocks as defined below. Only valid for v2 SKUs.
	RewriteRuleSets interface{}
	// A `sku` block as defined below.
	Sku interface{}
	// One or more `ssl_certificate` blocks as defined below.
	SslCertificates interface{}
	// a `ssl policy` block as defined below.
	SslPolicies interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// One or more `url_path_map` blocks as defined below.
	UrlPathMaps interface{}
	// A `waf_configuration` block as defined below.
	WafConfiguration interface{}
	// A collection of availability zones to spread the Application Gateway over.
	Zones interface{}
}

// The set of arguments for constructing a ApplicationGateway resource.
type ApplicationGatewayArgs struct {
	// One or more `authentication_certificate` blocks as defined below.
	AuthenticationCertificates interface{}
	// A `autoscale_configuration` block as defined below.
	AutoscaleConfiguration interface{}
	// One or more `backend_address_pool` blocks as defined below.
	BackendAddressPools interface{}
	// One or more `backend_http_settings` blocks as defined below.
	BackendHttpSettings interface{}
	// One or more `custom_error_configuration` blocks as defined below.
	CustomErrorConfigurations interface{}
	// A list of SSL Protocols which should be disabled on this Application Gateway. Possible values are `TLSv1_0`, `TLSv1_1` and `TLSv1_2`.
	// > **NOTE:** `disabled_ssl_protocols ` has been deprecated in favour of `disabled_protocols` in the `ssl_policy` block.
	DisabledSslProtocols interface{}
	// Is HTTP2 enabled on the application gateway resource? Defaults to `false`.
	EnableHttp2 interface{}
	// One or more `frontend_ip_configuration` blocks as defined below.
	FrontendIpConfigurations interface{}
	// One or more `frontend_port` blocks as defined below.
	FrontendPorts interface{}
	// One or more `gateway_ip_configuration` blocks as defined below.
	GatewayIpConfigurations interface{}
	// One or more `http_listener` blocks as defined below.
	HttpListeners interface{}
	// The Azure region where the Application Gateway should exist. Changing this forces a new resource to be created.
	Location interface{}
	// The name of the Application Gateway. Changing this forces a new resource to be created.
	Name interface{}
	// One or more `probe` blocks as defined below.
	Probes interface{}
	// A `redirect_configuration` block as defined below.
	RedirectConfigurations interface{}
	// One or more `request_routing_rule` blocks as defined below.
	RequestRoutingRules interface{}
	// The name of the resource group in which to the Application Gateway should exist. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// One or more `rewrite_rule_set` blocks as defined below. Only valid for v2 SKUs.
	RewriteRuleSets interface{}
	// A `sku` block as defined below.
	Sku interface{}
	// One or more `ssl_certificate` blocks as defined below.
	SslCertificates interface{}
	// a `ssl policy` block as defined below.
	SslPolicies interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// One or more `url_path_map` blocks as defined below.
	UrlPathMaps interface{}
	// A `waf_configuration` block as defined below.
	WafConfiguration interface{}
	// A collection of availability zones to spread the Application Gateway over.
	Zones interface{}
}
