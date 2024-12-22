<script setup lang="ts">

const isDark = useDark();
const toggleDark = useToggle(isDark);

</script>

<template>
  <div id="root">
    <NuxtRouteAnnouncer />
    <AppHeader @themeChange="toggleDark()" />
    <NuxtPage id="content"/>
    <AppFooter />
  </div>
</template>

<style>
@import 'primeicons/primeicons.css';

#root {
  min-height: 100vh;
  @apply bg-white text-black dark:bg-gray dark:text-white
}

#content {
  min-height: calc(100vh - 8rem);
  @apply bg-white text-black dark:bg-gray dark:text-white w-1/2 mx-auto my-16
}
</style>