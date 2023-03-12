-- create sequence for serial column "volume"
CREATE SEQUENCE IF NOT EXISTS "calculate"."candles_volume_seq" OWNED BY "calculate"."candles"."volume";
-- modify "candles" table
ALTER TABLE "calculate"."candles" ALTER COLUMN "volume" SET DEFAULT nextval('"calculate"."candles_volume_seq"'), ALTER COLUMN "volume" TYPE bigint, ALTER COLUMN "volume" SET NOT NULL;
