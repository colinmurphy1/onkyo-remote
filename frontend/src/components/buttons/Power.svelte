<script>
    import Icon from 'svelte-awesome'
    import powerOff from 'svelte-awesome/icons/powerOff';
    
    export let pwrStatus;
    
    const handlePower = async () => {
        let set = "on"
        // If receiver is on, power it off
        if (pwrStatus) set = "off"

        try {
            // Make an http request
            await fetch("/api/power/" + set, {
                method: "GET",
            })
        }
        catch(err) {
            console.log("ERROR powering off receiver:", err)
        }
    }

</script>

<button on:click={handlePower} class="inline-block py-1 px-2 {pwrStatus ? 'hover:bg-green-200' : 'hover:bg-red-200'}" title="Click to power {pwrStatus ? 'off' : 'on'} receiver">
    <Icon data={powerOff} scale="1" class="align-middle {pwrStatus ? 'text-green-600' : 'text-red-600'}"/>
    Power
</button>
