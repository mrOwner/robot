schema "calculate" {}

table "candles" {
  schema = schema.calculate
  column "uid" {
    null = false
    type = varchar(36)
  }
  column "date" {
    null = false
    type = timestamptz
  }
  column "open" {
    null = false
    type = real
  }
  column "close" {
    null = false
    type = real
  }
  column "high" {
    null = false
    type = real
  }
  column "low" {
    null = false
    type = real
  }
  column "volume" {
    null = false
    type = bigserial
  }
}
