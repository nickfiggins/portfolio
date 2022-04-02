export class ProjectImage {
    file_name: string;
    path: string;
    type?: string;

    constructor(file_name: string, path: string) {
      this.file_name = file_name;
      this.path = path;
    }
  }