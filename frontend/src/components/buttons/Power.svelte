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
            await fetch("http://localhost:8080/api/power/" + set, {
                method: "GET",
            })
        }
        catch(err) {
            console.log("ERROR powering off receiver:", err)
        }
    }

</script>



<button on:click={handlePower} class="block py-2 px-4 hover:text-black {pwrStatus ? 'hover:bg-green-200' : 'hover:bg-red-200'}">
    <Icon data={powerOff} scale="1.5" class="align-middle {pwrStatus ? 'text-green-600' : 'text-red-600'}"/>
</button>
