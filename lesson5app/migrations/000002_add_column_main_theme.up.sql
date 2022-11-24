ALTER TABLE snippets 
    ADD COLUMN main_theme VARCHAR(200);

UPDATE snippets SET main_theme = '';

ALTER TABLE snippets 
    DROP CONSTRAINT snippets_unique, 
    ADD CONSTRAINT snippets_unique UNIQUE(main_theme, header);
