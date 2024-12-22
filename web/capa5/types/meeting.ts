import { Uniform } from '~/uniforms';

export class Meeting {
    start: Date;
    end: Date;
    blocks: MeetingBlock[];
    uod: Uniform[];
    constructor(start: Date, end: Date, blocks: MeetingBlock[], uod: Uniform[]) {
        this.start = start;
        this.end = end;
        this.blocks = blocks;
        this.uod = uod;
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
    constructor(start: Date, end: Date, topic: string, location: string) {
        this.start = start
        this.end = end
        this.topic = topic
        this.location = location
    }
    formatStartTime(): string {
        const options = [{hour: '2-digit', hour12: false}, {minute: '2-digit'}];

        return options.map((option): string => {
            const formatter = new Intl.DateTimeFormat('en', option);
            return formatter.format(this.start)
        }).join("");
    }
}