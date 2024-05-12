<template>
    <div>
        <h1>table page</h1>
        <p>{{ users }}</p>
    </div>
</template>
  
<script setup lang="ts">

const {data: users} = await useFetch('/api/users')

</script>