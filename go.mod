module github.com/golang-devkit/lab-queue-pattern

go 1.26.5

retract [v0.0.0-0, v0.0.2] // published with wrong commit, use v0.0.3 instead

require go.uber.org/zap v1.28.0

require (
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/stretchr/testify v1.11.1 // indirect
	go.uber.org/multierr v1.10.0 // indirect
)
