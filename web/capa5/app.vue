<script setup lang="ts">

const isDark = useDark();
const toggleDark = useToggle(isDark);

</script>

<template>
  <div id="root">
    <NuxtRouteAnnouncer />
    <AppHeader @themeChange="toggleDark()" />
    <NuxtPage />
    <AppFooter />
  </div>
</template>

<style>
@import 'primeicons/primeicons.css';

#root {
  @apply bg-white text-black dark:bg-gray dark:text-white
}

h1 {
  @apply text-3xl my-2
}
</style>