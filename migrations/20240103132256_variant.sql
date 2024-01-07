-- +goose Up
-- +goose StatementBegin

-- The `chromosome` table stores information about chromosomes
-- we have variant information for.
CREATE table chromosome (
    id TEXT NOT NULL PRIMARY KEY,
    description TEXT
);

-- Populate the `chromosome` table with the chromosomes in dbSNP.
INSERT INTO chromosome (id, description) VALUES
    ('1', 'Chromosome 1'),
    ('2', 'Chromosome 2'),
    ('3', 'Chromosome 3'),
    ('4', 'Chromosome 4'),
    ('5', 'Chromosome 5'),
    ('6', 'Chromosome 6'),
    ('7', 'Chromosome 7'),
    ('8', 'Chromosome 8'),
    ('9', 'Chromosome 9'),
    ('10', 'Chromosome 10'),
    ('11', 'Chromosome 11'),
    ('12', 'Chromosome 12'),
    ('13', 'Chromosome 13'),
    ('14', 'Chromosome 14'),
    ('15', 'Chromosome 15'),
    ('16', 'Chromosome 16'),
    ('17', 'Chromosome 17'),
    ('18', 'Chromosome 18'),
    ('19', 'Chromosome 19'),
    ('20', 'Chromosome 20'),
    ('21', 'Chromosome 21'),
    ('22', 'Chromosome 22'),
    ('X', 'Chromosome X'),
    ('Y', 'Chromosome Y'),
    -- The Psuedo-autosomal regions are regions of the X and Y chromosomes
    -- that share homology and thus recombine during meiosis. Variants in 
    -- these regions are mapped to both sex chromosomes.
    ('PAR', 'Pseudo-autosomal region'),
    ('PAR2', 'Pseudo-autosomal region 2'),
    ('MT', 'Mitochondrial DNA');

-- The `variant_class` table stores information about the classes of
-- variants, eg. SNV, INDEL, INS, DEL, MNV.
CREATE TABLE variant_class (
    id TEXT NOT NULL PRIMARY KEY,
    description TEXT
);

-- Populate the `variant_class` table with the variant classes in dbSNP.
INSERT INTO variant_class (id, description) VALUES
    ('SNV', 'Single nucleotide variant'),
    ('INDEL', 'Insertion or deletion'),
    ('INS', 'Insertion'),
    ('DEL', 'Deletion'),
    ('MNV', 'Multiple nucleotide variant');

-- The `variant` table stores information about genetic variants 
-- (based on dbSNP).
CREATE TABLE variant (
    -- The RSID of the variant.
    id INTEGER NOT NULL PRIMARY KEY,
    -- The chromosome on which the variant is located.
    chromosome TEXT,
    -- The position of the variant on the chromosome.
    position INTEGER,
    -- The reference base(s) at the variant's position.
    ref TEXT,
    -- The class of the variant, e.g., SNV, INDEL, INS, DEL, MNV.
    class TEXT,
    FOREIGN KEY (chromosome) REFERENCES chromosome (id),
    FOREIGN KEY (class) REFERENCES variant_class (id)
);
CREATE INDEX variant_coordinate ON variant(chromosome, position);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE variant;

DROP TABLE variant_class;

DROP TABLE chromosome;

-- +goose StatementEnd
