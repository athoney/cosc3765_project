# Set-up Instructions
1. Download and set-up PostgreSQL
2. Open PostgreSQL shell (Psql)
   - In Linux:
      1. Open terminal
      2. $ `sudo -i -u <username>` (default username is postgres)
      3. $ `psql`
      4. $ `CREATE DATABASE vuln;`
      5. $ `CREATE DATABASE secure;`
      6. Run $`\l` to ensure your databases show up
      7. Exit psql $`\q`
      8. Exit postgres $`exit`
3. Set-up `.env` file
    1. Rename `sample.env` to `.env`
    2. Fill in postgres user information
4. Start sever (In server directory $`go run server.go`)
    1. Open psql again
    2. Run `c\ vuln` - This connects you to the vuln database
    3. Your prompt should now say: `vuln=# `
    4. Query the database: `SELECT * FROM <table>;` (users or contacts)
    5. You should see data in the terminal