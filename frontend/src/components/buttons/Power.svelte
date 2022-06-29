<script>
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



<button on:click={handlePower}>
    Power<br>
    {#if pwrStatus}
    <span class="pwrOn">
        ON
    </span>
    {:else}
    <span class="pwrOff">
        OFF
    </span>
    {/if}
</button>


<style>
    .pwrOn {
        color: green;
    }
    .pwrOff {
        color: red;
    }
</style>