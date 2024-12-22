<script setup lang="ts">
import { Meeting, MeetingBlock } from '~/types/meeting'
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
          <ol>
            <li>This meeting is operations night, with class sections for USAFA Admissions and Christmas fun!</li>
            <li>The meeting will be conducted between 1800 MST and 2030 MST (6:00 P.M. to 8:30 P.M.). The schedule
              is as follows:
              <table>
                <thead>
                  <tr>
                    <th>Time</th>
                    <th>Activity</th>
                    <th>Location</th>
                  </tr>
                </thead>
                <tbody>
                  <tr v-for="block in meeting.blocks" :key="block.start">
                    <td>{{ block.formatStartTime() }}</td>
                    <td>{{ block.topic }}</td>
                    <td>{{ block.location }}</td>
                  </tr>
                </tbody>
              </table>
            </li>
            <li>
              Uniform of the Day is ABU, CWU, CFU, or squadron t-shirt and slacks.
            </li>
            <li>
              All cadets must report to their chain of command with their attendance for this meeting, testing
              requests, and/or promotion request. See attached BCSW 503 Promotion Request Worksheet for promotions.
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