import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { TravelLogComponent } from './components/travel-log/travel-log.component';
import { HomeComponent } from './components/home/home.component';
import { AdminRoutingModule } from './admin/admin-routing.module';
import { StuffComponent } from './components/stuff/stuff.component';
import { TutoringComponent } from './components/tutoring/tutoring.component';
import { FreelanceComponent } from './components/freelance/freelance.component';

const routes: Routes = [
  { path: '', component: HomeComponent},
  { path: 'travel', component: TravelLogComponent },
  { path: 'stuff', component: StuffComponent},
  { path: 'tutoring', component: TutoringComponent},
  { path: 'freelance', component: FreelanceComponent},
  {
    path: 'tech',
    loadChildren: () => import('./components/tech/tech.module')
        .then(m => m.TechModule),
}
];

@NgModule({
  imports: [RouterModule.forRoot(routes), AdminRoutingModule],
  exports: [RouterModule]
})
export class AppRoutingModule { }
