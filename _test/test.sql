-- Used for testing skirmish.Restore

DROP SCHEMA IF EXISTS public CASCADE;
CREATE SCHEMA public;
DROP TABLE IF EXISTS example CASCADE;
CREATE TABLE example (
    pk SERIAL PRIMARY KEY,
    text_field TEXT
);

INSERT INTO example ("text_field") VALUES ('ex_1'), ('ex_2');