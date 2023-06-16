export class Feed {
    url:string;
    addedOn:Date;

    constructor(url:string, addedOn?:Date) {
        this.url = url;
        this.addedOn = addedOn ?? new Date();
    }
}