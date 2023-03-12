env "mysql" {
  // Declare where the schema definition resides.
  // Also supported: ["multi.hcl", "file.hcl"].
  src = "./db/mysql/schema/schema.hcl"

  // Define the URL of the database which is managed
  // in this environment.
  url = "mysql://root:pass@localhost:3306/robot"

  // Define the URL of the Dev Database for this environment
  // See: https://atlasgo.io/concepts/dev-database
  dev = "mysql://root:pass@localhost:3306/dev"
  // dev = "docker://mysql/5.7/robot"

  // A block defines the migration configuration of the env.
  // See: https://atlasgo.io/atlas-schema/projects#environments
  migration {
      dir = "file://db/mysql/migrations"
  }
}

env "postgres" {
  // Declare where the schema definition resides.
  // Also supported: ["multi.hcl", "file.hcl"].
  src = "./db/postgres/schema/schema.hcl"

  // Define the URL of the database which is managed
  // in this environment.
  url = "postgres://root:secret@localhost:5432/robot?search_path=calculate&sslmode=disable"

  // Define the URL of the Dev Database for this environment
  // See: https://atlasgo.io/concepts/dev-database
  dev = "postgres://root:secret@localhost:5432/dev?sslmode=disable"
  // dev = "docker://postgres/15"

  // A block defines the migration configuration of the env.
  // See: https://atlasgo.io/atlas-schema/projects#environments
  migration {
      dir = "file://db/postgres/migrations"
  }
}
