
.PHONY:
default: main consul info test 

main.go: main


main: 
	go build -o go_consul_template_integration


.PHONY: consul
consul:
	./build_consul


.PHONY: clean_docker
clean_docker:
	docker rm -f $$(docker ps -qa)


.PHONY: test
test:
	docker exec dev-consul consul kv put memcache \
  '{"memcache_servers": ["10.0.0.1"]}'

	docker exec dev-test \
	consul watch -type=key -key=memcache /app/handler.sh &	

	docker exec  dev-test /app/go_consul_template_integration

.PHONY: update_list
update_list:
	docker exec dev-consul consul kv put memcache \
	'{"memcache_servers": ["10.0.0.10", "10.0.2.10"]}'

.PHONY: info
info:
