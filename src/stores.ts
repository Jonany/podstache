import { writable } from 'svelte/store';
import { Feed } from './types';
import savedData from './data.json';

function createFeeds() {
    let defaultFeeds:Feed[] = [];

    if(savedData.feeds) {
        defaultFeeds = savedData.feeds
            .map(savedFeed => new Feed(
                savedFeed.url,
                new Date(Date.parse(savedFeed.addedOn))
            ));
    }

    const { subscribe, update } = writable<Feed[]>(defaultFeeds);

    return {
        subscribe,
        add: (feed:Feed) => update(feeds => feeds ? [feed, ...feeds] : [feed])
    };
}
export const feeds = createFeeds();