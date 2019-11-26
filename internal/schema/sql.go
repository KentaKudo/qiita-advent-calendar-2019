package schema

var schemas = map[int]string{
	0: `
BEGIN;

CREATE TABLE schema_version {
        id        SMALLSERIAL NOT NULL PRIMARY KEY,
	md_insert BIGINT      DEFAULT  EXTRACT(EPOCH FROM current_timestamp)::INT,
	md_update BIGINT      DEFAULT  0,
	md_curr   BOOL        DEFAULT  'true'
};

INSERT INTO schema_version VALUES (0);

COMMIT;
`,
}
