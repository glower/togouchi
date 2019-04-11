REPO=togouchi/
GO?=go

LINT_FLAGS := run -v --deadline=120s
LINTER_EXE := golangci-lint
LINTER := ./bin/$(LINTER_EXE)
TESTFLAGS := -v -cover -tags=integration -timeout 120s

$(LINTER):
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s v1.15.0

check: gofmt lint

test:
	$(GO) test $(TESTFLAGS) ./...

lint: $(LINTER)
	$(LINTER) $(LINT_FLAGS) ./...

GFMT=find . -not \( \( -wholename "./vendor" \) -prune \) -name "*.go" | xargs gofmt -l
gofmt:
	@UNFMT=$$($(GFMT)); if [ -n "$$UNFMT" ]; then echo "gofmt needed on" $$UNFMT && exit 1; fi

vendor:
	dep ensure -v