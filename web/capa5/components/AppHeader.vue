<script setup lang="ts">
const isDark = useDark()

const items = ref([
  {
    label: 'Orders of the Day',
    route: '/units/ut-080/ood'
  },
  {
    label: 'Calendar',
    route: '/units/ut-080/calendar'
  },
  {
    label: 'Unit Documents',
    route: '/units/ut-080/docs'
  },
])

</script>

<template>
  <header>
    <Menubar :model="items">
      <template #start>
        <NuxtLink to="/units">
          <NuxtImg v-if=false id="site-icon" src="BcsPatch-Color.png"/>
          <Button v-else icon="pi pi-home"/>
        </NuxtLink>
      </template>
      <template #item="{ item, props, hasSubmenu, root } ">
        <NuxtLink :to="item.route" class="menu-item" activeClass="active-menu-item">
          <span>{{ item.label }}</span>
        </NuxtLink>
      </template>
      <template #end>
        <Button v-if="isDark" icon="pi pi-moon" @click="$emit('themeChange')"/>
        <Button v-else icon="pi pi-sun" @click="$emit('themeChange')"/>
        <Button v-if=false>
          <Gravatar class="gravatar" email="robert.hawk@gmail.com" />
        </Button>
        <Button v-else icon="pi pi-user"/>
      </template>
    </Menubar>
  </header>
</template>

<style scoped>

#site-icon {
  height: 5em;
  @apply p-1
}

.active-menu-item {
  @apply px-4 py-8 text-2xl bg-scarletRed h-full
}

.gravatar {
  height: 2em;
}

.menu-item {
  //height: 6em;
  @apply px-4 py-8 text-2xl hover:bg-scarletRed h-full
}

.p-button {
  @apply text-2xl px-4 hover:bg-scarletRed
}

.p-menubar{
  height: 6em;
  @apply flex items-center px-3 py-2 gap-2
  bg-symbolBlue
  text-white
}

.p-menubar-end{
  @apply flex align-middle items-center px-3 py-2 gap-2
}

</style>