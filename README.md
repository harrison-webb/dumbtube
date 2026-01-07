# Dumbtube
The youtube "Subscriptions" feed is one of the last places on the "commercial" internet that is still decent. Unfortunately, accessing it requires dodging endless youtube shorts, recommendations that are irrelevant at best and ragebait at worst, and constant ai slop advertisements.

Dumbtube is a self-hosted project to escape all that and let you enjoy your subscriptions feed in peace. Basically what it does is keeps a list of your subscribed channels, downloads their videos to a local server that you set up, and provides a minimal ui for accessing and watching your untainted subscriptions feed, from the peace and comfort of your own local server.

# Components
Here is a rundown of the project architecture, roughly following the data flow:

## Bookmarklets
The primary bookmarklet lets you pull in your subscribed channels without having to manually add all of them. A secondary booklet serves the purpose of fetching some of the videos in your "Watch Later" playlist and sprinkling them into your Dumbtube feed

## Backend ingestion and endpoint
Receives POST from bookmarklets containing json of channels or watch-later-video urls. Transforms channel URLs to channel RSS feed URLs, gets videos for each channel by querying RSS feeds, and inserts channels + videos into database. Periodically checks channels for new videos

## Database
Schema draft that chatgpt made, still refining:

channels
- channel_id (pk)
- title
- rss_url
- last_checked_at

videos
- video_id (pk)
- channel_id (fk)
- title
- published_at
- watch_url
- source (rss | watch_later)
- downloaded (bool)
- file_path (nullable)

## Download worker
Async downloads videos from the database to server

## Frontend
Super simple feed just showing you the videos and nothing else
