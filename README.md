[![CircleCI](https://circleci.com/gh/PlagaMedicum/tournament/tree/master.svg?style=svg)](https://circleci.com/gh/PlagaMedicum/tournament/tree/master)
[![codecov](https://codecov.io/gh/PlagaMedicum/tournament/branch/master/graph/badge.svg)](https://codecov.io/gh/PlagaMedicum/tournament)

# Social tournament service

This is a RESTful web server with a
[social tournament service](https://gist.github.com/sashayakovtseva/ed84bb13fbdfd8ef43bf0229108ace78)
implementation.  
Created within the _ITechArt Golang Students Lab_.

### Deployment
If you want to deploy this application both locally or in GCloud,
you need to run the `./deploy.sh` script with flags.  
See `./deploy.sh --help` for more information.

To deploy project in GCloud follow these steps:  
First of all you will need to set the `PROJECT_ID` variable. Check 
[this](https://cloud.google.com/resource-manager/docs/creating-managing-projects#identifying_projects)
for more information.  
Then you need to run following command in your terminal:
```bash
export PROJECT_ID=PASTE_YOUR_PROJECT_ID_HERE
```
After this you'll need to create `google_credentials.json` file. Check
[this](https://cloud.google.com/docs/authentication/production#obtaining_and_providing_service_account_credentials_manually)
for more information.

Finally, following command will build all of project's images and create a cluster with deployment:  
```bash
./deploy.sh up
```
This will take some time, so you can spend it making a cup of hot tea/coffee (choose one).

If everything ready, you can try to access the server.
To get the external ip of server run:
```bash
kubectl describe service server
```
and copy the value of `LoadBalancer Ingres: `.

Then you can try to test some endpoints:
* Create user:
```json
POST http://PASTE_SERVICE_IP_HERE/user
BODY:

{
    "name": "Daniil Dankovskij"
}

Request BODY:

{
    "id": 1
}
```
* Get User:
```json
GET http://PASTE_SERVICE_IP_HERE/user/1
BODY:

NONE

Request BODY:

{
    "id": 1,
    "name": ​"Daniil Dankovskij",
    "balance": 700
}
```
If you want to cleanup your GCloud, just run simple command:
```bash
./deploy.sh annihilate
```
