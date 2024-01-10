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
	Chr1  Chromosome = "1"
	Chr2  Chromosome = "2"
	Chr3  Chromosome = "3"
	Chr4  Chromosome = "4"
	Chr5  Chromosome = "5"
	Chr6  Chromosome = "6"
	Chr7  Chromosome = "7"
	Chr8  Chromosome = "8"
	Chr9  Chromosome = "9"
	Chr10 Chromosome = "10"
	Chr11 Chromosome = "11"
	Chr12 Chromosome = "12"
	Chr13 Chromosome = "13"
	Chr14 Chromosome = "14"
	Chr15 Chromosome = "15"
	Chr16 Chromosome = "16"
	Chr17 Chromosome = "17"
	Chr18 Chromosome = "18"
	Chr19 Chromosome = "19"
	Chr20 Chromosome = "20"
	Chr21 Chromosome = "21"
	Chr22 Chromosome = "22"
	ChrX  Chromosome = "X"
	ChrY  Chromosome = "Y"
	// ChrMT is the mitochondrial DNA.
	ChrMT Chromosome = "MT"
	// ChrPAR and ChrPAR2 are the pseudoautosomal regions.
	// The Psuedoautosomal regions are regions of the X and Y chromosomes
	// that share homology and thus recombine during meiosis.
	// Variants in these regions are mapped to both sex chromosomes.
	ChrPAR  Chromosome = "PAR"
	ChrPAR2 Chromosome = "PAR2"
)

func (c Chromosome) String() string {
	return string(c)
}

// Less returns true if the chromosome is less than the argument.
func (c Chromosome) Less(comparison Chromosome) bool {
	return c.int() < comparison.int()
}

// Length returns the length of the chromosome in base pairs.
func (c Chromosome) Length(reference Reference) int64 {
	if reference != ReferenceGRCh38 {
		panic("Only GRCh38 is supported")
	}

	switch c {
	case Chr1:
		return 248956422
	case Chr2:
		return 242193529
	case Chr3:
		return 198295559
	case Chr4:
		return 190214555
	case Chr5:
		return 181538259
	case Chr6:
		return 170805979
	case Chr7:
		return 159345973
	case Chr8:
		return 145138636
	case Chr9:
		return 138394717
	case Chr10:
		return 133797422
	case Chr11:
		return 135086622
	case Chr12:
		return 133275309
	case Chr13:
		return 114364328
	case Chr14:
		return 107043718
	case Chr15:
		return 101991189
	case Chr16:
		return 90338345
	case Chr17:
		return 83257441
	case Chr18:
		return 80373285
	case Chr19:
		return 58617616
	case Chr20:
		return 64444167
	case Chr21:
		return 46709983
	case Chr22:
		return 50818468
	case ChrX:
		return 156040895
	case ChrY:
		return 57227415
	case ChrMT:
		return 16569
	case ChrPAR:
		return 2771479
	case ChrPAR2:
		return 329513
	default:
		panic("Unknown chromosome")
	}
}

func (c Chromosome) int() int {
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
