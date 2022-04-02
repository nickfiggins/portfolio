import { Component, OnInit } from '@angular/core';
import {EnumService } from '../shared/services/Enum.service';

export interface CircleContent {
  title: string,
  text: string,
  link: string
}

export interface Circle {
  top: string,
  left: string,
  content: CircleContent,
}

@Component({
  selector: 'app-stuff',
  templateUrl: './stuff.component.html',
  styleUrls: ['./stuff.component.scss']
})
export class StuffComponent implements OnInit {

  particlesOptions = {};
  id = "ts-particles";
  

  circles: Circle[] = [
  ];

  constructor(public enums: EnumService) {
    this.particlesOptions = enums.particlesOptions
   }

  ngOnInit(): void {
  }

}
