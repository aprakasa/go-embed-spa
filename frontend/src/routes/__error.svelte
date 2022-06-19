<script context="module">
  export function load({ status, error }) {
    return {
      props: { status, error },
    };
  }
</script>

<script>
  export let status, error;
</script>

<svelte:head>
  <title>{status}</title>
</svelte:head>

<h1>{status}</h1>

<p>{error.message}</p>

<style>
  h1 {
    color: crimson;
  }
</style>
