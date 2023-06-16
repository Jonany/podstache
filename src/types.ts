export class Feed {
    url:string;
    addedOn:Date = new Date();

    constructor(url:string) {
        this.url = url;
    }
}