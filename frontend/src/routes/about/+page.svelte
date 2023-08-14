<script>
    const status = fetch("/hello.json")
        .then((r) => r.json())
        .catch((err) => {
            throw new Error(err);
        });
</script>

<svelte:head>
    <title>About</title>
</svelte:head>

<h1>About</h1>

<p>
    <strong>server respond</strong>:
    {#await status}
        loading
    {:then data}
        {data.message}
    {:catch err}
        failed to load data
    {/await}
</p>

<p>This is about page</p>

<style>
    h1 {
        color: green;
    }
</style>
