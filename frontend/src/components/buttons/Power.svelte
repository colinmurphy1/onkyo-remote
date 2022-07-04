<script>
    import { PowerIcon } from 'svelte-feather-icons'
    
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

<button on:click={handlePower} class="py-1 px-2 {pwrStatus ? 'hover:bg-green-200' : 'hover:bg-red-200'}" title="Power {pwrStatus ? 'off' : 'on'} receiver">
    <PowerIcon class="inline-block h-full align-bottom {pwrStatus ? 'text-green-600' : 'text-red-600'}" />
</button>
