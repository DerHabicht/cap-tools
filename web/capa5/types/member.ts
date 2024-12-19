export class Member {
    lastName: string;
    firstName: string;
    grade: string;
    officeSymbol: string;

    constructor(lastName: string, firstName: string, grade: string, officeSymbol: string) {
        this.lastName = lastName;
        this.firstName = firstName;
        this.grade = grade;
        this.officeSymbol = officeSymbol;
    }
}