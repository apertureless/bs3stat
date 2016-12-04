<p align="center">
	<img src="/static/bs3.png" alt="bs3stat" title="bs3stat" />
</p>

[![CircleCI](https://circleci.com/gh/apertureless/bs3stat/tree/master.svg?style=svg&circle-token=7eee193cd82594280172e81cd4c54b98f7179351)](https://circleci.com/gh/apertureless/bs3stat/tree/master)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/apertureless/bs3stat/blob/master/LICENSE.txt)

# bs3stat

Is a small dashboard service in go and vue, to collect and visualize your backups made with [backup](https://github.com/backup/backup)

**🚧 It is still in early alpha and not production ready⚠**

## Features

- Cluster statistics by model name
- Bar height based on backup duration
- Bar color based on backup status
- TODO: S3 integration with buckets and bucket cost
- TODO: Auth for api endpoints
- TODO: Simple login for dashboard
- TODO: Crud operations in frontend
- TODO: Tooltips with more info
- TODO: Docker container


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
