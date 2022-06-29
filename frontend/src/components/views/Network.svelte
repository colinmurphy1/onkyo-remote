<script>
    // Get status from receiver
    export let status;

    // Import components
    import ProgressBar from '../ProgressBar.svelte'

    // Shorthand variables
    let STitle, SArtist, SAlbum, artURL, TCurrent, TTotal
    $: STitle = status.SongInfo.Title
    $: SArtist = status.SongInfo.Artist
    $: SAlbum = status.SongInfo.Album
    $: TCurrent = status.SongInfo.Track.Current
    $: TTotal = status.SongInfo.Track.Total

    const updateArt = () => {
        const random = Math.floor(Math.random() * 10000)
        artURL = "/api/art?" + random 
    }

    $: STitle, SArtist, SAlbum && setTimeout(() => updateArt(), 1000)
</script>


<!-- NET Source and TRACK Position -->
<div class="flex flex-row p-2">
    <div class="pl-0 px-2 text-lg">
        <strong class="font-semibold">{status.Input.NetSource}</strong>
    </div>
    <div class="px-2 align-middle h-auto">
        <div class="text-gray-800">{TCurrent} / {TTotal}</div>
    </div>
</div>
<!-- NET Source and TRACK Position End -->


<div class="grid md:grid-cols-2 sm:grid-cols-1">
    <div class="col-span-1 p-2">
        <!-- Album art -->
        {#if status.SongInfo.AlbumArt}
        <img src={artURL} alt="ALBUM ARTWORK" class="mx-auto">
        {:else}
        No Album art available
        {/if}
        <!-- Album art end -->
    </div>

    <div class="col-span-1 p-2 lg:align-middle">
        <!-- Song Information -->
        <div class="mb-4 text-center md:text-left">
            <div class="md:text-3xl text-2xl font-semibold pb-1">{STitle}</div>
            <div class="text-xl text-gray-800 pb-1">{SAlbum}</div>
            <div class="text-xl text-gray-800">{SArtist}</div>
        </div>
        <!-- Song Information End -->

        <!-- Progress Bar start -->
        <ProgressBar
            current={status.SongInfo.Time.Current}
            length={status.SongInfo.Time.Length}
        />
        <!-- Progress bar end -->

        <!-- Playback controls start -->

        <!-- Playback controls end -->
    </div>
</div>
