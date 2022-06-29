<script>
	import { onMount } from 'svelte'

	// Load API scripts
	import getStatus from './api/getStatus'

	// Load components
	import PowerButton from './components/buttons/Power.svelte'
	import InputButton from './components/buttons/Input.svelte'
	import VolumeButton from './components/buttons/Volume.svelte'
	import OSDButton from './components/buttons/OSD.svelte'
	import SrcNetwork from './components/views/Network.svelte'
	import SrcTuner from './components/views/Tuner.svelte'
	import StdSource from './components/views/StdSource.svelte'

	
	let loaded = false
	let recvStatus

	const updateStatus = async () => {
		try {
			recvStatus = await getStatus()
		}
		catch (err) {
			console.log('Could not get status:', err)
		}
	}

	// On page load, get receiver status  
	onMount(async () => {
		await updateStatus()
		loaded = true
	})

	// Update status every second
	setInterval(async () => await updateStatus(), 1000);
</script>

<main>
	{#if loaded}
		<!-- OSD Control -->
		<OSDButton />

		<!-- Volume Control -->
		<VolumeButton volume={recvStatus.Volume} />

		<!-- Source Control -->
		<InputButton input={recvStatus.Input} />

		<!-- Power button -->
		<PowerButton pwrStatus={recvStatus.Power.Status} />

		{#if recvStatus.Input.HexCode == "2B"}
			<!-- NET Source -->
			<SrcNetwork status={recvStatus} />
		{:else if recvStatus.Input.HexCode == "24" || recvStatus.Input.HexCode == "25" }
			<!-- TUNER Source-->
			<SrcTuner status={recvStatus} />
		{:else}
			<!-- Other source -->
			<StdSource status={recvStatus} />
		{/if}

	{/if}
</main>
