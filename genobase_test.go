/* SPDX-License-Identifier: AGPL-3.0-or-later
 *
 * Zymatik Genobase - A genomics database for the Zymatik project.
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
}
