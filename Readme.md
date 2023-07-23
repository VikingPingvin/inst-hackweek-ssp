# HackWeek Project
## Api Gateway Self Service Portal
A study project to demonstrate a partial set of usecases for the Instructure Api Gateway Self Service Portal using:
- HTMX
- Alpine.Js
- Golang backend
- TODO: Redis/Elasticache database

### Motivation:

- I do not like JavaScript
- I like creating somewhat interactive frontends
- I have no idea how to create good UI/UX
- Try to recreate a low-effort look and feel of InstUI

### Features
- Fetch list of available APIs on page load
- SSR with htmx together with Go-s `template` package to use HTML fragment views
- Click event on API buttons to fetch that particular APIs data from DB (local dummy data)
- Use Alpine.Js to handle state
- TODO: Use Go's embed.FS making the whole webapp a single binary, 
no more bundling of random html or static etc... files
- Detect unsaved changes in api config and disallow navigation

### Known issues
- `isRowMatching()` function in `apiSettingComponent.html` s script block is only accessible upon a second contentView update