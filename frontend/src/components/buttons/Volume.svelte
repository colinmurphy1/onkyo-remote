<script>
    import Icon from 'svelte-awesome'
    import volumeUp from 'svelte-awesome/icons/volumeUp';
    export let volume
</script>

<button class="inline-block py-1 px-2 hover:bg-slate-400">
    <Icon data={volumeUp} class="align-sub"/>
    {volume.Level}
</button>
