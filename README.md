# Template

| Tool | version |
| -- | -- |
| Go | 1.22.5 |
| Echo | 4.12.0 |

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

### How to create Migration file


```sh
$ migrate create -ext sql -dir src/driver/db/migrate/sql -seq {FILE_NAME}
e.g.) migrate create -ext sql -dir src/driver/db/migrate/sql -seq create_samples_table
```

TODO
- prd docker env
