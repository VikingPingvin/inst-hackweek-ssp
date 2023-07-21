# HackWeek Project
## Api Gateway Self Service Portal
A study project to demonstrate a partial set of usecases for the Instructure Api Gateway Self Service Portal using:
- HTMX
- Alpine.Js
- Golang backend
- Redis/Elasticache database

### Motivation:

- I do not like JavaScript
- I like creating somewhat interactive frontends
- I have no idea how to create good UI/UX
- Try to recreate a low-effort look and feel of InstUI

### Features
- Fetch list of available APIs on page load
- Click event on API buttons to fetch that particular APIs data from DB (local dummy data)
- Use Alpine.Js to handle state
- TODO: Detect unsaved changes in api config and disallow navigation