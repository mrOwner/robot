-- Drop schema named "public"
DROP SCHEMA "public" CASCADE;
-- Add new schema named "calculate"
CREATE SCHEMA "calculate";
-- create "candles" table
CREATE TABLE "calculate"."candles" ("uid" character varying(36) NOT NULL, "date" timestamptz NOT NULL, "open" real NOT NULL, "close" real NOT NULL, "high" real NOT NULL, "low" real NOT NULL, "volume" integer NULL);
