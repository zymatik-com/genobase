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

package genobase_test

import (
	"context"
	"testing"

	"github.com/neilotoole/slogt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zymatik-com/genobase"
	"github.com/zymatik-com/genobase/types"
)

func TestGenobase(t *testing.T) {
	ctx := context.Background()

	db, err := genobase.Open(ctx, slogt.New(t), "", false)
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, db.Close())
	})

	t.Run("Variants", func(t *testing.T) {
		variants := []types.Variant{
			{
				ID:         4680,
				Chromosome: "22",
				Position:   19963748,
				Reference:  "G",
				Class:      types.VariantClassSNV,
			},
			{
				ID:         334,
				Chromosome: "11",
				Position:   5227002,
				Reference:  "T",
				Class:      types.VariantClassSNV,
			},
		}

		require.NoError(t, db.StoreVariants(ctx, variants))

		variant, err := db.GetVariant(ctx, 4680)
		require.NoError(t, err)

		assert.Equal(t, variant.ID, int64(4680))
		assert.Equal(t, variant.Chromosome, "22")
		assert.Equal(t, variant.Position, int64(19963748))
		assert.Equal(t, variant.Reference, "G")
		assert.Equal(t, variant.Class, types.VariantClassSNV)
	})

	t.Run("Alleles", func(t *testing.T) {
		alleles := []types.Allele{
			{
				// rs4680 slow COMT.
				ID:        4680,
				Reference: "G",
				Alternate: "A",
				Frequency: 0.461091,
				Ancestry:  types.AncestryGroupAll,
			},
			{
				// rs334 sickle cell anemia.
				ID:        334,
				Reference: "T",
				Alternate: "A",
				Frequency: 0.003480,
				Ancestry:  types.AncestryGroupAll,
			},
		}

		require.NoError(t, db.StoreAlleles(ctx, alleles))

		allele, err := db.GetAllele(ctx, 4680, "G", "A", types.AncestryGroupAll)
		require.NoError(t, err)

		assert.Equal(t, allele.ID, int64(4680))
		assert.Equal(t, allele.Reference, "G")
		assert.Equal(t, allele.Alternate, "A")
		assert.Equal(t, allele.Ancestry, types.AncestryGroupAll)
		assert.Equal(t, allele.Frequency, 0.461091)

		// Update the frequency.
		alleles[0].Frequency = 0.5

		require.NoError(t, db.StoreAlleles(ctx, alleles[:1]))

		allele, err = db.GetAllele(ctx, 4680, "G", "A", types.AncestryGroupAll)
		require.NoError(t, err)

		assert.Equal(t, allele.ID, int64(4680))
		assert.Equal(t, allele.Ancestry, types.AncestryGroupAll)
		assert.Equal(t, allele.Frequency, 0.5)

		alleleList, err := db.GetAlleles(ctx, 4680)
		require.NoError(t, err)

		assert.Len(t, alleleList, 1)
		assert.Equal(t, alleleList[0].ID, int64(4680))
		assert.Equal(t, alleleList[0].Ancestry, types.AncestryGroupAll)
	})

	t.Run("Liftover", func(t *testing.T) {
		chain := &types.Chain{
			Score:       1,
			Ref:         types.ReferenceGRCh37,
			RefName:     "1",
			RefSize:     249250621,
			RefStrand:   "+",
			RefStart:    10000,
			RefEnd:      267719,
			QueryName:   "1",
			QuerySize:   248956422,
			QueryStrand: "+",
			QueryStart:  10000,
			QueryEnd:    297968,
		}

		chainID, err := db.StoreChain(ctx, types.ReferenceGRCh37, chain)
		require.NoError(t, err)

		assert.Equal(t, chainID, int64(1))

		alignments := []types.Alignment{{
			ChainID:     chainID,
			RefOffset:   0,
			QueryOffset: 0,
			Size:        167417,
		}, {
			ChainID:     chainID,
			RefOffset:   217417,
			QueryOffset: 247666,
			Size:        40302,
		}}

		err = db.StoreAlignments(ctx, chainID, alignments)
		require.NoError(t, err)

		retrievedChain, err := db.GetChain(ctx, types.ReferenceGRCh37, "1", 217480)
		require.NoError(t, err)

		assert.Equal(t, retrievedChain.ID, chainID)
		assert.Equal(t, retrievedChain.Score, int64(1))
		assert.Equal(t, retrievedChain.Ref, types.ReferenceGRCh37)
		assert.Equal(t, retrievedChain.RefName, "1")
		assert.Equal(t, retrievedChain.RefSize, int64(249250621))
		assert.Equal(t, retrievedChain.RefStrand, "+")
		assert.Equal(t, retrievedChain.RefStart, int64(10000))
		assert.Equal(t, retrievedChain.RefEnd, int64(267719))
		assert.Equal(t, retrievedChain.QueryName, "1")
		assert.Equal(t, retrievedChain.QuerySize, int64(248956422))
		assert.Equal(t, retrievedChain.QueryStrand, "+")
		assert.Equal(t, retrievedChain.QueryStart, int64(10000))
		assert.Equal(t, retrievedChain.QueryEnd, int64(297968))

		retrievedAlignment, err := db.GetAlignment(ctx, chainID, 217480)
		require.NoError(t, err)

		assert.NotZero(t, retrievedAlignment.ID)
		assert.Equal(t, retrievedAlignment.ChainID, chainID)
		assert.Equal(t, retrievedAlignment.RefOffset, int64(217417))
		assert.Equal(t, retrievedAlignment.QueryOffset, int64(247666))
		assert.Equal(t, retrievedAlignment.Size, int64(40302))
	})
}
