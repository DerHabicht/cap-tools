<script setup lang="ts">

</script>

<template>
  <div id="content">
    <slot/>
  </div>
</template>

<style scoped>
#content {
  min-height: calc(100vh - 8em);
  @apply bg-white text-black dark:bg-gray dark:text-white mx-32
}
</style>
