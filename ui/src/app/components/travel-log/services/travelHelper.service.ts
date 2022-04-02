
import { Injectable } from '@angular/core';
import { environment } from '../../../../environments/environment';
import {
  HttpClient
} from "@angular/common/http";
import * as _ from 'lodash';
import { tap } from 'rxjs/operators';
import { map } from 'lodash';
import { DriveImage } from '../../../models/drive-image';

interface driveWrapper{
  kind: string,
  incompleteSearch: boolean,
  files: DriveImage[]
}

@Injectable({
  providedIn: 'root'
})
export class TravelHelperService {
  private api_key: string = environment.drive_api_key;
  private google_url: string = "https://www.googleapis.com/drive/v3/files?q=%27";
  private images: DriveImage[] = [];

  constructor(private http: HttpClient
  ) { 
    this.api_key = environment.drive_api_key;
  }

  getFilesFromFolder(id: string){
    return this.http.get(this.google_url + id + "%27+in+parents&key=" + this.api_key);
  }
}
