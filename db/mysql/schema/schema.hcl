schema "robot" {
  charset = "utf8"
  collate = "utf8_general_ci"
  comment = "for accumulate information about stocks"
}

table "candles" {
  schema = schema.robot
  column "uid" {
    null = false
    type = varchar(36)
  }
  column "date" {
    null = false
    type = timestamp
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
    null = true
    type = int
  }
}
