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

package genobase

import (
	"context"
	"embed"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pressly/goose/v3"
	"github.com/zymatik-com/genobase/types"
)

type DB struct {
	db *sqlx.DB
}

//go:embed migrations/*.sql
var embedMigrations embed.FS

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

// Open opens a database connection and applies any necessary migrations.
func Open(ctx context.Context, logger *slog.Logger, dbPath string, opts ...Option) (*DB, error) {
	var connStr string

	// So that we can use an in-memory database for testing.
	if dbPath != "" {
		connStr = fmt.Sprintf("file:%s", dbPath)
	}

	for _, opt := range opts {
		connStr = opt(connStr)
	}

	db, err := sqlx.Connect("sqlite3", connStr)
	if err != nil {
		return nil, fmt.Errorf("could not open database: %w", err)
	}

	goose.SetLogger(&gooseLogger{logger})
	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("sqlite3"); err != nil {
		return nil, fmt.Errorf("could not set dialect: %w", err)
	}

	if err := goose.UpContext(ctx, db.DB, "migrations"); err != nil {
		return nil, fmt.Errorf("could not apply migrations: %w", err)
	}

	return &DB{
		db: db,
	}, nil
}

func (db *DB) Close() error {
	return db.db.Close()
}

func (db *DB) GetVariant(ctx context.Context, id int64) (*types.Variant, error) {
	rows, err := db.db.QueryxContext(ctx, "SELECT * FROM variant WHERE id = ? LIMIT 1", id)
	if err != nil {
		return nil, fmt.Errorf("could not query variants: %w", err)
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("no variant found: %w", os.ErrNotExist)
	}

	var variant types.Variant
	if err := rows.StructScan(&variant); err != nil {
		return nil, fmt.Errorf("could not unmarshal variant: %w", err)
	}

	return &variant, nil
}

func (db *DB) GetVariants(ctx context.Context, chromosome string, position int64) ([]types.Variant, error) {
	rows, err := db.db.QueryxContext(ctx, "SELECT * FROM variant WHERE chromosome = ? AND position = ?",
		chromosome, position)
	if err != nil {
		return nil, fmt.Errorf("could not query variants: %w", err)
	}
	defer rows.Close()

	var variants []types.Variant
	for rows.Next() {
		var variant types.Variant
		if err := rows.StructScan(&variant); err != nil {
			return nil, fmt.Errorf("could not unmarshal variant: %w", err)
		}

		variants = append(variants, variant)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("could not scan variants: %w", err)
	}

	return variants, nil
}

func (db *DB) StoreVariants(ctx context.Context, variants []types.Variant) error {
	tx, err := db.db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("could not start transaction: %w", err)
	}

	stmt, err := tx.PrepareNamedContext(ctx, `INSERT INTO variant (id, chromosome, position, class) 
	  VALUES (:id, :chromosome, :position, :class)
      ON CONFLICT(id) DO UPDATE SET 
        chromosome = excluded.chromosome, 
        position = excluded.position,
		class = excluded.class`)
	if err != nil {
		return fmt.Errorf("could not prepare statement: %w", err)
	}
	defer stmt.Close()

	for _, variant := range variants {
		if _, err := stmt.ExecContext(ctx, variant); err != nil {
			return fmt.Errorf("could not store variant: %w", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("could not commit transaction: %w", err)
	}

	return nil
}

func (db *DB) GetAllele(ctx context.Context, id int64, reference, alternate string, ancestry types.AncestryGroup) (*types.Allele, error) {
	rows, err := db.db.QueryxContext(ctx, "SELECT * FROM allele WHERE id = ? AND ref = ? AND alt = ? AND ancestry = ? LIMIT 1",
		id, reference, alternate, ancestry)
	if err != nil {
		return nil, fmt.Errorf("could not query alleles: %w", err)
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("no allele found: %w", os.ErrNotExist)
	}

	var allele types.Allele
	if err := rows.StructScan(&allele); err != nil {
		return nil, fmt.Errorf("could not unmarshal allele: %w", err)
	}

	return &allele, nil
}

func (db *DB) GetAlleles(ctx context.Context, id int64) ([]types.Allele, error) {
	rows, err := db.db.QueryxContext(ctx, "SELECT * FROM allele WHERE id = ?", id)
	if err != nil {
		return nil, fmt.Errorf("could not query alleles: %w", err)
	}
	defer rows.Close()

	var alleles []types.Allele
	for rows.Next() {
		var allele types.Allele
		if err := rows.StructScan(&allele); err != nil {
			return nil, fmt.Errorf("could not unmarshal allele: %w", err)
		}

		alleles = append(alleles, allele)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("could not scan alleles: %w", err)
	}

	return alleles, nil
}

func (db *DB) StoreAlleles(ctx context.Context, alleles []types.Allele) error {
	tx, err := db.db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("could not start transaction: %w", err)
	}

	stmt, err := tx.PrepareNamedContext(ctx, `INSERT INTO allele (id, ref, alt, ancestry, frequency) 
	  VALUES (:id, :ref, :alt, :ancestry, :frequency) 
	  ON CONFLICT(id, ref, alt, ancestry) DO UPDATE SET
		frequency = excluded.frequency`)
	if err != nil {
		return fmt.Errorf("could not prepare statement: %w", err)
	}
	defer stmt.Close()

	for _, allele := range alleles {
		if _, err := stmt.ExecContext(ctx, allele); err != nil {
			return fmt.Errorf("could not store allele: %w", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("could not commit transaction: %w", err)
	}

	return nil
}

func (db *DB) KnownAlleles(ctx context.Context) (map[int64]bool, error) {
	rows, err := db.db.QueryxContext(ctx, "SELECT DISTINCT id FROM allele")
	if err != nil {
		return nil, fmt.Errorf("could not query alleles: %w", err)
	}
	defer rows.Close()

	entries := make(map[int64]bool)
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			return nil, fmt.Errorf("could not scan allele: %w", err)
		}

		entries[id] = true
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("could not scan alleles: %w", err)
	}

	return entries, nil
}

