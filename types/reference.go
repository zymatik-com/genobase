/* SPDX-License-Identifier: AGPL-3.0-or-later
 *
 * Zymatik Genobase - A Human Genomics reference database.
 * Copyright (C) 2024 Damian Peckett <damian@pecke.tt>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published
 * by the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package types

// Reference is a reference genome assembly.
type Reference string

const (
	ReferenceNCBI36 Reference = "NCBI36"
	ReferenceGRCh37 Reference = "GRCh37"
	ReferenceGRCh38 Reference = "GRCh38"
)
