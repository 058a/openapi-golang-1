~/go/bin/oapi-codegen -generate types,server,spec -package oapicodegen ./../../api/hello.yaml > ./../../internal/infra/oapicodegen/hello/hello.go && go mod tidy

~/go/bin/oapi-codegen -generate types,server,spec -package oapicodegen ./../../api/stock.yaml > ./../../internal/infra/oapicodegen/stock/stock.go && go mod tidy

~/go/bin/oapi-codegen -generate types,server,spec -package oapicodegen ./../../api/inbound.yaml > ./../../internal/infra/oapicodegen/inbound/inbound.go && go mod tidy