# nglclone
   Clone of https://ngl.link/
   
   - create account with oauth (for example: username as mystica) 
   - link will be generated (test.com/mystica)
   - share the link on socials. 
   - anyone with that link will have form(text area, submit)
   - people with the link can able to send text message
   - users can able to view the text messages people sent, and mark as replied/delete/download them as image
   - users can enable/disable the link

## Installation and Setup Instruction

### Server
Uses
-  google oauth 
   - get clientID, clientSecret from google developer console by creating a new project
   - clientID looks like ```generatedcode.apps.googleusercontent.com```
- cloudinary
   - for storing images (profile picture).
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
![nglclone](https://user-images.githubusercontent.com/45729256/219892477-75653890-c5e2-4773-aebe-d91cee4b2299.png)

