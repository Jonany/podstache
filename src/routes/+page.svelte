<script lang="ts">
    let submittedUrls:string[] = ['http://sdsds'];
    let url:string = "";

    function addUrl(): any {
        submittedUrls = [...submittedUrls, url];
        url = "";
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
            required bind:value={url} />
        <button type="submit"
            class="p-1 border-solid border-2 border-slate-500">Add Feed</button>
    </form>

    {#if submittedUrls.length > 0}
        <h3 class="pt-10 text-2xl">Submitted URLs</h3>
    
        <ol class="pt-5 text-lime-800">
            {#each submittedUrls as submittedUrl}
                <li class="list-item">{submittedUrl}</li>
            {/each}
        </ol>
    {/if}
</main>