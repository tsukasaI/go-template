# Template

| Tool | version |
| -- | -- |
| Go | 1.22.5 |

## Develop

1. Install `Dev Containers` to your local VSCode
1. Press F1 key and exec `Reopen in Container`

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

- カラム追加、制約追加（ALTER TABLE）の場合

```sh
$ migrate create -ext sql -dir src/driver/db/migrate/sql -seq add_{カラム名}_to_{TABLE NAME}
例) migrate create -ext sql -dir src/driver/db/migrate/sql -seq add_name_to_samples_table
```

- データ追加、casbin追加（INSERT INTO）の場合

```sh
$ migrate create -ext sql -dir src/driver/db/migrate/sql -seq insert_{データ投入内容}
例) migrate create -ext sql -dir src/driver/db/migrate/sql -seq insert_samples
```

- カラム削除の場合

```sh
$ migrate create -ext sql -dir src/driver/db/migrate/sql -seq del_{カラム名}_from_{TABLE NAME}
例) migrate create -ext sql -dir src/driver/db/migrate/sql -seq del_name_from_samples
```

up には反映用のSQLで、down には切り戻し用の SQL を記述します。

TODO
- use echo instead of gin
- prd docker env
- debugger
