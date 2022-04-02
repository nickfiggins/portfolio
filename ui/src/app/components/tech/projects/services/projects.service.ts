import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Project } from '../models/Project.model';
import { ProjectImage } from '../models/ProjectImage.model';

@Injectable({
  providedIn: 'root'
})
export class ProjectsService {

  constructor(private http : HttpClient) { }

  getAllProjects(){
    return this.http.get<Project[]>('/api/projects');
  }

  getProjectImagesByProjectId(id: string){
    return this.http.get<ProjectImage[]>(`/api/projects/images/${id}`);
  }

  getAllProjectImages(){
    return this.http.get(`/api/projects/images`);
  }
}
