- Develop service architecture
  - Factor DB actions into service functions
- Complete authentication flow
- Pagination
- Handle errors gracefully
  - Log errors to a file or a DB table
  - Replace log.Error with a way to return HTTP 5XX response codes
- Move environment variables into .env (e.g. DB credentials)
