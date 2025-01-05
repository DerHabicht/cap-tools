SRC = $(shell find . -name "*.go")

BASE_ENV_CONFIG_KEY = CAPTOOLS
DEV_CONFIG_DIR=testdata/dev/config
TEST_CONFIG_DIR=testdata/test/config

# Credit to https://github.com/commissure/go-git-build-vars for giving me a starting point for this.
BUILD_TIME = `date +%Y%m%d%H%M%S`
GIT_REVISION = `git rev-parse --short HEAD`
GIT_BRANCH = `git rev-parse --symbolic-full-name --abbrev-ref HEAD | sed 's/\//-/g'`
GIT_DIRTY = `git diff-index --quiet HEAD -- || echo 'x-'`

LDFLAGS = -ldflags "-s -X main.BuildTime=$(BUILD_TIME) -X main.GitRevision=$(GIT_DIRTY)$(GIT_REVISION) -X main.GitBranch=$(GIT_BRANCH)"

.PHONY: all
all: bin/capa5srv

bin/capa5srv: $(foreach f, $(SRC), $(f))
	go build $(LD_FLAGS) -o bin/capa5srv cmd/capa5srv/main.go

.PHONY: dev_setup
dev_setup:
	mkdir -p $(DEV_CONFIG_DIR)
	mkdir -p $(TEST_CONFIG_DIR)

.PHONY: test_db_up
test_db_up:
	cd tools/migrate/ && $(BASE_ENV_CONFIG_KEY)_CONFIG=../../$(TEST_CONFIG_DIR) go run migrate.go up

.PHONY: test_db_down
test_db_down:
	cd tools/migrate/ && $(BASE_ENV_CONFIG_KEY)_CONFIG=../../$(TEST_CONFIG_DIR) go run migrate.go down

.PHONY: test_db_reset
test_db_reset:
	cd tools/migrate/ && $(BASE_ENV_CONFIG_KEY)_CONFIG=../../$(TEST_CONFIG_DIR) go run migrate.go drop
	cd tools/migrate/ && $(BASE_ENV_CONFIG_KEY)_CONFIG=../../$(TEST_CONFIG_DIR) go run migrate.go up

.PHONY: test
test: bin/capa5srv
	$(MAKE) test_db_reset
	go test -v -count=1 ./...

.PHONY: dev_srv_run
dev_srv_run: bin/capa5srv
	$(BASE_ENV_CONFIF_KEY)_CONFIG=./$(DEV_CONFIG_DIR) bin/capa5srv run

.PHONY: dev_db_up
dev_db_up:
	cd tools/migrate/ && $(BASE_ENV_CONFIG_KEY)_CONFIG=../../$(DEV_CONFIG_DIR) go run migrate.go up

.PHONY: dev_db_down
dev_db_down:
	cd tools/migrate/ && $(BASE_ENV_CONFIG_KEY)_CONFIG=../../$(DEV_CONFIG_DIR) go run migrate.go down

.PHONY: dev_db_reset
dev_db_reset:
	cd tools/migrate/ && $(BASE_ENV_CONFIG_KEY)_CONFIG=../../$(DEV_CONFIG_DIR) go run migrate.go drop
	cd tools/migrate/ && $(BASE_ENV_CONFIG_KEY)_CONFIG=../../$(DEV_CONFIG_DIR) go run migrate.go up

.PHONY: clean
clean:
	rm -rf bin/