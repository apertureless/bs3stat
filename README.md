<p align="center">
	<img src="/static/bs3.png" alt="bs3stat" title="bs3stat" />
</p>

# bs3stat

Is a small dashboard service in go and vue, to collect and visualize your backups made with [backup](https://github.com/backup/backup)

**🚧 It is still in early alpha and not production ready⚠**

<p align="center">
	<img src="/static/bs3dash.png" alt="bs3stat" title="bs3stat" />
</p>


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

## Backup Notifier Settings

Accepted post parameters are

```json
{
	"name":"model_name",
	"title": "Model Name",
	"status": "Backup::Success",
	"duration": "00:00:24"

}
```
