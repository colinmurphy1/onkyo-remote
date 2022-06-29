<script>
    import Icon from 'svelte-awesome'
    import music from 'svelte-awesome/icons/music';

    export let input
</script>

<button class="block py-2 px-4 hover:bg-gray-600">
    <Icon data={music} scale="1.5" class="align-sub"/>
    {input.Name}
</button>
