<script>
    import { createEventDispatcher } from "svelte";
    import { VolumeXIcon, Volume2Icon } from 'svelte-feather-icons'

    export let volume


    let showVolume = false

    // Create dispatcher
    const dispatch = createEventDispatcher()

    const handleShowVolume = () => {
        showVolume = ! showVolume
        dispatch('showvolume', showVolume)
    }

</script>

<button class="py-1 px-2 bg-gray-200 {volume.Mute ? "text-red-500" : "text-black"} hover:bg-blue-500 hover:text-white" on:click={handleShowVolume}>
    {#if volume.Mute}
        <!-- Show mute icon -->
        <VolumeXIcon class="inline-block h-full align-bottom" />
    {:else}
        <!-- Show regular volume icon -->
        <Volume2Icon class="inline-block h-full align-bottom" />
    {/if}
    <span class="h-full">
        {volume.Mute ? "Mute": volume.Level}
    </span>
</button>
