# Bookmarklets readme

channels_bookmarklet.js runs on the **_Youtube Homepage_** (but press "Show more" under your subscriptions so they are all on the screen). It grabs the channels you are subscribed to and sends them to your entered server in the following format:

```json
{
  "channels": [
    "https://youtube.com/@channel1",
    "https://youtube.com/@channel2",
    "https://youtube.com/@channel3"
  ]
}
```

watch_later_bookmarklet.js runs on your watch later playlist page (www.youtube.com/playlist?list=WL) and grabs the video URLs for all videos in Watch Later.

# Setup

- Run build_bookmarklets.sh to minify the javascript in channels_bookmarklet.js and watch_later_bookmarklet.js (the actual bookmarklet will not work if there are newlines)
- Create a bookmark in your browser and set its URL to the content of channels_bookmarklet.txt/watch_later_bookmarklet.txt
- Navigate to the youtube homepage or watch later page, run the desired bookmarklet, enter your server URL if you haven't before, and viola
