import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { BlurbsComponent } from './blurbs/blurbs.component';
import { ProjectsComponent } from './projects/projects.component';


const routes: Routes = [
  {
      path: '',
      children: [
          {
              path: 'projects',
              component: ProjectsComponent
          },
          {
              path: 'blurbs',
              component: BlurbsComponent
          }
      ]
  }
];

@NgModule({
    imports: [
        RouterModule.forChild(routes)
    ],
    exports: [
        RouterModule
    ]
})
export class TechRoutingModule { }