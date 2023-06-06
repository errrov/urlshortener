# URL shortener for Ozon internship

Created URL shortener service using Golang, Echo, Postgresql and Docker.

## How to launch 

```docker-compose -f docker-compose.yml up```

## How to change storage type

Using in-memory storage is by default. To change to Postgresql use flag 
- ```-Memory_type psql ```

To change type of memory in Dockerfile : 
- ```CMD ["/app/main", "-Memory_type","psql"]``` <--> ```CMD ["/app/main"]```


## Identifier algorithm 

If identifier is passed with Request and it is not taken, this identifier will be short link for passed URL.

If identifier is not passed, random generated :

- New UUID is generated
- UUID is converted to base63 (A-Z a-z 0-9 _)
- Result is the new Identifier for passed URL

## Concurrency safety

For Concurrency safety I used :

- In-memory storage : sync.map
- Postgresql : pgxpool



