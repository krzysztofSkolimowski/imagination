ensure_outside_container:
	if [ -f /.dockerenv ]; then echo -n $(RED)This target needs to be run outside server container$(NC)"\n"; false; fi

up: ensure_outside_container
	@docker-compose pull && docker-compose up | grep --line-buffered server_1

rm: ensure_outside_container
	docker-compose down -v

enter: ensure_outside_container
	@docker-compose exec server bash

mycli: SHELL:=/bin/bash
mycli: ensure_outside_container
	@source <(cat .env $(ENVS)) && mycli -h 127.0.0.1 -P $$IMAGINATION_MYSQL_PORT -u $$IMAGINATION_MYSQL_USER -p$$IMAGINATION_MYSQL_PASSWORD
