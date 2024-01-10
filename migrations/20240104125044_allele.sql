-- +goose Up
-- +goose StatementBegin

-- The `ancestry_group` table stores information about ancestry groups.
-- we have allele frequencies for.
CREATE TABLE ancestry_group (
    id TEXT NOT NULL PRIMARY KEY,
    description TEXT
);

-- Populate the `ancestry_group` table with the ancestry groups from gnomAD v3.
INSERT INTO ancestry_group (id, description) VALUES
    ('ALL', 'All populations'),
    ('AFR', 'African/African American'),
    ('AMI', 'Amish'),
    ('AMR', 'Admixed American (Latino)'),
    ('ASJ', 'Ashkenazi Jewish'),
    ('EAS', 'East Asian'),
    ('FIN', 'Finnish'),
    ('MID', 'Middle Eastern'),
    ('NFE', 'Non-Finnish European'),
    ('SAS', 'South Asian');

-- The `allele` table stores information about alleles.
-- Based on gnomAD v3.2.1.
CREATE TABLE allele (
    -- Unique ID of the variant the allele is associated with (RSID).
    id INTEGER,
    -- The reference base(s) at the variant's position.
    ref TEXT,
    -- Alternate base(s) at the variant's position, representing the allele.
    alt TEXT,
    -- Ancestry group the allele is associated with.
    ancestry TEXT,
    -- Frequency of the allele in the ancestry group.
    frequency REAL,
    PRIMARY KEY (id, ref, alt, ancestry),
    FOREIGN KEY (ancestry) REFERENCES ancestry_group (id)  
);
CREATE INDEX allele_id ON allele(id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE allele;

DROP TABLE ancestry_group;

-- +goose StatementEnd
