package schema

var schemas = map[int]string{
	0: `
BEGIN;

CREATE TABLE schema_version ( 
        id        SMALLSERIAL NOT NULL PRIMARY KEY,
	md_insert BIGINT      DEFAULT  EXTRACT(EPOCH FROM current_timestamp)::INT,
	md_update BIGINT      DEFAULT  0,
	md_curr   BOOL        DEFAULT  'true'
);

INSERT INTO schema_version VALUES (0);

COMMIT;
`,
	1: `
BEGIN;

CREATE TABLE todo ( 
        id          UUID   NOT NULL PRIMARY KEY,
	title       STRING NOT NULL,
	description STRING NOT NULL
);

UPDATE schema_version SET md_curr = false, md_update = EXTRACT(EPOCH FROM current_timestamp)::INT WHERE md_curr = true;
INSERT INTO schema_version VALUES (1);

COMMIT;
`,
}
