-- UP --
CREATE TABLE migrations(id INT, name VARCHAR(255), hash VARCHAR(255));
-- UP --

-- DOWN --
DROP TABLE IF EXISTS migrations;
-- DOWN --