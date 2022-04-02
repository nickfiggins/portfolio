import { Component, OnInit } from '@angular/core';
import { environment } from '../../../environments/environment';
import {TravelHelperService} from './services/travelHelper.service';
import { DriveImage } from '../../models/drive-image';
import { ImageModalComponent } from '../shared/components/image-modal/image-modal.component';
import { MatDialog, MatDialogConfig } from '@angular/material/dialog';
import { EnumService } from '../shared/services/Enum.service';
import {PlaceService} from '../../admin/services/place.service';
import { Place } from 'src/app/models/place';

interface driveWrapper{
  kind: string,
  incompleteSearch: boolean,
  files: DriveImage[]
}

@Component({
  selector: 'app-travel-log',
  templateUrl: './travel-log.component.html',
  styleUrls: ['./travel-log.component.scss']
})
export class TravelLogComponent implements OnInit {
  place: any = {};
  drive_api_key: string = environment.drive_api_key;
  images: DriveImage[] = [];
  imageWrapper:any;
  showButton: boolean[] = [];
  driveFolders: any;
  places: any[] = [];
  particlesOptions = {};
  breakpoint: any;

  constructor(private helperService : TravelHelperService, private modal: MatDialog, public enums: EnumService,
    public placeService: PlaceService) { 
      this.particlesOptions = enums.particlesOptions;
    }

    
  setPlaces(places: any[]){
    this.places = places;
  }

  onResize(event: any) {
    this.breakpoint = (event.target.innerWidth <= 400) ? 1 : 8;
  }
  
  ngOnInit(): void {
    this.breakpoint = (window.innerWidth <= 400) ? 1 : 8;
    this.driveFolders = this.enums.driveFolders;
    this.placeService.getAllPlaces().subscribe(
      (res: any) => {
        this.setPlaces(res.data);
      }
    )
  }

  changePlace(newPlace:string): void{
      this.place = this.driveFolders[newPlace];
      this.helperService.getFilesFromFolder(this.place.id).subscribe(
        (res: any) => {this.setImages(res.files);          // On emit
        });
  }

  openModal(curIndex: number, curId:string) {
    const imageModalConfig = new MatDialogConfig();
    imageModalConfig.scrollStrategy?.disable();
    imageModalConfig.autoFocus = false;
    
    imageModalConfig.data = {
      index: curIndex,
      id: curId
    }
    let dialogRef = this.modal.open(ImageModalComponent, imageModalConfig);
  }

  mouseEnter(idx: number, id: string){
    this.showButton[idx] = true;
  }

  mouseLeave(idx: number, id: string){
    this.showButton[idx] = false;
  }

  setImages(imageList: DriveImage[]){
    this.images = imageList.map(img => new DriveImage(img.kind, img.id, img.name, img.mimeType));
    this.showButton = new Array(this.images.length).fill(false);
  }

}
