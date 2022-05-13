/*
    Copyright (C) 2022 Tenable, Inc.

	Licensed under the Apache License, Version 2.0 (the "License");
    you may not use this file except in compliance with the License.
    You may obtain a copy of the License at

		http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
    distributed under the License is distributed on an "AS IS" BASIS,
    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
    See the License for the specific language governing permissions and
    limitations under the License.
*/

package functions

import (
	"strings"

	exp "github.com/VerbalExpressions/GoVerbalExpressions"
	"go.uber.org/zap"
)

// Variables function runs variable against a regular
// expression and return the variable key.
//
// For example:
// if var = Variables('identityName'),
// the function returns identityName as the key.
func Variables(variable string) string {
	const (
		start = "variables('"
		end   = "')"
	)

	key := strings.TrimPrefix(variable, "[")
	key = strings.TrimRight(key, "]")
	results := exp.New().
		StartOfLine().Find(start).
		BeginCapture().Anything().EndCapture().
		Find(end).EndOfLine().
		Captures(key)

	if len(results) == 0 {
		zap.S().Debugf("failed to parse expression: %s", variable)
		return ""
	}
	return results[0][1]
}
