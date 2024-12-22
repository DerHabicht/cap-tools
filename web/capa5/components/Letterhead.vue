<script setup lang="ts">

</script>

<template>
  <div id="address">
    <p>{{ props.unit.name.toUpperCase() }}</p>
    <p>CIVIL AIR PATROL</p>
    <p>UNITED STATES AIR FORCE AUXILIARY</p>
    <p>{{ props.unit.address }}</p>
    <p>{{ props.unit.city }}</p>
    <div class="h-4"/>
  </div>
</template>

<style scoped>

</style>