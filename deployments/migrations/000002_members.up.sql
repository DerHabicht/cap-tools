CREATE TYPE member_category AS ENUM (
    'SENIOR',
    'CADET',
    'CADET SPONSOR',
    'AEM',
    'STATE LEG',
    'LEGISLATIVE',
    'PATRON'
);

CREATE TYPE grade AS ENUM (
    'Maj Gen',
    'Brig Gen',
    'Col',
    'Lt Col',
    'Maj',
    'Capt',
    '1st Lt',
    '2d Lt',
    'SFO',
    'TFO',
    'FO',
    'CMSgt',
    'SMSgt',
    'MSgt',
    'TSgt',
    'SSgt',
    'SM',
    'C/Col',
    'C/Lt Col',
    'C/Maj',
    'C/Capt',
    'C/1st Lt',
    'C/2d Lt',
    'C/CMSgt',
    'C/SMSgt',
    'C/MSgt',
    'C/TSgt',
    'C/SSgt',
    'C/SrA',
    'C/A1C',
    'C/Amn',
    'C/AB'
);

CREATE TABLE members (
    capid INT PRIMARY KEY,
    member_category member_category NOT NULL,
    last_name TEXT NOT NULL,
    first_name TEXT NOT NULL,
    grade grade NOT NULL,
    unit_charter VARCHAR(11),
    FOREIGN KEY (unit_charter) REFERENCES units (charter_number) ON DELETE CASCADE
);