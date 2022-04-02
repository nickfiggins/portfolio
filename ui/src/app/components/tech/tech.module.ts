import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ProjectsComponent } from './projects/projects.component';
import { TechRoutingModule } from './tech-routing.module';
import { BlurbsComponent } from './blurbs/blurbs.component';
import { ComponentsModule } from '../components.module';
import { NgParticlesModule } from 'ng-particles';
import { MatCardModule } from '@angular/material/card';
import { FlexLayoutModule } from '@angular/flex-layout';
import { SharedModule } from '../shared/shared.module';
import { MatButtonModule } from '@angular/material/button';


@NgModule({
  declarations: [ProjectsComponent, BlurbsComponent],
  imports: [
    CommonModule,
    TechRoutingModule,
    NgParticlesModule,
    MatCardModule,
    FlexLayoutModule,
    SharedModule,
    MatButtonModule
  ]
})
export class TechModule { }
