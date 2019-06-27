# Hypatos Mail

Golang app for sending email via different email providers.

This API is using different email providers in the order in which they were initialized in `initService()` function inside`main.go` file.
#### How it works
Main goal of this app is to try to send message using first email provider. In case of failure the error is logging and app is trying to send message using another email provider(without returning error to user). If every provider returned error the last error is send to end user.

### Providers

By default app implements 2 mail providers:
 - `MailGun` via API, using MailGun API Client
 - `SendGrid` via API, using SendGrid API Client
 
#### Adding Providers
To add a new provider you need to create new type that satisfies the `provider.MailProvider` interface. Than you need to add it to `service` in `initService()` function inside`main.go` file. 

### API Description
This API have one endpoint `/v1/send_mail` which requires `POST` Request.

In case everything is ok and mail was successfully sent API returns empty response with status code `202`. Otherwise you'll get response with error message and appropriative status code.
##### Request Structure:
```json
{ 
    "from": {
        "name": "Tomasz Nowak", 
        "email": "tomasz.nowak@example.com" 
    },
    "subject": "Some Subject",
    "text": "This is a Message Body",
    "to": {
        "name": "Tomasz Nowak",
        "email": "tomasz.nowak@example.com"
    }
}
```
Every fields except `name` are required.

### Config
Before run you need to set 3 env variables in your dockerfile, or in your local environment if you don't want to use docker.
 - `MAIL_GUN_DOMAIN` Your MailGun domain.
 - `MAIL_GUN_API_KEY` Your API Key for MailGun
 - `SENDGRID_API_KEY` Your API Key for SendGrid

### Run
There is Dockerfile so you can build image eg.
```
docker build -t hypatos .
```
And run it:
```
docker run -p 8080:8080 hypatos
```

### Test
As there is no test file in root directory you can run tests using command:
```
go test ./...
```
### TODO:
- Logs - Right now app is sending log to console but in real world example the log place should be determined by the infrastructure.
- Metrics - There is no metrics but as above in real example metrics(and logs) should exist end be kept in some better place than stdout, eg. Kibana.