# Infrastructure Devops Technical Challenge

## API

Inside the [API](/api) folder resides the Go Code for a simple API REST that gatherers air quality measurements from a PostgreSQL database and serve them in a well-structured JSON.

To populate the database with measurements, the API also serves as an entry point, you can pass a filename as an argument to be parsed and inserted into the database. The supported format is CSV and a sample file is provided. See [environment_airq_measurand.csv](environment_airq_measurand.csv)

### Dockerfile

The provided [Dokerfile](Dockerfile) will install all the needed dependencies to build the API and then remove all the unnecessary elements to maintain a smaller image file using a multi-staged build.

It will also install "postgresql-client" to access the tool [pg_isready](https://www.postgresql.org/docs/9.3/app-pg-isready.html)

### docker-compose

There is also a [docker-compose](docker-compose.yml) environment that will set up a Postgres database and run a one-time migration container to populate the database with measurements before starting the API.

## CI (Continuous Integration)

Inside the [workflows](.github/workflows) directory there is a Github Actions file called [deployment.yml](.github/workflows/deployment.yml) that provide all the necessary instructions to deploy the API on Google Cloud (had to shut down the cluster because it was getting expensive).

## Deployment


I have chosen Kubernetes as the Deployment solution because it is the one I know. All the needed configuration to deploy a working API resides inside the [deployment](/deployment) folder. This config has been created with [kompose](https://kompose.io), a conversion tool to go from Docker Compose to Kubernetes, and then modified by hand to accomplish the desired behavior. The source file used as a template is [docker-compose.kubernetes.yml](docker-compose.kubernetes.yml)

**Note:** Normally you should not commit your secrets, I did it just to keep record of everything needed.

### Cache

[Varnish](https://varnish-cache.org/) was my go-to for caching as it is completely independent of the code so there was no need to modify it, as with redis or similar you have to manually call the cache service inside the code.

## Logs and Monitoring

As I used Google Cloud Engine, I had all the tools needed for seeing logs and monitoring the state of pods.
