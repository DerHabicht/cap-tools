import { DateTime } from 'luxon';
import { Uniform } from '~/types/uniforms';

export const enum MeetingType {
    RegularMeeting = 'regular meeting',
    SpecialActivity = 'special activity',
}

export class Meeting {
    location: string;
    address: string;
    start: DateTime;
    end: DateTime;
    meetingType: MeetingType;
    topic: string;
    blocks: MeetingBlock[];
    uod: Uniform[];
    additionalOrders: string;

    constructor(
        location: string,
        address: string,
        start: DateTime,
        end: DateTime,
        meetingType: MeetingType,
        topic: string,
        blocks: MeetingBlock[],
        uod: Uniform[],
        additionalOrders: string,
    ) {
        this.location = location;
        this.address = address;
        this.start = start;
        this.end = end;
        this.meetingType = meetingType;
        this.topic = topic;
        this.blocks = blocks;
        this.uod = uod;
        this.additionalOrders = additionalOrders;
    }
}

export class MeetingBlock {
    start: DateTime;
    end: DateTime;
    topic: string;
    location: string;

    constructor(start: DateTime, end: DateTime, topic: string, location: string) {
        this.start = start
        this.end = end
        this.topic = topic
        this.location = location
    }
}