# Makefile
CURRENT_DIR := $(shell pwd)
DBURL = 'postgres://postgres:0509@localhost:5432/nt?sslmode=disable'

mig-up:
	@echo "Migrating up"
	migrate -path migrations -database ${DBURL} -verbose up

mig-down:
	@echo "Migrating down"
	migrate -path migrations -database '$(DBURL)' -verbose down

mig-create:
	@echo "Creating new migration"
	migrate create -ext sql -dir migrations -seq create_table


mig-insert:
	@echo "Inserting database"
	migrate create -ext sql -dir migrations -seq inserting_data


mig-force:
	@echo "Migrating down"
	migrate -path migrations -database '$(DBURL)' -verbose force 1	

mig-goto:
	@echo "Migrating down"
	migrate -path migrations -database '$(DBURL)' -verbose goto 2

mig-run:
	go run main.go
