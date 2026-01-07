/* Javascript bookmarklet to grab your subscribed channels and send to Dumbtube backend */
/* Run this bookmarklet from the youtube HOMEPAGE with your left-side "Subscriptions" sidebar EXPANDED */
/* All comments must be block comments in order for build_bookmarklets.sh to properly convert code to one line*/
(function() {

    const blacklist = ["/@nytimes", "/@fox43"];

    function isVisibleAndInSubscriptions(element) {
        return !element.hasAttribute("hidden") && !blacklist.includes(element.getAttribute("href"));
    }

    alert("Be sure to click 'Show more' on your subscriptions sidebar so that all subscribed channels are shown");

    /* Find all <a> tags that have class yt-simple-endpoint and href attribute starting with /@ */
    const subscribedChannelsElements = document.querySelectorAll('a.yt-simple-endpoint[href^="/@"]');

    const visibleSubscribedChannelElements = Array.from(subscribedChannelsElements).filter(isVisibleAndInSubscriptions);

    const channels = visibleSubscribedChannelElements.map(a => ("https://youtube.com" + a.getAttribute("href")));
    
    console.log(channels);

    /* TODO: send json to endpoint */
})();