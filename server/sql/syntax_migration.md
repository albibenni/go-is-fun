# Migration with goose

## Command to run Migration

`goose postgres postgresql://albibenni:password@localhost:5432/go-db up`

- type of db selected
- complete url of the go-db
- example that should be run from `/schema` folder
- `up` migration

`goose postgres postgresql://albibenni:password@localhost:5432/go-db down`

- `down` migration
