# Project Description
### Alicia Thoney and Adeline Reichert
### November 15, 2022

We aim to implement a web server, database, and web page to demonstrate SQL injection attacks (and countermeasures). We will create a vulnerable and defended version of our website. Our vision for the project is to spin up a locally hosted Go web application, using the GIN library, including a web page with an authentication page and a page where a user can click on links that will send them to urls that query various portions of the database. Our backend will consist of a database that stores usernames and passwords as well as uses SQL code from the url to query other portions of the database (think, click on this link, run ‘SELECT * FROM Pictures.cats’).  
In terms of creating a more secure site version, we plan on investigating the SQL cleansing functionality included with Gin as well as potentially bringing in an outside library to cleanse any form data we have interacting with our database. In terms of fixing the button url routing vulnerabilities, we will investigate not pulling SQL code directly from the url or at least heavily cleansing it if we decide to still go that route. Finally, we plan on storing hashes of passwords rather than just passwords themselves in the database to add another layer of security (and because it’s always a terrible idea to store plaintext passwords in databases).  
For our expected results, we plan to have a demo-able secure and vulnerable site demonstrating SQL injection attacks. For the vulnerable site, this would include giving the “user” (attacker) the ability to drop a table from either our login form or just directly from the url. Our secure site would defend against these attacks.