# Nginx reverse proxy

The point here is be able to use Nginx as a reverse proxy with an independet
docker for the frontend and two different microservices saying hello and bye
bye.

I create a total of four dockers:

- two for the microservices hello and bye bye in golang using
  [mux](http://www.gorillatoolkit.org/pkg/mux)
- one for the front end using nginx to serve the static files
- one more nginx to use as a reverse proxy

The nginx configuration used here is the minimal options to use the reverse proxy
passing the host and the IP in the header to the microservice.

The frontend, only serve an `index.html` static file with links to the
hello and bye bye microservices.

To run everything you only need docker and docker compose installed.

1. Clone the repo
2. Build the images: `docker-compose build`
3. Run the containers: `docker-compose up -d`
4. Open http://localhost:3000 on your browser.
