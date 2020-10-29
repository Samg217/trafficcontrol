package cfgfile

/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

import (
	"errors"

	"github.com/apache/trafficcontrol/lib/go-atscfg"
	"github.com/apache/trafficcontrol/lib/go-tc"
	"github.com/apache/trafficcontrol/traffic_ops_ort/atstccfg/config"
)

func GetConfigFileCDNSSLMultiCertDotConfig(toData *config.TOData) (string, string, string, error) {
	if toData.Server.CDNName == nil {
		return "", "", "", errors.New("server missing CDNName")
	}

	cfgDSes := atscfg.DeliveryServicesToSSLMultiCertDSes(toData.DeliveryServices)

	return atscfg.MakeSSLMultiCertDotConfig(tc.CDNName(*toData.Server.CDNName), toData.TOToolName, toData.TOURL, cfgDSes), atscfg.ContentTypeSSLMultiCertDotConfig, atscfg.LineCommentSSLMultiCertDotConfig, nil
}