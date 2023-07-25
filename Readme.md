# HackWeek Project
## Api Gateway Self Service Portal
A study project to demonstrate a partial set of usecases for the Instructure Api Gateway Self Service Portal using:
- HTMX
- TailWindCSS
- Alpine.Js
- Golang backend
- TODO: Redis/Elasticache database

### Motivation:

- I do not like JavaScript
- I like creating somewhat interactive frontends
- I have no idea how to create good UI/UX
- Try to recreate a low-effort look and feel of InstUI

### Features
- Use Alpine.Js to handle state
- SSR with htmx together with Go-s `template` package to use HTML fragment views
- Fetch list of available APIs on page load, based on Google OAuth2 login cookie
- Click event on API buttons to fetch that particular APIs data from DB (local dummy data)
- Edit button to toggle editable settings table
- Submit button to submit table fields as Forms
- Detect unsaved changes in api config and disallow navigation
- TODO: Use Go's embed.FS making the whole webapp a single binary, 
no more bundling of random html or static etc... files

### Known issues
- `isRowMatching()` function in `apiSettingComponent.html` s script block is only accessible upon a second contentView update