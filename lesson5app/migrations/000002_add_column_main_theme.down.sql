ALTER TABLE snippets DROP CONSTRAINT snippets_unique;
ALTER TABLE snippets DROP COLUMN main_theme;
ALTER TABLE snippets ADD CONSTRAINT snippets_unique UNIQUE(header);