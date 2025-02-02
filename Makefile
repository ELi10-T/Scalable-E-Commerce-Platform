postgres:
	sudo docker rm eCommDB
	sudo docker run --name eCommDB -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=test -d postgres
	sudo docker cp ./init/init_db.sql eCommDB:/docker-entrypoint-initdb.d/init.sql