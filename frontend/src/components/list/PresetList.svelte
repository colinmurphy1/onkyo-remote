<script>
    import { createEventDispatcher } from "svelte";

    // Key and value for list options
    export let preset, frequency, band;

    // Band type (MHz or KHz)
    let bandTyp, presetInt

    // Create dispatcher
    const dispatch = createEventDispatcher();

    // Pass key and value to parent component

    const handleListSelection = () => {
        dispatch('selection', {
            preset
        })
    }

    // 1: FM 2: AM
    if (band == 1) bandTyp = "MHz"
    else if (band == 2) bandTyp = "KHz"

    // Convert preset from hexidecimal to decimal
    presetInt = parseInt(preset, 16)

</script>

<div class="py-1 px-2 bg-gray-100 text-black hover:bg-blue-500 hover:text-white cursor-pointer border border-gray-300" on:click={handleListSelection}>
    {presetInt}
    <span class="float-right">
        {frequency} {bandTyp}
    </span>
</div>