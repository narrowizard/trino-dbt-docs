# Trino-dbt-docs
Generate dbt source defination from trino

## Usage
- Setting your trino connection info in `main.go`
- Run `go build` to build project
- Run `mkdir ./dist` to create dist directory
- Run `./trino-dbt-docs` to generate dbt source config file at `./dist/source.yml`
- Copy `./dist/source.yml` to your dbt project's models directory
- Switch to your dbt project 
- Run `dbt run` to build your dbt project
- Run `dbt docs generate` to generate documents for you dbt project
- Run `dbt docs serve` to start document services for your dbt project
