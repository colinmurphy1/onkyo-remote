<script>
    import { createEventDispatcher } from "svelte";

    // Key and value for list options
    export let key, value;

    // Create dispatcher
    const dispatch = createEventDispatcher();

    // Pass key and value to parent component

    const handleListSelection = () => {
        dispatch('selection', {
            key,
            value
        })
    }

</script>

<div class="py-1 px-2 bg-gray-100 text-black hover:bg-blue-500 hover:text-white cursor-pointer border border-gray-300" on:click={handleListSelection}>
    {value}
</div>