/* Javascript bookmarklet to grab your subscribed channels and send to Dumbtube backend */
/* Run this bookmarklet from the youtube HOMEPAGE with your left-side "Subscriptions" sidebar EXPANDED */
/* All comments must be block comments in order for build_bookmarklets.sh to properly convert code to one line*/
(function() {

    const STORAGE_KEY = "dumbtube_ingest_endpoint";
    const blacklist = ["/@nytimes", "/@fox43"];

    function getEndpoint(){
        let endpoint = localStorage.getItem(STORAGE_KEY);
        if (!endpoint) {
            endpoint = prompt("Enter Dumbtube channel list ingest endpoint url: ");
            if (!endpoint) return null;
            localStorage.setItem(STORAGE_KEY, endpoint);
        }
        return endpoint;
    }

    function isVisibleAndInSubscriptions(element) {
        return !element.hasAttribute("hidden") && !blacklist.includes(element.getAttribute("href"));
    }

    const endpoint = getEndpoint();
    if (!endpoint) return;

    alert("Be sure to click 'Show more' on your subscriptions sidebar so that all subscribed channels are shown");

    /* Find all <a> tags that have class yt-simple-endpoint and href attribute starting with /@ */
    const subscribedChannelsElements = document.querySelectorAll('a.yt-simple-endpoint[href^="/@"]');

    const visibleSubscribedChannelElements = Array.from(subscribedChannelsElements).filter(isVisibleAndInSubscriptions);

    const channels = visibleSubscribedChannelElements.map(a => ("https://youtube.com" + a.getAttribute("href")));
    
    console.log(channels);

    /* send json to endpoint */
    fetch(endpoint, {
        method: "POST",
        mode: "no-cors",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ channels })
    });
})();