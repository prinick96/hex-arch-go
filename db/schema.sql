--
-- Struct for the table `todo`
--
CREATE TABLE IF NOT EXISTS todo (
    id uuid NOT NULL,
    _to varchar(100) NOT NULL,
    _do varchar(100) NOT NULL,
    PRIMARY KEY(id)
);