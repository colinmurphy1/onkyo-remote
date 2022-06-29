<script>
    import Icon from 'svelte-awesome'
    import volumeUp from 'svelte-awesome/icons/volumeUp';
    export let volume
</script>

<button class="block py-2 px-4 hover:bg-gray-600">
    <Icon data={volumeUp} scale="1.5" class="align-sub"/>
    {volume.Level}
</button>