func (db *DB) GetChain(ctx context.Context, from types.Reference, chromosome string, position int64) (*types.Chain, error) {
	rows, err := db.db.QueryxContext(ctx, "SELECT * FROM liftover_chain WHERE ref = ? AND ref_name = ? AND ref_start <= ? AND ref_end >= ? LIMIT 1",
		from, chromosome, position, position)
	if err != nil {
		return nil, fmt.Errorf("could not query chains: %w", err)
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("no chain found: %w", os.ErrNotExist)
	}

	var chain types.Chain
	if err := rows.StructScan(&chain); err != nil {
		return nil, fmt.Errorf("could not unmarshal chain: %w", err)
	}

	return &chain, nil
}

func (db *DB) StoreChain(ctx context.Context, from types.Reference, chain *types.Chain) (int64, error) {
	result, err := db.db.NamedExecContext(ctx, `
		INSERT INTO liftover_chain (
			score, ref, ref_name, ref_size, ref_strand, 
			ref_start, ref_end, query_name, query_size, 
			query_strand, query_start, query_end
		) VALUES (
			:score, :ref, :ref_name, :ref_size, :ref_strand, 
			:ref_start, :ref_end, :query_name, :query_size, 
			:query_strand, :query_start, :query_end
		)`, chain)
	if err != nil {
		return -1, fmt.Errorf("could not store chain: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, fmt.Errorf("could not get chain id: %w", err)
	}

	return id, nil
}

func (db *DB) GetAlignment(ctx context.Context, chainID, refOffset int64) (*types.Alignment, error) {
	rows, err := db.db.QueryxContext(ctx, "SELECT * FROM liftover_alignment WHERE chain_id = ? AND ref_offset + size >= ? ORDER BY ref_offset ASC LIMIT 1",
		chainID, refOffset)
	if err != nil {
		return nil, fmt.Errorf("could not query alignments: %w", err)
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("no alignment found: %w", os.ErrNotExist)
	}

	var alignment types.Alignment
	if err := rows.StructScan(&alignment); err != nil {
		return nil, fmt.Errorf("could not unmarshal alignment: %w", err)
	}

	return &alignment, nil
}

func (db *DB) StoreAlignments(ctx context.Context, chainID int64, alignments []types.Alignment) error {
	tx, err := db.db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("could not start transaction: %w", err)
	}

	for _, alignment := range alignments {
		alignment.ChainID = chainID

		if _, err := tx.NamedExecContext(ctx, `INSERT INTO liftover_alignment (chain_id, ref_offset, query_offset, size) 
			VALUES (:chain_id, :ref_offset, :query_offset, :size)`, alignment); err != nil {
			return fmt.Errorf("could not store alignment: %w", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("could not commit transaction: %w", err)
	}

	return nil
}
