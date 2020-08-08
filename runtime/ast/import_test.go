/*
 * Cadence - The resource-oriented smart contract programming language
 *
 * Copyright 2019-2020 Dapper Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ast

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIdentifierLocation_MarshalJSON(t *testing.T) {

	loc := IdentifierLocation("test")

	actual, err := json.Marshal(loc)
	require.NoError(t, err)

	assert.JSONEq(t,
		`
        {
            "Type": "IdentifierLocation",
            "Identifier": "test"
        }
        `,
		string(actual),
	)
}

func TestStringLocation_MarshalJSON(t *testing.T) {

	loc := StringLocation("test")

	actual, err := json.Marshal(loc)
	require.NoError(t, err)

	assert.JSONEq(t,
		`
        {
            "Type": "StringLocation",
            "String": "test"
        }
        `,
		string(actual),
	)
}

func TestAddressLocation_MarshalJSON(t *testing.T) {

	loc := AddressLocation([]byte{1})

	actual, err := json.Marshal(loc)
	require.NoError(t, err)

	assert.JSONEq(t,
		`
        {
            "Type": "AddressLocation",
            "Address": "0x1"
        }
        `,
		string(actual),
	)
}