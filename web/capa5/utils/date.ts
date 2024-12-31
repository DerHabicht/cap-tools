import { DateTime } from 'luxon';

function formatDateFull(date: DateTime): string {
    return date.toFormat('cccc, d LLLL yyyy');
}

function formatDateLong(date: DateTime): string {
    return date.toFormat('d LLLL yyyy');
}

export function formatDate(date: DateTime, format: 'full' | 'long'): string {
    switch (format) {
    case 'full':
        return formatDateFull(date);
    case 'long':
        return formatDateLong(date);
    }
}

function formatTimeFull(date: DateTime): string {
    return date.toFormat('HHmm ZZZZ');
}

function formatTimeShort(date: DateTime): string {
    return date.toFormat('HHmm');
}

function formatHeathenTime(date: DateTime): string {
    return date.toFormat('h:mm a')
}

export function formatTime(date: DateTime, format: 'full' | 'short' | 'heathen'): string {
    switch (format) {
    case 'full':
        return formatTimeFull(date);
    case 'short':
        return formatTimeShort(date);
    case 'heathen':
        return formatHeathenTime(date);
    }
}
