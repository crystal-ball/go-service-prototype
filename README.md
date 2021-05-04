# Go service prototype

## Running

Live reload is setup with `modd`, install with Homebrew:

```sh
brew install modd
```

Then run the service with `modd` from the repo root:

```
modd
```

The service will automatically restart on changes.

### Routes

- GET /healthcheck
- GET /user/account (get account if authenticated)
- POST /user/account (create a new account)
- POST /auth/signin
- POST /auth/signout
- GET /projects
- GET /projects/{projectID}
- POST /projects/{projectID}
- PUT /projects/{projectID}
- DELETE /projects/{projectID}
