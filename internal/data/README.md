# Data

# Data
## Requirement
- [ent](https://entgo.io/)
- [atlas](https://atlasgo.io/)

## Generate
### New model/schema
```bash
ent new --target ent/schema <model-name>
```

### generate ent code
```bash
go generate ./internal/data/ent
```

## Migration
### create migration
```bash
atlas migrate diff [migration_name] \
  --dir "file://db/migrations" \
  --to "ent://internal/data/ent/schema" \
  --dev-url "docker://postgres/10/test?search_path=public" 
```

### apply migration
```bash
atlas migrate apply \
  --dir "file://db/migrations" \
  --url "postgres://postgres:pass@localhost:5432/database?search_path=public&sslmode=disable" 
```

### clean up
```bash
atlas schema clean -u "postgres://postgres:pass@localhost:5432/database?search_path=public&sslmode=disable"
```