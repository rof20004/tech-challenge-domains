#!/usr/bin/env bash

echo "================= Executando testes de unidade =================="

go test -v ./tests/unit/... --bench . --benchmem -cover -coverpkg=./application/... -coverprofile=coverage.out
