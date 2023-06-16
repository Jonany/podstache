<script lang="ts">
    import type { Feed } from "../types";
    export let feed:Feed;
</script>

Added: {feed.addedOn.toLocaleString()} - {feed.url}