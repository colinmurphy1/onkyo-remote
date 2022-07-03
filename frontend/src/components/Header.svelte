<script>
	import PowerButton from './buttons/Power.svelte'
    import InputButton from './buttons/Input.svelte'
    import VolumeButton from './buttons/Volume.svelte'
    //import MuteButton from './buttons/Mute.svelte'

    export let status
</script>

<header class="mb-4 px-0.5 bg-gray-200 border-b border-gray-100 border-b-gray-300">
    <div class="max-w-4xl mx-auto relative">
        <div class="absolute left-0 flex flex-row pl-1">
            <InputButton />
            <VolumeButton volume={status.Volume} />
            <!--<MuteButton muteStatus={status.Volume.Mute} />-->
        </div>
    
        <!-- Power button -->
        <div class="absolute right-0 pr-1">
            <PowerButton pwrStatus={status.Power.Status} />
        </div>
        <!-- Power button end -->
        <div class="text-center font-semibold py-1">
            {#if status.Power.Status == false}
                &nbsp;
            {:else}
    
                <!-- If the source is NET, display the NET source -->
                {#if status.Input.HexCode == "2B"}
                    {status.Input.NetSource}
                {:else}
                    {status.Input.Name}
                {/if}
            {/if}
        </div>    
    </div>
</header>
