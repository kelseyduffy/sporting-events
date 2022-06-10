
-- create the table for teams and the corresponding table for their name history

CREATE TABLE IF NOT EXISTS public."Teams"
(
    "ID" integer NOT NULL DEFAULT nextval('"Teams_ID_seq"'::regclass),
    "FoundedYear" smallint NOT NULL,
    "DissolvedYear" smallint,
    CONSTRAINT "Teams_pkey" PRIMARY KEY ("ID")
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public."Teams"
    OWNER to postgres;

CREATE TABLE IF NOT EXISTS public."TeamNameHistory"
(
    "TeamID" integer NOT NULL,
    "Name" text NOT NULL,
    "StartYear" smallint NOT NULL,
    "EndYear" smallint,
    PRIMARY KEY ("TeamID", "StartYear"),
    CONSTRAINT "TeamID" FOREIGN KEY ("TeamID")
        REFERENCES public."Teams" ("ID") MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
);

ALTER TABLE IF EXISTS public."TeamNameHistory"
    OWNER to postgres;

-- insert a team (Houston Dynamo) into the Teams table
INSERT INTO public."Teams"(
	"FoundedYear")
	VALUES (2005);

-- insert the team name history entries into the TeamNameHistory table
INSERT INTO public."TeamNameHistory"(
	"TeamID", "Name", "StartYear", "EndYear")
	VALUES  (1, 'Houston Dynamo', 2005, 2020),
   	        (1, 'Houston Dynamo FC', 2020, NULL);

-- insert them all together in a robust way to grab the just created Team ID
WITH team AS (
    INSERT INTO "Teams"
        ("FoundedYear", "DissolvedYear")
    VALUES
        (1995, null)
    RETURNING "ID"
)
INSERT INTO "TeamNameHistory" ("TeamID", "Name", "StartYear", "EndYear")
	SELECT "ID", 'Kansas City Wiz', 1995, 1996
    FROM team
    UNION
    SELECT "ID", 'Kansas City Wizards', 1996, 2011
    FROM team
    UNION
    SELECT "ID", 'Sporting Kansas City', 2011, null
    FROM team

-- query the current names and founded years of all teams
SELECT n."Name", t."FoundedYear"
	FROM public."TeamNameHistory" AS n
    INNER JOIN public."Teams" AS t
    ON n."TeamID" = t."ID"
    WHERE n."EndYear" IS NULL