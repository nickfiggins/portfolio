import { Component, OnInit } from '@angular/core';
import { EnumService } from '../../shared/services/Enum.service';
import { ProjectsService } from './services/projects.service';
import { Input } from '@angular/core';
import {CarouselComponent} from '../../shared/carousel/carousel.component';
import { Project } from './models/Project.model';
import { ProjectImage } from './models/ProjectImage.model';

const IMAGE_DIRECTORY = 'assets/images/';

@Component({
  selector: 'app-projects',
  templateUrl: './projects.component.html',
  styleUrls: ['./projects.component.scss']
})
export class ProjectsComponent implements OnInit {

  particlesOptions = {};
  projects : Project[] = [];
  mobile: boolean = false;
  projectImages : any = {};
  constructor(private enums : EnumService, private projectsService : ProjectsService) { 
    this.particlesOptions = enums.particlesOptions;
  }

  ngOnInit(): void {
    this.projectsService.getAllProjects().toPromise().then((response) => {
      this.projects = response;
    }).then(() => {
      this.projects = this.projects.sort((projectA, projectB) => (projectA.id - projectB.id));
      for(let project of this.projects){
            let project_images: ProjectImage[] = [];
            for(let image of project.images){
              project_images.push(new ProjectImage(image, IMAGE_DIRECTORY + "/" + project.name + "/" + image))
            }
            this.projectImages[project.id] = project_images;
          }
          console.log(this.projects)
        });
  
    if (window.screen.width <= 400) { // 768px portrait
      this.mobile = true;
    }
  }

  onResize(event: any) {
    this.mobile = (event.target.innerWidth <= 400) ? true : false;
  }

}
