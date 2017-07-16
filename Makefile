all:
	go build src/diamond_pony.go

run:
	go run src/diamond_pony.go

clean:
	rm -f diamond_pony
