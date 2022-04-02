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

  readingCircleContent = {
    title: "currently reading...",
    text: "why we sleep by matthew walker",
    link: "https://www.goodreads.com/book/show/34466963-why-we-sleep"
  }

  interestedCircleContent = {
    title: "currently interested in...",
    text: "solidity/ethereum development",
    link: '',
  }

  blogCircleContent = {
    title: "blogging...",
    text: "eventually",
    link: '',
  }

  learningCircleContent = {
    title: "currently learning...",
    text: "a lot",
    link: ''
  }
  

  circles: Circle[] = [
    //{top: "0%", left: "7.5%", content: this.readingCircleContent},
    //{top: "-10%", left: "45%", content: this.interestedCircleContent},
    //{top: "-10%", left: "25%", content: this.learningCircleContent},
    //{top: "-40%", left: "75%", content: this.blogCircleContent},
  ];

  constructor(public enums: EnumService) {
    this.particlesOptions = enums.particlesOptions
   }

  ngOnInit(): void {
  }

}
