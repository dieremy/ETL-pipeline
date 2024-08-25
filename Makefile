name = sdesdo
CWD=$(shell pwd)

HOST_HAS_DOCKER := $(shell which docker >/dev/null 2>&1 && echo yes || echo no)

ifeq ($(HOST_HAS_DOCKER), yes)
	@printf "Host has docker engine installed..."
else
	@printf "Host has not docker engine installed..."
	@exit 1
endif

FIGLET_CMD := $(shell which figlet >/dev/null 2>&1 && echo yes || echo no)

ifeq ($(FIGLET_CMD), yes)
  PRINT := figlet
else
  PRINT := printf
endif

.PHONY: all
all:
	@printf "CHECKING VOLUMES ..."
	@mkdir -p $(CWD)/data/postgres
	@mkdir -p $(CWD)/data/grype
ifeq ($(shell docker ps -q), "")
	@$(PRINT) "RE LAUNCH..."
	@docker compose -f ./infrastructure/docker-compose.yml --env-file .env up -d
else
	@$(PRINT) "BUILD UP ..."
	@docker compose -f ./infrastructure/docker-compose.yml --env-file .env up -d --build
endif
	@printf "\n\n"
	@docker ps -a -s
	@printf "\n\n"
	@docker system df

.PHONY: build
build:
	@$(PRINT) "BUILD UP ..."
	@docker compose -f ./infrastructure/docker-compose.yml --env-file .env up -d --build

.PHONY: down
down:
	@$(PRINT) "DROP ..."
	@docker compose -f ./infrastructure/docker-compose.yml --env-file .env down

.PHONY: re
re: down build
	@$(PRINT) "RE-BUILD UP ..."

.PHONY: fclean
fclean: clean
	@$(PRINT) "FORCE CLEAN ALL INFRASTRUCTURE"
	@docker system prune --all --force --volumes
	@docker network prune --force
	@docker volume prune --force

.PHONY: cleanvolumes
cleanvolumes: down
	@$(PRINT) "CLEAN VOLUMES ..."
	@sudo rm -rf $(CWD)/data/
