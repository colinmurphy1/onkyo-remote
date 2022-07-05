<script>
    // Current position in HH:MM:SS
    export let current

    // Length of song in HH:MM:SS
    export let length

    // Calculate length based on HH:MM:SS
    const lengthInSeconds = (input) => {
        // Split the hours, minutes, and seconds into individual variables
        let [hour, minute, second] = input.split(':')

        // Convert hours and minutes to seconds
        hour = hour * 3600
        minute = minute * 60

        // Add up the hours, minutes, and seconds
        return +hour + +minute + +second
    }

    // Calculate rounded percentage of song completion
    let percentageCompleted
    $: {
        // If the current time is "--:--:--", make it be "00:00:00"
        if (current == '--:--:--') current = '00:00:00'

        percentageCompleted = Math.round(lengthInSeconds(current) / lengthInSeconds(length) * 100)
    }
</script>

<div class="flex flex-row">
    <div class="pl-0 pr-2 text-gray-800">
        {current}
    </div>

    <div class="w-full bg-gray-200 h-auto">
        <div class="bg-blue-500 h-full" style="width: {percentageCompleted}%"></div>
    </div>

    <div class="pr-0 pl-2 text-gray-800">
        {length}
    </div>
</div>
