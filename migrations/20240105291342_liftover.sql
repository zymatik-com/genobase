-- +goose Up
-- +goose StatementBegin

-- We store the liftover chains in a database so that we can query them
-- efficiently. The `liftover_chain` table stores the chain metadata, and 
-- the `liftover_alignment` table stores the alignment blocks.
-- See: https://genome.ucsc.edu/goldenPath/help/chain.html
CREATE TABLE liftover_chain (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    -- Alignment score.
    score INTEGER,
    -- Reference genome assembly name.
    ref TEXT,
    -- Reference chromosome name.
    ref_name TEXT,
    -- Size of the reference chromosome.
    ref_size INTEGER,
    -- Strand in the reference genome ('+' or '-').
    ref_strand TEXT,
    -- Start position in the reference genome.
    ref_start INTEGER,
    -- End position in the reference genome.
    ref_end INTEGER,
    -- Query chromosome name.
    query_name TEXT,
    -- Size of the query chromosome.
    query_size INTEGER,
    -- Strand in the query genome ('+' or '-').
    query_strand TEXT,
    -- Start position in the query genome.
    query_start INTEGER,
    -- End position in the query genome.
    query_end INTEGER,
    FOREIGN KEY (ref_name) REFERENCES chromosome (id),
    FOREIGN KEY (query_name) REFERENCES chromosome (id)
);
CREATE INDEX liftover_chain_ref_name ON liftover_chain(ref, ref_name);
CREATE INDEX liftover_chain_ref_start ON liftover_chain(ref_start);
CREATE INDEX liftover_chain_ref_end ON liftover_chain(ref_end);

-- See: https://genome.ucsc.edu/goldenPath/help/chain.html
CREATE TABLE liftover_alignment (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    -- The chain that this alignment belongs to.
    chain_id INTEGER,
    -- Offset from the start of the chain in the reference genome.
    ref_offset INTEGER,
    -- Offset from the start of the chain in the query genome.
    query_offset INTEGER,
    --  Size of the aligned block in bases.
    size INTEGER,
    FOREIGN KEY (chain_id) REFERENCES liftover_alignment (id)
);
CREATE INDEX liftover_alignment_chain_id ON liftover_alignment(chain_id);
CREATE INDEX liftover_alignment_ref_offset ON liftover_alignment(ref_offset);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE liftover_alignment;

DROP TABLE liftover_chain;

-- +goose StatementEnd
