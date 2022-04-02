import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HomeComponent } from './home/home.component';
import { MatCardModule } from '@angular/material/card';
import { MatIconModule } from '@angular/material/icon';
import { FlexLayoutModule } from '@angular/flex-layout';
import {MatButtonModule} from '@angular/material/button';
import { TravelLogComponent } from './travel-log/travel-log.component';
import {RouterModule} from '@angular/router';
import { NgParticlesModule } from 'ng-particles';
import { SidenavComponent } from './shared/components/sidenav/sidenav.component';
import {MatMenuModule} from '@angular/material/menu';
import {TravelHelperService} from './travel-log/services/travelHelper.service';
import {MatGridListModule} from '@angular/material/grid-list';
import { ImageModalComponent } from './shared/components/image-modal/image-modal.component';
import {MatDialogModule, MatDialogRef, MAT_DIALOG_DATA} from '@angular/material/dialog';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { StuffComponent } from './stuff/stuff.component';
import { TutoringComponent } from './tutoring/tutoring.component';
import { ContactComponent } from './contact/contact.component';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule} from '@angular/material/input';
import {MatRadioModule} from '@angular/material/radio';
import {TechModule} from './tech/tech.module';
import { CarouselComponent } from './shared/carousel/carousel.component';
import { SharedModule } from './shared/shared.module';
import { FreelanceComponent } from './freelance/freelance.component';

@NgModule({
  declarations: [HomeComponent, TravelLogComponent, SidenavComponent, ImageModalComponent, StuffComponent, TutoringComponent, ContactComponent, FreelanceComponent],
  imports: [
    CommonModule,
    MatCardModule,
    MatIconModule,
    FlexLayoutModule,
    MatButtonModule,
    RouterModule,
    NgParticlesModule,
    MatMenuModule,
    MatGridListModule,
    MatDialogModule,
    FormsModule,
    ReactiveFormsModule,
    MatFormFieldModule,
    MatInputModule,
    MatRadioModule,
    TechModule,
    SharedModule
  ],
  providers: [
    TravelHelperService,
    {
      provide: MatDialogRef,
       useValue: {}
    },
    {
      provide:MAT_DIALOG_DATA,
      useValue:{}
    }
  ],
  exports: [
    HomeComponent,
    ContactComponent,
    CarouselComponent
  ],
  entryComponents: [
    ImageModalComponent,
    CarouselComponent
  ]
})
export class ComponentsModule { }
