build:
		docker-compose up --build

start:
		docker-compose start

stop:
		docker-compose stop

clean:
		docker rm orders nats postgres

.PHONY: build, start, stop, clean