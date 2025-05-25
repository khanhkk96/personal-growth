data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "ariga.io/atlas-provider-gorm",
    "load",
    "--path", "./db/models",  // set this to your models' path       
    "--dialect", "postgres",          // use: mysql, postgres, sqlite, or  sqlserver
  ]
}

env "gorm" {
  url = "postgres://postgres:postgres@localhost:5432/pgw?sslmode=disable"
  src = data.external_schema.gorm.url
  dev = "docker://postgres/15"     // or your actual dev url, must match dialect
  migration {
    dir = "file://db/migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}

env "docker" {
  url = "postgres://postgres:postgres@localhost:5433/pgw?sslmode=disable"
  src = data.external_schema.gorm.url
  dev = "docker://postgres/15"     // or your actual dev url, must match dialect
  migration {
    dir = "file://db/migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}
  