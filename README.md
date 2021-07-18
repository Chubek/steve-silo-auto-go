# steve-silo-auto-go

Scrape silo sites. Made specially for Steve. If you're not Steve, you can still use it! But provide your own API key.

## How to Run:
1. Download the compiled binary from the [releases section](https://github.com/Chubek/steve-silo-auto-go/releases/tag/v1.0).
2. Extract it.
3. In the root folder, create a file called `.env` that contains these lines:
```
API_KEY=<API KEY, ASK ME OR GET IT YOURSELF FROM https://ocr.space/OCRAPI>
DATABASE_LOC= <Path to a .db file such as I:\my_db.db>
```
4. Run the application. Don't run it often, every half an hour is good. These sites impose a limit on logins. Don't overuse it!
5. Use an SQLite viewer such as DBeaver to view the database.
