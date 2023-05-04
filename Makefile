
test:
	+docker-compose -f dockerfiles/test/docker-compose.yml up
	
clear-test:
	+docker-compose -f dockerfiles/test/docker-compose.yml down -v --rmi all --remove-orphans

dev:
	+docker-compose -f dockerfiles/census3/docker-compose.yml up

clear-dev:
	+docker-compose -f dockerfiles/census3/docker-compose.yml down -v --rmi all --remove-orphans