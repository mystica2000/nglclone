# nglclone

## Installation and Setup Instruction


### Server
- Uses google oauth and cloudinary for storing images.

- google oauth
   - get clientID, clientSecret from google developer console by creating a new project
   - clientID looks like ```generatedcode.apps.googleusercontent.com```
- cloudinary
   - get API Environment Variable from the dashboard after creating a account there.
   - looks like ```cloudinary://yourvariable``` on the dashboard

create ```env.json``` inside the server folder and add the env variables to it.
```
{
  "clientID":"GOOGLE_CLIENT_ID",
  "clientSecret":"CLIENT_SECRET",
  "signKey":"secret",
  "CLOUDINARY_URL": "API_Environment_variable"
}
```

- go mod download

- Install SQLite

### Client
- run ```npm install```


## Database Schema
