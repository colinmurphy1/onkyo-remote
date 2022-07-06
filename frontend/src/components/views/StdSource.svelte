<script>
    export let status
</script>

<div>

    <h2 class="text-center text-2xl my-4 font-semibold">{status.Input.Name}</h2>

    <!--
        TODO: Add source information to be displayed here.

        Will need to adjust API endpoint to enable this feature.

        Display:
        * Source Analog/Digital
        * If digital, display sample rate, pcm/dsd, etc.
        * STEREO/DIRECT/5.1 Etc.
    -->

    <div class="grid grid-cols-2 md:grid-cols-5 gap-1 text-gray-600 m-2 text-center">
        <div>
            <strong>Port:</strong><br>
            {status.Input.Info.InputPort ? status.Input.Info.InputPort : "N/A"}
        </div>
        <div>
            <strong>Format:</strong><br>
            {status.Input.Info.InputFormat ? status.Input.Info.InputFormat : "N/A"}
        </div>
        <div>
            <strong>Sample rate:</strong><br>
            {status.Input.Info.SamplingFreq ? status.Input.Info.SamplingFreq : "N/A"}
        </div>
        <div>
            <strong>Channels:</strong><br>
            {status.Input.Info.InputSignalChannel ? status.Input.Info.InputSignalChannel : "N/A"}
        </div>
        <div>
            <strong>Mode:</strong><br>
            {status.Input.Info.ListenMode ? status.Input.Info.ListenMode : "N/A"}
        </div>
    </div>
</div>