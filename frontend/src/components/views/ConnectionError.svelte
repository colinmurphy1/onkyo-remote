<script>
    import { AlertTriangleIcon } from 'svelte-feather-icons'

    export let error = ""
</script>

<div class="text-center text-gray-800 text-lg my-8">
    <div class="pb-2 font-semibold text-red-600">
        <AlertTriangleIcon class="inline-block h-full align-bottom" />
    </div>
    <div class="font-semibold text-red-600">
        Unable to connect to receiver
    </div>
    <div>
        {error}
    </div>
</div>
