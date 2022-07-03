<script>    
    export let muteStatus;
    
    const handleMute = async () => {
        let set = "on"
        // If receiver is on, power it off
        if (muteStatus) set = "off"

        try {
            // Make an http request
            await fetch("/api/volume/mute/" + set, {
                method: "GET",
            })
        }
        catch(err) {
            console.log("ERROR powering off receiver:", err)
        }
    }

</script>

<button on:click={handleMute} class="inline-block py-1 px-2 {muteStatus ? 'bg-red-200 text-red-700' : 'hover:bg-red-200'}">
    Mute
</button>
