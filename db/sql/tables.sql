
-- create the table for teams and the corresponding table for their name history

CREATE TABLE public."Teams"
(
    "ID" serial NOT NULL,
    "Name" text NOT NULL,
    "FoundedDate" text NOT NULL,
    "DissolvedDate" text,
    "Sport" text NOT NULL,
    PRIMARY KEY ("ID")
);

ALTER TABLE IF EXISTS public."Teams"
    OWNER to postgres;