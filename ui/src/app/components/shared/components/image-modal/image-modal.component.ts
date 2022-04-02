import { Component, OnInit, Inject } from '@angular/core';
import {MatDialog} from '@angular/material/dialog';
import {MAT_DIALOG_DATA, MatDialogRef} from "@angular/material/dialog";

@Component({
  selector: 'app-image-modal',
  templateUrl: './image-modal.component.html',
  styleUrls: ['./image-modal.component.scss']
})
export class ImageModalComponent implements OnInit {
  id: string;
  index: number;
  src_url: string;

  constructor(private matModal: MatDialog, private dialogRef: MatDialogRef<ImageModalComponent>,
    @Inject(MAT_DIALOG_DATA) data: any) { 
      this.id = data.id;
      this.index = data.index;
      this.src_url = "https://drive.google.com/uc?export=view&id=" + this.id;
    }

  openModal() {
    const modalRef = this.matModal.open(matModalContent);

    modalRef.afterClosed().subscribe(result => {
      console.log(`Modal result: ${result}`);
    });
  }

  ngOnInit(): void {
  }

}

export class matModalContent{}
