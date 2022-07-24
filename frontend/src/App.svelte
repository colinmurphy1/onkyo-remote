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
	let recvSource
	let pageTitle = ""

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


	// Sets the page title to the receiver's friendly name if specified,
	// otherwise default to model number, and if that is not specified set a
	// default title
	const setPageTitle = (info) => {
		if (info.FriendlyName != "") {
			return info.FriendlyName
		} else if (info.ModelName != "") {
			return info.ModelName
		}
		return "Onkyo Remote"
	}

    // API call to get all available inputs
    const getSource = async () => {
        const req = await fetch("/api/source", { method: "GET" })
        .then(response => {
            if (!response.ok) {
                throw new Error("Bad network response")
            }
            return response.json()
        })
        .then(data => data)

        return req.data
    }

	// On page load, get receiver status and sources
	onMount(async () => {
		await updateStatus()
		recvSource = await getSource()
		pageTitle = setPageTitle(recvStatus.Info)
	})

	// Update status every second
	setInterval(async () => await updateStatus(), 1000);
</script>

<svelte:head>
	<title>{pageTitle}</title>
</svelte:head>


{#if loaded}
	<Header status={recvStatus} sources={recvSource} />
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
