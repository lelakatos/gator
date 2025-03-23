# Gator - an RSS aggreGator, built for the boot.dev course

## Requirements:
- go, version 1.22+
- Postgres, version 1.14+

## Installation:
```zsh
go install github.com/lelakatos/gator
```

## Setup:
- Create a .gatorconfig.json in the root directory
    - it should contain:
        - "db_url": the url of your postgres database
        - "current_user_name": a username

## Usage:
In the shell run `gator <command> <args>`

Useful commands:
- register: takes 1 argument, the username string
- login: takes 1 argument, the username (should already be registered)
- reset: takes no argument, resets the database
- agg: runs the aggregation, scrapes all the registered feeds - long running
- addfeed: takes 2 arguments, first the name of the feed, second the URL to the RSS feed. Arguments should be in ""-s
- feeds: takes no argument, returns the registered feeds
- follow: takes 1 argument, the URL of the feed, follows the feed provided for the current user
- following: takes no argument, returns all the followed feeds for the current user
- unfollow: takes 1 argument, the URL of the feed, unfollows it for the current user
- browse: takes 1 optional argument, the number of entries to be returned, returns the most recent news from across followed feeds
