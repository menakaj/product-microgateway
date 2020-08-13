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

import ballerina/http;

# Data holder for Application details available in Gateway Pilot node. 
# 
# + applications - Map of `Application` objects
type ApplicationDataStore object {
    //Tenant wise applications map
    map<Application> applications = {};

    private string pilotUsername;
    private string pilotPassword;
    private string serviceContext;

    public function __init(string username, string password, string context) {
        self.pilotUsername = username;
        self.pilotPassword = password;
        self.serviceContext = context + "/applications";
        if (apimEventHubEnabled) {
            future<()> applicationsFetch = start self.fetchApplications();
        }
    }

    # Retrieve a specific `Application` object from the Applicatio Data Store.
    # 
    # + appId - Application Id of the required `Application`
    # + return - `Application` with provided `appId`. If no match was found `()` is returned.
    function getApplication(string appId) returns (Application | ()) {
        if (self.applications.hasKey(appId)) {
            return self.applications.get(appId);
        }
        return ();
    }

    function addApplication(Application app) {
        string appKey = app.id.toString();
        lock {
            //Writing event should be locked, due to worker threads are reading the map during request validations
            self.applications[appKey] = app;
        }
    }

    function removeApplication(Application app) {
        lock {
            //Remove event should be locked, due to worker threads are reading the map during request validations
            Application removedApp = self.applications.remove(app.id.toString());
        }
    }

    function loadAppplicationFromService(int appId) returns Application? {
        string basicAuthHeader = buildBasicAuthHeader(self.pilotUsername, self.pilotPassword);
        string tenantDomain = getPilotAuthenticatedUserTenantDomain(self.pilotUsername);
        http:Request apiReq = new;
        apiReq.setHeader(AUTHORIZATION_HEADER, basicAuthHeader);
        string serviceContext = self.serviceContext + "?appId=" + appId.toString();
        var response = gatewayPilotEndpoint->get(serviceContext, message = apiReq);
        if (response is http:Response) {
            var payload = response.getJsonPayload();
            if (payload is json) {
                printDebug(KEY_APPLICATION_STORE, "Application list for id : " + appId.toString() + " is : " + payload.toJsonString());
                json[] list = <json[]>payload.list;
                if (list.length() > 0 ) {
                    Application app = {
                        uuid: list[0].uuid.toString(),
                        id: <int>list[0].id,
                        owner: list[0].subName.toString(),
                        name: list[0].name.toString(),
                        policyId: list[0].policy.toString(),
                        tokenType: list[0].tokenType.toString(),
                        tenantDomain: tenantDomain,
                        groupIds: <json[]>list[0].groupIds,
                        attributes: <map<json>>list[0].attributes
                    };
                    self.applications[appId.toString()] = <@untainted>app;
                    printDebug(KEY_APPLICATION_STORE, "Returned application from service is : " + app.toString());
                    return <@untainted>app;
                }
            } else {
              printError(KEY_APPLICATION_STORE, "Received invalid application data for id : " + appId.toString(), payload);
            }
        } else {
          printError(KEY_APPLICATION_STORE, "Failed to retrieve application data for id : " + appId.toString(), response);
        }
        return ();
    }

    private function fetchApplications() {
        string tenantDomain = getPilotAuthenticatedUserTenantDomain(self.pilotUsername);
        string basicAuthHeader = buildBasicAuthHeader(self.pilotUsername, self.pilotPassword);
        http:Request appReq = new;
        appReq.setHeader(AUTHORIZATION_HEADER, basicAuthHeader);
        var response = gatewayPilotEndpoint->get(self.serviceContext, message = appReq);
        if (response is http:Response) {
            map<Application> applicationMap = {};
            var payload = response.getJsonPayload();
            if (payload is json) {
                printDebug(KEY_APPLICATION_STORE, "Application list is : " + payload.toJsonString());
                json[] list = <json[]>payload.list;
                printDebug(KEY_APPLICATION_STORE, "Received valid application details");
                foreach json jsonApp in list {
                    Application app = {
                        uuid: jsonApp.uuid.toString(),
                        id: <int>jsonApp.id,
                        owner: jsonApp.subName.toString(),
                        name: jsonApp.name.toString(),
                        policyId: jsonApp.policy.toString(),
                        tokenType: jsonApp.tokenType.toString(),
                        tenantDomain: tenantDomain,
                        groupIds: <json[]>jsonApp.groupIds,
                        attributes: <map<json>>jsonApp.attributes
                    };
                    string appKey = app.id.toString();
                    self.applications[appKey] = app;
                }
            } else {
                printError(KEY_APPLICATION_STORE, "Received invalid application data", payload);
            }
        } else {
            printError(KEY_APPLICATION_STORE, "Failed to retrieve application data", response);
        }

    }
};
