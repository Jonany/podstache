import { writable } from 'svelte/store';
import type { Feed } from './types';

function createFeeds() {
    const { subscribe, set, update } = writable<Feed[]>();

    return {
        subscribe,
        add: (feed:Feed) => update(feeds => feeds ? [feed, ...feeds] : [feed])
    };
}
export const feeds = createFeeds();