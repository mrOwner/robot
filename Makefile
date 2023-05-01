# change this.
# 			↓ 	mysql or postgres
DB_USED := postgres

MIGRATIONS_PATH := db/${DB_USED}/migrations
SQLC_PATH := db/${DB_USED}/sqlc

.PHONY: default
default: help

##@ General

help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9\/-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Run

candles-grab: ## Run stock candles grabber.
	@go run ./cmd/robot/main.go grab
.PHONY: candles-grab

##@ Database (if you need to change RDMS, check DB_USED variable in the makefile).

migration/create: ## Create the migration sql files.
	@read -p 'Enter migration name: ' name; \
	atlas migrate diff $$name --env ${DB_USED}
.PHONY: migration/create

migration/apply: ## Apply the migration into the database.
	@atlas schema apply --env ${DB_USED} --exclude "atlas_schema_revisions,dev"; 	\
	mig=$$(ls -Art ${MIGRATIONS_PATH}/*.sql | tail -n 1 | grep -Eo '[0-9]+');		\
	atlas migrate set $$mig --env ${DB_USED}
.PHONY: migration/apply

migration/status: ## Check status of the migration.
	@atlas migrate status --env ${DB_USED}
.PHONY: migration/status

# Example for revert to a specific migration. But I had problem with it.
#		 atlas schema apply \
#		   --url "mysql://root:pass@localhost:3306" \
#		   --to "file://db/mysql/migrations?version=20230305153309" \
#		   --dev-url "mysql://root:pass@localhost:3306/dev" \
#		   --exclude "atlas_schema_revisions,dev"
migration/revert-all: ## Revert all of migrations.
	@atlas schema clean --env ${DB_USED}
.PHONY: migration/revert-all

migration/hash: ## Update migration hash (if you manual change it).
	@atlas migrate hash --env ${DB_USED}
.PHONY: migration/hash

migration/set: ## Set a vesrion of migration.
	@echo 'Please select a migartion file.';					\
	n=0;														\
	for file in ${MIGRATIONS_PATH}/*.sql; 						\
	do 															\
		n=$$((n+1)); 											\
		printf "[%s] %s\n" $$n $$file; 							\
		eval "files$${n}=\$$(echo $$file | grep -Eo '[0-9]+')"; \
	done; 														\
	if [ $$n -eq 0 ];											\
	then														\
		echo >&2 No migartion files found.;						\
		exit;													\
	fi;															\
	printf 'Enter File Index ID (1 to %s): ' $$n;				\
	read -r num; 												\
	num=$$(echo $$num | grep -Eo '[0-9]+'); 					\
	if [[ $$num -le 0 || $$num -gt $$n ]]; 						\
	then 														\
		echo >&2 Wrong selection.;								\
		exit 1;													\
	fi; 														\
	eval "mig=\$$files$${num}"; 								\
	atlas migrate set $$mig --env ${DB_USED}
.PHONY: migration/set

##@ Mocks

mocks: ## Generate mocks.
	@go generate ./...; \
	mockgen -package mock -destination ./mocks/querier.go github.com/mrOwner/robot/db/postgres/sqlc Querier

.PHONY: mocks

##@ Generates

sqlc: ## Generate sqlc files.
	@rm -rf ${SQLC_PATH};	\
	sqlc generate
.PHONY: sqlc

buf: ## Generate code by proto files.
	@rm -rf ./clients/grpc/proto/* ; \
	buf generate
.PHONY: buf

##@ Install

install/atlas: ## Install the atlas.
	@curl -sSf https://atlasgo.sh | sh
.PHONY: install/atlas

install/sqlc: ## Install the sqlc.
	@go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
.PHONY: install/sqlc
