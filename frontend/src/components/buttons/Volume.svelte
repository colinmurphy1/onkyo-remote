<script>
    import { createEventDispatcher } from "svelte";
    import { Volume2Icon } from 'svelte-feather-icons'

    export let volume


    let showVolume = false

    // Create dispatcher
    const dispatch = createEventDispatcher()

    const handleShowVolume = () => {
        showVolume = ! showVolume
        dispatch('showvolume', showVolume)
    }

</script>

<button class="py-1 px-2 bg-gray-200 text-black hover:bg-blue-500 hover:text-white" on:click={handleShowVolume}>
    <Volume2Icon class="inline-block h-full" />
    <span class="h-full inline-block align-middle">
        {volume.Level}
    </span>
</button>
