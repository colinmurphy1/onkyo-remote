<script>
    import Icon from 'svelte-awesome'
    import play from 'svelte-awesome/icons/play';
    import pause from 'svelte-awesome/icons/pause';
    import stop from 'svelte-awesome/icons/stop';
    import backward from 'svelte-awesome/icons/backward';
    import forward from 'svelte-awesome/icons/forward';

    export let action

    const buttonAction = async () => {
        // Make an API request to control the network playback
        try {
            const req = await fetch("/api/net/" + action)
            if (req.status != 200) {
                console.log("got bad response code", req.status)
            }
        } catch(err) {
            console.log("error:", err)
        }
    }
</script>


<button class="rounded p-2 bg-blue-400 hover:bg-blue-500 text-white font-semibold" on:click={buttonAction}>
    {#if action == "play"}
        <Icon data={play} scale="1.5" class="align-middle m-1" />
    {:else if action == "pause"}
        <Icon data={pause} scale="1.5" class="align-middle m-1" />
    {:else if action == "stop"}
        <Icon data={stop} scale="1.5" class="align-middle m-1" />
    {:else if action == "trdn"}
        <Icon data={backward} scale="1.5" class="align-middle m-1" />
    {:else if action == "trup"}
        <Icon data={forward} scale="1.5" class="align-middle m-1" />
    {:else}
        Unknown
    {/if}
</button>
