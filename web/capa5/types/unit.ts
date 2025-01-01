export enum UnitKind {
    Group = "group",
    Squadron = "squadron",
    Flight = "flight",
}

export enum UnitCategory {
    Admin = "admin",
    Composite = "composite",
    Cadet = "cadet",
    Senior  = "senior",
}

export class Unit {
    charterNumber: string;
    kind: UnitKind;
    category: UnitCategory;
    name: string;
    address: string;
    city: string;

    constructor(
        charterNumber: string,
        kind: string,
        category: string,
        name: string,
        address: string,
        city: string,
    ) {
        this.charterNumber = charterNumber;
        this.kind = kind
        this.category = category
        this.name = name
        this.address = address
        this.city = city
    }
}