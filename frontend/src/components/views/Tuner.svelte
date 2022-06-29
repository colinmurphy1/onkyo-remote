<script>
    // Get status from receiver
    export let status;

    let frequency, band

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

</script>

<div>
    <div>
        FREQ: {frequency} {band}
    </div>
    <div>
        PRS: {status.Tuner.Preset}
    </div>
</div>