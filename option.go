/* SPDX-License-Identifier: MPL-2.0
 *
 * Zymatik Genobase - A human genomics reference DB.
 * Copyright (C) 2024 Damian Peckett <damian@pecke.tt>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the Mozilla Public License v2.0.
 *
 * You should have received a copy of the Mozilla Public License v2.0
 * along with this program. If not, see <https://mozilla.org/MPL/2.0/>.
 */

package genobase

import "strings"

type Option func(string) string

// ReadOnly makes the database read-only (allowing for multiple concurrent readers).
func ReadOnly(connStr string) string {
	if !strings.Contains(connStr, "?") {
		connStr += "?"
	} else {
		connStr += "&"
	}

	return connStr + "mode=ro"
}

// NoSync disables flushing the database to disk after each write.
// This is unsafe, but significantly speeds up bulk imports.
func NoSync(connStr string) string {
	if !strings.Contains(connStr, "?") {
		connStr += "?"
	} else {
		connStr += "&"
	}

	return connStr + "_journal_mode=OFF&_synchronous=OFF"
}
