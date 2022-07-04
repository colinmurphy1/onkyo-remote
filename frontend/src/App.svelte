<script>
	import { onMount } from 'svelte'

	// Load API scripts
	import getStatus from './api/getStatus'

	// Load components
	import Page from './components/Page.svelte'
	import Header from './components/Header.svelte'
	import PowerOff from './components/views/PowerOff.svelte'
	import SrcNetwork from './components/views/Network.svelte'
	import SrcTuner from './components/views/Tuner.svelte'
	import StdSource from './components/views/StdSource.svelte'
	import ConnectionError from './components/views/ConnectionError.svelte'
	
	let loaded = false
	let loadError = ""
	let recvStatus

	const updateStatus = async () => {
		try {
			recvStatus = await getStatus()
			loadError = ""
			loaded = true
		}
		catch (err) {
			loaded = false
			loadError = err
		}
	}

	// On page load, get receiver status  
	onMount(async () => {
		await updateStatus()
	})

	// Update status every second
	setInterval(async () => await updateStatus(), 1000);
</script>

{#if loaded}
	<Header status={recvStatus} />
	<Page>
		{#if recvStatus.Power.Status}
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
		{:else}
			<PowerOff />
		{/if}
	</Page>
{:else}
	<Page>
		{#if loadError != ""}
			<ConnectionError error={loadError} />
		{:else}
			Loading...
		{/if}
	</Page>
{/if}



<style lang="postcss" global>
@tailwind base;
@tailwind components;
@tailwind utilities;
</style>
