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

import "strconv"

// Chromosome is a chromosome in a genome.
type Chromosome string

// Chromosome constants.
const (
	Chr1    Chromosome = "1"
	Chr2    Chromosome = "2"
	Chr3    Chromosome = "3"
	Chr4    Chromosome = "4"
	Chr5    Chromosome = "5"
	Chr6    Chromosome = "6"
	Chr7    Chromosome = "7"
	Chr8    Chromosome = "8"
	Chr9    Chromosome = "9"
	Chr10   Chromosome = "10"
	Chr11   Chromosome = "11"
	Chr12   Chromosome = "12"
	Chr13   Chromosome = "13"
	Chr14   Chromosome = "14"
	Chr15   Chromosome = "15"
	Chr16   Chromosome = "16"
	Chr17   Chromosome = "17"
	Chr18   Chromosome = "18"
	Chr19   Chromosome = "19"
	Chr20   Chromosome = "20"
	Chr21   Chromosome = "21"
	Chr22   Chromosome = "22"
	ChrX    Chromosome = "X"
	ChrY    Chromosome = "Y"
	ChrMT   Chromosome = "MT"
	ChrPAR  Chromosome = "PAR"
	ChrPAR2 Chromosome = "PAR2"
)

func (c Chromosome) Int() int {
	switch c {
	case "X":
		return 23
	case "Y":
		return 24
	case "MT":
		return 25
	case "PAR":
		return 26
	case "PAR2":
		return 27
	default:
		num, _ := strconv.Atoi(string(c))
		return num
	}
}
