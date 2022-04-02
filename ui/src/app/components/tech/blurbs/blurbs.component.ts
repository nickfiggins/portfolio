import { Component, OnInit } from '@angular/core';
import { EnumService } from '../../shared/services/Enum.service';

@Component({
  selector: 'app-blurbs',
  templateUrl: './blurbs.component.html',
  styleUrls: ['./blurbs.component.scss']
})
export class BlurbsComponent implements OnInit {

  particlesOptions = {};
  constructor(enums : EnumService) { 
    this.particlesOptions = enums.particlesOptions;
  }

  ngOnInit(): void {
  }

}
