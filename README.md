# 1.Migration 
### Generate migration file
```
atlas migrate diff --env gorm
```

### Migrate DB
```
atlas migrate apply --env gorm
```

### Check DB migration status
```
atlas migrate status --env gorm
```

### Down DB migration version
1. Revert DB to previous version
```
atlas migrate down --env gorm
```
2. Delete the lastest migration file
3. Re-hash file in atlas summary
```
atlas migrate hash --env gorm
```

# Config the .env file when running Docker
```
REDIS_HOST=pgw_redis
```
```
DB_HOST=pgw_postgres
```
