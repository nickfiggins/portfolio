import { Component, OnInit } from '@angular/core';
import { EnumService } from '../shared/services/Enum.service';

@Component({
  selector: 'app-tutoring',
  templateUrl: './tutoring.component.html',
  styleUrls: ['./tutoring.component.scss']
})
export class TutoringComponent implements OnInit {

  particlesOptions = {};

  constructor(public enums : EnumService) {
    this.particlesOptions = enums.particlesOptions;
   }

  ngOnInit(): void {
  }

}
