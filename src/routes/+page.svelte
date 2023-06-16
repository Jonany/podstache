<script lang="ts">
    import { feeds } from "../stores";
	import { Feed } from "../types";
    let feedUrl:string = "";

    function addUrl(): any {
        feeds.add(new Feed(feedUrl));
        feedUrl = "";
    }
</script>
  
<style lang="postcss">
    :global(html) {
        background-color: theme(colors.gray.100);
    }
</style>

<main class="px-4">
    <h1 class="text-3xl font-bold">
        Welcome to Podstache  🥸
    </h1>
    
    <h3 class="pt-4 text-2xl">Add Podcast Feed</h3>
    <form on:submit={addUrl} class="pt-4">
        <input type="url"
            class="p-1 border-solid border-2 border-slate-500"
            placeholder="Enter URL Here"
            required bind:value={feedUrl} />
        <button type="submit"
            class="p-1 border-solid border-2 border-slate-500">Add Feed</button>
    </form>

    <h3 class="pt-10 text-2xl">Submitted URLs</h3>

    {#if $feeds}
        <ol class="pt-5 text-lime-800">
            {#each $feeds as feed}
                <li class="list-item">Added: {feed.addedOn.toLocaleString()} - {feed.url}</li>
            {/each}
        </ol>
    {:else}
        <h4 class="pt-5 text-lime-800">Nothing here...yet 😏</h4>
    {/if}

</main>