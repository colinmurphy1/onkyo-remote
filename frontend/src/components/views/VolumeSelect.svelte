<script>
    import { VolumeXIcon, Volume1Icon, Volume2Icon } from 'svelte-feather-icons'

    // Get volume data
    export let volume;

    let volumeLevel

    // Handle muting
    const handleMute = async () => {
        // if Mute is set to true, set to off, otherwise on
        const set = volume.Mute ? "off" : "on"

        try {
            // Make an http request
            await fetch("/api/volume/mute/" + set, {
                method: "GET",
            })
        }
        catch(err) {
            console.log("ERROR muting receiver:", err)
        }
    }

    const setVolume = async (level) => {
        // Don't set volume if it's the current value
        if (level == volume.Level) return false

        await fetch("/api/volume/level/" + level, { method: "GET" })
        .then(response => {
            if (!response.ok) {
                throw new Error("Bad network response")
            }
            return response.json()
        })
        .then(data => data)

        return level
    }

    // Set volume on load
    volumeLevel = volume.Level

    // Update volume level on change
    $: setVolume(volumeLevel)

</script>

<div class="p-1 bg-gray-100 text-black border border-gray-300 flex flex-row gap-2">
    <div class="h-full block align-middle">
        <button class="p-2 {volume.Mute ? 'text-red-500' : 'text-black hover:bg-red-500'}" on:click={handleMute} title={volume.Mute ? 'Unmute' : 'Mute'}>
            <VolumeXIcon class="mx-auto" />
        </button>
    </div>
    <div class="h-full block align-middle">
        <button class="p-2 hover:bg-blue-500 hover:text-white" on:click={()=> {volumeLevel--}}>
            <Volume1Icon class="mx-auto" />
        </button>
    </div>
    <div class="w-full">
        <input type="range" min="0" max={volume.Max} bind:value={volumeLevel} class="h-full w-full align-middle block">
    </div>
    <div class="h-full align-bottom font-semibold font-mono text-lg">
        {volumeLevel}
    </div>
    <div class="h-full block align-middle">
        <button class="p-2 hover:bg-blue-500 hover:text-white" on:click={()=> {volumeLevel++}}>
            <Volume2Icon class="mx-auto" />
        </button>
    </div>

</div>
