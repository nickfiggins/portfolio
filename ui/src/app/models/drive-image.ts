export class DriveImage {
    kind: string = ""
    id: string = ""
    name: string = ""
    mimeType: string = ""
    src_url: string = ""
    constructor(kind: string = "", id: string = "", name: string = "", mimeType: string = ""){
        this.kind = kind;
        this.id = id;
        this.name = name;
        this.mimeType = mimeType;
        //this.src_url = "https://drive.google.com/uc?export=view&id=" + this.id;
        this.src_url = "https://drive.google.com/thumbnail?id=" + this.id;
    }
}
