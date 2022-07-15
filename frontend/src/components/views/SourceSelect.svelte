<script>
    import { createEventDispatcher } from "svelte";
    import ListOption from "../list/ListOption.svelte"

    export let sources;

    // Create dispatcher
    const dispatch = createEventDispatcher();

    // Handle source selection
    const handleSelection = async (event) => {
        const inputCode = event.detail.key
        // Make an http request
        await fetch("/api/source/" + inputCode, {
            method: "GET",
        })
        .then(response => {
            if (! response.ok) {
                throw new Exception("Response was not OK")
            }
            return response.json()
        })
        .then(data => data)

        dispatch('sourcelist', false)
        return true
    }

</script>

<div class="grid grid-cols-2 md:grid-cols-4 gap-1">
    {#each Object.entries(sources) as [id, name]}
        <ListOption key={id} value={name} on:selection={handleSelection} />
    {/each}
</div>
