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

// Chain is a chain of alignments between two genomes.
type Chain struct {
	ID          int64      `db:"id"`           // Unique ID of the chain.
	Score       int64      `db:"score"`        // Alignment score.
	Ref         Reference  `db:"ref"`          // Reference genome assembly name.
	RefName     Chromosome `db:"ref_name"`     // Reference chromosome name.
	RefSize     int64      `db:"ref_size"`     // Size of the reference chromosome in bases.
	RefStrand   string     `db:"ref_strand"`   // Strand in the reference genome ('+' or '-').
	RefStart    int64      `db:"ref_start"`    // Start position in the reference genome.
	RefEnd      int64      `db:"ref_end"`      // End position in the reference genome.
	QueryName   Chromosome `db:"query_name"`   // Query chromosome name.
	QuerySize   int64      `db:"query_size"`   // Size of the query chromosome in bases.
	QueryStrand string     `db:"query_strand"` // Strand in the query genome ('+' or '-').
	QueryStart  int64      `db:"query_start"`  // Start position in the query genome.
	QueryEnd    int64      `db:"query_end"`    // End position in the query genome.
}

// Alignment is an alignment block between two genomes.
type Alignment struct {
	ID          int64 `db:"id"`           // Unique ID of the alignment.
	ChainID     int64 `db:"chain_id"`     // The chain this alignment belongs to.
	RefOffset   int64 `db:"ref_offset"`   // Offset of the aligned block in the reference chromosome from the start of the chain.
	QueryOffset int64 `db:"query_offset"` // Offset of the aligned block in the query chromosome from the start of the chain.
	Size        int64 `db:"size"`         // Size of the aligned block in bases.
}
