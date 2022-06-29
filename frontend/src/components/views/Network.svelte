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

<div>

    <div>{status.Input.NetSource}</div>
    <div>{TCurrent}/{TTotal}</div>

    <!-- Album art -->
    {#if status.SongInfo.AlbumArt}
    <div>
        <img src={artURL} alt="ALBUM ARTWORK">
    </div>
    {/if}
    <!-- Album art end -->

    <ProgressBar
        current={status.SongInfo.Time.Current}
        length={status.SongInfo.Time.Length}
    />

    <div>{STitle}</div>
    <div>{SAlbum}</div>
    <div>{SArtist}</div>
</div>