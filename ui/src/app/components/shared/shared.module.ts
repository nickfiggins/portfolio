import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { CarouselComponent } from './carousel/carousel.component';
import {IvyCarouselModule} from 'angular-responsive-carousel';
import { MatButtonModule } from '@angular/material/button';

@NgModule({
  declarations: [CarouselComponent],
  imports: [
    CommonModule,
    IvyCarouselModule,
    MatButtonModule
  ],
  exports: [
    CarouselComponent
  ]
})
export class SharedModule { }
