<script setup lang="ts">
import { Meeting } from '~/types/meeting'
import { formatUniformList } from '~/types/uniforms'
import { Member } from '~/types/member'
import { Unit } from '~/types/unit'
import { formatDate, formatTime } from '~/utils/date'

const props = defineProps({
    unit: Unit,
    author: Member,
    meeting: Meeting,
})

</script>

<template>
  <div id="ood">
    <div id="document">
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
      <div id="memo">
        <div id="header">
          <div id="date">
            <p>{{ formatDate(props.meeting.start, 'long') }}</p>
          </div>
          <div class="h-4"/>
          <p>MEMORANDUM FOR {{ props.unit.name.toUpperCase() }}</p>
          <div class="h-4"/>
          <p>FROM: {{ props.author.officeSymbol }}</p>
          <div class="h-4"/>
          <p>SUBJECT: {{ formatDate(props.meeting.start, 'long').toUpperCase() }} ORDERS OF THE DAY</p>
        </div>
        <div id="body">
          <ol>
            <li>
              On {{ formatDate(props.meeting.start, 'full') }}, the {{ props.unit.name }} will be holding a
              {{ props.meeting.meetingType }} at {{ props.meeting.location }} ({{ props.meeting.address }}). The topic
              is {{ props.meeting.topic }}.
            </li>
            <li>The meeting will be conducted between {{ formatTime(props.meeting.start, 'full') }} and
              {{ formatTime(props.meeting.end, 'full') }} ({{ formatTime(props.meeting.start, 'heathen') }} to
              {{ formatTime(props.meeting.end, 'heathen') }}). The schedule is as follows:
              <table>
                <thead>
                  <tr>
                    <th>Start</th>
                    <th>End</th>
                    <th>Activity</th>
                    <th>Location</th>
                  </tr>
                </thead>
                <tbody>
                  <tr v-for="block in meeting.blocks" :key="block.start">
                    <td>{{ formatTime(block.start, 'short') }}</td>
                    <td>{{ formatTime(block.end, 'short') }}</td>
                    <td>{{ block.topic }}</td>
                    <td>{{ block.location }}</td>
                  </tr>
                </tbody>
              </table>
            </li>
            <li v-if="props.meeting.uod.length > 1">
              The uniforms of the day are the {{ formatUniformList(props.meeting.uod, 'or') }}.
            </li>
            <li v-else>
              The uniform of the day is the {{ props.meeting.uod[0] }}.
            </li>
            <li v-if="props.meeting.additionalOrders">
              {{ props.meeting.additionalOrders }}
            </li>
          </ol>
        </div>
        <div id="signature">
          <div class="w-1/2">&nbsp;</div>
          <div>
            <p class="mb-4">// SIGNED //</p>

            <p>RYAN FITZGERALD, Capt, CAP</p>
            <p>Commander</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
ol {
 counter-reset: para;
}

li {
  @apply my-4
}

li:before {
  counter-increment: para;
  content: "" counter(para) ".  ";
}

table {
  @apply my-4 w-4/5 mx-auto
}

#ood {
  @apply flex flex-row w-full
}

#document {
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
  @apply m-auto flex flex-col items-center text-symbolBlue dark:text-white
}

#memo {
  @apply px-12
}

#header {
  @apply pt-4
}

#body {
}

#signature {
  @apply flex flex-row mt-12
}

#unit-patch {
  @apply h-28
}

#date {
  @apply flex flex-row-reverse
}

</style>