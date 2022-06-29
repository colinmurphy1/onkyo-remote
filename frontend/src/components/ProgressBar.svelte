<script>
    // Current position in HH:MM:SS
    export let current

    // Length of song in HH:MM:SS
    export let length

    // Calculate length based on HH:MM:SS
    const lengthInSeconds = (input) => {
        // If there's no length specified, return 0 seconds
        if (input == '--:--:--') {
            return 0
        }

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
        percentageCompleted = Math.round(lengthInSeconds(current) / lengthInSeconds(length) * 100)
    }
</script>

<div class="flex flex-row">
    <div class="pl-0 pr-2 text-lg">
        {current}
    </div>

    <div class="w-full bg-gray-200 rounded-full h-auto dark:bg-gray-700">
        <div class="bg-blue-600 h-full rounded-full" style="width: {percentageCompleted}%"></div>
    </div>

    <div class="pr-0 pl-2 text-lg">
        {length}
    </div>
</div>
