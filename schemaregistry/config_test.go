/**
 * Copyright 2022 Confluent Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package schemaregistry

import (
	"testing"
)

func TestConfigWithAuthentication(t *testing.T) {
	maybeFail = initFailFunc(t)

	c := NewConfigWithAuthentication("mock://", "username", "password")

	maybeFail("BasicAuthCredentialsSource", expect(c.BasicAuthCredentialsSource, "USER_INFO"))
	maybeFail("BasicAuthUserInfo", expect(c.BasicAuthUserInfo, "username:password"))
}

func TestConfigWithSSLAuthentication(t *testing.T) {
	maybeFail = initFailFunc(t)

	sslCaString 	:= "CA"
	sslCertString	:= "CERTIFICATE"
	sslKeyString	:= "KEY"

	c := NewConfigWithSSLAuthentication("mock://", []byte(sslCaString), []byte(sslCertString), []byte(sslKeyString))

	maybeFail("SslCa", expect(c.SslCa, []byte(sslCaString)))
	maybeFail("SslCertificate", expect(c.SslCertificate, []byte(sslCertString)))
	maybeFail("SslKey", expect(c.SslKey, []byte(sslKeyString)))
}
