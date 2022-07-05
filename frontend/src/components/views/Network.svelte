<script>
    // Get status from receiver
    export let status;

    // Import components
    import ProgressBar from '../ProgressBar.svelte'
    import NetControl from '../buttons/NetControl.svelte';

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

<div class="grid md:grid-cols-2 sm:grid-cols-1">
    <div class="col-span-1 p-2">
        <!-- Album art -->
        {#if status.SongInfo.AlbumArt}
            <img src={artURL} alt="" class="mx-auto">
        {:else}
            <div class="w-40 h-40 bg-gray-100 mx-auto">
            </div>
        {/if}
        <!-- Album art end -->
    </div>

    <div class="col-span-1 p-2 lg:align-middle">
        <!-- Song Information -->
        <div class="mb-4 text-center md:text-left">
            <div class="md:text-3xl text-2xl font-semibold pb-1">{STitle ? STitle : "No Title"}</div>
            <div class="text-xl text-gray-800 pb-1">{SAlbum ? SAlbum : "No Album"}</div>
            <div class="text-xl text-gray-800">{SArtist ? SArtist : "No Artist"}</div>
        </div>
        <!-- Song Information End -->

        <!-- Progress Bar start -->
        <ProgressBar
            current={status.SongInfo.Time.Current}
            length={status.SongInfo.Time.Length}
        />
        <!-- Progress bar end -->

        <!-- Playback controls start -->
        <div class="text-center my-4">
            <NetControl action="trdn" />
            <!-- Hide paused if paused, play if playing etc. -->
            {#if status.SongInfo.Status == "Play"}
                <NetControl action="pause" />
            {:else}
                <NetControl action="play" />
            {/if}
            <NetControl action="trup" />
        </div>
        <!-- Playback controls end -->

        <!-- Track position -->
        <div class="text-gray-800 text-center pb-2">Track {TCurrent} of {TTotal}</div>
        <!-- Track position end -->
    </div>
</div>
