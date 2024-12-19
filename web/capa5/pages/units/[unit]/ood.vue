<script setup lang="ts">
import { Meeting, MeetingBlock } from '~/types/meeting'
import { Member } from '~/types/member'
import { Uniform } from '~/types/uniforms'
import { Unit } from '~/types/unit'

const meetingBlocks = [
      new MeetingBlock(
          new Date('2024-12-19T18:00-07:00'),
          new Date('2024-12-19T18:10-07:00'),
          'Opening Formation',
          'Gym',
      ),
      new MeetingBlock(
          new Date('2024-12-19T18:10-07:00'),
          new Date('2024-12-19T18:55-07:00'),
          'USAFA Admissions',
          'Auditorium',
      ),
      new MeetingBlock(
        new Date('2024-12-19T18:55-07:00'),
        new Date('2024-12-19T19:25-07:00'),
        'Flight Time',
        'Gym',
      ),
      new MeetingBlock(
        new Date('2024-12-19T18:55-07:00'),
        new Date('2024-12-19T19:25-07:00'),
        'White Elephant Gift Exchange',
        'Gym',
      ),
      new MeetingBlock(
          new Date('2024-12-19T19:30-07:00'),
          new Date('2024-12-19T20:15-07:00'),
          'Break',
          'Gym',
      ),
      new MeetingBlock(
          new Date('2024-12-19T20:15-07:00'),
          new Date('2024-12-19T20:30-07:00'),
          'Closing Formation',
          'Gym',
      ),
    ]

const meeting = new Meeting(
    new Date('2024-12-19T18:00-07:00'),
    new Date('2024-12-19T20:30-07:00'),
    meetingBlocks,
    [Uniform.ABU, Uniform.CWU, Uniform.CFU]
);

const unit = new Unit(
    'Blackhawk Cadet Squadron',
    'RMR-UT-080',
    '12953 S Minuteman Dr',
    'Draper, UT 84020'
);

const author = new Member(
    'Fitzgerald',
    'Ryan',
    'Capt',
    'UT080/CC'
);


</script>

<template>
  <OrdersOfTheDay :unit="unit" :author="author" :meeting="meeting"/>
</template>

<style scoped>

</style>