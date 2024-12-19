<script setup lang="ts">
import { Meeting } from '~/types/meeting'
import { Member } from '~/types/member'
import { Unit } from '~/types/unit'

const props = defineProps({
    unit: Unit,
    author: Member,
    meeting: Meeting,
})

</script>

<template>
  <div id="ood">
    <div class="w-3/4"/>
    <div id="memo">
      <div id="letterhead">
        <NuxtImg id="cap-seal" src="CapSeal-Color.png"/>
        <div id="address">
          <p>{{ props.unit.name.toUpperCase() }}</p>
          <p>CIVIL AIR PATROL</p>
          <p>UNITED STATES AIR FORCE AUXILIARY</p>
          <p>{{ props.unit.address }}</p>
          <p>{{ props.unit.city }}</p>
          <div class="h-4"/>
        </div>
        <NuxtImg id="unit-patch" src="BcsPatch-Color.png"/>
      </div>
      <div id="header">
        <div id="date">
          <p>{{ props.meeting.formatDateLong() }}</p>
        </div>
        <div class="h-4"/>
        <p>MEMORANDUM FOR {{ props.unit.name.toUpperCase() }}</p>
        <div class="h-4"/>
        <p>FROM: {{ props.author.officeSymbol }}</p>
        <div class="h-4"/>
        <p>SUBJECT: {{ props.meeting.formatDateLong().toUpperCase() }} ORDERS OF THE DAY</p>
      </div>
      <div id="body">

      </div>
    </div>
    <div class="w-3/4"/>
  </div>
</template>

<style scoped>
#ood {
  @apply flex flex-row
}

#memo {
  @apply w-full;
}

#letterhead {
  @apply w-full flex flex-row items-center content-between;
}

#cap-seal {
  @apply h-28
}

#address {
  font-weight: bold;
  @apply flex flex-col items-center text-symbolBlue dark:text-white
}

#unit-patch {
  @apply h-28
}

#date {
  @apply flex flex-row-reverse
}

</style>