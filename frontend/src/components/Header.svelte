<script>
	import PowerButton from './buttons/Power.svelte'
    import SourceButton from './buttons/Source.svelte'
    import VolumeButton from './buttons/Volume.svelte'
    import SourceSelect from './views/SourceSelect.svelte'
    import VolumeSelect from './views/VolumeSelect.svelte'

	let showSourceList = false
    let showVolume = false

	// Event handler for showing source
	const handleSourceList = (event) => {
        showVolume = false
        showSourceList = event.detail
	}

    // Event handler for volume slider
	const handleVolume = (event) => {
        showSourceList = false
        showVolume = event.detail
	}

    export let status, sources
</script>

<header class="px-0.5 bg-gray-200 border-b border-b-gray-300">
    <div class="max-w-4xl mx-auto relative">
        {#if status.Power.Status}
            <div class="absolute left-0 flex flex-row h-full">
                <SourceButton on:sourcelist={handleSourceList}/>
                <VolumeButton volume={status.Volume} on:showvolume={handleVolume} />
            </div>
        {/if}
    
        <!-- Power button -->
        <div class="absolute right-0 flex flex-row h-full">
            <PowerButton pwrStatus={status.Power.Status} />
        </div>
        <!-- Power button end -->
        <div class="text-center font-semibold py-1">
            {#if status.Power.Status == false}
                &nbsp;
            {:else}
                <!-- If the source is NET, display the NET source (if it is provided) -->
                {#if status.Input.HexCode == "2B" && status.Input.NetSource != "" }
                    {status.Input.NetSource}
                {:else}
                    {status.Input.Name}
                {/if}
            {/if}
        </div>    
    </div>
</header>

<div class="mt-1 mb-4 max-w-4xl mx-auto">
    {#if showSourceList }
        <SourceSelect sources={sources} on:sourcelist={handleSourceList} />
    {/if}
    {#if showVolume }
        <VolumeSelect volume={status.Volume} on:showvolume={handleVolume} />
    {/if}
</div>
