<script>
    import PresetList from "../list/PresetList.svelte"

    // Get status from receiver
    export let status;

    let frequency, band, preset

    $: {
        if (status.Input.HexCode == "24") {
            // If FM, add a period to the freqency and add MHz
            band = "MHz"
            // Add a period in the correct spot in the frequency (for example, 93.3 or 100.3)
            if (String(status.Tuner.Frequency).length == 4) {
                // xx.x
                frequency = String(status.Tuner.Frequency).slice(0,2) + "." + String(status.Tuner.Frequency).slice(2,3)
            } else {
                // xxx.x
                frequency = String(status.Tuner.Frequency).slice(0,3) + "." + String(status.Tuner.Frequency).slice(3,4)
            }
        } else {
            // AM
            band = "KHz"
            frequency = status.Tuner.Frequency
        }
    }

    const handlePreset = async (event) => {
        const preset = event.detail.preset
        await fetch("/api/tuner/preset/" + preset, {
            method: "GET",
        })
        .then(response => {
            if (!response.ok) {
                throw new Exception("Response was not OK")
            }
        })
    }
</script>

<div>
    <!-- TUNER Frequency and preset -->
    <div class="text-center">
        <div class="text-4xl font-semibold my-4">
            {frequency}
            <span class="text-gray-700 text-2xl">{band}</span>
        </div>
        <div class="text-gray-700 font-semibold">
            {#if status.Tuner.Preset != 0}
                Preset {status.Tuner.Preset}
            {:else}
                No preset
            {/if}
        </div>
    </div>
    <!-- TUNER Frequency and preset END -->

    <!-- TUNER Preset list -->
    <div class="grid grid-cols-2 md:grid-cols-5 gap-1 my-4">
        {#each Object.entries(status.Tuner.PresetList) as [id, preset]}
            <PresetList preset={id} frequency={preset.Frequency} band={preset.Band} on:selection={handlePreset} />
        {/each}
    </div>
    <!-- TUNER Preset list END -->
</div>
