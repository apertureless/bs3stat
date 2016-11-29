![bs3stat logo](/static/bs3.png)
# bs3stat

Is a small dashboard service in go and vue, to collect and visualize your backups made with [backup](https://github.com/backup/backup)

## Commands
**Run server:**
```
go run main.go
```

**Run server with custom port:**
```
go run main.go --port=3000
```

**Run with migrations:**
```
go run main.go --migrate
```

**Run with db seeds:**
```
go run main.go --seed
```

## Migration
**Create migration**
```
migrate --url sqlite3://storage.db --path ./db/migrations create add_example_field
```

**Migrate**
```
migrate --url sqlite3://storage.db --path ./db/migrations up # or down
```
