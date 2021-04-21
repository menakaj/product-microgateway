/*
 *  Copyright (c) 2020, WSO2 Inc. (http://www.wso2.org) All Rights Reserved.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

package model

import (
	"errors"
	"net/url"
	"regexp"
	"strings"

	parser "github.com/mitchellh/mapstructure"
	"github.com/wso2/adapter/config"
	"github.com/wso2/adapter/internal/svcdiscovery"
	logger "github.com/wso2/adapter/loggers"
)

// MgwSwagger represents the object structure holding the information related to the
// openAPI object. The values are populated from the extensions/properties mentioned at
// the root level of the openAPI definition. The pathItem level information is represented
// by the resources array which contains the MgwResource entries.
type MgwSwagger struct {
	id                  string
	apiType             string
	description         string
	title               string
	version             string
	vendorExtensions    map[string]interface{}
	productionUrls      []Endpoint //
	sandboxUrls         []Endpoint
	resources           []Resource
	xWso2Basepath       string
	xWso2Cors           *CorsConfig
	securityScheme      []string
	xWso2ThrottlingTier string
	xWso2AuthHeader     string
	disableSecurity     bool
	OrganizationID      string
}

// EndpointSecurity contains the SandBox/Production endpoint security
type EndpointSecurity struct {
	SandBox    SecurityInfo
	Production SecurityInfo
}

// SecurityInfo contains the parameters of endpoint security
type SecurityInfo struct {
	Password         string
	CustomParameters string
	SecurityType     string
	Enabled          bool
	Username         string
}

// Endpoint represents the structure of an endpoint.
type Endpoint struct {
	// Host name
	Host string
	// BasePath (which would be added as prefix to the path mentioned in openapi definition)
	// In openAPI v2, it is determined from the basePath property
	// In openAPi v3, it is determined from the server object's suffix
	Basepath string
	// https, http, ws, wss
	// In openAPI v2, it is fetched from the schemes entry
	// In openAPI v3, it is extracted from the server property under servers object
	// only https and http are supported at the moment.
	URLType string
	// Port of the endpoint.
	// If the port is not specified, 80 is assigned if URLType is http
	// 443 is assigned if URLType is https
	Port uint32
	//ServiceDiscoveryQuery consul query for service discovery
	ServiceDiscoveryString string
}

// CorsConfig represents the API level Cors Configuration
type CorsConfig struct {
	Enabled                       bool     `mapstructure:"corsConfigurationEnabled"`
	AccessControlAllowCredentials bool     `mapstructure:"accessControlAllowCredentials,omitempty"`
	AccessControlAllowHeaders     []string `mapstructure:"accessControlAllowHeaders"`
	AccessControlAllowMethods     []string `mapstructure:"accessControlAllowMethods"`
	AccessControlAllowOrigins     []string `mapstructure:"accessControlAllowOrigins"`
	AccessControlExposeHeaders    []string `mapstructure:"accessControlExposeHeaders"`
}

// GetCorsConfig returns the CorsConfiguration Object.
func (swagger *MgwSwagger) GetCorsConfig() *CorsConfig {
	return swagger.xWso2Cors
}

// GetAPIType returns the openapi version
func (swagger *MgwSwagger) GetAPIType() string {
	return swagger.apiType
}

// GetVersion returns the API version
func (swagger *MgwSwagger) GetVersion() string {
	return swagger.version
}

// GetTitle returns the API Title
func (swagger *MgwSwagger) GetTitle() string {
	return swagger.title
}

// GetXWso2Basepath returns the basepath set via the vendor extension.
func (swagger *MgwSwagger) GetXWso2Basepath() string {
	return swagger.xWso2Basepath
}

// GetVendorExtensions returns the map of vendor extensions which are defined
// at openAPI's root level.
func (swagger *MgwSwagger) GetVendorExtensions() map[string]interface{} {
	return swagger.vendorExtensions
}

// GetProdEndpoints returns the array of production endpoints.
func (swagger *MgwSwagger) GetProdEndpoints() []Endpoint {
	return swagger.productionUrls
}

// GetSandEndpoints returns the array of sandbox endpoints.
func (swagger *MgwSwagger) GetSandEndpoints() []Endpoint {
	return swagger.sandboxUrls
}

// GetResources returns the array of resources (openAPI path level info)
func (swagger *MgwSwagger) GetResources() []Resource {
	return swagger.resources
}

// GetDescription returns the description of the openapi
func (swagger *MgwSwagger) GetDescription() string {
	return swagger.description
}

// GetXWso2ThrottlingTier returns the Throttling tier via the vendor extension.
func (swagger *MgwSwagger) GetXWso2ThrottlingTier() string {
	return swagger.xWso2ThrottlingTier
}

// GetDisableSecurity returns the authType via the vendor extension.
func (swagger *MgwSwagger) GetDisableSecurity() bool {
	return swagger.disableSecurity
}

// GetID returns the Id of the API
func (swagger *MgwSwagger) GetID() string {
	return swagger.id
}

// SetID set the Id of the API
func (swagger *MgwSwagger) SetID(id string) {
	swagger.id = id
}

// SetName sets the name of the API
func (swagger *MgwSwagger) SetName(name string) {
	swagger.title = name
}

// SetSecurityScheme sets the securityscheme of the API
func (swagger *MgwSwagger) SetSecurityScheme(securityScheme []string) {
	swagger.securityScheme = securityScheme
}

// SetVersion sets the version of the API
func (swagger *MgwSwagger) SetVersion(version string) {
	swagger.version = version
}

// SetXWso2AuthHeader sets the authHeader of the API
func (swagger *MgwSwagger) SetXWso2AuthHeader(authHeader string) {
	if swagger.xWso2AuthHeader == "" {
		swagger.xWso2AuthHeader = authHeader
	}
}

// GetXWSO2AuthHeader returns the auth header set via the vendor extension.
func (swagger *MgwSwagger) GetXWSO2AuthHeader() string {
	return swagger.xWso2AuthHeader
}

// GetSetSecurityScheme returns the securityscheme of the API
func (swagger *MgwSwagger) GetSetSecurityScheme() []string {
	return swagger.securityScheme
}

// SetXWso2Extenstions set the MgwSwagger object with the properties
// extracted from vendor extensions.
// xWso2Basepath, xWso2ProductionEndpoints, and xWso2SandboxEndpoints are assigned
// based on the vendor extensions.
//
// Resource level properties (xwso2ProductionEndpoints and xWso2SandboxEndpoints are
// also populated at the same time).
func (swagger *MgwSwagger) SetXWso2Extenstions() {
	swagger.setXWso2Basepath()
	swagger.setXWso2PrdoductionEndpoint()
	swagger.setXWso2SandboxEndpoint()
	swagger.setXWso2Cors()
	swagger.setXWso2ThrottlingTier()
	swagger.setDisableSecurity()
	swagger.setXWso2AuthHeader()
}

// SetXWso2SandboxEndpointForMgwSwagger set the MgwSwagger object with the SandboxEndpoint when
// it is not populated by SetXWso2Extenstions
func (swagger *MgwSwagger) SetXWso2SandboxEndpointForMgwSwagger(sandBoxURL string) {
	var sandboxEndpoints []Endpoint
	sandboxEndpoints = append(sandboxEndpoints, getHostandBasepathandPort(sandBoxURL))
	swagger.sandboxUrls = sandboxEndpoints
}

// SetXWso2ProductionEndpointMgwSwagger set the MgwSwagger object with the productionEndpoint when
// it is not populated by SetXWso2Extenstions
func (swagger *MgwSwagger) SetXWso2ProductionEndpointMgwSwagger(productionURL string) {
	var productionEndpoints []Endpoint
	productionEndpoints = append(productionEndpoints, getHostandBasepathandPort(productionURL))
	swagger.productionUrls = productionEndpoints
}

func (swagger *MgwSwagger) setXWso2PrdoductionEndpoint() {
	xWso2APIEndpoints := getXWso2Endpoints(swagger.vendorExtensions, productionEndpoints)
	if xWso2APIEndpoints != nil && len(xWso2APIEndpoints) > 0 {
		swagger.productionUrls = xWso2APIEndpoints
	}

	//resources
	for i, resource := range swagger.resources {
		xwso2ResourceEndpoints := getXWso2Endpoints(resource.vendorExtensions, productionEndpoints)
		if xwso2ResourceEndpoints != nil {
			swagger.resources[i].productionUrls = xwso2ResourceEndpoints
		}
	}
}

func (swagger *MgwSwagger) setXWso2SandboxEndpoint() {
	xWso2APIEndpoints := getXWso2Endpoints(swagger.vendorExtensions, sandboxEndpoints)
	if xWso2APIEndpoints != nil && len(xWso2APIEndpoints) > 0 {
		swagger.sandboxUrls = xWso2APIEndpoints
	}

	// resources
	for i, resource := range swagger.resources {
		xwso2ResourceEndpoints := getXWso2Endpoints(resource.vendorExtensions, sandboxEndpoints)
		if xwso2ResourceEndpoints != nil {
			swagger.resources[i].sandboxUrls = xwso2ResourceEndpoints
		}
	}
}

func (swagger *MgwSwagger) setXWso2ThrottlingTier() {
	tier := ResolveThrottlingTier(swagger.vendorExtensions)
	if tier != "" {
		swagger.xWso2ThrottlingTier = tier
	}
}

// getXWso2AuthHeader extracts the value of xWso2AuthHeader extension.
// if the property is not available, an empty string is returned.
func getXWso2AuthHeader(vendorExtensions map[string]interface{}) string {
	xWso2AuthHeader := ""
	if y, found := vendorExtensions[xAuthHeader]; found {
		if val, ok := y.(string); ok {
			xWso2AuthHeader = val
		}
	}
	return xWso2AuthHeader
}

// SetXWSO2AuthHeader sets the AuthHeader of the API
func (swagger *MgwSwagger) setXWso2AuthHeader() {
	authorizationHeader := getXWso2AuthHeader(swagger.vendorExtensions)
	if authorizationHeader != "" {
		swagger.xWso2AuthHeader = authorizationHeader
	}
}

func (swagger *MgwSwagger) setDisableSecurity() {
	swagger.disableSecurity = ResolveDisableSecurity(swagger.vendorExtensions)
}

// Validate method confirms that the mgwSwagger has all required fields in the required format.
// This needs to be checked prior to generate router/enforcer related resources.
func (swagger *MgwSwagger) Validate() error {
	versionMatcher := regexp.MustCompile(`^[a-zA-Z0-9_.-]*$`)
	if !versionMatcher.Match([]byte(swagger.version)) {
		logger.LoggerOasparser.Errorf("API version should be in proper format '^[a-zA-Z0-9_.-]*$' %s:%s",
			swagger.title, swagger.version)
		return errors.New("API Version should be in proper format '^[a-zA-Z0-9_.-]*$'")
	}
	if len(swagger.productionUrls) == 0 && len(swagger.sandboxUrls) == 0 {
		logger.LoggerOasparser.Errorf("No Endpoints are provided for the API %s:%s",
			swagger.title, swagger.version)
		return errors.New("No Endpoints are provided for the API")
	}
	if len(swagger.productionUrls) > 0 {
		err := swagger.productionUrls[0].validateEndpoint()
		if err != nil {
			logger.LoggerOasparser.Errorf("Error while parsing the production endpoints of the API %s:%s - %v",
				swagger.title, swagger.version, err)
			return err
		}
	}

	if len(swagger.sandboxUrls) > 0 {
		err := swagger.sandboxUrls[0].validateEndpoint()
		if err != nil {
			logger.LoggerOasparser.Errorf("Error while parsing the sandbox endpoints of the API %s:%s - %v",
				swagger.title, swagger.version, err)
			return err
		}
	}
	return nil
}

func (endpoint *Endpoint) validateEndpoint() error {
	if len(endpoint.ServiceDiscoveryString) > 0 {
		return nil
	}
	if len(endpoint.Host) == 0 {
		return errors.New("Empty Hostname is provided")
	}
	if endpoint.Port >= 65535 {
		return errors.New("Port value should be less than 65535")
	}
	if strings.HasPrefix(endpoint.Host, "/") {
		return errors.New("Relative paths are not supported as endpoint URLs")
	}
	urlString := endpoint.URLType + "://" + endpoint.Host
	_, err := url.ParseRequestURI(urlString)
	return err
}

// getXWso2Endpoints extracts and generate the Endpoint Objects from the vendor extension map.
func getXWso2Endpoints(vendorExtensions map[string]interface{}, endpointType string) []Endpoint {
	var endpoints []Endpoint

	// TODO: (VirajSalaka) x-wso2-production-endpoint 's type does not represent http/https, instead it indicates loadbalance and failover
	if y, found := vendorExtensions[endpointType]; found {
		if val, ok := y.(map[string]interface{}); ok {
			urlsProperty, ok := val[urls]
			if !ok {
				// TODO: (VirajSalaka) Throw an error and catch from an upper layer where the API name is visible.
				logger.LoggerOasparser.Error("urls property is not provided with the x-wso2-production-endpoints/" +
					"x-wso2-sandbox-endpoints extension.")
			} else {
				castedUrlsInterface := urlsProperty.([]interface{})
				for _, v := range castedUrlsInterface {
					if svcdiscovery.IsServiceDiscoveryEnabled && svcdiscovery.IsDiscoveryServiceEndpoint(v.(string)) {
						logger.LoggerOasparser.Debug("consul query syntax found: ", v.(string))
						queryString, defHost, err := svcdiscovery.ParseConsulSyntax(v.(string))
						if err != nil {
							logger.LoggerOasparser.Error("consul syntax parse error ", err)
							continue
						}
						endpoint := getHostandBasepathandPort(defHost)
						endpoint.ServiceDiscoveryString = queryString
						endpoints = append(endpoints, endpoint)

					} else {
						endpoint := getHostandBasepathandPort(v.(string))
						endpoints = append(endpoints, endpoint)
					}

				}
				return endpoints
			}
		} else {
			logger.LoggerOasparser.Fatal("x-wso2-production/sandbox-endpoints is not having a correct map structure")
		}
	}
	return nil
}

// getXWso2Basepath extracts the value of xWso2BasePath extension.
// if the property is not available, an empty string is returned.
func getXWso2Basepath(vendorExtensions map[string]interface{}) string {
	xWso2basepath := ""
	if y, found := vendorExtensions[xWso2BasePath]; found {
		if val, ok := y.(string); ok {
			xWso2basepath = val
		}
	}
	return xWso2basepath
}

func (swagger *MgwSwagger) setXWso2Basepath() {
	extBasepath := getXWso2Basepath(swagger.vendorExtensions)
	if extBasepath != "" {
		swagger.xWso2Basepath = extBasepath
	}
}

func (swagger *MgwSwagger) setXWso2Cors() {
	if cors, corsFound := swagger.vendorExtensions[xWso2Cors]; corsFound {
		logger.LoggerOasparser.Debugf("%v configuration is available", xWso2Cors)
		if parsedCors, parsedCorsOk := cors.(map[string]interface{}); parsedCorsOk {
			//Default CorsConfiguration
			corsConfig := &CorsConfig{
				Enabled: true,
			}
			err := parser.Decode(parsedCors, &corsConfig)
			if err != nil {
				logger.LoggerOasparser.Errorf("Error while parsing %v: "+err.Error(), xWso2Cors)
				return
			}
			if corsConfig.Enabled {
				logger.LoggerOasparser.Debugf("API Level Cors Configuration is applied : %+v\n", corsConfig)
				swagger.xWso2Cors = corsConfig
				return
			}
			swagger.xWso2Cors = generateGlobalCors()
			return
		}
		logger.LoggerOasparser.Errorf("Error while parsing %v .", xWso2Cors)
	} else {
		swagger.xWso2Cors = generateGlobalCors()
	}
}

func generateGlobalCors() *CorsConfig {
	conf, _ := config.ReadConfigs()
	logger.LoggerOasparser.Debug("CORS policy is applied from global configuration.")
	return &CorsConfig{
		Enabled:                       conf.Envoy.Cors.Enabled,
		AccessControlAllowCredentials: conf.Envoy.Cors.AllowCredentials,
		AccessControlAllowOrigins:     conf.Envoy.Cors.AllowOrigins,
		AccessControlAllowHeaders:     conf.Envoy.Cors.AllowHeaders,
		AccessControlAllowMethods:     conf.Envoy.Cors.AllowMethods,
		AccessControlExposeHeaders:    conf.Envoy.Cors.ExposeHeaders,
	}
}

// ResolveThrottlingTier extracts the value of x-wso2-throttling-tier and
// x-throttling-tier extension. if x-wso2-throttling-tier is available it
// will be prioritized.
// if both the properties are not available, an empty string is returned.
func ResolveThrottlingTier(vendorExtensions map[string]interface{}) string {
	xTier := ""
	if x, found := vendorExtensions[xWso2ThrottlingTier]; found {
		if val, ok := x.(string); ok {
			xTier = val
		}
	} else if y, found := vendorExtensions[xThrottlingTier]; found {
		if val, ok := y.(string); ok {
			xTier = val
		}
	}
	return xTier
}

// ResolveDisableSecurity extracts the value of x-auth-type extension.
// if the property is not available, false is returned.
// If the API definition is fed from API manager, then API definition contains
// x-auth-type as "None" for non secured APIs. Then the return value would be true.
// If the API definition is fed through apictl, the users can use either
// x-wso2-disable-security : true/false to enable and disable security.
func ResolveDisableSecurity(vendorExtensions map[string]interface{}) bool {
	disableSecurity := false
	y, vExtAuthType := vendorExtensions[xAuthType]
	z, vExtDisableSecurity := vendorExtensions[xWso2DisableSecurity]
	if vExtDisableSecurity {
		// If x-wso2-disable-security is present, then disableSecurity = val
		if val, ok := z.(bool); ok {
			disableSecurity = val
		}
	} else if vExtAuthType {
		// If APIs are published through APIM, all resource levels contains x-auth-type
		// vendor extension.
		if val, ok := y.(string); ok {
			// If the x-auth-type vendor ext is None, then the API/resource is considerd
			// to be non secure
			if val == None {
				disableSecurity = true
			}
		}
	}
	return disableSecurity
}
