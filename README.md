# Geolocation
This project aims to create and manage a geolocation database and allow access to it through an api endpoint.

## Directories
- [.github](.github/) holds the ci workflows
- [cmd](cmd/) contains application code for geolocation api and importer task
- [deployment](deployment/) contains kubernetes manifests in order to create pods, services, cronjobs...
- [pkg](pkg/) contains geolocation package with logic for import and services to interact with geolocation data

## Runing

### Local
In order to run localy we need to execute the following command:
```bash
> docker-compose up -d
```
The importer will fill the database with data contained in file [cmd/importer/data_dump.csv](data_dump.csv)

Fetch geolocation data using the exposed api:

```bash
> curl --location --request GET 'localhost:8080/api/geolocations?ipaddress=200.106.141.15'
```

### Deploy into Kubernetes
When the code is pushed to main branch the Github action with jobs setup-build-publish-deploy-geolocation-api and 
setup-build-publish-deploy-importer-task will be executed.

setup-build-publish-deploy-geolocation-api is responsible for deploying the geolocation api and create a service that 
grants access to it from outside.

setup-build-publish-deploy-importer-task will create a cronjob responsible for populating the csv file with data from a 
csv file located in a goolge storage bucket.

we can fetch geolocation data from production using the following command:

```bash
> curl --location --request GET 'http://35.227.21.172:80/api/geolocations?ipaddress=181.231.183.23'
```