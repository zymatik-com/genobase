/* SPDX-License-Identifier: AGPL-3.0-or-later
 *
 * Zymatik Genobase - A human genomics reference DB.
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

// VariantClass is the class of a genomic variant.
type VariantClass string

const (
	// Single Nucleotide Variant: a variation in which a single nucleotide (base pair) is altered.
	VariantClassSNV VariantClass = "SNV"
	// INsertion/DELetion: a small insertion or deletion of bases in the DNA.
	VariantClassINDEL VariantClass = "INDEL"
	// INSertion: a type of variation where extra base pairs are inserted into a new place in the DNA.
	VariantClassINS VariantClass = "INS"
	// DELetion: a type of variation where some base pairs are deleted from the DNA.
	VariantClassDEL VariantClass = "DEL"
	// MultiNucleotide Variant: a variation where two or more nucleotides are replaced with other nucleotides.
	VariantClassMNV VariantClass = "MNV"
)

// Variant is a genomic variant (based on dbSNP).
type Variant struct {
	ID         int64        `db:"id"`         // Unique ID of the variant (RSID).
	Chromosome string       `db:"chromosome"` // Chromosome on which the variant is located.
	Position   int64        `db:"position"`   // Position of the variant on the chromosome.
	Class      VariantClass `db:"class"`      // Class of the variant, e.g., SNV, INDEL, INS, DEL, MNV.
}

// AncestryGroup identifies an ancestry group (as found in gnoMAD v3).
type AncestryGroup string

const (
	AncestryGroupAll AncestryGroup = "ALL"
	// AncestryGroupAfrican is the African/African American ancestry group.
	AncestryGroupAfrican AncestryGroup = "AFR"
	// AncestryGroupAmish is the Amish ancestry group.
	AncestryGroupAmish AncestryGroup = "AMI"
	// AncestryGroupAmerican is the Admixed American (Latino) ancestry group.
	AncestryGroupAmerican AncestryGroup = "AMR"
	// AncestryGroupAshkenazi is the Ashkenazi Jewish ancestry group.
	AncestryGroupAshkenazi AncestryGroup = "ASJ"
	// AncestryGroupEastAsian is the East Asian ancestry group.
	AncestryGroupEastAsian AncestryGroup = "EAS"
	// AncestryGroupFinnish is the Finnish ancestry group.
	AncestryGroupFinnish AncestryGroup = "FIN"
	// AncestryGroupMiddleEastern is the Middle Eastern ancestry group.
	AncestryGroupMiddleEastern AncestryGroup = "MID"
	// AncestryGroupEuropean is the Non-Finnish European ancestry group.
	AncestryGroupEuropean AncestryGroup = "NFE"
	// AncestryGroupSouthAsian is the South Asian ancestry group.
	AncestryGroupSouthAsian AncestryGroup = "SAS"
	// AncestryGroupOther encompasses all other ancestry groups.
	AncestryGroupOther AncestryGroup = "OTH"
)

// Allele is an allele of a genomic variant.
type Allele struct {
	ID        int64         `db:"id"`        // Unique ID of the variant the allele is associated with (RSID).
	Reference string        `db:"ref"`       // Reference base(s) at the variant's position.
	Alternate string        `db:"alt"`       // Alternate base(s) at the variant's position, representing the allele.
	Ancestry  AncestryGroup `db:"ancestry"`  // Ancestry group the allele is associated with.
	Frequency float64       `db:"frequency"` // Frequency of the allele in the ancestry group.
}
