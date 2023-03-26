# Go example Google Cloud Function

Result: https://deriabinhellohttp-7u7jkxiv6a-ey.a.run.app/

## Local development

For local development, should use the [Cloud Functions Framework](https://github.com/GoogleCloudPlatform/functions-framework-go).
```shell
# start function framework and serving locally
$ go run cmd/app/main.go
```

## Deploy

### Build deployable container

Information about variables and flags can be found in docs [Building a Go application](https://cloud.google.com/docs/buildpacks/go).

```shell
# build image with pack and gcr.io/buildpacks/builder
$ pack build --builder gcr.io/buildpacks/builder \
  --env GOOGLE_FUNCTION_SIGNATURE_TYPE=http \
  --env GOOGLE_FUNCTION_TARGET=HelloHTTP \
  gocloudfunction 
    
# run container locally
$ docker run --rm --memory-reservation 128M -p 8080:8080 --name gcfunction gocloudfunction
```

### Deploy to Google Cloud Platform

```shell
$ gcloud functions deploy --gen2 --region=europe-west3 \
  --allow-unauthenticated --runtime=go120 --memory=128Mi \
  --trigger-http --entry-point=HelloHTTP \
  DeriabinHelloHttp
```

### Deploy on push to GitHub

Can be founded in [GitHub Actions](.github/workflows/deploy.yml) file.