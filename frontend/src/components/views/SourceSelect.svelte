<script>
    import { onMount, createEventDispatcher } from "svelte";
    import ListOption from "../list/ListOption.svelte"

    // Create dispatcher
    const dispatch = createEventDispatcher();

    let sources, sourcesLoaded = false;

    // Handle source selection
    const handleSelection = async (event) => {
        const inputCode = event.detail.key
        // Make an http request
        const req = await fetch("/api/source/" + inputCode, {
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


    // API call to get all available inputs
    const getSource = async () => {
        const req = await fetch("/api/source", { method: "GET" })
        .then(response => {
            if (!response.ok) {
                throw new Error("Bad network response")
            }
            return response.json()
        })
        .then(data => data)

        return req.data
    }

    onMount(async () => {
        // Get sources
        try {
            sources = await getSource()
        } catch (err) {
            console.log("error loading sources:", err)
        }
        sourcesLoaded = true
    })

</script>

{#if sourcesLoaded}
<div class="grid grid-cols-2 lg:grid-cols-4 gap-1">
    {#each Object.entries(sources) as [id, name]}
        <ListOption key={id} value={name} on:selection={handleSelection} />
    {/each}
</div>
{/if}
