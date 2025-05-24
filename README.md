# 1.Migration 
### Generate migration file
```
atlas migrate diff --env gorm
```

### Migrate DB
```
atlas migrate apply --url "postgres://postgres:postgres@localhost:5432/pgw?sslmode=disable" --dir "file://db/migrations"
```

### Check DB migration status
```
atlas migrate status --url "postgres://postgres:postgres@localhost:5432/pgw?sslmode=disable" --dir "file://db/migrations"
```

### Down DB migration version
1. Revert DB to previous version
```
atlas migrate down --url "postgres://postgres:postgres@localhost:5432/pgw?sslmode=disable" --dir "file://db/migrations" --dev-url "docker://postgres/15"
```
2. Delete the lastest migration file
3. Re-hash file in atlas summary
```
atlas migrate hash --dir "file://db/migrations"
```
