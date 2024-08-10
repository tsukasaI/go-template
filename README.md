# Template

| Tool | version |
| -- | -- |
| Go | 1.22.5 |

## Develop

1. Install `Dev Containers` to your local VSCode
1. Press F1 key and select `Reopen in Container`

## debug

1. Execute 'air` command.
1. Press F5 key

## Production build

```
docker build -f production/Dockerfile .
```

## Migration

golang-migrate

https://github.com/golang-migrate/migrate

### Migration ファイルの作成方法

- CREATE TABLE

```sh
$ migrate create -ext sql -dir src/driver/db/migrate/sql -seq create_{TABLE NAME}
例) migrate create -ext sql -dir src/driver/db/migrate/sql -seq create_samples
```

- ALTER TABLE

```sh
$ migrate create -ext sql -dir src/driver/db/migrate/sql -seq add_{COLUMN NAME}_to_{TABLE NAME}
例) migrate create -ext sql -dir src/driver/db/migrate/sql -seq add_name_to_samples_table
```

- INSERT

```sh
$ migrate create -ext sql -dir src/driver/db/migrate/sql -seq insert_{DATA}
例) migrate create -ext sql -dir src/driver/db/migrate/sql -seq insert_samples
```

- DROP

```sh
$ migrate create -ext sql -dir src/driver/db/migrate/sql -seq drop_{カラム名}_from_{TABLE NAME}
例) migrate create -ext sql -dir src/driver/db/migrate/sql -seq drop_name_from_samples
```

TODO
- use echo instead of gin and graceful stop
- prd docker env
- slog
