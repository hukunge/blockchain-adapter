/*
Copyright ArxanChain Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package structs

// signature header
const (
	SignatureTimestamp = "X-Signature-Timestamp"
	SignatureMethod    = "X-Signature-Method"
	SignatureInfo      = "X-Signature"
	APIKeyHeader       = "API-Key"

	HMACSHA256 = "hmac-sha256"
)

// http method
const (
	PostMethod = "POST"
	GetMethod  = "GET"
)

// arxanchain api url
const (
	CreateAccountURL = "/api/v1/abox/account/create"

	CreateAssetURL      = "/api/v1/abox/asset/register"
	QueryAssetDetailURL = "/api/v1/abox/asset/detail"
)

// aranchain blockchain action type
const (
	CreateAccount = "create_acc"

	CreateAsset = "create_asset"
	QueryAsset  = "query_asset"
)
