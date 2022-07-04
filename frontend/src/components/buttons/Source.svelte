<script>
    import { createEventDispatcher } from "svelte";
    import { ListIcon } from 'svelte-feather-icons'

    // Create dispatcher
    const dispatch = createEventDispatcher();

    let sourceList = false

    const handleSourceList = () => {
        sourceList = ! sourceList
        dispatch('sourcelist', sourceList)
    }

</script>
<!-- {sourceList ? 'bg-blue-500 text-white font-semibold' : 'bg-gray-200 text-black'} -->
<button class="py-1 px-2 hover:bg-blue-500 hover:text-white" on:click={handleSourceList}>
    <ListIcon class="inline-block h-full" />
    <span class="h-full inline-block align-middle">
        Source
    </span>
</button>
