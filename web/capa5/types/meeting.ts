import { Uniform } from '~/uniforms';

export class Meeting {
    start: Date;
    end: Date;
    blocks: MeetingBlock[]
    constructor(start: Date, end: Date, blocks: MeetingBlock[]) {
        this.start = start;
        this.end = end;
        this.blocks = blocks;
    }
    formatDateLong(): string {
        const options = [{day: 'numeric'}, {month: 'long'}, {year: 'numeric'}];

        return options.map((option): string => {
            const formatter = new Intl.DateTimeFormat('en', option);
            return formatter.format(this.start)
        }).join(" ");
    }
}

export class MeetingBlock {
    start: Date;
    end: Date;
    topic: string;
    location: string;
    uod: Uniform[];
    constructor(start: Date, end: Date, topic: string, location: string, uod: Uniform[]) {
        this.start = start
        this.end = end
        this.topic = topic
        this.location = location
        this.uod = uod
    }
}