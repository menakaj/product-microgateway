// Copyright (c) 2020, WSO2 Inc. (http://www.wso2.org) All Rights Reserved.
//
// WSO2 Inc. licenses this file to you under the Apache License,
// Version 2.0 (the "License"); you may not use this file except
// in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

import ballerina/runtime;

# To retrieve claims via the user specific claim retrieve implementation.
# 
# + userInfo - Authentication Context of the user, which is provided as input to the claim retriever Implementation
# + return - ClaimListDTO if there are any claims added from the user specific implementation
function retrieveClaims (UserClaimRetrieverContextDTO? userInfo) returns @tainted RetrievedUserClaimsListDTO ? {
    //if claim retrieve variable is disabled, there is no need to run through the method.
    if (!claimRetrieverClassLoaded) {
        printDebug (CLAIM_RETRIEVER, "Claim Retriever class is not loaded. Hence remote claim retrieval " +
                    "process is skipped");
        return;
    }
    if (userInfo is UserClaimRetrieverContextDTO) {
        printDebug (CLAIM_RETRIEVER, "User Auth Context information provided to the claim retrieval implementation : " +
                    userInfo.toString());
        RetrievedUserClaimsListDTO? | error claimListDTO = trap retrieveClaimsFromImpl(userInfo);
        if (claimListDTO is RetrievedUserClaimsListDTO ) {
            printDebug (CLAIM_RETRIEVER, "Claims List received from the claim retrieval implementation : " +
                        claimListDTO.toString());
            return claimListDTO;
        } else if (claimListDTO is ()) {
            printDebug(CLAIM_RETRIEVER , "No user claims are received from the claim retrieval implementation");
        } else {
            printError(CLAIM_RETRIEVER , "Error while retrieving user claims from the claim retrieval implementation",
                claimListDTO);
        }
    }
}

# To do the class loading operation for the user specific claim retriever implementation.
# + return - true if claim retriever class loading is successful.
public function loadClaimRetrieverImpl() returns boolean {

    if (!isConfigAvailable(JWT_GENERATOR_CLAIM_RETRIEVAL_INSTANCE_ID, JWT_GENERATOR_CLAIM_RETRIEVAL_IMPLEMENTATION)) {
        printDebug(CLAIM_RETRIEVER, "Claim Retrieval related class loading is disabled as the implementation is not provided." +  
                    "Hence claim retrieval is disabled");
        return false;
    }
    GatewayConf gatewayConf = getGatewayConfInstance();
    string claimRetrieverImplClassName = gatewayConf.jwtGeneratorConfig.claimRetrieval.retrieverImpl;
    map<any> claimRetrieverConfig = gatewayConf.jwtGeneratorConfig.claimRetrieval.configuration;
    string unresolvedTrustStorePath = gatewayConf.listenerConfig.trustStorePath;
    string trustStorePassword = gatewayConf.listenerConfig.trustStorePassword;

    if (!claimRetrieverConfig.hasKey(APIM_CREDENTIALS_USERNAME)) {
        claimRetrieverConfig[APIM_CREDENTIALS_USERNAME] = gatewayConf.apimCredentials.username;
    }
    if (!claimRetrieverConfig.hasKey(APIM_CREDENTIALS_PASSWORD)) {
        claimRetrieverConfig[APIM_CREDENTIALS_PASSWORD] = gatewayConf.apimCredentials.password;
    }
    if (!claimRetrieverConfig.hasKey(KM_SERVER_URL)) {
        claimRetrieverConfig[KM_SERVER_URL] = gatewayConf.getKeyManagerConf().serverUrl;
    }
    return loadClaimRetrieverClass(claimRetrieverImplClassName, unresolvedTrustStorePath, trustStorePassword,
                                    claimRetrieverConfig);
}

# Populate the DTO required for the claim retrieval implementation from authContext and principal component.
# 
# + authContext - Authentication Context
# + principal - Principal component
# + issuer - Issuer related to KeyManager
# + return - populated UserClaimRetrieverContextDTO
function generateUserClaimRetrieverContextFromPrincipal(AuthenticationContext authContext, runtime:Principal principal,
                                            string issuer, boolean isJWT)
        returns UserClaimRetrieverContextDTO {
    UserClaimRetrieverContextDTO userAuthContextDTO = {};
    userAuthContextDTO.username = principal?.username ?: UNKNOWN_VALUE;
    if (isJWT) {
        userAuthContextDTO.token_type = "bearer JWT";
    } else {
        userAuthContextDTO.token_type = "bearer opaque";
    }

    userAuthContextDTO.issuer = issuer;
    userAuthContextDTO.token =  authContext.apiKey;
    map<any>? claims = principal?.claims;
    if (claims is map<any>) {
        userAuthContextDTO.customClaims = claims;
        //cannot use authContext here as the clientId is not populated if subscription is not enabled.
        any clientId = claims[CLIENT_ID];
        if (clientId is string) {
            userAuthContextDTO.client_id = clientId;
        }
    }
    return userAuthContextDTO;
}

function convertAnyToString(any variable) returns string{
    if (variable is string) {
        return variable;
    } 
    return UNKNOWN_VALUE;
}
