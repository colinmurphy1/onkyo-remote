<script>
    import { PlayIcon, PauseIcon, StopCircleIcon, SkipForwardIcon, SkipBackIcon } from 'svelte-feather-icons'
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


<button class="p-2 bg-gray-200 text-blue-400 hover:bg-blue-500 hover:text-white font-semibold align-sub" on:click={buttonAction}>
    {#if action == "play"}
        <PlayIcon class="inline-block h-full align-bottom" size="2x" />
    {:else if action == "pause"}
        <PauseIcon class="inline-block h-full align-bottom" size="2x" />
    {:else if action == "stop"}
        <StopCircleIcon class="inline-block h-full align-bottom" size="2x" />
    {:else if action == "trdn"}
        <SkipBackIcon class="inline-block h-full align-bottom" size="2x" />
    {:else if action == "trup"}
        <SkipForwardIcon class="inline-block h-full align-bottom" size="2x" />
    {:else}
        Unknown
    {/if}
</button>
