<template>
  <footer>
    <p>CAP/A5&mdash;Units v0.1.0</p>
    <div class="spacer"/>
    <p>&copy;{{ new Date().getFullYear() }} Robert Herschel Hawk</p>
  </footer>

</template>

<style scoped>
footer {
  height: 2em;
  @apply flex items-center justify-center text-white bg-symbolBlue
}

.spacer {
  width: 4em;
}
</style>