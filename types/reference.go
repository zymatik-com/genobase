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

package types

// Reference is a reference genome assembly.
type Reference string

const (
	ReferenceNCBI36               Reference = "NCBI36"
	ReferenceGRCh37               Reference = "GRCh37"
	ReferenceGRCh38               Reference = "GRCh38"
	ReferenceTelomereToTelomereV2 Reference = "T2T-CHM13v2.0"
)
